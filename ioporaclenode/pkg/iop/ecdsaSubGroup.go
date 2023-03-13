package iop

import (
	"context"
	"crypto/ecdsa"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	iota "github.com/iotaledger/iota.go/v2"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"go.dedis.ch/kyber/v3/suites"
	"math/big"
	"sync"
)

type SubgroupSigGenerator struct {
	sync.Mutex
	dkg               *dkg.DistKeyGenerator
	suite             suites.Suite
	connectionManager *ConnectionManager
	aggregator        *Aggregator
	mqttClient        mqtt.Client
	mqttTopic         []byte
	iotaClient        *iota.NodeHTTPAPIClient
	registryContract  *RegistryContractWrapper
	oracleContract    *OracleContract
	ecdsaPrivateKey   *ecdsa.PrivateKey
	blsPrivateKey     kyber.Scalar
	account           common.Address
	deals             map[uint32]*dkg.Deal
	pendingResp       map[uint32][]*dkg.Response
	index             uint32
	chainId           *big.Int
}

func NewSubgroupSigGenerator(
	connectionManager *ConnectionManager,
	aggregator *Aggregator,
	mqttClient mqtt.Client,
	mqttTopic []byte,
	iotaClient *iota.NodeHTTPAPIClient,
	registryContract *RegistryContractWrapper,
	oracleContract *OracleContract,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	account common.Address,
	chainId *big.Int,
) *SubgroupSigGenerator {
	return &SubgroupSigGenerator{
		connectionManager: connectionManager,
		aggregator:        aggregator,
		mqttClient:        mqttClient,
		mqttTopic:         mqttTopic,
		iotaClient:        iotaClient,
		registryContract:  registryContract,
		oracleContract:    oracleContract,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		account:           account,
		deals:             make(map[uint32]*dkg.Deal),
		pendingResp:       make(map[uint32][]*dkg.Response),
		chainId:           chainId,
	}
}

func (g *SubgroupSigGenerator) ListenAndProcess(ctx context.Context) error {
	// 启动协程监听并处理ECDSA交互；
	go func() {
		if err := g.WatchAndHandleECDSALog(ctx); err != nil {
			log.Errorf("Watch and handle DKG log: %v", err)
		}
	}()

	// 启动协程监听并处理交互过程中其他节点发送的Deal
	go func() {
		if err := g.ListenAndProcessResponse(); err != nil {
			log.Errorf("Listen and process response: %v", err)
		}
	}()
	return nil
}

func (g *SubgroupSigGenerator) WatchAndHandleECDSALog(ctx context.Context) error {
	sink := make(chan *OracleContractValidationBegin)
	defer close(sink)

	sub, err := g.oracleContract.WatchValidationBegin(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	for {
		select {
		case event := <-sink:
			log.Infof("Received Validation Begin event!")
			if err := g.HandleValidationBeginLog(event); err != nil {
				log.Errorf("Handle ValidationBegin log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (g *SubgroupSigGenerator) ListenAndProcessResponse() error {
	// TODO: 实现监听处理功能；
	return nil
}

func (g *SubgroupSigGenerator) HandleValidationBeginLog(event *OracleContractValidationBegin) error {
	// TODO:收到链上通知后，开始子分组ECDSA交互；
	return nil
}
