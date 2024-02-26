//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type createProjectSuite struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	res   *types.MsgCreateProjectResponse
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
	creditTypeAbbrev := base.GetCreditTypeAbbrevFromClassID(a)

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

func (s *createProjectSuite) AProjectWithProjectIdAndReferenceId(a, b string) {
	err := s.k.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:          a,
		ReferenceId: b,
	})
	require.NoError(s.t, err)
}

func (s *createProjectSuite) AliceAttemptsToCreateAProjectWithClassId(a string) {
	s.res, s.err = s.k.CreateProject(s.ctx, &types.MsgCreateProject{
		Admin:   s.alice.String(),
		ClassId: a,
	})
}

func (s *createProjectSuite) BobAttemptsToCreateAProjectWithClassId(a string) {
	s.res, s.err = s.k.CreateProject(s.ctx, &types.MsgCreateProject{
		Admin:   s.bob.String(),
		ClassId: a,
	})
}

func (s *createProjectSuite) AliceAttemptsToCreateAProjectWithClassIdAndReferenceId(a, b string) {
	s.res, s.err = s.k.CreateProject(s.ctx, &types.MsgCreateProject{
		Admin:       s.alice.String(),
		ClassId:     a,
		ReferenceId: b,
	})
}

func (s *createProjectSuite) AliceAttemptsToCreateAProjectWithProperties(a gocuke.DocString) {
	var msg types.MsgCreateProject
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateProject(s.ctx, &types.MsgCreateProject{
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

func (s *createProjectSuite) ExpectProjectProperties(a gocuke.DocString) {
	var expected types.Project
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.ProjectTable().GetById(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Metadata, batch.Metadata)
	require.Equal(s.t, expected.Jurisdiction, batch.Jurisdiction)
	require.Equal(s.t, expected.ReferenceId, batch.ReferenceId)
}

func (s *createProjectSuite) ExpectTheResponse(a gocuke.DocString) {
	var res types.MsgCreateProjectResponse
	err := jsonpb.UnmarshalString(a.Content, &res)
	require.NoError(s.t, err)

	require.Equal(s.t, &res, s.res)
}

func (s *createProjectSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventCreateProject
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}
