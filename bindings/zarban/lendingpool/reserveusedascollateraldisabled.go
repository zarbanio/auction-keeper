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

type ReserveUsedAsCollateralDisabledHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolReserveUsedAsCollateralDisabled]
}

func (h *ReserveUsedAsCollateralDisabledHandler) ID() string {
	return h.addr + ":" + "0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd"
}

func (h *ReserveUsedAsCollateralDisabledHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseReserveUsedAsCollateralDisabled(log.Log)
}

func (h *ReserveUsedAsCollateralDisabledHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolReserveUsedAsCollateralDisabled)
	if !ok {
		return errors.New("event type is not LendingpoolReserveUsedAsCollateralDisabled")
	}
	return h.callback(log, e)
}

func (h *ReserveUsedAsCollateralDisabledHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseReserveUsedAsCollateralDisabled(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewReserveUsedAsCollateralDisabledHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolReserveUsedAsCollateralDisabled]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &ReserveUsedAsCollateralDisabledHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
