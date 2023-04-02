// SPDX-License-Identifier: MIT
pragma solidity  ^0.8.0;

import "./RegistryContract.sol";
import "./crypto/Schnorr.sol";

contract OracleContract {
    struct EnrollNode {
        address addr;
        uint256 index;
    }

    // 当前是否有请求验证的标志
    bool private isValidateTime = false;

    // 请求验证需要的钱
    uint256 public constant BASE_FEE = 0.001 ether;

    // 验证奖励
    uint256 public constant VALIDATOR_FEE = 0.0001 ether;
    uint256 public constant TOTAL_FEE = BASE_FEE + VALIDATOR_FEE * 100;



    bool private isUrgeTime = false; // 标志当前是否正在催促验证器;
    uint private urgeBeginTime;
    uint public constant waitTime = 30;
    mapping(address => uint8) private lazyValidatorIndex;
    address[] private LazyValidators; // 聚合器提交的不工作的验证者地址;
    uint256 public constant URGE_FEE = 0.0001 ether;  // 每次催促的开销;
    uint256 public constant COMPENSATION_FEE = 0.001 ether;  // 无法提交结果时候的惩罚;


    // 保存验证结果的映射；
    mapping(bytes32 => bool) private blockValidationResults;
    mapping(bytes32 => bool) private txValidationResults;

    // 记录报名节点的映射；
    mapping(address => EnrollNode) private enrollNodes;
    address[] private enrollNodeIndices;


    // 验证类型的枚举：未知，区块存在验证，交易存在验证;
    enum ValidationType { UNKNOWN, BLOCK, TRANSACTION }

    // indexed属性是为了方便在日志结构中查找，这个是一个事件，会存储到交易的日志中，就是类似于挖矿上链
    event ValidationRequest(ValidationType typ, address indexed from, bytes32 hash);

    event urgeEvent(uint beginTime, address[] lazyValidators);

    event ValidationBegin();

    event ValidationResponse(
        ValidationType typ,
        address indexed aggregator,   
        bytes32 message,
        bool valid,
        uint256 fee
    );

    RegistryContract private registryContract;
    constructor(address _registryContract) {
        registryContract = RegistryContract(_registryContract);
    }

    modifier minFee(uint _min) {
        require(msg.value >= _min, "too few fee amount");
        _;
    }

    function validateBlock(bytes32 _hash) external payable minFee(TOTAL_FEE) {
        require(!isValidateTime, "Another validate is in progress!");
        isValidateTime = true;
        emit ValidationRequest(ValidationType.BLOCK, msg.sender, _hash);
    }

    function validateTransaction(bytes32 _hash) external payable minFee(TOTAL_FEE) {
        require(!isValidateTime, "Another validate is in progress!");
        isValidateTime = true;
        emit ValidationRequest(ValidationType.TRANSACTION, msg.sender, _hash);
    }

    function submitBlockValidationResult(bool _result, bytes32 message, uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash) external {
        // require(isValidateTime, "Not validate time!");
        submitValidationResult(ValidationType.BLOCK, _result, message, signature, pubKeyX, pubKeyY, RX, RY, _hash);
        isValidateTime = false;
    }

    function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash) external {
        // require(isValidateTime, "Not validate time!");
        submitValidationResult(ValidationType.TRANSACTION, _result, message, signature, pubKeyX, pubKeyY, RX, RY, _hash);
        isValidateTime = false;
    }


    function submitValidationResult(
        ValidationType _typ,
        bool _result,
        bytes32 message,
        uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash
    ) private {

        require(_typ != ValidationType.UNKNOWN, "unknown validation type");

        // RegistryContract.OracleNode memory aggregator = registryContract.getAggregator();
        // require(aggregator.addr == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器

        /******************
         *Schnorr签名的验证*
         ******************/

        require(Schnorr.verify(signature, pubKeyX, pubKeyY, RX, RY, _hash), "signature: address does not match");
        delete enrollNodeIndices;

        // 给当前合约的调用者（聚合器）转账 
        // payable(msg.sender).transfer(BASE_FEE);     //此处完成给聚合器的报酬转账

        // 给所有的参与验证的验证器节点转账
        for(uint32 i = 0 ; i < enrollNodeIndices.length ; i++){
            payable(enrollNodeIndices[i]).transfer(VALIDATOR_FEE);
        }
        
        if (_typ == ValidationType.BLOCK) {
            blockValidationResults[message] = _result;
        } else if (_typ == ValidationType.TRANSACTION) {
            txValidationResults[message] = _result;
        }

        emit ValidationResponse(_typ, msg.sender, message, _result, BASE_FEE + VALIDATOR_FEE * enrollNodeIndices.length );  // 通知链下，聚合奖励 + 验证奖励 * 验证器报名个数
    }

    // 报名
    function EnrollOracleNode() external payable{
        require(enrollNodeIndices.length <= registryContract.countOracleNodes()/2, "Enrollment closed!");
        require(registryContract.oracleNodeIsRegistered(msg.sender) , "The Oracle doesn't registered");
        require(!oracleNodeIsEnroll(msg.sender), "already enrolled");
        EnrollNode storage iopNode = enrollNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.index = enrollNodeIndices.length;
        enrollNodeIndices.push(iopNode.addr);
        if(enrollNodeIndices.length == registryContract.countOracleNodes()/2 + 1){
            emit ValidationBegin();
        }
    }
    // 是否报名
    function oracleNodeIsEnroll(address _addr) public view returns (bool){
        if (enrollNodeIndices.length == 0) return false;
        return (enrollNodeIndices[enrollNodes[_addr].index] == _addr);
    }

    // 由聚合器调用，用于催促未提交结果或参与交互的验证器；
    function urge(address[] calldata lazyValidators) public payable {
        RegistryContract.OracleNode memory aggregator = registryContract.getAggregator();
        require(aggregator.addr == msg.sender, "Caller is not the aggregator");  //判断当前合约的调用者是不是聚合器
        require(isValidateTime, "Not validate time!"); // 当前是不是验证时间；
        require(!isUrgeTime, "Already urge lazyValidators"); // 是否已经催促过，不允许重复调用;
        isUrgeTime = true;
        urgeBeginTime = getBlockTime();
        for(uint8 i = 0 ; i < lazyValidators.length ; i++){
            uint8 index = lazyValidatorIndex[lazyValidators[i]];
            index = i;
            LazyValidators.push(lazyValidators[i]);
        }
        emit urgeEvent(urgeBeginTime, LazyValidators);
    }

    // 审判超时还未提交的验证器节点，只能由聚合器调用，并且当前区块时间需要达到超时条件才会惩罚验证器;
    function judge() public {
        RegistryContract.OracleNode memory aggregator = registryContract.getAggregator();
        require(aggregator.addr == msg.sender, "Caller is not the aggregator");  //判断当前合约的调用者是不是聚合器
        require(isUrgeTime, "Haven't urge lazyValidator");
        require(urgeBeginTime + waitTime < getBlockTime(), "Be patient");
        // TODO: 惩罚LazyValidators中的验证器；
        delete(LazyValidators);
    }

    // 被催促的验证器提交证明的入口，提交证明并且通过之后，可以将自己移除出待惩罚的节点地址数组：
    function lazySubmit(uint256 provement) public {
        require(LazyValidators[lazyValidatorIndex[msg.sender]] == msg.sender, "Not be urged!");
        // TODO: 提交证明，使自己从LazyValidators中移除；
        LazyValidators[lazyValidatorIndex[msg.sender]] = 0x0000000000000000000000000000000000000000; // 把自己的地址修改为黑洞地址，惩罚时会忽略这个地址;
    }

    function getBlockTime() public view returns (uint) {
        return block.timestamp;
    }

    // 根据索引查找
    function findEnrollNodeByIndex(uint256 _index)
        public
        view
        returns (EnrollNode memory)
    {
        require(_index >= 0 && _index < enrollNodeIndices.length, "not found");
        return enrollNodes[enrollNodeIndices[_index]];
    }
    // 返回报名总数
    function countEnrollNodes() external view returns (uint256){
        return enrollNodeIndices.length;
    }

    function findBlockValidationResult(bytes32 _hash) public view returns (bool) {
        return blockValidationResults[_hash];
    }

    function findTransactionValidationResult(bytes32 _hash) public view returns (bool) {
        return txValidationResults[_hash];
    }
}
