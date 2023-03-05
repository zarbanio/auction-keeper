package store

import (
	"context"
	"fmt"
	"math/big"

	"github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
)

type redoModel struct {
	sale_id *big.Int
}

func (f redoModel) ToDomain() *entities.ClipperRedo {
	return &entities.ClipperRedo{
		SailId: f.sale_id,
	}
}

func NewRedo(saleId big.Int) *redoModel {
	return &redoModel{
		sale_id: &saleId,
	}

}

func (p postgres) CreateRedo(ctx context.Context, redo entities.ClipperRedo, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
		INSERT INTO clipper_redoes (sale_id, tx_id)
		VALUES ($1, $2)
		ON CONFLICT (tx_id)
		DO UPDATE 
			SET sale_id = EXCLUDED.sale_id
		RETURNING id
	`, redo.SailId.Int64(), tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
