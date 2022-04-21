package core

import (
	"testing"

	"github.com/regen-network/regen-ledger/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQueryBatchesByIssuer(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	_, _, otherAddr := testdata.KeyTestPubAddr()
	_, _, noBatches := testdata.KeyTestPubAddr()

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

	// insert project
	assert.NilError(t, s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id: "P01",
	}))

	// insert two batches with s.addr as the issuer
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, batch1))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		Issuer:     s.addr,
		ProjectKey: 1,
		BatchDenom: "C01-20200101-20200102-002",
	}))

	// insert one batch without s.addr as the issuer
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		Issuer:     otherAddr,
		ProjectKey: 1,
		BatchDenom: "C01-20200101-20200102-003",
	}))

	res, err := s.k.BatchesByIssuer(s.ctx, &core.QueryBatchesByIssuerRequest{
		Issuer:     s.addr.String(),
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assertBatchEqual(t, s.ctx, s.k, res.Batches[0], batch1)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	res, err = s.k.BatchesByIssuer(s.ctx, &core.QueryBatchesByIssuerRequest{Issuer: noBatches.String()})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Batches))

	_, err = s.k.BatchesByIssuer(s.ctx, &core.QueryBatchesByIssuerRequest{Issuer: "foobar"})
	assert.ErrorContains(t, err, sdkerrors.ErrInvalidAddress.Error())
}
