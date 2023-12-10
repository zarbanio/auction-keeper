package indexer

import (
	"context"
	"log"

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
	s    store.IBlockPtr
	name string
	min  uint64
}

func NewDBBlockPointer(s store.IBlockPtr, name string, min uint64) BlockPointer {
	return &dbBlockPointer{s: s, name: name, min: min}
}

func (d *dbBlockPointer) Update(u uint64) error {
	_, err := d.s.UpdateBlockPtr(context.Background(), d.name, u)
	return err
}

func (d *dbBlockPointer) Create() error {
	_, err := d.s.CreateBlockPtr(context.Background(), d.name, d.min)
	if err != nil {
		return err
	}
	return nil
}

func (d *dbBlockPointer) Exists() bool {
	exists, err := d.s.ExistsBlockPtr(context.Background(), d.name)
	if err != nil {
		return false
	}
	return exists
}

func (d *dbBlockPointer) Read() (uint64, error) {
	return d.s.GetBlockPtrByName(context.Background(), d.name)
}

func (d *dbBlockPointer) Inc() error {
	_, err := d.s.IncBlockPtr(context.Background(), d.name)
	return err
}

func CreateBlockPtrIfNotExists(s store.IBlockPtr, name string, min uint64) BlockPointer {
	blockPtr := NewDBBlockPointer(s, name, min)
	if !blockPtr.Exists() {
		log.Println(name, "block pointer doest not exits. creating a new one.")
		err := blockPtr.Create()
		if err != nil {
			log.Fatal("error creating history block pointer.", err)
		}
		log.Printf("new history block pointer created %d.\n", min)
	}
	return blockPtr
}
