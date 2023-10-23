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

type File0Handler struct {
	addr string
	binding  *Dog
	callback events.CallbackFn[DogFile0]
}

func (h *File0Handler) ID() string {
	return h.addr + ":" + "0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba"
}

func (h *File0Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile0(log.Log)
}

func (h *File0Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(DogFile0)
	if !ok {
		return errors.New("event type is not DogFile0")
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

func NewFile0Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[DogFile0]) events.Handler {
	b, err := NewDog(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File0Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
