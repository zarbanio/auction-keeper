package cmd

import (
	"context"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	"github.com/zarbanio/auction-keeper/bindings/zarban/median"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vow"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/collateral"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/domain/entities"
	"github.com/zarbanio/auction-keeper/services/actions"
	"github.com/zarbanio/auction-keeper/services/cachedeth"
	"github.com/zarbanio/auction-keeper/services/eventmanager"
	"github.com/zarbanio/auction-keeper/services/loaders"
	loggerPkg "github.com/zarbanio/auction-keeper/services/logger"
	"github.com/zarbanio/auction-keeper/services/processor"
	clipperServices "github.com/zarbanio/auction-keeper/services/processor/clipper"
	"github.com/zarbanio/auction-keeper/services/processor/clipper/vault"
	dogServices "github.com/zarbanio/auction-keeper/services/processor/dog"
	"github.com/zarbanio/auction-keeper/services/sender"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/services/uniswap_v3"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/chain"
	"github.com/zarbanio/auction-keeper/x/events"
)

func Execute() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := configs.ReadConfig("config.goerli.yaml")
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	// create a logger
	logger := loggerPkg.NewLogger(context.Background(), postgresStore)

	eth, err := ethclient.Dial(cfg.Network.Node.Api)
	if err != nil {
		log.Fatal(err)
	}
	// do database migrations
	err = postgresStore.Migrate(cfg.Postgres.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	// get application connection
	redisCache := cache.NewRedisStore(cfg.Redis.URL)
	bcache := cachedeth.NewBlockCache(eth)
	memCache := cache.NewMemCache()
	ceth := cachedeth.NewEthProxy(eth, postgresStore, bcache)

	// create new sender
	newSigner, err := signer.NewSigner(cfg.Wallet.Private, big.NewInt(cfg.Network.ChainId))
	if err != nil {
		log.Fatal(err)
	}

	newSender := sender.NewSender(newSigner, eth)

	// get loaders
	addressesLoader := loaders.NewAddressLoader(eth, memCache, cfg.Contracts.Deployment, cfg.Contracts.AddressProvider)
	addrs, err := addressesLoader.LoadAddresses(context.Background())
	if err != nil {
		log.Fatal("error loading addresses.", err)
	}

	addrs["cdp_manager"] = cfg.Contracts.CDPManager
	addrs["get_cdps"] = cfg.Contracts.GetCDPs
	addrs["ilk_registry"] = cfg.Contracts.IlkRegistry
	addrs["eth_a_join"] = cfg.Contracts.ETHAJoin
	addrs["eth_b_join"] = cfg.Contracts.ETHBJoin
	addrs["dai_a_join"] = cfg.Contracts.DAIAJoin
	addrs["dai_b_join"] = cfg.Contracts.DAIBJoin
	addrs["dai_median"] = cfg.Contracts.DAIMedian
	addrs["eth_median"] = cfg.Contracts.ETHMedian

	vaultLoader := loaders.NewVaultLoader(
		eth,
		postgresStore,
		addrs["cdp_manager"],
		addrs["get_cdps"],
		addrs["vat"],
	)

	ilksLoader := loaders.NewIlksLoader(
		ceth,
		postgresStore,
		addrs["vat"],
		addrs["jug"],
		addrs["spot"],
		addrs["dog"],
		addrs["ilk_registry"],
		[]common.Address{
			addrs["eth_a_join"],
			addrs["eth_b_join"],
			addrs["dai_a_join"],
			addrs["dai_b_join"],
		},
		map[common.Address]common.Address{
			cfg.Contracts.DAI:  addrs["dai_median"],
			cfg.Contracts.WETH: addrs["eth_median"],
		},
	)

	vatLoader := loaders.NewVatLoader(
		eth,
		cfg.Vat,
	)
	dogLoader := loaders.NewDogLoader(
		eth,
		cfg.Dog,
	)
	// services
	// 1. Dog Bark service
	dogBarkService := dogServices.NewDogBarkService(
		context.Background(),
		eth,
		postgresStore, addrs["dog"],
		addrs["spot"],
		vaultLoader, vatLoader,
		ilksLoader,
		newSender,
		logger)

	collaterals, err := getCollaterals(cfg, eth)
	if err != nil {
		log.Fatal(err)
	}

	clipperAddrs := make(map[string]common.Address)
	var clipperTakeServices []*clipperServices.ClipperTakeService
	var clipperRedoServices []*clipperServices.ClipperRedoService
	for _, c := range collaterals {
		clipperAddrs[c.Name] = c.Clipper.Address

		ct := clipperServices.NewClipperTakeService(eth, c.Clipper.Address, newSender, c, logger)
		cr := clipperServices.NewClipperRedoService(eth, c.Clipper.Address, newSender, c, logger)

		clipperTakeServices = append(clipperTakeServices, ct)
		clipperRedoServices = append(clipperRedoServices, cr)
	}

	eventManger := eventmanager.NewEventManager(
		postgresStore,
		ilksLoader,
		vaultLoader,
		dogBarkService,
		clipperTakeServices,
		clipperRedoServices,
	)

	actions, err := actions.NewActions(eth, newSender, cfg.Vat, cfg.Dog, cfg.Vow)
	if err != nil {
		log.Fatal(err)
	}

	/* -------------------------------------------------------------------------- */
	/*  get active auctions and add them in auction processor and start processor */
	liquidatorConfig := &processor.LiquidatorConfig{
		MinProfitPercentage: big.NewInt(cfg.Processor.MinProfitPercentage),
		MinLotZarValue:      big.NewInt(cfg.Processor.MinLotZarValue),
		MaxLotZarValue:      big.NewInt(cfg.Processor.MaxLotZarValue),
		ProfitAddress:       cfg.Wallet.Address,
	}
	liquidatorProcessor := processor.NewLiquidatorProcessor(eth, actions, collaterals, liquidatorConfig)
	err = getActiveAuctions(collaterals, liquidatorProcessor)
	if err != nil {
		log.Fatal(err)
	}

	handlers := getEventHandlers(eventManger, eth, addrs, clipperAddrs, liquidatorProcessor)

	var addressesArr []common.Address
	for k, v := range addrs {
		if k == "weth_gateway" || k == "weth" || k == "eth" || k == "zar" || k == "usdt" || k == "usdc" || k == "wbtc" || k == "dai" {
			continue
		}
		addressesArr = append(addressesArr, v)
	}

	eventHandlersMap := make(map[string]events.Handler)
	for _, h := range handlers {
		eventHandlersMap[h.ID()] = h
	}

	historyBlockPtr := createBlockPtrIfNotExists(postgresStore, "history", cfg.Indexer.StartBlock)
	startBlock, err := historyBlockPtr.Read()
	if err != nil {
		log.Fatal("error reading history block pointer.", err)
	}

	liveBlockPtr := createBlockPtrIfNotExists(postgresStore, "live", startBlock)

	indexer := chain.NewIndexer(ceth, cfg.Indexer.BlockInterval, liveBlockPtr, historyBlockPtr, cfg.Indexer.BlockRange, addressesArr, eventHandlersMap)

	for _, c := range collaterals {
		err = clipperAllowance(eth, c.Name, cfg.Vat, c.Clipper.Address, newSender, actions)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = zarJoinAllowance(eth, cfg.Vat, cfg.ZarJoin, newSender, actions)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(time.Duration(cfg.Times.LiquidatorTicker) * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				liquidatorProcessor.StartProcessing()
			}
		}
	}()

	vaultsChecker := vault.NewVaultsChecker(eth, redisCache, actions, dogLoader, vatLoader, indexer, postgresStore)
	vaultsCheckerTicker := time.NewTicker(time.Duration(cfg.Times.VaultTicker) * time.Second)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-vaultsCheckerTicker.C:
				vaultsChecker.Start()
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ticker.Stop()
}

