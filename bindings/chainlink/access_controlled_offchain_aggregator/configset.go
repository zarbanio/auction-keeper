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

type ConfigSetHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorConfigSet]
}

func (h *ConfigSetHandler) ID() string {
	return h.addr + ":" + "0x25d719d88a4512dd76c7442b910a83360845505894eb444ef299409e180f8fb9"
}

func (h *ConfigSetHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseConfigSet(log.Log)
}

func (h *ConfigSetHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorConfigSet)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorConfigSet")
	}
	return h.callback(log, e)
}

func (h *ConfigSetHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseConfigSet(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewConfigSetHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorConfigSet]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &ConfigSetHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
