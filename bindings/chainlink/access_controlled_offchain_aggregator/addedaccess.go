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

type AddedAccessHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorAddedAccess]
}

func (h *AddedAccessHandler) ID() string {
	return h.addr + ":" + "0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4"
}

func (h *AddedAccessHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseAddedAccess(log.Log)
}

func (h *AddedAccessHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorAddedAccess)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorAddedAccess")
	}
	return h.callback(log, e)
}

func (h *AddedAccessHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseAddedAccess(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewAddedAccessHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorAddedAccess]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &AddedAccessHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
