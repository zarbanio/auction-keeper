package actions

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/zarbanio/auction-keeper/bindings/zarban/dog"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vow"
	"github.com/zarbanio/auction-keeper/services/sender"
)

type Actions struct {
	vat    *vat.Vat
	Dog    *dog.Dog
	Vow    *vow.Vow
	sender sender.Sender
}

func NewActions(eth *ethclient.Client, sender sender.Sender, vatAddr, dogAddr common.Address, vowAddr common.Address) (IAction, error) {
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
