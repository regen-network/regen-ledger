package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) CreateBasket(goCtx context.Context, req *ecocredit.MsgCreateBasket) (*ecocredit.MsgCreateBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	basketDenom := getBasketDenom(req.Curator, req.Name)
	err := s.basketInfoTable.Create(ctx, &ecocredit.BasketInfo{
		BasketDenom:       basketDenom,
		DisplayName:       req.DisplayName,
		Exponent:          req.Exponent,
		AdmissionCriteria: req.AdmissionCriteria,
		RetireOnTake:      req.RetireOnTake,
		AllowPicking:      req.AllowPicking,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateBasketResponse{BasketDenom: basketDenom}, nil
}

func (s serverImpl) AddToBasket(goCtx context.Context, req *ecocredit.MsgAddToBasket) (*ecocredit.MsgAddToBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)

	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	if !s.basketInfoTable.Has(ctx, orm.RowID(req.BasketDenom)) {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid basket denom", req.BasketDenom)
	}

	var basket ecocredit.BasketInfo
	err = s.basketInfoTable.GetOne(ctx, orm.RowID(req.BasketDenom), &basket)
	if err != nil {
		return nil, err
	}

	amtReceived := math.NewDecFromInt64(0)

	for _, credit := range req.Credits {
		batchDenom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(batchDenom)) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid credit batch denom", batchDenom)
		}

		// TODO Add AdmissionCriteria validation here

		maxDecimalPlaces, err := s.getBatchPrecision(ctx, batchDenom)
		if err != nil {
			return nil, err
		}

		tradable, err := math.NewNonNegativeFixedDecFromString(credit.TradableAmount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		amtReceived, err = amtReceived.Add(tradable)
		if err != nil {
			return nil, err
		}

		err = addAndSetDecimal(store, BasketCreditsKey(basketDenomT(req.BasketDenom), owner, batchDenom), tradable)
		if err != nil {
			return nil, err
		}
	}

	// TODO Why 10?
	multiplier, err := math.NewNonNegativeFixedDecFromString(basket.AdmissionCriteria[0].Multiplier, 10)
	if err != nil {
		return nil, err
	}
	multipliedAmtReceived, err := amtReceived.Mul(multiplier)

	// TODO Is there another way than to convert from math.Dec to sdk.Int other than passing by int64?
	i, err := multipliedAmtReceived.Int64()
	if err != nil {
		return nil, err
	}
	amtAsInt := sdk.NewIntFromUint64(uint64(i))
	basketTokens := sdk.NewCoins(sdk.NewCoin(basket.BasketDenom, amtAsInt))

	// TODO don't hardcode ecocredit string.
	err = s.bankKeeper.MintCoins(ctx.Context, "ecocredit", basketTokens)
	if err != nil {
		return nil, err
	}
	err = s.bankKeeper.SendCoinsFromModuleToAccount(ctx.Context, "ecocredit", owner, basketTokens)
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgAddToBasketResponse{
		AmountReceived: amtReceived.String(),
	}, nil
}

func (s serverImpl) TakeFromBasket(goCtx context.Context, req *ecocredit.MsgTakeFromBasket) (*ecocredit.MsgTakeFromBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	if !s.basketInfoTable.Has(ctx, orm.RowID(req.BasketDenom)) {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid basket denom", req.BasketDenom)
	}

}

func (s serverImpl) PickFromBasket(goCtx context.Context, req *ecocredit.MsgPickFromBasket) (*ecocredit.MsgPickFromBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	if !s.basketInfoTable.Has(ctx, orm.RowID(req.BasketDenom)) {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid basket denom", req.BasketDenom)
	}
}

func getBasketDenom(curator, name string) string {
	return fmt.Sprintf("ecocredit:%s:%s", curator, name)
}
