package actions

import (
	"context"
	"math/big"

	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ClipperTake struct {
	Auction_id  *big.Int
	Amt         *big.Int
	Max         *big.Int
	Who         common.Address
	Data        []byte
	Opts        *bind.TransactOpts
	Transaction types.Transaction
}

func (a Actions) Take(clipper *clipper.Clipper, take ClipperTake) (string, error) {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return "", err
	}

	tx, err := clipper.ClipperTransactor.Take(opts, take.Auction_id, take.Amt, take.Max, take.Who, take.Data)
	if err != nil {
		return "", err
	}
	store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	tx
	cb := func(header types.Header, recipt *types.Receipt) error {
		return nil
	}
	// txHandler := NewHandler(*tx, cb)

	// s.watchTransactionHash(txHandler)

	return tx.Hash().Hex(), nil
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
