package actions

import (
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
)

func (a Actions) Hope(hope *entities.VatHope) (*types.Transaction, error) {
	opts, err := a.sender.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	tx, err := a.vat.VatTransactor.Hope(opts, hope.Usr)
	if err != nil {
		return nil, err
	}
	log.Println("Hope Tx Hash: ", tx.Hash().String())

	return tx, nil
}