func getEventHandlers(
	e *eventmanager.EventManager,
	eth *ethclient.Client,
	addresses map[string]common.Address,
	clipperAddrs map[string]common.Address,
	liquidatorProcessor *processor.LiquidatorProcessor) []events.Handler {

	eventHandlers := []events.Handler{
		vat.NewFrobHandler(addresses["vat"], eth, e.VatFrobCallback()),
		vat.NewForkHandler(addresses["vat"], eth, e.VatForkCallback()),
		vat.NewGrabHandler(addresses["vat"], eth, e.VatGrabCallback()),
		vow.NewFlogHandler(addresses["vow"], eth, e.VowFlogCallback()),
		vow.NewFessHandler(addresses["vow"], eth, e.VowFessCallback()),
		median.NewLogMedianPriceHandler(addresses["dai_median"], eth, e.MedianLogMedianPriceCallback()),
		median.NewLogMedianPriceHandler(addresses["eth_median"], eth, e.MedianLogMedianPriceCallback()),
	}

	for name, clipperAddr := range clipperAddrs {
		eventHandlers = append(eventHandlers, clipper.NewRedoHandler(clipperAddr, eth, e.ClipperRedoCallback(liquidatorProcessor, name)))
		eventHandlers = append(eventHandlers, clipper.NewTakeHandler(clipperAddr, eth, e.ClipperTakeCallback(liquidatorProcessor, name)))
		eventHandlers = append(eventHandlers, clipper.NewKickHandler(clipperAddr, eth, e.ClipperKickCallback(liquidatorProcessor, name)))
		// eventHandlers = append(eventHandlers, clipper.NewYankHandler(clipperAddr, eth, e.ClipperYankCallback(liquidatorProcessor, name)))

	}

	return eventHandlers
}

