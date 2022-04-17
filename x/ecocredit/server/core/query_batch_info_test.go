package core

import (
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_BatchInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	startTime, err := types.ParseDate("", "2020-01-01")
	assert.NilError(t, err)
	endTime, err := types.ParseDate("", "2022-01-01")
	assert.NilError(t, err)

	projectId := "P01"
	batchDenom := "C01-20200101-20220101-001"
	metadata := "data"
	startDate := timestamppb.New(startTime)
	endDate := timestamppb.New(endTime)

	projectKey, err := s.stateStore.ProjectInfoTable().InsertReturningID(s.ctx, &api.ProjectInfo{
		Name: projectId,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  projectKey,
		BatchDenom: batchDenom,
		Metadata:   metadata,
		StartDate:  startDate,
		EndDate:    endDate,
	}))

	// valid query
	res, err := s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: batchDenom})
	assert.NilError(t, err)
	assert.Equal(t, projectId, res.Batch.ProjectId)
	assert.Equal(t, batchDenom, res.Batch.BatchDenom)
	assert.Equal(t, startDate.Seconds, res.Batch.StartDate.Seconds)
	assert.Equal(t, endDate.Seconds, res.Batch.EndDate.Seconds)

	// invalid query
	_, err = s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: "A00-00000000-00000000-000"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
