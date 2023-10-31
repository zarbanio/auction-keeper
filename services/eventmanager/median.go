package eventmanager

import (
	"context"

	"github.com/zarbanio/auction-keeper/bindings/zarban/median"
	"github.com/zarbanio/auction-keeper/x/eth"
	"github.com/zarbanio/auction-keeper/x/events"
)

func (e *EventManager) MedianLogMedianPriceCallback() events.CallbackFn[median.MedianLogMedianPrice] {
	return func(raw eth.Log, median median.MedianLogMedianPrice) error {
		_, err := e.store.CreateLogMedianPrice(context.Background(), median, raw.Id)
		if err != nil {
			return err
		}

		err = e.dogBarkService.Start(context.Background())
		if err != nil {
			return err
		}

		return e.UpdateIlks()
	}
}
