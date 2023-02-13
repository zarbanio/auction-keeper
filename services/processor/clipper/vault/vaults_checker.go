package vault

import (
	"context"
	"errors"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/cache"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/IR-Digital-Token/auction-keeper/domain/math"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"log"
	"math/big"
	"sync"
	"time"
)

type VaultChecker struct {
	cache      cache.ICache
	sender     transaction.ISender
	dogLoader  *loaders.DogLoader
	vatLoader  *loaders.VaultLoader
	processing sync.Mutex
}

func NewVaultsChecker(cache cache.ICache, sender transaction.ISender, dogLoader *loaders.DogLoader, vatLoader *loaders.VaultLoader) *VaultChecker {
	vaultChecker := &VaultChecker{
		cache:      cache,
		sender:     sender,
		dogLoader:  dogLoader,
		vatLoader:  vatLoader,
		processing: sync.Mutex{},
	}

	return vaultChecker
}

func (vc *VaultChecker) Start() {
	vc.processing.Lock()
	defer vc.processing.Unlock()

	//TODO: check dog is relay or not in clipper (log for not and return)

	// get dog hole and dog dirt
	dogHole, err := vc.dogLoader.GetHole(context.Background())
	if err != nil {
		log.Println("error in getting dog hole.", err)
		return
	}

	dogDirt, err := vc.dogLoader.GetDirt(context.Background())
	if err != nil {
		log.Println("error in getting dog dirt.", err)
		return
	}

	//  crate map for dog ilks
	dogIlks := make(map[string]*entities.DogIlk)

	// get vaults from cache
	vaults, err := vc.cache.GetVaults(context.Background())
	if err != nil {
		log.Println("error in getting vaults from cache.", err)
		return
	}

	log.Printf("Evaluating %d urns to be bitten if any are unsafe\n", len(vaults))

	for _, vault := range vaults {
		ilkId := vault.IlkId
		ilkName := string(ilkId[:])

		// get dog ilk hole
		dogIlk, ok := dogIlks[ilkName]
		if !ok {
			dogIlk, err = vc.dogLoader.GetIlk(context.Background(), ilkId)
			if err != nil {
				log.Printf("error in getting dog ilk. ilkName:%s  error:%v", ilkName, err)
				continue
			}
			dogIlks[ilkName] = dogIlk
		}

		// vatIlk.rate changes every block, so try to read from cache if it doesn't exist get ilk from vat and save it in cache for 10 seconds
		vatIlk, err := vc.cache.GetVatIlkById(context.Background(), ilkId)
		if errors.Is(err, cache.ErrIlkNotFound) {
			vatIlk, err = vc.vatLoader.GetIlkById(context.Background(), ilkId)
			if err != nil {
				log.Printf("error in getting ilk from vat. ilkName:%s  error:%v", ilkName, err)
				continue
			}
			err := vc.cache.SaveVatIlk(context.Background(), *vatIlk, 10*time.Second)
			if err != nil {
				log.Printf("error in saving vat ilk to vat. ilkName:%s  error:%v", ilkName, err)
				continue
			}
		}
		if err != nil {
			log.Printf("error in getting vat ilk from cache. ilkName:%s  error:%v", ilkName, err)
			continue
		}

		// check can call bark
		if canBark(vault.Urn, *vatIlk, dogHole, dogDirt, dogIlk.Hole, dogIlk.Dirt, dogIlk.Chop) {
			txHash, err := vc.sender.Bark(ilkId, vault.UrnAddress)
			if err != nil {
				log.Println("error in sending bark transaction.", err)
				continue
			}

			// contract update ilk.dirt after bark
			delete(dogIlks, ilkName)
			fmt.Printf("\tBark Transaction Hash: %s\n", txHash)
		}
	}

	log.Printf("Checked %d urns", len(vaults))
}

func canBark(urn entities.Urn, vatIlk entities.VatIlk, dogHole, dogDirt, ilkHole, ilkDirt, chop *big.Int) bool {
	/*
			Collateral value should be less than the product of our stablecoin debt and the debt multiplier
		    require(spot > 0 && mul(ink, spot) < mul(art, rate), "Dog/not-unsafe");
	*/
	// TODO: in contract doesn't convert Ink and Art to RAY
	inkRay := math.IntToRay(urn.Ink) // convert to RAY
	artRay := math.IntToRay(urn.Art) // convert to RAY

	safe := new(big.Int).Mul(inkRay, vatIlk.Spot).Cmp(new(big.Int).Mul(artRay, vatIlk.Rate)) >= 0
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
	dart := math.BigMin(urn.Art, new(big.Int).Div(new(big.Int).Div(room, vatIlk.Rate), chop)) // TODO: Wad & Rad
	if urn.Art.Cmp(dart) > 0 {
		if new(big.Int).Sub(new(big.Int).Mul(urn.Art, dart), vatIlk.Rate).Cmp(vatIlk.Dust) < 0 { // TODO: Rad
			return true
		} else if new(big.Int).Mul(dart, vatIlk.Rate).Cmp(vatIlk.Dust) < 0 { // TODO: Rad
			// require(mul(dart, rate) >= dust, "Dog/dusty-auction-from-partial-liquidation");
			return false
		}
	}
	return true
}
