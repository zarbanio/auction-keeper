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

type OwnershipTransferredHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorOwnershipTransferred]
}

func (h *OwnershipTransferredHandler) ID() string {
	return h.addr + ":" + "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"
}

func (h *OwnershipTransferredHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseOwnershipTransferred(log.Log)
}

func (h *OwnershipTransferredHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorOwnershipTransferred)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorOwnershipTransferred")
	}
	return h.callback(log, e)
}

func (h *OwnershipTransferredHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseOwnershipTransferred(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewOwnershipTransferredHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorOwnershipTransferred]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &OwnershipTransferredHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
