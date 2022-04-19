package core

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestUpdateProjectMetadata_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	projectId := "VERRA1"
	assert.NilError(t, s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id:              projectId,
		Admin:           s.addr,
		ClassKey:        1,
		ProjectLocation: "US-NY",
		Metadata:        "hi",
	}))
	newMetadata := "hello world"

	_, err := s.k.UpdateProjectMetadata(s.ctx, &core.MsgUpdateProjectMetadata{
		Admin:       s.addr.String(),
		NewMetadata: newMetadata,
		ProjectId:   projectId,
	})
	assert.NilError(t, err)

	project, err := s.stateStore.ProjectInfoTable().GetById(s.ctx, projectId)
	assert.NilError(t, err)
	assert.Equal(t, newMetadata, project.Metadata)
}

func TestUpdateProjectMetadata_Unauthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	projectId := "VERRA1"
	assert.NilError(t, s.stateStore.ProjectInfoTable().Insert(s.ctx, &api.ProjectInfo{
		Id:              projectId,
		Admin:           s.addr,
		ClassKey:        1,
		ProjectLocation: "US-NY",
		Metadata:        "hi",
	}))
	notAdmin := sdk.AccAddress("addr1")

	_, err := s.k.UpdateProjectMetadata(s.ctx, &core.MsgUpdateProjectMetadata{
		Admin:       notAdmin.String(),
		NewMetadata: "yay",
		ProjectId:   projectId,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())
}
