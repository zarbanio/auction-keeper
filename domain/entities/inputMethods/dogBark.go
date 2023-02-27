package entities

import (
	"github.com/ethereum/go-ethereum/common"
)

type DogBark struct {
	Ilk [32]byte
	Urn common.Address
}
