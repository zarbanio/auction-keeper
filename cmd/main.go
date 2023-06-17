package cmd

import (
	"context"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	clipper "github.com/zarbanio/auction-keeper/bindings/clip"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/vat"
	"github.com/zarbanio/auction-keeper/bindings/vow"
	"github.com/zarbanio/auction-keeper/cache"
	"github.com/zarbanio/auction-keeper/collateral"
	"github.com/zarbanio/auction-keeper/configs"
	"github.com/zarbanio/auction-keeper/domain/entities"
	"github.com/zarbanio/auction-keeper/services/actions"
	"github.com/zarbanio/auction-keeper/services/callbacks"
	"github.com/zarbanio/auction-keeper/services/jobs"
	"github.com/zarbanio/auction-keeper/services/loaders"
	"github.com/zarbanio/auction-keeper/services/processor"
	"github.com/zarbanio/auction-keeper/services/processor/clipper/vault"
	"github.com/zarbanio/auction-keeper/services/transaction"
	"github.com/zarbanio/auction-keeper/services/uniswap_v3"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/chain"
	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/zarbanio/auction-keeper/x/messages"
	"github.com/zarbanio/auction-keeper/x/pubsub"
	"github.com/zarbanio/auction-keeper/x/pubsub/gochannel"
)

func startSubscribeEvents(ps pubsub.Pubsub, redisCache cache.ICache, vaultLoader *loaders.VaultLoader) {
	log.Println("subscribe Jobs for goChanel PubSub events.")
	_ = ps.Subscribe(context.Background(), "events.vat.frobs", func(msg *messages.Message) error {
		return jobs.Frobs(msg, redisCache, vaultLoader)
	})
	_ = ps.Subscribe(context.Background(), "events.vat.forks", func(msg *messages.Message) error {
		return jobs.Forks(msg, redisCache, vaultLoader)
	})
	_ = ps.Subscribe(context.Background(), "events.vat.grabs", func(msg *messages.Message) error {
		return jobs.Grabs(msg, redisCache, vaultLoader)
	})
	_ = ps.Subscribe(context.Background(), "events.vow.fess", func(msg *messages.Message) error {
		return jobs.Fess(msg, redisCache)
	})
	_ = ps.Subscribe(context.Background(), "events.vow.flog", func(msg *messages.Message) error {
		return jobs.Flog(msg, redisCache)
	})
}

func getEventHandlers(ps pubsub.Pubsub, eth *ethclient.Client, liquidatorProcessor *processor.LiquidatorProcessor, collaterals map[string]collateral.Collateral, vatAddress, vowAddress common.Address, startBlockNumber uint64) []events.Handler {
	var handlers []events.Handler

	for _, v := range collaterals {
		handlers = append(handlers, clipper.NewKickHandler(v.Clipper.Address, eth, callbacks.ClipperKickCallback(liquidatorProcessor, v.Name, startBlockNumber)))
		handlers = append(handlers, clipper.NewRedoHandler(v.Clipper.Address, eth, callbacks.ClipperRedoCallback(liquidatorProcessor, v.Name, startBlockNumber)))
		handlers = append(handlers, clipper.NewTakeHandler(v.Clipper.Address, eth, callbacks.ClipperTakeCallback(liquidatorProcessor, v.Name, startBlockNumber)))
	}
	handlers = append(handlers, vat.NewFrobHandler(vatAddress, eth, callbacks.VatFrobCallback(ps, 0)))
	handlers = append(handlers, vat.NewForkHandler(vatAddress, eth, callbacks.VatForkCallback(ps, 0)))
	handlers = append(handlers, vat.NewGrabHandler(vatAddress, eth, callbacks.VatGrabCallback(ps, 0)))
	handlers = append(handlers, vow.NewFessHandler(vowAddress, eth, callbacks.VowFessCallback(ps, 0)))
	handlers = append(handlers, vow.NewFlogHandler(vowAddress, eth, callbacks.VowFlogCallback(ps, 0)))
	return handlers
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
				return err
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

func clipperAllowance(eth *ethclient.Client, collateralName string, vatAddr, clipperAddr common.Address, sender *transaction.Sender, actions actions.IAction) error {
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

func zarJoinAllowance(eth *ethclient.Client, vatAddr, zarJoinAddr common.Address, sender *transaction.Sender, actions actions.IAction) error {
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

func Execute() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := configs.ReadConfig("config.yaml")
	postgresStore := store.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)

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

	// get loaders
	vaultLoader := loaders.NewVaultLoader(
		eth,
		cfg.Vat,
	)
	vatLoader := loaders.NewVatLoader(
		eth,
		cfg.Vat,
	)
	dogLoader := loaders.NewDogLoader(
		eth,
		cfg.Dog,
	)

	// This GoChannel is not persistent.
	//That means if you send a message to a topic to which no subscriber is subscribed, that message will be discarded.
	ps := gochannel.NewGoChannel(128)
	// start subscribe on chanels
	startSubscribeEvents(ps, redisCache, vaultLoader)

	blockPtr := NewDBBlockPointer(postgresStore, cfg.Indexer.BlockPtr)
	if !blockPtr.Exists() {
		log.Println("block pointer doest not exits. creating a new one.")
		err := blockPtr.Create()
		if err != nil {
			log.Fatal("error creating block pointer.", err)
		}
		log.Println("new block pointer created.", cfg.Indexer.BlockPtr)
	}

	collaterals, err := getCollaterals(cfg, eth)
	if err != nil {
		log.Fatal(err)
	}
	var addresses []common.Address
	for _, c := range collaterals {
		addresses = append(addresses, c.Clipper.Address)
	}
	addresses = append(addresses, cfg.Vat, cfg.Vow, cfg.Dog, cfg.Flopper)

	sender, err := transaction.NewSender(eth, cfg.Wallet.Private, big.NewInt(cfg.Network.ChainId))
	if err != nil {
		log.Fatal(err)
	}
	actions, err := actions.NewActions(eth, sender, cfg.Vat, cfg.Dog, cfg.Vow)
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

	lastBlack, err := eth.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	handlers := getEventHandlers(ps, eth, liquidatorProcessor, collaterals, cfg.Vat, cfg.Vow, lastBlack)
	eventHandlersMap := make(map[string]events.Handler)
	for _, h := range handlers {
		eventHandlersMap[h.ID()] = h
	}

	indexer := chain.NewIndexer(eth, chain.NewBlockCache(eth), cfg.Indexer.BlockInterval, blockPtr, addresses, eventHandlersMap)

	for _, c := range collaterals {
		err = clipperAllowance(eth, c.Name, cfg.Vat, c.Clipper.Address, sender, actions)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = zarJoinAllowance(eth, cfg.Vat, cfg.ZarJoin, sender, actions)
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

	vaultsChecker := vault.NewVaultsChecker(redisCache, actions, dogLoader, vatLoader, indexer, postgresStore)
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
