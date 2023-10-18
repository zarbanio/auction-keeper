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

type NewRoundHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorNewRound]
}

func (h *NewRoundHandler) ID() string {
	return h.addr + ":" + "0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271"
}

func (h *NewRoundHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseNewRound(log.Log)
}

func (h *NewRoundHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorNewRound)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorNewRound")
	}
	return h.callback(log, e)
}

func (h *NewRoundHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseNewRound(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewNewRoundHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorNewRound]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &NewRoundHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
