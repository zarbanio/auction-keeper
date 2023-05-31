package cmd

import (
	"context"

	"github.com/zarbanio/auction-keeper/store"
)

type BlockPointer interface {
	Update(uint64) error
	Create() error
	Exists() bool
	Read() (uint64, error)
	Inc() error
}

type dbBlockPointer struct {
	s   store.IBlockPtr
	id  int64
	min uint64
}

func NewDBBlockPointer(s store.IStore, min uint64) BlockPointer {
	return &dbBlockPointer{s: s, min: min}
}

func (d *dbBlockPointer) Update(u uint64) error {
	_, err := d.s.UpdateBlockPtr(context.Background(), d.id, u)
	return err
}

func (d *dbBlockPointer) Create() error {
	id, err := d.s.CreateBlockPtr(context.Background(), d.min)
	if err != nil {
		return err
	}
	d.id = id
	return nil
}

func (d *dbBlockPointer) Exists() bool {
	exists, err := d.s.ExistsBlockPtr(context.Background())
	if err != nil {
		return false
	}
	id, _, err := d.s.GetLastBlockPtr(context.Background())
	if err != nil {
		return false
	}
	d.id = id
	return exists
}

func (d *dbBlockPointer) Read() (uint64, error) {
	return d.s.GetBlockPtrById(context.Background(), d.id)
}

func (d *dbBlockPointer) Inc() error {
	_, err := d.s.IncBlockPtr(context.Background(), d.id)
	return err
}
