package core

import (
	"testing"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestBridgeReceive_ProjectAndBatchExist(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	projectRefId := "VCS-001"
	project, batch := setupBridgeTest(s, projectRefId)
	recipient := testutil.GenAddress()

	start, end := batch.StartDate.AsTime(), batch.EndDate.AsTime()
	msg := core.MsgBridgeReceive{
		ServiceAddress: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: recipient,
			Amount:    "3",
			OriginTx: &core.OriginTx{
				Id:     "0x1324092835908235",
				Source: "polygon:0x325325230958",
			},
			StartDate: &start,
			EndDate:   &end,
			Metadata:  "",
			Note:      "bridged from test",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  projectRefId,
			Jurisdiction: "US-KY",
			Metadata:     "hi",
			ClassId:      "C01",
		},
	}
	res, err := s.k.BridgeReceive(s.ctx, &msg)
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, project.Id)
	assert.Equal(t, res.BatchDenom, batch.Denom)

	// this was a fresh account, so we know their balance is only what was bridged to it.
	bal, err := s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    recipient,
		BatchDenom: batch.Denom,
	})
	assert.NilError(t, err)
	balAfter := bal.Balance
	assert.Equal(t, balAfter.TradableAmount, "3")
}

func TestBridgeReceive_ProjectNoBatch(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	recipient := testutil.GenAddress()
	refId := "VCS-001"
	project, batch := setupBridgeTest(s, refId)
	startDate, endDate := (&timestamppb.Timestamp{Seconds: 500}).AsTime(), (&timestamppb.Timestamp{Seconds: 600}).AsTime()

	msg := core.MsgBridgeReceive{
		ServiceAddress: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: recipient,
			Amount:    "3",
			OriginTx: &core.OriginTx{
				Id:     "0x12345",
				Source: "polygon:0x12345",
			},
			StartDate: &startDate,
			EndDate:   &endDate,
			Metadata:  "hi",
			Note:      "bridged test",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  refId,
			Jurisdiction: "US-KY",
			ClassId:      "C01",
		},
	}

	res, err := s.k.BridgeReceive(s.ctx, &msg)
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, project.Id)
	assert.Check(t, res.BatchDenom != batch.Denom)

	batch, err = s.stateStore.BatchTable().GetByDenom(s.ctx, res.BatchDenom)
	assert.NilError(t, err)

	bal, err := s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    recipient,
		BatchDenom: batch.Denom,
	})
	assert.NilError(t, err)
	assert.Equal(t, bal.Balance.TradableAmount, msg.Batch.Amount)
}

func TestBridgeReceive_None(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	setupBridgeTest(s, "VCS-002")
	recipient := testutil.GenAddress()
	start, end := time.Now(), time.Now()
	msg := core.MsgBridgeReceive{
		ServiceAddress: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: recipient,
			Amount:    "3",
			OriginTx: &core.OriginTx{
				Id:     "0x12345",
				Source: "polygon:0x12345",
			},
			StartDate: &start,
			EndDate:   &end,
			Metadata:  "bar",
			Note:      "bridged",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  "VCS-001",
			Jurisdiction: "US-KY",
			Metadata:     "foo",
			ClassId:      "C01",
		},
	}
	res, err := s.k.BridgeReceive(s.ctx, &msg)
	assert.NilError(t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, res.BatchDenom)
	assert.NilError(t, err)
	bal, err := s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    recipient,
		BatchDenom: batch.Denom,
	})
	assert.NilError(t, err)
	assert.Equal(t, bal.Balance.TradableAmount, msg.Batch.Amount)
}

func TestBridgeReceive_TooManyProjects(t *testing.T) {
	t.Parallel()
	refId := "VCS-001"
	s := setupBase(t)
	setupBridgeTest(s, refId)
	err := s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:           "C01-002",
		Admin:        s.addr,
		ClassKey:     1,
		Jurisdiction: "US-KY",
		Metadata:     "hi",
		ReferenceId:  refId,
	})
	assert.NilError(t, err)

	msg := core.MsgBridgeReceive{
		ServiceAddress: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: testutil.GenAddress(),
			Amount:    "3",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId: refId,
		},
	}
	_, err = s.k.BridgeReceive(s.ctx, &msg)
	assert.ErrorIs(t, err, sdkerrors.ErrInvalidRequest.Wrapf("fatal error: bridge service %s has %d projects registered "+
		"with reference id %s", s.addr.String(), 2, refId))
}

func TestBridgeReceive_TooManyBatches(t *testing.T) {
	t.Parallel()
	refId := "VCS-001"
	s := setupBase(t)
	project, batch := setupBridgeTest(s, refId)

	// create 2 batches with "hi" as metadata.
	batchMetadata := "hi"
	start, end := batch.StartDate.AsTime(), batch.EndDate.AsTime()
	_, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: project.Id,
		Issuance: []*core.BatchIssuance{
			{Recipient: s.addr.String(), TradableAmount: "10"},
		},
		Metadata:  batchMetadata,
		StartDate: &start,
		EndDate:   &end,
		Open:      true,
		OriginTx:  &core.OriginTx{Id: "0x12345", Source: "polygon:0x12345"},
		Note:      "hi",
	})
	assert.NilError(t, err)
	_, err = s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: project.Id,
		Issuance: []*core.BatchIssuance{
			{Recipient: s.addr.String(), TradableAmount: "10"},
		},
		Metadata:  batchMetadata,
		StartDate: &start,
		EndDate:   &end,
		Open:      true,
		OriginTx:  &core.OriginTx{Id: "0x123456", Source: "polygon:0x12345"},
		Note:      "hi",
	})
	assert.NilError(t, err)

	msg := core.MsgBridgeReceive{
		ServiceAddress: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: testutil.GenAddress(),
			Amount:    "3",
			StartDate: &start,
			EndDate:   &end,
			Metadata:  batchMetadata,
			Note:      "bridged",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  refId,
			Jurisdiction: "US-KY",
			Metadata:     "hi",
			ClassId:      "C01",
		},
	}
	_, err = s.k.BridgeReceive(s.ctx, &msg)
	assert.ErrorIs(t, err, sdkerrors.ErrInvalidRequest.Wrapf("fatal error: bridge service %s has %d batches issued "+
		"with start %v and end %v dates in project %s", s.addr.String(), 2, msg.Batch.StartDate.String(), msg.Batch.EndDate.String(), project.Id))
}

func setupBridgeTest(s *baseSuite, refId string) (project *api.Project, batch *api.Batch) {
	var err error
	_, projectId, batchDenom := s.setupClassProjectBatch(s.t)
	batch, err = s.stateStore.BatchTable().GetByDenom(s.ctx, batchDenom)
	assert.NilError(s.t, err)
	batch.Open = true
	assert.NilError(s.t, s.stateStore.BatchTable().Update(s.ctx, batch))
	project, err = s.stateStore.ProjectTable().GetById(s.ctx, projectId)
	assert.NilError(s.t, err)
	project.ReferenceId = refId
	assert.NilError(s.t, s.stateStore.ProjectTable().Update(s.ctx, project))
	return
}
