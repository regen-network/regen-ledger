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
	"sort"
	"strings"
	"time"
)

// CreateBasket creates a basket keyed by a given basketDenom.
func (s serverImpl) CreateBasket(goCtx context.Context, req *ecocredit.MsgCreateBasket) (*ecocredit.MsgCreateBasketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := s.validateFilterData(ctx, req.BasketCriteria); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("invalid basket filter: %s", err.Error())
	}

	// construct the basket denom - ecocredit:curatorAddress:basketName
	basketDenom := constructBasketDenom(req.Name)

	if err := s.basketTable.Create(ctx, &ecocredit.Basket{
		BasketDenom:       basketDenom,
		Curator:           req.Curator,
		Name:              req.Name,
		DisplayName:       req.DisplayName,
		Exponent:          req.Exponent,
		BasketCriteria:    req.BasketCriteria,
		DisableAutoRetire: req.DisableAutoRetire,
		AllowPicking:      req.AllowPicking,
	}); err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateBasketResponse{BasketDenom: basketDenom}, nil
}

// AddToBasket adds ecocredits to the basket if they comply with the basket's filter
// then mints and sends basket tokens to the depositor.
func (s serverImpl) AddToBasket(goCtx context.Context, req *ecocredit.MsgAddToBasket) (*ecocredit.MsgAddToBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	regenCtx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	owner, _ := types.AccAddressFromBech32(req.Owner)

	// get basket info
	var basket ecocredit.Basket
	err := s.basketTable.GetOne(ctx, orm.RowID(req.BasketDenom), &basket)
	if err != nil {
		return nil, err
	}

	totalCreditsDeposited := regenmath.NewDecFromInt64(0)
	for _, credit := range req.Credits {

		// get the batch
		batchDenom := batchDenomT(credit.BatchDenom)
		var batchInfo ecocredit.BatchInfo
		err := s.batchInfoTable.GetOne(ctx, orm.RowID(batchDenom), &batchInfo)
		if err != nil {
			return nil, err
		}

		// get project info
		res, err := s.ProjectInfo(goCtx, &ecocredit.QueryProjectInfoRequest{ProjectId: batchInfo.ProjectId})
		if err != nil {
			return nil, err
		}
		projectInfo := res.Info

		// get the class info
		res2, err := s.ClassInfo(goCtx, &ecocredit.QueryClassInfoRequest{ClassId: res.Info.ClassId})
		if err != nil {
			return nil, err
		}
		classInfo := res2.Info

		// TODO(Tyler): should we even check here for fast exiting? s.Send could take care of this.
		// verify the user has sufficient ecocredits to send
		err = verifyCreditBalance(store, owner, credit.BatchDenom, credit.TradableAmount)
		if err != nil {
			return nil, err
		}

		// check that the credits meet the filter
		if _, err = checkFilters([]*ecocredit.Filter{basket.BasketCriteria}, *classInfo, batchInfo, *projectInfo, req.Owner); err != nil {
			return nil, err
		}

		maxDec, err := s.getBatchPrecision(regenCtx, batchDenom)
		if err != nil {
			return nil, err
		}
		creditsToDeposit, err := regenmath.NewPositiveFixedDecFromString(credit.TradableAmount, maxDec)
		if err != nil {
			return nil, err
		}

		totalCreditsDeposited, err = totalCreditsDeposited.Add(creditsToDeposit)
		if err != nil {
			return nil, err
		}

		creditToAddToBasket := &ecocredit.MsgSend_SendCredits{
			BatchDenom:         credit.BatchDenom,
			TradableAmount:     credit.TradableAmount,
			RetiredAmount:      "",
			RetirementLocation: "",
		}

		// derive the address of the basket
		basketKey := BasketCreditsKey(basketDenomT(basket.BasketDenom), batchDenom)
		derivedKey := s.storeKey.Derive(basketKey)

		// send the credits do the derived address
		if _, err := s.Send(goCtx, &ecocredit.MsgSend{
			Sender:    owner.String(),
			Recipient: derivedKey.Address().String(),
			Credits:   []*ecocredit.MsgSend_SendCredits{creditToAddToBasket},
		}); err != nil {
			return nil, err
		}

	}

	// TODO(Tyler): should this return sdk.Dec? Is regen's dec okay for x/bank?
	// calculate total basket tokens to be awarded to the depositor
	basketTokensAmt, err := calculateBasketTokens(totalCreditsDeposited, basket.Exponent)
	if err != nil {
		return nil, err
	}
	basketTokensInt, err := basketTokensAmt.Int64()
	if err != nil {
		return nil, err
	}

	// mint basket tokens to send to basket depositor
	basketCoins := types.NewCoin(basket.BasketDenom, types.NewInt(basketTokensInt))
	if err = s.bankKeeper.MintCoins(ctx, ecocredit.ModuleName, types.Coins{basketCoins}); err != nil {
		return nil, err
	}

	// send the basket tokens to the basket depositor
	if err = s.bankKeeper.SendCoinsFromModuleToAccount(ctx, ecocredit.ModuleName, owner, types.Coins{basketCoins}); err != nil {
		return nil, err
	}

	return &ecocredit.MsgAddToBasketResponse{AmountReceived: basketTokensAmt.String()}, nil
}

