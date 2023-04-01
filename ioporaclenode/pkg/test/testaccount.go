package test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type TestAccount struct {
	UnsafeOracleNodeServer
	suite           *secp256k1.BitCurve
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

	suite := secp256k1.S256()
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

	// 待签名消息的hash值，格式为32位Bytes数组
	str := "1256901778428453331bd44b1619e05350b853610968f54da7338e4c331acbe2"
	message := crypto.Keccak256Hash([]byte(str))

	// 测试用的ECDSA公私钥，当作是最后聚合的公私钥
	testECDSAKey, _ := crypto.GenerateKey()
	testPrivateKey := testECDSAKey.D
	testPublicKey := testECDSAKey.PublicKey
	fmt.Print("测试私钥：")
	fmt.Println(testPrivateKey)
	fmt.Print("测试公钥：")
	fmt.Println(testPublicKey)

	r := make([]byte, 32)
	rand.Read(r)
	R := new(Point)
	R.X, R.Y = n.suite.ScalarBaseMult(r)

	RBytes := n.suite.Marshal(R.X, R.Y)
	PublicKeyBytes := n.suite.Marshal(testPublicKey.X, testPublicKey.Y)
	h := make([][]byte, 3)
	h[0] = RBytes
	h[1] = PublicKeyBytes
	h[2] = []byte(str)
	hash := sha256.New()
	e := hash.Sum(bytes.Join(h, []byte("")))[:32]
	_hash := new(big.Int)
	_hash.SetBytes(e)

	s := big.NewInt(1)
	s.Mul(testPrivateKey, new(big.Int).SetBytes(e)).Add(s, new(big.Int).SetBytes(r))
	s.Mod(s, n.suite.N)

	tmp := new(Point)
	tmp.X, tmp.Y = n.suite.ScalarMult(testPublicKey.X, testPublicKey.Y, e)

	/*	S1 := new(Point)
		S1.X, S1.Y = n.suite.Add(R.X, R.Y, tmp.X, tmp.Y)
		S2 := new(Point)
		S2.X, S2.Y = n.suite.ScalarBaseMult(s.Bytes())
		fmt.Println(S1.x, S2.x)*/

	_, err = n.oracleContract.SubmitTransactionValidationResult(auth, true, message, s, testPublicKey.X, testPublicKey.Y, R.X, R.Y, _hash)
	if err != nil {
		return fmt.Errorf("submit error: %w", err)
	}
	return nil
}

func (n *TestAccount) Stop() {
	n.targetEthClient.Close()
}
