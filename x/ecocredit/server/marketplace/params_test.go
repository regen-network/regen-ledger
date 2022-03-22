package marketplace

import (
	"testing"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	assert.NilError(t, s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom:    "ufoo",
		DisplayDenom: "foo",
		Exponent:     5,
	}))
	assert.NilError(t, s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom:    "ubar",
		DisplayDenom: "bar",
		Exponent:     11,
	}))

	s.accountKeeper.EXPECT().GetModuleAddress("gov").Return(s.addr).Times(4)
	_, err := s.k.AllowAskDenom(s.ctx, &marketplace.MsgAllowAskDenom{
		RootAddress: s.addr.String(),
		AddDenoms: []*marketplace.MsgAllowAskDenom_DenomInfo{
			{Denom: "ubaz", DisplayDenom: "baz", Exponent: 3},
		},
		RemoveDenoms: []string{"ufoo", "ubar"},
	})
	assert.NilError(t, err)

	it, err := s.marketStore.AllowedDenomTable().List(s.ctx, &api.AllowedDenomPrimaryKey{})
	assert.NilError(t, err)
	count := 0
	for it.Next() {
		count++
		v, err := it.Value()
		assert.NilError(t, err)
		assert.Equal(t, "ubaz", v.BankDenom)
	}
	assert.Equal(t, 1, count)

	// can't add duplicates
	_, err = s.k.AllowAskDenom(s.ctx, &marketplace.MsgAllowAskDenom{
		RootAddress: s.addr.String(),
		AddDenoms: []*marketplace.MsgAllowAskDenom_DenomInfo{
			{Denom: "ubaz", DisplayDenom: "baz", Exponent: 15},
		},
		RemoveDenoms: nil,
	})
	assert.ErrorContains(t, err, ormerrors.PrimaryKeyConstraintViolation.Error())

	// should be able to delete then add
	_, err = s.k.AllowAskDenom(s.ctx, &marketplace.MsgAllowAskDenom{
		RootAddress: s.addr.String(),
		AddDenoms: []*marketplace.MsgAllowAskDenom_DenomInfo{
			{Denom: "ubaz", DisplayDenom: "baz", Exponent: 15},
		},
		RemoveDenoms: []string{"ubaz"},
	})
	assert.NilError(t, err)
	v, err := s.marketStore.AllowedDenomTable().Get(s.ctx, "ubaz")
	assert.NilError(t, err)
	assert.Equal(t, uint32(15), v.Exponent)

	// unauthorized
	_, err = s.k.AllowAskDenom(s.ctx, &marketplace.MsgAllowAskDenom{RootAddress: sdk.AccAddress("not_governance").String()})
	assert.ErrorContains(t, err, "params can only be updated via governance")
}
