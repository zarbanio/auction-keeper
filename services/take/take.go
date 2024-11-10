package take

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/abacus"
	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/domain/entities"
	inputMethods "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/services/logger"
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

type Options struct {
	IlkName       string
	Eth           *ethclient.Client
	ClipperAddr   common.Address
	Quoter        *uniswap_v3.UniswapV3Quoter
	Path          []entities.UniswapV3Route
	Callee        common.Address
	GemJoin       common.Address
	AssetDecimals *big.Int
	Sender        sender.Sender
	Logger        *logger.Logger
	UseUniswap    bool
}

type Option func(*Options)

func WithIlkName(ilkName string) Option {
	return func(opts *Options) {
		opts.IlkName = ilkName
	}
}

func WithEthClient(eth *ethclient.Client) Option {
	return func(opts *Options) {
		opts.Eth = eth
	}
}

func WithClipperAddr(clipperAddr common.Address) Option {
	return func(opts *Options) {
		opts.ClipperAddr = clipperAddr
	}
}

func WithQuoter(quoter *uniswap_v3.UniswapV3Quoter) Option {
	return func(opts *Options) {
		opts.Quoter = quoter
	}
}

func WithPath(path []entities.UniswapV3Route) Option {
	return func(opts *Options) {
		opts.Path = path
	}
}

func WithCallee(callee common.Address) Option {
	return func(opts *Options) {
		opts.Callee = callee
	}
}

func WithGemJoin(gemJoin common.Address) Option {
	return func(opts *Options) {
		opts.GemJoin = gemJoin
	}
}

func WithAssetDecimals(assetDecimals *big.Int) Option {
	return func(opts *Options) {
		opts.AssetDecimals = assetDecimals
	}
}

func WithSender(sender sender.Sender) Option {
	return func(opts *Options) {
		opts.Sender = sender
	}
}

func WithLogger(logger *logger.Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}

func WithUseUniswap(useUniswap bool) Option {
	return func(opts *Options) {
		opts.UseUniswap = useUniswap
	}
}

type Service struct {
	IlkName string
	eth     *ethclient.Client
	sender  sender.Sender

	clip   *clipper.Clipper
	abacus *abacus.Abacus
	vat    *vat.Vat
	chost  *big.Int

	quoter        *uniswap_v3.UniswapV3Quoter
	path          []entities.UniswapV3Route
	callee        common.Address
	gemjoin       common.Address
	assetDecimals *big.Int
	l             *logger.Logger
	useUniswap    bool
}

func NewService(
	opts ...Option,
) *Service {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}

	clip, err := clipper.NewClipper(options.ClipperAddr, options.Eth)
	if err != nil {
		options.Logger.Logger.Fatal().Str("service", "take").Str("method", "NewClipperTakeService").Msg("err while instancing a clipper contract")
	}
	abacusAddr, err := clip.Calc(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		options.Logger.Logger.Fatal().Str("service", "take").Str("method", "NewClipperTakeService").Msg("err while getting the abacus address")
	}
	abacus, err := abacus.NewAbacus(abacusAddr, options.Eth)
	if err != nil {
		options.Logger.Logger.Fatal().Str("service", "take").Str("method", "NewClipperTakeService").Msg("err while instancing an abacus contract")
	}

	chost, err := clip.Chost(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		options.Logger.Logger.Fatal().Str("service", "take").Str("method", "NewClipperTakeService").Msg("err while getting the chost")
	}

	vatAddr, err := clip.Vat(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		options.Logger.Logger.Fatal().Str("service", "take").Str("method", "NewClipperTakeService").Msg("err while getting the vat address")
	}
	vat, err := vat.NewVat(vatAddr, options.Eth)
	if err != nil {
		options.Logger.Logger.Fatal().Str("service", "take").Str("method", "NewClipperTakeService").Msg("err while instancing a vat contract")
	}

	return &Service{
		IlkName:       options.IlkName,
		eth:           options.Eth,
		abacus:        abacus,
		clip:          clip,
		vat:           vat,
		chost:         chost,
		sender:        options.Sender,
		quoter:        options.Quoter,
		assetDecimals: options.AssetDecimals,
		callee:        options.Callee,
		gemjoin:       options.GemJoin,
		path:          options.Path,
		l:             options.Logger,
		useUniswap:    options.UseUniswap,
	}
}

