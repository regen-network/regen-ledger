package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"
	regentypes "github.com/regen-network/regen-ledger/types"
	regenmath "github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"math"
	"sort"
	"time"
)

// CreateBasket creates a basket keyed by a given basketDenom.
func (s serverImpl) CreateBasket(goCtx context.Context, req *ecocredit.MsgCreateBasket) (*ecocredit.MsgCreateBasketResponse, error) {
	ctx := regentypes.UnwrapSDKContext(goCtx)
	sdkCtx := sdktypes.UnwrapSDKContext(goCtx)

	// basket denom = name of basket? TODO(Tyler)
	basketDenom := req.Name

	// check if a basket with this name already exists
	var basket ecocredit.Basket
	err := s.basketTable.GetOne(ctx, orm.RowID(basketDenom), &basket)
	if err == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("basket with name %s already exists", req.Name)
	}

	// TODO(Tyler): generate this
	displayName := ""

	ct, err := s.getCreditType(sdkCtx, req.CreditTypeName)
	if err != nil {
		return nil, err
	}
	if req.Exponent < ct.Precision {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("creating a basket with credit type %s requires exponent >= %d", req.CreditTypeName, ct.Precision)
	}

	// stateful validation of basket criteria. checks that the filters actually apply to existing types.
	if err := s.validateFilterData(ctx, req.BasketCriteria); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("invalid basket filter: %s", err.Error())
	}

	if err := s.basketTable.Create(ctx, &ecocredit.Basket{
		BasketDenom:       basketDenom,
		Curator:           req.Curator,
		Name:              req.Name,
		DisplayName:       displayName,
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
// TODO(Tyler): this method mints tokens based on the basket name. some issues:
// if someone names a basket OSMO, that could be confusing on chain.
// should we instead use something akin to wrapped token names?
// i.e. (wBTC for wrapped bitcoin -> bsktBasketName for basket tokens)?
func (s serverImpl) AddToBasket(goCtx context.Context, req *ecocredit.MsgAddToBasket) (*ecocredit.MsgAddToBasketResponse, error) {
	ctx := sdktypes.UnwrapSDKContext(goCtx)
	regenCtx := regentypes.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	owner, _ := sdktypes.AccAddressFromBech32(req.Owner)

	// get basket info
	var basket ecocredit.Basket
	err := s.basketTable.GetOne(ctx, orm.RowID(req.BasketDenom), &basket)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("basket %s not found", req.BasketDenom)
	}

	totalCreditsDeposited := regenmath.NewDecFromInt64(0)
	for _, credit := range req.Credits {

		// get the batch
		batchDenom := batchDenomT(credit.BatchDenom)
		var batchInfo ecocredit.BatchInfo
		err := s.batchInfoTable.GetOne(ctx, orm.RowID(batchDenom), &batchInfo)
		if err != nil {
			return nil, sdkerrors.ErrNotFound.Wrapf("batch %s not found", credit.BatchDenom)
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

		// verify the user has sufficient ecocredits to send
		err = verifyCreditBalance(store, owner, credit.BatchDenom, credit.TradableAmount)
		if err != nil {
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

		// TODO(Tyler): enforce criteria
		if basket.BasketCriteria != nil {
			// check that the credits meet the filter
			if err = checkFilterMatch(basket.BasketCriteria, *classInfo, batchInfo, *projectInfo, req.Owner); err != nil {
				return nil, err
			}
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

		if err = s.depositCreditToBasket(regenCtx, store, owner, basket, creditToAddToBasket); err != nil {
			return nil, err
		}

	}

	// calculate total basket tokens to be awarded to the depositor
	basketTokensAmt, err := CalculateBasketTokens(totalCreditsDeposited, basket.Exponent)
	if err != nil {
		return nil, err
	}
	basketTokensInt, err := basketTokensAmt.Int64()
	if err != nil {
		return nil, err
	}

	// mint basket tokens
	basketCoins := sdktypes.NewCoin(basket.BasketDenom, sdktypes.NewInt(basketTokensInt))
	if err = s.bankKeeper.MintCoins(ctx, ecocredit.ModuleName, sdktypes.Coins{basketCoins}); err != nil {
		return nil, err
	}

	// send the basket tokens to the basket depositor
	if err = s.bankKeeper.SendCoinsFromModuleToAccount(ctx, ecocredit.ModuleName, owner, sdktypes.Coins{basketCoins}); err != nil {
		return nil, err
	}

	return &ecocredit.MsgAddToBasketResponse{AmountReceived: basketTokensAmt.String()}, nil
}

// TakeFromBasket will take the oldest credit from the batch
// TODO(Tyler): the response only indicates tradable amounts. Should we add retired amounts here?
func (s serverImpl) TakeFromBasket(goCtx context.Context, req *ecocredit.MsgTakeFromBasket) (*ecocredit.MsgTakeFromBasketResponse, error) {
	// setup vars
	regenCtx := regentypes.UnwrapSDKContext(goCtx)
	sdkCtx := sdktypes.UnwrapSDKContext(goCtx)
	owner, _ := sdktypes.AccAddressFromBech32(req.Owner)
	store := sdkCtx.KVStore(s.storeKey)

	// get the basket
	var basket ecocredit.Basket
	if err := s.basketTable.GetOne(sdkCtx, orm.RowID(req.BasketDenom), &basket); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("unable to get basket with denom %s: %s", req.BasketDenom, err.Error())
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
	tokensRequiredDec, err := CalculateBasketTokens(requestedCreditAmount, basket.Exponent)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	// check they have enough basket tokens to complete this transaction
	if userBalanceDec.Cmp(tokensRequiredDec) == -1 {
		return nil, sdkerrors.ErrInsufficientFunds.Wrapf("insufficient basket token balance, got: %s, needed at least: %s", userBalanceDec.String(), tokensRequiredDec.String())
	}

	creditsInBasket := make([]struct {
		startTime time.Time
		amount    string
		denom     string
	}, 0)

	it := s.basketCreditsIterator(basket, store)

	// loop over the balances and store them in a slice
	for ; it.Valid(); it.Next() {
		// get the denom and deconstruct it
		_, creditDenom := ParseBalanceKey(it.Key())
		deconstructedDenom, err := ecocredit.DeconstructBatchDenom(string(creditDenom))
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
		creditsInBasket = append(creditsInBasket, struct {
			startTime time.Time
			amount    string
			denom     string
		}{
			startTime: t,
			amount:    string(it.Value()),
			denom:     string(creditDenom),
		})
	}

	// close the iterator
	if err = it.Close(); err != nil {
		return nil, err
	}

	// sort the slice based on start time (we want to take the OLDEST credits first)
	sort.Slice(creditsInBasket, func(i int, j int) bool {
		return creditsInBasket[i].startTime.Before(creditsInBasket[j].startTime)
	})

	// response slice
	basketCreditsSent := make([]*ecocredit.BasketCredit, 0)

	// begin sending the credits
	// keep track of how many credits we need to send and update each iteration
	creditsNeeded := requestedCreditAmount
	for _, credit := range creditsInBasket {

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

		if err := s.sendCreditFromBasket(regenCtx, store, owner, basket, &creditsToSend); err != nil {
			return nil, err
		}

		// if we don't need anymore credits, we can break
		if creditsNeeded.IsZero() {
			break
		}
	}

	// TODO(Tyler): check bank supply to ensure this tx can be fulfilled
	// check if we ended on zero. if it was not zero, the swap could not be fully executed, so we should error out.
	if !creditsNeeded.IsZero() {
		return nil, sdkerrors.ErrInsufficientFunds.Wrap("the basket does not have enough credits to complete this transaction")
	}

	// burn the basket coins
	amountI64, err := tokensRequiredDec.Int64()
	if err != nil {
		return nil, err
	}
	basketToken := sdktypes.NewCoin(basket.BasketDenom, sdktypes.NewInt(amountI64))
	if err = s.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, owner, ecocredit.ModuleName, sdktypes.NewCoins(basketToken)); err != nil {
		return nil, err
	}
	if err = s.bankKeeper.BurnCoins(sdkCtx, ecocredit.ModuleName, sdktypes.NewCoins(basketToken)); err != nil {
		return nil, err
	}

	return &ecocredit.MsgTakeFromBasketResponse{Credits: basketCreditsSent}, nil
}

// PickFromBasket allows picking a specific ecocredit from a basket.
func (s serverImpl) PickFromBasket(goCtx context.Context, req *ecocredit.MsgPickFromBasket) (*ecocredit.MsgPickFromBasketResponse, error) {
	// setup
	sdkCtx := sdktypes.UnwrapSDKContext(goCtx)
	regenCtx := regentypes.UnwrapSDKContext(goCtx)
	owner, _ := sdktypes.AccAddressFromBech32(req.Owner)
	store := regenCtx.KVStore(s.storeKey)

	// get the basket
	var basket ecocredit.Basket
	if err := s.basketTable.GetOne(sdkCtx, orm.RowID(req.BasketDenom), &basket); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get basket with denom %s: %s", req.BasketDenom, err.Error())
	}

	// fail fast if they didn't specify a retirement location for an auto-retirement basket
	if !basket.DisableAutoRetire && req.RetirementLocation == "" {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("basket %s has auto-retirement enabled, but the request did not include a retirement location.", basket.BasketDenom)
	}

	// get the basket token balance of the requester
	tokenBalance := s.bankKeeper.GetBalance(sdkCtx, owner, basket.BasketDenom).Amount.String()
	basketTokenBalance, err := regenmath.NewDecFromString(tokenBalance)
	if err != nil {
		return nil, err
	}

	if !basket.AllowPicking { // can only pick if the basket allows it!
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("basket %s does not allow picking", basket.BasketDenom)
	} else {
		tokensOwed := regenmath.NewDecFromInt64(0)
		for _, credit := range req.Credits {
			//prefix := BasketCreditsKey(basketDenom, batchDenomT(credit.BatchDenom))
			creditAmtRequested, _ := regenmath.NewDecFromString(credit.TradableAmount)

			// calculate the token cost for this specific credit
			requiredTokens, err := CalculateBasketTokens(creditAmtRequested, basket.Exponent)
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
				msgSendCredits.RetiredAmount = "0"
			} else {
				msgSendCredits.TradableAmount = "0"
				msgSendCredits.RetiredAmount = credit.TradableAmount
				msgSendCredits.RetirementLocation = req.RetirementLocation
			}

			// send credits from the basket to the user
			if err := s.sendCreditFromBasket(regenCtx, store, owner, basket, msgSendCredits); err != nil {
				return nil, err
			}
		}

		// burn the coins
		ti64, err := tokensOwed.Int64()
		if err != nil {
			return nil, err
		}
		tokenCost := sdktypes.NewCoin(basket.BasketDenom, sdktypes.NewInt(ti64))
		if err := s.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, owner, ecocredit.ModuleName, sdktypes.NewCoins(tokenCost)); err != nil {
			return nil, err
		}
		if err := s.bankKeeper.BurnCoins(sdkCtx, ecocredit.ModuleName, sdktypes.NewCoins(tokenCost)); err != nil {
			return nil, err
		}
	}
	return &ecocredit.MsgPickFromBasketResponse{}, nil
}

