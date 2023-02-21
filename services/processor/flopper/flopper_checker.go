package flopper

import (
	"context"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/cache"
	"github.com/IR-Digital-Token/auction-keeper/domain/math"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"sync"
)

type FlopperChecker struct {
	eth           *ethclient.Client
	cache         cache.ICache
	sender        transaction.ISender
	vowAddress    common.Address
	vowLoader     *loaders.VowLoader
	vatLoader     *loaders.VatLoader
	flopperLoader *loaders.FlopperLoader
	processing    sync.Mutex
}

func NewFlopperChecker(eth *ethclient.Client, cache cache.ICache, sender transaction.ISender, vowAddr common.Address, vowLoader *loaders.VowLoader, vatLoader *loaders.VatLoader, flopperLoader *loaders.FlopperLoader) *FlopperChecker {
	flopperChecker := &FlopperChecker{
		eth:           eth,
		cache:         cache,
		sender:        sender,
		vowAddress:    vowAddr,
		vowLoader:     vowLoader,
		vatLoader:     vatLoader,
		flopperLoader: flopperLoader,
		processing:    sync.Mutex{},
	}

	return flopperChecker
}

func (fc *FlopperChecker) Start() {
	fc.processing.Lock()
	defer fc.processing.Unlock()

	isRely, err := fc.flopperLoader.IsRely(context.Background(), fc.vowAddress)
	if err != nil {
		log.Println("error in getting flop wards.", err)
		return
	}
	if !isRely {
		log.Println("vow is not authorized to kick on this flopper.")
		return
	}

	vowZarBalance, err := fc.vatLoader.GetZarBalance(context.Background(), fc.vowAddress)
	if err != nil {
		log.Println("error in getting vow zar balance.", err)
		return
	}

	vatSin, err := fc.vatLoader.GetSin(context.Background(), fc.vowAddress) // Total Deficit
	if err != nil {
		log.Println("error in getting vat.sin(vowAddress).", err)
		return
	}

	// Check if Vow has bad debt in excess: vowZarBalance < vatSin
	excessDebt := vowZarBalance.Cmp(vatSin) == -1
	if !excessDebt {
		return
	}

	vowSin, err := fc.vowLoader.GetSin(context.Background())
	if err != nil {
		log.Println("error in getting vow.Sin.", err)
		return
	}

	ash, err := fc.vowLoader.GetAsh(context.Background())
	if err != nil {
		log.Println("error in getting vow.Ash.", err)
		return
	}

	woe := calculateWoe(vatSin, vowSin, ash) // woe is unqueued, pre-auction debt

	sump, err := fc.vowLoader.GetSump(context.Background())
	if err != nil {
		log.Println("error in getting vow.sump.", err)
		return
	}

	wait, err := fc.vowLoader.GetWait(context.Background())
	if err != nil {
		log.Println("error in getting vow.wait.", err)
		return
	}

	// Check if Vow has enough bad debt and that we have enough dai balance
	if new(big.Int).Add(woe, vowSin).Cmp(sump) >= 0 { // woe + vowSin >= sump
		// We need to bring vowZarBalance to 0 and Woe to at least sump

		if vowZarBalance.Cmp(big.NewInt(0)) > 1 {
			err = fc.ReconcileDebt(vowZarBalance, ash, woe)
			if err != nil {
				log.Println("error in reconcile debt:", err)
				return
			}
		}

		// Convert enough sin in woe to have woe >= sump + vowZarBalance
		if woe.Cmp(new(big.Int).Add(sump, vowZarBalance)) < 0 { // if woe < (sump + vowZarBalance)
			eras, err := fc.cache.GetEras(context.Background())
			if err != nil {
				log.Println("error in get eras from redis:", err)
				return
			}
			if len(eras) == 0 {
				log.Println("eras length is zero", err)
				return
			}

			for _, era := range eras {
				latestBlock, err := fc.eth.BlockByNumber(context.Background(), nil)
				if err != nil {
					// Handle error
				}
				now := new(big.Int).SetUint64(latestBlock.Time())
				eraSin, err := fc.vowLoader.GetSinOf(context.Background(), era)
				if err != nil {
					log.Println("error in get era sin.", err)
					return
				}

				if eraSin.Cmp(math.Zero) > 0 && new(big.Int).Add(era, wait).Cmp(now) <= 0 { // if sin > 0 and era + wait <= now:
					txHash, err := fc.sender.Flog(era)
					if err != nil {
						log.Println("error in sending flog transaction.", err)
						return
					}
					fmt.Printf("\tFlog Transaction Hash: %s\n", txHash)

					// flog() sin until woe is above sump + joy
					vowZarBalance, err = fc.vatLoader.GetZarBalance(context.Background(), fc.vowAddress)
					if err != nil {
						log.Println("error in getting vow zar balance after flog.", err)
						return
					}

					woe, err = fc.GetWoe()
					if err != nil {
						log.Println("error in getting woe after flog.", err)
						return
					}
					if new(big.Int).Sub(woe, vowZarBalance).Cmp(sump) >= 0 {
						break
					}
				}
			}
		}

		// Reduce on-auction debt and reconcile remaining joy
		vowZarBalance, err = fc.vatLoader.GetZarBalance(context.Background(), fc.vowAddress)
		if err != nil {
			log.Println("error in getting vow zar balance.", err)
			return
		}

		if vowZarBalance.Cmp(math.Zero) > 0 {
			ash, err = fc.vowLoader.GetAsh(context.Background())
			if err != nil {
				log.Println("error in getting vow.Ash.", err)
				return
			}

			woe, err = fc.GetWoe()
			if err != nil {
				log.Println("error in getting woe after flog.", err)
				return
			}

			err = fc.ReconcileDebt(vowZarBalance, ash, woe)
			if err != nil {
				log.Println("error in reconcile debt:", err)
				return
			}

			vowZarBalance, err = fc.vatLoader.GetZarBalance(context.Background(), fc.vowAddress)
			if err != nil {
				log.Println("error in getting vow zar balance.", err)
				return
			}
		}

		woe, err = fc.GetWoe()
		if err != nil {
			log.Println("error in getting woe after flog.", err)
			return
		}

		if sump.Cmp(woe) <= 0 && vowZarBalance.Cmp(math.Zero) == 0 {
			txHash, err := fc.sender.Flop()
			if err != nil {
				log.Println("error in sending flop transaction.", err)
			}
			fmt.Printf("\tFlop Transaction Hash: %s\n", txHash)
		}
	}

}

