package keeper

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type createOrUpdateApplicationSuite struct {
	*baseSuite
	addrs map[string]sdk.AccAddress
	err   error
}

func TestCreateOrUpdateApplication(t *testing.T) {
	gocuke.NewRunner(t, &createOrUpdateApplicationSuite{}).
		Path("./features/msg_create_or_update_application.feature").
		Run()
}

func (s *createOrUpdateApplicationSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.addrs = make(map[string]sdk.AccAddress)
	s.addrs["Alice"] = s.addr
	s.addrs["Bob"] = s.addr2
}

func (s *createOrUpdateApplicationSuite) ProjectWithAdmin(projId, admin string) {
	require.NoError(s.t, s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:    projId,
		Admin: s.addrs[admin],
	}))
}

func (s *createOrUpdateApplicationSuite) CreditClass(clsId string) {
	require.NoError(s.t, s.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id: clsId,
	}))
}

func (s *createOrUpdateApplicationSuite) AnApplicationForToWithMetadata(projId, clsId, metadata string) {
	proj, err := s.stateStore.ProjectTable().GetById(s.ctx, projId)
	require.NoError(s.t, err)
	cls, err := s.stateStore.ClassTable().GetById(s.ctx, clsId)
	require.NoError(s.t, err)
	require.NoError(s.t, s.stateStore.ProjectEnrollmentTable().Insert(s.ctx, &api.ProjectEnrollment{
		ProjectKey:          proj.Key,
		ClassKey:            cls.Key,
		ApplicationMetadata: metadata,
	}))
}

func (s *createOrUpdateApplicationSuite) TheApplicationForToIsAccepted(projId, clsId string) {
	enrollment, err := s.getEnrollment(projId, clsId)
	require.NoError(s.t, err)
	enrollment.Status = api.ProjectEnrollmentStatus_PROJECT_ENROLLMENT_STATUS_ACCEPTED
	require.NoError(s.t, s.stateStore.ProjectEnrollmentTable().Save(s.ctx, enrollment))
}

func (s *createOrUpdateApplicationSuite) AnApplicationForToDoesNotExist(projId, clsId string) {
	enrollment, err := s.getEnrollment(projId, clsId)
	if !ormerrors.IsNotFound(err) {
		s.t.Fatalf("expected project enrollment not found,  got %v", enrollment)
	}
}

func (s *createOrUpdateApplicationSuite) CreatesOrUpdatesAnApplicationForToWithMetadata(admin, projId, clsId, metadata string) {
	_, s.err = s.k.CreateOrUpdateApplication(s.ctx, &v1.MsgCreateOrUpdateApplication{
		ProjectAdmin: s.addrs[admin].String(),
		ProjectId:    projId,
		ClassId:      clsId,
		Metadata:     metadata,
		Withdraw:     false,
	})
}

func (s *createOrUpdateApplicationSuite) AttemptsToWithdrawTheApplicationForToWithMetadata(admin, projId, clsId, metadata string) {
	_, s.err = s.k.CreateOrUpdateApplication(s.ctx, &v1.MsgCreateOrUpdateApplication{
		ProjectAdmin: s.addrs[admin].String(),
		ProjectId:    projId,
		ClassId:      clsId,
		Metadata:     metadata,
		Withdraw:     true,
	})
}

func (s *createOrUpdateApplicationSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createOrUpdateApplicationSuite) ExpectErrorContains(x string) {
	if x == "" {
		require.NoError(s.t, s.err)
	} else {
		require.ErrorContains(s.t, s.err, x)
	}
}

func (s *createOrUpdateApplicationSuite) ExpectTheApplicationForToExistsWithMetadata(projId, clsId, metadata string) {
	enrollment, err := s.getEnrollment(projId, clsId)
	require.NoError(s.t, err)
	require.Equal(s.t, metadata, enrollment.ApplicationMetadata)
}

func (s *createOrUpdateApplicationSuite) ExpectEventupdateapplicationWithProperties(a gocuke.DocString) {
	var evtExpected v1.EventUpdateApplication
	err := jsonpb.UnmarshalString(a.Content, &evtExpected)
	require.NoError(s.t, err)

	evtActual, found := testutil.GetEvent(&evtExpected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&evtExpected, evtActual)
	require.NoError(s.t, err)
}

func (s *createOrUpdateApplicationSuite) getEnrollment(projId, clsId string) (*api.ProjectEnrollment, error) {
	proj, err := s.stateStore.ProjectTable().GetById(s.ctx, projId)
	require.NoError(s.t, err)
	cls, err := s.stateStore.ClassTable().GetById(s.ctx, clsId)
	require.NoError(s.t, err)
	return s.stateStore.ProjectEnrollmentTable().Get(s.ctx, proj.Key, cls.Key)
}
