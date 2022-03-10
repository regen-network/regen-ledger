package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/types"
	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

type Keeper struct {
	stateStore marketplacev1.StateStore
	coreStore  ecocreditv1.StateStore // TODO: do we want a ref to the core's store? or should we pass it a keeper iface?
	bankKeeper ecocredit.BankKeeper
	params     ParamKeeper
}

type ParamKeeper interface {
	GetParamSet(ctx sdk.Context, ps types.ParamSet)
}

func NewKeeper(ss marketplacev1.StateStore, cs ecocreditv1.StateStore, bk ecocredit.BankKeeper, params ParamKeeper) Keeper {
	return Keeper{
		coreStore:  cs,
		stateStore: ss,
		bankKeeper: bk,
		params:     params,
	}
}

// TODO: uncomment when impl
// var _ v1.MsgServer = Keeper{}
