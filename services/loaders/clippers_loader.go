package loaders

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/abacus"
	clipper "github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/domain/entities"
)

type ClipperLoader struct {
	eth            *ethclient.Client
	ClipperAddress common.Address
	Clipper        *clipper.Clipper
}

func NewClipperLoader(eth *ethclient.Client, clipperAddr common.Address) (*ClipperLoader, error) {
	_clipper, err := clipper.NewClipper(clipperAddr, eth)
	if err != nil {
		return nil, err
	}

	return &ClipperLoader{
		eth:            eth,
		ClipperAddress: clipperAddr,
		Clipper:        _clipper,
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
		Tic: time.Unix(sale.Tic.Int64(), 0),
	}
	return &auction, nil
}

func (cp *ClipperLoader) GetAuctionStatus(id *big.Int) (bool, error) {
	status, err := cp.Clipper.ClipperCaller.GetStatus(nil, id)
	if err != nil {
		return false, err
	}
	return status.NeedsRedo, nil
}

func (cp *ClipperLoader) GetChost() (*big.Int, error) {
	// chost = ilk.dust * ilk.chop(liquidation penalty) [rad] (fixed point decimal with 45 decimals )
	chost, err := cp.Clipper.Chost(nil)
	if err != nil {
		return nil, err
	}
	return chost, nil
}

func (cp *ClipperLoader) GetAbacusInstance() (*abacus.Abacus, error) {
	// clipper.calc() returns the abacus address of the collateral
	abacusAddress, err := cp.Clipper.Calc(nil)
	if err != nil {
		return nil, err
	}

	abacusInstance, err := abacus.NewAbacus(abacusAddress, cp.eth)
	if err != nil {
		return nil, err
	}

	return abacusInstance, nil
}
