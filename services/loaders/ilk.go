package loaders

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/bindings/zarban/dog"
	"github.com/zarbanio/auction-keeper/bindings/zarban/ilkregistry"
	"github.com/zarbanio/auction-keeper/bindings/zarban/jug"
	"github.com/zarbanio/auction-keeper/bindings/zarban/osmregistry"
	"github.com/zarbanio/auction-keeper/bindings/zarban/spot"
	"github.com/zarbanio/auction-keeper/bindings/zarban/vat"
	"github.com/zarbanio/auction-keeper/domain"
	"github.com/zarbanio/auction-keeper/domain/math"
	"github.com/zarbanio/auction-keeper/store"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type IlksLoader struct {
	store       store.IStore
	vat         *vat.Vat
	jug         *jug.Jug
	spot        *spot.Spot
	ilkregistry *ilkregistry.Ilkregistry
	osmregistry *osmregistry.Osmregistry
	dog         *dog.Dog
}

func NewIlksLoader(
	ceth eth.Ethereum,
	store store.IStore,
	vatAddr,
	jugAddr,
	spotAddr,
	dogAddr,
	ilkregistryAddr common.Address,
	osmregistryAddr common.Address,
) *IlksLoader {

	v, err := vat.NewVat(vatAddr, ceth)
	if err != nil {
		log.Fatal(err)
	}
	j, err := jug.NewJug(jugAddr, ceth)
	if err != nil {
		log.Fatal(err)
	}
	s, err := spot.NewSpot(spotAddr, ceth)
	if err != nil {
		log.Fatal(err)
	}
	ir, err := ilkregistry.NewIlkregistry(ilkregistryAddr, ceth)
	if err != nil {
		log.Fatal(err)
	}
	d, err := dog.NewDog(dogAddr, ceth)
	if err != nil {
		log.Fatal(err)
	}

	or, err := osmregistry.NewOsmregistry(osmregistryAddr, ceth)
	if err != nil {
		log.Fatal("error creating osm registry.", err)
	}

	return &IlksLoader{
		store:       store,
		jug:         j,
		spot:        s,
		dog:         d,
		ilkregistry: ir,
		osmregistry: or,
		vat:         v,
	}
}

func (l *IlksLoader) LoadIlkByName(ctx context.Context, name string) (*domain.Ilk, error) {
	id := domain.StringToBytes32(name)
	return l.LoadIlkById(ctx, id)
}

func (l *IlksLoader) LoadIlkById(ctx context.Context, id [32]byte) (*domain.Ilk, error) {
	vatInfo, err := l.vat.Ilks(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, fmt.Errorf("error getting vat info. %w", err)
	}
	jugInfo, err := l.jug.Ilks(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, fmt.Errorf("error getting jug info. %w", err)
	}
	spotInfo, err := l.spot.Ilks(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, fmt.Errorf("error getting spot info. %w", err)
	}
	dogInfo, err := l.dog.Ilks(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, fmt.Errorf("error getting dog info. %w", err)
	}
	par, err := l.spot.Par(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting par. %w", err)
	}
	data, err := l.ilkregistry.IlkData(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, fmt.Errorf("error getting ilk data. %w", err)
	}
	if data.Gem == common.HexToAddress("0x") {
		return nil, errors.New("ilk is not supported")
	}

	_, median, err := l.osmregistry.Get(&bind.CallOpts{Context: ctx}, data.Gem)
	if err != nil {
		return nil, fmt.Errorf("error getting osm src. %w", err)
	}

	price := price(par, vatInfo.Spot, spotInfo.Mat)

	name := domain.Bytes32ToString(id)
	ilk := domain.Ilk{
		Bytes32:                       id,
		Name:                          name,
		Symbol:                        domain.Symbol(data.Symbol),
		MinimumCollateralizationRatio: liquidationRatio(spotInfo.Mat),
		AnnualStabilityFee:            annualStabilityFee(jugInfo.Duty),
		DustLimit:                     math.Normalize(vatInfo.Dust, int64(math.RadDecimals)),
		LiquidationPenalty:            math.Normalize(dogInfo.Chop, int64(math.WadDecimals)),
		Hole:                          math.Normalize(dogInfo.Hole, int64(math.RadDecimals)),
		Dirt:                          math.Normalize(dogInfo.Dirt, int64(math.RadDecimals)),
		DebtCeiling:                   debtCeiling(vatInfo.Line),
		Price:                         price,
		Debt:                          debtValue(vatInfo.Art, vatInfo.Rate),
		Rate:                          vatInfo.Rate,
		Join:                          data.Join,
		Gem:                           data.Gem,
		Median:                        median,
		Clipper:                       data.Xlip,
		Pip:                           data.Pip,
	}

	return &ilk, nil
}

func (l *IlksLoader) LoadIlks(ctx context.Context) ([]domain.Ilk, error) {
	ids, err := l.ilkregistry.List(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting ilk count. %w", err)
	}
	var ilks []domain.Ilk
	for _, id := range ids {
		ilk, err := l.LoadIlkById(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("error getting ilk. %w", err)
		}
		ilks = append(ilks, *ilk)
	}
	return ilks, nil
}