func (s *Service) Start(ctx context.Context, minProfitPercentage, minLotZarValue, maxLotZarValue *big.Int, profitAddress common.Address) error {
	s.l.Logger.Info().
		Str("service", "take").
		Str("method", "start").
		Str("ilk", s.IlkName).
		Msg("take Service is starting")

	// 1) Get the IDs of active auctions
	activeAuctions, err := s.clip.List(&bind.CallOpts{Context: ctx})
	if err != nil {
		s.l.Logger.Error().
			Err(err).Str("service", "take").
			Str("method", "start").
			Str("ilk", s.IlkName).
			Msg("error while getting the list of active auctions")
		return err
	}

	for _, auctionId := range activeAuctions {
		err = s.TakeById(ctx, auctionId, minProfitPercentage, minLotZarValue, maxLotZarValue, profitAddress)
		if err != nil {
			s.l.Logger.Error().Err(err).
				Str("service", "take").
				Str("method", "start").
				Str("ilk", s.IlkName).
				Int64("auctionId", auctionId.Int64()).
				Msg("error while executing an auction")
			continue
		}

		s.l.Logger.Info().
			Str("service", "take").
			Str("method", "start").
			Str("ilk", s.IlkName).
			Int64("auctionId", auctionId.Int64()).
			Msg("auction executed successfully")
	}
	s.l.Logger.Info().
		Str("service", "take").
		Str("method", "start").
		Str("ilk", s.IlkName).
		Msg("take Service iteration completed")

	return nil
}

