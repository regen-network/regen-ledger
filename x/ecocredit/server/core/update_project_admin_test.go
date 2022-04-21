package core

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestUpdateProjectAdmin_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	projectId := "VERRA1"
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:                  projectId,
		Admin:               s.addr,
		ClassKey:            1,
		ProjectJurisdiction: "US-NY",
		Metadata:            "hi",
	}))
	newAdmin := sdk.AccAddress("addr1")

	_, err := s.k.UpdateProjectAdmin(s.ctx, &core.MsgUpdateProjectAdmin{
		Admin:     s.addr.String(),
		NewAdmin:  newAdmin.String(),
		ProjectId: projectId,
	})
	assert.NilError(t, err)

	project, err := s.stateStore.ProjectTable().GetById(s.ctx, projectId)
	assert.NilError(t, err)
	assert.Check(t, sdk.AccAddress(project.Admin).Equals(newAdmin))
}

func TestUpdateProjectAdmin_Unauthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	projectId := "VERRA1"
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:                  projectId,
		Admin:               s.addr,
		ClassKey:            1,
		ProjectJurisdiction: "US-NY",
		Metadata:            "hi",
	}))
	notAdmin := sdk.AccAddress("addr1c")

	_, err := s.k.UpdateProjectAdmin(s.ctx, &core.MsgUpdateProjectAdmin{
		Admin:     notAdmin.String(),
		NewAdmin:  notAdmin.String(),
		ProjectId: projectId,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())
}
