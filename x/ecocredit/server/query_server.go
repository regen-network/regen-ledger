package server

import (
	"context"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Classes queries for all credit classes with pagination.
func (s serverImpl) Classes(goCtx context.Context, request *ecocredit.QueryClassesRequest) (*ecocredit.QueryClassesResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	classesIter, err := s.classInfoTable.PrefixScan(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	var classes []*ecocredit.ClassInfo
	pageResp, err := orm.Paginate(classesIter, request.Pagination, &classes)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryClassesResponse{
		Classes:    classes,
		Pagination: pageResp,
	}, nil
}

// ClassInfo queries for information on a credit class.
func (s serverImpl) ClassInfo(goCtx context.Context, request *ecocredit.QueryClassInfoRequest) (*ecocredit.QueryClassInfoResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	classInfo, err := s.getClassInfo(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryClassInfoResponse{Info: classInfo}, nil
}

func (s serverImpl) getClassInfo(ctx types.Context, classID string) (*ecocredit.ClassInfo, error) {
	var classInfo ecocredit.ClassInfo
	err := s.classInfoTable.GetOne(ctx, orm.RowID(classID), &classInfo)
	return &classInfo, err
}

// Batches queries for all batches in the given credit class.
func (s serverImpl) Batches(goCtx context.Context, request *ecocredit.QueryBatchesRequest) (*ecocredit.QueryBatchesResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if err := ecocredit.ValidateClassID(request.ClassId); err != nil {
		return nil, err
	}

	// Only read IDs that have a prefix match with the ClassID
	ctx := types.UnwrapSDKContext(goCtx)
	start, end := orm.PrefixRange([]byte(request.ClassId))
	batchesIter, err := s.batchInfoTable.PrefixScan(ctx, start, end)
	if err != nil {
		return nil, err
	}

	var batches []*ecocredit.BatchInfo
	pageResp, err := orm.Paginate(batchesIter, request.Pagination, &batches)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBatchesResponse{
		Batches:    batches,
		Pagination: pageResp,
	}, nil
}

// BatchInfo queries for information on a credit batch.
func (s serverImpl) BatchInfo(goCtx context.Context, request *ecocredit.QueryBatchInfoRequest) (*ecocredit.QueryBatchInfoResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	ctx := types.UnwrapSDKContext(goCtx)
	var batchInfo ecocredit.BatchInfo
	err := s.batchInfoTable.GetOne(ctx, orm.RowID(request.BatchDenom), &batchInfo)

	return &ecocredit.QueryBatchInfoResponse{Info: &batchInfo}, err
}

// Balance queries the balance (both tradable and retired) of a given credit
// batch for a given account.
func (s serverImpl) Balance(goCtx context.Context, request *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	ctx := types.UnwrapSDKContext(goCtx)
	acc := request.Account
	denom := batchDenomT(request.BatchDenom)
	store := ctx.KVStore(s.storeKey)
	accAddr, err := sdk.AccAddressFromBech32(acc)
	if err != nil {
		return nil, err
	}

	tradable, err := getDecimal(store, TradableBalanceKey(accAddr, denom))
	if err != nil {
		return nil, err
	}

	retired, err := getDecimal(store, RetiredBalanceKey(accAddr, denom))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBalanceResponse{
		TradableAmount: tradable.String(),
		RetiredAmount:  retired.String(),
	}, nil
}

// Supply queries the tradable and retired supply of a credit batch.
func (s serverImpl) Supply(goCtx context.Context, request *ecocredit.QuerySupplyRequest) (*ecocredit.QuerySupplyResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	denom := batchDenomT(request.BatchDenom)

	tradable, err := getDecimal(store, TradableSupplyKey(denom))
	if err != nil {
		return nil, err
	}

	retired, err := getDecimal(store, RetiredSupplyKey(denom))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QuerySupplyResponse{
		TradableSupply: tradable.String(),
		RetiredSupply:  retired.String(),
	}, nil
}

// CreditTypes queries the list of allowed types that credit classes can have.
func (s serverImpl) CreditTypes(goCtx context.Context, _ *ecocredit.QueryCreditTypesRequest) (*ecocredit.QueryCreditTypesResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx).Context
	creditTypes := s.getAllCreditTypes(ctx)
	return &ecocredit.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}

// Params queries the ecocredit module parameters.
func (s serverImpl) Params(goCtx context.Context, req *ecocredit.QueryParamsRequest) (*ecocredit.QueryParamsResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx).Context
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx, &params)
	return &ecocredit.QueryParamsResponse{Params: &params}, nil
}

// SellOrder queries for information about a sell order by its ID
func (s serverImpl) SellOrder(goCtx context.Context, request *ecocredit.QuerySellOrderRequest) (*ecocredit.QuerySellOrderResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	sellOrder, err := s.getSellOrder(ctx, request.SellOrderId)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QuerySellOrderResponse{SellOrder: sellOrder}, nil
}

