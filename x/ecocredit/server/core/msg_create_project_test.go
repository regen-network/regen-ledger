package core

import (
	"strconv"
	"strings"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type createProjectSuite struct {
	*baseSuite
	alice sdk.AccAddress
	res   *core.MsgCreateProjectResponse
	err   error
}

func TestCreateProject(t *testing.T) {
	gocuke.NewRunner(t, &createProjectSuite{}).Path("./features/msg_create_project.feature").Run()
}

func (s *createProjectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
}

func (s *createProjectSuite) ACreditTypeWithAbbreviation(a string) {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) ACreditClassWithClassIdAndIssuerAlice(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(a)

	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               a,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cKey,
		Issuer:   s.alice,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AProjectSequenceForCreditClass(a, b string) {
	class, err := s.k.stateStore.ClassTable().GetById(s.ctx, b)
	require.NoError(s.t, err)

	nextSequence, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	err = s.k.stateStore.ProjectSequenceTable().Insert(s.ctx, &api.ProjectSequence{
		ClassKey:     class.Key,
		NextSequence: nextSequence,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AProjectWithProjectId(a string) {
	classId := core.GetClassIdFromProjectId(a)

	class, err := s.k.stateStore.ClassTable().GetById(s.ctx, classId)
	require.NoError(s.t, err)

	err = s.k.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:       a,
		ClassKey: class.Key,
	})
	require.NoError(s.t, err)

	seq := s.getProjectSequence(a)

	// Save because project sequence may already exist
	err = s.k.stateStore.ProjectSequenceTable().Save(s.ctx, &api.ProjectSequence{
		ClassKey:     class.Key,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AProjectWithProjectIdAndReferenceId(a, b string) {
	classId := core.GetClassIdFromProjectId(a)

	class, err := s.k.stateStore.ClassTable().GetById(s.ctx, classId)
	require.NoError(s.t, err)

	err = s.k.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:          a,
		ClassKey:    class.Key,
		ReferenceId: b,
	})
	require.NoError(s.t, err)

	seq := s.getProjectSequence(a)

	// Save because project sequence may already exist
	err = s.k.stateStore.ProjectSequenceTable().Save(s.ctx, &api.ProjectSequence{
		ClassKey:     class.Key,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AliceAttemptsToCreateAProjectWithClassId(a string) {
	s.res, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Admin:   s.alice.String(),
		ClassId: a,
	})
}

func (s *createProjectSuite) AliceAttemptsToCreateAProjectWithClassIdAndReferenceId(a, b string) {
	s.res, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Admin:       s.alice.String(),
		ClassId:     a,
		ReferenceId: b,
	})
}

func (s *createProjectSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createProjectSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createProjectSuite) ExpectProjectWithProjectId(a string) {
	project, err := s.k.stateStore.ProjectTable().GetById(s.ctx, a)
	require.NoError(s.t, err)
	require.Equal(s.t, a, project.Id)
}

func (s *createProjectSuite) getProjectSequence(projectId string) uint64 {
	str := strings.Split(projectId, "-")
	seq, err := strconv.ParseUint(str[1], 10, 32)
	require.NoError(s.t, err)
	return seq
}
