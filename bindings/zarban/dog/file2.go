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

type File2Handler struct {
	addr string
	binding  *Dog
	callback events.CallbackFn[DogFile2]
}

func (h *File2Handler) ID() string {
	return h.addr + ":" + "0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d"
}

func (h *File2Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile2(log.Log)
}

func (h *File2Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(DogFile2)
	if !ok {
		return errors.New("event type is not DogFile2")
	}
	return h.callback(log, e)
}

func (h *File2Handler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseFile2(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewFile2Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[DogFile2]) events.Handler {
	b, err := NewDog(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File2Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
