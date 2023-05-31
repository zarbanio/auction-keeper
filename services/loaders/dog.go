package loaders

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/bindings/dog"
	"github.com/zarbanio/auction-keeper/domain/entities"
)

type DogLoader struct {
	dog *dog.Dog
}

func NewDogLoader(eth *ethclient.Client, dogAddr common.Address) *DogLoader {

	d, err := dog.NewDog(dogAddr, eth)
	if err != nil {
		log.Fatal(err)
	}

	return &DogLoader{
		dog: d,
	}
}

func (dl *DogLoader) GetHole(ctx context.Context) (*big.Int, error) {
	hole, err := dl.dog.Hole(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	return hole, nil
}

func (dl *DogLoader) GetDirt(ctx context.Context) (*big.Int, error) {
	dirt, err := dl.dog.Dirt(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	return dirt, nil
}

func (dl *DogLoader) GetIlk(ctx context.Context, ilkId [32]byte) (*entities.DogIlk, error) {
	ilkInfo, err := dl.dog.Ilks(&bind.CallOpts{Context: ctx}, ilkId)
	if err != nil {
		return nil, err
	}

	return &entities.DogIlk{
		Clip: ilkInfo.Clip,
		Chop: ilkInfo.Chop,
		Hole: ilkInfo.Hole,
		Dirt: ilkInfo.Dirt,
	}, nil
}
