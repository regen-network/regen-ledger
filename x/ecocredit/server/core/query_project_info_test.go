package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_ProjectInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert 1 project
	err := s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1.ProjectInfo{
		Name:            "P01",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)

	// valid query
	res, err := s.k.ProjectInfo(s.ctx, &v1.QueryProjectInfoRequest{ProjectId: "P01"})
	assert.NilError(t, err)
	assert.Equal(t, "P01", res.Info.Name)

	// invalid query
	_, err = s.k.ProjectInfo(s.ctx, &v1.QueryProjectInfoRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
