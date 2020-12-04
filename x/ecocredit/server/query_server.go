package server

import (
	"github.com/regen-network/regen-ledger/util/storehelpers"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/bank/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) ClassInfo(ctx sdk.Context, request *ecocredit.QueryClassInfoRequest) (*ecocredit.QueryClassInfoResponse, error) {
	classInfo, err := s.getClassInfo(ctx, request.ClassId)
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

func (s serverImpl) BatchInfo(ctx sdk.Context, request *ecocredit.QueryBatchInfoRequest) (*ecocredit.QueryBatchInfoResponse, error) {
	var batchInfo ecocredit.BatchInfo
	err := s.batchInfoTable.GetOne(ctx, orm.RowID(request.BatchDenom), &batchInfo)
	return &ecocredit.QueryBatchInfoResponse{Info: &batchInfo}, err
}

func (s serverImpl) Balance(ctx sdk.Context, request *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	acc := request.Account
	denom := batchDenomT(request.BatchDenom)
	store := ctx.KVStore(s.key)

	tradable, err := storehelpers.GetDecimal(store, TradableBalanceKey(acc, denom))
	if err != nil {
		return nil, err
	}

	retired, err := storehelpers.GetDecimal(store, RetiredBalanceKey(acc, denom))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBalanceResponse{
		TradableUnits: math.DecimalString(tradable),
		RetiredUnits:  math.DecimalString(retired),
	}, nil
}

func (s serverImpl) Supply(ctx sdk.Context, request *ecocredit.QuerySupplyRequest) (*ecocredit.QuerySupplyResponse, error) {
	store := ctx.KVStore(s.key)
	denom := batchDenomT(request.BatchDenom)

	tradable, err := storehelpers.GetDecimal(store, TradableSupplyKey(denom))
	if err != nil {
		return nil, err
	}

	retired, err := storehelpers.GetDecimal(store, RetiredSupplyKey(denom))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QuerySupplyResponse{
		TradableSupply: math.DecimalString(tradable),
		RetiredSupply:  math.DecimalString(retired),
	}, nil
}
