package processor

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/dog"
	"github.com/zarbanio/auction-keeper/bindings/zarban/osm"
	"github.com/zarbanio/auction-keeper/bindings/zarban/spot"
	inputMethods "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/services"
	"github.com/zarbanio/auction-keeper/services/loaders"
	sender "github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/store"
)

type DogBarkService struct {
	eth         *ethclient.Client
	store       store.IStore
	dog         *dog.Dog
	spot        *spot.Spot
	osms        map[common.Address]*osm.Osm
	vaultLoader *loaders.VaultLoader
	vatLoader   *loaders.VatLoader
	sender      sender.Sender
}

func NewDogBarkService(
	ctx context.Context,
	eth *ethclient.Client,
	store store.IStore,
	dogAddr common.Address,
	spotAddr common.Address,
	vaultLoader *loaders.VaultLoader,
	vatLoader *loaders.VatLoader,
	ilkLoader *loaders.IlksLoader,
	sender sender.Sender) *DogBarkService {
	d, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		services.Logger.Fatal().Str("service", "bark").Str("method", "NewDogBarkService").Msg("err while instancing a dog contract")
	}

	s, err := spot.NewSpot(spotAddr, eth)
	if err != nil {
		services.Logger.Fatal().Str("service", "bark").Str("method", "NewDogBarkService").Msg("err while instancing a spot contract")
	}

	ilks, err := ilkLoader.LoadIlks(ctx)
	if err != nil {
		services.Logger.Fatal().Str("service", "bark").Str("method", "NewDogBarkService").Msg("err while loading ilks")
	}

	osms := make(map[common.Address]*osm.Osm)
	for _, ilk := range ilks {
		osm, err := osm.NewOsm(ilk.Pip, eth)
		if err != nil {
			services.Logger.Fatal().Str("service", "bark").Str("method", "NewDogBarkService").Msg("err while instancing an osm contract")
		}
		osms[ilk.Pip] = osm
	}

	return &DogBarkService{
		eth:         eth,
		dog:         d,
		spot:        s,
		store:       store,
		vaultLoader: vaultLoader,
		vatLoader:   vatLoader,
		osms:        osms,
		sender:      sender,
	}
}

func (s *DogBarkService) Start(ctx context.Context) error {

	services.Logger.Info().Str("service", "bark").Str("method", "Start").Msg("bark Service is starting")

	// Fetch all vaults
	vaults, err := s.vaultLoader.FetchAllVaults(ctx)
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "bark").Str("method", "Start").Msg("error while fetching all vaults")
		return err
	}

	for _, vault := range vaults {

		log.Println("*** VAULT ID: ", vault.Id, " ***")

		ilk, err := s.store.GetIlkByName(ctx, vault.IlkName)
		if err != nil {
			services.Logger.Error().Err(err).Str("service", "bark").Str("method", "Start").Msg("error while getting an ilk")
			return err
		}

		// Convert the ilk name to a byte array of length 32
		var ilkName [32]byte
		copy(ilkName[:], []byte(ilk.Name))

		if vault.CollateralizationRatio.LessThan(ilk.MinimumCollateralizationRatio) {
			// create a new Oracle Security Manager(OSM) and call the Poke method
			err := s.osmPoke(ilk.Pip)
			if err != nil {
				services.Logger.Error().Err(err).Str("service", "bark").Str("method", "Start").Msg("error while calling the osmPoke method")
				return err
			}
			// call the Poke method from spotter
			err = s.spotterPoke(ilkName)
			if err != nil {
				services.Logger.Error().Err(err).Str("service", "bark").Str("method", "Start").Msg("error while calling the spotterPoke method")
				return err
			}

			log.Println("bark is running")
			bark := &inputMethods.DogBark{
				Ilk: ilkName,
				Urn: vault.Urn,
			}
			s.Run(bark)
		}

	}
	return nil
}

// Stop the service
func (s *DogBarkService) Stop() {
	services.Logger.Panic().Str("service", "bark").Str("method", "Stop").Msg("Implement me!")
}

// Run the service logic
func (s *DogBarkService) Run(bark *inputMethods.DogBark) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "bark").Str("method", "Run").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := s.dog.Bark(opts, bark.Ilk, bark.Urn, s.sender.GetAddress())
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "bark").Str("method", "Run").Msg("error while calling the dog bark method")
		return err
	}
	return s.sender.HandleSentTx(tx)
}

func (s *DogBarkService) osmPoke(pip common.Address) error {
	osm := s.osms[pip]

	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "bark").Str("method", "osmPoke").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := osm.Poke(opts)
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "bark").Str("method", "osmPoke").Msg("error while calling the osm poke method")
		return err
	}

	return s.sender.HandleSentTx(tx)
}

func (s *DogBarkService) spotterPoke(ilk [32]byte) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "bark").Str("method", "spotterPoke").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := s.spot.Poke(opts, ilk)
	if err != nil {
		services.Logger.Error().Err(err).Str("service", "bark").Str("method", "spotterPoke").Msg("error while calling the spotter poke method")
		return err
	}

	return s.sender.HandleSentTx(tx)
}
