// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package ilkregistry

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type File2Handler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistryFile2]
}

func (h *File2Handler) ID() string {
	return h.addr + ":" + "0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8"
}

func (h *File2Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile2(log.Log)
}

func (h *File2Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistryFile2)
	if !ok {
		return errors.New("event type is not IlkregistryFile2")
	}
	return h.callback(log, e)
}

func (h *File2Handler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFile2(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFile2Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistryFile2]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File2Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
