package store

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type takeModel struct {
	auction_id string
	amt        string
	max        string
	who        string
	data       string
}

func (f takeModel) ToDomain() *entities.ClipperTake {
	return &entities.ClipperTake{
		Auction_id: math.BigIntFromString(f.auction_id),
		Amt:        math.BigIntFromString(f.amt),
		Max:        math.BigIntFromString(f.max),
		Who:        common.HexToAddress(f.who),
		Data:       []byte(f.data),
	}
}

func NewTake(id *big.Int, amt *big.Int, max *big.Int, who common.Address, data []byte) *takeModel {
	return &takeModel{
		auction_id: id.String(),
		amt:        amt.String(),
		max:        max.String(),
		who:        who.String(),
		data:       string(data),
	}
}

func (p postgres) CreateTake(ctx context.Context, take *entities.ClipperTake, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO clipper_takes (auction_id, amt, max, who, data, tx_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (tx_id)
        DO UPDATE 
			SET auction_id = EXCLUDED.auction_id
			amt = EXCLUDED.amt,
			max = EXCLUDED.max,
			who = EXCLUDED.who,
			data = EXCLUDED.data
        RETURNING id
    `, take.Auction_id, take.Amt, take.Max, take.Who, take.Data, tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
