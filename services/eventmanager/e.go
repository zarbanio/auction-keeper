package eventmanager

import (
	"github.com/zarbanio/auction-keeper/services/loaders"
	processor "github.com/zarbanio/auction-keeper/services/processor/dog"
	"github.com/zarbanio/auction-keeper/store"
)

type EventManager struct {
	store          store.IStore
	ilksLoader     *loaders.IlksLoader
	vaultLoader    *loaders.VaultLoader
	dogBarkService *processor.DogBarkService
}

func NewEventManager(s store.IStore, i *loaders.IlksLoader, v *loaders.VaultLoader, dogBarkService *processor.DogBarkService) *EventManager {
	return &EventManager{
		store:          s,
		ilksLoader:     i,
		vaultLoader:    v,
		dogBarkService: dogBarkService,
	}
}
