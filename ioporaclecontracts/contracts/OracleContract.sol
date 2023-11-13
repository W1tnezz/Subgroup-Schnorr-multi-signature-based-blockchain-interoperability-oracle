// SPDX-License-Identifier: MIT
pragma solidity  ^0.8.0;

import "./RegistryContract.sol";
import "./crypto/Schnorr.sol";

contract OracleContract {

    // 当前是否有请求验证的标志
    bool private isValidateTime = false;

    // 请求验证需要的钱
    uint256 public constant BASE_FEE = 0.1 ether;

    // 验证奖励
    uint256 public constant VALIDATOR_FEE = 0.01 ether;

    // 质询费用
    uint256 public constant INQUIRY_FEE = 0.01 ether;

    uint256 public constant RE_PROVIDE_FEE = 0.05 ether;


    uint256 public subgroupSize;
    
    address private aggregatorAddr;

    bool private isInquiryTime = false; // 标志当前是否正在催促验证器;
    uint private inquiryBeginTime;

    uint public constant WAIT_TIME = 10;
    mapping(address => uint) private lazyValidatorIndex;
    address[] private lazyValidators; // 聚合器提交的lazy验证者地址;

    mapping(address => uint) private dishonestValidatorIndex;
    address[] private dishonestValidators; // 聚合器提交的lazy验证者地址;


    // 保存验证结果的映射；
    mapping(bytes32 => bool) private blockValidationResults;
    mapping(bytes32 => bool) private txValidationResults;

    // 记录报名节点的地址；
    address[] private validators;

    // 验证类型的枚举：未知，区块存在验证，交易存在验证;
    enum ValidationType { UNKNOWN, BLOCK, TRANSACTION }

    enum InquiryType { LAZY, DISHONEST }

    // indexed属性是为了方便在日志结构中查找，这个是一个事件，会存储到交易的日志中，就是类似于挖矿上链
    event ValidationRequest(ValidationType typ, address indexed from, bytes32 hash, uint size);

    event InquiryEvent(InquiryType typ, uint beginTime, address[] lazyValidators);

    event Explain(string exp, address addr);

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

    modifier minFee(uint _size) {
        require(msg.value >= BASE_FEE + VALIDATOR_FEE * _size, "too few fee amount");
        _;
    }

    function validateBlock(bytes32 _hash, uint size) external payable minFee(size) {
        require(!isValidateTime, "Another validate is in progress!");
        isValidateTime = true;
        subgroupSize = size;
        RegistryContract.OracleNode memory aggregator = registryContract.getAggregator();
        aggregatorAddr = aggregator.addr;
        emit ValidationRequest(ValidationType.BLOCK, msg.sender, _hash, size);
    }

    function validateTransaction(bytes32 _hash, uint size) external payable minFee(size) {
        require(!isValidateTime, "Another validate is in progress!");
        isValidateTime = true;
        subgroupSize = size;
        RegistryContract.OracleNode memory aggregator = registryContract.getAggregator();
        aggregatorAddr = aggregator.addr;
        emit ValidationRequest(ValidationType.TRANSACTION, msg.sender, _hash, size);
    }

    function submitBlockValidationResult(bool _result, bytes32 message, uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash) external {
        require(isValidateTime, "Not validate time!");
        submitValidationResult(ValidationType.BLOCK, _result, message, signature, pubKeyX, pubKeyY, RX, RY, _hash);
        isValidateTime = false;
    }

    function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash) external {
        require(isValidateTime, "Not validate time!");
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
        require(aggregatorAddr == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器
        /*Schnorr签名的验证*/
        require(Schnorr.verify(signature, pubKeyX, pubKeyY, RX, RY, _hash), "sig: address doesn't match");

        if (_typ == ValidationType.BLOCK) {
            blockValidationResults[message] = _result;
        } else if (_typ == ValidationType.TRANSACTION) {
            txValidationResults[message] = _result;
        }

        // 给当前合约的调用者（聚合器）转账 
        payable(msg.sender).transfer(BASE_FEE);     //此处完成给聚合器的报酬转账
        // 给所有的参与验证的验证器节点转账
        for(uint32 i = 0 ; i < validators.length ; i++){
            payable(validators[i]).transfer(VALIDATOR_FEE);
        }

        delete validators;
        emit ValidationResponse(_typ, msg.sender, message, _result, BASE_FEE + VALIDATOR_FEE * validators.length );  // 通知链下，聚合奖励 + 验证奖励 * 验证器报名个数
    }

    // 报名
    function enroll() external {
        require(validators.length < subgroupSize, "Enrollment closed!");
        require(registryContract.oracleNodeIsRegistered(msg.sender) , "The Oracle doesn't registered");
        require(!oracleNodeIsEnroll(msg.sender), "already enrolled");
        validators.push(msg.sender);
        if(validators.length == subgroupSize){
            emit ValidationBegin();
        }
    }
    // 是否报名
    function oracleNodeIsEnroll(address _addr) public view returns (bool){
        if (validators.length == 0) return false;
        for(uint32 i = 0 ; i < validators.length ; i++){
            if(validators[i] == _addr){
                return true;
            }
        }
        return false;
    }


    // 由聚合器调用，用于质询可能在偷懒的验证器以及未提交结果或提交错误单签名的验证器；
    function inquiry(address[] calldata lazys, address[] calldata dishonests) external payable {
        require(aggregatorAddr == msg.sender, "not the aggregator"); //判断当前合约的调用者是不是聚合器
        require(isValidateTime, "Not validate time!"); // 当前是不是验证时间；
        require(!isInquiryTime, "Already inquiry lazyValidators"); // 是否已经质询过，不允许重复调用;
        require(msg.value >= INQUIRY_FEE, "Inquiry fee not enough");
        isInquiryTime = true;
        inquiryBeginTime = getBlockTime();
        for(uint8 i = 0 ; i < lazys.length ; i++){
            lazyValidatorIndex[lazys[i]] = i;
            lazyValidators.push(lazys[i]);
        }
        if (lazys.length != 0){
            emit InquiryEvent(InquiryType.LAZY, inquiryBeginTime, lazyValidators);
        }
        
        for(uint8 i = 0 ; i < dishonests.length ; i++){
            dishonestValidatorIndex[dishonests[i]] = i;
            dishonestValidators.push(dishonests[i]);
        }
        if(dishonests.length != 0){
            emit InquiryEvent(InquiryType.DISHONEST, inquiryBeginTime, dishonestValidators);
        }
    }

    // 审判超时还未提交的验证器节点，只能由聚合器调用，并且当前区块时间需要达到超时条件才会惩罚验证器;
    function punish() external {
        require(aggregatorAddr == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器
        require(isInquiryTime, "Haven't urge lazyValidator");
        require(inquiryBeginTime + WAIT_TIME < getBlockTime(), "Be patient");
        for(uint i = 0 ; i < lazyValidators.length ; i++){
            if(lazyValidators[i] != 0x0000000000000000000000000000000000000000){
                payable(msg.sender).transfer(registryContract.findOracleNodeByAddress(lazyValidators[i]).stake);
                registryContract.deleteNode(lazyValidators[i]);
            }
        }
        delete(lazyValidators);

        for(uint i = 0 ; i < dishonestValidators.length ; i++){
            if(dishonestValidators[i] != 0x0000000000000000000000000000000000000000){
                payable(msg.sender).transfer(registryContract.findOracleNodeByAddress(dishonestValidators[i]).stake);
                registryContract.deleteNode(dishonestValidators[i]);
            }
        }
        delete(dishonestValidators);
    }

    // 被催促的验证器提交证明的入口，提交证明并且通过之后，可以将自己移除出待惩罚的节点地址数组：
    function reProvide(bytes32 targetBlock, string calldata exp) external payable{
        require(lazyValidators[lazyValidatorIndex[msg.sender]] == msg.sender, "Not be inquiry!");
        require(msg.value >= RE_PROVIDE_FEE, "Re-provide fee not enough!");
        require(targetBlock != 0 || bytes(exp).length != 0, "Need target block or explain!");
        if(targetBlock != 0){
            lazyValidators[lazyValidatorIndex[msg.sender]] = 0x0000000000000000000000000000000000000000; // 把自己的地址修改为黑洞地址，惩罚时会忽略这个地址;
        } else {
            emit Explain(exp, msg.sender);
        }
    }

    function cancelPunishment(address addr) external {
        require(aggregatorAddr == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器
        if(addr == lazyValidators[lazyValidatorIndex[addr]]){
            lazyValidators[lazyValidatorIndex[addr]] = 0x0000000000000000000000000000000000000000;
        } else if(addr == dishonestValidators[dishonestValidatorIndex[addr]]){
            dishonestValidators[dishonestValidatorIndex[addr]] = 0x0000000000000000000000000000000000000000;
        }
    }

    function getBlockTime() public view returns (uint) {
        return block.timestamp;
    }

    // 根据索引查找
    function findEnrollNodeByIndex(uint256 _index)
        public
        view
        returns (address)
    {
        require(_index >= 0 && _index < validators.length, "not found");
        return validators[_index];
    }
    // 返回报名总数
    function countEnrollNodes() external view returns (uint256){
        return validators.length;
    }

    function findBlockValidationResult(bytes32 _hash) public view returns (bool) {
        return blockValidationResults[_hash];
    }

    function findTransactionValidationResult(bytes32 _hash) public view returns (bool) {
        return txValidationResults[_hash];
    }
}
