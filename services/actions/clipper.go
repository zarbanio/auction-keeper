package actions

import (
	"log"

	clipper "github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"

	"github.com/ethereum/go-ethereum/core/types"
)

func (a Actions) Take(clipper *clipper.Clipper, take *entities.ClipperTake) (*types.Transaction, error) {
	opts, err := a.sender.GetTransactOpts()

	if err != nil {
		return nil, err
	}

	tx, err := clipper.ClipperTransactor.Take(opts, take.AuctionId, take.Amt, take.Max, take.Who, take.Data)
	if err != nil {
		return nil, err
	}
	log.Println("Take Tx Hash: ", tx.Hash().String())
	return tx, nil
}

func (a Actions) Redo(clipper *clipper.Clipper, redo *entities.ClipperRedo) (*types.Transaction, error) {
	opts, err := a.sender.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	tx, err := clipper.ClipperTransactor.Redo(opts, redo.SailId, a.sender.GetAddress())
	if err != nil {
		return nil, err
	}
	log.Println("Redo Tx Hash: ", tx.Hash().String())

	return tx, nil
}
