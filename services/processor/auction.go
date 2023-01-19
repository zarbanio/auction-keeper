package processor

import (
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"math/big"
	"sync"
)

type auctionIterator struct {
	auctions       map[*big.Int]entities.Auction
	updateAuctions sync.Locker
}

func (ai *auctionIterator) AuctionIsExist(Id *big.Int) bool {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	if _, ok := ai.auctions[Id]; ok {
		return true
	}
	return false
}

func (ai *auctionIterator) AddAuction(auction entities.Auction) {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	ai.auctions[auction.Id] = auction
}

func (ai *auctionIterator) UpdateAuctionAfterTake(Id *big.Int, Tab *big.Int, Lot *big.Int) {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	auction := ai.auctions[Id]
	auction.Tab = Tab
	auction.Lot = Lot
}

func (ai *auctionIterator) UpdateAuctionAfterRedo(Id *big.Int, Top *big.Int, Tic uint64) {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	auction := ai.auctions[Id]
	auction.Top = Top
	auction.Tic = Tic
}

func (ai *auctionIterator) DeleteAuction(Id *big.Int) {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	delete(ai.auctions, Id)
}

//func (ap *AuctionProcessor) StartProcessXana() {
//	ap.processAuctions.Lock()
//	defer ap.processAuctions.Unlock()
//
//	blockNum, err := ap.eth.BlockNumber(context.Background())
//	if err != nil {
//		log.Println("error getting last block number.", err)
//		return
//	}
//	block, err := ap.eth.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
//	if err != nil {
//		log.Println("error getting head.", err)
//		return
//	}
//
//	currentTime := block.Header().Time
//	minProfitPercentage := new(big.Int).Mul(ap.config.MinProfitPercentage, Decimals18)
//
//	for _, c := range ap.clippers {
//		c.updateAuctions.Lock()
//
//		fmt.Println("processing opportunities for:", c.clipperLoader.Name)
//		fmt.Println(c.clipperLoader.Name, "\tactive auctions qty: ", len(c.auctions))
//
//		for _, a := range c.auctions {
//			fmt.Println("\tprocessing auction id:", a.Id)
//
//			// TODO
//			//needsRedo, err := c.clipperLoader.GetAuctionStatus(a.Id)
//			//if err != nil {
//			//	continue
//			//}
//
//			collateralPrice, err := c.clipperLoader.Abacus.Price(nil, a.Top, big.NewInt(int64(currentTime-a.Tic)))
//			if err != nil {
//				continue
//			}
//
//			// determine configured lot sizes in Gem terms
//			minLotDaiValue := new(big.Int).Mul(ap.config.MinLotDaiValue, Decimals18)
//			minLot := new(big.Int).Div(minLotDaiValue, new(big.Int).Div(collateralPrice, Decimals9))
//			maxLotDaiValue := new(big.Int).Mul(ap.config.MaxLotDaiValue, Decimals18)
//			maxLot := new(big.Int).Div(maxLotDaiValue, new(big.Int).Div(collateralPrice, Decimals9))
//
//			// adjust lot based upon slice taken at the current auction price
//			slice18 := math.BigMin(maxLot, a.Lot)
//			owe27 := new(big.Int).Div(new(big.Int).Mul(slice18, collateralPrice), Decimals18)
//			tab27 := new(big.Int).Div(a.Tab, Decimals18)
//
//			// adjust covered debt to tab, such that slice better reflects amount of collateral we'd receive
//			if owe27.Cmp(tab27) > 0 {
//				owe27 = tab27
//				slice18 = new(big.Int).Mul(owe27, new(big.Int).Mul(collateralPrice, Decimals18))
//			} else if owe27.Cmp(tab27) < 0 && slice18.Cmp(a.Lot) < 0 {
//				chost27 := new(big.Int).Div(c.clipperLoader.Chost, Decimals18)
//				if new(big.Int).Sub(tab27, owe27).Cmp(chost27) < 0 {
//					if tab27.Cmp(chost27) <= 0 {
//						// If tab <= chost, buyers have to take the entire lot.
//						owe27 = chost27
//					} else {
//						// adjust amount to pay
//						owe27 = new(big.Int).Sub(tab27, chost27)
//					}
//					slice18 = new(big.Int).Div(owe27, new(big.Int).Div(collateralPrice, Decimals18))
//				}
//				if slice18.Cmp(maxLot) > 0 { // handle corner case where maxLotDaiValue is set too low
//					fmt.Printf("Ignoring auction %d whose chost-adjusted slice of %d exceeds our maximum lot of %d}\n", a.Id, slice18, maxLot)
//					continue
//				}
//			}
//
//			if slice18.Cmp(a.Lot) > 0 {
//				// HACK: I suspect the issue involves interplay between reading price from the abacus and not having multicall.
//				slice18 = a.Lot
//				owe27 = new(big.Int).Div(new(big.Int).Mul(slice18, collateralPrice), Decimals18)
//			}
//
//			lot := slice18
//			if lot.Cmp(minLot) < 0 {
//				fmt.Printf("Ignoring auction %d while slice is smaller than our minimum lot\n", a.Id)
//				// slice approaches lot as auction price decreases towards owe == tab
//				continue
//			}
//
//			// Find the minimum effective exchange rate between collateral/Dai
//			// e.x. ETH price 1000 DAI -> minimum profit of 1% -> new ETH price is 1000*1.01 = 1010
//			calcMinProfit45 := new(big.Int).Mul(owe27, minProfitPercentage)
//			totalMinProfit45 := new(big.Int).Sub(calcMinProfit45, new(big.Int).Mul(owe27, Decimals18))
//			minProfit := new(big.Int).Div(totalMinProfit45, Decimals27)
//
//			debtToCover := new(big.Int).Div(owe27, Decimals9)
//
//			// Increase actual take amount to account for rounding errors and edge cases.
//			// Do not increase too much to not significantly go over configured maxAmt.
//			amt := new(big.Int).Div(lot, big.NewInt(1000001))
//
//			profitAddr := common.Address{0}
//			executeAuction(a.Id, amt, collateralPrice, minProfit, profitAddr)
//		}
//
//		c.updateAuctions.Unlock()
//	}
//}
//
//func executeAuction(auctionId, _amt, _maxPrice, _minProfit *big.Int, _profitAddr, _gemJoinAdapter, exchangeCalleeAddress common.Address) {
//
//}
