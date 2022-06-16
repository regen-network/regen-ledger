package core

import (
	"testing"
	"time"

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
		Issuer: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: recipient,
			Amount:    "3",
			StartDate: &start,
			EndDate:   &end,
			Metadata:  "",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  projectRefId,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
		},
		OriginTx: &core.OriginTx{
			Id:     "0x1324092835908235",
			Source: "polygon:0x325325230958",
		},
		Note:    "bridged from test",
		ClassId: "C01",
	}
	res, err := s.k.BridgeReceive(s.ctx, &msg)
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, project.Id)
	assert.Equal(t, res.BatchDenom, batch.Denom)
	assertBatchBridged(t, msg.Batch, batch)
	assertProjectBridged(t, msg.Project, project)

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
		Issuer: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: recipient,
			Amount:    "3",
			StartDate: &startDate,
			EndDate:   &endDate,
			Metadata:  "hi",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  refId,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
		},
		OriginTx: &core.OriginTx{
			Id:     "0x12345",
			Source: "polygon:0x12345",
		},
		ClassId: "C01",
		Note:    "bridged test",
	}

	res, err := s.k.BridgeReceive(s.ctx, &msg)
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, project.Id)
	assertProjectBridged(t, msg.Project, project)

	assert.Check(t, res.BatchDenom != batch.Denom)
	batch, err = s.stateStore.BatchTable().GetByDenom(s.ctx, res.BatchDenom)
	assert.NilError(t, err)
	assertBatchBridged(t, msg.Batch, batch)

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
	refId := "VCS-001"
	msg := core.MsgBridgeReceive{
		Issuer: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: recipient,
			Amount:    "3",

			StartDate: &start,
			EndDate:   &end,
			Metadata:  "bar",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  refId,
			Jurisdiction: "US-KY",
			Metadata:     "foo",
		},
		OriginTx: &core.OriginTx{
			Id:     "0x12345",
			Source: "polygon:0x12345",
		},
		ClassId: "C01",
		Note:    "bridged",
	}
	res, err := s.k.BridgeReceive(s.ctx, &msg)
	assert.NilError(t, err)

	// check ref id is correct
	project, err := s.stateStore.ProjectTable().GetById(s.ctx, res.ProjectId)
	assert.NilError(t, err)
	assert.Equal(t, project.ReferenceId, refId)
	assertProjectBridged(t, msg.Project, project)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, res.BatchDenom)
	assert.NilError(t, err)
	bal, err := s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    recipient,
		BatchDenom: batch.Denom,
	})
	assert.NilError(t, err)
	assert.Equal(t, bal.Balance.TradableAmount, msg.Batch.Amount)
	assertBatchBridged(t, msg.Batch, batch)
}

func TestBridgeReceive_MultipleProjects(t *testing.T) {
	t.Parallel()
	refId := "VCS-001"
	s := setupBase(t)
	project, batch := setupBridgeTest(s, refId)

	project2 := &api.Project{
		Id:           "C01-002",
		Admin:        s.addr,
		ClassKey:     project.ClassKey,
		Jurisdiction: "US-KY",
		Metadata:     "project2",
		ReferenceId:  refId,
	}
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, project2))

	start, end := batch.StartDate.AsTime(), batch.EndDate.AsTime()
	msg := core.MsgBridgeReceive{
		Issuer: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: s.addr.String(),
			Amount:    "3",
			StartDate: &start,
			EndDate:   &end,
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  refId,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
		},
		OriginTx: &core.OriginTx{
			Id:     "0x12345",
			Source: "polygon",
		},
		ClassId: "C01",
	}
	res, err := s.k.BridgeReceive(s.ctx, &msg)
	assert.NilError(t, err)
	// check to make sure the first project is selected
	assert.Equal(t, res.ProjectId, project.Id)
	assertProjectBridged(t, msg.Project, project)
	assertBatchBridged(t, msg.Batch, batch)
}

func TestBridgeReceive_ChoosesOldestBatch(t *testing.T) {
	t.Parallel()
	refId := "VCS-001"
	s := setupBase(t)
	project, batch := setupBridgeTest(s, refId)

	// set up a 2nd batch with same data as first, but an older issuance date.
	// the method should pick this first.
	oldTime := batch.IssuanceDate.AsTime().Add(time.Hour * -3)
	denom2 := batch.Denom[:len(batch.Denom)-1] + "2" // the previous batch denom but -002 instead of -001
	batch2 := &api.Batch{
		Issuer:       s.addr,
		ProjectKey:   batch.ProjectKey,
		Denom:        denom2,
		Metadata:     batch.Metadata,
		StartDate:    batch.StartDate,
		EndDate:      batch.EndDate,
		IssuanceDate: timestamppb.New(oldTime),
		Open:         true,
	}
	b2key, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, batch2)
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        b2key,
		TradableAmount:  "1",
		RetiredAmount:   "1",
		CancelledAmount: "1",
	}))

	start, end := batch.StartDate.AsTime(), batch.EndDate.AsTime()
	msg := &core.MsgBridgeReceive{
		Issuer: s.addr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: s.addr.String(),
			Amount:    "3",
			StartDate: &start,
			EndDate:   &end,
			Metadata:  batch.Metadata,
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  project.ReferenceId,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
		},
		OriginTx: &core.OriginTx{
			Id:     "0x12345",
			Source: "polygon",
		},
		Note:    "test",
		ClassId: "C01",
	}

	res, err := s.k.BridgeReceive(s.ctx, msg)
	assert.NilError(t, err)
	// ensure the 2nd batch was picked, since it was manually set to be an older issuance date than the first.
	assert.Equal(t, res.BatchDenom, batch2.Denom)
	assertBatchBridged(t, msg.Batch, batch2)
	assertProjectBridged(t, msg.Project, project)

}

func assertProjectBridged(t *testing.T, bridgedProject *core.MsgBridgeReceive_Project, project *api.Project) {
	assert.Equal(t, bridgedProject.Jurisdiction, project.Jurisdiction)
	assert.Equal(t, bridgedProject.Metadata, project.Metadata)
	assert.Equal(t, bridgedProject.ReferenceId, project.ReferenceId)
}

func assertBatchBridged(t *testing.T, bridgedBatch *core.MsgBridgeReceive_Batch, batch *api.Batch) {
	assert.Equal(t, bridgedBatch.Metadata, batch.Metadata)
	assert.Check(t, bridgedBatch.StartDate.UTC().Equal(batch.StartDate.AsTime().UTC()))
	assert.Check(t, bridgedBatch.EndDate.UTC().Equal(batch.EndDate.AsTime().UTC()))
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
