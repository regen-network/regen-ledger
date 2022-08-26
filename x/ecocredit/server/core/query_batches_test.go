package core

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Batches(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert project
	pKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: "C01-001",
	})
	assert.NilError(t, err)

	startTime, err := types.ParseDate("start date", "2020-01-01")
	assert.NilError(t, err)

	endTime, err := types.ParseDate("end date", "2021-01-01")
	assert.NilError(t, err)

	issuanceTime, err := types.ParseDate("issuance date", "2022-01-01")
	assert.NilError(t, err)

	batch := &api.Batch{
		Issuer:       s.addr,
		ProjectKey:   pKey,
		Denom:        "C01-001-20200101-20210101-001",
		Metadata:     "data",
		StartDate:    timestamppb.New(startTime),
		EndDate:      timestamppb.New(endTime),
		IssuanceDate: timestamppb.New(issuanceTime),
	}

	// insert two batches issued under the "C01-001" project
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, batch))
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{
		ProjectKey: pKey,
		Denom:      "C01-001-20200101-20220101-002",
	}))

	// query all batches with pagination
	res, err := s.k.Batches(s.ctx, &core.QueryBatchesRequest{
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assertBatchEqual(s.ctx, t, s.k, res.Batches[0], batch)
	assert.Equal(t, uint64(2), res.Pagination.Total)
}

func assertBatchEqual(ctx context.Context, t *testing.T, k Keeper, received *core.BatchInfo, batch *api.Batch) {
	issuer := sdk.AccAddress(batch.Issuer)

	project, err := k.stateStore.ProjectTable().Get(ctx, batch.ProjectKey)
	assert.NilError(t, err)

	info := core.BatchInfo{
		Issuer:       issuer.String(),
		ProjectId:    project.Id,
		Denom:        batch.Denom,
		Metadata:     batch.Metadata,
		StartDate:    types.ProtobufToGogoTimestamp(batch.StartDate),
		EndDate:      types.ProtobufToGogoTimestamp(batch.EndDate),
		IssuanceDate: types.ProtobufToGogoTimestamp(batch.IssuanceDate),
		Open:         batch.Open,
	}

	assert.DeepEqual(t, info, *received)
}
