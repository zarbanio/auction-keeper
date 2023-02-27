package store

import (
	"context"
	"fmt"
)

func (p postgres) CreateFlop(ctx context.Context, tx_id int64) (int64, error) {
	// rest of the code here
	var id int64
	err := p.conn.QueryRow(ctx, `
        INSERT INTO vow_flops (tx_id)
        VALUES ($1)
        ON CONFLICT (tx_id)
        RETURNING id
    `, tx_id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert . %w", err)
	}
	return id, nil
}
