package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/modules/incubator/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
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

func (s serverImpl) getRetiredSupply(ctx sdk.Context, batchDenom string) (sdk.Int, error) {
	store := ctx.KVStore(s.storeKey)

	var intProto sdk.IntProto
	bz := store.Get(RetiredSupplyKey(batchDenom))
	if bz == nil {
		return sdk.ZeroInt(), nil
	}

	err := intProto.Unmarshal(bz)
	if err != nil {
		return sdk.Int{}, err
	}

	return intProto.Int, nil
}

func (s serverImpl) Balance(goCtx context.Context, request *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	acc := request.Account
	denom := request.BatchDenom

	addr, err := sdk.AccAddressFromBech32(acc)
	if err != nil {
		return nil, err
	}

	coin := s.bankKeeper.GetBalance(ctx, addr, denom)

	retiredBalance, err := s.getRetiredBalance(ctx, acc, denom)
	if err != nil {
		return nil, err
	}

	return &ecocredit.QueryBalanceResponse{
		TradeableUnits: coin.Amount.String(),
		RetiredUnits:   retiredBalance.String(),
	}, nil
}

func (s serverImpl) getRetiredBalance(ctx sdk.Context, holder string, batchDenom string) (sdk.Int, error) {
	store := ctx.KVStore(s.storeKey)

	var intProto sdk.IntProto
	bz := store.Get(RetiredBalanceKey(holder, batchDenom))
	if bz == nil {
		return sdk.ZeroInt(), nil
	}

	err := intProto.Unmarshal(bz)
	if err != nil {
		return sdk.Int{}, err
	}

	return intProto.Int, nil
}
