package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestCancel_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, batchDenom := s.setupClassProjectBatch(t)

	// Supply -> tradable: 10.5 , retired: 10.5
	// s.addr balance -> tradable 10.5 , retired 10.5

	_, err := s.k.Cancel(s.ctx, &core.MsgCancel{
		Owner: s.addr.String(),
		Credits: []*core.MsgCancel_CancelCredits{
			{
				BatchDenom: batchDenom,
				Amount:     "10.5",
			},
		},
	})
	assert.NilError(t, err)

	// we cancel 10.5 credits, removing them from the s.addr balance, as well as supply, resulting in 0 to both.

	sup, err := s.stateStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, sup.TradableAmount, "0.0")
	assert.Equal(t, sup.RetiredAmount, "10.5")

	bal, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, bal.TradableAmount, "0.0")
	assert.Equal(t, bal.RetiredAmount, "10.5")
}

func TestCancel_InsufficientFunds(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	_, err := s.k.Cancel(s.ctx, &core.MsgCancel{
		Owner: s.addr.String(),
		Credits: []*core.MsgCancel_CancelCredits{
			{
				BatchDenom: "C01-001-20200101-20210101-01",
				Amount:     "100000",
			},
		},
	})
	assert.ErrorContains(t, err, "insufficient funds")

}

func TestCancel_BadPrecision(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	_, err := s.k.Cancel(s.ctx, &core.MsgCancel{
		Owner: s.addr.String(),
		Credits: []*core.MsgCancel_CancelCredits{
			{
				BatchDenom: "C01-001-20200101-20210101-01",
				Amount:     "10.5290385029385820935",
			},
		},
	})
	assert.ErrorContains(t, err, "exceeds maximum decimal places")
}

func TestCancel_InvalidBatch(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	_, err := s.k.Cancel(s.ctx, &core.MsgCancel{
		Owner: s.addr.String(),
		Credits: []*core.MsgCancel_CancelCredits{
			{
				BatchDenom: "C00-00000000-00000000-01",
				Amount:     "100000",
			},
		},
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
