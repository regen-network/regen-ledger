package core

import (
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestBridge_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	batchDenom1 := "C01-001-20200101-20210101-001"
	batchDenom2 := "C01-001-20200101-20210101-002"
	recipient := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	contract := "0x06012c8cf97bead5deae237070f9587f8e7a266d"

	// insert credit class (used for both batches)
	cKey, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               "C01",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)

	// insert project (used for both batches)
	pKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		ClassKey: cKey,
	})
	assert.NilError(t, err)

	// insert credit batch (one with batch contract mapping)
	bKey1, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: pKey,
		Denom:      batchDenom1,
	})
	assert.NilError(t, err)

	// insert credit batch (one without batch contract mapping)
	bKey2, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: pKey,
		Denom:      batchDenom2,
	})
	assert.NilError(t, err)

	// insert batch supply (for batch with batch contract mapping)
	assert.NilError(t, s.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        bKey1,
		TradableAmount:  "10.5",
		RetiredAmount:   "10.5",
		CancelledAmount: "0",
	}))

	// insert batch supply (for batch without batch contract mapping)
	assert.NilError(t, s.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        bKey2,
		TradableAmount:  "10.5",
		RetiredAmount:   "10.5",
		CancelledAmount: "0",
	}))

	// insert batch balance (for batch with batch contract mapping)
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       bKey1,
		Address:        s.addr,
		TradableAmount: "10.5",
		RetiredAmount:  "10.5",
	}))

	// insert batch balance (for batch without batch contract mapping)
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       bKey2,
		Address:        s.addr,
		TradableAmount: "10.5",
		RetiredAmount:  "10.5",
	}))

	// insert batch contract mapping (only for first batch)
	err = s.stateStore.BatchContractTable().Insert(s.ctx, &api.BatchContract{
		BatchKey: bKey1,
		Contract: contract,
	})
	assert.NilError(t, err)

	// bridge credits from batch with contract mapping
	_, err = s.k.Bridge(s.ctx, &core.MsgBridge{
		Owner: s.addr.String(),
		Credits: []*core.Credits{
			{
				BatchDenom: batchDenom1,
				Amount:     "10.5",
			},
		},
		Target:    "polygon",
		Recipient: recipient,
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

	// bridge credits from batch without contract mapping
	_, err = s.k.Bridge(s.ctx, &core.MsgBridge{
		Owner: s.addr.String(),
		Credits: []*core.Credits{
			{
				BatchDenom: batchDenom2,
				Amount:     "10.5",
			},
		},
		Target:    "polygon",
		Recipient: recipient,
	})
	assert.ErrorContains(t, err, "only credits previously bridged from another chain")
}
