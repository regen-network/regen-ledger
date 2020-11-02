package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/modules/incubator/orm"
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

func (s serverImpl) getClassInfo(ctx sdk.Context, classId string) (*ecocredit.ClassInfo, error) {
	var classInfo ecocredit.ClassInfo
	err := s.classInfoTable.GetOne(ctx, orm.RowID(classId), &classInfo)
	if err != nil {
		return nil, err
	}

	return &classInfo, nil
}

func (s serverImpl) BatchInfo(goCtx context.Context, request *ecocredit.QueryBatchInfoRequest) (*ecocredit.QueryBatchInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var batchInfo ecocredit.BatchInfo
	err := s.batchInfoTable.GetOne(ctx, orm.RowID(request.BatchDenom), &batchInfo)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBatchInfoResponse{Info: &batchInfo}, nil
}

func (s serverImpl) Balance(goCtx context.Context, request *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	acc := request.Account
	denom := batchDenomT(request.BatchDenom)

	store := ctx.KVStore(s.storeKey)

	tradeable, err := storeGetDec(store, TradeableBalanceKey(acc, denom))
	if err != nil {
		return nil, err
	}

	retired, err := storeGetDec(store, RetiredBalanceKey(acc, denom))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBalanceResponse{
		TradeableUnits: math.DecString(tradeable),
		RetiredUnits:   math.DecString(retired),
	}, nil
}

func (s serverImpl) Supply(goCtx context.Context, request *ecocredit.QuerySupplyRequest) (*ecocredit.QuerySupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	denom := batchDenomT(request.BatchDenom)

	tradeable, err := storeGetDec(store, TradeableSupplyKey(denom))
	if err != nil {
		return nil, err
	}

	retired, err := storeGetDec(store, RetiredSupplyKey(denom))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QuerySupplyResponse{
		TradeableSupply: math.DecString(tradeable),
		RetiredSupply:   math.DecString(retired),
	}, nil
}
