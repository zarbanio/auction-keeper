package store

import (
	"context"
	"fmt"
	"math/big"

	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type flogModel struct {
	era string
}

func (f flogModel) ToDomain() *entities.VowFlog {
	return &entities.VowFlog{
		Era: math.BigIntFromString(f.era),
	}
}

func NewFlog(era *big.Int) *flogModel {
	return &flogModel{
		era: era.String(),
	}
}

func (p postgres) CreateFlog(ctx context.Context, flog *entities.VowFlog, tx_id uint64) (int64, error) {
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vow_flogs (era, tx_id)
        VALUES ($1, $2)
        ON CONFLICT (tx_id)
        DO UPDATE 
			SET era = EXCLUDED.era
        RETURNING id 
    `, flog.Era.String(), tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert flogs. %w", err)
	}
	return id, nil
}
