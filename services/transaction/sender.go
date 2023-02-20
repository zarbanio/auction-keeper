package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/x/chain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Sender struct {
	eth        *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainId    *big.Int
	indexer    *chain.Indexer
}

func NewSender(eth *ethclient.Client, privateKey string, chainId *big.Int, vatAddr, dogAddr common.Address, indexer *chain.Indexer) (*Sender, error) {

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
		indexer:    indexer,
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
	//opts.GasLimit = uint64(300000) // in units //TODO
	opts.GasPrice = gasPrice

	return opts, nil
}

func (s *Sender) SendTakeTx(clipper *clipper.Clipper, id, amt, maxPrice *big.Int, exchangeCalleeAddress common.Address, flashData []byte) error {

	opts, err := s.GetOpts()
	if err != nil {
		return err
	}

	tx, err := clipper.ClipperTransactor.Take(opts, id, amt, maxPrice, exchangeCalleeAddress, flashData)
	if err != nil {
		return err
	}
	cb := func(header types.Header, recipt *types.Receipt) error {
		return nil
	}
	txHandler := NewHandler(*tx, cb)
	s.WatchTransactionHash(txHandler)
	// store.StoreTakeTransaction()

	fmt.Printf("Take Transaction Sent: %s\n", tx.Hash().Hex())

	return nil
}
