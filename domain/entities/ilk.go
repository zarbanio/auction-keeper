package entities

import "math/big"

type Ilk struct {
	Id   [32]byte `json:"-"`
	Name string   `json:"name"`
	Art  *big.Int `json:"art"`
	Rate *big.Int `json:"rate"`
	Spot *big.Int `json:"spot"`
	Line *big.Int `json:"line"`
	Dust *big.Int `json:"dust"`
}
