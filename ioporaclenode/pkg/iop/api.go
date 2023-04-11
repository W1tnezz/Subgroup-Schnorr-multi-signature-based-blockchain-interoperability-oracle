package iop

import (
	"bytes"
	"context"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *OracleNode) Chanllenge(_ context.Context, request *ChanllengeRequest) (*ChanllengeResponse, error) {
	//这个函数的中，此时验证器需要构造出一个证明，并且将证明作为结果返回,当前的问题就是找出 ：这个函数的调用场景
	var proof []byte
	if bytes.Equal(request.Type, []byte("chanllenge_result")) {
		proof = []byte("This is a proof for result") // 这是关于消息的回复
	} else if bytes.Equal(request.Type, []byte("chanllenge_lazyNode")) {
		proof = []byte("This is a proof for lazyNode") //这是关于是否工作的回复
	} else {
		proof = []byte("This is a proof for default") //这是关于是否工作的回复
	}
	return &ChanllengeResponse{Proof: proof}, nil
}

func (n *OracleNode) SendR(_ context.Context, request *SendRRequest) (*SendRResponse, error) {
	// 这里接收到了传递过来的参数R
	n.validator.HandleR(request.R)
	return &SendRResponse{}, nil
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

func ValidateResultToResponse(result *ValidateResult) *ValidateResponse {
	resp := &ValidateResponse{
		Hash:      result.hash[:],
		Valid:     result.valid,
		Signature: result.signature,
		R:         result.R,
	}

	if result.blockNumber != nil {
		resp.BlockNumber = result.blockNumber.Int64()
	}

	return resp
}
