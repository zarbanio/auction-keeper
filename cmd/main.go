package cmd

import (
	"context"
	"github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/collateral"
	"github.com/IR-Digital-Token/auction-keeper/configs"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/callbacks"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/IR-Digital-Token/auction-keeper/services/processor"
	"github.com/IR-Digital-Token/auction-keeper/services/router"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/IR-Digital-Token/x/chain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func registerEventHandlers(indexer *chain.Indexer, eth *ethclient.Client, liquidatorProcessor *processor.LiquidatorProcessor, collaterals map[string]collateral.Collateral) {
	println("Register callbacks on event triggers come from the indexer.")

	for _, v := range collaterals {
		indexer.RegisterEventHandler(clipper.NewKickHandler(v.Clipper.Address, eth, callbacks.ClipperKickCallback(liquidatorProcessor, v.Name)))
		indexer.RegisterEventHandler(clipper.NewRedoHandler(v.Clipper.Address, eth, callbacks.ClipperRedoCallback(liquidatorProcessor, v.Name)))
		indexer.RegisterEventHandler(clipper.NewTakeHandler(v.Clipper.Address, eth, callbacks.ClipperTakeCallback(liquidatorProcessor, v.Name)))
	}

	println("Done\n")
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
	collaterals := make(map[string]collateral.Collateral)

	type cConfig struct {
		name                 string
		clipperAddress       common.Address
		gemJoinAdapter       common.Address
		UniswapV3Callee      common.Address
		UniswapV3CalleeRoute []entities.UniswapRoute
	}
	collateralsConfig := make([]cConfig, 0)

	collateralsConfig = append(collateralsConfig, cConfig{
		"ETHA",
		cfg.Collaterals.ETHA.Clipper,
		cfg.Collaterals.ETHA.GemJoinAdapter,
		cfg.Collaterals.ETHA.UniswapV3Callee,
		[]entities.UniswapRoute{
			{
				cfg.Collaterals.ETHA.UniV3Path[0].Fee,
				cfg.Collaterals.ETHA.UniV3Path[0].TokenA,
				cfg.Collaterals.ETHA.UniV3Path[0].TokenB,
			},
			{
				cfg.Collaterals.ETHA.UniV3Path[1].Fee,
				cfg.Collaterals.ETHA.UniV3Path[1].TokenA,
				cfg.Collaterals.ETHA.UniV3Path[1].TokenB,
			},
		},
	})
	collateralsConfig = append(collateralsConfig, cConfig{
		"ETHB",
		cfg.Collaterals.ETHB.Clipper,
		cfg.Collaterals.ETHB.GemJoinAdapter,
		cfg.Collaterals.ETHB.UniswapV3Callee,
		[]entities.UniswapRoute{
			{
				cfg.Collaterals.ETHA.UniV3Path[0].Fee,
				cfg.Collaterals.ETHA.UniV3Path[0].TokenA,
				cfg.Collaterals.ETHA.UniV3Path[0].TokenB,
			},
			{
				cfg.Collaterals.ETHA.UniV3Path[1].Fee,
				cfg.Collaterals.ETHA.UniV3Path[1].TokenA,
				cfg.Collaterals.ETHA.UniV3Path[1].TokenB,
			},
		},
	})

	for _, collateralConfig := range collateralsConfig {
		clipperLoader, err := loaders.NewClipperLoader(eth, collateralConfig.clipperAddress)
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

		uniswapV3CalleeRoute, err := router.GetUniswapV3Router(collateralConfig.UniswapV3CalleeRoute)
		if err != nil {
			return nil, err
		}

		collaterals[collateralConfig.name] = collateral.Collateral{
			Name: collateralConfig.name,
			Clipper: entities.Clipper{
				Address: collateralConfig.clipperAddress,
				Chost:   chost,
				Abacus:  abacus,
			},
			ClipperLoader:        clipperLoader,
			GemJoinAdapter:       collateralConfig.gemJoinAdapter,
			UniswapV3Callee:      collateralConfig.UniswapV3Callee,
			UniswapV3CalleeRoute: uniswapV3CalleeRoute,
		}
	}

	return collaterals, nil
}

func Execute() {
	cfg := configs.ReadConfig("config.yaml")

	eth, err := ethclient.Dial(cfg.Network.Node.Api)
	if err != nil {
		panic(err)
	}

	blockPtr := chain.NewFileBlockPointer(".", "goerli.ptr", cfg.Indexer.BlockPtr)
	if !blockPtr.Exists() {
		log.Println("block pointer file doest not exits. creating a new one.")
		err = blockPtr.Create()
		if err != nil {
			panic(err)
		}

		log.Println("new block pointer file created.")
	}
	// write last block number in block pointer file
	lastBlack, err := eth.BlockNumber(context.Background())
	err = blockPtr.Update(lastBlack)
	if err != nil {
		panic(err)
	}

	/***************************************
	 			get collaterals
	***************************************/
	collaterals, err := getCollaterals(cfg, eth)
	if err != nil {
		panic(err)
	}

	/***************************************
	 			import wallet
	***************************************/
	sender, err := transaction.NewSender(eth, cfg.Wallet.Private, big.NewInt(cfg.Network.ChainId))
	if err != nil {
		panic(err)
	}

	/* -------------------------------------------------------------------------- */
	/*  get active auctions and add them in auction processor and start processor */
	liquidatorConfig := &processor.LiquidatorConfig{
		MinProfitPercentage: big.NewInt(cfg.Processor.MinProfitPercentage),
		MinLotZarValue:      big.NewInt(cfg.Processor.MinLotZarValue),
		MaxLotZarValue:      big.NewInt(cfg.Processor.MaxLotZarValue),
		ProfitAddress:       cfg.Wallet.Address,
	}
	liquidatorProcessor := processor.NewLiquidatorProcessor(eth, sender, collaterals, liquidatorConfig)
	err = getActiveAuctions(collaterals, liquidatorProcessor)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(60 * time.Second)
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
	/* -------------------------------------------------------------------------- */

	/* -------------------------------------------------------------------------- */
	/*                     register contract events on indexer                    */
	indexer := chain.NewIndexer(eth, blockPtr, cfg.Indexer.PoolSize)
	registerEventHandlers(indexer, eth, liquidatorProcessor, collaterals)

	for _, c := range collaterals {
		indexer.RegisterAddress(c.Clipper.Address)
	}
	/* -------------------------------------------------------------------------- */

	indexer.Init(cfg.Indexer.BlockInterval)
	go func() {
		for {
			err = indexer.Start()
			if err != nil {
				log.Println("indexer error.", indexer.Ptr(), err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ticker.Stop()
	done <- true
}
