package store

import (
	"context"
	"fmt"
	"math/big"

	"github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
)

type flogModel struct {
	era *big.Int
}

func (f flogModel) ToDomain() *entities.VowFlog {

	return &entities.VowFlog{
		Era: f.era,
	}
}

func NewFlog(era *big.Int) *flogModel {
	return &flogModel{
		era: era,
	}
}

func (p postgres) CreateFlog(ctx context.Context, flog *entities.VowFlog, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vow_flogs (era, tx_id)
        VALUES ($1, $2)
        ON CONFLICT (tx_id)
        DO UPDATE 
			SET era = EXCLUDED.era
        RETURNING id
    `, flog.Era.Uint64(), tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
