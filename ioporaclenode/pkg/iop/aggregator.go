package iop

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/share"
	"go.dedis.ch/kyber/v3/sign/tbls"
)

type Aggregator struct {
	suite             pairing.Suite
	ethClient         *ethclient.Client
	dkg               *DistKeyGenerator
	connectionManager *ConnectionManager
	oracleContract    *OracleContractWrapper
	registryContract  *RegistryContractWrapper
	account           common.Address
	ecdsaPrivateKey   *ecdsa.PrivateKey
	chainId           *big.Int
	t                 int
}

func NewAggregator(
	suite pairing.Suite,
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	oracleContract *OracleContractWrapper,
	registryContract *RegistryContractWrapper,
	account common.Address,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	chainId *big.Int,
) *Aggregator {
	return &Aggregator{
		suite:             suite,
		ethClient:         ethClient,
		connectionManager: connectionManager,
		oracleContract:    oracleContract,
		registryContract:  registryContract,
		account:           account,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		chainId:           chainId,
	}
}

func (a *Aggregator) WatchAndHandleValidationRequestsLog(ctx context.Context) error {
	sink := make(chan *OracleContractValidationRequest)
	defer close(sink)

	sub, err := a.oracleContract.WatchValidationRequest(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			typ := ValidateRequest_Type(event.Typ)

			log.Infof("Received ValidationRequest event for %s type with hash %s", typ, common.Hash(event.Hash))
			isAggregator, err := a.registryContract.IsAggregator(nil, a.account)
			if err != nil {
				log.Errorf("Is aggregator: %v", err)
				continue
			}
			if !isAggregator {
				// 报名函数
				node, err := a.registryContract.FindOracleNodeByAddress(nil, a.account)
				time.Sleep(time.Duration(node.Index.Int64()) * time.Second)
				err = a.Enroll()
				if err != nil {
					log.Errorf("Node Enroll log: %v", err)
				} else {
					log.Infof("Enroll success")
				}
				continue
			}
			if err := a.HandleValidationRequest(ctx, event, typ); err != nil {
				log.Errorf("Handle ValidationRequest log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// 报名函数
func (a *Aggregator) Enroll() error {
	isEnroll, err := a.oracleContract.OracleNodeIsEnroll(nil, a.account)
	if err != nil {
		return fmt.Errorf("is enrolled: %w", err)
	}
	if !isEnroll {
		auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
		_, err = a.oracleContract.EnrollOracleNode(auth)
		if err != nil {
			return fmt.Errorf("enroll iop node: %w", err)
		}
	}
	return nil
}

func (a *Aggregator) HandleValidationRequest(ctx context.Context, event *OracleContractValidationRequest, typ ValidateRequest_Type) error {
	result, s, err := a.AggregateValidationResults(ctx, event.Hash, typ)
	if err != nil {
		return fmt.Errorf("aggregate validation results: %w", err)
	}

	sig, err := SignatureToBig(s)
	if err != nil {
		return fmt.Errorf("signature to big int: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}

	switch typ {
	case ValidateRequest_block:
		_, err = a.oracleContract.SubmitBlockValidationResult(auth, event.Hash, result, sig)
	case ValidateRequest_transaction:
		_, err = a.oracleContract.SubmitTransactionValidationResult(auth, event.Hash, result, sig)
	default:
		return fmt.Errorf("unknown validation request type %s", typ)
	}

	if err != nil {
		return fmt.Errorf("submit verification: %w", err)
	}

	resultStr := "valid"
	if !result {
		resultStr = "invalid"
	}
	log.Infof("Submitted validation result (%s) for hash %s of type %s", resultStr, common.Hash(event.Hash), typ)

	return nil
}

func (a *Aggregator) AggregateValidationResults(ctx context.Context, txHash common.Hash, typ ValidateRequest_Type) (bool, []byte, error) {

	positiveResults := make([][]byte, 0)
	negativeResults := make([][]byte, 0)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	// 获取到了报名的节点数
	time.Sleep(time.Duration(10) * time.Second)
	enrollNodeCount, err := a.oracleContract.CountEnrollNodes(nil)
	if err != nil {
		log.Errorf("Find connection by address: %v", err)
	}
	// 当报名数小于等于门限数值
	if enrollNodeCount.Int64() < int64(a.t) {
		nodes, err := a.registryContract.FindOracleNodes()
		if err != nil {
			return false, nil, fmt.Errorf("find nodes: %w", err)
		}
		for _, node := range nodes {
			// 如果节点没有报名，就质询
			isEnroll, err := a.oracleContract.OracleNodeIsEnroll(nil, node.Addr)
			if err != nil {
				log.Errorf("Oracle Enroll: %v", err)
				continue
			}
			if !isEnroll {
				conn, err := a.connectionManager.FindByAddress(node.Addr)
				if err != nil {
					log.Errorf("Find connection by address: %v", err)
					continue
				}

				wg.Add(1)
				go func() {
					defer wg.Done()

					client := NewOracleNodeClient(conn)
					ctxTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
					_, err = a.oracleContract.Chanllenge(nil) // 质询请求扣费
					if err != nil {
						log.Errorf("chanllenge is error: %v", err)
						// cancel()
						// return
					}
					// 质询验证器不工作
					proof, err := client.Chanllenge(ctxTimeout, &ChanllengeRequest{Chanllenge: []byte("chanllenge")})
					if !bytes.Equal(proof.Proof, []byte("This is a proof for lazyNode")) {
						_, err = a.oracleContract.Fine(nil) // 质询结果错误，扣费
						if err != nil {
							log.Errorf("Fine is Error : %v", err)
						}
					}

					if err != nil {
						log.Error("Chanllenge : %v", err)
					}
					cancel()
					if err != nil {
						log.Errorf("Validate %s: %v", typ, err)
						return
					}

				}()
			}
		}
		wg.Wait()
	} else {
		enrollNodes, err := a.oracleContract.FindEnrollNodes()

		if err != nil {
			return false, nil, fmt.Errorf("find enrollnodes: %w", err)
		}
		rand.Seed(time.Now().Unix())

		for _, enrollNode := range enrollNodes {
			node, err := a.registryContract.FindOracleNodeByAddress(nil, enrollNode.Addr)
			conn, err := a.connectionManager.FindByAddress(node.Addr)
			if err != nil {
				log.Errorf("Find connection by address: %v", err)
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				client := NewOracleNodeClient(conn)
				ctxTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
				result, err := client.Validate(ctxTimeout, &ValidateRequest{
					Type: typ,
					Hash: txHash[:],
				})
				if err != nil {
					log.Errorf("Validate %s: %v", typ, err)
					return
				}
				// 在这里考虑加上质询需求
				chanllenge := 1

				if chanllenge == 1 {
					auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
					if err != nil {
						log.Error("auth : %v ", err)
					}
					_, err = a.oracleContract.Chanllenge(auth) // 质询请求扣费
					if err != nil {
						log.Error("262 Chanllenge : %v ", err) //质询扣费出错
						// cancel()
						// return
					}
					proof, err := client.Chanllenge(ctxTimeout, &ChanllengeRequest{Type: []byte("chanllenge_result"), Chanllenge: []byte("chanllenge")})
					log.Infof(string(proof.Proof))
					if err != nil {
						log.Error("271 Chanllenge : %v ", err)
					}
					if bytes.Equal(proof.Proof, []byte("This is a proof for result")) {
						result.Valid = true
					} else {
						result.Valid = false
						auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
						if err != nil {
							log.Error("auth : %v ", err)
						}
						_, err = a.oracleContract.Fine(auth) // 质询结果错误，扣费     惩罚出错
					}
					if err != nil {
						log.Error("Chanllenge : %v ", err)
					}
				}
				cancel()
				if err != nil {
					log.Errorf("Validate %s: %v", typ, err)
					return
				}
				mutex.Lock()
				if result.Valid {
					positiveResults = append(positiveResults, result.Signature)
				} else {
					negativeResults = append(negativeResults, result.Signature)
				}
				mutex.Unlock()
			}()
		}

		wg.Wait()

		distKey, err := a.dkg.DistKeyShare()
		if err != nil {
			return false, nil, fmt.Errorf("dist key share: %w", err)
		}

		pubPoly := share.NewPubPoly(a.suite.G2(), a.suite.G2().Point().Base(), distKey.Commits)

		result := false
		sigShares := negativeResults

		if len(positiveResults) >= len(negativeResults) {
			result = true
			sigShares = positiveResults
		}

		message, err := encodeValidateResult(txHash, result, typ)
		if err != nil {
			return false, nil, fmt.Errorf("encode validation result: %w", err)
		}

		signature, err := tbls.Recover(a.suite, pubPoly, message, sigShares, a.t, len(enrollNodes))
		if err != nil {
			return false, nil, fmt.Errorf("recover signature: %w", err)
		}

		return result, signature, nil
	}
	return false, nil, nil
}

func (a *Aggregator) SetDistKeyGenerator(dkg *DistKeyGenerator) {
	a.dkg = dkg
}

func (a *Aggregator) SetThreshold(threshold int) {
	a.t = threshold
}
