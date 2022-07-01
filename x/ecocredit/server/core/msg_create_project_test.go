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
	alice            sdk.AccAddress
	creditTypeAbbrev string
	classId          string
	classKey         uint64
	projectId        string
	res              *core.MsgCreateProjectResponse
	err              error
}

func TestCreateProject(t *testing.T) {
	gocuke.NewRunner(t, &createProjectSuite{}).Path("./features/msg_create_project.feature").Run()
}

func (s *createProjectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
	s.projectId = "C01-001"
}

func (s *createProjectSuite) ACreditType() {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Name:         s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) ACreditClassWithIssuerAlice() {
	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               s.classId,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cKey,
		Issuer:   s.alice,
	})
	require.NoError(s.t, err)

	s.classKey = cKey
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

	err = s.k.stateStore.ProjectSequenceTable().Save(s.ctx, &api.ProjectSequence{
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

	err = s.k.stateStore.ProjectSequenceTable().Save(s.ctx, &api.ProjectSequence{
		ClassKey:     class.Key,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AProjectWithReferenceId(a string) {
	err := s.k.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:          s.projectId,
		ClassKey:    s.classKey,
		ReferenceId: a,
	})
	require.NoError(s.t, err)

	seq := s.getProjectSequence(s.projectId)

	err = s.k.stateStore.ProjectSequenceTable().Save(s.ctx, &api.ProjectSequence{
		ClassKey:     s.classKey,
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

func (s *createProjectSuite) AliceAttemptsToCreateAProjectWithReferenceId(a string) {
	s.res, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Admin:       s.alice.String(),
		ClassId:     s.classId,
		ReferenceId: a,
	})
}

func (s *createProjectSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createProjectSuite) ExpectProjectWithProjectId(a string) {
	project, err := s.k.stateStore.ProjectTable().GetById(s.ctx, a)
	require.NoError(s.t, err)
	require.Equal(s.t, a, project.Id)
}

func (s *createProjectSuite) getProjectSequence(a string) uint64 {
	str := strings.Split(a, "-")
	seq, err := strconv.ParseUint(str[1], 10, 32)
	require.NoError(s.t, err)
	return seq
}
