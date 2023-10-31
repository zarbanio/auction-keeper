package actions

import (
	"log"

	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"

	"github.com/ethereum/go-ethereum/core/types"
)

func (a Actions) Heal(heal *entities.VowHeal) (*types.Transaction, error) {
	opts, err := a.sender.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	tx, err := a.Vow.Heal(opts, heal.Rad)
	if err != nil {
		return nil, err
	}
	log.Println("Heal Tx Hash: ", tx.Hash().String())

	return tx, nil
}

func (a Actions) Kiss(kiss *entities.VowKiss) (*types.Transaction, error) {
	opts, err := a.sender.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	tx, err := a.Vow.Kiss(opts, kiss.Rad)
	if err != nil {
		return nil, err
	}
	log.Println("Kiss Tx Hash: ", tx.Hash().String())

	return tx, nil
}

func (a Actions) Flop() (*types.Transaction, error) {
	opts, err := a.sender.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	tx, err := a.Vow.Flop(opts)
	if err != nil {
		return nil, err
	}
	log.Println("Flop Tx Hash: ", tx.Hash().String())

	return tx, nil
}

func (a Actions) Flog(flog *entities.VowFlog) (*types.Transaction, error) {
	opts, err := a.sender.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	tx, err := a.Vow.Flog(opts, flog.Era)
	if err != nil {
		return nil, err
	}
	log.Println("Flog Tx Hash: ", tx.Hash().String())

	return tx, nil
}
