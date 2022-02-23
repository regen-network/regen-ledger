package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Classes queries for all credit classes with pagination.
func (k Keeper) Classes(ctx context.Context, request *v1beta1.QueryClassesRequest) (*v1beta1.QueryClassesResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.ClassInfoStore().List(ctx, &ecocreditv1beta1.ClassInfoPrimaryKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	infos := make([]*v1beta1.ClassInfo, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}
		var ci v1beta1.ClassInfo
		if err = PulsarToGogoSlow(info, &ci); err != nil {
			return nil, err
		}
		infos = append(infos, &ci)
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1beta1.QueryClassesResponse{
		Classes:    infos,
		Pagination: pr,
	}, err
}

// ClassInfo queries for information on a credit class.
func (k Keeper) ClassInfo(ctx context.Context, request *v1beta1.QueryClassInfoRequest) (*v1beta1.QueryClassInfoResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if err := ecocredit.ValidateClassID(request.ClassId); err != nil {
		return nil, err
	}
	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}
	var ci v1beta1.ClassInfo
	if err = PulsarToGogoSlow(classInfo, &ci); err != nil {
		return nil, err
	}
	return &v1beta1.QueryClassInfoResponse{Info: &ci}, nil
}

func (k Keeper) ClassIssuers(ctx context.Context, request *v1beta1.QueryClassIssuersRequest) (*v1beta1.QueryClassIssuersResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	if err := ecocredit.ValidateClassID(request.ClassId); err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.ClassIssuerStore().List(ctx, ecocreditv1beta1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(classInfo.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	issuers := make([]string, 0)
	for it.Next() {
		val, err := it.Value()
		if err != nil {
			return nil, err
		}
		issuers = append(issuers, sdk.AccAddress(val.Issuer).String())
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1beta1.QueryClassIssuersResponse{
		Issuers:    issuers,
		Pagination: pr,
	}, nil
}

// Projects queries all projects of a given credit class.
func (k Keeper) Projects(ctx context.Context, request *v1beta1.QueryProjectsRequest) (*v1beta1.QueryProjectsResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	cInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.ProjectInfoStore().List(ctx, ecocreditv1beta1.ProjectInfoClassIdNameIndexKey{}.WithClassId(cInfo.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	projectInfos := make([]*v1beta1.ProjectInfo, 0)
	for it.Next() {
		info, err := it.Value()
		if err != nil {
			return nil, err
		}
		var pi v1beta1.ProjectInfo
		if err = PulsarToGogoSlow(info, &pi); err != nil {
			return nil, err
		}
		projectInfos = append(projectInfos, &pi)
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1beta1.QueryProjectsResponse{
		Projects:   projectInfos,
		Pagination: pr,
	}, nil
}

func (k Keeper) ProjectInfo(ctx context.Context, request *v1beta1.QueryProjectInfoRequest) (*v1beta1.QueryProjectInfoResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if err := ecocredit.ValidateProjectID(request.ProjectId); err != nil {
		return nil, err
	}
	pInfo, err := k.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	var pi v1beta1.ProjectInfo
	if err = PulsarToGogoSlow(pInfo, &pi); err != nil {
		return nil, err
	}
	return &v1beta1.QueryProjectInfoResponse{Info: &pi}, nil
}

// Batches queries for all batches in the given credit class.
func (k Keeper) Batches(ctx context.Context, request *v1beta1.QueryBatchesRequest) (*v1beta1.QueryBatchesResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	pg, err := GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}
	project, err := k.stateStore.ProjectInfoStore().GetByName(ctx, request.ProjectId)
	if err != nil {
		return nil, err
	}
	it, err := k.stateStore.BatchInfoStore().List(ctx, ecocreditv1beta1.BatchInfoProjectIdIndexKey{}.WithProjectId(project.Id), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	batches := make([]*v1beta1.BatchInfo, 0)
	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return nil, err
		}
		var bi v1beta1.BatchInfo
		if err = PulsarToGogoSlow(batch, &bi); err != nil {
			return nil, err
		}
		batches = append(batches, &bi)
	}
	pr, err := PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}
	return &v1beta1.QueryBatchesResponse{
		Batches:    batches,
		Pagination: pr,
	}, nil
}

// BatchInfo queries for information on a credit batch.
func (k Keeper) BatchInfo(ctx context.Context, request *v1beta1.QueryBatchInfoRequest) (*v1beta1.QueryBatchInfoResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
	}
	var bi v1beta1.BatchInfo
	if err = PulsarToGogoSlow(batch, &bi); err != nil {
		return nil, err
	}
	return &v1beta1.QueryBatchInfoResponse{Info: &bi}, nil
}

// Balance queries the balance (both tradable and retired) of a given credit
// batch for a given account.
func (k Keeper) Balance(ctx context.Context, req *v1beta1.QueryBalanceRequest) (*v1beta1.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if err := ecocredit.ValidateDenom(req.BatchDenom); err != nil {
		return nil, err
	}
	batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, err
	}
	if batch == nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("batch with denom %s not found", req.BatchDenom)
	}
	addr, _ := sdk.AccAddressFromBech32(req.Account)

	balance, err := k.stateStore.BatchBalanceStore().Get(ctx, addr, batch.Id)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return &v1beta1.QueryBalanceResponse{
				TradableAmount: "0",
				RetiredAmount:  "0",
			}, nil
		}
		return nil, err
	}
	return &v1beta1.QueryBalanceResponse{
		TradableAmount: balance.Tradable,
		RetiredAmount:  balance.Retired,
	}, nil
}

