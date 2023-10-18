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

type RemovedAccessHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorRemovedAccess]
}

func (h *RemovedAccessHandler) ID() string {
	return h.addr + ":" + "0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1"
}

func (h *RemovedAccessHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseRemovedAccess(log.Log)
}

func (h *RemovedAccessHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorRemovedAccess)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorRemovedAccess")
	}
	return h.callback(log, e)
}

func (h *RemovedAccessHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseRemovedAccess(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewRemovedAccessHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorRemovedAccess]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RemovedAccessHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
