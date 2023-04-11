package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go.dedis.ch/kyber/v3/util/random"
	"math/big"
	"net"

	"ioporaclenode/internal/pkg/kyber/pairing/bn256"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	iota "github.com/iotaledger/iota.go/v2"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
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
	schnorrPrivateKey kyber.Scalar
	account           common.Address
	connectionManager *ConnectionManager
	validator         *Validator
	aggregator        *Aggregator
	isAggregator      bool
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

	if err != nil {
		return nil, fmt.Errorf("dist key contract: %v", err)
	}

	suite := bn256.NewSuiteG1()

	ecdsaPrivateKey, err := crypto.HexToECDSA(c.Ethereum.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa: %v", err)
	}
	schnorrPrivateKey := suite.G1().Scalar().Pick(random.New())
	if err != nil {
		return nil, fmt.Errorf("hex to scalar: %v", err)
	}

	hexAddress, err := AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	connectionManager := NewConnectionManager(registryContractWrapper, account)
	validator := NewValidator(
		suite,
		registryContractWrapper,
		oracleContractWrapper,
		ecdsaPrivateKey,
		sourceEthClient,
		connectionManager,
		account,
		mqttClient,
		mqttTopic,
		iotaAPI,
		schnorrPrivateKey,
	)
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

	node := &OracleNode{
		server:            server,
		serverLis:         serverLis,
		targetEthClient:   targetEthClient,
		sourceEthClient:   sourceEthClient,
		registryContract:  registryContractWrapper,
		oracleContract:    oracleContractWrapper,
		suite:             suite,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		schnorrPrivateKey: schnorrPrivateKey,
		account:           account,
		connectionManager: connectionManager,
		validator:         validator,
		aggregator:        aggregator,
		isAggregator:      false,
		chainId:           chainId,
	}
	RegisterOracleNodeServer(server, node)

	return node, nil
}

func (n *OracleNode) Run() error {
	// 创建连接
	if err := n.connectionManager.InitConnections(); err != nil {
		return fmt.Errorf("init connections: %w", err)
	}

	go func() {
		if err := n.validator.ListenAndProcess(n); err != nil {
			log.Errorf("Watch and handle DKG log: %v", err)
		}
	}()

	go func() {
		if err := n.connectionManager.WatchAndHandleRegisterOracleNodeLog(context.Background()); err != nil {
			log.Errorf("Watch and handle register oracle node log: %v", err)
		}
	}()

	go func() {
		if err := n.aggregator.WatchAndHandleValidationRequestsLog(context.Background(), n); err != nil {
			log.Errorf("Watch and handle ValidationRequest log: %v", err)
		}
	}()

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

	schnorrPublicKey := n.suite.Point().Mul(n.schnorrPrivateKey, nil)
	b, err := schnorrPublicKey.MarshalBinary()
	if err != nil {
		return fmt.Errorf("marshal bls public key: %v", err)
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
		_, err = n.registryContract.RegisterOracleNode(auth, ipAddr, b)
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
