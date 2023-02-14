package cmd

import (
	"context"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/bindings/vat"
	"github.com/IR-Digital-Token/auction-keeper/collateral"
	"github.com/IR-Digital-Token/auction-keeper/configs"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/callbacks"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/IR-Digital-Token/auction-keeper/services/processor"
	"github.com/IR-Digital-Token/auction-keeper/services/transaction"
	"github.com/IR-Digital-Token/auction-keeper/services/uniswap_v3"
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

func clipperAllowance(eth *ethclient.Client, collateralName string, vatAddr, clipperAddr common.Address, sender transaction.ISender) error {
	vatInstance, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		return err
	}

	allowance, err := vatInstance.Can(nil, sender.GetAddress(), clipperAddr)
	if err != nil {
		return err
	}

	if allowance.Cmp(big.NewInt(1)) != 0 { // if allowance != 1
		fmt.Printf("HOPING %s CLIPPER IN VAT\n", collateralName)
		txHash, err := sender.Hope(vatInstance, clipperAddr)
		if err != nil {
			return err
		}

		fmt.Printf("Hoped %s Clipper in VAT: %s\n", collateralName, txHash)
	} else {
		fmt.Printf("%s Clipper is Hoped in VAT \n", collateralName)
	}
	return nil
}

func zarJoinAllowance(eth *ethclient.Client, vatAddr, zarJoinAddr common.Address, sender transaction.ISender) error {
	vatInstance, err := vat.NewVat(vatAddr, eth)
	if err != nil {
		return err
	}

	allowance, err := vatInstance.Can(nil, sender.GetAddress(), zarJoinAddr)
	if err != nil {
		return err
	}

	if allowance.Cmp(big.NewInt(1)) != 0 { // if allowance != 1
		fmt.Println("HOPING ZAR_JOIN IN VAT")
		txHash, err := sender.Hope(vatInstance, zarJoinAddr)
		if err != nil {
			return err
		}
		fmt.Println("Hoped ZAR_JOIN in VAT ", txHash)
	} else {
		fmt.Println("ZAR_JOIN is Hoped in VAT ")
	}
	return nil
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

	/***************************************
	  clipper and zarJoin Allowance
	***************************************/
	for _, c := range collaterals {
		err = clipperAllowance(eth, c.Name, cfg.Vat, c.Clipper.Address, sender)
		if err != nil {
			panic(err)
		}
	}
	err = zarJoinAllowance(eth, cfg.Vat, cfg.ZarJoin, sender)
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
