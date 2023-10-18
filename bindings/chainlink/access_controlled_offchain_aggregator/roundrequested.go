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

type RoundRequestedHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorRoundRequested]
}

func (h *RoundRequestedHandler) ID() string {
	return h.addr + ":" + "0x3ea16a923ff4b1df6526e854c9e3a995c43385d70e73359e10623c74f0b52037"
}

func (h *RoundRequestedHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseRoundRequested(log.Log)
}

func (h *RoundRequestedHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorRoundRequested)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorRoundRequested")
	}
	return h.callback(log, e)
}

func (h *RoundRequestedHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseRoundRequested(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewRoundRequestedHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorRoundRequested]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RoundRequestedHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
