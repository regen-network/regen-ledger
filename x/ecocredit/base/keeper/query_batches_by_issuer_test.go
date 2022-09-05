package keeper

import (
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestQueryBatchesByIssuer(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	startTime, err := regentypes.ParseDate("start date", "2020-01-01")
	assert.NilError(t, err)

	endTime, err := regentypes.ParseDate("end date", "2021-01-01")
	assert.NilError(t, err)

	issuanceTime, err := regentypes.ParseDate("issuance date", "2022-01-01")
	assert.NilError(t, err)

	// insert project
	pKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: "C01-001",
	})
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

	// insert two batches with s.addr as the issuer
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, batch1))
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{
		Issuer: s.addr,
		Denom:  "C01-001-20200101-20210101-002",
	}))

	// query batches by issuer s.addr
	res, err := s.k.BatchesByIssuer(s.ctx, &types.QueryBatchesByIssuerRequest{
		Issuer:     s.addr.String(),
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assert.Equal(t, uint64(2), res.Pagination.Total)
	assertBatchEqual(s.ctx, t, s.k, res.Batches[0], batch1)

	_, _, notIssuer := testdata.KeyTestPubAddr()

	// query batches by an address that is not an issuer
	res, err = s.k.BatchesByIssuer(s.ctx, &types.QueryBatchesByIssuerRequest{Issuer: notIssuer.String()})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Batches))

	// query batches by an invalid address
	_, err = s.k.BatchesByIssuer(s.ctx, &types.QueryBatchesByIssuerRequest{Issuer: "foobar"})
	assert.ErrorContains(t, err, sdkerrors.ErrInvalidAddress.Error())
}
