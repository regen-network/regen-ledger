package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestQuery_Projects_By_Admin(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, admin2 := testdata.KeyTestPubAddr()

	// insert class
	classKey, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id: "C01",
	})
	assert.NilError(t, err)

	// create two projects
	project1 := &api.Project{
		Id:           "C01-001",
		ClassKey:     classKey,
		Admin:        s.addr,
		Jurisdiction: "US-CA",
		Metadata:     "data",
	}

	err = s.stateStore.ProjectTable().Insert(s.ctx, project1)
	assert.NilError(t, err)

	project := &api.Project{
		Id:           "C01-002",
		ClassKey:     classKey,
		Admin:        s.addr,
		Jurisdiction: "US-CA",
		Metadata:     "data",
	}

	err = s.stateStore.ProjectTable().Insert(s.ctx, project)
	assert.NilError(t, err)

	// create project with different admin
	project = &api.Project{
		Id:           "C01-003",
		ClassKey:     classKey,
		Admin:        admin2,
		Jurisdiction: "US-CA",
		Metadata:     "data",
	}

	err = s.stateStore.ProjectTable().Insert(s.ctx, project)
	assert.NilError(t, err)

	// query project by admin1 expect 2 projects
	res, err := s.k.ProjectsByAdmin(s.ctx, &types.QueryProjectsByAdminRequest{Admin: s.addr.String()})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Projects), 2)

	// query project by admin1 with page limit 1 expect 1 project
	res, err = s.k.ProjectsByAdmin(s.ctx, &types.QueryProjectsByAdminRequest{Admin: s.addr.String(),
		Pagination: &query.PageRequest{
			Limit:      1,
			CountTotal: true,
		}})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Projects), 1)
	assert.Equal(t, project1.Id, res.Projects[0].Id)
	assert.Equal(t, "C01", res.Projects[0].ClassId)
	assert.Equal(t, sdk.AccAddress(project1.Admin).String(), res.Projects[0].Admin)
	assert.Equal(t, project1.Jurisdiction, res.Projects[0].Jurisdiction)
	assert.Equal(t, res.Pagination.Total, uint64(2))

	// query project by admin2 expect 1 project
	res, err = s.k.ProjectsByAdmin(s.ctx, &types.QueryProjectsByAdminRequest{Admin: admin2.String()})
	assert.NilError(t, err)
	assert.Equal(t, len(res.Projects), 1)

	// query project by invalid bech32 address
	_, err = s.k.ProjectsByAdmin(s.ctx, &types.QueryProjectsByAdminRequest{Admin: "address"})
	require.Error(t, err)
	require.Equal(t, "decoding bech32 failed: invalid bech32 string length 7: invalid argument", err.Error())
}
