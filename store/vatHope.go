package store

import (
	"context"
	"fmt"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/ethereum/go-ethereum/common"
)

type hopeModel struct {
	usr common.Address
}

func (f hopeModel) ToDomain() *entities.VatHope {
	return &entities.VatHope{Usr: f.usr}
}

func NewHope(usr common.Address) *hopeModel {
	return &hopeModel{
		usr: usr,
	}

}

func (p postgres) CreateHope(ctx context.Context, hope entities.VatHope, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
		INSERT INTO vat_hopes (usr, tx_id)
		VALUES ($1, $2)
		ON CONFLICT (tx_id)
		DO UPDATE 
			SET usr = EXCLUDED.usr
		RETURNING id
	`, hope.Usr, tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
