package core

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestCreateProject_ValidProjectState(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	makeClass(t, s.ctx, s.stateStore, s.addr)
	res, err := s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		Metadata:        "",
		ProjectLocation: "US-NY",
		ProjectId:       "FOO",
	})
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, "FOO")

	project, err := s.stateStore.ProjectInfoTable().GetByName(s.ctx, "FOO")
	assert.NilError(t, err)
	assert.Equal(t, project.ProjectLocation, "US-NY")
}

func TestCreateProject_GeneratedProjectID(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	makeClass(t, s.ctx, s.stateStore, s.addr)
	res, err := s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		Metadata:        "",
		ProjectLocation: "US-NY",
		ProjectId:       "",
	})
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, "C0101", "got project id: %s", res.ProjectId)

	res, err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		Metadata:        "",
		ProjectLocation: "US-NY",
		ProjectId:       "",
	})
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, "C0102", "got project id: %s", res.ProjectId)
}

func TestCreateProject_BadClassID(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, err := s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "NOPE",
		ProjectLocation: "US-NY",
		ProjectId:       "",
	})
	assert.ErrorContains(t, err, "not found")
}

func TestCreateProject_NoDuplicates(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	makeClass(t, s.ctx, s.stateStore, s.addr)
	_, err := s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		ProjectLocation: "US-NY",
		ProjectId:       "FOO",
	})
	assert.NilError(t, err)

	_, err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		ProjectLocation: "US-NY",
		ProjectId:       "FOO",
	})
	assert.ErrorContains(t, err, ormerrors.UniqueKeyViolation.Error())
}

func makeClass(t *testing.T, ctx context.Context, ss api.StateStore, addr types.AccAddress) {
	assert.NilError(t, ss.ClassInfoTable().Insert(ctx, &api.ClassInfo{
		Name:       "C01",
		Admin:      addr,
		Metadata:   "",
		CreditType: "C",
	}))
	assert.NilError(t, ss.ClassIssuerTable().Insert(ctx, &api.ClassIssuer{
		ClassId: 1,
		Issuer:  addr,
	}))
}
