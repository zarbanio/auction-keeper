package actions

import (
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
	Take(clipper *clipper.Clipper, take *entities.ClipperTake) error
	Redo(clipper *clipper.Clipper, redo *entities.ClipperRedo) error
}

type Vat interface {
	Hope(hope *entities.VatHope) error
}

type Dog interface {
	Bark(bark *entities.DogBark) error
}

type Vow interface {
	Heal(heal *entities.VowHeal) error
	Kiss(kiss *entities.VowKiss) error
	Flop() error
	Flog(flog *entities.VowFlog) error
}
