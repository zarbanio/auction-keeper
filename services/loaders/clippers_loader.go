package loaders

import (
	clipper "github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
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
		Id:  decimal.NewFromBigInt(auctionId, 0),
		Top: decimal.NewFromBigInt(sale.Top, 0),
		Tab: decimal.NewFromBigInt(sale.Tab, 0),
		Lot: decimal.NewFromBigInt(sale.Lot, 0),
		Usr: sale.Usr,
		Tic: sale.Tic.Uint64(),

		ClipperName: cp.Name,
	}
	return &auction, nil
}
