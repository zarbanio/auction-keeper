package store

import (
	"context"
	"fmt"
	"math/big"

	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type healModel struct {
	rad string
}

func (f healModel) ToDomain() *entities.VowHeal {

	return &entities.VowHeal{
		Rad: math.BigIntFromString(f.rad),
	}
}

func NewHeal(rad *big.Int) *healModel {
	return &healModel{
		rad: rad.String(),
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
    `, heal.Rad.String(), tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
