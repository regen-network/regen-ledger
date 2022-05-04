package core

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
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

	batch := &api.Batch{
		ProjectKey: pKey,
		Denom:      "C01-20200101-20220101-001",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}

	// insert two batches issued under the "C01-001" project
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, batch))
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{
		ProjectKey: pKey,
		Denom:      "C01-20200101-20220101-002",
	}))

	// query batches by the "C01-001" project
	res, err := s.k.Batches(s.ctx, &core.QueryBatchesRequest{
		ProjectId:  "C01-001",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assertBatchEqual(t, s.ctx, s.k, res.Batches[0], batch)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	// query batches by unknown project
	_, err = s.k.Batches(s.ctx, &core.QueryBatchesRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func assertBatchEqual(t *testing.T, ctx context.Context, k Keeper, received *core.BatchInfo, batch *api.Batch) {
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
