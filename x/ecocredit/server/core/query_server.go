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
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	creditTypes := s.getAllCreditTypes(sdkCtx)
	return &v1beta1.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}

// Params queries the ecocredit module parameters.
func (s serverImpl) Params(ctx context.Context, _ *v1beta1.QueryParamsRequest) (*v1beta1.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var params ecocredit.Params
	s.paramSpace.GetParamSet(sdkCtx, &params)
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

// SellOrder queries for information about a sell order by its ID
// TODO: impl with ORM https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) SellOrder(goCtx context.Context, request *ecocredit.QuerySellOrderRequest) (*ecocredit.QuerySellOrderResponse, error) {
	panic("impl me!")
	//if request == nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "empty request")
	//}
	//
	//ctx := types.UnwrapSDKContext(goCtx)
	//sellOrder, err := s.getSellOrder(ctx, request.SellOrderId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.QuerySellOrderResponse{SellOrder: sellOrder}, nil
}

// SellOrders queries for all sell orders with pagination.
// TODO: impl with ORM https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) SellOrders(goCtx context.Context, request *ecocredit.QuerySellOrdersRequest) (*ecocredit.QuerySellOrdersResponse, error) {
	panic("impl me!")
	//if request == nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "empty request")
	//}
	//
	//ctx := types.UnwrapSDKContext(goCtx)
	//ordersIter, err := s.sellOrderTable.PrefixScan(ctx, 1, math.MaxUint64)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var orders []*ecocredit.SellOrder
	//pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.QuerySellOrdersResponse{
	//	SellOrders: orders,
	//	Pagination: pageResp,
	//}, nil
}

// SellOrdersByAddress queries for all sell orders by address with pagination.
// TODO: impl with ORM https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) SellOrdersByAddress(goCtx context.Context, request *ecocredit.QuerySellOrdersByAddressRequest) (*ecocredit.QuerySellOrdersByAddressResponse, error) {
	panic("impl me!")
	//if request == nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "empty request")
	//}
	//
	//addr, err := sdk.AccAddressFromBech32(request.Address)
	//if err != nil {
	//	return nil, err
	//}
	//
	//ctx := types.UnwrapSDKContext(goCtx)
	//ordersIter, err := s.sellOrderByAddressIndex.GetPaginated(ctx, addr.Bytes(), request.Pagination)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var orders []*ecocredit.SellOrder
	//pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.QuerySellOrdersByAddressResponse{
	//	SellOrders: orders,
	//	Pagination: pageResp,
	//}, nil
}

// SellOrdersByBatchDenom queries for all sell orders by address with pagination.
// TODO: impl with ORM https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) SellOrdersByBatchDenom(goCtx context.Context, request *ecocredit.QuerySellOrdersByBatchDenomRequest) (*ecocredit.QuerySellOrdersByBatchDenomResponse, error) {
	panic("impl me!")
	//if request == nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "empty request")
	//}
	//
	//if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
	//	return nil, err
	//}
	//
	//ctx := types.UnwrapSDKContext(goCtx)
	//ordersIter, err := s.sellOrderByBatchDenomIndex.GetPaginated(ctx, request.BatchDenom, request.Pagination)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var orders []*ecocredit.SellOrder
	//pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.QuerySellOrdersByBatchDenomResponse{
	//	SellOrders: orders,
	//	Pagination: pageResp,
	//}, nil
}

// BuyOrder queries for information about a buy order by its ID
// TODO: impl with ORM https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) BuyOrder(goCtx context.Context, request *ecocredit.QueryBuyOrderRequest) (*ecocredit.QueryBuyOrderResponse, error) {
	panic("impl me!")
	//if request == nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "empty request")
	//}
	//
	//ctx := types.UnwrapSDKContext(goCtx)
	//buyOrder, err := s.getBuyOrder(ctx, request.BuyOrderId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.QueryBuyOrderResponse{BuyOrder: buyOrder}, nil
}

// BuyOrders queries for all buy orders with pagination.
// TODO: impl with ORM https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) BuyOrders(goCtx context.Context, request *ecocredit.QueryBuyOrdersRequest) (*ecocredit.QueryBuyOrdersResponse, error) {
	panic("impl me!")
	//if request == nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "empty request")
	//}
	//
	//ctx := types.UnwrapSDKContext(goCtx)
	//ordersIter, err := s.buyOrderTable.PrefixScan(ctx, 1, math.MaxUint64)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var orders []*ecocredit.BuyOrder
	//pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.QueryBuyOrdersResponse{
	//	BuyOrders:  orders,
	//	Pagination: pageResp,
	//}, nil
}

// BuyOrdersByAddress queries for all buy orders by address with pagination.
// TODO: impl with ORM https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) BuyOrdersByAddress(goCtx context.Context, request *ecocredit.QueryBuyOrdersByAddressRequest) (*ecocredit.QueryBuyOrdersByAddressResponse, error) {
	panic("impl me!")
	//if request == nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "empty request")
	//}
	//
	//ctx := types.UnwrapSDKContext(goCtx)
	//addr, err := sdk.AccAddressFromBech32(request.Address)
	//if err != nil {
	//	return nil, err
	//}
	//
	//ordersIter, err := s.buyOrderByAddressIndex.GetPaginated(ctx, addr.Bytes(), request.Pagination)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var orders []*ecocredit.BuyOrder
	//pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.QueryBuyOrdersByAddressResponse{
	//	BuyOrders:  orders,
	//	Pagination: pageResp,
	//}, nil
}

// AllowedAskDenoms queries for all allowed ask denoms with pagination.
// TODO: impl with ORM https://github.com/regen-network/regen-ledger/issues/728
func (s serverImpl) AllowedAskDenoms(goCtx context.Context, request *ecocredit.QueryAllowedAskDenomsRequest) (*ecocredit.QueryAllowedAskDenomsResponse, error) {
	panic("impl me!")
	//if request == nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "empty request")
	//}
	//
	//ctx := types.UnwrapSDKContext(goCtx)
	//denomsIter, err := s.askDenomTable.PrefixScan(ctx, nil, nil)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var denoms []*ecocredit.AskDenom
	//pageResp, err := orm.Paginate(denomsIter, request.Pagination, &denoms)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.QueryAllowedAskDenomsResponse{
	//	AskDenoms:  denoms,
	//	Pagination: pageResp,
	//}, nil
}

// TODO: baskets PR queries

func (s serverImpl) Basket(ctx context.Context, request *ecocredit.QueryBasketRequest) (*ecocredit.QueryBasketResponse, error) {
	panic("implement me")
}

func (s serverImpl) Baskets(ctx context.Context, request *ecocredit.QueryBasketsRequest) (*ecocredit.QueryBasketsResponse, error) {
	panic("implement me")
}

func (s serverImpl) BasketCredits(ctx context.Context, request *ecocredit.QueryBasketCreditsRequest) (*ecocredit.QueryBasketCreditsResponse, error) {
	panic("implement me")
}
