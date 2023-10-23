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

type BarkHandler struct {
	addr string
	binding  *Dog
	callback events.CallbackFn[DogBark]
}

func (h *BarkHandler) ID() string {
	return h.addr + ":" + "0x85258d09e1e4ef299ff3fc11e74af99563f022d21f3f940db982229dc2a3358c"
}

func (h *BarkHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseBark(log.Log)
}

func (h *BarkHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(DogBark)
	if !ok {
		return errors.New("event type is not DogBark")
	}
	return h.callback(log, e)
}

func (h *BarkHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseBark(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewBarkHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[DogBark]) events.Handler {
	b, err := NewDog(addr, eth)
	if err != nil {
		panic(err)
	}
	return &BarkHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
