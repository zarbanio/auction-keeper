package actions

import (
	"context"
	"math/big"

	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"

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
			uint64(txId),
			recipt,
			header.Time,
			*recipt.BlockNumber,
			recipt.BlockHash)
	})
	a.sender.WatchTransactionHash(txHandler)

	return nil
}

func (a Actions) Redo(clipper *clipper.Clipper, id *big.Int) (string, error) {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return "", err
	}

	tx, err := clipper.ClipperTransactor.Redo(opts, id, a.sender.GetAddress())
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
