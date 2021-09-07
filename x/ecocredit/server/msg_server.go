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

// CreateClass creates a new class of ecocredit
//
// The admin is charged a fee for creating the class. This is controlled by
// the global parameter CreditClassFee, which can be updated through the
// governance process.
func (s serverImpl) CreateClass(goCtx context.Context, req *ecocredit.MsgCreateClass) (*ecocredit.MsgCreateClassResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	// Charge the admin a fee to create the credit class
	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx.Context, &params)
	if params.AllowlistEnabled && !s.isCreatorAllowListed(params.AllowedClassCreators, adminAddress) {
		return nil, fmt.Errorf("%s is not allowed to create credit classes", adminAddress.String())
	}

	err = s.chargeCreditClassFee(ctx.Context, adminAddress)
	if err != nil {
		return nil, err
	}

	creditType, err := s.getCreditType(ctx.Context, req.CreditTypeName)
	if err != nil {
		return nil, err
	}

	classSeqNo, err := s.getCreditTypeSeqNextVal(ctx.Context, creditType)
	if err != nil {
		return nil, err
	}

	classID, err := ecocredit.FormatClassID(creditType, classSeqNo)
	if err != nil {
		return nil, err
	}

	err = s.classInfoTable.Create(ctx, &ecocredit.ClassInfo{
		ClassId:    classID,
		Admin:      req.Admin,
		Issuers:    req.Issuers,
		Metadata:   req.Metadata,
		CreditType: &creditType,
	})
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateClass{
		ClassId: classID,
		Admin:   req.Admin,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateClassResponse{ClassId: classID}, nil
}

// CreateBatch creates a new batch of credits.
// Credits in the batch must not have more decimal places than the credit type's specified precision.
func (s serverImpl) CreateBatch(goCtx context.Context, req *ecocredit.MsgCreateBatch) (*ecocredit.MsgCreateBatchResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	classID := req.ClassId
	classInfo, err := s.getClassInfo(ctx, classID)
	if err != nil {
		return nil, err
	}

	if err = classInfo.AssertClassIssuer(req.Issuer); err != nil {
		return nil, err
	}

	maxDecimalPlaces := classInfo.CreditType.Precision
	batchSeqNo, err := s.nextBatchInClass(ctx, classInfo)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	batchDenomStr, err := ecocredit.FormatDenom(classID, batchSeqNo, req.StartDate, req.EndDate)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	batchDenom := batchDenomT(batchDenomStr)
	tradableSupply := math.NewDecFromInt64(0)
	retiredSupply := math.NewDecFromInt64(0)

	store := ctx.KVStore(s.storeKey)

	for _, issuance := range req.Issuance {
		var err error
		tradable, retired := math.NewDecFromInt64(0), math.NewDecFromInt64(0)

		if issuance.TradableAmount != "" {
			tradable, err = math.NewNonNegativeDecFromString(issuance.TradableAmount)
			if err != nil {
				return nil, err
			}

			decPlaces := tradable.NumDecimalPlaces()
			if decPlaces > maxDecimalPlaces {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("tradable amount exceeds precision for credit type: "+
					"is %v, should be < %v", decPlaces, maxDecimalPlaces)
			}
		}

		if issuance.RetiredAmount != "" {
			retired, err = math.NewNonNegativeDecFromString(issuance.RetiredAmount)
			if err != nil {
				return nil, err
			}

			decPlaces := retired.NumDecimalPlaces()
			if decPlaces > maxDecimalPlaces {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("retired amount does not conform to credit type "+
					"precision: %v should be %v", decPlaces, maxDecimalPlaces)
			}
		}

		recipient := issuance.Recipient
		recipientAddr, err := sdk.AccAddressFromBech32(recipient)
		if err != nil {
			return nil, err
		}

		if !tradable.IsZero() {
			tradableSupply, err = tradableSupply.Add(tradable)
			if err != nil {
				return nil, err
			}

			err = addAndSetDecimal(store, TradableBalanceKey(recipientAddr, batchDenom), tradable)
			if err != nil {
				return nil, err
			}
		}

		if !retired.IsZero() {
			retiredSupply, err = retiredSupply.Add(retired)
			if err != nil {
				return nil, err
			}

			err = retire(ctx, store, recipientAddr, batchDenom, retired, issuance.RetirementLocation)
			if err != nil {
				return nil, err
			}
		}

		sum, err := tradable.Add(retired)
		if err != nil {
			return nil, err
		}

		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Recipient:  recipient,
			BatchDenom: string(batchDenom),
			Amount:     sum.String(),
		})
		if err != nil {
			return nil, err
		}
	}

	setDecimal(store, TradableSupplyKey(batchDenom), tradableSupply)
	setDecimal(store, RetiredSupplyKey(batchDenom), retiredSupply)

	totalSupply, err := tradableSupply.Add(retiredSupply)
	if err != nil {
		return nil, err
	}
	totalSupplyStr := totalSupply.String()

	amountCancelledStr := math.NewDecFromInt64(0).String()

	err = s.batchInfoTable.Create(ctx, &ecocredit.BatchInfo{
		ClassId:         classID,
		BatchDenom:      string(batchDenom),
		Issuer:          req.Issuer,
		TotalAmount:     totalSupplyStr,
		Metadata:        req.Metadata,
		AmountCancelled: amountCancelledStr,
		StartDate:       req.StartDate,
		EndDate:         req.EndDate,
		ProjectLocation: req.ProjectLocation,
	})
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateBatch{
		ClassId:         classID,
		BatchDenom:      string(batchDenom),
		Issuer:          req.Issuer,
		TotalAmount:     totalSupplyStr,
		StartDate:       req.StartDate.Format("2006-01-02"),
		EndDate:         req.EndDate.Format("2006-01-02"),
		ProjectLocation: req.ProjectLocation,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateBatchResponse{BatchDenom: string(batchDenom)}, nil
}

