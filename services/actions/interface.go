package actions

import (
	"github.com/ethereum/go-ethereum/core/types"
	clipper "github.com/zarbanio/auction-keeper/bindings/clip"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
)

type IAction interface {
	Clipper
	Vat
	Dog
	Vow
}

type Clipper interface {
	Take(clipper *clipper.Clipper, take *entities.ClipperTake) (*types.Transaction, error)
	Redo(clipper *clipper.Clipper, redo *entities.ClipperRedo) (*types.Transaction, error)
}

type Vat interface {
	Hope(hope *entities.VatHope) (*types.Transaction, error)
}

type Dog interface {
	Bark(bark *entities.DogBark) (*types.Transaction, error)
}

type Vow interface {
	Heal(heal *entities.VowHeal) (*types.Transaction, error)
	Kiss(kiss *entities.VowKiss) (*types.Transaction, error)
	Flop() (*types.Transaction, error)
	Flog(flog *entities.VowFlog) (*types.Transaction, error)
}