func (s *Service) TakeById(ctx context.Context, auctionId *big.Int, minProfitPercentage, minLotZarValue, maxLotZarValue *big.Int, profitAddress common.Address) error {
	currentTime, err := s.getCurrentTime(ctx)
	if err != nil {
		s.l.Logger.Error().Err(err).
			Str("service", "take").
			Str("method", "TakeById").
			Msg("error while getting current time")
		return err
	}

	auction, err := s.clip.Sales(&bind.CallOpts{Context: ctx}, auctionId)
	if err != nil {
		s.l.Logger.Error().Err(err).
			Str("service", "take").
			Str("method", "TakeById").
			Str("ilk", s.IlkName).
			Int64("auctionId", auctionId.Int64()).
			Msg("error while getting an auction")
		return err
	}

	collateralPrice, err := s.abacus.Price(&bind.CallOpts{Context: context.Background()}, auction.Top, big.NewInt(int64(currentTime)-auction.Tic.Int64()))
	if err != nil {
		s.l.Logger.Error().Err(err).
			Str("service", "take").
			Str("method", "TakeById").
			Str("ilk", s.IlkName).
			Int64("auctionId", auctionId.Int64()).
			Msg("error while getting the collateral price")
		return err
	}
	if collateralPrice.Cmp(big.NewInt(0)) <= 0 {
		s.l.Logger.Error().
			Err(err).
			Str("service", "take").
			Str("method", "TakeById").
			Str("ilk", s.IlkName).
			Msg("collateral price is less than or equal to zero")
		return err
	}

	s.l.Logger.Info().
		Str("service", "take").
		Str("method", "TakeById").
		Int64("auctionId", auctionId.Int64()).
		Str("ilk", s.IlkName).
		Msg("auction is active")
	s.l.Logger.Info().Str("service", "take").
		Str("method", "TakeById").
		Int64("auctionId", auctionId.Int64()).
		Str("collateralPrice", collateralPrice.String()).
		Str("ilk", s.IlkName).
		Msg("collateral price")

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
		chost27 := new(big.Int).Div(s.chost, Decimals18)
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
			log.Printf("Ignoring auction %d whose chost-adjusted slice of %d exceeds our maximum lot of %d\n", auctionId, slice18, maxLot)
			s.l.Logger.Info().Str("service", "take").
				Str("method", "TakeById").
				Str("ilk", s.IlkName).
				Int64("auctionId", auctionId.Int64()).
				Str("slice18", slice18.String()).
				Str("maxLot", maxLot.String()).
				Msg("ignoring auction whose chost-adjusted slice exceeds our maximum lot")
			return err
		}
	}

	if slice18.Cmp(auction.Lot) > 0 {
		// HACK: I suspect the issue involves interplay between reading price from the abacus and not having multicall.
		slice18 = auction.Lot
		owe27 = new(big.Int).Div(new(big.Int).Mul(slice18, collateralPrice), Decimals18)
	}

	lot := slice18
	if lot.Cmp(minLot) < 0 {
		// slice approaches lot as auction price decreases towards owe == tab
		s.l.Logger.Info().
			Str("service", "take").
			Str("method", "TakeById").
			Str("ilk", s.IlkName).
			Int64("auctionId", auctionId.Int64()).
			Str("slice18", slice18.String()).
			Str("minLot", minLot.String()).
			Msg("ignoring auction while slice is smaller than our minimum lot")
		return err
	}

	// Find the minimum effective exchange rate between collateral/Zar
	// e.x. ETH price 1000 Zar -> minimum profit of 1% -> new ETH price is 1000*1.01 = 1010
	owe45 := new(big.Int).Mul(owe27, Decimals18)
	minProfitPercentage18 := new(big.Int).Mul(minProfitPercentage, Decimals15)

	calcMinProfit45 := new(big.Int).Mul(owe27, minProfitPercentage18) // owe * minProfitPercentage
	totalMinProfit45 := new(big.Int).Sub(calcMinProfit45, owe45)
	minProfit := new(big.Int).Div(totalMinProfit45, Decimals27)

	// Increase actual take amount to account for rounding errors and edge cases.
	// Do not increase too much to not significantly go over configured maxAmt.
	amt := new(big.Int).Div(new(big.Int).Mul(lot, big.NewInt(1000001)), big.NewInt(1000000))

	if s.useUniswap {
		debtToCover := new(big.Int).Div(owe27, Decimals9)

		// Determine proceeds from swapping gem for Zar on Uniswap
		uniswapV3Proceeds, err := s.quoter.GetQuotedAmountOut(amt, s.path, s.assetDecimals)
		if err != nil {
			s.l.Logger.Error().Err(err).Str("service", "take").Str("method", "Start").Msg("error while getting the uniswapV3 proceeds")
			return err
		}
		minUniV3Proceeds := new(big.Int).Sub(uniswapV3Proceeds, minProfit)

		if debtToCover.Cmp(minUniV3Proceeds) > 0 {
			s.l.Logger.Info().
				Str("service", "take").
				Str("method", "start").
				Str("ilk", s.IlkName).
				Int64("auctionId", auctionId.Int64()).
				Str("debtToCover", debtToCover.String()).
				Str("minUniV3Proceeds", minUniV3Proceeds.String()).
				Msg("profit amount is less than cost")
			return errors.New("profit amount is less than cost")
		}

		take := inputMethods.ClipperTake{
			AuctionId: auctionId,
			Amt:       amt,
			Max:       collateralPrice,
			Who:       s.callee,
		}
		err = s.TakeUsingUniswap(take, minProfit, profitAddress, s.gemjoin)
		if err != nil {
			s.l.Logger.Error().Err(err).
				Str("service", "take").
				Str("method", "TakeById").
				Str("ilk", s.IlkName).
				Int64("auctionId", auctionId.Int64()).
				Msg("error while executing an auction")
			return err
		}
	} else {
		take := inputMethods.ClipperTake{
			AuctionId: auctionId,
			Amt:       amt,
			Max:       collateralPrice,
			Who:       profitAddress, // be careful with this. user should have access to profitAddress to exit gems
		}
		err = s.TakeUsingWalletFunds(ctx, take)
		if err != nil {
			s.l.Logger.Error().Err(err).
				Str("service", "take").
				Str("method", "TakeById").
				Str("ilk", s.IlkName).
				Int64("auctionId", auctionId.Int64()).
				Msg("error while executing an auction")
			return err
		}
	}

	s.l.Logger.Info().
		Str("service", "take").
		Str("method", "TakeById").
		Str("ilk", s.IlkName).
		Msg("take Service iteration completed")

	return nil
}

