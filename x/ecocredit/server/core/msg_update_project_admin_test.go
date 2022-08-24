//nolint:revive,stylecheck
package core

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type updateProjectAdmin struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	res   *core.MsgUpdateProjectAdminResponse
	err   error
}

func TestUpdateProjectAdmin(t *testing.T) {
	gocuke.NewRunner(t, &updateProjectAdmin{}).Path("./features/msg_update_project_admin.feature").Run()
}

func (s *updateProjectAdmin) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *updateProjectAdmin) ACreditTypeWithAbbreviation(a string) {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectAdmin) ACreditClassWithClassId(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassID(a)

	err := s.k.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               a,
		Admin:            s.alice,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectAdmin) AProjectWithProjectIdAndAdminAlice(a string) {
	err := s.k.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:    a,
		Admin: s.alice,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectAdmin) AliceAttemptsToUpdateProjectAdminWithProjectId(a string) {
	s.res, s.err = s.k.UpdateProjectAdmin(s.ctx, &core.MsgUpdateProjectAdmin{
		Admin:     s.alice.String(),
		ProjectId: a,
		NewAdmin:  s.bob.String(),
	})
}

func (s *updateProjectAdmin) BobAttemptsToUpdateProjectAdminWithProjectId(a string) {
	s.res, s.err = s.k.UpdateProjectAdmin(s.ctx, &core.MsgUpdateProjectAdmin{
		Admin:     s.bob.String(),
		ProjectId: a,
		NewAdmin:  s.bob.String(),
	})
}

func (s *updateProjectAdmin) AliceAttemptsToUpdateProjectAdminWithProjectIdAndNewAdminBob(a string) {
	s.res, s.err = s.k.UpdateProjectAdmin(s.ctx, &core.MsgUpdateProjectAdmin{
		Admin:     s.alice.String(),
		ProjectId: a,
		NewAdmin:  s.bob.String(),
	})
}

func (s *updateProjectAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateProjectAdmin) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateProjectAdmin) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *updateProjectAdmin) ExpectProjectWithProjectIdAndAdminBob(a string) {
	project, err := s.stateStore.ProjectTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	require.Equal(s.t, s.bob.Bytes(), project.Admin)
}

func (s *updateProjectAdmin) ExpectEventWithProperties(a gocuke.DocString) {
	var event core.EventUpdateProjectAdmin
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}
