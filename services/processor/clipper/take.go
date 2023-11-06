package processor

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/collateral"
	inputMethods "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/services/processor"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/uniswap_v3"
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

type ClipperTakeService struct {
	eth        *ethclient.Client
	clipper    *clipper.Clipper
	sender     sender.Sender
	collateral collateral.Collateral
	cp         *processor.CollateralProcessor
}

func NewClipperTakeService(
	eth *ethclient.Client,
	clipperAddr common.Address,
	sender sender.Sender,
	collateral collateral.Collateral,
) *ClipperTakeService {

	c, err := clipper.NewClipper(clipperAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &ClipperTakeService{
		eth:     eth,
		clipper: c,
		sender:  sender,
		cp:      processor.NewCollateralProcessor(eth, collateral),
	}
}

func (s *ClipperTakeService) Start(ctx context.Context, minProfitPercentage, minLotZarValue, maxLotZarValue *big.Int, profitAddress common.Address) error {
	currentTime, err := s.getCurrentTime(ctx)
	if err != nil {
		return nil
	}

	// 1) Get the IDs of the auctions
	auctionIds, err := s.clipper.List(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}

	for id, auctionId := range auctionIds {
		log.Printf("\tprocessing auction id: %d\n", id)
		auction, err := s.clipper.Sales(&bind.CallOpts{Context: ctx}, auctionId)
		if err != nil {
			return err
		}

		collateralPrice, err := s.collateral.Clipper.Abacus.Price(nil, auction.Top, big.NewInt(int64(currentTime)-auction.Tic.Int64()))
		if err != nil {
			log.Printf("error in get %s collateral price from abacus: %v\n", s.collateral.Name, err)
			continue
		}
		if collateralPrice.Cmp(big.NewInt(0)) <= 0 {
			log.Printf("Collateral price is zero")
			continue
		}

		// determine configured lot sizes in Gem terms
		minLot, maxLot := s.getMinAndMaxLot(minLotZarValue, maxLotZarValue, collateralPrice)

		// adjust lot based upon slice taken at the current auction price
		slice18 := math.BigMin(maxLot, auction.Lot)
		owe27 := new(big.Int).Div(new(big.Int).Mul(slice18, collateralPrice), Decimals18)
		tab27 := new(big.Int).Div(auction.Tab, Decimals18)

		// adjust covered debt to tab, such that slice better reflects amount of collateral we'd receive
		if owe27.Cmp(tab27) > 0 {
			owe27 = tab27
			slice18 = new(big.Int).Div(owe27, new(big.Int).Div(collateralPrice, Decimals18))
		} else if owe27.Cmp(tab27) < 0 && slice18.Cmp(auction.Lot) < 0 {
			chost27 := new(big.Int).Div(s.collateral.Clipper.Chost, Decimals18)
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
				log.Printf("Ignoring auction %d whose chost-adjusted slice of %d exceeds our maximum lot of %d}\n", auctionId, slice18, maxLot)
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
			log.Printf("Ignoring auction %d while slice is smaller than our minimum lot\n", auctionId)
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
		uniswapV3Proceeds, err := s.collateral.UniswapV3Quoter.GetQuotedAmountOut(lot, s.collateral.UniswapV3Path, s.collateral.Decimal)
		if err != nil {
			log.Printf("error in get uniswapV3 proceeds: %v\n", err)
			continue
		}
		minUniV3Proceeds := new(big.Int).Sub(uniswapV3Proceeds, minProfit)

		// Increase actual take amount to account for rounding errors and edge cases.
		// Do not increase too much to not significantly go over configured maxAmt.
		amt := new(big.Int).Div(new(big.Int).Mul(lot, big.NewInt(1000001)), big.NewInt(1000000))

		if debtToCover.Cmp(minUniV3Proceeds) <= 0 {
			// Uniswap tx executes only if the return amount also covers the minProfit %
			take := inputMethods.ClipperTake{
				Auction_id: auctionId,
				Amt:        amt,
				Max:        collateralPrice,
				Who:        s.collateral.UniswapV3Callee,
			}
			err = s.Run(take, minProfit, profitAddress, s.collateral.GemJoinAdapter)
			if err != nil {
				log.Printf("error in executeAuction: %v\n", err)
				continue
			}
		} else {
			log.Println("Uniswap V3 proceeds - profit amount is less than cost.")
		}
	}
	return nil
}

var (
	Uint256, _ = abi.NewType("uint256", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
	Address, _ = abi.NewType("address", "", nil)
)

// Run the service logic
func (s *ClipperTakeService) Run(take inputMethods.ClipperTake, minProfit *big.Int, profitAddr, gemJoinAdapter common.Address) error {
	// Uniswap v3 swap
	// typesArray := ['address', 'address', 'uint256', 'bytes', 'address'];
	args := abi.Arguments{
		{Name: "to", Type: Address},
		{Name: "gemJoin", Type: Address},
		{Name: "minProfit", Type: Uint256},
		{Name: "path", Type: Bytes},
		{Name: "charterManager", Type: Address},
	}
	route, err := uniswap_v3.GetRouter(s.collateral.UniswapV3Path)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("error in get route: %v", err))
	}

	flashData, err := args.Pack(profitAddr, gemJoinAdapter, minProfit, route, common.Address{0})
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("error in pack flash data: : %v", err))
	}

	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		return err
	}

	tx, err := s.clipper.Take(opts, take.Auction_id, take.Amt, take.Max, take.Who, flashData)
	if err != nil {
		return err
	}
	return s.sender.HandleSentTx(tx)
}

// Stop the service
func (s *ClipperTakeService) Stop() {
	panic("Implement me!")
}

func (s *ClipperTakeService) getCurrentTime(ctx context.Context) (uint64, error) {
	blockNum, err := s.eth.BlockNumber(context.Background())
	if err != nil {
		log.Println("error getting last block number.", err)
		return 0, err
	}
	block, err := s.eth.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		log.Println("error getting head.", err)
		return 0, err
	}

	currentTime := block.Header().Time

	return currentTime, nil
}

func (s *ClipperTakeService) getMinAndMaxLot(minLotZarValue, maxLotZarValue, collateralPrice *big.Int) (*big.Int, *big.Int) {
	// determine configured lot sizes in Gem terms
	minLotZarValue18 := new(big.Int).Mul(minLotZarValue, Decimals18)
	minLot := new(big.Int).Div(minLotZarValue18, new(big.Int).Div(collateralPrice, Decimals9))
	maxLotZarValue18 := new(big.Int).Mul(new(big.Int).Mul(maxLotZarValue, Decimals18), Decimals18)
	maxLot := new(big.Int).Div(maxLotZarValue18, new(big.Int).Div(collateralPrice, Decimals9))

	return minLot, maxLot
}
