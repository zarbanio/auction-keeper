package store

import (
	"context"
	"math/big"

	"github.com/IR-Digital-Token/auction-keeper/domain/math"
	"github.com/IR-Digital-Token/auction-keeper/services/actions"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type takeModel struct {
	id          string
	amt         string
	max         string
	who         string
	data        string
	opts        *bind.TransactOpts
	transaction types.Transaction
}

func (f takeModel) toDomain() *actions.ClipperTake {
	var D []byte
	copy(D[:], common.Hex2Bytes(f.data))
	return &actions.ClipperTake{
		Id:          math.BigIntFromString(f.id),
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
		id:   id.String(),
		amt:  amt.String(),
		max:  max.String(),
		who:  who.String(),
		data: string(data),
	}
}

func (p postgres) CreateTake(ctx context.Context, take actions.ClipperTake, newTx types.Transaction, opts *bind.TransactOpts) (int64, error) {
	// rest of the code here
	return 0, nil
}
