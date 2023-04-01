// SPDX-License-Identifier: MIT
pragma solidity  ^0.8.0;

import "./RegistryContract.sol";
import "./crypto/Schnorr.sol";

contract OracleContract {
    struct EnrollNode {
        address addr;
        uint256 index;
    }

    uint256 public constant BASE_FEE = 0.001 ether;

    // 验证奖励
    uint256 public constant VALIDATOR_FEE = 0.0001 ether;

    //每次挑战的开销
    uint256 public constant CHANLLENGE_FEE = 0.0001 ether;

    // 无法提供证明时候的惩罚
    uint256 public constant COMPENSATION_FEE = 0.001 ether;

    uint256 public constant TOTAL_FEE = BASE_FEE + VALIDATOR_FEE * 100;

    uint256 private requestCounter;
    uint256 private requestsSinceLastPayout;

    mapping(bytes32 => bool) private blockValidationResults;
    mapping(bytes32 => bool) private txValidationResults;
    mapping(address => EnrollNode) private enrollNodes;
    address[] private enrollNodeIndices;


    // 定义了一个报名的事件
    // event enrollOracleNode(address indexed sender);
    enum ValidationType { UNKNOWN, BLOCK, TRANSACTION }

    event ValidationRequest(ValidationType typ, address indexed from, bytes32 hash);
    // indexed属性是为了方便在日志结构中查找，这个是一个事件，会存储到交易的日志中，就是类似于挖矿上链

    event ValidationBegin();

    event ValidationResponse(
        ValidationType typ,
        address indexed aggregator,   
        bytes32 hash,
        bool valid,
        uint256 fee
    );

    RegistryContract private registryContract;
    // Schnorr private schnorr;
    constructor(address _registryContract) {
        registryContract = RegistryContract(_registryContract);
        // schnorr = Schnorr(_schnorr);
    }

    modifier minFee(uint _min) {
        require(msg.value >= _min, "too few fee amount");
        _;
    }

    function validateBlock(bytes32 _hash) external payable minFee(TOTAL_FEE) {
        emit ValidationRequest(ValidationType.BLOCK, msg.sender, _hash);
    }

    function validateTransaction(bytes32 _hash) external payable minFee(TOTAL_FEE) {
        emit ValidationRequest(ValidationType.TRANSACTION, msg.sender, _hash);
    }

    function submitBlockValidationResult(bool _result, bytes32 message, uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash) external {
        submitValidationResult(ValidationType.BLOCK, _result, message, signature, pubKeyX, pubKeyY, RX, RY, _hash);
    }

    function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash) external {
        submitValidationResult(ValidationType.TRANSACTION, _result, message, signature, pubKeyX, pubKeyY, RX, RY, _hash);
    }


    function submitValidationResult(
        ValidationType _typ,
        bool _result,  // 验证结果？
        bytes32 message,
        uint256 signature, uint256 pubKeyX, uint256 pubKeyY, uint256 RX , uint256 RY, uint256 _hash
    ) private {

        require(_typ != ValidationType.UNKNOWN, "unknown validation type");

        // RegistryContract.OracleNode memory aggregator = registryContract.getAggregator();

        // require(aggregator.addr == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器， 因此对于验证器的奖励和惩罚要放在前面

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

        emit ValidationResponse(_typ, msg.sender, message, _result, BASE_FEE + VALIDATOR_FEE * enrollNodeIndices.length );  // 通知连下，聚合奖励 + 验证奖励 * 验证器报名个数
    }

    // 关于聚合器质询扣费，是每质询一次就扣费一次,给合约转账(暂时存在问题)
    function chanllenge() public payable {
        RegistryContract.OracleNode memory aggregator =
            registryContract.getAggregator();
        require(aggregator.stake > CHANLLENGE_FEE , "stake is too low!");
        aggregator.stake = aggregator.stake - CHANLLENGE_FEE;
    }
    // 处罚金
    function fine() public payable {
        // 此时需要验证器给合约转账，作为处罚金(存在问题)
        RegistryContract.OracleNode memory oracle =
            registryContract.findOracleNodeByAddress(msg.sender);
        require(oracle.stake > COMPENSATION_FEE , "stake is too low!");

        oracle.stake = oracle.stake - COMPENSATION_FEE;
        
        // 合约给聚合器转账,该函数是聚合器调用的
        payable(msg.sender).transfer(COMPENSATION_FEE);
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
