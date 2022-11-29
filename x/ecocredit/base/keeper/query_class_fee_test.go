package keeper

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	sdkbase "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_ClassFee(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	result, err := s.k.ClassFee(s.ctx, &types.QueryClassFeeRequest{})
	require.NoError(t, err)
	require.Empty(t, result.Fee)

	// initialize credit class fee
	err = s.stateStore.ClassFeeTable().Save(s.ctx, &api.ClassFee{
		Fee: &sdkbase.Coin{
			Denom:  "uregen",
			Amount: "20000000",
		},
	})
	require.NoError(t, err)

	result, err = s.k.ClassFee(s.ctx, &types.QueryClassFeeRequest{})
	require.NoError(t, err)

	require.NotEmpty(t, result.Fee)
	require.Equal(t, "uregen", result.Fee.Denom)
	require.Equal(t, math.NewInt(2e7), result.Fee.Amount)
}
