package test

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/random"
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
	"math/big"
)

type TestAccount struct {
	UnsafeOracleNodeServer
	suite           pairing.Suite
	targetEthClient *ethclient.Client
	oracleContract  *OracleContractWrapper
	ecdsaPrivateKey *ecdsa.PrivateKey
	account         common.Address
	chainId         *big.Int
}

type Point struct {
	X *big.Int
	Y *big.Int
}

func NewTestAccount(c Config) (*TestAccount, error) {
	// 创建一个连接以太坊的客户端，TargetAddress是以太坊的目标地址
	targetEthClient, err := ethclient.Dial(c.Ethereum.TargetAddress)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}
	// 这个也是连接以太坊的连接客户端

	// 区块链的ID
	chainId := big.NewInt(c.Ethereum.ChainID)

	oracleContract, err := NewOracleContract(common.HexToAddress(c.Contracts.OracleContractAddress), targetEthClient)
	oracleContractWrapper := &OracleContractWrapper{
		OracleContract: oracleContract,
	}
	if err != nil {
		return nil, fmt.Errorf("oracle contract: %v", err)
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(c.Ethereum.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa: %v", err)
	}

	hexAddress, err := AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	suite := bn256.NewSuiteG1()
	node := &TestAccount{
		suite:           suite,
		targetEthClient: targetEthClient,
		oracleContract:  oracleContractWrapper,
		ecdsaPrivateKey: ecdsaPrivateKey,
		account:         account,
		chainId:         chainId,
	}
	return node, nil
}

func (n *TestAccount) Run() error {
	// 打印信息：
	fmt.Println("节点信息：")
	fmt.Print("节点以太坊账户：")
	fmt.Println(n.account)

	// 调用合约的权限对象
	auth, err := bind.NewKeyedTransactorWithChainID(n.ecdsaPrivateKey, n.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}

	// schnorr公私钥对，私钥为随机标量，pubkey = privateKey * G
	privateKey := n.suite.G1().Scalar().Pick(random.New())
	pubkey := n.suite.G1().Point().Mul(privateKey, nil)
	pubkeyBytes, _ := pubkey.MarshalBinary()

	// r为随机数，R = r * G
	r := n.suite.G1().Scalar().Pick(random.New())
	R := n.suite.G1().Point().Mul(r, nil)
	RBytes, _ := R.MarshalBinary()

	// 随机消息_message，哈希为byte[32]数组 --> message
	_message := "1256901778428453331b853610968234rw543f54da4c331e2"
	message := crypto.Keccak256Hash([]byte(_message))

	// e为随机消息映射到曲线上的标量,即_message --> message --> e, _hash为e的big.Int形式
	e := n.suite.G1().Scalar().SetBytes(message.Bytes())
	_hash := new(big.Int)
	_hash.SetString(e.String(), 16)
	fmt.Println(_hash)

	// tmp = e * privateKey, 签名_s = r + tmp = r + e * privateKey, s为_s的big.Int形式
	tmp := n.suite.G1().Scalar().Mul(e, privateKey)
	_s := n.suite.G1().Scalar().Add(r, tmp)
	_sBytes, _ := _s.MarshalBinary()
	s := new(big.Int)
	s.SetBytes(_sBytes)
	fmt.Println(s)

	_, err = n.oracleContract.SubmitTransactionValidationResult(auth, true, message, s, new(big.Int).SetBytes(pubkeyBytes[:32]), new(big.Int).SetBytes(pubkeyBytes[32:64]), new(big.Int).SetBytes(RBytes[:32]), new(big.Int).SetBytes(RBytes[32:64]), _hash)
	if err != nil {
		return fmt.Errorf("submit error: %w", err)
	}

	res, err := n.oracleContract.GetBlockTime(nil)
	fmt.Println(res)
	if err != nil {
		return fmt.Errorf("submit error: %w", err)
	}
	return nil
}

func (n *TestAccount) Stop() {
	n.targetEthClient.Close()
}
