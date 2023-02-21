package transaction

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ISender interface {
	GetAddress() common.Address
	Clipper
	Vat
	Vow
	Dog
}

type Clipper interface {
	Take(clipper *clipper.Clipper, id, amt, maxPrice *big.Int, exchangeCalleeAddress common.Address, flashData []byte) (string, error)
	Redo(clipper *clipper.Clipper, id *big.Int) (string, error)
}

type Vat interface {
	Hope(usr common.Address) (string, error)
}

type Vow interface {
	Flop() (string, error)
	Flog(era *big.Int) (string, error)
	Kiss(rad *big.Int) (string, error)
	Heal(rad *big.Int) (string, error)
}

type Dog interface {
	Bark(ilk [32]byte, urn common.Address) (string, error)
}
