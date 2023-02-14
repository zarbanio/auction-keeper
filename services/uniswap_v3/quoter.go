package uniswap_v3

import (
	"context"
	"fmt"
	"github.com/IR-Digital-Token/auction-keeper/bindings/uniswap_v3_quoter"
	"github.com/IR-Digital-Token/auction-keeper/domain/entities"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type UniswapV3Quoter struct {
	eth           *ethclient.Client
	Quoter        *uniswap_v3_quoter.UniswapV3Quoter
	QuoterAddress common.Address
}

func NewUniswapV3Quoter(eth *ethclient.Client, quoterAddr common.Address) (*UniswapV3Quoter, error) {
	quoter, err := uniswap_v3_quoter.NewUniswapV3Quoter(quoterAddr, eth)
	if err != nil {
		return nil, err
	}

	return &UniswapV3Quoter{
		eth:           eth,
		Quoter:        quoter,
		QuoterAddress: quoterAddr,
	}, nil
}

func (q *UniswapV3Quoter) GetQuotedAmountOut(amount *big.Int, path []entities.UniswapV3Route, assetDecimal *big.Int) (*big.Int, error) {
	decNormalized := math.BigPow(10, new(big.Int).Sub(big.NewInt(18), assetDecimal).Int64())
	quotedAmountOut := new(big.Int).Div(amount, decNormalized)

	quoterABI, _ := uniswap_v3_quoter.UniswapV3QuoterMetaData.GetAbi()

	for _, r := range path {
		data, err := quoterABI.Pack("quoteExactInputSingle", r.TokenA, r.TokenB, big.NewInt(r.Fee), quotedAmountOut, big.NewInt(0))
		if err != nil {
			fmt.Println("error in pack quoteExactInputSingle data: ", err)
			return nil, err
		}

		output, err := q.eth.CallContract(context.Background(), ethereum.CallMsg{
			To:   &q.QuoterAddress,
			Data: data,
		}, nil)
		if err != nil {
			fmt.Println("error in get quoted amount: ", err)
			return nil, err
		}

		quotedAmountOut = new(big.Int).SetBytes(output)
	}

	return quotedAmountOut, nil
}
