package marketplace

import (
	"testing"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	assert.NilError(t, s.marketStore.AllowedDenomTable().Insert(s.ctx, &marketplacev1.AllowedDenom{
		BankDenom:    "ufoo",
		DisplayDenom: "foo",
		Exponent:     5,
	}))
	assert.NilError(t, s.marketStore.AllowedDenomTable().Insert(s.ctx, &marketplacev1.AllowedDenom{
		BankDenom:    "ubar",
		DisplayDenom: "bar",
		Exponent:     11,
	}))

	s.accountKeeper.EXPECT().GetModuleAddress("gov").Return(s.addr).Times(2)
	_, err := s.k.AllowAskDenom(s.ctx, &marketplace.MsgAllowAskDenom{
		RootAddress: s.addr.String(),
		AddDenoms: []*marketplace.MsgAllowAskDenom_DenomInfo{
			{Denom: "ubaz", DisplayDenom: "baz", Exponent: 3},
		},
		RemoveDenoms: []string{"ufoo", "ubar"},
	})
	assert.NilError(t, err)

	it, err := s.marketStore.AllowedDenomTable().List(s.ctx, &marketplacev1.AllowedDenomPrimaryKey{})
	assert.NilError(t, err)
	count := 0
	for it.Next() {
		count++
		v, err := it.Value()
		assert.NilError(t, err)
		assert.Equal(t, "ubaz", v.BankDenom)
	}
	assert.Equal(t, 1, count)

	_, err = s.k.AllowAskDenom(s.ctx, &marketplace.MsgAllowAskDenom{RootAddress: sdk.AccAddress("not_governance").String()})
	assert.ErrorContains(t, err, "params can only be updated via governance")
}
