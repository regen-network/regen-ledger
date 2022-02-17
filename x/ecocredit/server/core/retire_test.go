package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestRetireState(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(_ interface{}, p *ecocredit.Params)  {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	_, err := s.k.Retire(s.ctx, &v1beta1.MsgRetire{
		Holder:   s.addr.String(),
		Credits:  []*v1beta1.MsgRetire_RetireCredits{
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
