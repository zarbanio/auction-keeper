package processor

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/zarban/dog"
	entities "github.com/zarbanio/auction-keeper/domain/entities/inputMethods"
	"github.com/zarbanio/auction-keeper/services/transaction"
)

type dogBarkService struct {
	dog    *dog.Dog
	sender *transaction.Sender
}

func NewDogService(eth *ethclient.Client, dogAddr common.Address, sender *transaction.Sender) *dogBarkService {
	d, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &dogBarkService{
		dog:    d,
		sender: sender,
	}
}

func (s *dogBarkService) Start(bark *entities.DogBark) error {
	log.Println("Dog Bark Service is starting")
	go s.Run(bark) // Run the service logic in a goroutine
	return nil
}

// Stop the service
func (s *dogBarkService) Stop() {
	panic("Implement me!")
}

// Run the service logic
func (s *dogBarkService) Run(bark *entities.DogBark) (*types.Transaction, error) {
	opts, err := s.sender.GetOpts()
	if err != nil {
		return nil, err
	}

	tx, err := s.dog.Bark(opts, bark.Ilk, bark.Urn, s.sender.GetAddress())
	if err != nil {
		log.Fatal(err)
	}

	// Print the transaction hash
	log.Println(tx.Hash().Hex())
	return tx, nil
}
