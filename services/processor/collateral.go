package processor

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/collateral"
	"github.com/zarbanio/auction-keeper/domain/entities"
	"github.com/zarbanio/auction-keeper/services/actions"
	"github.com/zarbanio/auction-keeper/services/uniswap_v3"
	"github.com/zarbanio/auction-keeper/store"
)

var (
	Decimals9  = math.BigPow(10, 9)
	Decimals15 = math.BigPow(10, 15)
	Decimals18 = math.BigPow(10, 18)
	Decimals27 = math.BigPow(10, 27)

	FDecimals9  = new(big.Float).SetInt(Decimals9)
	FDecimals18 = new(big.Float).SetInt(Decimals18)
	FDecimals27 = new(big.Float).SetInt(Decimals27)
)

type collateralProcessor struct {
	eth               *ethclient.Client
	collateral        collateral.Collateral
	auctionCollection *AuctionCollection
}

func NewCollateralProcessor(eth *ethclient.Client, collateral collateral.Collateral) *collateralProcessor {
	return &collateralProcessor{
		eth:               eth,
		collateral:        collateral,
		auctionCollection: NewAuctionCollection(),
	}
}

func (cp *collateralProcessor) addAuction(auction entities.Auction) {
	cp.auctionCollection.addAuction(auction)
}

func (cp *collateralProcessor) deleteAuction(id *big.Int) {
	cp.auctionCollection.deleteAuction(id)
}

func (cp *collateralProcessor) updateAuctionAfterTake(id, tab, lot *big.Int) {
	cp.auctionCollection.updateAuctionAfterTake(id, tab, lot)
}

func (cp *collateralProcessor) updateAuctionAfterRedo(id, top *big.Int, tic uint64) {
	cp.auctionCollection.updateAuctionAfterRedo(id, top, tic)
}

