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

type CheckAccessDisabledHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorCheckAccessDisabled]
}

func (h *CheckAccessDisabledHandler) ID() string {
	return h.addr + ":" + "0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638"
}

func (h *CheckAccessDisabledHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseCheckAccessDisabled(log.Log)
}

func (h *CheckAccessDisabledHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorCheckAccessDisabled)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorCheckAccessDisabled")
	}
	return h.callback(log, e)
}

func (h *CheckAccessDisabledHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseCheckAccessDisabled(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewCheckAccessDisabledHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorCheckAccessDisabled]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &CheckAccessDisabledHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
