package store

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type forkModel struct {
	ilk  string
	src  string
	dst  string
	dink string
	dart string
	raw  evmLogModel
}

func (f forkModel) toDomain() *vat.VatFork {
	var I [32]byte
	copy(I[:], common.Hex2Bytes(f.ilk))
	return &vat.VatFork{
		Ilk:  I,
		Src:  common.HexToAddress(f.src),
		Dst:  common.HexToAddress(f.dst),
		Dink: math.BigIntFromString(f.dink),
		Dart: math.BigIntFromString(f.dart),
		Raw:  *f.raw.toRaw(),
	}
}

func (p postgres) CreateFork(ctx context.Context, fork vat.VatFork, logId uint64) (int64, error) {
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vat_forks (ilk, src, dst, dink, dart, log_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (log_id)
        DO UPDATE 
			SET i = EXCLUDED.i,
			u = EXCLUDED.u,
			rate = EXCLUDED.rate
        RETURNING id
    `, fork.Ilk, fork.Src.String(), fork.Dst.String(), fork.Dink.String(), fork.Dart.String(), logId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert fork. %w", err)
	}
	return id, nil
}
