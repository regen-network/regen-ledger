package marketplace

import (
	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

type Keeper struct {
	stateStore marketplacev1.StateStore
	coreStore  ecocreditv1.StateStore
	bankKeeper ecocredit.BankKeeper
	params     server.ParamKeeper
}

func NewKeeper(ss marketplacev1.StateStore, cs ecocreditv1.StateStore, bk ecocredit.BankKeeper, params server.ParamKeeper) Keeper {
	return Keeper{
		coreStore:  cs,
		stateStore: ss,
		bankKeeper: bk,
		params:     params,
	}
}

// TODO: uncomment when impl
// var _ v1.MsgServer = Keeper{}
// var _ v1.QueryServer = Keeper{}
