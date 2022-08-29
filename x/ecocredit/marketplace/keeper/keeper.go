package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	marketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

type Keeper struct {
	stateStore   marketapi.StateStore
	coreStore    ecoApi.StateStore
	bankKeeper   ecocredit.BankKeeper
	paramsKeeper ecocredit.ParamKeeper
	authority    sdk.AccAddress
}

func NewKeeper(ss marketapi.StateStore, cs ecoApi.StateStore, bk ecocredit.BankKeeper,
	params ecocredit.ParamKeeper, authority sdk.AccAddress) Keeper {
	return Keeper{
		coreStore:    cs,
		stateStore:   ss,
		bankKeeper:   bk,
		paramsKeeper: params,
		authority:    authority,
	}
}

var _ types.MsgServer = Keeper{}
var _ types.QueryServer = Keeper{}
