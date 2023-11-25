package store

import (
	"context"
	"fmt"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/domain"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type vatGrabModel struct {
	I    string
	U    string
	V    string
	W    string
	Dink string
	Dart string
	Raw  evmLogModel
}

func (g vatGrabModel) toDomain() *vat.VatGrab {
	var ilk [32]byte
	copy(ilk[:], common.Hex2Bytes(g.I))
	vg := &vat.VatGrab{
		I:    ilk,
		U:    common.HexToAddress(g.U),
		V:    common.HexToAddress(g.V),
		W:    common.HexToAddress(g.W),
		Dink: math.BigIntFromString(g.Dink),
		Dart: math.BigIntFromString(g.Dart),
		Raw: types.Log{
			Address:     common.HexToAddress(g.Raw.Address),
			BlockNumber: g.Raw.BlockNumber,
			TxHash:      common.HexToHash(g.Raw.TxHash),
			BlockHash:   common.HexToHash(g.Raw.BlockHash),
			Index:       g.Raw.Index,
			Removed:     false,
		},
	}
	return vg
}

func (p postgres) CreateVatGrab(ctx context.Context, vatGrab vat.VatGrab, logId uint64) (int64, error) {
	ilk := domain.Bytes32ToString(vatGrab.I)
	if !utf8.ValidString(ilk) {
		return 0, ErrInvalidUTF8String
	}
	var id int64
	err := p.conn.QueryRow(ctx, `
		INSERT INTO grabs (i, u, v, w, dink, dart, log_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (log_id)
		DO UPDATE SET
			i = EXCLUDED.i,
			u = EXCLUDED.u,
			v = EXCLUDED.v,
			w = EXCLUDED.w,
			dink = EXCLUDED.dink,
			dart = EXCLUDED.dart
		RETURNING id;`,
		ilk, vatGrab.U.String(), vatGrab.V.String(), vatGrab.W.String(), vatGrab.Dink.String(), vatGrab.Dart.String(), logId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert grab. %w", err)
	}
	return id, nil
}
