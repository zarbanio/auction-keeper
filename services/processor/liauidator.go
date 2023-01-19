package processor

import (
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
)

var (
	Decimals9  = math.BigPow(10, 9)
	Decimals18 = math.BigPow(10, 18)
	Decimals27 = math.BigPow(10, 27)
)

type LiquidatorConfig struct {
	MinProfitPercentage *big.Int
	MinLotZarValue      *big.Int
	MaxLotZarValue      *big.Int
	ProfitAddress       common.Address
}

type LiquidatorProcessor struct {
	collateralsProcessor map[string]collateralProcessor
	config               LiquidatorConfig
	processing           sync.Mutex
}

func NewLiquidatorProcessor(eth *ethclient.Client, collaterals map[string]entities.Collateral, liquidatorConfig *LiquidatorConfig) *LiquidatorProcessor {
	liquidatorProcessor := &LiquidatorProcessor{
		config: *liquidatorConfig,
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

func (lp *LiquidatorProcessor) StartProcessing() {
	lp.processing.Lock()
	defer lp.processing.Unlock()

	minProfitPercentage := new(big.Int).Mul(lp.config.MinProfitPercentage, Decimals18)

	for _, cp := range lp.collateralsProcessor {
		cp.processCollateral(minProfitPercentage, lp.config.MinLotZarValue, lp.config.MaxLotZarValue, lp.config.ProfitAddress)
	}
}
