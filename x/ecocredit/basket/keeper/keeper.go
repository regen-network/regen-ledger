package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

var (
	_ types.MsgServer   = Keeper{}
	_ types.QueryServer = Keeper{}
)

// Keeper is the basket keeper.
type Keeper struct {
	stateStore    api.StateStore
	coreStore     ecoApi.StateStore
	bankKeeper    ecocredit.BankKeeper
	paramsKeeper  ecocredit.ParamKeeper
	moduleAddress sdk.AccAddress
	authority     sdk.AccAddress
}

// NewKeeper returns a new keeper instance.
func NewKeeper(
	ss api.StateStore,
	cs ecoApi.StateStore,
	bk ecocredit.BankKeeper,
	pk ecocredit.ParamKeeper,
	ma sdk.AccAddress,
	authority sdk.AccAddress,
) Keeper {
	return Keeper{
		stateStore:    ss,
		coreStore:     cs,
		bankKeeper:    bk,
		paramsKeeper:  pk,
		moduleAddress: ma,
		authority:     authority,
	}
}
