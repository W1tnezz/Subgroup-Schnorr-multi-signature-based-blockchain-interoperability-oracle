package iop

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"go.dedis.ch/kyber/v3"
	"math/big"
)

func AddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("could not cast to public key ecdsa")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
}

func HexToScalar(suite kyber.Group, hexScalar string) (kyber.Scalar, error) {
	b, err := hex.DecodeString(hexScalar)
	if byteErr, ok := err.(hex.InvalidByteError); ok {
		return nil, fmt.Errorf("invalid hex character %q in scalar", byte(byteErr))
	} else if err != nil {
		return nil, errors.New("invalid hex data for scalar")
	}
	s := suite.Scalar()
	if err := s.UnmarshalBinary(b); err != nil {
		return nil, fmt.Errorf("unmarshal scalar binary: %w", err)
	}
	return s, nil
}

func PointToBig(point kyber.Point) ([2]*big.Int, error) {
	bytes, err := point.MarshalBinary()
	if err != nil {
		return [2]*big.Int{}, fmt.Errorf("marshal public key: %w", err)
	}

	if len(bytes) != 64 {
		return [2]*big.Int{}, fmt.Errorf("invalid public key length")
	}

	return [2]*big.Int{
		new(big.Int).SetBytes(bytes[:32]),
		new(big.Int).SetBytes(bytes[32:64]),
	}, nil
}

func ScalarToBig(scalar kyber.Scalar) (*big.Int, error) {
	bytes, err := scalar.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("marshal signature: %w", err)
	}
	if len(bytes) != 32 {
		return nil, fmt.Errorf("invalid signature length")
	}
	return new(big.Int).SetBytes(bytes), nil
}