func (s serverImpl) basketCreditsIterator(basket ecocredit.Basket, store sdktypes.KVStore) sdktypes.Iterator {
	// get the iterator to scan all balances
	key := BasketIteratorKey(basketDenomT(basket.BasketDenom))
	return types.KVStorePrefixIterator(store, key)
}

// depositCreditToBasket deposits a set of credits to a basket
func (s serverImpl) depositCreditToBasket(ctx regentypes.Context, store sdktypes.KVStore, senderAddr sdktypes.AccAddress, basket ecocredit.Basket, credit *ecocredit.MsgSend_SendCredits) error {
	return s.sendEcocredits(ctx, credit, store, EcocreditEOA(senderAddr), basketDenomT(basket.BasketDenom))
}

// sendCreditFromBasket sends credits from basket to the `to` address
func (s serverImpl) sendCreditFromBasket(regenCtx regentypes.Context, store sdktypes.KVStore, to sdktypes.AccAddress, basket ecocredit.Basket, credit *ecocredit.MsgSend_SendCredits) error {
	return s.sendEcocredits(regenCtx, credit, store, basketDenomT(basket.BasketDenom), EcocreditEOA(to))
}

// validateFilterData is a recursive, stateful filter validation.
// it ensures all filters relative to other state (classes, batches, projects, etc) in the blockchain are valid.
func (s serverImpl) validateFilterData(ctx regentypes.Context, filters ...*ecocredit.Filter) error {
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

// CalculateBasketTokens calculates the basket tokens to be awarded based on how many ecocredits were added to the basket.
// the equation for calculating the award amount is as follows:
// total_credits_deposited * 10^(basket.Exponent)
// TODO(Tyler): not too convinced on this function's name...
func CalculateBasketTokens(credits regenmath.Dec, exponent uint32) (regenmath.Dec, error) {
	// calculate the multiplier and convert to string
	multi := math.Pow(10, float64(exponent))
	multiStr := fmt.Sprintf("%f", multi)

	// get the multiplier in dec form
	multiplier, err := regenmath.NewPositiveDecFromString(multiStr)
	if err != nil {
		return regenmath.Dec{}, err
	}

	// return the credits deposited * 10^(exponent)
	return credits.Mul(multiplier)
}

// checkCreditMatchesFilter recursively checks filters using `depth` to ensure valid filters.
// it sets the depth equal the length of the filter slice, and subtracts by 1 for each valid filter encountered.
// for AND filters, we require the depth to return 0, as each filter in the slice should subtract 1.
// for OR filters, we simply require that the slice of it's inner filter is not equal to the depth returned.
// this is because we only need ONE tree to be valid, thus, an invalid OR tree would be if 0 filters passed.
// TODO(Tyler): assume depth limit is enforced here already
func checkFilterMatch(filter *ecocredit.Filter, classInfo ecocredit.ClassInfo, batchInfo ecocredit.BatchInfo, projectInfo ecocredit.ProjectInfo, owner string) error {
	switch f := filter.Sum.(type) {
	case *ecocredit.Filter_And_:
		for _, f := range f.And.Filters {
			err := checkFilterMatch(f, classInfo, batchInfo, projectInfo, owner)
			if err != nil {
				return err
			}
		}
	case *ecocredit.Filter_Or_:
		var err1 error
		for _, f := range f.Or.Filters {
			err := checkFilterMatch(f, classInfo, batchInfo, projectInfo, owner)
			if err == nil {
				return nil
			} else {
				err1 = err
			}
		}
		return err1
	case *ecocredit.Filter_ClassId:
		if classInfo.ClassId == f.ClassId {
			return nil
		} else {
			return formatFilterError("class id", f.ClassId, classInfo.ClassId)
		}
	case *ecocredit.Filter_ProjectId:
		if f.ProjectId == projectInfo.ProjectId {
			return nil
		} else {
			return formatFilterError("project id", f.ProjectId, projectInfo.ProjectId)
		}
	case *ecocredit.Filter_BatchDenom:
		if batchInfo.BatchDenom == f.BatchDenom {
			return nil
		} else {
			return formatFilterError("batch denom", f.BatchDenom, batchInfo.BatchDenom)
		}
	case *ecocredit.Filter_ClassAdmin:
		if classInfo.Admin == f.ClassAdmin {
			return nil
		} else {
			return formatFilterError("class admin", f.ClassAdmin, classInfo.Admin)
		}
	case *ecocredit.Filter_Issuer:
		for _, issuer := range classInfo.Issuers {
			if f.Issuer == issuer {
				return nil
			}
		}
		return fmt.Errorf("credit class %s does not contain issuer %s", classInfo.ClassId, f.Issuer)

	case *ecocredit.Filter_Owner:
		if owner == f.Owner {
			return nil
		} else {
			return formatFilterError("credit owner", f.Owner, owner)
		}
	case *ecocredit.Filter_ProjectLocation:
		if f.ProjectLocation == projectInfo.ProjectLocation {
			return nil
		} else {
			return formatFilterError("project location", f.ProjectLocation, projectInfo.ProjectLocation)
		}
	case *ecocredit.Filter_DateRange_:
		if batchInfo.StartDate.Equal(*f.DateRange.StartDate) || batchInfo.StartDate.After(*f.DateRange.StartDate) {
			if batchInfo.EndDate.Equal(*f.DateRange.EndDate) || batchInfo.EndDate.Before(*f.DateRange.EndDate) {
				return nil
			} else {
				return formatFilterError("date range", f.DateRange.StartDate.String()+" to "+f.DateRange.EndDate.String(), batchInfo.StartDate.String()+" to "+batchInfo.EndDate.String())
			}
		} else {
			return formatFilterError("date range", f.DateRange.StartDate.String()+" to "+f.DateRange.EndDate.String(), batchInfo.StartDate.String()+" to "+batchInfo.EndDate.String())
		}
	//case *ecocredit.Filter_Tag:
	// depth -= 1 TODO: wait for tags PR
	default:
		return errors.New("no valid filter given")
	}

	return nil
}

// formatFilterError is a helper method for formatting filter errors in a repeatable fashion.
func formatFilterError(item, want, got string) error {
	return fmt.Errorf("basket filter requires %s %s, but a credit with %s %s was given", item, want, item, got)
}
