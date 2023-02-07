package processor

import (
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"math/big"
	"sync"
)

type AuctionCollection struct {
	auctions       map[*big.Int]entities.Auction
	updateAuctions sync.Mutex
}

func NewAuctionCollection() *AuctionCollection {
	return &AuctionCollection{
		auctions:       make(map[*big.Int]entities.Auction),
		updateAuctions: sync.Mutex{},
	}
}

func (ai *AuctionCollection) AuctionIsExist(Id *big.Int) bool {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	if _, ok := ai.auctions[Id]; ok {
		return true
	}
	return false
}

func (ai *AuctionCollection) addAuction(auction entities.Auction) {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	ai.auctions[auction.Id] = auction
}

func (ai *AuctionCollection) deleteAuction(Id *big.Int) {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	delete(ai.auctions, Id)
}

func (ai *AuctionCollection) updateAuctionAfterTake(id *big.Int, tab *big.Int, lot *big.Int) {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	auction := ai.auctions[id]
	auction.Tab = tab
	auction.Lot = lot
}

func (ai *AuctionCollection) updateAuctionAfterRedo(id *big.Int, top *big.Int, tic uint64) {
	ai.updateAuctions.Lock()
	defer ai.updateAuctions.Unlock()

	auction := ai.auctions[id]
	auction.Top = top
	auction.Tic = tic
}
