package actions

import (
	"context"
	"log"

	clipper "github.com/zarbanio/auction-keeper/bindings/clip"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/services/transaction"

	"github.com/ethereum/go-ethereum/core/types"
)

func (a Actions) Take(clipper *clipper.Clipper, take *entities.ClipperTake) error {
	opts, err := a.sender.GetOpts()
	if err != nil {
		return err
	}

	tx, err := clipper.ClipperTransactor.Take(opts, take.Auction_id, take.Amt, take.Max, take.Who, take.Data)
	if err != nil {
		return err
	}
	log.Println("Take Tx Hash: ", tx.Hash().String())

	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}
	_, err = a.store.CreateTake(context.Background(), take, int64(txId))
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

func (a Actions) Redo(clipper *clipper.Clipper, redo *entities.ClipperRedo) error {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return err
	}

	tx, err := clipper.ClipperTransactor.Redo(opts, redo.SailId, a.sender.GetAddress())
	if err != nil {
		return err
	}
	log.Println("Redo Tx Hash: ", tx.Hash().String())

	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}
	_, err = a.store.CreateRedo(context.Background(), *redo, int64(txId))
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
