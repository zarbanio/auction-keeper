package actions

import (
	"context"
	"log"

	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
)

func (a Actions) Hope(hope *entities.VatHope) error {
	opts, err := a.sender.GetOpts()
	if err != nil {
		return err
	}

	tx, err := a.vat.VatTransactor.Hope(opts, hope.Usr)
	if err != nil {
		return err
	}
	log.Println("Hope Tx Hash: ", tx.Hash().String())

	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}
	_, err = a.store.CreateHope(context.Background(), *hope, int64(txId))
	if err != nil {
		return err
	}

	// return a.store.UpdateTransactionBlock(
	// 	context.Background(),
	// 	txId,
	// 	recipt,
	// 	header.Time,
	// 	*recipt.BlockNumber,
	// 	recipt.BlockHash)

	return nil
}
