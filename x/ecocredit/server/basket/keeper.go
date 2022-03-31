package basket

import (
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
)

// Keeper is the basket keeper.
type Keeper struct {
	stateStore   api.StateStore
	coreStore    ecoApi.StateStore
	bankKeeper   ecocredit.BankKeeper
	distKeeper   ecocredit.DistributionKeeper
	paramsKeeper ecocredit.ParamKeeper
}

var _ baskettypes.MsgServer = Keeper{}
var _ baskettypes.QueryServer = Keeper{}

// NewKeeper returns a new keeper instance.
func NewKeeper(
	db ormdb.ModuleDB,
	bankKeeper ecocredit.BankKeeper,
	distKeeper ecocredit.DistributionKeeper,
	pk ecocredit.ParamKeeper,
) Keeper {
	basketStore, err := api.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	coreStore, err := ecoApi.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	return Keeper{
		bankKeeper:   bankKeeper,
		distKeeper:   distKeeper,
		stateStore:   basketStore,
		coreStore:    coreStore,
		paramsKeeper: pk,
	}
}
