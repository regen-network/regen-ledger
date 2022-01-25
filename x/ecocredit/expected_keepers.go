package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

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
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// Keeper defines the expected interface needed to prune expired buy and sell orders.
type Keeper interface {
	PruneOrders(ctx sdk.Context) error
}
