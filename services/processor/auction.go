package processor

import (
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"math/big"
	"sync"
)

type clipper struct {
	clipperLoader   *loaders.ClipperLoader
	auctions        map[*big.Int]entities.Auction
	updateAuctions  sync.Mutex
	processAuctions sync.Mutex
}

type AuctionProcessor struct {
	clippers map[string]*clipper
}

func NewAuctionProcessor(_clippersLoader map[string]*loaders.ClipperLoader) *AuctionProcessor {
	auctionProcessor := &AuctionProcessor{}

	for _, ca := range _clippersLoader {
		auctionProcessor.clippers[ca.Name] = &clipper{clipperLoader: ca}
	}

	return auctionProcessor
}

func (ap *AuctionProcessor) AddAuction(auction entities.Auction, clipperName string) {
	c := ap.clippers[clipperName]
	c.updateAuctions.Lock()
	defer c.updateAuctions.Unlock()

	c.auctions[auction.Id] = auction
}

func (ap *AuctionProcessor) DeleteAuction(auction entities.Auction, clipperName string) {
	c := ap.clippers[clipperName]
	c.updateAuctions.Lock()
	defer c.updateAuctions.Unlock()

	delete(c.auctions, auction.Id)
}

func (ap *AuctionProcessor) StartProcess() {
	for clipperName, clipperAuctions := range ap.clippers {
		clipperAuctions.processAuctions.Lock()
		clipperAuctions.updateAuctions.Lock()

		fmt.Println("processing opportunities for:", clipperName)
		fmt.Println(clipperName, "\tactive auctions qty: ", len(clipperAuctions.auctions))

		for _, auction := range clipperAuctions.auctions {
			fmt.Println("\tprocessing auction id:", auction.Id)

		}

		clipperAuctions.updateAuctions.Unlock()
		clipperAuctions.processAuctions.Unlock()
	}
}
