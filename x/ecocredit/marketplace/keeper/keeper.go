package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	marketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

var (
	_ types.MsgServer   = Keeper{}
	_ types.QueryServer = Keeper{}
)

type Keeper struct {
	stateStore marketapi.StateStore
	baseStore  baseapi.StateStore
	bankKeeper ecocredit.BankKeeper
	authority  sdk.AccAddress
}

func NewKeeper(ss marketapi.StateStore, cs baseapi.StateStore, bk ecocredit.BankKeeper,
	authority sdk.AccAddress) Keeper {
	return Keeper{
		baseStore:  cs,
		stateStore: ss,
		bankKeeper: bk,
		authority:  authority,
	}
}