func (s *Service) TakeUsingWalletFunds(ctx context.Context, take inputMethods.ClipperTake) error {
	s.l.Logger.Info().Str("service", "take").Str("method", "TakeUsingWalletFunds").Msg("Calling take method")
	err := s.Take(take)
	if err != nil {
		s.l.Logger.Error().Err(err).
			Str("service", "take").
			Str("method", "TakeUsingWalletFunds").
			Str("profitReceiver", take.Who.String()).
			Str("auctionId", take.AuctionId.String()).
			Msg("error while calling take")
		return err
	}

	return nil
}

var (
	Uint256, _ = abi.NewType("uint256", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
	Address, _ = abi.NewType("address", "", nil)
)

func (s *Service) TakeUsingUniswap(take inputMethods.ClipperTake, minProfit *big.Int, profitAddr, gemJoinAdapter common.Address) error {
	args := abi.Arguments{
		{Name: "to", Type: Address},
		{Name: "gemJoin", Type: Address},
		{Name: "minProfit", Type: Uint256},
		{Name: "path", Type: Bytes},
		{Name: "charterManager", Type: Address},
	}
	route, err := uniswap_v3.GetRouter(s.path)
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "take").Str("method", "take").Msg("error while getting the route from uniswapV3")
		return fmt.Errorf("error in get route: %v", err)
	}

	take.Data, err = args.Pack(profitAddr, gemJoinAdapter, minProfit, route, common.HexToAddress("0x"))
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "take").Str("method", "take").Msg("error while getting the flashData")
		return fmt.Errorf("error in pack flash data: : %v", err)
	}

	err = s.Take(take)
	if err != nil {
		s.l.Logger.Error().Err(err).
			Str("service", "take").
			Str("method", "TakeUsingUniswap").
			Str("profitReceiver", take.Who.String()).
			Str("auctionId", take.AuctionId.String()).
			Msg("error while calling take")
		return err
	}

	return nil
}

func (s *Service) Take(take inputMethods.ClipperTake) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "take").Str("method", "take").Msg("error while getting a transaction opts")
		return err
	}

	tx, err := s.clip.Take(opts, take.AuctionId, take.Amt, take.Max, take.Who, take.Data)
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "take").Str("method", "take").Msg("error while calling the clipper take method")
		return err
	}
	fmt.Println("take tx sent:", tx.Hash().String())
	return s.sender.HandleSentTx(tx)
}

// Stop the service
func (s *Service) Stop() {
	s.l.Logger.Panic().Str("service", "take").Str("method", "Stop").Msg("Implement me!")
}

func (s *Service) getCurrentTime(ctx context.Context) (uint64, error) {
	blockNum, err := s.eth.BlockNumber(ctx)
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "take").Str("method", "getCurrentTime").Msg("error while getting the last block number")
		return 0, err
	}
	block, err := s.eth.BlockByNumber(ctx, big.NewInt(int64(blockNum)))
	if err != nil {
		s.l.Logger.Error().Err(err).Str("service", "take").Str("method", "getCurrentTime").Msg("error while getting the head")
		return 0, err
	}

	return block.Header().Time, nil
}

func (s *Service) getMinAndMaxLot(minLotZarValue, maxLotZarValue, collateralPrice *big.Int) (*big.Int, *big.Int) {
	// determine configured lot sizes in Gem terms
	minLotZarValue18 := new(big.Int).Mul(minLotZarValue, Decimals18)
	minLot := new(big.Int).Div(minLotZarValue18, new(big.Int).Div(collateralPrice, Decimals9))
	maxLotZarValue18 := new(big.Int).Mul(new(big.Int).Mul(maxLotZarValue, Decimals18), Decimals18)
	maxLot := new(big.Int).Div(maxLotZarValue18, new(big.Int).Div(collateralPrice, Decimals9))

	return minLot, maxLot
}
