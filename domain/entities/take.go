package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ClipperTake struct {
	Auction_id  *big.Int
	Amt         *big.Int
	Max         *big.Int
	Who         common.Address
	Data        []byte
	Opts        *bind.TransactOpts
	Transaction types.Transaction
}
