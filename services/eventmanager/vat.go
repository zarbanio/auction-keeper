package eventmanager

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/x/eth"
	"github.com/zarbanio/auction-keeper/x/events"
)

func (e *EventManager) VatFrobCallback() events.CallbackFn[vat.VatFrob] {
	return func(raw eth.Log, frob vat.VatFrob) error {
		_, err := e.store.CreateFrob(context.Background(), frob, raw.Id)
		if err != nil {
			return err
		}
		return e.UrnManipulation(frob.U, frob.V, frob.W)
	}
}

func (e *EventManager) VatGrabCallback() events.CallbackFn[vat.VatGrab] {
	return func(raw eth.Log, vatGrab vat.VatGrab) error {
		_, err := e.store.CreateVatGrab(context.Background(), vatGrab, raw.Id)
		if err != nil {
			return err
		}
		return e.UrnManipulation(vatGrab.U, vatGrab.V, vatGrab.W)
	}
}

func (e *EventManager) VatForkCallback() events.CallbackFn[vat.VatFork] {
	return func(raw eth.Log, fork vat.VatFork) error {
		_, err := e.store.CreateFork(context.Background(), fork, raw.Id)
		if err != nil {
			return err
		}
		return e.UrnManipulation(fork.Src, fork.Dst)
	}
}

func (e *EventManager) UpdateIlks() error {
	ilks, err := e.ilksLoader.LoadIlks(context.Background())
	if err != nil {
		return err
	}

	for _, ilk := range ilks {
		err := e.store.CreateOrUpdateIlk(context.Background(), ilk)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *EventManager) UrnManipulation(urns ...common.Address) error {
	panic("implement me")
}
