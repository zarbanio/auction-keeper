package actions

import (
	"context"

	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/ethereum/go-ethereum/core/types"
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
	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}
	_, err = a.store.CreateHope(context.Background(), *hope, int64(txId))
	if err != nil {
		return err
	}
	txHandler := transaction.NewHandler(*tx, func(header types.Header, recipt *types.Receipt) error {
		return a.store.UpdateTransactionBlock(
			context.Background(),
			uint64(txId),
			recipt,
			header.Time,
			*recipt.BlockNumber,
			recipt.BlockHash)
	})
	a.sender.WatchTransactionHash(txHandler)

	return nil
}
