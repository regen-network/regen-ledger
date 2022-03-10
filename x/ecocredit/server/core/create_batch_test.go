package core

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestCreateBatch_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, projectName := batchTestSetup(t, s.ctx, s.stateStore, s.addr)
	_, _, addr2 := testdata.KeyTestPubAddr()

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).Times(1)

	start, end := time.Now(), time.Now()
	res, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "PRO",
		Issuance: []*core.MsgCreateBatch_BatchIssuance{
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
		Metadata:  nil,
		StartDate: &start,
		EndDate:   &end,
	})
	totalTradable := "12.4"
	totalRetired := "8.7"

	// check the batch
	batch, err := s.stateStore.BatchInfoStore().Get(s.ctx, 1)
	assert.NilError(t, err, "unexpected error: %w", err)
	assert.Equal(t, res.BatchDenom, batch.BatchDenom)

	// check the supply was set
	sup, err := s.stateStore.BatchSupplyStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, totalTradable, sup.TradableAmount, "got %s", sup.TradableAmount)
	assert.Equal(t, totalRetired, sup.RetiredAmount, "got %s", sup.RetiredAmount)
	assert.Equal(t, "0", sup.CancelledAmount, "got %s", sup.CancelledAmount)

	// check balances were allocated
	bal, err := s.stateStore.BatchBalanceStore().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, "10", bal.Tradable)
	assert.Equal(t, "5.3", bal.Retired)

	bal2, err := s.stateStore.BatchBalanceStore().Get(s.ctx, addr2, 1)
	assert.NilError(t, err)
	assert.Equal(t, "2.4", bal2.Tradable)
	assert.Equal(t, "3.4", bal2.Retired)

	// check sequence number
	seq, err := s.stateStore.BatchSequenceStore().Get(s.ctx, projectName)
	assert.NilError(t, err)
	assert.Equal(t, uint64(2), seq.NextBatchId)
}

func TestCreateBatch_BadPrecision(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).Times(1)

	start, end := time.Now(), time.Now()
	_, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "PRO",
		Issuance: []*core.MsgCreateBatch_BatchIssuance{
			{
				Recipient:      s.addr.String(),
				TradableAmount: "10.1234567891111",
			},
		},
		Metadata:  nil,
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
func batchTestSetup(t *testing.T, ctx context.Context, ss api.StateStore, addr types.AccAddress) (className, projectName string) {
	className, projectName = "C01", "PRO"
	cid, err := ss.ClassInfoStore().InsertReturningID(ctx, &api.ClassInfo{
		Name:       className,
		Admin:      addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	err = ss.ClassIssuerStore().Insert(ctx, &api.ClassIssuer{
		ClassId: cid,
		Issuer:  addr,
	})
	assert.NilError(t, err)
	_, err = ss.ProjectInfoStore().InsertReturningID(ctx, &api.ProjectInfo{
		Name:            projectName,
		ClassId:         1,
		ProjectLocation: "",
		Metadata:        nil,
	})
	assert.NilError(t, err)
	return
}
