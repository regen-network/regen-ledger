package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
	"testing"
)

func TestCancelValid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any, p *ecocredit.Params)  {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	_, err := s.k.Cancel(s.ctx, &v1beta1.MsgCancel{
		Holder: s.addr.String(),
		Credits: []*v1beta1.MsgCancel_CancelCredits{
			{
				BatchDenom: "C01-20200101-20210101-01",
				Amount: "10.5",
			},
		},
	})
	assert.NilError(t, err)

	sup, err := s.stateStore.BatchSupplyStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, sup.TradableAmount, "0.0")
	assert.Equal(t, sup.RetiredAmount, "10.5")

	bal, err := s.stateStore.BatchBalanceStore().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, bal.Tradable, "0.0")
	assert.Equal(t, bal.Retired, "10.5")
}

func TestInsufficientFunds(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any, p *ecocredit.Params)  {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	_, err := s.k.Cancel(s.ctx, &v1beta1.MsgCancel{
		Holder: s.addr.String(),
		Credits: []*v1beta1.MsgCancel_CancelCredits{
			{
				BatchDenom: "C01-20200101-20210101-01",
				Amount: "100000",
			},
		},
	})
	assert.ErrorContains(t, err, "insufficient funds")

}

func (s baseSuite) setupClassProjectBatch(t *testing.T) {
	assert.NilError(t, s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	}))
	assert.NilError(t, s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "PRO",
		ClassId:         1,
		ProjectLocation: "US-OR",
		Metadata:        nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20200101-20210101-01",
		Metadata:   nil,
		StartDate:  &timestamppb.Timestamp{Seconds: 2},
		EndDate:    &timestamppb.Timestamp{Seconds: 2},
	}))
	assert.NilError(t, s.stateStore.BatchSupplyStore().Insert(s.ctx, &ecocreditv1beta1.BatchSupply{
		BatchId:         1,
		TradableAmount:  "10.5",
		RetiredAmount:   "10.5",
		CancelledAmount: "",
	}))
	assert.NilError(t, s.stateStore.BatchBalanceStore().Insert(s.ctx, &ecocreditv1beta1.BatchBalance{
		Address:  s.addr,
		BatchId:  1,
		Tradable: "10.5",
		Retired:  "10.5",
	}))
}
