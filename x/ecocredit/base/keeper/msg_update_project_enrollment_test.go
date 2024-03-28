package keeper

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
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
		"Bob": s.addr2,
	}
}

func (s *updateProjectEnrollmentSuite) ClassWithIssuer(clsId, issuer string) {
	clsKey, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:    clsId,
		Admin: s.addrs[issuer],
	})
	require.NoError(s.t, err)

	err = s.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: clsKey,
		Issuer:   s.addrs[issuer],
	})
	require.NoError(s.t, err)
}

func (s *updateProjectEnrollmentSuite) AnApplicationForProjectToClassWithStatusAndMetadata(projId, clsId, status, metadata string) {
	projKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: projId,
	})
	require.NoError(s.t, err)
	cls, err := s.stateStore.ClassTable().GetById(s.ctx, clsId)
	require.NoError(s.t, err)
	err = s.stateStore.ProjectEnrollmentTable().Insert(s.ctx, &api.ProjectEnrollment{
		ProjectKey:         projKey,
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

func (s *updateProjectEnrollmentSuite) ExpectNoEnrollmentForTo(projId, clsId string) {
	enrollment, err := s.getEnrollment(projId, clsId)
	if !ormerrors.IsNotFound(err) {
		s.t.Fatalf("expected project enrollment not found,  got %v", enrollment)
	}
}

func (s *updateProjectEnrollmentSuite) ExpectErrorContains(a string) {
	if a == "" {
		require.Error(s.t, s.err)
	} else {
		require.ErrorContains(s.t, s.err, a)
	}
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
