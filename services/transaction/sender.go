package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/bindings/dog"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vow"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type Sender struct {
	eth        *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainId    *big.Int
	vat        *vat.Vat
	vow        *vow.Vow
	dog        *dog.Dog
}

func NewSender(eth *ethclient.Client, privateKey string, chainId *big.Int, vatAddr, vowAddr, dogAddr common.Address) (ISender, error) {

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

	vatInstance, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	vowInstance, err := vow.NewVow(vowAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	dogInstance, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &Sender{
		eth:        eth,
		privateKey: prvKey,
		address:    address,
		chainId:    chainId,
		vat:        vatInstance,
		vow:        vowInstance,
		dog:        dogInstance,
	}, nil
}

func (s Sender) GetAddress() common.Address {
	return s.address
}

func (s Sender) getOpts() (*bind.TransactOpts, error) {

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
