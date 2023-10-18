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

type AnswerUpdatedHandler struct {
	addr string
	binding  *AccessControlledOffchainAggregator
	callback events.CallbackFn[AccessControlledOffchainAggregatorAnswerUpdated]
}

func (h *AnswerUpdatedHandler) ID() string {
	return h.addr + ":" + "0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f"
}

func (h *AnswerUpdatedHandler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.ParseAnswerUpdated(log.Log)
}

func (h *AnswerUpdatedHandler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.(AccessControlledOffchainAggregatorAnswerUpdated)
	if !ok {
		return errors.New("event type is not AccessControlledOffchainAggregatorAnswerUpdated")
	}
	return h.callback(log, e)
}

func (h *AnswerUpdatedHandler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.ParseAnswerUpdated(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func NewAnswerUpdatedHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[AccessControlledOffchainAggregatorAnswerUpdated]) events.Handler {
	b, err := NewAccessControlledOffchainAggregator(addr, eth)
	if err != nil {
		panic(err)
	}
	return &AnswerUpdatedHandler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
