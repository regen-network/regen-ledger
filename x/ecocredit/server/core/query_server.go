package core

import (
	"context"
	queryv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gogo/protobuf/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Classes queries for all credit classes with pagination.
func (s serverImpl) Classes(ctx context.Context, request *v1beta1.QueryClassesRequest) (*v1beta1.QueryClassesResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if request.Pagination == nil {
		request.Pagination = &query.PageRequest{}
	}
	p := request.Pagination
	it, err := s.stateStore.ClassInfoStore().List(ctx, &ecocreditv1beta1.ClassInfoPrimaryKey{}, ormlist.Paginate(&queryv1beta1.PageRequest{
		Key:        p.Key,
		Offset:     p.Offset,
		Limit:      p.Limit,
		CountTotal: p.CountTotal,
		Reverse:    p.Reverse,
	}))
	if err != nil {
		return nil, err
	}
	infos := make([]*v1beta1.ClassInfo, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}
		infos = append(infos, &v1beta1.ClassInfo{
			Id:         info.Id,
			Name:       info.Name,
			Admin:      info.Admin,
			Metadata:   info.Metadata,
			CreditType: info.CreditType,
		})
	}
	return nil, err
}

// ClassInfo queries for information on a credit class.
func (s serverImpl) ClassInfo(ctx context.Context, request *v1beta1.QueryClassInfoRequest) (*v1beta1.QueryClassInfoResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if err := ecocredit.ValidateClassID(request.ClassId); err != nil {
		return nil, err
	}
	classInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	issuers := make([]sdk.AccAddress, 0)
	it, err := s.stateStore.ClassIssuerStore().List(ctx, ecocreditv1beta1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(classInfo.Id))
	if err != nil {
		return nil, err
	}
	for it.Next() {
		val, err := it.Value()
		if err != nil {
			return nil, err
		}
		issuers = append(issuers, val.Issuer)
	}

	return &v1beta1.QueryClassInfoResponse{Info: &v1beta1.ClassInfo{
		Id:         classInfo.Id,
		Name:       request.ClassId,
		Admin:      classInfo.Admin,
		Metadata:   classInfo.Metadata,
		CreditType: classInfo.CreditType,
	}}, nil
}

func (s serverImpl) ClassIssuers(ctx context.Context, request *v1beta1.QueryClassIssuersRequest) (*v1beta1.QueryClassIssuersResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if request.Pagination == nil {
		request.Pagination = &query.PageRequest{}
	}
	p := request.Pagination
	if err := ecocredit.ValidateClassID(request.ClassId); err != nil {
		return nil, err
	}

	classInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.ClassIssuerStore().List(ctx, ecocreditv1beta1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(classInfo.Id), ormlist.Paginate(&queryv1beta1.PageRequest{
		Key:        p.Key,
		Offset:     p.Offset,
		Limit:      p.Limit,
		CountTotal: p.CountTotal,
		Reverse:    p.Reverse,
	}))
	if err != nil {
		return nil, err
	}

	issuers := make([]string, 0)
	for it.Next() {
		issuer, err := it.Value()
		if err != nil {
			return nil, err
		}

		issuers = append(issuers, string(issuer.Issuer))
	}
	pr := it.PageResponse()

	return &v1beta1.QueryClassIssuersResponse{
		Issuers: issuers,
		Pagination: &query.PageResponse{
			NextKey: pr.NextKey,
			Total:   pr.Total,
		},
	}, nil
}

// Projects queries projects of a given credit batch.
func (s serverImpl) Projects(ctx context.Context, request *v1beta1.QueryProjectsRequest) (*v1beta1.QueryProjectsResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if request.Pagination == nil {
		request.Pagination = &query.PageRequest{}
	}
	p := request.Pagination
	cInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}
	it, err := s.stateStore.ProjectInfoStore().List(ctx, ecocreditv1beta1.ProjectInfoClassIdNameIndexKey{}.WithClassId(cInfo.Id), ormlist.Paginate(&queryv1beta1.PageRequest{
		Key:        p.Key,
		Offset:     p.Offset,
		Limit:      p.Limit,
		CountTotal: p.CountTotal,
		Reverse:    p.Reverse,
	}))
	if err != nil {
		return nil, err
	}
	projectInfos := make([]*v1beta1.ProjectInfo, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}
		classInfo, err := s.stateStore.ClassInfoStore().Get(ctx, info.ClassId)
		if err != nil {
			return nil, err
		}
		projectInfos = append(projectInfos, &v1beta1.ProjectInfo{
			Id:              info.Id,
			Name:            info.Name,
			ClassId:         classInfo.Id,
			ProjectLocation: info.ProjectLocation,
			Metadata:        info.Metadata,
		})
	}
	pg := it.PageResponse()
	return &v1beta1.QueryProjectsResponse{
		Projects: projectInfos,
		Pagination: &query.PageResponse{
			NextKey: pg.NextKey,
			Total:   pg.Total,
		},
	}, nil
}

func (s serverImpl) ProjectInfo(ctx context.Context, request *v1beta1.QueryProjectInfoRequest) (*v1beta1.QueryProjectInfoResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if err := ecocredit.ValidateProjectID(request.ProjectId); err != nil {
		return nil, err
	}
	pInfo, err := s.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}

	cInfo, err := s.stateStore.ClassInfoStore().Get(ctx, pInfo.ClassId)
	if err != nil {
		return nil, err
	}

	return &v1beta1.QueryProjectInfoResponse{Info: &v1beta1.ProjectInfo{
		Id:              pInfo.Id,
		Name:            request.ProjectId,
		ClassId:         cInfo.Id,
		ProjectLocation: pInfo.ProjectLocation,
		Metadata:        pInfo.Metadata,
	}}, nil
}

