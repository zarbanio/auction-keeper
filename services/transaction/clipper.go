package transaction

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (s Sender) Take(clipper *clipper.Clipper, id, amt, maxPrice *big.Int, exchangeCalleeAddress common.Address, flashData []byte) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := clipper.ClipperTransactor.Take(opts, id, amt, maxPrice, exchangeCalleeAddress, flashData)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s Sender) Redo(clipper *clipper.Clipper, id *big.Int) (string, error) {

	opts, err := s.getOpts()
	if err != nil {
		return "", err
	}

	tx, err := clipper.ClipperTransactor.Redo(opts, id, s.GetAddress())
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
