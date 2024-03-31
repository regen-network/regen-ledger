package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestQuery_ProjectEnrollments(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// insert projects
	projId1 := "P001"
	projKey1, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: projId1,
	})
	require.NoError(t, err)

	projId2 := "P002"
	projKey2, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: projId2,
	})
	require.NoError(t, err)

	// insert classes
	clsId1 := "C01"
	clsKey1, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id: clsId1,
	})

	clsId2 := "BIO01"
	clsKey2, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id: clsId2,
	})

	// insert enrollments
	enrollment1 := &api.ProjectEnrollment{
		ProjectKey: projKey1,
		ClassKey:   clsKey1,
	}

	enrollment2 := &api.ProjectEnrollment{
		ProjectKey: projKey1,
		ClassKey:   clsKey2,
	}

	enrollment3 := &api.ProjectEnrollment{
		ProjectKey: projKey2,
		ClassKey:   clsKey2,
	}

	for _, enrollment := range []*api.ProjectEnrollment{enrollment1, enrollment2, enrollment3} {
		require.NoError(t, s.stateStore.ProjectEnrollmentTable().Insert(s.ctx, enrollment))
	}

	// query project enrollments by project id
	res, err := s.k.ProjectEnrollments(s.ctx, &types.QueryProjectEnrollmentsRequest{
		ProjectId: projId1,
	})
	require.NoError(t, err)
	require.Len(t, res.Enrollments, 2)
	require.Equal(t, projId1, res.Enrollments[0].ProjectId)
	require.Equal(t, clsId1, res.Enrollments[0].ClassId)
	require.Equal(t, projId1, res.Enrollments[1].ProjectId)
	require.Equal(t, clsId2, res.Enrollments[1].ClassId)

	// query project enrollments by class id
	res, err = s.k.ProjectEnrollments(s.ctx, &types.QueryProjectEnrollmentsRequest{
		ClassId: clsId2,
	})
	require.NoError(t, err)
	require.Len(t, res.Enrollments, 2)
	require.Equal(t, projId1, res.Enrollments[0].ProjectId)
	require.Equal(t, clsId2, res.Enrollments[0].ClassId)
	require.Equal(t, projId2, res.Enrollments[1].ProjectId)
	require.Equal(t, clsId2, res.Enrollments[1].ClassId)
}
