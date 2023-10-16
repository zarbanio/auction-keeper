package eventmanager

import (
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/store"
)

type EventManager struct {
	store        store.IStore
	ilksLoader   *loaders.IlksLoader
	clipperLoder *loaders.ClipperLoader
}

func NewEventManager(s store.IStore, i *loaders.IlksLoader, a *loaders.ClipperLoader) *EventManager {
	return &EventManager{
		store:        s,
		ilksLoader:   i,
		clipperLoder: a,
	}
}
