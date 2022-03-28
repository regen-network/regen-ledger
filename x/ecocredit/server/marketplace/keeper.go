package marketplace

import (
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
)

type Keeper struct {
	stateStore    marketApi.StateStore
	coreStore     ecoApi.StateStore
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper
	paramsKeeper  ecocredit.ParamKeeper
}

func NewKeeper(db ormdb.ModuleDB, cs ecoApi.StateStore, bk ecocredit.BankKeeper, params ecocredit.ParamKeeper, ak ecocredit.AccountKeeper) Keeper {
	marketplaceStore, err := marketApi.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	return Keeper{
		coreStore:     cs,
		stateStore:    marketplaceStore,
		bankKeeper:    bk,
		paramsKeeper:  params,
		accountKeeper: ak,
	}
}

var _ marketplace.MsgServer = Keeper{}
var _ marketplace.QueryServer = Keeper{}
