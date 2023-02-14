package entities

import "math/big"

type Urn struct {
	Ink *big.Int `json:"ink"`
	Art *big.Int `json:"art"`
}
