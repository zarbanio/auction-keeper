package loaders

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type ClipperLoader struct {
	eth     *ethclient.Client
	Name    string
	Address common.Address
	Clipper *clipper.Clipper
}

func NewClipperLoader(eth *ethclient.Client, name string, clipperAddr common.Address) (*ClipperLoader, error) {
	_clipper, err := clipper.NewClipper(clipperAddr, eth)
	if err != nil {
		return nil, err
	}
	return &ClipperLoader{
		eth:     eth,
		Name:    name,
		Address: clipperAddr,
		Clipper: _clipper,
	}, nil
}

func (cp *ClipperLoader) GetActiveAuctions() ([]*big.Int, error) {
	activeAuctionsId, err := cp.Clipper.ClipperCaller.List(nil)
	if err != nil {
		return nil, err
	}
	return activeAuctionsId, nil
}

func (cp *ClipperLoader) GetSale(auctionId *big.Int) (*entities.Auction, error) {
	sale, err := cp.Clipper.ClipperCaller.Sales(nil, auctionId)
	if err != nil {
		return nil, err
	}

	auction := entities.Auction{
		Id:  auctionId,
		Top: sale.Top,
		Tab: sale.Tab,
		Lot: sale.Lot,
		Usr: sale.Usr,
		Tic: sale.Tic.Uint64(),
	}
	return &auction, nil
}
