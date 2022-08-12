package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

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
	moduleAddress sdk.AccAddress

	// the address capable of executing ecocredit params messages. Typically, this
	// should be the x/gov module account.
	authority sdk.AccAddress
}

func NewKeeper(
	ss api.StateStore,
	bk ecocredit.BankKeeper,
	pk ecocredit.ParamKeeper,
	ma sdk.AccAddress,
	authority sdk.AccAddress,
) Keeper {
	return Keeper{
		stateStore:    ss,
		bankKeeper:    bk,
		paramsKeeper:  pk,
		moduleAddress: ma,
		authority:     authority,
	}
}
