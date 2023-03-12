package iop

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	vss "go.dedis.ch/kyber/v3/share/vss/pedersen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *OracleNode) Chanllenge(_ context.Context, request *ChanllengeRequest) (*ChanllengeResponse, error) {
	//这个函数的中，此时验证器需要构造出一个证明，并且将证明作为结果返回,当前的问题就是找出 ：这个函数的调用场景
	var proof []byte
	if bytes.Equal(request.Type, []byte("chanllenge_result")){
		proof = []byte("This is a proof for result") // 这是关于消息的回复
	}else if bytes.Equal(request.Type, []byte("chanllenge_lazyNode")){
		proof = []byte("This is a proof for lazyNode") //这是关于是否工作的回复
	}else {
		proof = []byte("This is a proof for default") //这是关于是否工作的回复
	}
	return &ChanllengeResponse{Proof: proof}, nil
}

func (n *OracleNode) SendDeal(_ context.Context, request *SendDealRequest) (*SendDealResponse, error) {
	_, err := n.dkg.ProcessDeal(PbToDeal(request.Deal))
	if err != nil {
		return nil, fmt.Errorf("handle deal: %w", err)
	}
	return &SendDealResponse{}, nil
}

// 这个函数的功能是验证器来验证的过程，以及构造出应答
func (n *OracleNode) Validate(ctx context.Context, request *ValidateRequest) (*ValidateResponse, error) {

	var result *ValidateResult
	var err error

	switch request.Type {
	case ValidateRequest_block:
		result, err = n.validator.ValidateBlock(
			ctx,
			common.BytesToHash(request.Hash),
		)
	case ValidateRequest_transaction:
		result, err = n.validator.ValidateTransaction(
			ctx,
			common.BytesToHash(request.Hash),
		)
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "validate %s: %v", request.Type, err)
	}

	resultStr := "valid"
	if !result.valid {
		resultStr = "invalid"
	}
	log.Infof("Validated hash %s of type %s with result: %s", common.BytesToHash(request.Hash), request.Type, resultStr)

	return ValidateResultToResponse(result), nil
}

func DealToPb(deal *dkg.Deal) *Deal {
	return &Deal{
		Index:     deal.Index,
		Deal:      EncryptedDealToPb(deal.Deal),
		Signature: deal.Signature,
	}
}

func PbToDeal(deal *Deal) *dkg.Deal {
	return &dkg.Deal{
		Index:     deal.Index,
		Deal:      PbToEncryptedDeal(deal.Deal),
		Signature: deal.Signature,
	}
}

func EncryptedDealToPb(encryptedDeal *vss.EncryptedDeal) *EncryptedDeal {
	return &EncryptedDeal{
		DhKey:     encryptedDeal.DHKey,
		Signature: encryptedDeal.Signature,
		Nonce:     encryptedDeal.Nonce,
		Cipher:    encryptedDeal.Cipher,
	}
}

func PbToEncryptedDeal(encryptedDeal *EncryptedDeal) *vss.EncryptedDeal {
	return &vss.EncryptedDeal{
		DHKey:     encryptedDeal.DhKey,
		Signature: encryptedDeal.Signature,
		Nonce:     encryptedDeal.Nonce,
		Cipher:    encryptedDeal.Cipher,
	}
}

func ValidateResultToResponse(result *ValidateResult) *ValidateResponse {
	resp := &ValidateResponse{
		Hash:      result.hash[:],
		Valid:     result.valid,
		Signature: result.signature,
	}

	if result.blockNumber != nil {
		resp.BlockNumber = result.blockNumber.Int64()
	}

	return resp
}