func (fc *FlopperChecker) ReconcileDebt(zarBalance, ash, woe *big.Int) error {
	var err error

	if ash.Cmp(math.Zero) > 0 {
		if zarBalance.Cmp(ash) > 0 {
			txHash, err := fc.sender.Kiss(ash)
			if err != nil {
				return err
			}
			fmt.Printf("\tKiss Transaction Hash: %s\n", txHash)
		} else {
			txHash, err := fc.sender.Kiss(zarBalance)
			if err != nil {
				return err
			}
			fmt.Printf("\tKiss Transaction Hash: %s\n", txHash)
		}
	}
	if woe.Cmp(math.Zero) > 0 {
		zarBalance, err = fc.vatLoader.GetZarBalance(context.Background(), fc.vowAddress)
		if err != nil {
			log.Println("[ReconcileDebt] error in getting vow zar balance.", err)
			return err
		}

		if zarBalance.Cmp(woe) > 0 {
			txHash, err := fc.sender.Heal(woe)
			if err != nil {
				return err
			}
			fmt.Printf("\tHeal Transaction Hash: %s\n", txHash)
		} else {
			txHash, err := fc.sender.Heal(zarBalance)
			if err != nil {
				return err
			}
			fmt.Printf("\tHeal Transaction Hash: %s\n", txHash)
		}
	}

	return nil
}

func (fc *FlopperChecker) GetWoe() (*big.Int, error) {
	// Woe is unqueued, pre-auction debt

	vatSin, err := fc.vatLoader.GetSin(context.Background(), fc.vowAddress) // Total Deficit
	if err != nil {
		return nil, err
	}

	vowSin, err := fc.vowLoader.GetSin(context.Background())
	if err != nil {
		return nil, err
	}

	ash, err := fc.vowLoader.GetAsh(context.Background())
	if err != nil {
		return nil, err
	}

	return calculateWoe(vatSin, vowSin, ash), nil
}

func calculateWoe(vatSin, vowSin, ash *big.Int) *big.Int {
	return new(big.Int).Sub(new(big.Int).Sub(vatSin, vowSin), ash)
}
