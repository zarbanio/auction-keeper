// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package lendingpool

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type UnpausedHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolUnpaused]
}

func (h *UnpausedHandler) ID() string {
	return h.addr + ":" + "0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933"
}

func (h *UnpausedHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseUnpaused(log.Log)
}

func (h *UnpausedHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolUnpaused)
	if !ok {
		return errors.New("event type is not LendingpoolUnpaused")
	}
	return h.callback(log, e)
}

func (h *UnpausedHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseUnpaused(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewUnpausedHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolUnpaused]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &UnpausedHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
