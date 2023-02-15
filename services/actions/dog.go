package actions

import (
	"github.com/ethereum/go-ethereum/common"
)

func (a Actions) Bark(ilk [32]byte, urn common.Address) (string, error) {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return "", err
	}

	tx, err := a.Dog.DogTransactor.Bark(opts, ilk, urn, a.sender.GetAddress())
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
