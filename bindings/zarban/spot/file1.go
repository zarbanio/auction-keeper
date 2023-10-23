// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package spot

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type File1Handler struct {
	addr string
	binding  *Spot
	callback events.CallbackFn[SpotFile1]
}

func (h *File1Handler) ID() string {
	return h.addr + ":" + "0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7"
}

func (h *File1Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseFile1(log.Log)
}

func (h *File1Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(SpotFile1)
	if !ok {
		return errors.New("event type is not SpotFile1")
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

func NewFile1Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[SpotFile1]) events.Handler {
	b, err := NewSpot(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File1Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
