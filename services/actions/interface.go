package actions

import (
	"math/big"

	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/ethereum/go-ethereum/common"
)

type IAction interface {
	Clipper
	Vat
	Dog
}

type Clipper interface {
	Take(clipper *clipper.Clipper, take ClipperTake) (string, error)
	Redo(clipper *clipper.Clipper, id *big.Int) (string, error)
}

type Vat interface {
	Hope(usr common.Address) (string, error)
}

type Dog interface {
	Bark(ilk [32]byte, urn common.Address) (string, error)
}
