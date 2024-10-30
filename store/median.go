package store

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	"github.com/zarbanio/auction-keeper/bindings/zarban/median"
	"github.com/zarbanio/auction-keeper/domain/math"
)

type logMedianPriceModel struct {
	val string
	age int64
	raw evmLogModel
}

func (l *logMedianPriceModel) toDomain() *median.MedianLogMedianPrice {
	return &median.MedianLogMedianPrice{
		Val: math.BigIntFromString(l.val),
		Age: big.NewInt(l.age),
		Raw: *l.raw.toRaw(),
	}
}

func (p postgres) CreateLogMedianPrice(ctx context.Context, logmedianPrice median.MedianLogMedianPrice, logId uint64) (int64, error) {
	query := `
		INSERT INTO median_prices (val, age, log_id) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (log_id)
		DO UPDATE SET val = excluded.val, age = excluded.age
		RETURNING id
	`
	var id int64
	err := p.conn.QueryRow(ctx, query, logmedianPrice.Val.String(), logmedianPrice.Age.String(), logId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p postgres) GetLogMedianPriceById(ctx context.Context, id int64) (*median.MedianLogMedianPrice, error) {
	row := p.conn.QueryRow(ctx, `
		SELECT val, age, address, block_number, tx_hash, block_hash, index
		FROM median_prices join evm_logs l on l.id = median_prices.log_id
		WHERE median_prices.id = $1
	`, id)

	var m logMedianPriceModel
	err := row.Scan(&m.val, &m.age,
		&m.raw.Address, &m.raw.BlockNumber, &m.raw.TxHash, &m.raw.BlockHash, &m.raw.Index)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrMedianLogPriceMedianNotFound
		}
		return nil, fmt.Errorf("failed to get logMedianPrice. %w", err)
	}
	return m.toDomain(), nil
}

func (p postgres) GetLogMedianPriceByOrder(ctx context.Context, address common.Address, cursor, limit int64) ([]median.MedianLogMedianPrice, error) {
	rows, err := p.conn.Query(ctx, `
		SELECT val, age, address, block_number, tx_hash, block_hash, index
		FROM median_prices join evm_logs l on l.id = median_prices.log_id
		WHERE l.address = $1 AND median_prices.id > $2
		ORDER BY l.block_number DESC
		LIMIT $3
	`, address.Hex(), cursor, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var arr []median.MedianLogMedianPrice
	for rows.Next() {
		var m logMedianPriceModel
		err = rows.Scan(
			&m.val, &m.age,
			&m.raw.Address, &m.raw.BlockNumber, &m.raw.TxHash, &m.raw.BlockHash, &m.raw.Index)

		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, ErrMedianLogPriceMedianNotFound
			}
			return nil, fmt.Errorf("failed to get logMedianPrice. %w", err)
		}

		arr = append(arr, *m.toDomain())
	}
	return arr, nil
}

func (p postgres) GetLastLogMedianPrice(ctx context.Context, address common.Address) (*median.MedianLogMedianPrice, error) {
	row := p.conn.QueryRow(ctx, `
		SELECT val, age, address, block_number, tx_hash, block_hash, index
		FROM median_prices join evm_logs l on l.id = median_prices.log_id
		WHERE l.address = $1
		ORDER BY l.block_number DESC
		LIMIT 1
	`, address.Hex())

	var m logMedianPriceModel
	err := row.Scan(&m.val, &m.age,
		&m.raw.Address, &m.raw.BlockNumber, &m.raw.TxHash, &m.raw.BlockHash, &m.raw.Index)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrMedianLogPriceMedianNotFound
		}
		return nil, fmt.Errorf("failed to get logMedianPrice. %w", err)
	}
	return m.toDomain(), nil
}
