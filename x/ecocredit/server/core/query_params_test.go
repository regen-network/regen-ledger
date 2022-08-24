package core

import (
	"testing"

	"gotest.tools/v3/assert"

	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Params(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	err := s.stateStore.AllowedClassCreatorTable().Insert(s.ctx, &ecocreditv1.AllowedClassCreator{
		Address: s.addr,
	})
	assert.NilError(t, err)

	err = s.stateStore.AllowListEnabledTable().Save(s.ctx, &ecocreditv1.AllowListEnabled{
		Enabled: true,
	})
	assert.NilError(t, err)

	err = s.stateStore.ClassFeesTable().Save(s.ctx, &ecocreditv1.ClassFees{
		Fees: []*basev1beta1.Coin{
			{
				Denom:  sdk.DefaultBondDenom,
				Amount: "100",
			},
		},
	})
	assert.NilError(t, err)

	err = s.k.basketStore.BasketFeesTable().Save(s.ctx, &basketv1.BasketFees{
		Fees: []*basev1beta1.Coin{
			{
				Denom:  sdk.DefaultBondDenom,
				Amount: "1000",
			},
		},
	})
	assert.NilError(t, err)

	result, err := s.k.Params(s.ctx, &core.QueryParamsRequest{})
	assert.NilError(t, err)

	assert.Equal(t, result.Params.AllowlistEnabled, true)
	assert.DeepEqual(t, result.Params.AllowedClassCreators, []string{s.addr.String()})
	assert.Equal(t, result.Params.CreditClassFee.String(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))).String())
	assert.Equal(t, result.Params.BasketFee.String(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))).String())
}
