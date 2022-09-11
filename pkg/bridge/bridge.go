package bridge

import (
	"context"
)

type Bridge struct{}

func New() (*Bridge, error) {
	return &Bridge{}, nil
}

const dummyTxHash = "0x0000000000000000000000000000000000000001"

func (b *Bridge) ShibcMove(ctx context.Context, txHash string) (string, error) {
	return dummyTxHash, nil
}

func (b *Bridge) ETHMove(ctx context.Context, txHash string) (string, error) {
	return dummyTxHash, nil
}
