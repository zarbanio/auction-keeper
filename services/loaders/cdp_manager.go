package loaders

import (
	"context"

	"github.com/zarbanio/auction-keeper/bindings/zarban/cdpmanager"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type CdpMangerLoader struct {
	cdpManger eth.Contract[cdpmanager.Cdpmanager]
}

func NewCdpMangerLoader(ceth eth.CachedEthereum, cdpManger eth.Contract[cdpmanager.Cdpmanager]) *CdpMangerLoader {
	return &CdpMangerLoader{cdpManger: cdpManger}
}

type CdpmanagerState struct{}

func (l CdpMangerLoader) Load(ctx context.Context) (CdpmanagerState, error) {
	panic("Impelent me")
}
