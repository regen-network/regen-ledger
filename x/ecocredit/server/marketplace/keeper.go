package marketplace

import (
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type Keeper struct {
	stateStore   marketApi.StateStore
	coreStore    ecoApi.StateStore
	bankKeeper   ecocredit.BankKeeper
	paramsKeeper ecocredit.ParamKeeper
}

func NewKeeper(ss marketApi.StateStore, cs ecoApi.StateStore, bk ecocredit.BankKeeper, params ecocredit.ParamKeeper) Keeper {
	return Keeper{
		coreStore:    cs,
		stateStore:   ss,
		bankKeeper:   bk,
		paramsKeeper: params,
	}
}

var _ marketplace.MsgServer = Keeper{}
var _ marketplace.QueryServer = Keeper{}