// Send sends credits to a recipient.
// Send also retires credits if the amount to retire is specified in the request.
func (s serverImpl) Send(goCtx context.Context, req *ecocredit.MsgSend) (*ecocredit.MsgSendResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	sender := req.Sender
	recipient := req.Recipient

	senderAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return nil, err
	}

	recipientAddr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return nil, err
	}

	for _, credit := range req.Credits {
		denom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid credit batch denom", denom)
		}

		maxDecimalPlaces, err := s.getBatchPrecision(ctx, denom)
		if err != nil {
			return nil, err
		}

		tradable, err := math.NewNonNegativeFixedDecFromString(credit.TradableAmount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		retired, err := math.NewNonNegativeFixedDecFromString(credit.RetiredAmount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		sum, err := tradable.Add(retired)
		if err != nil {
			return nil, err
		}

		// subtract balance
		err = subAndSetDecimal(store, TradableBalanceKey(senderAddr, denom), sum)
		if err != nil {
			return nil, err
		}

		// Add tradable balance
		err = addAndSetDecimal(store, TradableBalanceKey(recipientAddr, denom), tradable)
		if err != nil {
			return nil, err
		}

		if !retired.IsZero() {
			// subtract retired from tradable supply
			err = subAndSetDecimal(store, TradableSupplyKey(denom), retired)
			if err != nil {
				return nil, err
			}

			// Add retired balance
			err = retire(ctx, store, recipientAddr, denom, retired, credit.RetirementLocation)
			if err != nil {
				return nil, err
			}

			// Add retired supply
			err = addAndSetDecimal(store, RetiredSupplyKey(denom), retired)
			if err != nil {
				return nil, err
			}
		}

		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Sender:     sender,
			Recipient:  recipient,
			BatchDenom: string(denom),
			Amount:     sum.String(),
		})
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgSendResponse{}, nil
}

