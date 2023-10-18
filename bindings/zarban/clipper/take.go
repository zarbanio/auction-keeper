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

type TakeHandler struct {
	addr string
	binding  *Clipper
	callback events.CallbackFn[ClipperTake]
}

func (h *TakeHandler) ID() string {
	return h.addr + ":" + "0x05e309fd6ce72f2ab888a20056bb4210df08daed86f21f95053deb19964d86b1"
}

func (h *TakeHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseTake(log.Log)
}

func (h *TakeHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(ClipperTake)
	if !ok {
		return errors.New("event type is not ClipperTake")
	}
	return h.callback(log, e)
}

func (h *TakeHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseTake(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewTakeHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperTake]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &TakeHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
