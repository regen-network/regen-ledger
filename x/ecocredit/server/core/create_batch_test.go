package core

import (
	"context"
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestCreateBatch_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)
	_, _, addr2 := testdata.KeyTestPubAddr()

	blockTime, err := time.Parse("2006-01-02", "2049-01-30")
	assert.NilError(t, err)
	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = types.WrapSDKContext(s.sdkCtx)

	start, end := time.Now(), time.Now()
	res, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "PRO",
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.addr.String(),
				TradableAmount: "10",
				RetiredAmount:  "5.3",
			},
			{
				Recipient:      addr2.String(),
				TradableAmount: "2.4",
				RetiredAmount:  "3.4",
			},
		},
		Metadata:  "",
		StartDate: &start,
		EndDate:   &end,
	})
	totalTradable := "12.4"
	totalRetired := "8.7"

	// check the batch
	batch, err := s.stateStore.BatchTable().Get(s.ctx, 1)
	assert.NilError(t, err, "unexpected error: %w", err)
	assert.Equal(t, res.BatchDenom, batch.Denom)
	assert.Check(t, batch.IssuanceDate.AsTime().Equal(blockTime))

	// check the supply was set
	sup, err := s.stateStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, totalTradable, sup.TradableAmount, "got %s", sup.TradableAmount)
	assert.Equal(t, totalRetired, sup.RetiredAmount, "got %s", sup.RetiredAmount)
	assert.Equal(t, "0", sup.CancelledAmount, "got %s", sup.CancelledAmount)

	// check balances were allocated
	bal, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, "10", bal.Tradable)
	assert.Equal(t, "5.3", bal.Retired)

	bal2, err := s.stateStore.BatchBalanceTable().Get(s.ctx, addr2, 1)
	assert.NilError(t, err)
	assert.Equal(t, "2.4", bal2.Tradable)
	assert.Equal(t, "3.4", bal2.Retired)

	// check sequence number
	seq, err := s.stateStore.BatchSequenceTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, uint64(2), seq.NextSequence)
}

func TestCreateBatch_BadPrecision(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)

	start, end := time.Now(), time.Now()
	_, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "PRO",
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.addr.String(),
				TradableAmount: "10.1234567891111",
			},
		},
		Metadata:  "",
		StartDate: &start,
		EndDate:   &end,
	})
	assert.ErrorContains(t, err, "exceeds maximum decimal places")
}

func TestCreateBatch_UnauthorizedIssuer(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)
	_, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		ProjectId: "PRO",
		Issuer:    types.AccAddress("FooBarBaz").String(),
	})
	assert.ErrorContains(t, err, "is not an issuer for the class")
}

func TestCreateBatch_ProjectNotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	_, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		ProjectId: "none",
	})
	assert.ErrorContains(t, err, "not found")
}

// creates a class "C01", with a single class issuer, and a project "PRO"
func batchTestSetup(t *testing.T, ctx context.Context, ss api.StateStore, addr types.AccAddress) (classId, projectId string) {
	classId, projectId = "C01", "PRO"
	classKey, err := ss.ClassTable().InsertReturningID(ctx, &api.Class{
		Id:               classId,
		Admin:            addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)
	err = ss.ClassIssuerTable().Insert(ctx, &api.ClassIssuer{
		ClassKey: classKey,
		Issuer:   addr,
	})
	assert.NilError(t, err)
	_, err = ss.ProjectTable().InsertReturningID(ctx, &api.Project{
		Id:                  projectId,
		ClassKey:            classKey,
		ProjectJurisdiction: "",
		Metadata:            "",
	})
	assert.NilError(t, err)
	return
}
