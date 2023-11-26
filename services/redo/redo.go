package redo

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/services/logger"
	"github.com/zarbanio/auction-keeper/services/sender"
)

type Service struct {
	ilkName string
	eth     *ethclient.Client
	clipper *clipper.Clipper
	sender  sender.Sender
	l       *logger.Logger
}

func NewService(
	ilkName string,
	eth *ethclient.Client,
	clipperAddr common.Address,
	sender sender.Sender,
	l *logger.Logger,
) *Service {
	c, err := clipper.NewClipper(clipperAddr, eth)
	if err != nil {
		l.Logger.Fatal().Str("service", "redo").Str("method", "NewClipperTakeService").Msg("err while instancing a clipper contract")
	}

	return &Service{
		ilkName: ilkName,
		eth:     eth,
		clipper: c,
		sender:  sender,
		l:       l,
	}
}

func (s *Service) Run(ctx context.Context) error {
	auctionIds, err := s.clipper.List(&bind.CallOpts{Context: ctx})
	if err != nil {
		s.l.Logger.Error().
			Err(err).Str("service", "redo").
			Str("method", "run").
			Str("ilkName", s.ilkName).
			Msg("error while getting the list of auction IDs")
		return err
	}

	for _, auctionId := range auctionIds {
		status, err := s.clipper.GetStatus(&bind.CallOpts{Context: ctx}, auctionId)
		if err != nil {
			s.l.Logger.Error().Err(err).
				Str("service", "redo").
				Str("method", "run").
				Str("ilkName", s.ilkName).
				Int64("auctionId", auctionId.Int64()).
				Msg("error while getting the status of the auction")
			continue
		}

		if status.NeedsRedo {
			s.l.Logger.Info().
				Str("service", "redo").
				Str("method", "run").
				Str("ilkName", s.ilkName).
				Int64("auctionId", auctionId.Int64()).
				Msg("auction needs redo")
			err = s.Redo(auctionId)
			if err != nil {
				s.l.Logger.Error().Err(err).
					Str("service", "redo").
					Str("method", "run").
					Str("ilkName", s.ilkName).
					Int64("auctionId", auctionId.Int64()).
					Msg("error while redoing the auction")
			} else {
				s.l.Logger.Info().
					Str("service", "redo").
					Str("method", "run").
					Str("ilkName", s.ilkName).
					Int64("auctionId", auctionId.Int64()).
					Msg("auction redo successfully")
			}
		}
	}

	s.l.Logger.Info().
		Str("service", "redo").
		Str("method", "run").
		Str("ilkName", s.ilkName).
		Msg("auction redo service iteration completed")

	return nil
}

// Run the service logic
func (s *Service) Redo(auctionId *big.Int) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "redo").Str("method", "Run").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := s.clipper.Redo(opts, auctionId, s.sender.GetAddress())
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "redo").Str("method", "Run").Msg("error while calling the cliiper redo method")
		return err
	}
	return s.sender.HandleSentTx(tx)
}

func (s *Service) Stop() {
	s.l.Logger.Panic().Str("service", "redo").Str("method", "Stop").Msg("Implement me!")
}
