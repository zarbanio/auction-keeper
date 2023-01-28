package processor

import (
	"context"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/collateral"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var (
	Decimals9  = math.BigPow(10, 9)
	Decimals18 = math.BigPow(10, 18)
	Decimals27 = math.BigPow(10, 27)
)

type collateralProcessor struct {
	eth               *ethclient.Client
	collateral        collateral.Collateral
	auctionCollection AuctionCollection
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

func (cp *collateralProcessor) processCollateral(sender *transaction.Sender, minProfitPercentage, minLotZarValue, maxLotZarValue *big.Int, profitAddress common.Address) {
	fmt.Printf("processing opportunities for: %s", cp.collateral.Name)
	fmt.Printf("%d active auctions qty: %d\n", cp.collateral.Name, len(cp.auctionCollection.auctions))

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
		fmt.Printf("\tprocessing auction id: %d", id)

		// TODO
		//needsRedo, err := c.clipperLoader.GetAuctionStatus(a.Id)
		//if err != nil {
		//	continue
		//}

		collateralPrice, err := cp.collateral.Clipper.Abacus.Price(nil, auction.Top, big.NewInt(int64(currentTime-auction.Tic)))
		if err != nil {
			fmt.Printf("error in get %s collateral price from abacus: %v\n", err)
			continue
		}

		// determine configured lot sizes in Gem terms
		minLotZarValue18 := new(big.Int).Mul(minLotZarValue, Decimals18)
		minLot := new(big.Int).Div(minLotZarValue18, new(big.Int).Div(collateralPrice, Decimals9))
		maxLotZarValue18 := new(big.Int).Mul(maxLotZarValue, Decimals18)
		maxLot := new(big.Int).Div(maxLotZarValue18, new(big.Int).Div(collateralPrice, Decimals9))

		// adjust lot based upon slice taken at the current auction price
		slice18 := math.BigMin(maxLot, auction.Lot)
		owe27 := new(big.Int).Div(new(big.Int).Mul(slice18, collateralPrice), Decimals18)
		tab27 := new(big.Int).Div(auction.Tab, Decimals18)

		// adjust covered debt to tab, such that slice better reflects amount of collateral we'd receive
		if owe27.Cmp(tab27) > 0 {
			owe27 = tab27
			slice18 = new(big.Int).Mul(owe27, new(big.Int).Mul(collateralPrice, Decimals18))
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
				fmt.Printf("Ignoring auction %d whose chost-adjusted slice of %d exceeds our maximum lot of %d}\n", auction.Id, slice18, maxLot)
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
			fmt.Printf("Ignoring auction %d while slice is smaller than our minimum lot\n", auction.Id)
			// slice approaches lot as auction price decreases towards owe == tab
			continue
		}

		// Find the minimum effective exchange rate between collateral/Zar
		// e.x. ETH price 1000 Zar -> minimum profit of 1% -> new ETH price is 1000*1.01 = 1010
		calcMinProfit45 := new(big.Int).Mul(owe27, minProfitPercentage)
		totalMinProfit45 := new(big.Int).Sub(calcMinProfit45, new(big.Int).Mul(owe27, Decimals18))
		minProfit := new(big.Int).Div(totalMinProfit45, Decimals27)

		//debtToCover := new(big.Int).Div(owe27, Decimals9) // TODO

		// Increase actual take amount to account for rounding errors and edge cases.
		// Do not increase too much to not significantly go over configured maxAmt.
		amt := new(big.Int).Div(lot, big.NewInt(1000001))

		// TODO: print Auction Summary

		cp.executeAuction(sender, auction.Id, amt, collateralPrice, minProfit, profitAddress, cp.collateral.GemJoinAdapter, cp.collateral.UniswapV3Callee)
	}
}

var (
	Uint256, _ = abi.NewType("uint256", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
	Address, _ = abi.NewType("address", "", nil)
)

func (cp *collateralProcessor) executeAuction(sender *transaction.Sender, auctionId, amt, maxPrice, _minProfit *big.Int, _profitAddr, _gemJoinAdapter, exchangeCalleeAddress common.Address) {

	// Uniswap v3 swap
	// typesArray := ['address', 'address', 'uint256', 'bytes', 'address'];
	args := abi.Arguments{
		{Name: "to", Type: Address},
		{Name: "gemJoin", Type: Address},
		{Name: "minProfit", Type: Uint256},
		{Name: "path", Type: Bytes},
		{Name: "charterManager", Type: Address},
	}

	route := []byte{} // TODO
	flashData, err := args.Pack(_profitAddr, _gemJoinAdapter, _minProfit, route, common.Address{0})
	if err != nil {
		fmt.Printf("error in pack flash data: ", err)
	}

	err = sender.SendTakeTx(cp.collateral.ClipperLoader.Clipper, auctionId, amt, maxPrice, exchangeCalleeAddress, flashData)
	if err != nil {
		fmt.Printf("error in sending take transaction: ", err)
	}

}
