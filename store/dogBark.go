package store

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
)

type barkModel struct {
	ilk string
	urn string
}

func (f barkModel) ToDomain() *entities.DogBark {
	var i [32]byte
	copy(i[:], f.ilk)

	return &entities.DogBark{
		Ilk: i,
		Urn: common.HexToAddress(f.urn),
	}
}

func NewBark(ilk [32]byte, urn common.Address) *barkModel {
	return &barkModel{
		ilk: string(ilk[:]),
		urn: urn.Hex(),
	}

}

func (p postgres) CreateBark(ctx context.Context, bark entities.DogBark, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
		INSERT INTO dog_barks (ilk, urn, tx_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (tx_id)
		DO UPDATE 
			SET ilk = EXCLUDED.ilk,
			urn = EXCLUDED.urn
		RETURNING id
	`, string(bark.Ilk[:]), bark.Urn.Hex(), tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
