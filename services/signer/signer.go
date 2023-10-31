package signer

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer interface {
	GetAddress() common.Address
	GetTransactOpts() (*bind.TransactOpts, error)
}

type signer struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainId    *big.Int
}

func NewSigner(privateKey string, chainId *big.Int) (Signer, error) {
	prvKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := prvKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &signer{
		privateKey: prvKey,
		address:    address,
		chainId:    chainId,
	}, nil
}

func (s signer) GetAddress() common.Address {
	return s.address
}

func (s signer) GetTransactOpts() (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.chainId)
	if err != nil {
		return nil, err
	}
	return opts, nil
}
