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
	hash  string
	from  common.Address
	block uint64
}

// id      			BIGINT PRIMARY KEY,
// hash    			TEXT,
// from    			TEXT,
// time    			timestamp default current_timestamp,
// block   			BIGINT
// cost 				TEXT,
// data 				TEXT,
// value 				TEXT,
// gas_price			TEXT,
// gas 				TEXT,
// nonce 				TEXT,
// to 					TEXT,
// RawSignatureValues 	TEXT

func (p postgres) CreateTransaction(ctx context.Context, transaction *types.Transaction, from common.Address) (error, uint) {
	var id uint
	v, r, s := transaction.RawSignatureValues()
	q := `
	INSERT INTO transactions (hash, from, cost, data, value, gas_price, gas, nonce, to, v, r, s)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	ON CONFLICT (hash)
	DO UPDATE 
		SET from = EXCLUDED.from
		block = EXCLUDED.block
	RETURNING id
	`
	err := p.conn.QueryRow(ctx, q, transaction.Hash(),
		from.String(),
		transaction.Cost(),
		string(transaction.Data()),
		transaction.Value(),
		transaction.GasPrice(),
		transaction.Gas(),
		transaction.Nonce(),
		transaction.To(),
		v.String(),
		r.String(),
		s.String()).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err), 0
	}
	return nil, id
}
func (p postgres) UpdateTransactionBlock(ctx context.Context, id uint64, block *types.Block, recipt *types.Receipt) error {
	q := `UPDATE transactions 
			SET block_timestamp = $1,
				block_number = $2, 
				status = $3, 
				cumulative_gas_used = $4, 
				block_hash = $5, 

			WHERE id = $6`
	err := p.conn.QueryRow(ctx, q, block.Time(), block.Number(), id)
	test := recipt.Status
	if err != nil {
		return fmt.Errorf("failed to update block data in transaction: %w", err)
	}
	return nil
}

func (p postgres) GetTransactionById(ctx context.Context, id uint64) (*TransactionModel, error) {
	row := p.conn.QueryRow(ctx, `
		SELECT id, hash, from, block
		FROM transactions
		WHERE id = $1
	`, id)

	var transaction *TransactionModel
	err := row.Scan(&transaction.id, &transaction.hash, &transaction.from, &transaction.block)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrTransactionNotFound
		}
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}
	return transaction, nil
}