// Supply queries the supply (tradable, retired, cancelled) of a given credit batch.
func (k Keeper) Supply(ctx context.Context, request *v1beta1.QuerySupplyRequest) (*v1beta1.QuerySupplyResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchInfoStore().GetByBatchDenom(ctx, request.BatchDenom)
	if err != nil {
		return nil, err
	}

	supply, err := k.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
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
func (k Keeper) CreditTypes(ctx context.Context, _ *v1beta1.QueryCreditTypesRequest) (*v1beta1.QueryCreditTypesResponse, error) {
	creditTypes := make([]*v1beta1.CreditType, 0)
	it, err := k.stateStore.CreditTypeStore().List(ctx, ecocreditv1beta1.CreditTypePrimaryKey{})
	if err != nil {
		return nil, err
	}
	for it.Next() {
		ct, err := it.Value()
		if err != nil {
			return nil, err
		}
		var cType v1beta1.CreditType
		if err = PulsarToGogoSlow(ct, &cType); err != nil {
			return nil, err
		}
		creditTypes = append(creditTypes, &cType)
	}
	return &v1beta1.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}

// Params queries the ecocredit module parameters.
// TODO: remove params https://github.com/regen-network/regen-ledger/issues/729
// Currently this is an ugly hack that grabs v1alpha types and converts them into v1beta types.
// will be gone with #729.
func (k Keeper) Params(ctx context.Context, _ *v1beta1.QueryParamsRequest) (*v1beta1.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var params ecocredit.Params
	k.params.GetParamSet(sdkCtx, &params)
	v1beta1types := make([]*v1beta1.CreditType, len(params.CreditTypes))
	for i, typ := range params.CreditTypes {
		v1beta1types[i] = &v1beta1.CreditType{
			Abbreviation: typ.Abbreviation,
			Name:         typ.Name,
			Unit:         typ.Unit,
			Precision:    typ.Precision,
		}
	}
	v1beta1Params := v1beta1.Params{
		CreditClassFee:       params.CreditClassFee,
		AllowedClassCreators: params.AllowedClassCreators,
		AllowlistEnabled:     params.AllowlistEnabled,
		CreditTypes:          v1beta1types,
	}
	return &v1beta1.QueryParamsResponse{Params: &v1beta1Params}, nil
}
