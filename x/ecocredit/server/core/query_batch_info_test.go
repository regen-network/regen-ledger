package core

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	project := &api.Project{
		Id: "P01",
	}

	pKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, project)
	assert.NilError(t, err)

	issuer, err := sdk.AccAddressFromBech32(s.addr.String())
	assert.NilError(t, err)

	startTime, err := types.ParseDate("", "2020-01-01")
	assert.NilError(t, err)

	endTime, err := types.ParseDate("", "2021-01-01")
	assert.NilError(t, err)

	issuanceTime, err := types.ParseDate("", "2022-01-01")
	assert.NilError(t, err)

	batch := &api.Batch{
		Issuer:       issuer,
		ProjectKey:   pKey,
		Denom:        "C01-20200101-20220101-001",
		Metadata:     "data",
		StartDate:    timestamppb.New(startTime),
		EndDate:      timestamppb.New(endTime),
		IssuanceDate: timestamppb.New(issuanceTime),
	}

	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, batch))

	// valid query
	res, err := s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: batch.Denom})
	assert.NilError(t, err)
	assert.Equal(t, project.Id, res.Batch.ProjectId)
	assert.Equal(t, batch.Denom, res.Batch.Denom)
	assert.Equal(t, batch.Metadata, res.Batch.Metadata)
	assert.Equal(t, issuer.String(), res.Batch.Issuer)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(batch.StartDate), res.Batch.StartDate)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(batch.EndDate), res.Batch.EndDate)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(batch.IssuanceDate), res.Batch.IssuanceDate)

	// invalid query
	_, err = s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: "A00-00000000-00000000-000"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