// TakeFromBasket will take the oldest credit from the batch
// TODO(Tyler): do we need to more actively check for batch precision?
// TODO(Tyler): the response only indicates tradable amounts. Should we add retired amounts here?
func (s serverImpl) TakeFromBasket(goCtx context.Context, req *ecocredit.MsgTakeFromBasket) (*ecocredit.MsgTakeFromBasketResponse, error) {
	// setup vars
	sdkCtx := types.UnwrapSDKContext(goCtx)
	owner, _ := types.AccAddressFromBech32(req.Owner)
	store := sdkCtx.KVStore(s.storeKey)
	basketDenom := basketDenomT(req.BasketDenom)

	// get the basket
	var basket ecocredit.Basket
	if err := s.basketTable.GetOne(sdkCtx, orm.RowID(req.BasketDenom), &basket); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	// we fail fast in the event they didn't provide a retirement location when this basket requires retiring on swaps
	if !basket.DisableAutoRetire && req.RetirementLocation == "" {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("basket %s has auto-retirement enabled, but the request did not include a retirement location.", basket.BasketDenom)
	}

	// get the basket token balance of the caller in dec form
	userTokenBalance := s.bankKeeper.GetBalance(sdkCtx, owner, basket.BasketDenom)
	userBalanceDec, err := regenmath.NewDecFromString(userTokenBalance.Amount.String())
	if err != nil {
		return nil, err
	}

	// calculate how many basket tokens the user will need to fulfil the requested amount of credits
	requestedCreditAmount, _ := regenmath.NewDecFromString(req.Amount)
	tokensRequiredDec, err := calculateBasketTokens(requestedCreditAmount, basket.Exponent)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	// check they have enough basket tokens to complete this transaction
	if userBalanceDec.Cmp(tokensRequiredDec) == -1 {
		return nil, sdkerrors.ErrInsufficientFunds.Wrapf("insufficient basket token balance, got: %s, needed at least: %s", userBalanceDec.String(), tokensRequiredDec.String())
	}

	// put all the credits from this basket into a slice
	prefix := BasketCreditsKey(basketDenom, "")
	it := types.KVStorePrefixIterator(store, prefix)
	credits := make([]struct {
		startTime time.Time
		amount    string
		denom     string
	}, 0)
	for ; it.Valid(); it.Next() {
		// get the denom and deconstruct it
		creditDenom := string(it.Key())
		deconstructedDenom, err := ecocredit.DeconstructDenom(creditDenom)
		if err != nil {
			return nil, err
		}

		// get the start date and parse it into a time object
		batchStartDate := deconstructedDenom[1]
		t, err := time.Parse(ecocredit.TimeLayout, batchStartDate)
		if err != nil {
			return nil, err
		}

		// add the credit to the slice
		credits = append(credits, struct {
			startTime time.Time
			amount    string
			denom     string
		}{
			startTime: t,
			amount:    string(it.Value()),
			denom:     creditDenom,
		})
	}

	// close the iterator
	if err = it.Close(); err != nil {
		return nil, err
	}

	// sort the slice based on start time (we want to take the OLDEST credits first)
	sort.Slice(credits, func(i int, j int) bool {
		return credits[i].startTime.After(credits[j].startTime)
	})

	// response slice
	basketCreditsSent := make([]*ecocredit.BasketCredit, 0)

	// begin sending the credits
	// keep track of how many credits we need to send and update each iteration
	creditsNeeded := requestedCreditAmount
	for _, credit := range credits {

		// get the basket's credit amount in dec
		creditAmtDec, err := regenmath.NewDecFromString(credit.amount)
		if err != nil {
			return nil, err
		}

		// check how much we need to send from this batch
		var sendAmtStr string
		if creditAmtDec.Cmp(creditsNeeded) == -1 { // if theres not enough of this credit batch to fill the entire order,
			sendAmtStr = credit.amount                           // transfer all from this batch and move on to the next
			creditsNeeded, err = creditsNeeded.Sub(creditAmtDec) // update the needed credits
			if err != nil {
				return nil, err
			}
		} else { // the credits in the batch are either equal to or greater than the needed credits, so we just take the creditsNeeded amount and end.
			sendAmtStr = creditsNeeded.String()
			creditsNeeded = regenmath.NewDecFromInt64(0)
		}

		// fill in either tradable or retired amounts based on the basket settings
		creditsToSend := ecocredit.MsgSend_SendCredits{BatchDenom: credit.denom}
		if !basket.DisableAutoRetire {
			creditsToSend.RetiredAmount = sendAmtStr
			creditsToSend.RetirementLocation = req.RetirementLocation
		} else {
			creditsToSend.TradableAmount = sendAmtStr

			// add this to the response slice
			basketCreditsSent = append(basketCreditsSent, &ecocredit.BasketCredit{
				BatchDenom:     credit.denom,
				TradableAmount: sendAmtStr,
			})
		}

		// derive the address for this basket and send from basket to the caller
		basketKey := BasketCreditsKey(basketDenomT(basket.BasketDenom), batchDenomT(credit.denom))
		derivedKey := s.storeKey.Derive(basketKey)
		if _, err = s.Send(goCtx, &ecocredit.MsgSend{
			Sender:    derivedKey.Address().String(),
			Recipient: req.Owner,
			Credits:   []*ecocredit.MsgSend_SendCredits{&creditsToSend},
		}); err != nil {
			return nil, err
		}

		// if we don't need anymore credits, we can break
		if creditsNeeded.IsZero() {
			break
		}
	}

	// burn the basket tokens
	amountI64, err := requestedCreditAmount.Int64()
	if err != nil {
		return nil, err
	}
	basketToken := types.NewCoin(basket.BasketDenom, types.NewInt(amountI64))
	if err = s.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, owner, ecocredit.ModuleName, types.NewCoins(basketToken)); err != nil {
		return nil, err
	}
	if err = s.bankKeeper.BurnCoins(sdkCtx, ecocredit.ModuleName, types.NewCoins(basketToken)); err != nil {
		return nil, err
	}

	// TODO(Tyler): should we cache the total amount of credits, regardless of batch, in the basket table? that way we can fail fast and avoid a lot of unnecessary calculations.
	// check if we ended on zero. if it was not zero, the swap could not be fully executed, so we should error out.
	if !creditsNeeded.IsZero() {
		return nil, sdkerrors.ErrInsufficientFunds.Wrap("the basket does not have enough credits to complete this transaction")
	}

	return &ecocredit.MsgTakeFromBasketResponse{Credits: basketCreditsSent}, nil
}