func (cp *collateralProcessor) processCollateral(actions actions.IAction, minProfitPercentage, minLotZarValue, maxLotZarValue *big.Int, profitAddress common.Address) {
	log.Printf("processing opportunities for: %s\n", cp.collateral.Name)
	log.Printf("%s active auctions qty: %d\n", cp.collateral.Name, len(cp.auctionCollection.auctions))

	blockNum, err := cp.eth.BlockNumber(context.Background())
	if err != nil {
		log.Println("error getting last block number.", err)
		return
	}
	block, err := cp.eth.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		log.Println("error getting head.", err)
		return
	}

	currentTime := block.Header().Time

	for id, auction := range cp.auctionCollection.auctions {
		log.Printf("\tprocessing auction id: %d\n", id)

		needsRedo, err := cp.collateral.ClipperLoader.GetAuctionStatus(auction.Id)
		if err != nil {
			log.Printf("error in get auction status: %v\n", err)
			continue
		}
		if needsRedo {
			redo := store.NewRedo(*auction.Id).ToDomain()
			err := actions.Redo(cp.collateral.ClipperLoader.Clipper, redo)
			if err != nil {
				log.Printf("error in sending redo transaction: %v\n", err)
			}
			continue
		}

		collateralPrice, err := cp.collateral.Clipper.Abacus.Price(nil, auction.Top, big.NewInt(int64(currentTime-auction.Tic)))
		if err != nil {
			log.Printf("error in get %s collateral price from abacus: %v\n", cp.collateral.Name, err)
			continue
		}
		if collateralPrice.Cmp(big.NewInt(0)) <= 0 {
			log.Printf("Collateral price is zero")
			continue
		}

		// determine configured lot sizes in Gem terms
		minLotZarValue18 := new(big.Int).Mul(minLotZarValue, Decimals18)
		minLot := new(big.Int).Div(minLotZarValue18, new(big.Int).Div(collateralPrice, Decimals9))
		maxLotZarValue18 := new(big.Int).Mul(new(big.Int).Mul(maxLotZarValue, Decimals18), Decimals18)
		maxLot := new(big.Int).Div(maxLotZarValue18, new(big.Int).Div(collateralPrice, Decimals9))

		// adjust lot based upon slice taken at the current auction price
		slice18 := math.BigMin(maxLot, auction.Lot)
		owe27 := new(big.Int).Div(new(big.Int).Mul(slice18, collateralPrice), Decimals18)
		tab27 := new(big.Int).Div(auction.Tab, Decimals18)

		// adjust covered debt to tab, such that slice better reflects amount of collateral we'd receive
		if owe27.Cmp(tab27) > 0 {
			owe27 = tab27
			slice18 = new(big.Int).Div(owe27, new(big.Int).Div(collateralPrice, Decimals18))
		} else if owe27.Cmp(tab27) < 0 && slice18.Cmp(auction.Lot) < 0 {
			chost27 := new(big.Int).Div(cp.collateral.Clipper.Chost, Decimals18)
			if new(big.Int).Sub(tab27, owe27).Cmp(chost27) < 0 {
				if tab27.Cmp(chost27) <= 0 {
					// If tab <= chost, buyers have to take the entire lot.
					owe27 = chost27
				} else {
					// adjust amount to pay
					owe27 = new(big.Int).Sub(tab27, chost27)
				}
				slice18 = new(big.Int).Div(owe27, new(big.Int).Div(collateralPrice, Decimals18))
			}
			if slice18.Cmp(maxLot) > 0 { // handle corner case where maxLotZarValue is set too low
				log.Printf("Ignoring auction %d whose chost-adjusted slice of %d exceeds our maximum lot of %d}\n", auction.Id, slice18, maxLot)
				continue
			}
		}

		if slice18.Cmp(auction.Lot) > 0 {
			// HACK: I suspect the issue involves interplay between reading price from the abacus and not having multicall.
			slice18 = auction.Lot
			owe27 = new(big.Int).Div(new(big.Int).Mul(slice18, collateralPrice), Decimals18)
		}

		lot := slice18
		if lot.Cmp(minLot) < 0 {
			log.Printf("Ignoring auction %d while slice is smaller than our minimum lot\n", auction.Id)
			// slice approaches lot as auction price decreases towards owe == tab
			continue
		}

		// Find the minimum effective exchange rate between collateral/Zar
		// e.x. ETH price 1000 Zar -> minimum profit of 1% -> new ETH price is 1000*1.01 = 1010
		calcMinProfit45 := new(big.Int).Mul(owe27, minProfitPercentage)
		totalMinProfit45 := new(big.Int).Sub(calcMinProfit45, new(big.Int).Mul(owe27, Decimals18))
		minProfit := new(big.Int).Div(totalMinProfit45, Decimals27)

		debtToCover := new(big.Int).Div(owe27, Decimals9)

		// Determine proceeds from swapping gem for Zar on Uniswap
		uniswapV3Proceeds, err := cp.collateral.UniswapV3Quoter.GetQuotedAmountOut(lot, cp.collateral.UniswapV3Path, cp.collateral.Decimal)
		if err != nil {
			log.Printf("error in get uniswapV3 proceeds: %v\n", err)
			continue
		}
		minUniV3Proceeds := new(big.Int).Sub(uniswapV3Proceeds, minProfit)

		// Increase actual take amount to account for rounding errors and edge cases.
		// Do not increase too much to not significantly go over configured maxAmt.
		amt := new(big.Int).Div(new(big.Int).Mul(lot, big.NewInt(1000001)), big.NewInt(1000000))

		printAuctionSummary(cp.collateral.Name, auction, minLot, maxLot, lot, collateralPrice, debtToCover, minProfit, uniswapV3Proceeds, minUniV3Proceeds)

		if debtToCover.Cmp(minUniV3Proceeds) <= 0 {
			// Uniswap tx executes only if the return amount also covers the minProfit %
			err = cp.executeAuction(actions, auction.Id, amt, collateralPrice, minProfit, profitAddress, cp.collateral.GemJoinAdapter, cp.collateral.UniswapV3Callee)
			if err != nil {
				log.Printf("error in executeAuction: %v\n", err)
				continue
			}
		} else {
			log.Println("Uniswap V3 proceeds - profit amount is less than cost.")
		}
	}
}

