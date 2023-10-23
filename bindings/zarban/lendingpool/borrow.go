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

type BorrowHandler struct {
	addr string
	binding  *Lendingpool
	callback events.CallbackFn[LendingpoolBorrow]
}

func (h *BorrowHandler) ID() string {
	return h.addr + ":" + "0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b"
}

func (h *BorrowHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseBorrow(log.Log)
}

func (h *BorrowHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(LendingpoolBorrow)
	if !ok {
		return errors.New("event type is not LendingpoolBorrow")
	}
	return h.callback(log, e)
}

func (h *BorrowHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseBorrow(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewBorrowHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[LendingpoolBorrow]) events.Handler {
	b, err := NewLendingpool(addr, eth)
	if err != nil {
		panic(err)
	}
	return &BorrowHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
