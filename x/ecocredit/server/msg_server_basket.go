package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"
	sdk "github.com/regen-network/regen-ledger/types"
	regenmath "github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"strings"
)

func (s serverImpl) CreateBasket(goCtx context.Context, req *ecocredit.MsgCreateBasket) (*ecocredit.MsgCreateBasketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// stateful validation of filters
	for _, criteria := range req.BasketCriteria {
		if err := s.validateFilterData(ctx, criteria.Filter); err != nil {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("invalid basket filter: %s", err.Error())
		}
	}

	basketDenom := getBasketDenom(req.Curator, req.Name)
	err := s.basketTable.Create(ctx, &ecocredit.Basket{
		BasketDenom:       basketDenom,
		Curator:           req.Curator,
		Name:              req.Name,
		DisplayName:       req.DisplayName,
		Exponent:          req.Exponent,
		BasketCriteria:    req.BasketCriteria,
		DisableAutoRetire: req.DisableAutoRetire,
		AllowPicking:      req.AllowPicking,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateBasketResponse{BasketDenom: basketDenom}, nil
}

func (s serverImpl) AddToBasket(goCtx context.Context, req *ecocredit.MsgAddToBasket) (*ecocredit.MsgAddToBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	//sdkCtx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	owner, _ := types.AccAddressFromBech32(req.Owner)

	// try to get the basket info
	var basket ecocredit.Basket
	err := s.basketTable.GetOne(ctx, orm.RowID(req.BasketDenom), &basket)
	if err != nil {
		return nil, err
	}

	totalCreditsDeposited := regenmath.NewDecFromInt64(0)
	//creditsToSend := make([]*ecocredit.MsgSend_SendCredits, len(req.Credits))
	//basketKeys := make([][]byte, len(req.Credits))
	for _, credit := range req.Credits {

		// verify this is a legit ecocredit batch
		batchDenom := batchDenomT(credit.BatchDenom)
		var batchInfo ecocredit.BatchInfo
		err := s.batchInfoTable.GetOne(ctx, orm.RowID(batchDenom), &batchInfo)
		if err != nil {
			return nil, err
		}

		res, err := s.ClassInfo(goCtx, &ecocredit.QueryClassInfoRequest{ClassId: batchInfo.ClassId})
		if err != nil {
			return nil, err
		}
		classInfo := res.Info

		// verify the user has sufficient ecocredits to send
		err = verifyCreditBalance(store, owner, credit.BatchDenom, credit.TradableAmount)
		if err != nil {
			return nil, err
		}

		// TODO(tyler): verify this credit matches the filter
		filters := make([]*ecocredit.Filter, len(basket.BasketCriteria))
		for i, bc := range basket.BasketCriteria {
			filters[i] = bc.Filter
		}
		_, err = checkFilters(filters, *classInfo, batchInfo, basket, req.Owner)
		if err != nil {
			return nil, err
		}

		// we dont have to check for error cause we already did in verifyCreditBalance
		creditsToDeposit, _ := regenmath.NewDecFromString(credit.TradableAmount)

		totalCreditsDeposited, err = totalCreditsDeposited.Add(creditsToDeposit)
		if err != nil {
			return nil, err
		}

		creditsToAddToBasket := &ecocredit.MsgSend_SendCredits{
			BatchDenom:         credit.BatchDenom,
			TradableAmount:     credit.TradableAmount,
			RetiredAmount:      "",
			RetirementLocation: "",
		}

		// send the ecocredits to the basket
		basketKey := BasketCreditsKey(basketDenomT(basket.BasketDenom), owner.Bytes(), batchDenom)
		derivedKey := s.storeKey.Derive(basketKey)
		if _, err := s.Send(goCtx, &ecocredit.MsgSend{
			Sender:    owner.String(),
			Recipient: derivedKey.Address().String(),
			Credits:   []*ecocredit.MsgSend_SendCredits{creditsToAddToBasket},
		}); err != nil {
			return nil, err
		}

	}

	// calculate how many basket tokens are to be awarded to the basket depositor
	basketTokens, err := calculateBasketTokens(totalCreditsDeposited, basket.Exponent)
	if err != nil {
		return nil, err
	}
	basketTokensInt, err := basketTokens.Int64()
	if err != nil {
		return nil, err
	}

	// mint basket tokens to send to basket depositor
	basketCoins := types.NewCoin(basket.BasketDenom, types.NewInt(basketTokensInt))
	err = s.bankKeeper.MintCoins(ctx, ecocredit.ModuleName, types.Coins{basketCoins})
	if err != nil {
		return nil, err
	}

	// send the basket tokens to the basket depositor
	err = s.bankKeeper.SendCoinsFromModuleToAccount(ctx, ecocredit.ModuleName, owner, types.Coins{basketCoins})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgAddToBasketResponse{AmountReceived: basketTokens.String()}, nil
}

func (s serverImpl) TakeFromBasket(goCtx context.Context, req *ecocredit.MsgTakeFromBasket) (*ecocredit.MsgTakeFromBasketResponse, error) {
	panic("implement me")
}

func (s serverImpl) PickFromBasket(goCtx context.Context, req *ecocredit.MsgPickFromBasket) (*ecocredit.MsgPickFromBasketResponse, error) {
	panic("implement me")
}

// ----- HELPER METHODS -----

func (s serverImpl) validateFilterData(ctx sdk.Context, filters ...*ecocredit.Filter) error {
	for _, filter := range filters {
		switch f := filter.Sum.(type) {
		case *ecocredit.Filter_And_:
			if err := s.validateFilterData(ctx, f.And.Filters...); err != nil {
				return err
			}
		case *ecocredit.Filter_Or_:
			if err := s.validateFilterData(ctx, f.Or.Filters...); err != nil {
				return err
			}
		case *ecocredit.Filter_CreditTypeName:
			if _, err := s.getCreditType(ctx.Context, f.CreditTypeName); err != nil {
				return sdkerrors.ErrNotFound.Wrapf("credit type %s not found", f.CreditTypeName)
			}
		case *ecocredit.Filter_ClassId:
			if exists := s.classInfoTable.Has(ctx, orm.RowID(f.ClassId)); !exists {
				return sdkerrors.ErrNotFound.Wrapf("credit class with id %s not found", f.ClassId)
			}
		case *ecocredit.Filter_BatchDenom:
			if exists := s.batchInfoTable.Has(ctx, orm.RowID(f.BatchDenom)); !exists {
				return sdkerrors.ErrNotFound.Wrapf("batch with denom %s not found", f.BatchDenom)
			}
		}
	}
	return nil
}

func padRight(length int, prefix, add string) string {
	builder := strings.Builder{}
	builder.Grow(length + len(prefix))

	builder.WriteString(prefix)
	for i := 0; i < length-1; i++ {
		builder.WriteString(add)
	}

	return builder.String()
}

// calculateBasketTokens calculates the basket tokens to be awarded based on how many ecocredits were added to the basket.
// the equation for calculating the award amount is as follows:
// total_credits_deposited * 10^(basket.Exponent)
func calculateBasketTokens(creditsDeposited regenmath.Dec, exponent uint32) (regenmath.Dec, error) {
	// get the str to use in the multiplier
	multiStr := padRight(int(exponent), "10", "0")

	// get the multiplier in dec form
	multiplier, err := regenmath.NewPositiveDecFromString(multiStr)
	if err != nil {
		return regenmath.Dec{}, err
	}

	// return the credits deposited * 10^(exponent)
	return creditsDeposited.Mul(multiplier)
}

func checkFilters(filters []*ecocredit.Filter, classInfo ecocredit.ClassInfo, batchInfo ecocredit.BatchInfo, basketInfo ecocredit.Basket, owner string) (int, error) {
	depth := len(filters)
	var err error
	for _, filter := range filters {
		switch f := filter.Sum.(type) {
		case *ecocredit.Filter_And_:
			andFilter := f.And.Filters
			andDepth := len(andFilter)
			innerDepth, err := checkFilters(andFilter, classInfo, batchInfo, basketInfo, owner)
			if andDepth != 0 || err != nil {
				return innerDepth, err
			} else {
				depth -= 1
			}
		case *ecocredit.Filter_Or_:
			orFilter := f.Or.Filters
			orDepth := len(orFilter)
			innerDepth, err := checkFilters(orFilter, classInfo, batchInfo, basketInfo, owner)
			if orDepth == innerDepth {
				return innerDepth, err
			} else {
				depth -= 1
			}
		case *ecocredit.Filter_CreditTypeName:
			if classInfo.CreditType.Name == f.CreditTypeName {
				depth -= 1
			} else {
				err = formatFilterError("credit type name", f.CreditTypeName, classInfo.CreditType.Name)
			}
		case *ecocredit.Filter_ClassId:
			if batchInfo.ClassId == f.ClassId {
				depth -= 1
			} else {
				err = formatFilterError("class id", f.ClassId, batchInfo.ClassId)
			}
		case *ecocredit.Filter_ProjectId:
			//  depth -= 1 TODO: need projects PR
		case *ecocredit.Filter_BatchDenom:
			if batchInfo.BatchDenom == f.BatchDenom {
				depth -= 1
			} else {
				err = formatFilterError("batch denom", f.BatchDenom, batchInfo.BatchDenom)
			}
		case *ecocredit.Filter_ClassAdmin:
			if classInfo.Admin == f.ClassAdmin {
				depth -= 1
			} else {
				err = formatFilterError("class admin", f.ClassAdmin, classInfo.Admin)
			}
		case *ecocredit.Filter_Issuer:
			if batchInfo.Issuer == f.Issuer {
				depth -= 1
			} else {
				err = formatFilterError("issuer", f.Issuer, batchInfo.Issuer)
			}
		case *ecocredit.Filter_Owner:
			if owner == f.Owner {
				depth -= 1
			} else {
				err = formatFilterError("credit owner", f.Owner, owner)
			}
		case *ecocredit.Filter_ProjectLocation:
			// depth -= 1 TODO: wait for projects PR
		case *ecocredit.Filter_DateRange_:
			if batchInfo.StartDate.Equal(*f.DateRange.StartDate) || batchInfo.StartDate.After(*f.DateRange.StartDate) {
				if batchInfo.EndDate.Equal(*f.DateRange.EndDate) || batchInfo.EndDate.Before(*f.DateRange.EndDate) {
					depth -= 1
				} else {
					err = formatFilterError("date range", f.DateRange.StartDate.String()+" to "+f.DateRange.EndDate.String(), batchInfo.StartDate.String()+" to "+batchInfo.EndDate.String())
				}
			} else {
				err = formatFilterError("date range", f.DateRange.StartDate.String()+" to "+f.DateRange.EndDate.String(), batchInfo.StartDate.String()+" to "+batchInfo.EndDate.String())
			}
		case *ecocredit.Filter_Tag:
		// depth -= 1 TODO: wait for tags PR
		default:
			err = errors.New("no valid filter given")
		}

	}

	if depth != 0 || err != nil {
		return depth, err
	}
	return depth, nil
}

func formatFilterError(item, want, got string) error {
	return fmt.Errorf("basket filter requires %s %s, but a credit with %s %s was given", item, got, item, want)
}

func getBasketDenom(curator, name string) string {
	return fmt.Sprintf("%s:%s:%s", ecocredit.ModuleName, curator, name)
}
