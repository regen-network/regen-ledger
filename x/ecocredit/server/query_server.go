package server

import (
	"context"

	"github.com/regen-network/regen-ledger/util/storehelpers"
	"github.com/regen-network/regen-ledger/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/bank/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s keeper) ClassInfo(ctx context.Context, request *ecocredit.QueryClassInfoRequest) (*ecocredit.QueryClassInfoResponse, error) {
	classInfo, err := s.getClassInfo(sdk.UnwrapSDKContext(ctx), request.ClassId)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryClassInfoResponse{Info: classInfo}, nil
}

func (s keeper) getClassInfo(ctx sdk.Context, classID string) (*ecocredit.ClassInfo, error) {
	var classInfo ecocredit.ClassInfo
	err := s.classInfoTable.GetOne(ctx, orm.RowID(classID), &classInfo)
	return &classInfo, err
}

func (s keeper) BatchInfo(goCtx context.Context, request *ecocredit.QueryBatchInfoRequest) (*ecocredit.QueryBatchInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var batchInfo ecocredit.BatchInfo
	err := s.batchInfoTable.GetOne(ctx, orm.RowID(request.BatchDenom), &batchInfo)
	return &ecocredit.QueryBatchInfoResponse{Info: &batchInfo}, err
}

func (s keeper) Balance(goCtx context.Context, request *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	acc := request.Account
	denom := batchDenomT(request.BatchDenom)
	store := ctx.KVStore(s.key)

	//tradable, err := storehelpers.GetDecimal(store, TradableBalanceKey(acc, denom))
	//if err != nil {
	//	return nil, err
	//}

	res, err := s.bankQueryClient.Balance(goCtx, &bank.QueryBalanceRequest{
		Address: acc,
		Denom:   request.BatchDenom,
	})
	if err != nil {
		return nil, err
	}

	retired, err := storehelpers.GetDecimal(store, RetiredBalanceKey(acc, denom))
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBalanceResponse{
		TradableUnits: res.Balance.Amount,
		RetiredUnits:  math.DecimalString(retired),
	}, nil
}

func (s keeper) Supply(goCtx context.Context, request *ecocredit.QuerySupplyRequest) (*ecocredit.QuerySupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
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
