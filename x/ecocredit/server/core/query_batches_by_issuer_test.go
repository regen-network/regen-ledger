package core

import (
	"testing"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func TestQueryBatchesByIssuer(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, otherAddr := testdata.KeyTestPubAddr()
	_, _, noBatches := testdata.KeyTestPubAddr()
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{Issuer: s.addr, Denom: "1"}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{Issuer: s.addr, Denom: "2"}))
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{Issuer: otherAddr, Denom: "3"}))

	res, err := s.k.BatchesByIssuer(s.ctx, &core.QueryBatchesByIssuerRequest{Issuer: s.addr.String(), Pagination: &query.PageRequest{Limit: 1, CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assert.Equal(t, "1", res.Batches[0].Denom)
	assert.Equal(t, uint64(2), res.Pagination.Total)

	res, err = s.k.BatchesByIssuer(s.ctx, &core.QueryBatchesByIssuerRequest{Issuer: noBatches.String()})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.Batches))

	_, err = s.k.BatchesByIssuer(s.ctx, &core.QueryBatchesByIssuerRequest{Issuer: "foobar"})
	assert.ErrorContains(t, err, sdkerrors.ErrInvalidAddress.Error())
}
