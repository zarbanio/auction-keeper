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

type AddIlkHandler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistryAddIlk]
}

func (h *AddIlkHandler) ID() string {
	return h.addr + ":" + "0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7"
}

func (h *AddIlkHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseAddIlk(log.Log)
}

func (h *AddIlkHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistryAddIlk)
	if !ok {
		return errors.New("event type is not IlkregistryAddIlk")
	}
	return h.callback(log, e)
}

func (h *AddIlkHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseAddIlk(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewAddIlkHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistryAddIlk]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &AddIlkHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
