package events

import "github.com/zarbanio/auction-keeper/x/eth"

type CallbackFn[T any] func(raw eth.Log, event T) error

type Handler interface {
	ID() string
	DecodeLog(log eth.Log) (interface{}, error)
	HandleEvent(log eth.Log, event interface{}) error
	DecodeAndHandle(log eth.Log) error
}
