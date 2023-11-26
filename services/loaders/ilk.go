package loaders

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/bindings/zarban/dog"
	"github.com/zarbanio/auction-keeper/bindings/zarban/gemjoin"
	"github.com/zarbanio/auction-keeper/bindings/zarban/ilkregistry"
	"github.com/zarbanio/auction-keeper/bindings/zarban/jug"
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
	dog         *dog.Dog
	joins       map[common.Address]join           // join address to join contract
	medians     map[common.Address]common.Address // gem address to median address
}

type join struct {
	bind    *gemjoin.Gemjoin
	address common.Address
	ilk     [32]byte
	gem     common.Address
}

func NewIlksLoader(
	ceth eth.Ethereum,
	store store.IStore,
	vatAddr,
	jugAddr,
	spotAddr,
	dogAddr,
	ilkregistryAddr common.Address,
	joinsAddr []common.Address,
	medians map[common.Address]common.Address) *IlksLoader {

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

	joins := make(map[common.Address]join)
	for _, addr := range joinsAddr {
		gj, err := gemjoin.NewGemjoin(addr, ceth)
		if err != nil {
			log.Fatal(err)
		}
		i, err := gj.Ilk(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			log.Fatal(err)
		}
		gem, err := gj.Gem(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			log.Fatal(err)
		}

		j := join{
			bind:    gj,
			address: addr,
			ilk:     i,
			gem:     gem,
		}

		joins[addr] = j
	}

	return &IlksLoader{
		store:       store,
		jug:         j,
		spot:        s,
		dog:         d,
		ilkregistry: ir,
		vat:         v,
		joins:       joins,
		medians:     medians,
	}
}

func (l *IlksLoader) fetchIlkByJoin(ctx context.Context, j join) (*domain.Ilk, error) {
	vatInfo, err := l.vat.Ilks(&bind.CallOpts{Context: ctx}, j.ilk)
	if err != nil {
		return nil, fmt.Errorf("error getting vat info. %w", err)
	}
	jugInfo, err := l.jug.Ilks(&bind.CallOpts{Context: ctx}, j.ilk)
	if err != nil {
		return nil, fmt.Errorf("error getting jug info. %w", err)
	}
	spotInfo, err := l.spot.Ilks(&bind.CallOpts{Context: ctx}, j.ilk)
	if err != nil {
		return nil, fmt.Errorf("error getting spot info. %w", err)
	}
	par, err := l.spot.Par(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("error getting par. %w", err)
	}
	clipAddr, err := l.ilkregistry.Xlip(&bind.CallOpts{Context: ctx}, j.ilk)
	if err != nil {
		return nil, fmt.Errorf("error getting clipper. %w", err)
	}
	pip, err := l.ilkregistry.Pip(&bind.CallOpts{Context: ctx}, j.ilk)
	if err != nil {
		return nil, fmt.Errorf("error getting pip. %w", err)
	}
	dogInfo, err := l.dog.Ilks(&bind.CallOpts{Context: ctx}, j.ilk)
	if err != nil {
		return nil, fmt.Errorf("error getting dog info. %w", err)
	}

	medain, ok := l.medians[j.gem]
	if !ok {
		return nil, errors.New("median not found")
	}

	medainLog, err := l.store.GetLastLogMedianPrice(ctx, medain)
	if err != nil {
		if !errors.Is(err, store.ErrMedianLogPriceMedianNotFound) {
			return nil, fmt.Errorf("error getting median log. %w", err)
		}
	}

	price := price(par, vatInfo.Spot, spotInfo.Mat)
	if medainLog != nil {
		price = math.Normalize(medainLog.Val, int64(math.WadDecimals))
	}

	name := domain.Bytes32ToString(j.ilk)
	ilk := &domain.Ilk{
		Bytes32:                       j.ilk,
		Name:                          name,
		Symbol:                        IlkToSymbol(name),
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
		Join:                          j.address,
		Gem:                           j.gem,
		Median:                        medain,
		Clipper:                       clipAddr,
		Pip:                           pip,
	}

	return ilk, nil
}

func (l *IlksLoader) LoadIlks(ctx context.Context) ([]domain.Ilk, error) {
	var ilks []domain.Ilk
	for _, join := range l.joins {
		ilk, err := l.fetchIlkByJoin(ctx, join)
		if err != nil {
			return nil, fmt.Errorf("error getting ilk. %w", err)
		}
		ilks = append(ilks, *ilk)
	}
	return ilks, nil
}
