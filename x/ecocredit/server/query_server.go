package server

import (
	"context"

	"github.com/regen-network/regen-ledger/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) Classes(goCtx context.Context, request *ecocredit.QueryClassesRequest) (*ecocredit.QueryClassesResponse, error) {
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

func (s serverImpl) ClassInfo(goCtx context.Context, request *ecocredit.QueryClassInfoRequest) (*ecocredit.QueryClassInfoResponse, error) {
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

func (s serverImpl) Batches(goCtx context.Context, request *ecocredit.QueryBatchesRequest) (*ecocredit.QueryBatchesResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	// Only read IDs that have a prefix match with the ClassID
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

func (s serverImpl) BatchInfo(goCtx context.Context, request *ecocredit.QueryBatchInfoRequest) (*ecocredit.QueryBatchInfoResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	var batchInfo ecocredit.BatchInfo
	err := s.batchInfoTable.GetOne(ctx, orm.RowID(request.BatchDenom), &batchInfo)
	return &ecocredit.QueryBatchInfoResponse{Info: &batchInfo}, err
}

func (s serverImpl) Balance(goCtx context.Context, request *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	acc := request.Account
	denom := batchDenomT(request.BatchDenom)
	store := ctx.KVStore(s.storeKey)

	tradable, err := getDecimal(store, TradableBalanceKey(acc, denom))
	if err != nil {
		return nil, err
	}

	retired, err := getDecimal(store, RetiredBalanceKey(acc, denom))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBalanceResponse{
		TradableAmount: tradable.String(),
		RetiredAmount:  retired.String(),
	}, nil
}

func (s serverImpl) Supply(goCtx context.Context, request *ecocredit.QuerySupplyRequest) (*ecocredit.QuerySupplyResponse, error) {
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

func (s serverImpl) CreditTypes(goCtx context.Context, _ *ecocredit.QueryCreditTypesRequest) (*ecocredit.QueryCreditTypesResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx).Context
	creditTypes := s.getAllCreditTypes(ctx)
	return &ecocredit.QueryCreditTypesResponse{CreditTypes: creditTypes}, nil
}
