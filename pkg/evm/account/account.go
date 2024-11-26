package account

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EVMAccountDetail struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

func NewEVMAccount(privateKeyHex string) (*EVMAccountDetail, error) {
	ecdsaPKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := ecdsaPKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	operatorAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &EVMAccountDetail{
		Address:    operatorAddr,
		PrivateKey: ecdsaPKey,
	}, nil
}
