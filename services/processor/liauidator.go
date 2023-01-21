package processor

import (
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
)

type LiquidatorConfig struct {
	MinProfitPercentage *big.Int
	MinLotZarValue      *big.Int
	MaxLotZarValue      *big.Int
	ProfitAddress       common.Address
}

type LiquidatorProcessor struct {
	collateralsProcessor map[string]collateralProcessor
	config               *LiquidatorConfig
	sender               *transaction.Sender
	processing           sync.Mutex
}

func NewLiquidatorProcessor(eth *ethclient.Client, sender *transaction.Sender, collaterals map[string]entities.Collateral, liquidatorConfig *LiquidatorConfig) *LiquidatorProcessor {
	liquidatorProcessor := &LiquidatorProcessor{
		config: liquidatorConfig,
		sender: sender,
	}

	for key, collateral := range collaterals {
		cp := collateralProcessor{
			eth:        eth,
			collateral: collateral,
		}
		liquidatorProcessor.collateralsProcessor[key] = cp
	}

	return liquidatorProcessor
}

func (lp *LiquidatorProcessor) AddAuction(auction entities.Auction, collateralName string) {
	cp, _ := lp.collateralsProcessor[collateralName]
	cp.addAuction(auction)
}

func (lp *LiquidatorProcessor) DeleteAuction(id *big.Int, collateralName string) {
	cp, _ := lp.collateralsProcessor[collateralName]
	cp.deleteAuction(id)
}

func (lp *LiquidatorProcessor) UpdateAuctionAfterTake(id, tab, lot *big.Int, collateralName string) {
	cp, _ := lp.collateralsProcessor[collateralName]
	cp.updateAuctionAfterTake(id, tab, lot)
}

func (lp *LiquidatorProcessor) UpdateAuctionAfterRedo(id, top *big.Int, tic uint64, collateralName string) {
	cp, _ := lp.collateralsProcessor[collateralName]
	cp.updateAuctionAfterRedo(id, top, tic)
}

func (lp *LiquidatorProcessor) StartProcessing() {
	lp.processing.Lock()
	defer lp.processing.Unlock()

	minProfitPercentage := new(big.Int).Mul(lp.config.MinProfitPercentage, Decimals18)

	for _, cp := range lp.collateralsProcessor {
		cp.processCollateral(lp.sender, minProfitPercentage, lp.config.MinLotZarValue, lp.config.MaxLotZarValue, lp.config.ProfitAddress)
	}
}
