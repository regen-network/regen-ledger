package v4

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
)

// MigrateState performs in-place store migrations from v4.0 to v5.0.
func MigrateState(sdkCtx sdk.Context, storeKey storetypes.StoreKey, cdc codec.Codec, ss api.StateStore, basketStore basketapi.StateStore, subspace paramtypes.Subspace) error {

	return nil
}
