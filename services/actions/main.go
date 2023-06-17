package actions

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/dog"
	"github.com/zarbanio/auction-keeper/bindings/vat"
	"github.com/zarbanio/auction-keeper/bindings/vow"
	"github.com/zarbanio/auction-keeper/services/transaction"
	"github.com/zarbanio/auction-keeper/x/chain"
)

type Actions struct {
	vat    *vat.Vat
	Dog    *dog.Dog
	Vow    *vow.Vow
	sender *transaction.Sender
}

func NewActions(eth *ethclient.Client, sender *transaction.Sender, vatAddr, dogAddr common.Address, vowAddr common.Address, indexer *chain.Indexer) (IAction, error) {
	v, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	d, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		log.Fatal(err)
	}
	vo, err := vow.NewVow(vowAddr, eth)
	if err != nil {
		log.Fatal(err)
	}
	return &Actions{
		vat:    v,
		Dog:    d,
		Vow:    vo,
		sender: sender,
	}, nil
}
