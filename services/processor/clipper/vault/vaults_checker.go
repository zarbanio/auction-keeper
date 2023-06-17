package vault

import (
	"context"
	"errors"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/domain/entities"
	"github.com/zarbanio/auction-keeper/domain/math"
	"github.com/zarbanio/auction-keeper/services/actions"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/chain"
)

type VaultChecker struct {
	cache      cache.ICache
	actions    actions.IAction
	dogLoader  *loaders.DogLoader
	vatLoader  *loaders.VatLoader
	processing sync.Mutex
	indexer    *chain.Indexer
	store      store.IStore
}

func NewVaultsChecker(cache cache.ICache, actions actions.IAction, dogLoader *loaders.DogLoader, vatLoader *loaders.VatLoader, indexer *chain.Indexer, store store.IStore) *VaultChecker {
	vaultChecker := &VaultChecker{
		cache:      cache,
		actions:    actions,
		dogLoader:  dogLoader,
		vatLoader:  vatLoader,
		processing: sync.Mutex{},
		indexer:    indexer,
		store:      store,
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

		// vatIlk.rate changes every block, so try to read from cache if it doesn't exist get ilk from vat and save it in cache for 12 seconds
		vatIlk, err := vc.cache.GetVatIlkByName(context.Background(), ilkName)
		if errors.Is(err, cache.ErrIlkNotFound) {
			vatIlk, err = vc.vatLoader.GetIlkById(context.Background(), ilkId)
			if err != nil {
				log.Printf("error in getting ilk from vat. ilkName:%s  error:%v", ilkName, err)
				continue
			}
			err := vc.cache.SaveVatIlk(context.Background(), *vatIlk, 12*time.Second)
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
			bark := store.NewBark(ilkId, vault.UrnAddress).ToDomain()
			tx, err := vc.actions.Bark(bark)
			if err != nil {
				log.Println("error in sending bark transaction.", err)
				continue
			}

			err, txId := vc.store.CreateTransaction(context.Background(), tx)
			if err != nil {
				log.Println("error in saving bark transaction.", err)
				continue
			}
			_, err = vc.store.CreateBark(context.Background(), *bark, int64(txId))
			if err != nil {
				log.Println("error in saving bark.", err)
				continue
			}

			receipt, header, err := vc.indexer.WaitForReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Println("error in getting bark transaction receipt.", err)
				continue
			}
			log.Printf("Bark transaction mined. TxHash:%s BlockNumber:%d BlockHash:%s", receipt.TxHash.Hex(), header.Number, header.Hash().Hex())

			err = vc.store.UpdateTransactionBlock(
				context.Background(),
				txId,
				receipt,
				header.Time,
				*receipt.BlockNumber,
				receipt.BlockHash)

			if err != nil {
				log.Println("error in updating bark transaction receipt.", err)
				continue
			}

			// contract update ilk.dirt after bark
			delete(dogIlks, ilkName)
		}
	}

	log.Printf("Checked %d urns", len(vaults))
}

func canBark(urn entities.Urn, vatIlk entities.VatIlk, dogHole, dogDirt, ilkHole, ilkDirt, chop *big.Int) bool {
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
