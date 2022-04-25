package core

import (
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_BatchInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert project
	pKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: "P01",
	})
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

	// insert batch
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, batch))

	// query batch by "C01-20200101-20220101-001" batch denom
	res, err := s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: batch.Denom})
	assert.NilError(t, err)
	assert.Equal(t, "P01", res.Batch.ProjectId)
	assert.Equal(t, batch.Denom, res.Batch.Denom)
	assert.Equal(t, batch.Metadata, res.Batch.Metadata)
	assert.Equal(t, issuer.String(), res.Batch.Issuer)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(batch.StartDate), res.Batch.StartDate)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(batch.EndDate), res.Batch.EndDate)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(batch.IssuanceDate), res.Batch.IssuanceDate)

	// query batch by unknown batch denom
	_, err = s.k.BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: "A00-00000000-00000000-000"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
