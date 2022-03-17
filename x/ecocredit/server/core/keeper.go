package core

import (
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

var _ core.MsgServer = &Keeper{}
var _ core.QueryServer = &Keeper{}

type Keeper struct {
	stateStore api.StateStore
	bankKeeper ecocredit.BankKeeper
	params     server.ParamKeeper
}

func NewKeeper(ss api.StateStore, bk ecocredit.BankKeeper, params server.ParamKeeper) Keeper {
	return Keeper{
		stateStore: ss,
		bankKeeper: bk,
		params:     params,
	}
}
