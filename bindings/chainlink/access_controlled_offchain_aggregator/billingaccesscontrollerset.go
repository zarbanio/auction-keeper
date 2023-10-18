// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package access_controlled_offchain_aggregator

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type BillingAccessControllerSetHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorBillingAccessControllerSet]
}

func (h *BillingAccessControllerSetHandler) ID() string {
	return h.addr + ":" + "0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912"
}

func (h *BillingAccessControllerSetHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseBillingAccessControllerSet(log.Log)
}

func (h *BillingAccessControllerSetHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorBillingAccessControllerSet)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorBillingAccessControllerSet")
	}
	return h.callback(log, e)
}

func (h *BillingAccessControllerSetHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseBillingAccessControllerSet(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewBillingAccessControllerSetHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorBillingAccessControllerSet]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &BillingAccessControllerSetHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
