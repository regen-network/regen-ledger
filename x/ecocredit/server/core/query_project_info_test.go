package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_ProjectInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert 1 project
	err := s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id:                  "P01",
		ClassKey:            1,
		ProjectJurisdiction: "US-CA",
		Metadata:            "",
	})
	assert.NilError(t, err)

	// valid query
	res, err := s.k.ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: "P01"})
	assert.NilError(t, err)
	assert.Equal(t, "P01", res.Info.Id)

	// invalid query
	_, err = s.k.ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
