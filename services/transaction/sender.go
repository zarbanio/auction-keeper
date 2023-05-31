package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Sender struct {
	eth        *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainId    *big.Int
}

func NewSender(eth *ethclient.Client, privateKey string, chainId *big.Int) (*Sender, error) {
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

	return &Sender{
		eth:        eth,
		privateKey: prvKey,
		address:    address,
		chainId:    chainId,
	}, nil
}

func (s Sender) GetAddress() common.Address {
	return s.address
}

func (s Sender) GetOpts() (*bind.TransactOpts, error) {
	nonce, err := s.eth.PendingNonceAt(context.Background(), s.GetAddress())
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.eth.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.chainId)
	if err != nil {
		return nil, err
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasLimit = uint64(800_000) //TODO
	opts.GasPrice = gasPrice

	return opts, nil
}
