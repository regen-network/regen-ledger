package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

type Keeper struct {
	stateStore ecocreditv1beta1.StateStore
	bankKeeper ecocredit.BankKeeper
	params     ParamKeeper
}

type ParamKeeper interface {
	GetParamSet(ctx sdk.Context, ps types.ParamSet)
}

func NewKeeper(ss ecocreditv1beta1.StateStore, bk ecocredit.BankKeeper, params ParamKeeper) Keeper {
	return Keeper{
		stateStore: ss,
		bankKeeper: bk,
		params:     params,
	}
}

// TODO: uncomment when impl
//var _ v1beta1.MsgServer = &Keeper{}
