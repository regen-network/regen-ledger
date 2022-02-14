package server

import (
	"context"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/v2/x/bond"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Bonds Bonds queries all bonds.
func (s serverImpl) Bonds(goCtx context.Context, request *bond.QueryBondsRequest) (*bond.QueryBondsResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	bondsIter, err := s.bondInfoTable.PrefixScan(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	var bonds []*bond.BondInfo
	pageResp, err := orm.Paginate(bondsIter, request.Pagination, &bonds)
	if err != nil {
		return nil, err
	}

	return &bond.QueryBondsResponse{
		Bonds:      bonds,
		Pagination: pageResp,
	}, nil
}

func (s serverImpl) Params(goCtx context.Context, request *bond.QueryParamsRequest) (*bond.QueryParamsResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx).Context
	var params bond.Params
	s.paramSpace.GetParamSet(ctx, &params)
	return &bond.QueryParamsResponse{}, nil
}
