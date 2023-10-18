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

type ReserveUsedAsCollateralEnabledHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolReserveUsedAsCollateralEnabled]
}

func (h *ReserveUsedAsCollateralEnabledHandler) ID() string {
	return h.addr + ":" + "0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2"
}

func (h *ReserveUsedAsCollateralEnabledHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseReserveUsedAsCollateralEnabled(log.Log)
}

func (h *ReserveUsedAsCollateralEnabledHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolReserveUsedAsCollateralEnabled)
	if !ok {
		return errors.New("event type is not LendingpoolReserveUsedAsCollateralEnabled")
	}
	return h.callback(log, e)
}

func (h *ReserveUsedAsCollateralEnabledHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseReserveUsedAsCollateralEnabled(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewReserveUsedAsCollateralEnabledHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolReserveUsedAsCollateralEnabled]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &ReserveUsedAsCollateralEnabledHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
