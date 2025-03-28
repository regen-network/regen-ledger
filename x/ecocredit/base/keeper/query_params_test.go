package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	sdkbase "cosmossdk.io/api/cosmos/base/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"

	baskettypes "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	markettypes "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_Params(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	err := s.stateStore.AllowedClassCreatorTable().Insert(s.ctx, &api.AllowedClassCreator{
		Address: s.addr,
	})
	assert.NilError(t, err)

	err = s.stateStore.ClassCreatorAllowlistTable().Save(s.ctx, &api.ClassCreatorAllowlist{
		Enabled: true,
	})
	assert.NilError(t, err)

	err = s.stateStore.ClassFeeTable().Save(s.ctx, &api.ClassFee{
		Fee: &sdkbase.Coin{
			Denom:  sdk.DefaultBondDenom,
			Amount: "100",
		},
	})
	assert.NilError(t, err)

	err = s.k.basketStore.BasketFeeTable().Save(s.ctx, &baskettypes.BasketFee{
		Fee: &sdkbase.Coin{
			Denom:  sdk.DefaultBondDenom,
			Amount: "1000",
		},
	})
	assert.NilError(t, err)

	err = s.k.marketStore.AllowedDenomTable().Insert(s.ctx, &markettypes.AllowedDenom{
		BankDenom:    "uregen",
		DisplayDenom: "REGEN",
		Exponent:     6,
	})
	assert.NilError(t, err)

	allowedChain := "polygon"
	assert.NilError(t, s.stateStore.AllowedBridgeChainTable().Insert(s.ctx, &api.AllowedBridgeChain{ChainName: allowedChain}))

	result, err := s.k.Params(s.ctx, &types.QueryParamsRequest{})
	assert.NilError(t, err)

	assert.Equal(t, result.Params.AllowlistEnabled, true)
	assert.DeepEqual(t, result.Params.AllowedClassCreators, []string{s.addr.String()})
	assert.Equal(t, result.Params.CreditClassFee.String(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))).String())
	assert.Equal(t, result.Params.BasketFee.String(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))).String())
	assert.Equal(t, len(result.Params.AllowedDenoms), 1)
	assert.Equal(t, result.Params.AllowedDenoms[0].BankDenom, "uregen")
	assert.Equal(t, len(result.Params.AllowedBridgeChains), 1)
	assert.Equal(t, result.Params.AllowedBridgeChains[0], allowedChain)
}
