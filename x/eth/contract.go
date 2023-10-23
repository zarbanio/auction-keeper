package eth

import "github.com/ethereum/go-ethereum/common"

type Contract[T any] struct {
	Address common.Address
	Binding T
}
