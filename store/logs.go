package store

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type logModel struct {
	Id          uint64
	Address     string
	Timestamp   int64
	BlockNumber uint64
	TxHash      string
	TxIndex     uint
	BlockHash   string
	Index       uint
	Data        []byte
	Topics      []string
}

func (l *logModel) toDomain() *eth.Log {
	return &eth.Log{
		Id:        l.Id,
		Log:       *l.toRaw(),
		Timestamp: time.Unix(l.Timestamp, 0),
	}
}

func (l *logModel) toRaw() *types.Log {
	topics := make([]common.Hash, len(l.Topics))
	for i, t := range l.Topics {
		topics[i] = common.HexToHash(t)
	}
	log := types.Log{
		Address:     common.HexToAddress(l.Address),
		BlockNumber: l.BlockNumber,
		TxHash:      common.HexToHash(l.TxHash),
		TxIndex:     l.TxIndex,
		BlockHash:   common.HexToHash(l.BlockHash),
		Index:       l.Index,
		Removed:     false,
		Data:        l.Data,
		Topics:      topics,
	}
	return &log
}

func (p postgres) CreateLog(ctx context.Context, log eth.Log) (uint64, error) {
	var id uint64
	topics := make([]string, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = t.String()
	}
	err := p.conn.QueryRow(ctx, `
        INSERT INTO logs (address, timestamp, block_number, tx_hash, tx_index, block_hash, index, data, topics)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (tx_hash, index)
        DO UPDATE SET
		 	address = EXCLUDED.address,
			block_number = EXCLUDED.block_number,
			tx_hash = EXCLUDED.tx_hash,
			tx_index = EXCLUDED.tx_index,
			block_hash = EXCLUDED.block_hash,
			index = EXCLUDED.index,
			data = EXCLUDED.data,
			topics = EXCLUDED.topics
        RETURNING id
    `,
		log.Address.String(),
		log.Timestamp.Unix(),
		log.BlockNumber,
		log.TxHash.String(),
		log.TxIndex,
		log.BlockHash.String(),
		log.Index,
		log.Data,
		topics,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert log. %w", err)
	}
	return id, nil
}

func (p postgres) GetLogById(ctx context.Context, id uint64) (*eth.Log, error) {
	row := p.conn.QueryRow(ctx, `
		SELECT address, timestamp, block_number, tx_hash, tx_index, block_hash, index, data, topics
		FROM logs
		WHERE id = $1
	`, id)

	var log logModel
	err := row.Scan(
		&log.Address,
		&log.Timestamp,
		&log.BlockNumber,
		&log.TxHash,
		&log.TxIndex,
		&log.BlockHash,
		&log.Index,
		&log.Data,
		&log.Topics,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrLogNotFound
		}
		return nil, fmt.Errorf("failed to get log. %w", err)
	}
	return log.toDomain(), nil
}

func (p postgres) GetLogsByQuery(ctx context.Context, q EthereumLogsQuery) ([]eth.Log, error) {
	var logs []eth.Log

	addressStrings := make([]string, len(q.Addresses))
	for i, addr := range q.Addresses {
		addressStrings[i] = addr.Hex()
	}

	query := `
		SELECT
			id, timestamp, address, block_number, tx_hash,
			tx_index, block_hash, index, data, topics
		FROM
			logs
		WHERE
			block_number BETWEEN $1 AND $2
			AND address = ANY($3::TEXT[])
	`

	rows, err := p.conn.Query(ctx, query, q.FromBlock, q.ToBlock, addressStrings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var log logModel
		if err := rows.Scan(
			&log.Id,
			&log.Timestamp,
			&log.Address,
			&log.BlockNumber,
			&log.TxHash,
			&log.TxIndex,
			&log.BlockHash,
			&log.Index,
			&log.Data,
			&log.Topics); err != nil {
			return nil, err
		}
		logs = append(logs, *log.toDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
