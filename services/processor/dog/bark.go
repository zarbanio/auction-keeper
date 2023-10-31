package processor

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/dog"
	"github.com/zarbanio/auction-keeper/bindings/zarban/osm"
	"github.com/zarbanio/auction-keeper/bindings/zarban/spot"
	inputMethods "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/services/loaders"
	sender "github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/store"
)

type DogBarkService struct {
	eth           *ethclient.Client
	blockInterval time.Duration
	store         store.IStore
	dog           *dog.Dog
	spot          *spot.Spot
	osms          map[common.Address]*osm.Osm
	vaultLoader   *loaders.VaultLoader
	vatLoader     *loaders.VatLoader
	sender        sender.Sender
}

func NewDogBarkService(
	ctx context.Context,
	eth *ethclient.Client,
	blockInterval time.Duration,
	store store.IStore,
	dogAddr common.Address,
	spotAddr common.Address,
	vaultLoader *loaders.VaultLoader,
	vatLoader *loaders.VatLoader,
	ilkLoader *loaders.IlksLoader,
	sender sender.Sender) *DogBarkService {

	d, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	s, err := spot.NewSpot(spotAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	ilks, err := ilkLoader.LoadIlks(ctx)
	if err != nil {
		log.Fatal(err)
	}

	osms := make(map[common.Address]*osm.Osm)
	for _, ilk := range ilks {
		osm, err := osm.NewOsm(ilk.Pip, eth)
		if err != nil {
			log.Fatal(err)
		}
		osms[ilk.Pip] = osm
	}

	return &DogBarkService{
		eth:           eth,
		blockInterval: blockInterval,
		dog:           d,
		spot:          s,
		store:         store,
		vaultLoader:   vaultLoader,
		vatLoader:     vatLoader,
		osms:          osms,
		sender:        sender,
	}
}

func (s *DogBarkService) Start(ctx context.Context) error {

	log.Println("Dog Bark Service is starting")

	// Fetch all vaults
	vaults, err := s.vaultLoader.FetchAllVaults(ctx)
	if err != nil {
		return err
	}

	for _, vault := range vaults {

		log.Println("*** VAULT ID: ", vault.Id, " ***")

		ilk, err := s.store.GetIlkByName(ctx, vault.IlkName)
		if err != nil {
			return err
		}

		// Convert the ilk name to a byte array of length 32
		var ilkName [32]byte
		copy(ilkName[:], []byte(ilk.Name))

		if vault.CollateralizationRatio.LessThan(ilk.MinimumCollateralizationRatio) {
			if err != nil {
				return err
			}

			// create a new Oracle Security Manager(OSM) and call the Poke method
			err := s.osmPoke(ctx, ilk.Pip)
			if err != nil {
				return nil
			}
			// call the Poke method from spotter
			err = s.spotterPoke(ctx, ilkName)
			if err != nil {
				return nil
			}

			log.Println("Dog Bark is running")
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
	panic("Implement me!")
}

// Run the service logic
func (s *DogBarkService) Run(bark *inputMethods.DogBark) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		return err
	}

	tx, err := s.dog.Bark(opts, bark.Ilk, bark.Urn, s.sender.GetAddress())
	if err != nil {
		return err
	}
	return s.sender.HandleSentTx(tx)
}

func (s *DogBarkService) osmPoke(ctx context.Context, pip common.Address) error {
	osm := s.osms[pip]

	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		return err
	}

	tx, err := osm.Poke(opts)
	if err != nil {
		return err
	}

	return s.sender.HandleSentTx(tx)
}

func (s *DogBarkService) spotterPoke(ctx context.Context, ilk [32]byte) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		return err
	}

	tx, err := s.spot.Poke(opts, ilk)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return s.sender.HandleSentTx(tx)
}
