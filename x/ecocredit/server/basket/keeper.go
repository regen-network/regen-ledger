package basket

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// Keeper is the basket keeper.
type Keeper struct {
	stateStore        basketv1.StateStore
	bankKeeper        BankKeeper
	ecocreditKeeper   EcocreditKeeper
	moduleAccountName string
	storeKey          sdk.StoreKey
}

var _ baskettypes.MsgServer = Keeper{}
var _ baskettypes.QueryServer = Keeper{}

// NewKeeper returns a new keeper instance.
func NewKeeper(db ormdb.ModuleDB, ecocreditKeeper EcocreditKeeper, bankKeeper BankKeeper, storeKey sdk.StoreKey, moduleAccountName string) Keeper {
	basketStore, err := basketv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	return Keeper{bankKeeper: bankKeeper, ecocreditKeeper: ecocreditKeeper, stateStore: basketStore, storeKey: storeKey, moduleAccountName: moduleAccountName}
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
}

// BankKeeper abstracts over methods that the main bank keeper
// needs to expose to the basket keeper.
//
// NOTE: run `make mocks` whenever you add methods here
type BankKeeper interface {
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}
