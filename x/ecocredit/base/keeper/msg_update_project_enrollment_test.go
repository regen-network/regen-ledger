package keeper

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type updateProjectEnrollmentSuite struct {
	*baseSuite
	err   error
	addrs map[string]sdk.AccAddress
}

func TestUpdateProjectEnrollment(t *testing.T) {
	gocuke.NewRunner(t, &updateProjectEnrollmentSuite{}).
		Path("./features/msg_update_project_enrollment.feature").
		Run()
}

func (s *updateProjectEnrollmentSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.addrs = map[string]sdk.AccAddress{
		"I01": s.addr,
		"I02": s.addr2,
		"Bob": s.addr3,
	}
}

func (s *updateProjectEnrollmentSuite) Project(projId string) {
	err := s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id: projId,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectEnrollmentSuite) Class(clsId string) {
	err := s.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id: clsId,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectEnrollmentSuite) ClassIssuerFor(issuer, clsId string) {
	cls, err := s.stateStore.ClassTable().GetById(s.ctx, clsId)
	require.NoError(s.t, err)
	err = s.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cls.Key,
		Issuer:   s.addrs[issuer],
	})
	require.NoError(s.t, err)
}

func (s *updateProjectEnrollmentSuite) EnrollmentForToIsWithMetadata(projId, clsId, status, metadata string) {
	proj, err := s.stateStore.ProjectTable().GetById(s.ctx, projId)
	require.NoError(s.t, err)
	cls, err := s.stateStore.ClassTable().GetById(s.ctx, clsId)
	require.NoError(s.t, err)
	err = s.stateStore.ProjectEnrollmentTable().Insert(s.ctx, &api.ProjectEnrollment{
		ProjectKey:         proj.Key,
		ClassKey:           cls.Key,
		Status:             statusFromString(status),
		EnrollmentMetadata: metadata,
	})
	require.NoError(s.t, err)
}

func (s *updateProjectEnrollmentSuite) UpdatesEnrollmentForToWithStatusAndMetadata(issuer, projId, clsId, status, metadata string) {
	_, s.err = s.k.UpdateProjectEnrollment(s.ctx, &v1.MsgUpdateProjectEnrollment{
		Issuer:    s.addrs[issuer].String(),
		ProjectId: projId,
		ClassId:   clsId,
		NewStatus: v1.ProjectEnrollmentStatus(statusFromString(status)),
		Metadata:  metadata,
	})
}

func (s *updateProjectEnrollmentSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateProjectEnrollmentSuite) ExpectEnrollmentForToToBeWithMetadata(projId, clsId, status, metadata string) {
	enrollment, err := s.getEnrollment(projId, clsId)
	require.NoError(s.t, err)
	require.Equal(s.t, statusFromString(status), enrollment.Status)
	require.Equal(s.t, metadata, enrollment.EnrollmentMetadata)
}

func (s *updateProjectEnrollmentSuite) ExpectEnrollmentExistsForToToBe(projId, clsId, exists string) {
	enrollment, err := s.getEnrollment(projId, clsId)
	if exists == "true" {
		require.NoError(s.t, err)
	} else if exists == "false" {
		if !ormerrors.IsNotFound(err) {
			s.t.Fatalf("expected project enrollment not found,  got %v", enrollment)
		}
	} else {
		s.t.Fatalf("invalid exists value: %s", exists)
	}
}

func (s *updateProjectEnrollmentSuite) ExpectErrorContains(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.ErrorContains(s.t, s.err, a)
	}
}

func (s *updateProjectEnrollmentSuite) IfNoErrorExpectEventupdateprojectenrollmentWithProperties(a gocuke.DocString) {
	if s.err != nil {
		return
	}

	s.ExpectEventupdateprojectenrollmentWithProperties(a)
}

func (s *updateProjectEnrollmentSuite) ExpectEventupdateprojectenrollmentWithProperties(a gocuke.DocString) {
	var evtExpected v1.EventUpdateProjectEnrollment
	err := jsonpb.UnmarshalString(a.Content, &evtExpected)
	require.NoError(s.t, err)

	// update issuer to actual address
	evtExpected.Issuer = s.addrs[evtExpected.Issuer].String()

	evtActual, found := testutil.GetEvent(&evtExpected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&evtExpected, evtActual)
	require.NoError(s.t, err)
}

func (s *updateProjectEnrollmentSuite) getEnrollment(projId, clsId string) (*api.ProjectEnrollment, error) {
	proj, err := s.stateStore.ProjectTable().GetById(s.ctx, projId)
	require.NoError(s.t, err)
	cls, err := s.stateStore.ClassTable().GetById(s.ctx, clsId)
	require.NoError(s.t, err)
	return s.stateStore.ProjectEnrollmentTable().Get(s.ctx, proj.Key, cls.Key)
}

func statusFromString(statusName string) api.ProjectEnrollmentStatus {
	var status api.ProjectEnrollmentStatus
	return api.ProjectEnrollmentStatus(
		status.Descriptor().Values().
			ByName(protoreflect.Name(fmt.Sprintf("PROJECT_ENROLLMENT_STATUS_%s", statusName))).
			Number())
}
