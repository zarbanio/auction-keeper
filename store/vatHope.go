package store

import (
	"context"
	"fmt"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
	"github.com/ethereum/go-ethereum/common"
)

type hopeModel struct {
	usr string
}

func (f hopeModel) ToDomain() *entities.VatHope {
	return &entities.VatHope{Usr: common.HexToAddress(f.usr)}
}

func NewHope(usr common.Address) *hopeModel {
	return &hopeModel{
		usr: usr.Hex(),
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
	`, hope.Usr.String(), tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
