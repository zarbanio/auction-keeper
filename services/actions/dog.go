package actions

import (
	"context"
	"log"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/ethereum/go-ethereum/core/types"
)

func (a Actions) Bark(bark *entities.DogBark) error {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return err
	}

	tx, err := a.Dog.DogTransactor.Bark(opts, bark.Ilk, bark.Urn, a.sender.GetAddress())
	if err != nil {
		return err
	}
	log.Println("Bark Tx Hash: ", tx.Hash().String())

	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}

	_, err = a.store.CreateBark(context.Background(), *bark, int64(txId))
	if err != nil {
		return err
	}
	txHandler := transaction.NewHandler(*tx, func(header types.Header, recipt *types.Receipt) error {
		return a.store.UpdateTransactionBlock(
			context.Background(),
			txId,
			recipt,
			header.Time,
			*recipt.BlockNumber,
			recipt.BlockHash)
	})
	a.sender.WatchTransactionHash(txHandler)
	return nil
}
