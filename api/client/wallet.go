package client

import (
	"context"
	"fmt"

	"github.com/textileio/powergate/wallet/rpc"
)

// Wallet provides an API for managing filecoin wallets.
type Wallet struct {
	client rpc.RPCServiceClient
}

// NewWallet creates a new filecoin address [bls|secp256k1].
func (w *Wallet) NewWallet(ctx context.Context, typ string) (string, error) {
	resp, err := w.client.NewAddress(ctx, &rpc.NewAddressRequest{Type: typ})
	if err != nil {
		return "", fmt.Errorf("calling NewAddress: %v", err)
	}
	return resp.GetAddress(), nil
}

// List returns all wallet addresses.
func (w *Wallet) List(ctx context.Context) ([]string, error) {
	resp, err := w.client.List(ctx, &rpc.ListRequest{})
	if err != nil {
		return nil, fmt.Errorf("calling List: %v", err)
	}
	return resp.Addresses, nil
}

// WalletBalance gets a filecoin wallet's balance.
func (w *Wallet) WalletBalance(ctx context.Context, address string) (uint64, error) {
	resp, err := w.client.WalletBalance(ctx, &rpc.WalletBalanceRequest{Address: address})
	if err != nil {
		return 0, fmt.Errorf("calling WalletBalance: %v", err)
	}
	return resp.GetBalance(), nil
}
