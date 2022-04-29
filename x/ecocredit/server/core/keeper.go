package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

var _ core.MsgServer = &Keeper{}
var _ core.QueryServer = &Keeper{}

type Keeper struct {
	stateStore   api.StateStore
	bankKeeper   ecocredit.BankKeeper
	paramsKeeper ecocredit.ParamKeeper
}

func (k Keeper) AddAskDenom(ctx sdk.Context, proposal marketplace.AskDenomProposal) error {
	//TODO implement me
	panic("implement me")
}

func NewKeeper(ss api.StateStore, bk ecocredit.BankKeeper, pk ecocredit.ParamKeeper) Keeper {
	return Keeper{
		stateStore:   ss,
		bankKeeper:   bk,
		paramsKeeper: pk,
	}
}
