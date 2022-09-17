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

func TestQueryBasketFee(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// check empty basket fees
	res, err := s.k.BasketFee(s.ctx, &types.QueryBasketFeeRequest{})
	require.NoError(s.t, err)
	require.Equal(s.t, res.Fee, &sdk.Coin{})

	// add a basket fee
	require.NoError(t, s.stateStore.BasketFeeTable().Save(s.ctx, &api.BasketFee{
		Fee: &basev1beta1.Coin{
			Amount: "100",
			Denom:  "uregen",
		},
	}))

	// query basket fee
	res, err = s.k.BasketFee(s.ctx, &types.QueryBasketFeeRequest{})
	require.NoError(s.t, err)

	require.NotEmpty(s.t, res.Fee)
	require.Equal(s.t, "uregen", res.Fee.Denom)
	require.Equal(s.t, math.NewInt(100), res.Fee.Amount)
}
