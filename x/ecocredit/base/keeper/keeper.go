package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketplaceapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

var (
	_ types.MsgServer = &Keeper{}
	_ types.QueryServer = &Keeper{}
)

type Keeper struct {
	stateStore    api.StateStore
	bankKeeper    ecocredit.BankKeeper
	moduleAddress sdk.AccAddress

	basketStore basketapi.StateStore
	marketStore marketplaceapi.StateStore

	// the address capable of executing ecocredit params messages. Typically, this
	// should be the x/gov module account.
	authority sdk.AccAddress
}

func NewKeeper(
	ss api.StateStore,
	bk ecocredit.BankKeeper,
	ma sdk.AccAddress,
	basketStore basketapi.StateStore,
	marketStore marketplaceapi.StateStore,
	authority sdk.AccAddress,
) Keeper {
	return Keeper{
		stateStore:    ss,
		bankKeeper:    bk,
		moduleAddress: ma,
		basketStore:   basketStore,
		authority:     authority,
		marketStore:   marketStore,
	}
}
