package core

import (
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

var _ core.MsgServer = &Keeper{}
var _ core.QueryServer = &Keeper{}

type Keeper struct {
	stateStore    api.StateStore
	bankKeeper    ecocredit.BankKeeper
	paramsKeeper  ecocredit.ParamKeeper
	accountKeeper ecocredit.AccountKeeper
}

func NewKeeper(ss api.StateStore, bk ecocredit.BankKeeper, pk ecocredit.ParamKeeper, ak ecocredit.AccountKeeper) Keeper {
	return Keeper{
		stateStore:    ss,
		bankKeeper:    bk,
		paramsKeeper:  pk,
		accountKeeper: ak,
	}
}
