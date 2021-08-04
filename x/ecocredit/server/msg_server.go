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
	"github.com/regen-network/regen-ledger/x/ecocredit/util"
)

// CreateClass creates a new class of ecocredit
//
// The designer is charged a fee for creating the class. This is controlled by
// the global parameter CreditClassFee, which can be updated through the
// governance process.
func (s serverImpl) CreateClass(goCtx context.Context, req *ecocredit.MsgCreateClass) (*ecocredit.MsgCreateClassResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	classID := s.idSeq.NextVal(ctx)
	classIDStr := util.Uint64ToBase58Check(classID)

	// Charge the designer a fee to create the credit class
	designerAddress, err := sdk.AccAddressFromBech32(req.Designer)
	if err != nil {
		return nil, err
	}

	err = s.chargeCreditClassFee(ctx.Context, designerAddress)
	if err != nil {
		return nil, err
	}

	err = s.classInfoTable.Create(ctx, &ecocredit.ClassInfo{
		ClassId:  classIDStr,
		Designer: req.Designer,
		Issuers:  req.Issuers,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateClass{
		ClassId:  classIDStr,
		Designer: req.Designer,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateClassResponse{ClassId: classIDStr}, nil
}

func (s serverImpl) CreateBatch(goCtx context.Context, req *ecocredit.MsgCreateBatch) (*ecocredit.MsgCreateBatchResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	classID := req.ClassId
	if err := s.assertClassIssuer(ctx, classID, req.Issuer); err != nil {
		return nil, err
	}

	batchID := s.idSeq.NextVal(ctx)
	batchDenom := batchDenomT(fmt.Sprintf("%s/%s", classID, util.Uint64ToBase58Check(batchID)))
	tradableSupply := math.NewDecFromInt64(0)
	retiredSupply := math.NewDecFromInt64(0)
	var maxDecimalPlaces uint32 = 0

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
				maxDecimalPlaces = decPlaces
			}
		}

		if issuance.RetiredAmount != "" {
			retired, err = math.NewNonNegativeDecFromString(issuance.RetiredAmount)
			if err != nil {
				return nil, err
			}

			decPlaces := retired.NumDecimalPlaces()
			if decPlaces > maxDecimalPlaces {
				maxDecimalPlaces = decPlaces
			}
		}

		recipient := issuance.Recipient

		if !tradable.IsZero() {
			tradableSupply, err = tradableSupply.Add(tradable)
			if err != nil {
				return nil, err
			}

			err := getAddAndSetDecimal(store, TradableBalanceKey(recipient, batchDenom), tradable)
			if err != nil {
				return nil, err
			}
		}

		if !retired.IsZero() {
			retiredSupply, err = retiredSupply.Add(retired)
			if err != nil {
				return nil, err
			}

			err = retire(ctx, store, recipient, batchDenom, retired, issuance.RetirementLocation)
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

	err = setUInt32(store, MaxDecimalPlacesKey(batchDenom), maxDecimalPlaces)
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

func (s serverImpl) Send(goCtx context.Context, req *ecocredit.MsgSend) (*ecocredit.MsgSendResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	sender := req.Sender
	recipient := req.Recipient

	for _, credit := range req.Credits {
		denom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid credit batch denom", denom)
		}

		maxDecimalPlaces, err := getUint32(store, MaxDecimalPlacesKey(denom))
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
		err = getSubAndSetDecimal(store, TradableBalanceKey(sender, denom), sum)
		if err != nil {
			return nil, err
		}

		// Add tradable balance
		err = getAddAndSetDecimal(store, TradableBalanceKey(recipient, denom), tradable)
		if err != nil {
			return nil, err
		}

		if !retired.IsZero() {
			// subtract retired from tradable supply
			err = getSubAndSetDecimal(store, TradableSupplyKey(denom), retired)
			if err != nil {
				return nil, err
			}

			// Add retired balance
			err = retire(ctx, store, recipient, denom, retired, credit.RetirementLocation)
			if err != nil {
				return nil, err
			}

			// Add retired supply
			err = getAddAndSetDecimal(store, RetiredSupplyKey(denom), retired)
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

func (s serverImpl) Retire(goCtx context.Context, req *ecocredit.MsgRetire) (*ecocredit.MsgRetireResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	holder := req.Holder

	for _, credit := range req.Credits {
		denom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid credit batch denom", denom)
		}

		maxDecimalPlaces, err := getUint32(store, MaxDecimalPlacesKey(denom))
		if err != nil {
			return nil, err
		}

		toRetire, err := math.NewPositiveFixedDecFromString(credit.Amount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		err = subtractTradableBalanceAndSupply(store, holder, denom, toRetire)
		if err != nil {
			return nil, err
		}

		//  Add retired balance
		err = retire(ctx, store, holder, denom, toRetire, req.Location)
		if err != nil {
			return nil, err
		}

		//  Add retired supply
		err = getAddAndSetDecimal(store, RetiredSupplyKey(denom), toRetire)
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgRetireResponse{}, nil
}

func (s serverImpl) Cancel(goCtx context.Context, req *ecocredit.MsgCancel) (*ecocredit.MsgCancelResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	holder := req.Holder
	for _, credit := range req.Credits {

		// Check that the batch that were trying to cancel credits from
		// exists
		denom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid credit batch denom", denom)
		}

		// Fetch the max precision of decimal values in this batch
		maxDecimalPlaces, err := getUint32(store, MaxDecimalPlacesKey(denom))
		if err != nil {
			return nil, err
		}

		// Parse the amount of credits to cancel, checking it conforms
		// to the precision
		toCancel, err := math.NewPositiveFixedDecFromString(credit.Amount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		// Remove the credits from the balance of the holder and the
		// overall supply
		err = subtractTradableBalanceAndSupply(store, holder, denom, toCancel)
		if err != nil {
			return nil, err
		}

		// Remove the credits from the total_amount in the batch and add
		// them to amount_cancelled
		var batchInfo ecocredit.BatchInfo
		err = s.batchInfoTable.GetOne(ctx, orm.RowID(denom), &batchInfo)
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

		s.batchInfoTable.Save(ctx, &batchInfo)

		// Emit the cancellation event
		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCancel{
			Canceller:  holder,
			BatchDenom: string(denom),
			Amount:     toCancel.String(),
		})
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgCancelResponse{}, nil
}

func (s serverImpl) SetPrecision(goCtx context.Context, req *ecocredit.MsgSetPrecision) (*ecocredit.MsgSetPrecisionResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	var batchInfo ecocredit.BatchInfo

	err := s.batchInfoTable.GetOne(ctx, orm.RowID(req.BatchDenom), &batchInfo)
	if err != nil {
		return nil, err
	}

	if req.Issuer != batchInfo.Issuer {
		return nil, sdkerrors.ErrUnauthorized
	}

	store := ctx.KVStore(s.storeKey)
	key := MaxDecimalPlacesKey(batchDenomT(req.BatchDenom))
	x, err := getUint32(store, key)
	if err != nil {
		return nil, err
	}

	if req.MaxDecimalPlaces <= x {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Maximum decimal can only be increased, it is currently %d, and %d was requested", x, req.MaxDecimalPlaces))
	}

	err = setUInt32(store, key, req.MaxDecimalPlaces)
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgSetPrecisionResponse{}, nil
}

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (s serverImpl) assertClassIssuer(goCtx context.Context, classID, issuer string) error {
	ctx := types.UnwrapSDKContext(goCtx)
	classInfo, err := s.getClassInfo(ctx, classID)
	if err != nil {
		return err
	}
	for _, i := range classInfo.Issuers {
		if issuer == i {
			return nil
		}
	}
	return sdkerrors.ErrUnauthorized
}

func retire(ctx types.Context, store sdk.KVStore, recipient string, batchDenom batchDenomT, retired math.Dec, location string) error {
	err := getAddAndSetDecimal(store, RetiredBalanceKey(recipient, batchDenom), retired)
	if err != nil {
		return err
	}

	return ctx.EventManager().EmitTypedEvent(&ecocredit.EventRetire{
		Retirer:    recipient,
		BatchDenom: string(batchDenom),
		Amount:     retired.String(),
		Location:   location,
	})
}

func subtractTradableBalanceAndSupply(store sdk.KVStore, holder string, batchDenom batchDenomT, amount math.Dec) error {
	// subtract tradable balance
	err := getSubAndSetDecimal(store, TradableBalanceKey(holder, batchDenom), amount)
	if err != nil {
		return err
	}

	// subtract tradable supply
	err = getSubAndSetDecimal(store, TradableSupplyKey(batchDenom), amount)
	if err != nil {
		return err
	}

	return nil
}
