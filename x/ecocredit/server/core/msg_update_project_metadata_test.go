package core

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type updateProjectMetadata struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	res   *core.MsgUpdateProjectMetadataResponse
	err   error
}

func TestUpdateProjectMetadata(t *testing.T) {
	gocuke.NewRunner(t, &updateProjectMetadata{}).Path("./features/msg_update_project_metadata.feature").Run()
}

func (s *updateProjectMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *updateProjectMetadata) ACreditTypeWithAbbreviation(a string) {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectMetadata) ACreditClassWithClassId(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(a)

	err := s.k.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               a,
		Admin:            s.alice,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectMetadata) AProjectWithProjectIdAndAdminAlice(a string) {
	err := s.k.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:    a,
		Admin: s.alice,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectMetadata) AliceAttemptsToUpdateProjectMetadataWithProjectId(a string) {
	s.res, s.err = s.k.UpdateProjectMetadata(s.ctx, &core.MsgUpdateProjectMetadata{
		Admin:     s.alice.String(),
		ProjectId: a,
	})
}

func (s *updateProjectMetadata) BobAttemptsToUpdateProjectMetadataWithProjectId(a string) {
	s.res, s.err = s.k.UpdateProjectMetadata(s.ctx, &core.MsgUpdateProjectMetadata{
		Admin:     s.bob.String(),
		ProjectId: a,
	})
}

func (s *updateProjectMetadata) AliceAttemptsToUpdateProjectMetadataWithProjectIdAndNewMetadata(a string, b gocuke.DocString) {
	s.res, s.err = s.k.UpdateProjectMetadata(s.ctx, &core.MsgUpdateProjectMetadata{
		Admin:       s.alice.String(),
		ProjectId:   a,
		NewMetadata: b.Content,
	})
}

func (s *updateProjectMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateProjectMetadata) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateProjectMetadata) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *updateProjectMetadata) ExpectProjectWithProjectIdAndMetadata(a string, b gocuke.DocString) {
	project, err := s.stateStore.ProjectTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	require.Equal(s.t, b.Content, project.Metadata)
}
