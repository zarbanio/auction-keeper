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

type File0Handler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistryFile0]
}

func (h *File0Handler) ID() string {
	return h.addr + ":" + "0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d"
}

func (h *File0Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile0(log.Log)
}

func (h *File0Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistryFile0)
	if !ok {
		return errors.New("event type is not IlkregistryFile0")
	}
	return h.callback(log, e)
}

func (h *File0Handler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFile0(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFile0Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistryFile0]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File0Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