// Batches queries for all batches in the given credit class.
func (s serverImpl) Batches(ctx context.Context, request *v1beta1.QueryBatchesRequest) (*v1beta1.QueryBatchesResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if request.Pagination == nil {
		request.Pagination = &query.PageRequest{}
	}
	p := request.Pagination
	project, err := s.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	it, err := s.stateStore.BatchInfoStore().List(ctx, ecocreditv1beta1.BatchInfoProjectIdIndexKey{}.WithProjectId(project.Id), ormlist.Paginate(&queryv1beta1.PageRequest{
		Key:        p.Key,
		Offset:     p.Offset,
		Limit:      p.Limit,
		CountTotal: p.CountTotal,
		Reverse:    p.Reverse,
	}))
	if err != nil {
		return nil, err
	}

	projectName := request.ProjectId
	pinfo, err := s.stateStore.ProjectInfoStore().GetByName(ctx, projectName)
	if err != nil {
		return nil, err
	}

	batches := make([]*v1beta1.BatchInfo, 0)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}

		protoStart, err := types.TimestampProto(batch.StartDate.AsTime())
		if err != nil {
			return nil, err
		}
		protoEnd, err := types.TimestampProto(batch.EndDate.AsTime())
		if err != nil {
			return nil, err
		}
		batches = append(batches, &v1beta1.BatchInfo{
			ProjectId:  pinfo.Id,
			BatchDenom: batch.BatchDenom,
			Metadata:   batch.Metadata,
			StartDate:  protoStart,
			EndDate:    protoEnd,
		})
	}
	pr := it.PageResponse()
	return &v1beta1.QueryBatchesResponse{
		Batches: batches,
		Pagination: &query.PageResponse{
			NextKey: pr.NextKey,
			Total:   pr.Total,
		},
	}, nil
}

// BatchInfo queries for information on a credit batch.
func (s serverImpl) BatchInfo(ctx context.Context, request *v1beta1.QueryBatchInfoRequest) (*v1beta1.QueryBatchInfoResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	batch, err := s.stateStore.BatchInfoStore().GetByBatchDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
	}

	project, err := s.stateStore.ProjectInfoStore().Get(ctx, batch.ProjectId)
	if err != nil {
		return nil, err
	}

	protoStart, err := types.TimestampProto(batch.StartDate.AsTime())
	if err != nil {
		return nil, err
	}
	protoEnd, err := types.TimestampProto(batch.EndDate.AsTime())
	if err != nil {
		return nil, err
	}
	return &v1beta1.QueryBatchInfoResponse{
		Info: &v1beta1.BatchInfo{
			Id:         batch.Id,
			ProjectId:  project.Id,
			BatchDenom: request.BatchDenom,
			Metadata:   batch.Metadata,
			StartDate:  protoStart,
			EndDate:    protoEnd,
		},
	}, nil
}

// Balance queries the balance (both tradable and retired) of a given credit
// batch for a given account.
func (s serverImpl) Balance(ctx context.Context, req *v1beta1.QueryBalanceRequest) (*v1beta1.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if err := ecocredit.ValidateDenom(req.BatchDenom); err != nil {
		return nil, err
	}
	batch, err := s.stateStore.BatchInfoStore().GetByBatchDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, err
	}
	if batch == nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("batch with denom %s not found", req.BatchDenom)
	}
	addr, _ := sdk.AccAddressFromBech32(req.Account)

	balance, err := s.stateStore.BatchBalanceStore().Get(ctx, addr, batch.Id)
	if err != nil {
		return nil, err
	}
	if balance == nil {
		return &v1beta1.QueryBalanceResponse{
			TradableAmount: "0",
			RetiredAmount:  "0",
		}, nil
	}
	return &v1beta1.QueryBalanceResponse{
		TradableAmount: balance.Tradable,
		RetiredAmount:  balance.Retired,
	}, nil
}

// Supply queries the supply (tradable, retired, cancelled) of a given credit batch.
func (s serverImpl) Supply(ctx context.Context, request *v1beta1.QuerySupplyRequest) (*v1beta1.QuerySupplyResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	batch, err := s.stateStore.BatchInfoStore().GetByBatchDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
	}

	supply, err := s.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
	if err != nil {
		return nil, err
	}

	return &v1beta1.QuerySupplyResponse{
		TradableSupply:  supply.TradableAmount,
		RetiredSupply:   supply.RetiredAmount,
		CancelledAmount: supply.CancelledAmount,
	}, nil
}

// CreditTypes queries the list of allowed types that credit classes can have.
func (s serverImpl) CreditTypes(ctx context.Context, _ *v1beta1.QueryCreditTypesRequest) (*v1beta1.QueryCreditTypesResponse, error) {
	creditTypes := make([]*v1beta1.CreditType, 0)
	it, err := s.stateStore.CreditTypeStore().List(ctx, ecocreditv1beta1.CreditTypePrimaryKey{})
	if err != nil {
		return nil, err
	}
	for it.Next() {
		ct, err := it.Value()
		if err != nil {
			return nil, err
		}
		creditTypes = append(creditTypes, &v1beta1.CreditType{
			Abbreviation: ct.Abbreviation,
			Name:         ct.Name,
			Unit:         ct.Unit,
			Precision:    ct.Precision,
		})
	}
	return &v1beta1.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}
