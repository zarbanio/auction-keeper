package store

import (
	"context"
	"fmt"
	"math/big"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
	"github.com/IR-Digital-Token/auction-keeper/domain/math"
)

type kissModel struct {
	rad string
}

func (f kissModel) ToDomain() *entities.VowKiss {

	return &entities.VowKiss{
		Rad: math.BigIntFromString(f.rad),
	}
}

func NewKiss(rad *big.Int) *kissModel {
	return &kissModel{
		rad: rad.String(),
	}
}

func (p postgres) CreateKiss(ctx context.Context, kiss *entities.VowKiss, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vow_kisses (rad, tx_id)
        VALUES ($1, $2)
        ON CONFLICT (tx_id)
        DO UPDATE 
			SET rad = EXCLUDED.rad
        RETURNING id
    `, kiss.Rad, tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
