package core

import (
	"context"
	"strings"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_BatchesByClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	startTime, err := types.ParseDate("", "2020-01-01")
	assert.NilError(t, err)
	endTime, err := types.ParseDate("", "2021-01-01")
	assert.NilError(t, err)
	issuanceTime, err := types.ParseDate("", "2022-01-01")
	assert.NilError(t, err)

	batch1 := &api.BatchInfo{
		Issuer:       s.addr,
		ProjectKey:   1,
		BatchDenom:   "C01-20200101-20200102-001",
		Metadata:     "data",
		StartDate:    timestamppb.New(startTime),
		EndDate:      timestamppb.New(endTime),
		IssuanceDate: timestamppb.New(issuanceTime),
	}

	// insert class
	assert.NilError(t, s.stateStore.ClassInfoTable().Insert(s.ctx, &api.ClassInfo{
		Id: "C01",
	}))

	// insert project
	assert.NilError(t, s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id: "P01",
	}))

	// insert three batches under the project
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, batch1))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectKey: 1,
		BatchDenom: "C01-20190203-20200102-002",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectKey: 1,
		BatchDenom: "C01-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	// classes that SHOULD NOT show up from a query for "C01"
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectKey: 1,
		BatchDenom: "C011-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		ProjectKey: 1,
		BatchDenom: "BIO1-20500404-20900102-003",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	res, err := s.k.BatchesByClass(s.ctx, &core.QueryBatchesByClassRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{CountTotal: true, Limit: 2},
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Batches))
	assertBatchEqual(t, s.ctx, s.k, res.Batches[1], batch1)
	assert.Equal(t, uint64(3), res.Pagination.Total)
	for _, batch := range res.Batches {
		assert.Check(t, strings.Contains(batch.BatchDenom, "C01"))
	}
}

func assertBatchEqual(t *testing.T, ctx context.Context, k Keeper, received *core.BatchDetails, batch *api.BatchInfo) {
	issuer := sdk.AccAddress(batch.Issuer)

	project, err := k.stateStore.ProjectInfoTable().Get(ctx, batch.ProjectKey)
	assert.NilError(t, err)

	info := core.BatchDetails{
		Issuer:       issuer.String(),
		ProjectId:    project.Id,
		BatchDenom:   batch.BatchDenom,
		Metadata:     batch.Metadata,
		StartDate:    types.ProtobufToGogoTimestamp(batch.StartDate),
		EndDate:      types.ProtobufToGogoTimestamp(batch.EndDate),
		IssuanceDate: types.ProtobufToGogoTimestamp(batch.IssuanceDate),
		Open:         batch.Open,
	}

	assert.DeepEqual(t, info, *received)
}
