package transaction

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ISender interface {
	GetAddress() common.Address
	Clipper
	Vat
}

type Clipper interface {
	Take(clipper *clipper.Clipper, id, amt, maxPrice *big.Int, exchangeCalleeAddress common.Address, flashData []byte) (string, error)
	Redo(clipper *clipper.Clipper, id *big.Int) (string, error)
}

type Vat interface {
	Hope(vat *vat.Vat, usr common.Address) (string, error)
}
