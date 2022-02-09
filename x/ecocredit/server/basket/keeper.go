package basket

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Keeper is the basket keeper.
type Keeper struct {
	basketStore     basketv1.StateStore
	bankKeeper      BankKeeper
	ecocreditKeeper EcocreditKeeper
}

// NewKeeper returns a new keeper instance.
func NewKeeper(db ormdb.ModuleDB, ecocreditKeeper EcocreditKeeper, bankKeeper BankKeeper) Keeper {
	basketStore, err := basketv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	return Keeper{bankKeeper: bankKeeper, ecocreditKeeper: ecocreditKeeper, basketStore: basketStore}
}

// EcocreditKeeper abstracts over methods that the main eco-credit keeper
// needs to expose to the basket keeper.
//
// NOTE: run `make mocks` whenever you add methods here
type EcocreditKeeper interface {
	// we embed a query server directly here rather than trying to go through
	// ADR 033 for simplicity
	ecocredit.QueryServer

	GetCreateBasketFee(ctx context.Context) sdk.Coins
	// add additional keeper methods here
}

type BankKeeper interface {
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}