func (s serverImpl) getSellOrder(ctx types.Context, orderID uint64) (*ecocredit.SellOrder, error) {
	var sellOrder ecocredit.SellOrder
	_, err := s.sellOrderTable.GetOne(ctx, orderID, &sellOrder)
	return &sellOrder, err
}

// SellOrders queries for all sell orders with pagination.
func (s serverImpl) SellOrders(goCtx context.Context, request *ecocredit.QuerySellOrdersRequest) (*ecocredit.QuerySellOrdersResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	ordersIter, err := s.sellOrderTable.PrefixScan(ctx, 1, math.MaxUint64)
	if err != nil {
		return nil, err
	}

	var orders []*ecocredit.SellOrder
	pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QuerySellOrdersResponse{
		SellOrders: orders,
		Pagination: pageResp,
	}, nil
}

// SellOrdersByAddress queries for all sell orders by address with pagination.
func (s serverImpl) SellOrdersByAddress(goCtx context.Context, request *ecocredit.QuerySellOrdersByAddressRequest) (*ecocredit.QuerySellOrdersByAddressResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, err
	}

	ctx := types.UnwrapSDKContext(goCtx)
	ordersIter, err := s.sellOrderByAddressIndex.GetPaginated(ctx, addr.Bytes(), request.Pagination)
	if err != nil {
		return nil, err
	}

	var orders []*ecocredit.SellOrder
	pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QuerySellOrdersByAddressResponse{
		SellOrders: orders,
		Pagination: pageResp,
	}, nil
}

// SellOrdersByBatchDenom queries for all sell orders by address with pagination.
func (s serverImpl) SellOrdersByBatchDenom(goCtx context.Context, request *ecocredit.QuerySellOrdersByBatchDenomRequest) (*ecocredit.QuerySellOrdersByBatchDenomResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if err := ecocredit.ValidateDenom(request.BatchDenom); err != nil {
		return nil, err
	}

	ctx := types.UnwrapSDKContext(goCtx)
	ordersIter, err := s.sellOrderByBatchDenomIndex.GetPaginated(ctx, request.BatchDenom, request.Pagination)
	if err != nil {
		return nil, err
	}

	var orders []*ecocredit.SellOrder
	pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QuerySellOrdersByBatchDenomResponse{
		SellOrders: orders,
		Pagination: pageResp,
	}, nil
}

// BuyOrder queries for information about a buy order by its ID
func (s serverImpl) BuyOrder(goCtx context.Context, request *ecocredit.QueryBuyOrderRequest) (*ecocredit.QueryBuyOrderResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	buyOrder, err := s.getBuyOrder(ctx, request.BuyOrderId)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBuyOrderResponse{BuyOrder: buyOrder}, nil
}

func (s serverImpl) getBuyOrder(ctx types.Context, orderID uint64) (*ecocredit.BuyOrder, error) {
	var buyOrder ecocredit.BuyOrder
	_, err := s.buyOrderTable.GetOne(ctx, orderID, &buyOrder)
	return &buyOrder, err
}

// BuyOrders queries for all buy orders with pagination.
func (s serverImpl) BuyOrders(goCtx context.Context, request *ecocredit.QueryBuyOrdersRequest) (*ecocredit.QueryBuyOrdersResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	ordersIter, err := s.buyOrderTable.PrefixScan(ctx, 1, math.MaxUint64)
	if err != nil {
		return nil, err
	}

	var orders []*ecocredit.BuyOrder
	pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBuyOrdersResponse{
		BuyOrders:  orders,
		Pagination: pageResp,
	}, nil
}

// BuyOrdersByAddress queries for all buy orders by address with pagination.
func (s serverImpl) BuyOrdersByAddress(goCtx context.Context, request *ecocredit.QueryBuyOrdersByAddressRequest) (*ecocredit.QueryBuyOrdersByAddressResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	addr, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, err
	}

	ordersIter, err := s.buyOrderByAddressIndex.GetPaginated(ctx, addr.Bytes(), request.Pagination)
	if err != nil {
		return nil, err
	}

	var orders []*ecocredit.BuyOrder
	pageResp, err := orm.Paginate(ordersIter, request.Pagination, &orders)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBuyOrdersByAddressResponse{
		BuyOrders:  orders,
		Pagination: pageResp,
	}, nil
}

// AllowedAskDenoms queries for all allowed ask denoms with pagination.
func (s serverImpl) AllowedAskDenoms(goCtx context.Context, request *ecocredit.QueryAllowedAskDenomsRequest) (*ecocredit.QueryAllowedAskDenomsResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := types.UnwrapSDKContext(goCtx)
	denomsIter, err := s.askDenomTable.PrefixScan(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	var denoms []*ecocredit.AskDenom
	pageResp, err := orm.Paginate(denomsIter, request.Pagination, &denoms)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryAllowedAskDenomsResponse{
		AskDenoms:  denoms,
		Pagination: pageResp,
	}, nil
}
