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

type ReserveDataUpdatedHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolReserveDataUpdated]
}

func (h *ReserveDataUpdatedHandler) ID() string {
	return h.addr + ":" + "0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"
}

func (h *ReserveDataUpdatedHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseReserveDataUpdated(log.Log)
}

func (h *ReserveDataUpdatedHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolReserveDataUpdated)
	if !ok {
		return errors.New("event type is not LendingpoolReserveDataUpdated")
	}
	return h.callback(log, e)
}

func (h *ReserveDataUpdatedHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseReserveDataUpdated(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewReserveDataUpdatedHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolReserveDataUpdated]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &ReserveDataUpdatedHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
