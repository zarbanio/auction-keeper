package sender

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/services/signer"
	"github.com/zarbanio/auction-keeper/store"
)

type Sender interface {
	GetAddress() common.Address
	GetTransactOpts() (*bind.TransactOpts, error)
	HandleSentTx(tx *types.Transaction) error
}

type sender struct {
	signer  signer.Signer
	eth     *ethclient.Client // TODO: create a custom interface with only pendingStateReader and GasPricer
	txs     store.ITranasaction
	timeout time.Duration
}

func NewSender(signer signer.Signer, eth *ethclient.Client) Sender {
	return &sender{
		signer:  signer,
		eth:     eth,
		timeout: 1 * time.Minute, // TODO: make this configurable
	}
}

func (s sender) GetAddress() common.Address {
	return s.signer.GetAddress()
}

func (s sender) GetTransactOpts() (*bind.TransactOpts, error) {
	opts, err := s.signer.GetTransactOpts()
	if err != nil {
		return nil, err
	}
	nonce, err := s.eth.PendingNonceAt(context.Background(), s.signer.GetAddress())
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.eth.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = gasPrice
	return opts, nil
}

func (s sender) HandleSentTx(tx *types.Transaction) error {
	ctx := context.Background()

	err, id := s.txs.CreateTransaction(ctx, tx)
	if err != nil {
		log.Println("error in saving the transaction.", err)
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	receipt, header, err := WaitForReceipt(ctx, s.eth, tx.Hash())
	if err != nil {
		return err
	}

	err = s.txs.UpdateTransactionReceipt(context.Background(), id, receipt, header)
	if err != nil {
		log.Println("error in updating the transaction receipt.", err)
		return err
	}

	return nil
}

func WaitForReceipt(ctx context.Context, eth *ethclient.Client, txHash common.Hash) (*types.Receipt, *types.Header, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-ticker.C:
			receipt, err := eth.TransactionReceipt(ctx, txHash)
			if err == nil {
				header, err := eth.HeaderByHash(ctx, receipt.BlockHash)
				if err != nil {
					return nil, nil, err
				}
				return receipt, header, nil
			}
		}
	}
}
