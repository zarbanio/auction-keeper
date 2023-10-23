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

type BillingSetHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorBillingSet]
}

func (h *BillingSetHandler) ID() string {
	return h.addr + ":" + "0xd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b6"
}

func (h *BillingSetHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseBillingSet(log.Log)
}

func (h *BillingSetHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorBillingSet)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorBillingSet")
	}
	return h.callback(log, e)
}

func (h *BillingSetHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseBillingSet(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewBillingSetHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorBillingSet]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &BillingSetHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
