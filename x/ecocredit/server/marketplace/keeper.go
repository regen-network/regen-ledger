package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type Keeper struct {
	stateStore   marketApi.StateStore
	coreStore    ecoApi.StateStore
	bankKeeper   ecocredit.BankKeeper
	paramsKeeper ecocredit.ParamKeeper
	authority    sdk.AccAddress
}

func NewKeeper(ss marketApi.StateStore, cs ecoApi.StateStore, bk ecocredit.BankKeeper,
	params ecocredit.ParamKeeper, authority sdk.AccAddress) Keeper {
	return Keeper{
		coreStore:    cs,
		stateStore:   ss,
		bankKeeper:   bk,
		paramsKeeper: params,
		authority:    authority,
	}
}

var _ marketplace.MsgServer = Keeper{}
var _ marketplace.QueryServer = Keeper{}
