package basket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
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

var _ baskettypes.MsgServer = Keeper{}
var _ baskettypes.QueryServer = Keeper{}

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
