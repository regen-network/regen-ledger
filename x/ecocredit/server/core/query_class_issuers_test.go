package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_ClassIssuers(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// make a class with 3 issuers.
	issuers := genAddrs(2)
	issuers = append(issuers, s.addr)
	err := s.stateStore.ClassTable().Insert(s.ctx, &ecocreditv1.Class{
		Id:               "C01",
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.ClassIssuerTable().Insert(s.ctx, &ecocreditv1.ClassIssuer{
		ClassKey: 1,
		Issuer:   s.addr,
	}))
	assert.NilError(t, s.stateStore.ClassIssuerTable().Insert(s.ctx, &ecocreditv1.ClassIssuer{
		ClassKey: 1,
		Issuer:   issuers[0],
	}))
	assert.NilError(t, s.stateStore.ClassIssuerTable().Insert(s.ctx, &ecocreditv1.ClassIssuer{
		ClassKey: 1,
		Issuer:   issuers[1],
	}))

	// base request
	res, err := s.k.ClassIssuers(s.ctx, &core.QueryClassIssuersRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, len(issuers), len(res.Issuers))

	// bad request
	_, err = s.k.ClassIssuers(s.ctx, &core.QueryClassIssuersRequest{ClassId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// paginated request
	res, err = s.k.ClassIssuers(s.ctx, &core.QueryClassIssuersRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Issuers))
	assert.Equal(t, uint64(3), res.Pagination.Total)
}

func genAddrs(x int) []types.AccAddress {
	addrs := make([]types.AccAddress, x)
	for i := 0; i < x; i++ {
		_, _, addr := testdata.KeyTestPubAddr()
		addrs[i] = addr
	}
	return addrs
}