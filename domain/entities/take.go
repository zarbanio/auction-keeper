package entities

import (
	"math/big"

	"github.com/IR-Digital-Token/auction-keeper/domain/math"
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
type takeModel struct {
	auction_id  string
	amt         string
	max         string
	who         string
	data        string
	opts        *bind.TransactOpts
	transaction types.Transaction
}

func (f takeModel) ToDomain() *ClipperTake {
	var D []byte
	copy(D[:], common.Hex2Bytes(f.data))
	return &ClipperTake{
		Auction_id:  math.BigIntFromString(f.auction_id),
		Amt:         math.BigIntFromString(f.amt),
		Max:         math.BigIntFromString(f.max),
		Who:         common.HexToAddress(f.who),
		Data:        D,
		Opts:        f.opts,
		Transaction: f.transaction,
	}
}

func NewTake(id *big.Int, amt *big.Int, max *big.Int, who common.Address, data []byte) *takeModel {
	return &takeModel{
		auction_id: id.String(),
		amt:        amt.String(),
		max:        max.String(),
		who:        who.String(),
		data:       string(data),
	}
}
