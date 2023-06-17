package actions

import (
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
)

func (a Actions) Bark(bark *entities.DogBark) (*types.Transaction, error) {
	opts, err := a.sender.GetOpts()
	if err != nil {
		return nil, err
	}

	tx, err := a.Dog.DogTransactor.Bark(opts, bark.Ilk, bark.Urn, a.sender.GetAddress())
	if err != nil {
		return nil, err
	}
	log.Println("Bark Tx Hash: ", tx.Hash().String())

	return tx, nil
}
