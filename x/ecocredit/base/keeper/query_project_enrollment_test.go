package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_ProjectEnrollment(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert project
	projId := "P001"
	projKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: projId,
	})

	// insert class
	clsId := "C01"
	clsKey, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id: clsId,
	})
	require.NoError(t, err)

	// insert project enrollment
	appMetadata := "foobar"
	enrollment := &api.ProjectEnrollment{
		ProjectKey:          projKey,
		ClassKey:            clsKey,
		Status:              0,
		ApplicationMetadata: appMetadata,
		EnrollmentMetadata:  "",
	}
	require.NoError(t, s.stateStore.ProjectEnrollmentTable().Insert(s.ctx, enrollment))

	// query project enrollment by project and class id
	res, err := s.k.ProjectEnrollment(s.ctx, &types.QueryProjectEnrollmentRequest{
		ProjectId: projId,
		ClassId:   clsId,
	})
	require.NoError(t, err)
	require.Equal(t, projId, res.Enrollment.ProjectId)
	require.Equal(t, clsId, res.Enrollment.ClassId)
	require.Equal(t, 0, int(res.Enrollment.Status))
	require.Equal(t, appMetadata, res.Enrollment.ApplicationMetadata)
	require.Equal(t, "", res.Enrollment.EnrollmentMetadata)
}
