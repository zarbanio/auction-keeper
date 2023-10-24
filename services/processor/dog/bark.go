package processor

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/dog"
	"github.com/zarbanio/auction-keeper/bindings/zarban/osm"
	"github.com/zarbanio/auction-keeper/bindings/zarban/spot"
	"github.com/zarbanio/auction-keeper/domain/entities"
	inputMethods "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/domain/math"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/transaction"
	"github.com/zarbanio/auction-keeper/store"
)

type dogBarkService struct {
	eth         *ethclient.Client
	store       store.IStore
	dog         *dog.Dog
	spot        *spot.Spot
	vaultLoader *loaders.VaultLoader
	vatLoader   *loaders.VatLoader
	keeper      *transaction.Sender
	opts        *bind.TransactOpts
}

func NewDogService(
	eth *ethclient.Client,
	store store.IStore,
	dogAddr common.Address,
	spotAddr common.Address,
	vaultLoader *loaders.VaultLoader,
	vatLoader *loaders.VatLoader,
	keeper *transaction.Sender) *dogBarkService {

	d, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	s, err := spot.NewSpot(spotAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	opts, err := keeper.GetOpts()
	if err != nil {
		log.Fatal(err)
	}

	return &dogBarkService{
		eth:         eth,
		dog:         d,
		spot:        s,
		store:       store,
		vaultLoader: vaultLoader,
		vatLoader:   vatLoader,
		keeper:      keeper,
		opts:        opts,
	}
}

func (s *dogBarkService) Start(ctx context.Context) error {

	log.Println("Dog Bark Service is starting")

	// Fetch all vaults
	vaults, err := s.vaultLoader.FetchAllVaults(ctx)

	if err != nil {
		return err
	}
	// obtain params of isVaultUnsafe func PHASE 1

	// 1) dog Hole
	dogHole, err := s.dog.Hole(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}

	// 2) dog Dirt
	dogDirt, err := s.dog.Dirt(&bind.CallOpts{Context: ctx})
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

			osm, err := osm.NewOsm(ilk.Pip, s.eth)
			if err != nil {
				log.Fatal(err)
			}

			tx, err := osm.Poke(s.opts)
			if err != nil {
				return nil
			}

			log.Println("\tOSM Transaction: ", tx.Hash().Hex())

			tx, err = s.spot.Poke(s.opts, ilkName)

			if err != nil {
				return nil
			}

			log.Println("\tSpotter Transaction: ", tx.Hash().Hex())

			// obtain params of isVaultUnsafe func PHASE 2
			v, err := s.vaultLoader.GetVaultByIlkUrn(ctx, ilkName, vault.Urn)
			if err != nil {
				return nil
			}
			// 1) urn
			urn := entities.Urn{
				Ink: v.Urn.Ink,
				Art: v.Urn.Art,
			}
			// 2) vat ilk
			vatIlk, err := s.vatLoader.GetIlkById(ctx, ilkName)
			if err != nil {
				return nil
			}

			// 3) dog chop
			dogChop, err := s.dog.Chop(&bind.CallOpts{Context: ctx}, ilkName)
			if err != nil {
				return err
			}
			if s.isVaultUnSafe(urn, *vatIlk, dogHole, dogDirt, ilk.Hole.BigInt(), ilk.Dirt.BigInt(), dogChop) {
				log.Println("Dog Bark is running")
				bark := &inputMethods.DogBark{
					Ilk: ilkName,
					Urn: vault.Urn,
				}
				go s.Run(bark) // Run the service logic in a goroutine
			}
		}

	}
	return nil
}

// Stop the service
func (s *dogBarkService) Stop() {
	panic("Implement me!")
}

// Run the service logic
func (s *dogBarkService) Run(bark *inputMethods.DogBark) (*types.Transaction, error) {

	tx, err := s.dog.Bark(s.opts, bark.Ilk, bark.Urn, s.keeper.GetAddress())
	if err != nil {
		log.Fatal(err)
	}

	// Print the transaction hash
	log.Println(tx.Hash().Hex())
	return tx, nil
}

func (s *dogBarkService) isVaultUnSafe(urn entities.Urn, vatIlk entities.VatIlk, dogHole, dogDirt, ilkHole, ilkDirt, chop *big.Int) bool {
	/*
			Collateral value should be less than the product of our stablecoin debt and the debt multiplier
		    require(spot > 0 && mul(ink, spot) < mul(art, rate), "Dog/not-unsafe");
	*/
	if vatIlk.Spot.Cmp(big.NewInt(0)) <= 0 {
		return false
	}

	safe := new(big.Int).Mul(urn.Ink, vatIlk.Spot).Cmp(new(big.Int).Mul(urn.Art, vatIlk.Rate)) >= 0
	if safe {
		return false
	}

	/*
	   Ensure there's room in the dog.hole
	   require(Hole > Dirt && milk.hole > milk.dirt, "Dog/liquidation-limit-hit");
	*/
	if dogHole.Cmp(dogDirt) <= 0 {
		return false
	}
	dogRoom := new(big.Int).Sub(dogHole, dogDirt)

	/*
	   Ensure there's room in the collateral-specific hole
	   require(Hole > Dirt && milk.hole > milk.dirt, "Dog/liquidation-limit-hit");
	*/
	if ilkHole.Cmp(ilkDirt) <= 0 {
		return false
	}
	ilkRoom := new(big.Int).Sub(ilkHole, ilkDirt)

	/*
	   Prevent dusty partial liquidation
	*/
	room := math.BigMin(dogRoom, ilkRoom)
	wadRoom := new(big.Int).Mul(room, math.Wad)
	dart := math.BigMin(urn.Art, new(big.Int).Div(new(big.Int).Div(wadRoom, vatIlk.Rate), chop))
	if urn.Art.Cmp(dart) > 0 {
		if new(big.Int).Mul(new(big.Int).Sub(urn.Art, dart), vatIlk.Rate).Cmp(vatIlk.Dust) < 0 {
			dart = urn.Art
		} else if new(big.Int).Mul(dart, vatIlk.Rate).Cmp(vatIlk.Dust) < 0 {
			// require(mul(dart, rate) >= dust, "Dog/dusty-auction-from-partial-liquidation");
			return false
		}
	}

	dink := new(big.Int).Div(new(big.Int).Mul(urn.Ink, dart), urn.Art)
	if dink.Cmp(math.Zero) <= 0 {
		return false // require(dink > 0, "Dog/null-auction");
	}

	return true
}
