package marketplace

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplaceapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type Keeper struct {
	stateStore marketplaceapi.StateStore
	coreStore  ecocreditapi.StateStore
	bankKeeper ecocredit.BankKeeper
	params     ecocredit.ParamKeeper
}

func NewKeeper(db ormdb.ModuleDB, cs ecocreditapi.StateStore, bk ecocredit.BankKeeper, params ecocredit.ParamKeeper) Keeper {
	marketplaceStore, err := marketplaceapi.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	return Keeper{
		coreStore:  cs,
		stateStore: marketplaceStore,
		bankKeeper: bk,
		params:     params,
	}
}

var _ marketplace.MsgServer = Keeper{}
var _ marketplace.QueryServer = Keeper{}
