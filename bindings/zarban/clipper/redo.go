// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package clipper

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type RedoHandler struct {
	addr string
	binding  *Clipper
	callback events.CallbackFn[ClipperRedo]
}

func (h *RedoHandler) ID() string {
	return h.addr + ":" + "0x275de7ecdd375b5e8049319f8b350686131c219dd4dc450a08e9cf83b03c865f"
}

func (h *RedoHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseRedo(log.Log)
}

func (h *RedoHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(ClipperRedo)
	if !ok {
		return errors.New("event type is not ClipperRedo")
	}
	return h.callback(log, e)
}

func (h *RedoHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseRedo(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewRedoHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperRedo]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RedoHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
