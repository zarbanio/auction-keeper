package actions

import (
	"log"

	"github.com/IR-Digital-Token/auction-keeper/bindings/dog"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vow"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/IR-Digital-Token/auction-keeper/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Actions struct {
	vat    *vat.Vat
	Dog    *dog.Dog
	Vow    *vow.Vow
	sender *transaction.Sender
	store  store.IStore
}

func NewActions(eth *ethclient.Client, sender *transaction.Sender, store store.IStore, vatAddr, dogAddr common.Address, vowAddr common.Address) (IAction, error) {
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
		store:  store,
	}, nil
}
