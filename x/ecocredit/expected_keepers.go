package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BankKeeper interface {
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
