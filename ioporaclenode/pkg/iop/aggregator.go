package iop

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
)

type Aggregator struct {
	suite             pairing.Suite
	ethClient         *ethclient.Client
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

func (a *Aggregator) WatchAndHandleValidationRequestsLog(ctx context.Context, o *OracleNode) error {
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
			o.isAggregator = isAggregator
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
	result, MulSig, MulR, apk, _hash, err := a.AggregateValidationResults(ctx, event.Hash, typ)
	fmt.Println("result:               ", result)
	fmt.Println("MultiSignature:       ", MulSig)
	fmt.Println("MultiR:               ", MulR)
	fmt.Println("Aggregate Public Key: ", apk)
	fmt.Println("hash:                 ", _hash)

	if err != nil {
		return fmt.Errorf("aggregate validation results: %w", err)
	}
	if err != nil {
		return fmt.Errorf("signature to big int: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}

	sig, err := ScalarToBig(MulSig)
	if err != nil {
		return fmt.Errorf("signature tranform to big int: %w", err)
	}
	pubKey, err := PointToBig(apk)
	if err != nil {
		return fmt.Errorf("public key tranform to big int: %w", err)
	}
	R, err := PointToBig(MulR)
	if err != nil {
		return fmt.Errorf("multi R tranform to big int: %w", err)
	}
	hash, err := ScalarToBig(_hash)
	if err != nil {
		return fmt.Errorf("hash tranform to big int: %w", err)
	}
	switch typ {
	case ValidateRequest_block:
		_, err = a.oracleContract.SubmitBlockValidationResult(auth, result, event.Hash, sig, pubKey[0], pubKey[1], R[0], R[1], hash)
	case ValidateRequest_transaction:
		_, err = a.oracleContract.SubmitTransactionValidationResult(auth, result, event.Hash, sig, pubKey[0], pubKey[1], R[0], R[1], hash)
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

func (a *Aggregator) AggregateValidationResults(ctx context.Context, txHash common.Hash, typ ValidateRequest_Type) (bool, kyber.Scalar, kyber.Point, kyber.Point, kyber.Scalar, error) {

	Signatures := make([]kyber.Scalar, 0)
	Rs := make([]kyber.Point, 0)
	J := make([][]byte, 0)

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
		wg.Wait()
	} else {
		enrollNodes, err := a.oracleContract.FindEnrollNodes()

		if err != nil {
			return false, nil, nil, nil, nil, fmt.Errorf("find enrollnodes: %w", err)
		}
		rand.Seed(time.Now().Unix())

		for _, enrollNode := range enrollNodes {
			node, err := a.registryContract.FindOracleNodeByAddress(nil, enrollNode)
			conn, err := a.connectionManager.FindByAddress(node.Addr)
			if err != nil {
				log.Errorf("Find connection by address: %v", err)
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				client := NewOracleNodeClient(conn)
				ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
				result, err := client.Validate(ctxTimeout, &ValidateRequest{
					Type: typ,
					Hash: txHash[:],
				})
				if err != nil {
					log.Errorf("Validate %s: %v", typ, err)
					return
				}

				cancel()
				if err != nil {
					log.Errorf("Validate %s: %v", typ, err)
					return
				}

				mutex.Lock()
				if result.Valid {
					s := a.suite.G1().Scalar()
					err := s.UnmarshalBinary(result.Signature)
					if err != nil {
						fmt.Errorf("s transform to Scalar: %w", err)
					}
					Signatures = append(Signatures, s) //
					RPoint := a.suite.G1().Point()
					err = RPoint.UnmarshalBinary(result.R)
					if err != nil {
						fmt.Errorf("R transform to Point: %w", err)
					}
					Rs = append(Rs, RPoint)

					J = append(J, node.PubKey)
				}
				mutex.Unlock()
			}()
		}

		wg.Wait()

		R := a.suite.G1().Point().Null()
		for i := 0; i < len(Rs); i++ {
			R.Add(R, Rs[i])
		}
		PK, err := a.registryContract.GetAllPk()
		if err != nil {
			fmt.Errorf("get PK: %w", err)
		}
		hash_1 := sha256.New()
		message, err := encodeValidateResult(txHash, true, typ)

		gt := hash_1.Sum(bytes.Join(PK, []byte("")))

		m := make([][]byte, 3)
		m[0] = message
		m[1], err = R.MarshalBinary()
		m[2] = gt
		hash := sha256.New()
		e := hash.Sum(bytes.Join(m, []byte("")))
		MulSignature := a.suite.G1().Scalar().Zero()
		MulR := a.suite.G1().Point().Null()
		apk := a.suite.G1().Point().Null()

		for i := 0; i < len(J); i++ {
			pub := a.suite.G1().Point()
			err = pub.UnmarshalBinary(J[i])
			verify_R := Rs[i].Clone()
			verify_R.Add(verify_R, a.suite.G1().Point().Mul(a.suite.G1().Scalar().SetBytes(e), pub))
			S2 := a.suite.G1().Point().Mul(Signatures[i], nil)
			if !verify_R.Equal(S2) {
				return false, nil, nil, nil, nil, fmt.Errorf("签名验证失败 ，该签名的公钥为：", J[i])
			}
			hash := sha256.New()
			h := make([][]byte, 3)
			h[0] = J[i]
			h[1] = bytes.Join(J, []byte(""))
			h[2] = bytes.Join(PK, []byte(""))
			a_j := hash.Sum(bytes.Join(h, []byte("")))
			aScalar := a.suite.G1().Scalar().SetBytes(a_j)

			MulSignature.Add(MulSignature, a.suite.G1().Scalar().Mul(aScalar, Signatures[i]))
			MulR.Add(MulR, a.suite.G1().Point().Mul(aScalar, Rs[i]))
			apk.Add(apk, a.suite.G1().Point().Mul(aScalar, pub))

		}
		_hash := a.suite.G1().Scalar().SetBytes(e)
		tempR := MulR.Clone()
		S3 := tempR.Add(tempR, a.suite.G1().Point().Mul(_hash, apk))
		S4 := a.suite.G1().Point().Mul(MulSignature, nil)
		if S3.Equal(S4) {
			fmt.Println("多重签名链下验证成功")
		} else {
			fmt.Println("多重签名链下验证失败")
		}
		return true, MulSignature, MulR, apk, _hash, nil
	}
	return false, nil, nil, nil, nil, nil
}

func (a *Aggregator) SetThreshold(threshold int) {
	a.t = threshold
}
