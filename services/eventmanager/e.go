package eventmanager

import (
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/store"
)

type EventManager struct {
	store       store.IStore
	ilksLoader  *loaders.IlksLoader
	vaultLoader *loaders.VaultLoader
}

func NewEventManager(s store.IStore, i *loaders.IlksLoader, v *loaders.VaultLoader) *EventManager {
	return &EventManager{
		store:       s,
		ilksLoader:  i,
		vaultLoader: v,
	}
}
