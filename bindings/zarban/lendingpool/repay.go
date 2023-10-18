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

type RepayHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolRepay]
}

func (h *RepayHandler) ID() string {
	return h.addr + ":" + "0x4cdde6e09bb755c9a5589ebaec640bbfedff1362d4b255ebf8339782b9942faa"
}

func (h *RepayHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseRepay(log.Log)
}

func (h *RepayHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolRepay)
	if !ok {
		return errors.New("event type is not LendingpoolRepay")
	}
	return h.callback(log, e)
}

func (h *RepayHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseRepay(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewRepayHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolRepay]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RepayHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
