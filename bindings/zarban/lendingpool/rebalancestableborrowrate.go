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

type RebalanceStableBorrowRateHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolRebalanceStableBorrowRate]
}

func (h *RebalanceStableBorrowRateHandler) ID() string {
	return h.addr + ":" + "0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300"
}

func (h *RebalanceStableBorrowRateHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseRebalanceStableBorrowRate(log.Log)
}

func (h *RebalanceStableBorrowRateHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolRebalanceStableBorrowRate)
	if !ok {
		return errors.New("event type is not LendingpoolRebalanceStableBorrowRate")
	}
	return h.callback(log, e)
}

func (h *RebalanceStableBorrowRateHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseRebalanceStableBorrowRate(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewRebalanceStableBorrowRateHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolRebalanceStableBorrowRate]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RebalanceStableBorrowRateHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
