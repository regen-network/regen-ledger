package core

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

func TestMintBatchCredits_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	ctx := s.ctx
	batch := setupMintBatchTest(s, true)

	balBefore, err := s.stateStore.BatchBalanceTable().Get(ctx, s.addr, batch.Key)
	assert.NilError(t, err)
	supplyBefore, err := s.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
	assert.NilError(t, err)

	mintTradable, mintRetired := math.NewDecFromInt64(10), math.NewDecFromInt64(10)
	issuance := core.BatchIssuance{
		Recipient:              s.addr.String(),
		TradableAmount:         mintTradable.String(),
		RetiredAmount:          mintRetired.String(),
		RetirementJurisdiction: "US-OR",
	}
	msg := core.MsgMintBatchCredits{
		Issuer:     s.addr.String(),
		BatchDenom: batch.Denom,
		Issuance:   []*core.BatchIssuance{&issuance},
		OriginTx: &core.OriginTx{
			Id:     "210985091248",
			Source: "Ethereum",
		},
		Note: "bridged credits",
	}

	_, err = s.k.MintBatchCredits(ctx, &msg)
	assert.NilError(t, err)

	balAfter, err := s.stateStore.BatchBalanceTable().Get(ctx, s.addr, batch.Key)
	assert.NilError(t, err)
	supplyAfter, err := s.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
	assert.NilError(t, err)

	assertCreditsMinted(t, balBefore, balAfter, supplyBefore, supplyAfter, issuance, 6)
}

func TestMintBatchCredits_Unauthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batch := setupMintBatchTest(s, true)
	addr := sdk.AccAddress("foobar")

	_, err := s.k.MintBatchCredits(s.ctx, &core.MsgMintBatchCredits{
		Issuer:     addr.String(),
		BatchDenom: batch.Denom,
		OriginTx: &core.OriginTx{
			Id:     "210985091248",
			Source: "Ethereum",
		},
	})
	assert.ErrorContains(t, err, "unauthorized")
}

func TestMintBatchCredits_ClosedBatch(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batch := setupMintBatchTest(s, false)
	addr := sdk.AccAddress("foobar")

	_, err := s.k.MintBatchCredits(s.ctx, &core.MsgMintBatchCredits{
		Issuer:     addr.String(),
		BatchDenom: batch.Denom,
		OriginTx: &core.OriginTx{
			Id:     "210985091248",
			Source: "Ethereum",
		},
	})
	assert.ErrorContains(t, err, "credits cannot be minted in a closed batch")
}

func TestMintBatchCredits_NotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	setupMintBatchTest(s, true)
	addr := sdk.AccAddress("foobar")

	_, err := s.k.MintBatchCredits(s.ctx, &core.MsgMintBatchCredits{
		Issuer:     addr.String(),
		BatchDenom: "C05-00000000-00000000-001",
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func TestMintBatchCredits_SameTxId(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	ctx := s.ctx
	batch := setupMintBatchTest(s, true)

	mintTradable, mintRetired := math.NewDecFromInt64(10), math.NewDecFromInt64(10)
	issuance := core.BatchIssuance{
		Recipient:              s.addr.String(),
		TradableAmount:         mintTradable.String(),
		RetiredAmount:          mintRetired.String(),
		RetirementJurisdiction: "US-OR",
	}
	msg := core.MsgMintBatchCredits{
		Issuer:     s.addr.String(),
		BatchDenom: batch.Denom,
		Issuance:   []*core.BatchIssuance{&issuance},
		OriginTx: &core.OriginTx{
			Id:     "210985091248",
			Source: "Ethereum",
		},
		Note: "bridged credits",
	}

	_, err := s.k.MintBatchCredits(ctx, &msg)
	assert.NilError(t, err)

	_, err = s.k.MintBatchCredits(ctx, &msg)
	assert.ErrorContains(t, err, "credits already issued with tx id")
}

func setupMintBatchTest(s *baseSuite, open bool) *api.Batch {
	ctx := s.ctx
	_, _, batchDenom := s.setupClassProjectBatch(s.t)
	batch, err := s.stateStore.BatchTable().GetByDenom(ctx, batchDenom)
	assert.NilError(s.t, err)
	batch.Open = open
	assert.NilError(s.t, s.stateStore.BatchTable().Update(ctx, batch))
	return batch
}

func assertCreditsMinted(t *testing.T, balBefore, balAfter *api.BatchBalance, supBefore, supAfter *api.BatchSupply, issuance core.BatchIssuance, precision uint32) {
	checkFunc := func(before, after, change math.Dec) {
		expected, err := before.Add(change)
		assert.NilError(t, err)
		assert.Check(t, after.Equal(expected), fmt.Sprintf("expected %s got %s", expected.String(), after.String()))
	}

	issuanceDecs, err := utils.GetNonNegativeFixedDecs(precision, issuance.TradableAmount, issuance.RetiredAmount)
	assert.NilError(t, err)
	tradable, retired := issuanceDecs[0], issuanceDecs[1]

	tradableBefore, retiredBefore, _ := extractBalanceDecs(t, balBefore, precision)
	tradableAfter, retiredAfter, _ := extractBalanceDecs(t, balAfter, precision)
	checkFunc(tradableBefore, tradableAfter, tradable)
	checkFunc(retiredBefore, retiredAfter, retired)

	supTBefore, supRBefore, _ := extractSupplyDecs(t, supBefore, precision)
	supTAfter, supRAfter, _ := extractSupplyDecs(t, supAfter, precision)
	checkFunc(supTBefore, supTAfter, tradable)
	checkFunc(supRBefore, supRAfter, retired)

}

func extractBalanceDecs(t *testing.T, b *api.BatchBalance, precision uint32) (tradable, retired, escrowed math.Dec) {
	decs, err := utils.GetNonNegativeFixedDecs(precision, b.TradableAmount, b.RetiredAmount, b.EscrowedAmount)
	assert.NilError(t, err)
	return decs[0], decs[1], decs[2]
}

func extractSupplyDecs(t *testing.T, s *api.BatchSupply, precision uint32) (tradable, retired, cancelled math.Dec) {
	decs, err := utils.GetNonNegativeFixedDecs(precision, s.TradableAmount, s.RetiredAmount, s.CancelledAmount)
	assert.NilError(t, err)
	return decs[0], decs[1], decs[2]
}
