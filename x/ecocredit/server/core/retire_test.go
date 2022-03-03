package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestRetire_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(_ interface{}, p *ecocredit.Params) {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	_, err := s.k.Retire(s.ctx, &v1beta1.MsgRetire{
		Holder: s.addr.String(),
		Credits: []*v1beta1.MsgRetire_RetireCredits{
			{BatchDenom: "C01-20200101-20210101-01", Amount: "10.0"},
		},
		Location: "US-NY",
	})
	assert.NilError(t, err)

	bal, err := s.stateStore.BatchBalanceStore().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, bal.Tradable, "0.5")
	assert.Equal(t, bal.Retired, "20.5")

	sup, err := s.stateStore.BatchSupplyStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, sup.TradableAmount, "0.5")
	assert.Equal(t, sup.RetiredAmount, "20.5")
}

func TestRetire_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, batchDenom := s.setupClassProjectBatch(t)

	// invalid batch denom
	_, err := s.k.Retire(s.ctx, &v1beta1.MsgRetire{
		Holder:   s.addr.String(),
		Credits:  []*v1beta1.MsgRetire_RetireCredits{
			{BatchDenom: "A00-00000000-00000000-01", Amount: "10.35"},
		},
		Location: "US-NY",
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(_ interface{}, p *ecocredit.Params) {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	// out of precision
	_, err = s.k.Retire(s.ctx, &v1beta1.MsgRetire{
		Holder:   s.addr.String(),
		Credits:  []*v1beta1.MsgRetire_RetireCredits{
			{BatchDenom: batchDenom, Amount: "10.35250982359823095"},
		},
		Location: "US-NY",
	})
	assert.ErrorContains(t, err,"exceeds maximum decimal places")

	// not enough credits
	_, err = s.k.Retire(s.ctx, &v1beta1.MsgRetire{
		Holder:   s.addr.String(),
		Credits:  []*v1beta1.MsgRetire_RetireCredits{
			{BatchDenom: batchDenom, Amount: "150"},
		},
		Location: "US-NY",
	})
	assert.ErrorContains(t, err,errors.ErrInsufficientFunds.Error())
}
