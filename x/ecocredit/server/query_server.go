package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/math"
)

func (s serverImpl) ClassInfo(ctx context.Context, request *ecocredit.QueryClassInfoRequest) (*ecocredit.QueryClassInfoResponse, error) {
	classInfo, err := s.getClassInfo(sdk.UnwrapSDKContext(ctx), request.ClassId)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryClassInfoResponse{Info: classInfo}, nil
}

func (s serverImpl) getClassInfo(ctx sdk.Context, classID string) (*ecocredit.ClassInfo, error) {
	var classInfo ecocredit.ClassInfo
	err := s.classInfoTable.GetOne(ctx, orm.RowID(classID), &classInfo)
	return &classInfo, err
}

func (s serverImpl) BatchInfo(goCtx context.Context, request *ecocredit.QueryBatchInfoRequest) (*ecocredit.QueryBatchInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var batchInfo ecocredit.BatchInfo
	err := s.batchInfoTable.GetOne(ctx, orm.RowID(request.BatchDenom), &batchInfo)
	return &ecocredit.QueryBatchInfoResponse{Info: &batchInfo}, err
}

func (s serverImpl) Balance(goCtx context.Context, request *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
		TradableUnits: math.DecimalString(tradable),
		RetiredUnits:  math.DecimalString(retired),
	}, nil
}

func (s serverImpl) Supply(goCtx context.Context, request *ecocredit.QuerySupplyRequest) (*ecocredit.QuerySupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
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
		TradableSupply: math.DecimalString(tradable),
		RetiredSupply:  math.DecimalString(retired),
	}, nil
}

func (s serverImpl) Precision(goCtx context.Context, request *ecocredit.QueryPrecisionRequest) (*ecocredit.QueryPrecisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	x, err := getUint32(store, MaxDecimalPlacesKey(batchDenomT(request.BatchDenom)))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryPrecisionResponse{MaxDecimalPlaces: x}, nil
}
