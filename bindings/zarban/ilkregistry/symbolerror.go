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

type SymbolErrorHandler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistrySymbolError]
}

func (h *SymbolErrorHandler) ID() string {
	return h.addr + ":" + "0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d"
}

func (h *SymbolErrorHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseSymbolError(log.Log)
}

func (h *SymbolErrorHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistrySymbolError)
	if !ok {
		return errors.New("event type is not IlkregistrySymbolError")
	}
	return h.callback(log, e)
}

func (h *SymbolErrorHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseSymbolError(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewSymbolErrorHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistrySymbolError]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &SymbolErrorHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
