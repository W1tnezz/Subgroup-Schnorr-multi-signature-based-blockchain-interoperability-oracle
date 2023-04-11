package test

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type TestAccount struct {
	UnsafeOracleNodeServer
	targetEthClient *ethclient.Client
	oracleContract  *OracleContractWrapper
	ecdsaPrivateKey *ecdsa.PrivateKey
	account         common.Address
	chainId         *big.Int
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

	node := &TestAccount{
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
	hash := crypto.Keccak256Hash([]byte(str))

	// 测试用的ECDSA公私钥，当作是最后聚合的公私钥
	testECDSAKey, _ := crypto.GenerateKey()
	testPrivateKey := testECDSAKey.D
	testPublicKey := testECDSAKey.PublicKey
	fmt.Print("测试私钥：")
	fmt.Println(testPrivateKey)
	fmt.Print("测试公钥：")
	fmt.Println(testPublicKey)

	testAddress := crypto.PubkeyToAddress(testPublicKey)

	sig, _ := crypto.Sign(hash.Bytes(), testECDSAKey)
	fmt.Print("signature:")
	fmt.Println(sig)
	var r [32]byte
	for i := 0; i < 32; i++ {
		r[i] = sig[i]
	}
	fmt.Print("r：")
	fmt.Println(r)
	var s [32]byte
	for i := 0; i < 32; i++ {
		s[i] = sig[i+32]
	}
	fmt.Print("s：")
	fmt.Println(s)

	// 巨坑：以太坊黄皮书更新v从0/1为27/28
	var v = sig[64] + 27
	fmt.Print("v：")
	fmt.Println(v)

	_, err = n.oracleContract.SubmitTransactionValidationResult(auth, true, testAddress, v, r, s, hash)
	if err != nil {
		return fmt.Errorf("submit error: %w", err)
	}
	return nil
}

func (n *TestAccount) Stop() {
	n.targetEthClient.Close()
}
