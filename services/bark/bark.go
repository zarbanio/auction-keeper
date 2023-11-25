package bark

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/dog"
	"github.com/zarbanio/auction-keeper/bindings/zarban/osm"
	"github.com/zarbanio/auction-keeper/bindings/zarban/spot"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/logger"
	sender "github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/store"
)

type Service struct {
	eth         *ethclient.Client
	store       store.IStore
	dog         *dog.Dog
	spot        *spot.Spot
	osms        map[common.Address]*osm.Osm
	vaultLoader *loaders.VaultLoader
	vatLoader   *loaders.VatLoader
	sender      sender.Sender
	l           *logger.Logger
}

func NewService(
	ctx context.Context,
	eth *ethclient.Client,
	store store.IStore,
	dogAddr common.Address,
	spotAddr common.Address,
	vaultLoader *loaders.VaultLoader,
	vatLoader *loaders.VatLoader,
	ilkLoader *loaders.IlksLoader,
	sender sender.Sender,
	l *logger.Logger,
) *Service {
	d, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		l.Logger.Fatal().Str("service", "bark").Str("method", "NewDogBarkService").Msg("err while instancing a dog contract")
	}

	s, err := spot.NewSpot(spotAddr, eth)
	if err != nil {
		l.Logger.Fatal().Str("service", "bark").Str("method", "NewDogBarkService").Msg("err while instancing a spot contract")
	}

	ilks, err := ilkLoader.LoadIlks(ctx)
	if err != nil {
		l.Logger.Fatal().Str("service", "bark").Str("method", "NewDogBarkService").Msg("err while loading ilks")
	}

	osms := make(map[common.Address]*osm.Osm)
	for _, ilk := range ilks {
		osm, err := osm.NewOsm(ilk.Pip, eth)
		if err != nil {
			l.Logger.Fatal().Str("service", "bark").Str("method", "NewDogBarkService").Msg("err while instancing an osm contract")
		}
		osms[ilk.Pip] = osm
	}

	return &Service{
		eth:         eth,
		dog:         d,
		spot:        s,
		store:       store,
		vaultLoader: vaultLoader,
		vatLoader:   vatLoader,
		osms:        osms,
		sender:      sender,
		l:           l,
	}
}

func (s *Service) Start(ctx context.Context) error {
	s.l.Logger.Info().Str("service", "bark").Str("method", "Start").Msg("bark Service is starting")

	vaults, err := s.vaultLoader.FetchAllVaults(ctx)
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "Start").Msg("error while fetching all vaults")
		return err
	}

	for _, vault := range vaults {
		ilk, err := s.store.GetIlkByName(ctx, vault.IlkName)
		if err != nil {
			s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "Start").Msg("error while getting an ilk")
			return err
		}

		// Convert the ilk name to a byte array of length 32
		var ilkBytes32 [32]byte
		copy(ilkBytes32[:], []byte(ilk.Name))

		if !vault.CollateralizationRatio.LessThan(ilk.MinimumCollateralizationRatio) {
			s.l.Logger.Info().Str("service", "bark").Str("method", "Start").Int64("id", vault.Id).Msg("vault is not undercollateralized, skipping")
			continue
		}

		s.l.Logger.Info().Str("service", "bark").Str("method", "Start").Int64("id", vault.Id).Msg("vault is undercollateralized, poking the osm")
		err = s.osmPoke(ilk.Pip)
		if err != nil {
			s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "Start").Msg("error while calling the osmPoke method")
			return err
		}

		s.l.Logger.Info().Str("service", "bark").Str("method", "Start").Int64("id", vault.Id).Msg("poking the spotter")
		err = s.spotterPoke(ilkBytes32)
		if err != nil {
			s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "Start").Msg("error while calling the spotterPoke method")
			return err
		}

		err = s.Bark(ilkBytes32, vault.Urn)
		if err != nil {
			s.l.Logger.Error().Err(err).
				Str("service", "bark").
				Str("method", "Start").
				Int64("id", vault.Id).
				Msg("error while calling the bark method")
			return err
		}
	}
	return nil
}

func (s *Service) Stop() {
	s.l.Logger.Panic().Str("service", "bark").Str("method", "Stop").Msg("Implement me!")
}

// Run the service logic
func (s *Service) Bark(ilkBytes32 [32]byte, urn common.Address) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "Run").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := s.dog.Bark(opts, ilkBytes32, urn, s.sender.GetAddress())
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "Run").Msg("error while calling the dog bark method")
		return err
	}
	return s.sender.HandleSentTx(tx)
}

func (s *Service) osmPoke(pip common.Address) error {
	osm := s.osms[pip]

	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "osmPoke").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := osm.Poke(opts)
	if err != nil {
		if err.Error() != "execution reverted: OSM/not-passed" {
			s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "osmPoke").Msg("error while calling the osm poke method")
			return err
		}
		return nil
	}

	return s.sender.HandleSentTx(tx)
}

func (s *Service) spotterPoke(ilk [32]byte) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "spotterPoke").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := s.spot.Poke(opts, ilk)
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "bark").Str("method", "spotterPoke").Msg("error while calling the spotter poke method")
		return err
	}

	return s.sender.HandleSentTx(tx)
}
