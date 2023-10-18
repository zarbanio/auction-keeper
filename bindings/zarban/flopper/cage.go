// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package flopper

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type CageHandler struct {
	addr string
	binding  *Flopper
	callback events.CallbackFn[FlopperCage]
}

func (h *CageHandler) ID() string {
	return h.addr + ":" + "0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda"
}

func (h *CageHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseCage(log.Log)
}

func (h *CageHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(FlopperCage)
	if !ok {
		return errors.New("event type is not FlopperCage")
	}
	return h.callback(log, e)
}

func (h *CageHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseCage(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewCageHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[FlopperCage]) events.Handler {
	b, err := NewFlopper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &CageHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
