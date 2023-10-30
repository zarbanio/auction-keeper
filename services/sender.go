package sender

import (
	"context"
	"crypto"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Sender struct {
	eth           *ethclient.Client
	blockInterval time.Duration
	privateKey    crypto.PrivateKey
}

// CallBacks
type BroadcastCb func(ctx context.Context, tx *types.Transaction) (uint64, error)
type ReceiptCb func(ctx context.Context, receipt *types.Receipt)

func NewSender(
	eth *ethclient.Client,
	privateKey crypto.PrivateKey,
	blockInterval time.Duration,
) *Sender {
	return &Sender{
		eth,
		blockInterval,
		privateKey,
	}
}

func (s *Sender) SendTransaction(ctx context.Context, tx *types.Transaction, onBroadcast BroadcastCb, onReceipt ReceiptCb) error {

	// Call the OnBroadcast CB
	if onBroadcast != nil {
		onBroadcast(ctx, tx)
	}

	// Wait for the transaction receipt
	receipt, err := s.waitForReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}

	// Call the onReceipt CB
	if onReceipt != nil {
		onReceipt(ctx, receipt)
	}

	return nil
}

func (s *Sender) waitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ticker := time.NewTicker(s.blockInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			receipt, err := s.eth.TransactionReceipt(ctx, txHash)
			if err == nil {
				return receipt, nil
			}
		}
	}
}
