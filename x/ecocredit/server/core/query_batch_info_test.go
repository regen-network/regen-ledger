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
	endTime, err := types.ParseDate("", "2021-01-01")
	assert.NilError(t, err)
	issuanceTime, err := types.ParseDate("", "2022-01-01")
	assert.NilError(t, err)

	projectId := "P01"
	issuer := s.addr
	batchDenom := "C01-20200101-20220101-001"
	metadata := "data"
	startDate := timestamppb.New(startTime)
	endDate := timestamppb.New(endTime)
	issuanceDate := timestamppb.New(issuanceTime)

	assert.NilError(t, s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id: projectId,
	}))

	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		Issuer:       issuer,
		ProjectKey:   1,
		BatchDenom:   batchDenom,
		Metadata:     metadata,
		StartDate:    startDate,
		EndDate:      endDate,
		IssuanceDate: issuanceDate,
	}))

	// valid query
	res, err := s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: batchDenom})
	assert.NilError(t, err)
	assert.Equal(t, projectId, res.Batch.ProjectId)
	assert.Equal(t, batchDenom, res.Batch.BatchDenom)
	assert.Equal(t, metadata, res.Batch.Metadata)
	assert.Equal(t, issuer.String(), res.Batch.Issuer)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(startDate), res.Batch.StartDate)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(endDate), res.Batch.EndDate)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(issuanceDate), res.Batch.IssuanceDate)

	// invalid query
	_, err = s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: "A00-00000000-00000000-000"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// good query
	res, err = s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: batchDenom})
	assert.NilError(t, err)
	assert.Equal(t, projectId, res.Batch.ProjectId)
}
