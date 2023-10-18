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

type File1Handler struct {
	addr string
	binding  *Dog
	callback events.CallbackFn[DogFile1]
}

func (h *File1Handler) ID() string {
	return h.addr + ":" + "0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1"
}

func (h *File1Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile1(log.Log)
}

func (h *File1Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(DogFile1)
	if !ok {
		return errors.New("event type is not DogFile1")
	}
	return h.callback(log, e)
}

func (h *File1Handler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFile1(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFile1Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[DogFile1]) events.Handler {
	b, err := NewDog(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File1Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
