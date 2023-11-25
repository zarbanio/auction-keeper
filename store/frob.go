package store

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type frobModel struct {
	I    string
	U    string
	V    string
	W    string
	Dink string
	Dart string
	Raw  evmLogModel
}

func (f frobModel) toDomain() *vat.VatFrob {
	var ilk [32]byte
	copy(ilk[:], common.Hex2Bytes(f.I))
	i := &vat.VatFrob{
		I:    ilk,
		U:    common.HexToAddress(f.U),
		V:    common.HexToAddress(f.V),
		W:    common.HexToAddress(f.W),
		Dink: math.BigIntFromString(f.Dink),
		Dart: math.BigIntFromString(f.Dart),
		Raw: types.Log{
			Address:     common.HexToAddress(f.Raw.Address),
			BlockNumber: f.Raw.BlockNumber,
			TxHash:      common.HexToHash(f.Raw.TxHash),
			BlockHash:   common.HexToHash(f.Raw.BlockHash),
			Index:       f.Raw.Index,
			Removed:     false,
		},
	}
	return i
}

func (p postgres) CreateFrob(ctx context.Context, frob vat.VatFrob, logId uint64) (int64, error) {
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vat_frobs (i, u, v, w, dink, dart, log_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (log_id)
        DO UPDATE SET i = EXCLUDED.i, u = EXCLUDED.u, v = EXCLUDED.v, w = EXCLUDED.w, dink = EXCLUDED.dink, dart = EXCLUDED.dart
        RETURNING id
    `, common.Bytes2Hex(frob.I[:]), frob.U.String(), frob.V.String(), frob.W.String(), frob.Dink.String(), frob.Dart.String(), logId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert frob. %w", err)
	}
	return id, nil
}
