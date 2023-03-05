package actions

import (
	"context"
	"log"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"

	"github.com/ethereum/go-ethereum/core/types"
)

func (a Actions) Heal(heal *entities.VowHeal) error {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return err
	}

	tx, err := a.Vow.Heal(opts, heal.Rad)
	if err != nil {
		return err
	}
	log.Println("Heal Tx Hash: ", tx.Hash().String())

	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}
	_, err = a.store.CreateHeal(context.Background(), heal, int64(txId))
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

func (a Actions) Kiss(kiss *entities.VowKiss) error {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return err
	}

	tx, err := a.Vow.Kiss(opts, kiss.Rad)
	if err != nil {
		return err
	}
	log.Println("Kiss Tx Hash: ", tx.Hash().String())

	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}
	_, err = a.store.CreateKiss(context.Background(), kiss, int64(txId))
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

func (a Actions) Flop() error {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return err
	}

	tx, err := a.Vow.Flop(opts)
	if err != nil {
		return err
	}
	log.Println("Flop Tx Hash: ", tx.Hash().String())

	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}
	_, err = a.store.CreateFlop(context.Background(), int64(txId))
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

func (a Actions) Flog(flog *entities.VowFlog) error {

	opts, err := a.sender.GetOpts()
	if err != nil {
		return err
	}

	tx, err := a.Vow.Flog(opts, flog.Era)
	if err != nil {
		return err
	}
	log.Println("Flog Tx Hash: ", tx.Hash().String())

	err, txId := a.store.CreateTransaction(context.Background(), tx, a.sender.GetAddress())
	if err != nil {
		return err
	}
	_, err = a.store.CreateFlog(context.Background(), flog, int64(txId))
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
