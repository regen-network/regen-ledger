package data

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//go:generate mockgen -source=expected_keepers.go -package mocks -destination mocks/expected_keepers.go

// AccountKeeper defines the expected interface needed to create and retrieve accounts.
type AccountKeeper interface {
	// NewAccount returns a new account with the next account number. Does not save the new account to the store.
	NewAccount(context.Context, sdk.AccountI) sdk.AccountI

	// GetAccount retrieves an account from the store.
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI

	// SetAccount sets an account in the store.
	SetAccount(context.Context, sdk.AccountI)
}

// BankKeeper defines the expected interface needed to send coins and retrieve account balances.
type BankKeeper interface {
	SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetSupply(ctx context.Context, denom string) sdk.Coin
}
