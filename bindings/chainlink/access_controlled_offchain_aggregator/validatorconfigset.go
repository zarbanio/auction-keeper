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

type ValidatorConfigSetHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorValidatorConfigSet]
}

func (h *ValidatorConfigSetHandler) ID() string {
	return h.addr + ":" + "0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541"
}

func (h *ValidatorConfigSetHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseValidatorConfigSet(log.Log)
}

func (h *ValidatorConfigSetHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorValidatorConfigSet)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorValidatorConfigSet")
	}
	return h.callback(log, e)
}

func (h *ValidatorConfigSetHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseValidatorConfigSet(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewValidatorConfigSetHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorValidatorConfigSet]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &ValidatorConfigSetHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
