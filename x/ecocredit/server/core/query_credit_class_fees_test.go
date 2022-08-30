package core

import (
	"testing"

	"cosmossdk.io/math"
	"gotest.tools/v3/assert"

	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_CreditClassFees(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	result, err := s.k.CreditClassFees(s.ctx, &core.QueryCreditClassFeesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, result.Fees.Len(), 0)

	// initialize credit class fees
	err = s.stateStore.ClassFeesTable().Save(s.ctx, &ecocreditv1.ClassFees{
		Fees: []*basev1beta1.Coin{
			{
				Denom:  "uatom",
				Amount: "20000000",
			},
		},
	})
	assert.NilError(t, err)

	result, err = s.k.CreditClassFees(s.ctx, &core.QueryCreditClassFeesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, result.Fees.Len(), 1)
	assert.Equal(t, result.Fees.AmountOf("uatom").Equal(math.NewInt(2e7)), true)

	// initialize credit class fees
	err = s.stateStore.ClassFeesTable().Save(s.ctx, &ecocreditv1.ClassFees{
		Fees: []*basev1beta1.Coin{
			{
				Denom:  "uatom",
				Amount: "20000000",
			},
			{
				Denom:  "uregen",
				Amount: "10000000",
			},
		},
	})
	assert.NilError(t, err)
	result, err = s.k.CreditClassFees(s.ctx, &core.QueryCreditClassFeesRequest{})
	assert.NilError(t, err)

	assert.Equal(t, result.Fees.Len(), 2)
	assert.Equal(t, result.Fees.AmountOf("uregen").Equal(math.NewInt(1e7)), true)

}