// PickFromBasket allows picking a specific ecocredit from a basket.
func (s serverImpl) PickFromBasket(goCtx context.Context, req *ecocredit.MsgPickFromBasket) (*ecocredit.MsgPickFromBasketResponse, error) {
	// setup
	sdkCtx := types.UnwrapSDKContext(goCtx)
	owner, _ := types.AccAddressFromBech32(req.Owner)

	// get the basket
	var basket ecocredit.Basket
	if err := s.basketTable.GetOne(sdkCtx, orm.RowID(req.BasketDenom), &basket); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	// fail fast if they didn't specify a retirement location for an auto-retirement basket
	if !basket.DisableAutoRetire && req.RetirementLocation == "" {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("basket %s has auto-retirement enabled, but the request did not include a retirement location.", basket.BasketDenom)
	}

	// get the basket token balance of the requester
	basketTokenBalanceStr := s.bankKeeper.GetBalance(sdkCtx, owner, basket.BasketDenom).String()
	basketTokenBalance, err := regenmath.NewDecFromString(basketTokenBalanceStr)
	if err != nil {
		return nil, err
	}

	// TODO(tyler): should the basket curator be able to pick even if its disabled?
	if !basket.AllowPicking { // can only pick if the basket allows it!
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("basket %s does not allow picking", basket.BasketDenom)
	} else {
		tokensOwed := regenmath.NewDecFromInt64(0)
		for _, credit := range req.Credits {
			//prefix := BasketCreditsKey(basketDenom, batchDenomT(credit.BatchDenom))
			creditAmtRequested, _ := regenmath.NewDecFromString(credit.TradableAmount)

			// get the basket's balance of the requested credit
			basketKey := BasketCreditsKey(basketDenomT(basket.BasketDenom), batchDenomT(credit.BatchDenom))
			derivedKey := s.storeKey.Derive(basketKey)
			res, err := s.Balance(goCtx, &ecocredit.QueryBalanceRequest{
				Account:    derivedKey.Address().String(),
				BatchDenom: credit.BatchDenom,
			})
			if err != nil {
				return nil, err
			}
			basketBalance, _ := regenmath.NewDecFromString(res.TradableAmount)

			// TODO(Tyler): should we partial fill?
			// check if the basket has the credit balance to support this tx
			if creditAmtRequested.Cmp(basketBalance) == 1 { // if the requested credits is more than what's in the basket..
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("requested %s credits but basket %s only has %s credits from batch %s", credit.TradableAmount, basket.BasketDenom, res.TradableAmount, credit.BatchDenom)
			}

			// calculate the token cost for this specific credit
			requiredTokens, err := calculateBasketTokens(creditAmtRequested, basket.Exponent)
			if err != nil {
				return nil, err
			}

			// add it to the overall cost of this tx
			tokensOwed, err = tokensOwed.Add(requiredTokens)
			if err != nil {
				return nil, err
			}

			// check to see if their balance can handle it
			basketTokenBalance, err = basketTokenBalance.Sub(tokensOwed)
			if err != nil {
				return nil, err
			}
			if basketTokenBalance.IsNegative() {
				return nil, sdkerrors.ErrInsufficientFunds.Wrapf("transaction failed after calculating tokens required for credit batch %s", credit.BatchDenom)
			}

			// send the credit from the basket to the requester
			msgSendCredits := &ecocredit.MsgSend_SendCredits{BatchDenom: credit.BatchDenom}
			if basket.DisableAutoRetire {
				msgSendCredits.TradableAmount = credit.TradableAmount
			} else {
				msgSendCredits.RetiredAmount = credit.TradableAmount
				msgSendCredits.RetirementLocation = req.RetirementLocation
			}
			send := &ecocredit.MsgSend{
				Sender:    derivedKey.Address().String(),
				Recipient: req.Owner,
				Credits:   []*ecocredit.MsgSend_SendCredits{msgSendCredits},
			}
			if _, err = s.Send(goCtx, send); err != nil {
				return nil, err
			}
		}

		// burn the coins
		ti64, err := tokensOwed.Int64()
		if err != nil {
			return nil, err
		}
		// TODO(Tyler) should we send to module then burn? or should we just send to some sort of 0 address? what to do?
		tokenCost := types.NewCoin(basket.BasketDenom, types.NewInt(ti64))
		if err := s.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, owner, ecocredit.ModuleName, types.NewCoins(tokenCost)); err != nil {
			return nil, err
		}
		if err := s.bankKeeper.BurnCoins(sdkCtx, ecocredit.ModuleName, types.NewCoins(tokenCost)); err != nil {
			return nil, err
		}

	}
	return &ecocredit.MsgPickFromBasketResponse{}, nil
}

