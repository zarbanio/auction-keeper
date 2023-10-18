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

type OwnershipTransferRequestedHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorOwnershipTransferRequested]
}

func (h *OwnershipTransferRequestedHandler) ID() string {
	return h.addr + ":" + "0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278"
}

func (h *OwnershipTransferRequestedHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseOwnershipTransferRequested(log.Log)
}

func (h *OwnershipTransferRequestedHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorOwnershipTransferRequested)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorOwnershipTransferRequested")
	}
	return h.callback(log, e)
}

func (h *OwnershipTransferRequestedHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseOwnershipTransferRequested(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewOwnershipTransferRequestedHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorOwnershipTransferRequested]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &OwnershipTransferRequestedHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
