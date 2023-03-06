package store

import (
	"context"
	"fmt"
	"math/big"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities/inputMethods"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type takeModel struct {
	auction_id  *big.Int
	amt         *big.Int
	max         *big.Int
	who         common.Address
	data        []byte
	opts        *bind.TransactOpts
	transaction types.Transaction
}

func (f takeModel) ToDomain() *entities.ClipperTake {
	return &entities.ClipperTake{
		Auction_id: f.auction_id,
		Amt:        f.amt,
		Max:        f.max,
		Who:        f.who,
		Data:       f.data,
	}
}

func NewTake(id *big.Int, amt *big.Int, max *big.Int, who common.Address, data []byte) *takeModel {
	return &takeModel{
		auction_id: id,
		amt:        amt,
		max:        max,
		who:        who,
		data:       data,
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
