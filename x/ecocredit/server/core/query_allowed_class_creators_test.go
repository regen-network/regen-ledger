package core

import (
	"testing"

	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_AllowedClassCreators(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	res, err := s.k.AllowedClassCreators(s.ctx, &core.QueryAllowedClassCreatorsRequest{})
	assert.NilError(t, err)
	assert.Equal(t, len(res.ClassCreators), 0)

	// add one class creator
	err = s.stateStore.AllowedClassCreatorTable().Insert(s.ctx, &ecocreditv1.AllowedClassCreator{
		Address: sdk.AccAddress("creator1"),
	})
	assert.NilError(t, err)

	res, err = s.k.AllowedClassCreators(s.ctx, &core.QueryAllowedClassCreatorsRequest{})
	assert.NilError(t, err)
	assert.Equal(t, len(res.ClassCreators), 1)

	// add another class creator
	err = s.stateStore.AllowedClassCreatorTable().Insert(s.ctx, &ecocreditv1.AllowedClassCreator{
		Address: sdk.AccAddress("creator2"),
	})
	assert.NilError(t, err)

	res, err = s.k.AllowedClassCreators(s.ctx, &core.QueryAllowedClassCreatorsRequest{})
	assert.NilError(t, err)
	assert.Equal(t, len(res.ClassCreators), 2)

	// test pagination
	res, err = s.k.AllowedClassCreators(s.ctx, &core.QueryAllowedClassCreatorsRequest{
		Pagination: &query.PageRequest{
			Limit:      1,
			CountTotal: true,
		},
	})
	assert.NilError(t, err)
	assert.Equal(t, len(res.ClassCreators), 1)
	assert.Equal(t, res.Pagination.Total, uint64(2))
}