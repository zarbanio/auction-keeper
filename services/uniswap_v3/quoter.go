package uniswap_v3

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	uniswap_v3_quoter "github.com/zarbanio/auction-keeper/bindings/uniswapv3/quoter"
	"github.com/zarbanio/auction-keeper/domain"
	"github.com/zarbanio/auction-keeper/domain/entities"
)

type UniswapV3Quoter struct {
	eth           *ethclient.Client
	Quoter        *uniswap_v3_quoter.Quoter
	QuoterAddress common.Address
}

func NewUniswapV3Quoter(eth *ethclient.Client, quoterAddr common.Address) (*UniswapV3Quoter, error) {
	quoter, err := uniswap_v3_quoter.NewQuoter(quoterAddr, eth)
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

	quoterABI, _ := uniswap_v3_quoter.QuoterMetaData.GetAbi()

	for _, r := range path {
		data, err := quoterABI.Pack("quoteExactInputSingle", domain.QuoteExactInputSingleParams{
			TokenIn:           r.TokenA,
			TokenOut:          r.TokenB,
			AmountIn:          quotedAmountOut,
			Fee:               big.NewInt(r.Fee),
			SqrtPriceLimitX96: big.NewInt(0),
		})
		if err != nil {
			log.Println("error in pack quoteExactInputSingle data: ", err)
			return nil, err
		}

		result, err := q.eth.CallContract(
			context.Background(),
			ethereum.CallMsg{
				To:   &q.QuoterAddress,
				Data: data,
			}, nil,
		)
		if err != nil {
			if err.Error() == "execution reverted" {
				var revertData rpc.DataError
				success := errors.As(err, &revertData)
				if !success {
					return nil, fmt.Errorf("failed to parse error message: %w", err)
				}
			}
			log.Println("error in get quoted amount: ", err)
			return nil, err
		}

		outputs, err := quoterABI.Unpack("quoteExactInputSingle", result)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack Quote: %w", err)
		}

		quotedAmountOut = outputs[0].(*big.Int)
	}

	return quotedAmountOut, nil
}