// Retire credits to the specified location.
// WARNING: retiring credits is permanent. Retired credits cannot be un-retired.
func (s serverImpl) Retire(goCtx context.Context, req *ecocredit.MsgRetire) (*ecocredit.MsgRetireResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	holderAddr, err := sdk.AccAddressFromBech32(req.Holder)
	if err != nil {
		return nil, err
	}

	for _, credit := range req.Credits {
		denom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid credit batch denom", denom)
		}

		maxDecimalPlaces, err := s.getBatchPrecision(ctx, denom)
		if err != nil {
			return nil, err
		}

		toRetire, err := math.NewPositiveFixedDecFromString(credit.Amount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		err = subtractTradableBalanceAndSupply(store, holderAddr, denom, toRetire)
		if err != nil {
			return nil, err
		}

		//  Add retired balance
		err = retire(ctx, store, holderAddr, denom, toRetire, req.Location)
		if err != nil {
			return nil, err
		}

		//  Add retired supply
		err = addAndSetDecimal(store, RetiredSupplyKey(denom), toRetire)
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgRetireResponse{}, nil
}

// Cancel credits, removing them from the supply and balance of the holder
func (s serverImpl) Cancel(goCtx context.Context, req *ecocredit.MsgCancel) (*ecocredit.MsgCancelResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	holderAddr, err := sdk.AccAddressFromBech32(req.Holder)
	if err != nil {
		return nil, err
	}

	for _, credit := range req.Credits {

		// Check that the batch that were trying to cancel credits from
		// exists
		denom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid credit batch denom", denom)
		}

		// Remove the credits from the total_amount in the batch and add
		// them to amount_cancelled
		var batchInfo ecocredit.BatchInfo
		err := s.batchInfoTable.GetOne(ctx, orm.RowID(denom), &batchInfo)
		if err != nil {
			return nil, err
		}

		classInfo, err := s.getClassInfo(ctx, batchInfo.ClassId)
		if err != nil {
			return nil, err
		}

		maxDecimalPlaces := classInfo.CreditType.Precision

		// Parse the amount of credits to cancel, checking it conforms
		// to the precision
		toCancel, err := math.NewPositiveFixedDecFromString(credit.Amount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		// Remove the credits from the balance of the holder and the
		// overall supply
		err = subtractTradableBalanceAndSupply(store, holderAddr, denom, toCancel)
		if err != nil {
			return nil, err
		}

		totalAmount, err := math.NewPositiveFixedDecFromString(batchInfo.TotalAmount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		totalAmount, err = math.SafeSubBalance(totalAmount, toCancel)
		if err != nil {
			return nil, err
		}
		batchInfo.TotalAmount = totalAmount.String()

		amountCancelled, err := math.NewNonNegativeFixedDecFromString(batchInfo.AmountCancelled, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		amountCancelled, err = amountCancelled.Add(toCancel)
		if err != nil {
			return nil, err
		}
		batchInfo.AmountCancelled = amountCancelled.String()

		if err = s.batchInfoTable.Save(ctx, &batchInfo); err != nil {
			return nil, err
		}

		// Emit the cancellation event
		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCancel{
			Canceller:  req.Holder,
			BatchDenom: string(denom),
			Amount:     toCancel.String(),
		})
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgCancelResponse{}, nil
}

// nextBatchInClass gets the sequence number for the next batch in the credit
// class and updates the class info with the new batch number
func (s serverImpl) nextBatchInClass(ctx types.Context, classInfo *ecocredit.ClassInfo) (uint64, error) {
	// Get the next value
	nextVal := classInfo.NumBatches + 1

	// Update the ClassInfo
	classInfo.NumBatches = nextVal
	err := s.classInfoTable.Save(ctx, classInfo)
	if err != nil {
		return 0, err
	}

	return nextVal, nil
}

func retire(ctx types.Context, store sdk.KVStore, recipient sdk.AccAddress, batchDenom batchDenomT, retired math.Dec, location string) error {
	err := addAndSetDecimal(store, RetiredBalanceKey(recipient, batchDenom), retired)
	if err != nil {
		return err
	}

	return ctx.EventManager().EmitTypedEvent(&ecocredit.EventRetire{
		Retirer:    recipient.String(),
		BatchDenom: string(batchDenom),
		Amount:     retired.String(),
		Location:   location,
	})
}

// subtracts `amount` from the tradable balance and tradable supply
func subtractTradableBalanceAndSupply(store sdk.KVStore, holder sdk.AccAddress, batchDenom batchDenomT, amount math.Dec) error {
	// subtract tradable balance
	err := subAndSetDecimal(store, TradableBalanceKey(holder, batchDenom), amount)
	if err != nil {
		return err
	}

	// subtract tradable supply
	err = subAndSetDecimal(store, TradableSupplyKey(batchDenom), amount)
	if err != nil {
		return err
	}

	return nil
}

// gets the precision of the credit type associated with the batch
func (s serverImpl) getBatchPrecision(ctx types.Context, denom batchDenomT) (uint32, error) {
	var batchInfo ecocredit.BatchInfo
	err := s.batchInfoTable.GetOne(ctx, orm.RowID(denom), &batchInfo)
	if err != nil {
		return 0, err
	}

	classInfo, err := s.getClassInfo(ctx, batchInfo.ClassId)
	if err != nil {
		return 0, err
	}

	return classInfo.CreditType.Precision, nil
}

// Checks if the given address is in the allowlist of credit class designers
func (s serverImpl) isCreatorAllowListed(allowlist []string, designer sdk.Address) bool {
	for _, addr := range allowlist {
		allowListedAddr, _ := sdk.AccAddressFromBech32(addr)
		if designer.Equals(allowListedAddr) {
			return true
		}
	}
	return false
}
