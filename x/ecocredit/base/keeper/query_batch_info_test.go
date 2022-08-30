package keeper

import (
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestQuery_Batch(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert project
	pKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: "C01-001",
	})
	assert.NilError(t, err)

	issuer, err := sdk.AccAddressFromBech32(s.addr.String())
	assert.NilError(t, err)

	startTime, err := regentypes.ParseDate("", "2020-01-01")
	assert.NilError(t, err)

	endTime, err := regentypes.ParseDate("", "2021-01-01")
	assert.NilError(t, err)

	issuanceTime, err := regentypes.ParseDate("", "2022-01-01")
	assert.NilError(t, err)

	batch := &api.Batch{
		Issuer:       issuer,
		ProjectKey:   pKey,
		Denom:        "C01-001-20200101-20220101-001",
		Metadata:     "data",
		StartDate:    timestamppb.New(startTime),
		EndDate:      timestamppb.New(endTime),
		IssuanceDate: timestamppb.New(issuanceTime),
	}

	// insert batch
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, batch))

	// query batch by "C01-001-20200101-20220101-001" batch denom
	res, err := s.k.Batch(s.ctx, &types.QueryBatchRequest{BatchDenom: batch.Denom})
	assert.NilError(t, err)
	assert.Equal(t, "C01-001", res.Batch.ProjectId)
	assert.Equal(t, batch.Denom, res.Batch.Denom)
	assert.Equal(t, batch.Metadata, res.Batch.Metadata)
	assert.Equal(t, issuer.String(), res.Batch.Issuer)
	assert.DeepEqual(t, regentypes.ProtobufToGogoTimestamp(batch.StartDate), res.Batch.StartDate)
	assert.DeepEqual(t, regentypes.ProtobufToGogoTimestamp(batch.EndDate), res.Batch.EndDate)
	assert.DeepEqual(t, regentypes.ProtobufToGogoTimestamp(batch.IssuanceDate), res.Batch.IssuanceDate)

	// query batch by unknown batch denom
	_, err = s.k.Batch(s.ctx, &types.QueryBatchRequest{BatchDenom: "A00-000-00000000-00000000-000"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
