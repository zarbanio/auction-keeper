package store

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx"
)

type TransactionModel struct {
	id    uint64
	time  uint64
	hash  string
	from  common.Address
	block uint64
}

func (p postgres) CreateTransaction(ctx context.Context, transaction *types.Transaction, from common.Address) error {
	_, err := p.conn.Exec(ctx, `
		INSERT INTO transactions (time, hash, from, block)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (hash)
		DO UPDATE SET time = EXCLUDED.time, from = EXCLUDED.from
	`, 1222222, transaction.Hash(), from.String())
	if err != nil {
		return fmt.Errorf("failed to upsert transaction. %w", err)
	}
	return nil
}

func (p postgres) GetTransactionById(ctx context.Context, id uint64) (*TransactionModel, error) {
	row := p.conn.QueryRow(ctx, `
		SELECT id, time, hash, from
		FROM transactions
		WHERE id = $1
	`, id)

	var transaction *TransactionModel
	err := row.Scan(&transaction.id, &transaction.time, &transaction.hash, &transaction.from)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrTransactionNotFound
		}
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}
	return transaction, nil
}
