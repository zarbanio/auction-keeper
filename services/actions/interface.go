package actions

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
)

type IAction interface {
	Clipper
	Vat
	Dog
}

type Clipper interface {
	Take(clipper *clipper.Clipper, take *entities.ClipperTake) error
	Redo(clipper *clipper.Clipper, redo *entities.ClipperRedo) error
}

type Vat interface {
	Hope(hope *entities.VatHope) error
}

type Dog interface {
	Bark(bark *entities.DogBark) error
}
