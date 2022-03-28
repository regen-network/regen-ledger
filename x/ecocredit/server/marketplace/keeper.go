package marketplace

import (
	marketplaceapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type Keeper struct {
	stateStore    marketplaceapi.StateStore
	coreStore     ecocreditapi.StateStore
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper
	paramsKeeper  ecocredit.ParamKeeper
}

func NewKeeper(db ormdb.ModuleDB, cs ecocreditapi.StateStore, bk ecocredit.BankKeeper, params ecocredit.ParamKeeper, ak ecocredit.AccountKeeper) Keeper {
	marketplaceStore, err := marketplaceapi.NewStateStore(db)
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
