package keeper

import (
	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	marketplaceapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

var (
	_ types.MsgServer   = &Keeper{}
	_ types.QueryServer = &Keeper{}
)

type Keeper struct {
	*types.UnimplementedMsgServer
	*types.UnimplementedQueryServer

	stateStore    api.StateStore
	bankKeeper    ecocredit.BankKeeper
	moduleAddress sdk.AccAddress

	basketStore basketapi.StateStore
	marketStore marketplaceapi.StateStore

	// the address capable of executing ecocredit params messages. Typically, this
	// should be the x/gov module account.
	authority sdk.AccAddress
	ac        address.Codec
}

func NewKeeper(
	ss api.StateStore,
	bk ecocredit.BankKeeper,
	ma sdk.AccAddress,
	basketStore basketapi.StateStore,
	marketStore marketplaceapi.StateStore,
	authority sdk.AccAddress,
	ac address.Codec,
) Keeper {
	return Keeper{
		stateStore:    ss,
		bankKeeper:    bk,
		moduleAddress: ma,
		basketStore:   basketStore,
		authority:     authority,
		marketStore:   marketStore,
		ac:            ac,
	}
}