// ----- HELPER METHODS -----

// validateFilterData is a recursive, stateful filter validation.
// it ensures all filters relative to other state (classes, batches, projects, etc) in the blockchain are valid.
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
		case *ecocredit.Filter_ProjectId:
			if exists := s.projectInfoTable.Has(ctx, orm.RowID(f.ProjectId)); !exists {
				return sdkerrors.ErrNotFound.Wrapf("project with id %s not found", f.ProjectId)
			}
		}
	}
	return nil
}

// padRight is a helper function to construct a string of length
// `length` with a prefix, and a given `add` string, which serves
// as the string to continuously add until len(string) == length
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
func calculateBasketTokens(credits regenmath.Dec, exponent uint32) (regenmath.Dec, error) {
	// get the str to use in the multiplier
	multiStr := padRight(int(exponent), "10", "0")

	// get the multiplier in dec form
	multiplier, err := regenmath.NewPositiveDecFromString(multiStr)
	if err != nil {
		return regenmath.Dec{}, err
	}

	// return the credits deposited * 10^(exponent)
	return credits.Mul(multiplier)
}

// checkFilters recursively checks filters using `depth` to ensure valid filters.
// it sets the depth equal the length of the filter slice, and subtracts by 1 for each valid filter encountered.
// for AND filters, we require the depth to return 0, as each filter in the slice should subtract 1.
// for OR filters, we simply require that the slice of it's inner filter is not equal to the depth returned.
// this is because we only need ONE tree to be valid, thus, an invalid OR tree would be if 0 filters passed.
// TODO(Tyler): should we enforce a depth limit on OR/AND filters?
func checkFilters(filters []*ecocredit.Filter, classInfo ecocredit.ClassInfo, batchInfo ecocredit.BatchInfo, projectInfo ecocredit.ProjectInfo, owner string) (int, error) {
	depth := len(filters)
	var err error
	for _, filter := range filters {
		switch f := filter.Sum.(type) {
		case *ecocredit.Filter_And_:
			andFilter := f.And.Filters
			innerDepth, err := checkFilters(andFilter, classInfo, batchInfo, projectInfo, owner)
			if innerDepth != 0 || err != nil {
				return innerDepth, sdkerrors.ErrInvalidRequest.Wrap("invalid AND filter")
			} else {
				depth -= 1
			}
		case *ecocredit.Filter_Or_:
			orFilter := f.Or.Filters
			orDepth := len(orFilter)
			innerDepth, err := checkFilters(orFilter, classInfo, batchInfo, projectInfo, owner)

			// when orDepth == innerDepth, none of the filters in the OR got a match. we need AT LEAST 1 match for a valid OR filter.
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
			if classInfo.ClassId == f.ClassId {
				depth -= 1
			} else {
				err = formatFilterError("class id", f.ClassId, classInfo.ClassId)
			}
		case *ecocredit.Filter_ProjectId:
			if f.ProjectId == projectInfo.ProjectId {
				depth -= 1
			} else {
				err = formatFilterError("project id", f.ProjectId, projectInfo.ProjectId)
			}
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
			found := false
			for _, issuer := range classInfo.Issuers {
				if f.Issuer == issuer {
					depth -= 1
					found = true
					break
				}
			}
			if !found {
				err = fmt.Errorf("credit class %s does not contain issuer %s", classInfo.ClassId, f.Issuer)
			}
		case *ecocredit.Filter_Owner:
			if owner == f.Owner {
				depth -= 1
			} else {
				err = formatFilterError("credit owner", f.Owner, owner)
			}
		case *ecocredit.Filter_ProjectLocation:
			if f.ProjectLocation == projectInfo.ProjectLocation {
				depth -= 1
			} else {
				err = formatFilterError("project location", f.ProjectLocation, projectInfo.ProjectLocation)
			}
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
		//case *ecocredit.Filter_Tag:
		// depth -= 1 TODO: wait for tags PR
		default:
			err = errors.New("no valid filter given")
		}

	}

	if err != nil {
		return depth, err
	}
	if depth != 0 {
		return depth, fmt.Errorf("the filter could not be matched with depth %d", depth)
	}
	return depth, nil
}

// formatFilterError is a helper method for formatting filter errors in a repeatable fashion.
func formatFilterError(item, want, got string) error {
	return fmt.Errorf("basket filter requires %s %s, but a credit with %s %s was given", item, got, item, want)
}

// constructBasketDenom constructs the denom for a basket token. it takes the form of `ecocredit:basketName`
func constructBasketDenom(name string) string {
	return fmt.Sprintf("%s:%s", ecocredit.ModuleName, name)
}
