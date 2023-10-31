package store

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx"
	"github.com/zarbanio/auction-keeper/services/transaction"
)

type TransactionModel struct {
	id    uint64
	hash  string
	from  string
	block uint64
}

func (p postgres) CreateTransaction(ctx context.Context, tx *types.Transaction) (error, uint64) {
	from, err := transaction.SenderAddressFromTransaction(tx)
	if err != nil {
		return err, 0
	}
	var id uint64
	v, r, s := tx.RawSignatureValues()
	q := `
	INSERT INTO transactions (hash, from_address, cost, data, value, gas_price, gas, nonce, to_address, chain_id, v, r, s)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	ON CONFLICT (hash)
	DO UPDATE 
		SET from_address = EXCLUDED.from_address,
		cost = EXCLUDED.cost,
		data = EXCLUDED.data,
		value = EXCLUDED.value,
		gas_price = EXCLUDED.gas_price,
		gas = EXCLUDED.gas,
		nonce = EXCLUDED.nonce,
		to_address = EXCLUDED.to_address,
		chain_id = EXCLUDED.chain_id,
		v = EXCLUDED.v,
		r = EXCLUDED.r,
		s = EXCLUDED.s
	RETURNING id
	`
	tx.ChainId()
	err = p.conn.QueryRow(ctx, q,
		tx.Hash().String(),
		from.String(),
		tx.Cost().Int64(),
		hex.EncodeToString(tx.Data()),
		tx.Value().Int64(),
		tx.GasPrice().Int64(),
		tx.Gas(),
		tx.Nonce(),
		tx.To().String(),
		tx.ChainId().Int64(),
		v.String(),
		r.String(),
		s.String()).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err), 0
	}
	return nil, id
}
func (p postgres) UpdateTransactionReceipt(ctx context.Context, id uint64, recipt *types.Receipt, header *types.Header) error {
	q := `UPDATE transactions 
			SET block_timestamp = $1,
				block_number = $2, 
				status = $3, 
				cumulative_gas_used = $4, 
				block_hash = $5 
			WHERE id = $6`
	err := p.conn.QueryRow(ctx, q,
		header.Time,
		recipt.BlockNumber,
		recipt.Status,
		recipt.CumulativeGasUsed,
		header.Hash().String(),
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to update block data in transaction: %w", err)
	}
	return nil
}

func (p postgres) GetTransactionById(ctx context.Context, id uint64) (*TransactionModel, error) {
	row := p.conn.QueryRow(ctx, `
		SELECT *
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
