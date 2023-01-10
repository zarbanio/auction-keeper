package cmd

import (
	"context"
	"github.com/IR-Digital-Token/auction-keeper/bindings/clip"
	"github.com/IR-Digital-Token/auction-keeper/configs"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/callbacks"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"github.com/IR-Digital-Token/auction-keeper/services/processor"
	"github.com/IR-Digital-Token/x/chain"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func registerEventHandlers(indexer *chain.Indexer, eth *ethclient.Client, auctionProcessor *processor.AuctionProcessor, clippers map[string]*loaders.ClipperLoader) {
	println("Register callbacks on event triggers come from the indexer.")

	for _, v := range clippers {
		indexer.RegisterEventHandler(clipper.NewKickHandler(v.Address, eth, callbacks.ClipperKickCallback(auctionProcessor, v.Name)))
		indexer.RegisterEventHandler(clipper.NewRedoHandler(v.Address, eth, callbacks.ClipperRedoCallback()))
		indexer.RegisterEventHandler(clipper.NewTakeHandler(v.Address, eth, callbacks.ClipperTakeCallback()))
	}

	println("Done\n")
}

func getActiveAuctions(clippers map[string]*loaders.ClipperLoader) (map[decimal.Decimal]entities.Auction, error) {
	println("Get active auctions from clippers.")

	auctions := make(map[decimal.Decimal]entities.Auction)
	for _, v := range clippers {
		auctionsIds, err := v.GetActiveAuctions()
		if err != nil {
			return nil, err
		}

		for _, auctionId := range auctionsIds {
			sale, err := v.GetSale(auctionId)
			if err != nil {
				return nil, err
			}

			auctions[decimal.NewFromBigInt(auctionId, 0)] = *sale
		}
	}

	defer println("Done\n")
	return auctions, nil
}

func Execute() {
	cfg := configs.ReadConfig("config.yaml")

	eth, err := ethclient.Dial(cfg.Network.Node.Api)
	if err != nil {
		panic(err)
	}

	/***************************************
	 create clippers loaders
	***************************************/
	clippersLoader := make(map[string]*loaders.ClipperLoader)
	clippersLoader["ETHAClipper"], err = loaders.NewClipperLoader(eth, "ETHAClipper", cfg.Contracts.ETHAClipper)
	if err != nil {
		panic(err)
	}
	clippersLoader["ETHBClipper"], err = loaders.NewClipperLoader(eth, "ETHBClipper", cfg.Contracts.ETHBClipper)
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

	/* -------------------------------------------------------------------------- */
	/*  get active auctions and add them in auction processor and start processor */
	activeAuctions, _err := getActiveAuctions(clippersLoader)
	if _err != nil {
		panic(_err)
	}

	auctionProcessor := processor.NewAuctionProcessor(clippersLoader)
	for _, auction := range activeAuctions {
		auctionProcessor.AddAuction(auction)
	}

	auctionProcessor.StartProcess()
	/* -------------------------------------------------------------------------- */

	/* -------------------------------------------------------------------------- */
	/*                     register contract events on indexer                    */
	indexer := chain.NewIndexer(eth, blockPtr, cfg.Indexer.PoolSize)
	registerEventHandlers(indexer, eth, auctionProcessor, clippersLoader)

	for _, v := range clippersLoader {
		indexer.RegisterAddress(v.Address)
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
}
