package store

import (
	"context"
	"fmt"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
	"github.com/ethereum/go-ethereum/common"
)

type dogModel struct {
	ilk string
	urn string
}

func (f dogModel) ToDomain() *entities.DogBark {
	var i [32]byte
	copy(i[:], common.Hex2Bytes(f.ilk))
	return &entities.DogBark{
		Ilk: i,
		Urn: common.HexToAddress(f.urn),
	}
}

func NewDog(ilk [32]byte, urn common.Address) *dogModel {
	return &dogModel{
		ilk: string(ilk[:]),
		urn: urn.String(),
	}
}

func (p postgres) Createdog(ctx context.Context, dog entities.DogBark, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
		INSERT INTO clipper_dogs (ilk, urn, tx_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (tx_id)
		DO UPDATE 
			SET ilk = EXCLUDED.ilk
			urn = EXCLUDED.urn
		RETURNING id
	`, string(dog.Ilk[:]), dog.Urn.String(), tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
