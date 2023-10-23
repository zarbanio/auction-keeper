// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package dog

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type DigsHandler struct {
	addr string
	binding  *Dog
	callback events.CallbackFn[DogDigs]
}

func (h *DigsHandler) ID() string {
	return h.addr + ":" + "0x54f095dc7308776bf01e8580e4dd40fd959ea4bf50b069975768320ef8d77d8a"
}

func (h *DigsHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseDigs(log.Log)
}

func (h *DigsHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(DogDigs)
	if !ok {
		return errors.New("event type is not DogDigs")
	}
	return h.callback(log, e)
}

func (h *DigsHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseDigs(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewDigsHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[DogDigs]) events.Handler {
	b, err := NewDog(addr, eth)
	if err != nil {
		panic(err)
	}
	return &DigsHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
