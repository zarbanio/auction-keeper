package store

import (
	"context"
	"sort"
)

func (e postgres) CreateEthereumLogsCacheKey(ctx context.Context, key EthereumLogsQuery) error {
	addresses := make([]string, len(key.Addresses))
	for i, a := range key.Addresses {
		addresses[i] = a.Hex()
	}
	sort.Strings(addresses)

	sql := `
	WITH overlapping_keys AS (
		SELECT id, from_block, to_block
		FROM logs_cache_keys
		WHERE
			(from_block BETWEEN $1 AND $2 OR
			to_block BETWEEN $1 AND $2 OR
			($1 BETWEEN from_block AND to_block AND
			$2 BETWEEN from_block AND to_block))
			AND addresses::TEXT[] @> $3::TEXT[]
	)
	, merged_range AS (
		SELECT
			LEAST($1, MIN(from_block)) AS new_from_block,
			GREATEST($2, MAX(to_block)) AS new_to_block
		FROM overlapping_keys
	)
	, deleted_keys AS (
		DELETE FROM logs_cache_keys
		WHERE id IN (SELECT id FROM overlapping_keys)
		RETURNING id
	)
	INSERT INTO logs_cache_keys (from_block, to_block, addresses)
	SELECT new_from_block, new_to_block, $3::TEXT[]
	FROM merged_range;
	`

	_, err := e.conn.Exec(ctx, sql, key.FromBlock, key.ToBlock, addresses)
	if err != nil {
		return err
	}

	return nil
}

func (p postgres) IsEthereumLogsCached(ctx context.Context, key EthereumLogsQuery) (bool, error) {
	addresses := make([]string, len(key.Addresses))
	for i, a := range key.Addresses {
		addresses[i] = a.Hex()
	}
	sort.Strings(addresses)

	var cacheHit bool
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM logs_cache_keys
            WHERE (from_block <= $1 AND to_block >= $2)
            AND addresses::TEXT[] @> $3::TEXT[]
        )
    `
	err := p.conn.QueryRow(ctx, query, key.FromBlock, key.ToBlock, addresses).Scan(&cacheHit)
	if err != nil {
		return false, err
	}

	return cacheHit, nil
}
