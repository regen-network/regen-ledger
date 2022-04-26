package core

import (
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

func TestMintBatchCredits_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	ctx := s.ctx
	_, _, batchDenom := s.setupClassProjectBatch(t)
	batch, err := s.stateStore.BatchTable().GetByDenom(ctx, batchDenom)
	assert.NilError(t, err)
	batch.Open = true
	assert.NilError(t, s.stateStore.BatchTable().Update(ctx, batch))

	balBefore, err := s.stateStore.BatchBalanceTable().Get(ctx, s.addr, batch.Key)
	assert.NilError(t, err)
	supplyBefore, err := s.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
	assert.NilError(t, err)

	msg := core.MsgMintBatchCredits{
		Issuer:     s.addr.String(),
		BatchDenom: batchDenom,
		Issuance: []*core.BatchIssuance{
			{Recipient: s.addr.String(), TradableAmount: "10", RetiredAmount: "10", RetirementJurisdiction: "US-OR"},
		},
		OriginTx: &core.OriginTx{
			Typ: "Ethereum",
			Id:  "210985091248",
		},
		Note: "bridged credits",
	}
	_, err = s.k.MintBatchCredits(ctx, &msg)
	assert.NilError(t, err)

	balAfter, err := s.stateStore.BatchBalanceTable().Get(ctx, s.addr, batch.Key)
	assert.NilError(t, err)
	supplyAfter, err := s.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
	assert.NilError(t, err)

	assertCreditsMoved(t, balBefore, balAfter, supplyBefore, supplyAfter, "10", "10", "0", 6)
}

func assertCreditsMoved(t *testing.T, balBefore, balAfter *api.BatchBalance, supBefore, supAfter *api.BatchSupply, tradable, retired, escrowed string, precision uint32) {
	decs, err := utils.GetNonNegativeFixedDecs(precision, tradable, retired, escrowed)
	assert.NilError(t, err)
	tradableDec, retiredDec, escrowedDec := decs[0], decs[1], decs[2]

	decs2, err := utils.GetNonNegativeFixedDecs(precision, balBefore.)
}
