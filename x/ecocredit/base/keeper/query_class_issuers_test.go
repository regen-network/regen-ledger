package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_ClassIssuers(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// make a class with 3 issuers.
	issuers := genAddrs(2)
	issuers = append(issuers, s.addr)
	err := s.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               "C01",
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: 1,
		Issuer:   s.addr,
	}))
	assert.NilError(t, s.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: 1,
		Issuer:   issuers[0],
	}))
	assert.NilError(t, s.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: 1,
		Issuer:   issuers[1],
	}))

	// base request
	res, err := s.k.ClassIssuers(s.ctx, &types.QueryClassIssuersRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, len(issuers), len(res.Issuers))

	// bad request
	_, err = s.k.ClassIssuers(s.ctx, &types.QueryClassIssuersRequest{ClassId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// paginated request
	res, err = s.k.ClassIssuers(s.ctx, &types.QueryClassIssuersRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Issuers))
	assert.Equal(t, uint64(3), res.Pagination.Total)
}

func genAddrs(x int) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, x)
	for i := 0; i < x; i++ {
		_, _, addr := testdata.KeyTestPubAddr()
		addrs[i] = addr
	}
	return addrs
}
