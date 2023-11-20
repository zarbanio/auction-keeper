package store

import (
	"context"
	"fmt"

	"github.com/zarbanio/auction-keeper/bindings/zarban/vow"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type flogModel struct {
	era string
	raw evmLogModel
}

func (f flogModel) toDomain() *vow.VowFlog {
	return &vow.VowFlog{
		Era: math.BigIntFromString(f.era),
		Raw: *f.raw.toRaw(),
	}
}

func (p postgres) CreateFlog(ctx context.Context, flog vow.VowFlog, logId uint64) (int64, error) {
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vow_flogs (era, log_id)
        VALUES ($1, $2)
        ON CONFLICT (log_id)
        DO UPDATE 
			SET era = EXCLUDED.era
        RETURNING id 
    `, flog.Era.String(), logId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert flogs. %w", err)
	}
	return id, nil
}
