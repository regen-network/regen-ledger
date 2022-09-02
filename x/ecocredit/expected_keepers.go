package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/params/types"
)

//go:generate mockgen -source=expected_keepers.go -package mocks -destination mocks/expected_keepers.go

// AccountKeeper defines the expected interface needed to create and retrieve accounts.
type AccountKeeper interface {
	// NewAccount returns a new account with the next account number. Does not save the new account to the store.
	NewAccount(sdk.Context, authtypes.AccountI) authtypes.AccountI

	// GetAccount retrieves an account from the store.
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI

	// GetModuleAddress retrieves a module account address from the store.
	GetModuleAddress(moduleName string) sdk.AccAddress

	// SetAccount sets an account in the store.
	SetAccount(sdk.Context, authtypes.AccountI)
}

// BankKeeper defines the expected interface needed to burn and send coins and to retrieve account balances.
type BankKeeper interface {
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

// GovKeeper defines the expected interface needed to query governance params
type GovKeeper interface {
	// GetDepositParams queries governance module deposit params
	GetDepositParams(ctx sdk.Context) govv1.DepositParams
}

type ParamKeeper interface {

	// Get fetches a parameter by key from the Subspace's KVStore and sets the provided pointer to the fetched value.
	// If the value does not exist, this method will panic.
	Get(ctx sdk.Context, key []byte, ptr interface{})

	// GetParamSet fetches each parameter in the ParamSet.
	GetParamSet(ctx sdk.Context, ps types.ParamSet)
}
