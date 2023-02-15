package store

import (
	"context"
	"math/big"

	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/IR-Digital-Token/auction-keeper/domain/math"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type takeModel struct {
	opts        *bind.TransactOpts
	id          string
	amt         string
	max         string
	who         string
	data        string
	transaction types.Transaction
}

func (f takeModel) toDomain() *vat.VatFork {
	var I [32]byte
	copy(I[:], common.Hex2Bytes(f.ilk))
	return &vat.VatFork{
		Ilk:  I,
		Src:  common.HexToAddress(f.src),
		Dst:  common.HexToAddress(f.dst),
		Dink: math.BigIntFromString(f.dink),
		Dart: math.BigIntFromString(f.dart),
		Raw:  *f.Raw.toDomain(),
	}
}

func (p postgres) StoreTakeTransaction(ctx context.Context, newTx types.Transaction, opts *bind.TransactOpts, id *big.Int, amt *big.Int, max *big.Int, who common.Address, data []byte) (int64, error) {
	// rest of the code here
	return 0, nil
}
