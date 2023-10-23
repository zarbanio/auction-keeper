package store

import (
	"context"

	"github.com/jackc/pgx/v4"

	"errors"
)

func (p postgres) ExistsBlockPtr(ctx context.Context, name string) (bool, error) {
	var exists bool
	err := p.conn.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM block_ptrs WHERE name = $1)", name).Scan(&exists)
	return exists, err
}

func (p postgres) CreateBlockPtr(ctx context.Context, name string, ptr uint64) (int64, error) {
	var id int64
	err := p.conn.QueryRow(ctx, "INSERT INTO block_ptrs (ptr, name) VALUES ($1, $2) RETURNING id", ptr, name).Scan(&id)
	return id, err
}

func (p postgres) GetBlockPtrByName(ctx context.Context, name string) (uint64, error) {
	var ptr uint64
	err := p.conn.QueryRow(ctx, "SELECT ptr FROM block_ptrs WHERE name = $1", name).Scan(&ptr)
	if err == pgx.ErrNoRows {
		return 0, errors.New("block_ptr not found")
	}
	return ptr, err
}

func (p postgres) GetBlockPtrById(ctx context.Context, id int64) (uint64, error) {
	var ptr uint64
	err := p.conn.QueryRow(ctx, "SELECT ptr FROM block_ptrs WHERE id = $1", id).Scan(&ptr)
	if err == pgx.ErrNoRows {
		return 0, errors.New("block_ptr not found")
	}
	return ptr, err
}

func (p postgres) UpdateBlockPtr(ctx context.Context, name string, ptr uint64) (uint64, error) {
	tag, err := p.conn.Exec(ctx, "UPDATE block_ptrs SET ptr = $1 WHERE name = $2", ptr, name)
	if err != nil {
		return 0, err
	}
	if tag.RowsAffected() == 0 {
		return 0, errors.New("block_ptr not found")
	}
	return ptr, nil
}

func (p postgres) IncBlockPtr(ctx context.Context, name string) (uint64, error) {
	var ptr uint64
	err := p.conn.QueryRow(ctx, `UPDATE block_ptrs SET ptr = ptr + 1 WHERE name = $1 RETURNING ptr`, name).Scan(&ptr)
	if err == pgx.ErrNoRows {
		return 0, errors.New("block_ptr not found")
	}
	return ptr, err
}

func (p postgres) DeleteBlockPtr(ctx context.Context, name string) error {
	tag, err := p.conn.Exec(ctx, "DELETE FROM block_ptrs WHERE name = $1", name)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("block_ptr not found")
	}
	return nil
}

func (p postgres) GetLastBlockPtr(ctx context.Context) (int64, uint64, error) {
	var ptr uint64
	var id int64
	err := p.conn.QueryRow(ctx, "SELECT id, ptr FROM block_ptrs ORDER BY id DESC LIMIT 1").Scan(&id, &ptr)
	if err == pgx.ErrNoRows {
		return 0, 0, errors.New("no block_ptrs found")
	}
	return id, ptr, err
}
