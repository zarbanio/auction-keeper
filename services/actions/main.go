package actions

import (
	"log"

	"github.com/IR-Digital-Token/auction-keeper/bindings/dog"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Actions struct {
	vat    *vat.Vat
	Dog    *dog.Dog
	sender *transaction.Sender
}

func NewActions(eth *ethclient.Client, sender *transaction.Sender, vatAddr, dogAddr common.Address) (IAction, error) {
	v, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	d, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		log.Fatal(err)
	}
	return &Actions{
		vat:    v,
		Dog:    d,
		sender: sender,
	}, nil
}
