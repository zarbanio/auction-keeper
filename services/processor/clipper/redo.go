package processor

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/collateral"
	"github.com/zarbanio/auction-keeper/services"
	"github.com/zarbanio/auction-keeper/services/sender"
)

type ClipperRedoService struct {
	eth        *ethclient.Client
	clipper    *clipper.Clipper
	sender     sender.Sender
	collateral collateral.Collateral
}

func NewClipperRedoService(
	eth *ethclient.Client,
	clipperAddr common.Address,
	sender sender.Sender,
	collateral collateral.Collateral,
) *ClipperRedoService {

	c, err := clipper.NewClipper(clipperAddr, eth)
	if err != nil {
		services.Logger.Fatal().Str("service", "redo").Str("method", "NewClipperTakeService").Msg("err while instancing a clipper contract")
	}

	return &ClipperRedoService{
		eth:        eth,
		clipper:    c,
		sender:     sender,
		collateral: collateral,
	}
}

func (s *ClipperRedoService) Start(ctx context.Context) error {

	// 1) Get the IDs of the auctions
	auctionIds, err := s.clipper.List(&bind.CallOpts{Context: ctx})
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "redo").Str("method", "Start").Msg("error while getting the list of auction IDs")
		return err
	}

	for id, auctionId := range auctionIds {
		log.Printf("\tprocessing auction id: %d\n", id)
		needsRedo, err := s.collateral.ClipperLoader.GetAuctionStatus(auctionId)
		if err != nil {
			services.Logger.Error().Err(err).Str("service", "redo").Str("method", "Start").Msg("error while getting the status of the auction")
			continue
		}

		if needsRedo {
			s.Run(auctionId)
		}
	}

	return nil
}

// Run the service logic
func (s *ClipperRedoService) Run(auctionId *big.Int) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "redo").Str("method", "Run").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := s.clipper.Redo(opts, auctionId, s.sender.GetAddress())
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "redo").Str("method", "Run").Msg("error while calling the cliiper redo method")
		return err
	}
	return s.sender.HandleSentTx(tx)
}

// Stop the service
func (s *ClipperRedoService) Stop() {
	services.Logger.Panic().Str("service", "redo").Str("method", "Stop").Msg("Implement me!")
}
