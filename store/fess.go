package store

import (
	"context"
	"fmt"

	"github.com/zarbanio/auction-keeper/bindings/zarban/vow"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type fessModel struct {
	tab string
	raw evmLogModel
}

func (f fessModel) toDomain() *vow.VowFess {
	return &vow.VowFess{
		Tab: math.BigIntFromString(f.tab),
		Raw: *f.raw.toRaw(),
	}
}

func (p postgres) CreateFess(ctx context.Context, fess vow.VowFess, logId uint64) (int64, error) {
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vow_fesses (tab, log_id)
        VALUES ($1, $2)
        ON CONFLICT (log_id)
        DO UPDATE 
			SET tab = EXCLUDED.tab
        RETURNING id 
    `, fess.Tab.String(), logId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert fess. %w", err)
	}
	return id, nil
}
