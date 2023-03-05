package store

import (
	"context"
	"fmt"
	"math/big"

	"github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
)

type healModel struct {
	rad *big.Int
}

func (f healModel) ToDomain() *entities.VowHeal {

	return &entities.VowHeal{
		Rad: f.rad,
	}
}

func NewHeal(rad *big.Int) *healModel {
	return &healModel{
		rad: rad,
	}
}

func (p postgres) CreateHeal(ctx context.Context, heal *entities.VowHeal, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vow_heals (rad, tx_id)
        VALUES ($1, $2)
        ON CONFLICT (tx_id)
        DO UPDATE 
			SET rad = EXCLUDED.rad
        RETURNING id
    `, heal.Rad.Uint64(), tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
