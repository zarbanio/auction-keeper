package eventmanager

import (
	"context"

	"github.com/zarbanio/auction-keeper/bindings/zarban/vow"
	"github.com/zarbanio/auction-keeper/x/eth"
	"github.com/zarbanio/auction-keeper/x/events"
)

func (e *EventManager) VowFlogCallback() events.CallbackFn[vow.VowFlog] {
	return func(raw eth.Log, flog vow.VowFlog) error {
		_, err := e.store.CreateFlog(context.Background(), flog, raw.Id)
		if err != nil {
			return err
		}
		return nil
	}
}

func (e *EventManager) VowFessCallback() events.CallbackFn[vow.VowFess] {
	return func(raw eth.Log, fess vow.VowFess) error {
		_, err := e.store.CreateFess(context.Background(), fess, raw.Id)
		if err != nil {
			return err
		}
		return nil
	}
}
