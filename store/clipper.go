package store

import (
	"context"
	"fmt"

	entities "github.com/IR-Digital-Token/auction-keeper/domain/entities"
)

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
	// common.Bytes2Hex(init.Ilk[:]), logId
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
