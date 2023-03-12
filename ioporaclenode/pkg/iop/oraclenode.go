package iop

/*
1、修改新的OracleNode结构，不再需要BLS；
2、不再需要DKG，但是参照DKG的实现方法实现ECDSA聚合通信，其实也就是几轮广播；
3、
*/
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
	"math/big"
	"net"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	iota "github.com/iotaledger/iota.go/v2"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3/suites"
	"google.golang.org/grpc"
)

type OracleNode struct {
	UnsafeOracleNodeServer
	server            *grpc.Server
	serverLis         net.Listener
	targetEthClient   *ethclient.Client
	sourceEthClient   *ethclient.Client
	registryContract  *RegistryContractWrapper
	oracleContract    *OracleContractWrapper
	suite             suites.Suite
	ecdsaPrivateKey   *ecdsa.PrivateKey
	account           common.Address
	dkg               *DistKeyGenerator
	connectionManager *ConnectionManager
	validator         *Validator
	aggregator        *Aggregator
	chainId           *big.Int
}

func NewOracleNode(c Config) (*OracleNode, error) {
	server := grpc.NewServer()
	serverLis, err := net.Listen("tcp", c.BindAddress)
	if err != nil {
		return nil, fmt.Errorf("listen on %s: %v", c.BindAddress, err)
	}
	// 创建一个连接以太坊的客户端，TargetAddress是以太坊的目标地址
	targetEthClient, err := ethclient.Dial(c.Ethereum.TargetAddress)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}
	// 这个也是连接以太坊的连接客户端
	sourceEthClient, err := ethclient.Dial(c.Ethereum.SourceAddress)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}
	// 区块链的ID
	chainId := big.NewInt(c.Ethereum.ChainID)

	// 创建新的iota客户端
	iotaAPI := iota.NewNodeHTTPAPIClient(c.IOTA.Rest)
	if err != nil {
		return nil, fmt.Errorf("iota client: %v", err)
	}

	// 配置mqtt,创建公共服务器
	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.IOTA.Mqtt)
	opts.SetClientID(c.BindAddress)
	mqttClient := mqtt.NewClient(opts)
	mqttTopic := []byte(c.IOTA.Topic)

	// 注册
	registryContract, err := NewRegistryContract(common.HexToAddress(c.Contracts.RegistryContractAddress), targetEthClient)
	if err != nil {
		return nil, fmt.Errorf("registry contract: %v", err)
	}

	registryContractWrapper := &RegistryContractWrapper{
		RegistryContract: registryContract,
	}

	oracleContract, err := NewOracleContract(common.HexToAddress(c.Contracts.OracleContractAddress), targetEthClient)
	oracleContractWrapper := &OracleContractWrapper{
		OracleContract: oracleContract,
	}
	if err != nil {
		return nil, fmt.Errorf("oracle contract: %v", err)
	}

	suite := bn256.NewSuiteG2()

	ecdsaPrivateKey, err := crypto.HexToECDSA(c.Ethereum.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa: %v", err)
	}

	hexAddress, err := AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	connectionManager := NewConnectionManager(registryContractWrapper, account)
	validator := NewValidator(suite, oracleContract, sourceEthClient)
	aggregator := NewAggregator(
		suite,
		targetEthClient,
		connectionManager,
		oracleContractWrapper,
		registryContractWrapper,
		account,
		ecdsaPrivateKey,
		chainId,
	)
	dkg := NewDistKeyGenerator(
		connectionManager,
		aggregator,
		mqttClient,
		mqttTopic,
		iotaAPI,
		registryContractWrapper,
		ecdsaPrivateKey,
		account,
		chainId,
	)
	validator.SetDistKeyGenerator(dkg)
	aggregator.SetDistKeyGenerator(dkg)

	node := &OracleNode{
		server:            server,
		serverLis:         serverLis,
		targetEthClient:   targetEthClient,
		sourceEthClient:   sourceEthClient,
		registryContract:  registryContractWrapper,
		oracleContract:    oracleContractWrapper,
		suite:             suite,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		account:           account,
		dkg:               dkg,
		connectionManager: connectionManager,
		validator:         validator,
		aggregator:        aggregator,
		chainId:           chainId,
	}
	RegisterOracleNodeServer(server, node)

	return node, nil
}

func (n *OracleNode) Run() error {
	// 打印信息：
	fmt.Println("节点信息：")
	fmt.Print("节点以太坊账户：")
	fmt.Println(n.account)

	fmt.Print("节点ECDSA私钥：")
	fmt.Println(n.ecdsaPrivateKey.D)

	fmt.Println("ECDSA公钥：")
	fmt.Print("横坐标：")

	fmt.Println(n.ecdsaPrivateKey.X.Bytes())
	fmt.Print("纵坐标：")
	fmt.Println(n.ecdsaPrivateKey.Y)

	// 创建连接
	if err := n.connectionManager.InitConnections(); err != nil {
		return fmt.Errorf("init connections: %w", err)
	}

	/*启动3个协程*/

	// 协程1：监听并处理DKG event，这里应该不需要了；
	go func() {
		if err := n.dkg.ListenAndProcess(context.Background()); err != nil {
			log.Errorf("Watch and handle DKG log: %v", err)
		}
	}()

	// 协程2：监听并处理新注册节点event，主要负责与新注册的节点建立连接，这里应该不用动；
	go func() {
		if err := n.connectionManager.WatchAndHandleRegisterOracleNodeLog(context.Background()); err != nil {
			log.Errorf("Watch and handle register oracle node log: %v", err)
		}
	}()

	// 协程3：监听并处理验证请求event
	go func() {
		if err := n.aggregator.WatchAndHandleValidationRequestsLog(context.Background()); err != nil {
			log.Errorf("Watch and handle ValidationRequest log: %v", err)
		}
	}()

	// 注册链上身份
	if err := n.register(n.serverLis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return n.server.Serve(n.serverLis)
}

func (n *OracleNode) register(ipAddr string) error {
	isRegistered, err := n.registryContract.OracleNodeIsRegistered(nil, n.account)
	if err != nil {
		return fmt.Errorf("is registered: %w", err)
	}

	minStake, err := n.registryContract.MINSTAKE(nil)
	if err != nil {
		return fmt.Errorf("min stake: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(n.ecdsaPrivateKey, n.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}
	auth.Value = minStake

	if !isRegistered {
		_, err = n.registryContract.RegisterOracleNode(auth, ipAddr, n.ecdsaPrivateKey.X.Bytes())
		if err != nil {
			return fmt.Errorf("register iop node: %w", err)
		}
	}
	return nil
}

func (n *OracleNode) Stop() {
	n.server.Stop()
	n.targetEthClient.Close()
	n.sourceEthClient.Close()
	n.connectionManager.Close()
}
