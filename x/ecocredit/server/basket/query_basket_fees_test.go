package basket

import (
	"testing"

	"cosmossdk.io/math"
	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestQueryBasketFees(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// check empty basket fees
	res, err := s.k.BasketFees(s.ctx, &baskettypes.QueryBasketFeesRequest{})
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
	res, err = s.k.BasketFees(s.ctx, &baskettypes.QueryBasketFeesRequest{})
	require.NoError(s.t, err)
	require.Equal(s.t, res.Fees.Len(), 2)

	require.Equal(s.t, res.Fees.AmountOf("uregen"), math.NewInt(100))
	require.Equal(s.t, res.Fees.AmountOf("uatom"), math.NewInt(10))
}
