package domain

import "github.com/ethereum/go-ethereum/common"

type Token struct {
	Decimals int64
	Symbol   Symbol
	Address  common.Address
}
