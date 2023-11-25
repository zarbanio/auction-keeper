package eventmanager

import (
	"github.com/zarbanio/auction-keeper/services/bark"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/redo"
	"github.com/zarbanio/auction-keeper/services/take"
	"github.com/zarbanio/auction-keeper/store"
)

type EventManager struct {
	store store.IStore
	// loaders
	ilksLoader  *loaders.IlksLoader
	vaultLoader *loaders.VaultLoader
	// services
	dogBarkService      *bark.Service
	clipperTakeServices []*take.Service
	clipperRedoServices []*redo.Service
}

func NewEventManager(
	s store.IStore,
	i *loaders.IlksLoader,
	v *loaders.VaultLoader,
	dogBarkService *bark.Service,
	clipperTakeServices []*take.Service,
	clipperRedoServices []*redo.Service,
) *EventManager {
	return &EventManager{
		store:               s,
		ilksLoader:          i,
		vaultLoader:         v,
		dogBarkService:      dogBarkService,
		clipperTakeServices: clipperTakeServices,
		clipperRedoServices: clipperRedoServices,
	}
}
