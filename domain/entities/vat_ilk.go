package entities

import "math/big"

type VatIlk struct {
	Id   [32]byte `json:"id"`
	Name string   `json:"name"`
	Art  *big.Int `json:"art"`
	Rate *big.Int `json:"rate"`
	Spot *big.Int `json:"spot"`
	Line *big.Int `json:"line"`
	Dust *big.Int `json:"dust"`
}
