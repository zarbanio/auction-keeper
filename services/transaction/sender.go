package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Sender struct {
	eth        *ethclient.Client
	privateKey *ecdsa.PrivateKey
	Address    common.Address
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
		Address:    address,
		chainId:    chainId,
	}, nil
}

func (s *Sender) getOpts() (*bind.TransactOpts, error) {

	nonce, err := s.eth.PendingNonceAt(context.Background(), s.Address)
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
	opts.GasLimit = uint64(300000) // in units //TODO
	opts.GasPrice = gasPrice

	return opts, nil
}

func (s *Sender) SendTakeTx(clipper *clipper.Clipper, id, amt, maxPrice *big.Int, exchangeCalleeAddress common.Address, flashData []byte) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := clipper.ClipperTransactor.Take(opts, id, amt, maxPrice, exchangeCalleeAddress, flashData)
	if err != nil {
		return "", err
	}

	fmt.Printf("Take Transaction Sent: %s\n", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (s *Sender) SendVatHopeTx(vat *vat.Vat, usr common.Address) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := vat.VatTransactor.Hope(opts, usr)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
