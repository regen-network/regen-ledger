package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ProjectsByReferenceId queries projects by reference id.
func (k Keeper) ProjectsByReferenceId(ctx context.Context, req *core.QueryProjectsByReferenceIdRequest) (*core.QueryProjectsByReferenceIdResponse, error) {
	if req.ReferenceId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "reference-id is empty")
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.ProjectTable().List(ctx, api.ProjectReferenceIdIndexKey{}.WithReferenceId(req.ReferenceId), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	projects := make([]*core.ProjectInfo, 0)
	for it.Next() {
		project, err := it.Value()
		if err != nil {
			return nil, err
		}

		class, err := k.stateStore.ClassTable().Get(ctx, project.ClassKey)
		if err != nil {
			return nil, err
		}

		info := &core.ProjectInfo{
			Id:           project.Id,
			Admin:        sdk.AccAddress(project.Admin).String(),
			ClassId:      class.Id,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
			ReferenceId:  project.ReferenceId,
		}

		projects = append(projects, info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &core.QueryProjectsByReferenceIdResponse{
		Projects:   projects,
		Pagination: pr,
	}, nil
}
