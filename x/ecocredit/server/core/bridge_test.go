package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestBridge_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, batchDenom := s.setupClassProjectBatch(t)
	recipient := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	contract := "0x06012c8cf97bead5deae237070f9587f8e7a266d"

	// Supply -> tradable: 10.5 , retired: 10.5
	// s.addr balance -> tradable 10.5 , retired 10.5

	_, err := s.k.Bridge(s.ctx, &core.MsgBridge{
		Holder: s.addr.String(),
		Credits: []*core.MsgBridge_CancelCredits{
			{
				BatchDenom: batchDenom,
				Amount:     "10.5",
			},
		},
		Target:    "polygon",
		Recipient: recipient,
		Contract:  contract,
	})
	assert.NilError(t, err)

	// we cancel 10.5 credits, removing them from the s.addr balance, as well as supply, resulting in 0 to both.

	sup, err := s.stateStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, sup.TradableAmount, "0.0")
	assert.Equal(t, sup.RetiredAmount, "10.5")

	bal, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, bal.Tradable, "0.0")
	assert.Equal(t, bal.Retired, "10.5")
}
