package eventmanager

import (
	"github.com/zarbanio/auction-keeper/services/loaders"
	clipperServices "github.com/zarbanio/auction-keeper/services/processor/clipper"
	dogServices "github.com/zarbanio/auction-keeper/services/processor/dog"
	"github.com/zarbanio/auction-keeper/store"
)

type EventManager struct {
	store store.IStore
	// loaders
	ilksLoader  *loaders.IlksLoader
	vaultLoader *loaders.VaultLoader
	// services
	dogBarkService      *dogServices.DogBarkService
	clipperTakeServices []*clipperServices.ClipperTakeService
	clipperRedoServices []*clipperServices.ClipperRedoService
}

func NewEventManager(
	s store.IStore,
	i *loaders.IlksLoader,
	v *loaders.VaultLoader,
	dogBarkService *dogServices.DogBarkService,
	clipperTakeServices []*clipperServices.ClipperTakeService,
	clipperRedoServices []*clipperServices.ClipperRedoService,
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
