package core

import (
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_BatchesByClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert class
	assert.NilError(t, s.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id: "C01",
	}))

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

	batch1 := &api.Batch{
		Issuer:       s.addr,
		ProjectKey:   pKey,
		Denom:        "C01-001-20200101-20210101-001",
		Metadata:     "data",
		StartDate:    timestamppb.New(startTime),
		EndDate:      timestamppb.New(endTime),
		IssuanceDate: timestamppb.New(issuanceTime),
	}

	// insert two batches that are "C01" credit batches
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, batch1))
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{Denom: "C01-001-20200101-20210101-002"}))

	// insert two batches that are not "C01" credit batches
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{Denom: "C011-001-20200101-20210101-001"}))
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{Denom: "BIO1-001-20200101-20210101-001"}))

	// query batches by "C01" credit class
	res, err := s.k.BatchesByClass(s.ctx, &core.QueryBatchesByClassRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assertBatchEqual(t, s.ctx, s.k, res.Batches[0], batch1)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// query batches by unknown credit class
	_, err = s.k.BatchesByClass(s.ctx, &core.QueryBatchesByClassRequest{ClassId: "A00"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
