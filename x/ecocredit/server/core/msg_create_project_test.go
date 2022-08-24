//nolint:revive,stylecheck
package core

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type createProjectSuite struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	res   *core.MsgCreateProjectResponse
	err   error
}

func TestCreateProject(t *testing.T) {
	gocuke.NewRunner(t, &createProjectSuite{}).Path("./features/msg_create_project.feature").Run()
}

func (s *createProjectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *createProjectSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) ACreditClassWithClassIdAndIssuerAlice(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassID(a)

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

func (s *createProjectSuite) AProjectSequenceWithClassIdAndNextSequence(a, b string) {
	class, err := s.k.stateStore.ClassTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	nextSequence, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	err = s.k.stateStore.ProjectSequenceTable().Insert(s.ctx, &api.ProjectSequence{
		ClassKey:     class.Key,
		NextSequence: nextSequence,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AProjectWithProjectIdAndReferenceId(a, b string) {
	classID := core.GetClassIDFromProjectID(a)

	class, err := s.k.stateStore.ClassTable().GetById(s.ctx, classID)
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

func (s *createProjectSuite) BobAttemptsToCreateAProjectWithClassId(a string) {
	s.res, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Admin:   s.bob.String(),
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

func (s *createProjectSuite) AliceAttemptsToCreateAProjectWithProperties(a gocuke.DocString) {
	var msg core.MsgCreateProject
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateProject(s.ctx, &core.MsgCreateProject{
		Admin:        s.alice.String(),
		ClassId:      msg.ClassId,
		Metadata:     msg.Metadata,
		Jurisdiction: msg.Jurisdiction,
		ReferenceId:  msg.ReferenceId,
	})
}

func (s *createProjectSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createProjectSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createProjectSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *createProjectSuite) ExpectProjectSequenceWithClassIdAndNextSequence(a string, b string) {
	project, err := s.stateStore.ClassTable().GetById(s.ctx, a)
	require.NoError(s.t, err)

	nextSequence, err := strconv.ParseUint(b, 10, 64)
	require.NoError(s.t, err)

	projectSequence, err := s.stateStore.ProjectSequenceTable().Get(s.ctx, project.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, nextSequence, projectSequence.NextSequence)
}

func (s *createProjectSuite) ExpectProjectProperties(a gocuke.DocString) {
	var expected core.Project
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.ProjectTable().GetById(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Metadata, batch.Metadata)
	require.Equal(s.t, expected.Jurisdiction, batch.Jurisdiction)
	require.Equal(s.t, expected.ReferenceId, batch.ReferenceId)
}

func (s *createProjectSuite) ExpectTheResponse(a gocuke.DocString) {
	var res core.MsgCreateProjectResponse
	err := jsonpb.UnmarshalString(a.Content, &res)
	require.NoError(s.t, err)

	require.Equal(s.t, &res, s.res)
}

func (s *createProjectSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event core.EventCreateProject
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *createProjectSuite) getProjectSequence(projectID string) uint64 {
	str := strings.Split(projectID, "-")
	seq, err := strconv.ParseUint(str[1], 10, 32)
	require.NoError(s.t, err)
	return seq
}
