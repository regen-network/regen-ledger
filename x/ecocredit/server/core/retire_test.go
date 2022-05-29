package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestRetire_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	// starting balance
	// tradable: 10.5
	// retired: 10.5

	// starting balance -> 10.5 tradable, 10.5 retired
	// retire 10.0 -> 0.5 leftover in tradable, retired becomes 20.5

	_, err := s.k.Retire(s.ctx, &core.MsgRetire{
		Holder: s.addr.String(),
		Credits: []*core.MsgRetire_RetireCredits{
			{BatchDenom: "C01-001-20200101-20210101-01", Amount: "10.0"},
		},
		Jurisdiction: "US-NY",
	})
	assert.NilError(t, err)

	// check both balance and supply reflect the change

	bal, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, bal.TradableAmount, "0.5")
	assert.Equal(t, bal.RetiredAmount, "20.5")

	sup, err := s.stateStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, sup.TradableAmount, "0.5")
	assert.Equal(t, sup.RetiredAmount, "20.5")
}

func TestRetire_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, batchDenom := s.setupClassProjectBatch(t)

	// invalid batch denom
	_, err := s.k.Retire(s.ctx, &core.MsgRetire{
		Holder: s.addr.String(),
		Credits: []*core.MsgRetire_RetireCredits{
			{BatchDenom: "A00-00000000-00000000-01", Amount: "10.35"},
		},
		Jurisdiction: "US-NY",
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// out of precision
	_, err = s.k.Retire(s.ctx, &core.MsgRetire{
		Holder: s.addr.String(),
		Credits: []*core.MsgRetire_RetireCredits{
			{BatchDenom: batchDenom, Amount: "10.35250982359823095"},
		},
		Jurisdiction: "US-NY",
	})
	assert.ErrorContains(t, err, "exceeds maximum decimal places")

	// not enough credits
	_, err = s.k.Retire(s.ctx, &core.MsgRetire{
		Holder: s.addr.String(),
		Credits: []*core.MsgRetire_RetireCredits{
			{BatchDenom: batchDenom, Amount: "150"},
		},
		Jurisdiction: "US-NY",
	})
	assert.ErrorContains(t, err, errors.ErrInsufficientFunds.Error())
}
