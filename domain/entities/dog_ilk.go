package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type DogIlk struct {
	Clip common.Address `json:"clip"`
	Chop *big.Int       `json:"chop"`
	Hole *big.Int       `json:"hole"`
	Dirt *big.Int       `json:"dirt"`
}