var (
	Uint256, _ = abi.NewType("uint256", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
	Address, _ = abi.NewType("address", "", nil)
)

func (cp *collateralProcessor) executeAuction(actions actions.IAction, auctionId, amt, maxPrice, minProfit *big.Int, profitAddr, gemJoinAdapter, exchangeCalleeAddress common.Address) error {

	// Uniswap v3 swap
	// typesArray := ['address', 'address', 'uint256', 'bytes', 'address'];
	args := abi.Arguments{
		{Name: "to", Type: Address},
		{Name: "gemJoin", Type: Address},
		{Name: "minProfit", Type: Uint256},
		{Name: "path", Type: Bytes},
		{Name: "charterManager", Type: Address},
	}

	route, err := uniswap_v3.GetRouter(cp.collateral.UniswapV3Path)
	if err != nil {
		return errors.New(fmt.Sprintf("error in get route: %v", err))
	}

	flashData, err := args.Pack(profitAddr, gemJoinAdapter, minProfit, route, common.Address{0})
	if err != nil {
		return errors.New(fmt.Sprintf("error in pack flash data: : %v", err))
	}
	take := store.NewTake(auctionId, amt, maxPrice, exchangeCalleeAddress, flashData).ToDomain()
	err = actions.Take(cp.collateral.ClipperLoader.Clipper, take)
	if err != nil {
		return errors.New(fmt.Sprintf("error in sending take transaction: : %v", err))
	}
	return nil
}

func printAuctionSummary(collateralName string, auction entities.Auction, minLot, maxLot, lot, collateralPrice, debtToCover, minProfit, uniswapV3Proceeds, minUniV3Proceeds *big.Int) {
	fAuctionTab := new(big.Float).SetInt(auction.Tab)
	fAuctionLot := new(big.Float).SetInt(auction.Lot)

	fMinLot := new(big.Float).SetInt(minLot)
	fMaxLot := new(big.Float).SetInt(maxLot)

	fLot := new(big.Float).SetInt(lot)
	fCollateralPrice := new(big.Float).SetInt(collateralPrice)
	fDebtToCover := new(big.Float).SetInt(debtToCover)
	fMinProfit := new(big.Float).SetInt(minProfit)

	fUniswapV3Proceeds := new(big.Float).SetInt(uniswapV3Proceeds)
	fMinUniV3Proceeds := new(big.Float).SetInt(minUniV3Proceeds)

	log.Printf("  #%s Auction\n", collateralName)
	log.Printf("\tAuction Id:      %d\n", auction.Id)
	log.Printf("\tAuction Tab:     %.18f Zar\n", new(big.Float).Quo(new(big.Float).Quo(fAuctionTab, FDecimals27), FDecimals18))
	log.Printf("\tAuction Lot:     %.18f\n", new(big.Float).Quo(fAuctionLot, FDecimals18))
	log.Printf("\tConfigured Lot:  between %.18f and %.18f\n", new(big.Float).Quo(fMinLot, FDecimals18), new(big.Float).Quo(fMaxLot, FDecimals18))
	log.Printf("\tSlice to Take:   %.18f\n", new(big.Float).Quo(fLot, FDecimals18))

	log.Printf("\tAuction Price:   %.18f Zar\n\n", new(big.Float).Quo(new(big.Float).Quo(fCollateralPrice, FDecimals9), FDecimals18))
	log.Printf("\tDebt to Cover:   %.18f Zar\n", new(big.Float).Quo(fDebtToCover, FDecimals18))
	log.Printf("\tMinimum profit:  %.18f Zar\n", new(big.Float).Quo(fMinProfit, FDecimals18))
	log.Printf("\tUniswapV3 proceeds:  %.18f Zar\n", new(big.Float).Quo(fUniswapV3Proceeds, FDecimals18))
	log.Printf("\tLess min profit:     %.18f\n", new(big.Float).Quo(fMinUniV3Proceeds, FDecimals18))
}
