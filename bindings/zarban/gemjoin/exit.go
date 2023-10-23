// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package gemjoin

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type ExitHandler struct {
	addr string
	binding  *Gemjoin
	callback events.CallbackFn[GemjoinExit]
}

func (h *ExitHandler) ID() string {
	return h.addr + ":" + "0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631"
}

func (h *ExitHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseExit(log.Log)
}

func (h *ExitHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(GemjoinExit)
	if !ok {
		return errors.New("event type is not GemjoinExit")
	}
	return h.callback(log, e)
}

func (h *ExitHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseExit(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewExitHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[GemjoinExit]) events.Handler {
	b, err := NewGemjoin(addr, eth)
	if err != nil {
		panic(err)
	}
	return &ExitHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
