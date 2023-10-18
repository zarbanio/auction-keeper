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

type RemoveIlkHandler struct {
	addr string
	binding  *Ilkregistry
	callback events.CallbackFn[IlkregistryRemoveIlk]
}

func (h *RemoveIlkHandler) ID() string {
	return h.addr + ":" + "0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796"
}

func (h *RemoveIlkHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseRemoveIlk(log.Log)
}

func (h *RemoveIlkHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(IlkregistryRemoveIlk)
	if !ok {
		return errors.New("event type is not IlkregistryRemoveIlk")
	}
	return h.callback(log, e)
}

func (h *RemoveIlkHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseRemoveIlk(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewRemoveIlkHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[IlkregistryRemoveIlk]) events.Handler {
	b, err := NewIlkregistry(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RemoveIlkHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
