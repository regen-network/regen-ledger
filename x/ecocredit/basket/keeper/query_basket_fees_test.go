package keeper

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func TestQueryBasketFees(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// check empty basket fees
	res, err := s.k.BasketFees(s.ctx, &types.QueryBasketFeesRequest{})
	require.NoError(s.t, err)
	require.Equal(s.t, res.Fees, sdk.NewCoins())

	// add some basket fees
	require.NoError(t, s.stateStore.BasketFeesTable().Save(s.ctx, &api.BasketFees{
		Fees: []*basev1beta1.Coin{
			{
				Amount: "10",
				Denom:  "uatom",
			},
			{
				Amount: "100",
				Denom:  "uregen",
			},
		},
	}))

	// query basket fees
	res, err = s.k.BasketFees(s.ctx, &types.QueryBasketFeesRequest{})
	require.NoError(s.t, err)
	require.Equal(s.t, res.Fees.Len(), 2)

	require.Equal(s.t, res.Fees.AmountOf("uregen"), math.NewInt(100))
	require.Equal(s.t, res.Fees.AmountOf("uatom"), math.NewInt(10))
}
