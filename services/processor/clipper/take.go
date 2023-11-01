package processor

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/clipper"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/services/sender"
)

type ClipperTakeService struct {
	eth     *ethclient.Client
	clipper *clipper.Clipper
	sender  sender.Sender
}

func NewClipperTakeService(
	eth *ethclient.Client,
	clipperAddr common.Address,
	sender sender.Sender,
) *ClipperTakeService {

	c, err := clipper.NewClipper(clipperAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &ClipperTakeService{
		eth:     eth,
		clipper: c,
		sender:  sender,
	}
}

func (s *ClipperTakeService) Start() error {
	log.Println("Clipper Service is starting")
	return nil
	// s.Run()
}

// Run the service logic
func (s *ClipperTakeService) Run(take *entities.ClipperTake) error {
	opts, err := s.sender.GetTransactOpts()
	if err != nil {
		return err
	}

	tx, err := s.clipper.Take(opts, take.Auction_id, take.Amt, take.Max, take.Who, take.Data)
	if err != nil {
		return err
	}
	return s.sender.HandleSentTx(tx)
}

// Stop the service
func (s *ClipperTakeService) Stop() {
	panic("Implement me!")
}
