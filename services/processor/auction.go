package processor

import (
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/entities"
	"github.com/IR-Digital-Token/auction-keeper/services/loaders"
	"sync"
)

type AuctionProcessor struct {
	clippersLoader map[string]*loaders.ClipperLoader
	auctions       chan entities.Auction
	addAuction     sync.Mutex
}

func NewAuctionProcessor(_clippersLoader map[string]*loaders.ClipperLoader) *AuctionProcessor {
	return &AuctionProcessor{
		clippersLoader: _clippersLoader,
		auctions:       make(chan entities.Auction, 128),
	}
}

func (ap *AuctionProcessor) AddAuction(auction entities.Auction) {
	ap.addAuction.Lock()
	defer ap.addAuction.Unlock()

	ap.auctions <- auction
}

func (ap *AuctionProcessor) StartProcess() {
	go func() {
		for {
			select {
			case auction := <-ap.auctions:
				fmt.Println("Start Process Auction")
				fmt.Println("\tAuction ID: ", auction.Id)
				fmt.Println("\tClipper Name: ", ap.clippersLoader[auction.ClipperName].Name)
				fmt.Println("\tClipper Address: ", ap.clippersLoader[auction.ClipperName].Address)
			}
		}
	}()
}