func getActiveAuctions(collaterals map[string]collateral.Collateral, liquidatorProcessor *processor.LiquidatorProcessor) error {
	println("Get active auctions from clippers.")

	for _, c := range collaterals {
		auctionsIds, err := c.ClipperLoader.GetActiveAuctions()
		if err != nil {
			return err
		}

		for _, auctionId := range auctionsIds {
			sale, err := c.ClipperLoader.GetSale(auctionId)
			if err != nil {
				continue
			}
			liquidatorProcessor.AddAuction(*sale, c.Name)
		}
	}

	println("Done\n")
	return nil
}

func getCollaterals(cfg configs.Config, eth *ethclient.Client) (map[string]collateral.Collateral, error) {
	uniswapV3Quoter, err := uniswap_v3.NewUniswapV3Quoter(eth, cfg.UniswapV3QuoterAddress)
	if err != nil {
		return nil, err
	}

	collaterals := make(map[string]collateral.Collateral)

	for _, collateralConfig := range cfg.Collaterals {
		clipperLoader, err := loaders.NewClipperLoader(eth, collateralConfig.Clipper)
		if err != nil {
			return nil, err
		}

		chost, err := clipperLoader.GetChost()
		if err != nil {
			return nil, err
		}

		abacus, err := clipperLoader.GetAbacusInstance()
		if err != nil {
			return nil, err
		}

		collaterals[collateralConfig.Name] = collateral.Collateral{
			Name:    collateralConfig.Name,
			Decimal: big.NewInt(collateralConfig.Decimals),
			Clipper: entities.Clipper{
				Address: collateralConfig.Clipper,
				Chost:   chost,
				Abacus:  abacus,
			},
			ClipperLoader:   clipperLoader,
			GemJoinAdapter:  collateralConfig.GemJoinAdapter,
			UniswapV3Callee: collateralConfig.UniswapV3Callee,
			UniswapV3Path:   collateralConfig.UniswapV3Path,
			UniswapV3Quoter: uniswapV3Quoter,
		}
	}

	return collaterals, nil
}

func clipperAllowance(eth *ethclient.Client, collateralName string, vatAddr, clipperAddr common.Address, sender sender.Sender, actions actions.IAction) error {
	vatInstance, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		return err
	}

	allowance, err := vatInstance.Can(nil, sender.GetAddress(), clipperAddr)
	if err != nil {
		return err
	}

	if allowance.Cmp(big.NewInt(1)) != 0 { // if allowance != 1
		log.Printf("HOPING %s CLIPPER IN VAT\n", collateralName)
		hope := store.NewHope(clipperAddr).ToDomain()
		tx, err := actions.Hope(hope)
		if err != nil {
			return err
		}
		log.Printf("HOPING %s CLIPPER IN VAT TX: %s\n", collateralName, tx.Hash().Hex())
	} else {
		log.Printf("%s Clipper is Hoped in VAT \n", collateralName)
	}
	return nil
}

func zarJoinAllowance(eth *ethclient.Client, vatAddr, zarJoinAddr common.Address, sender sender.Sender, actions actions.IAction) error {
	vatInstance, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		return err
	}

	allowance, err := vatInstance.Can(nil, sender.GetAddress(), zarJoinAddr)
	if err != nil {
		return err
	}

	if allowance.Cmp(big.NewInt(1)) != 0 { // if allowance != 1
		log.Println("HOPING ZAR_JOIN IN VAT")
		hope := store.NewHope(zarJoinAddr).ToDomain()
		tx, err := actions.Hope(hope)
		if err != nil {
			return err
		}
		log.Printf("HOPING ZAR_JOIN IN VAT TX: %s\n", tx.Hash().Hex())
	} else {
		log.Println("ZAR_JOIN is Hoped in VAT ")
	}
	return nil
}
