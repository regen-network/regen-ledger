package basket

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// Keeper is the basket keeper.
type Keeper struct {
	stateStore      basketv1.StateStore
	bankKeeper      ecocredit.BankKeeper
	ecocreditKeeper EcocreditKeeper
	storeKey        sdk.StoreKey
	distKeeper      ecocredit.DistributionKeeper
}

var _ baskettypes.MsgServer = Keeper{}
var _ baskettypes.QueryServer = Keeper{}

// NewKeeper returns a new keeper instance.
func NewKeeper(
	db ormdb.ModuleDB,
	ecocreditKeeper EcocreditKeeper,
	bankKeeper ecocredit.BankKeeper,
	distKeeper ecocredit.DistributionKeeper,
	storeKey sdk.StoreKey,
) Keeper {
	basketStore, err := basketv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	return Keeper{
		bankKeeper:      bankKeeper,
		ecocreditKeeper: ecocreditKeeper,
		distKeeper:      distKeeper,
		stateStore:      basketStore,
		storeKey:        storeKey,
	}
}

// EcocreditKeeper abstracts over methods that the main ecocredit keeper
// needs to expose to the basket keeper.
//
// NOTE: run `make mocks` whenever you add methods here
type EcocreditKeeper interface {
	// we embed a query server directly here rather than trying to go through
	// ADR 033 for simplicity
	ecocredit.QueryServer

	GetCreateBasketFee(ctx context.Context) sdk.Coins
	HasClassInfo(ctx types.Context, classID string) bool
}
