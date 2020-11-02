package server

import (
	"context"
	"fmt"
	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/modules/incubator/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/math"
)

func (s serverImpl) CreateClass(goCtx context.Context, req *ecocredit.MsgCreateClassRequest) (*ecocredit.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	classId := s.classInfoSeq.NextVal(ctx)

	classIdStr := fmt.Sprintf("%x", classId)

	err := s.classInfoTable.Create(ctx, &ecocredit.ClassInfo{
		ClassId:  classIdStr,
		Designer: req.Designer,
		Issuers:  req.Issuers,
		Metadata: req.Metadata,
	})

	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateClass{
		ClassId:  classIdStr,
		Designer: req.Designer,
	})

	return &ecocredit.MsgCreateClassResponse{ClassId: classIdStr}, nil
}

func (s serverImpl) CreateBatch(goCtx context.Context, req *ecocredit.MsgCreateBatchRequest) (*ecocredit.MsgCreateBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	classId := req.ClassId
	classInfo, err := s.getClassInfo(ctx, classId)
	if err != nil {
		return nil, err
	}

	var found bool
	issuer := req.Issuer
	for _, issuer := range classInfo.Issuers {
		if issuer == issuer {
			found = true
			break
		}
	}

	if !found {
		return nil, sdkerrors.ErrUnauthorized
	}

	batchId := s.batchInfoSeq.NextVal(ctx)

	batchDenom := batchDenomT(fmt.Sprintf("%s/%x", classId, batchId))

	tradeableSupply := apd.New(0, 0)
	retiredSupply := apd.New(0, 0)
	var maxDecimalPlaces uint32 = 0

	store := ctx.KVStore(s.storeKey)

	for _, issuance := range req.Issuance {
		tradeable, err := math.ParseNonNegativeDecimal(issuance.TradeableUnits)
		if err != nil {
			return nil, err
		}

		decPlaces := math.NumDecimalPlaces(tradeable)
		if decPlaces > maxDecimalPlaces {
			maxDecimalPlaces = decPlaces
		}

		retired, err := math.ParseNonNegativeDecimal(issuance.RetiredUnits)
		if err != nil {
			return nil, err
		}

		decPlaces = math.NumDecimalPlaces(retired)
		if decPlaces > maxDecimalPlaces {
			maxDecimalPlaces = decPlaces
		}

		recipient := issuance.Recipient

		if !tradeable.IsZero() {
			err = math.Add(tradeableSupply, tradeableSupply, tradeable)
			if err != nil {
				return nil, err
			}

			err = s.receiveTradeable(ctx, store, recipient, batchDenom, tradeable)
			if err != nil {
				return nil, err
			}
		}

		if !retired.IsZero() {
			err = math.Add(retiredSupply, retiredSupply, retired)
			if err != nil {
				return nil, err
			}

			err = s.retire(ctx, store, recipient, batchDenom, retired)
			if err != nil {
				return nil, err
			}
		}
	}

	storeSetDec(store, TradeableSupplyKey(batchDenom), tradeableSupply)
	storeSetDec(store, RetiredSupplyKey(batchDenom), retiredSupply)

	var totalSupply apd.Decimal
	err = math.Add(&totalSupply, tradeableSupply, retiredSupply)
	if err != nil {
		return nil, err
	}

	totalSupplyStr := math.DecString(&totalSupply)

	err = s.batchInfoTable.Create(ctx, &ecocredit.BatchInfo{
		ClassId:    classId,
		BatchDenom: string(batchDenom),
		Issuer:     issuer,
		TotalUnits: totalSupplyStr,
	})
	if err != nil {
		return nil, err
	}

	err = storeSetUInt32(store, MaxDecimalPlacesKey(batchDenom), maxDecimalPlaces)
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateBatch{
		ClassId:    classId,
		BatchDenom: string(batchDenom),
		Issuer:     issuer,
		TotalUnits: totalSupplyStr,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateBatchResponse{BatchDenom: string(batchDenom)}, nil
}

func (s serverImpl) Send(goCtx context.Context, req *ecocredit.MsgSendRequest) (*ecocredit.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender := req.Sender
	recipient := req.Recipient

	store := ctx.KVStore(s.storeKey)

	for _, credit := range req.Credits {
		denom := batchDenomT(credit.BatchDenom)

		maxDecimalPlaces, err := storeGetUInt32(store, MaxDecimalPlacesKey(denom))
		if err != nil {
			return nil, err
		}

		tradeable, err := math.ParseNonNegativeFixedDecimal(credit.TradeableUnits, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		retired, err := math.ParseNonNegativeFixedDecimal(credit.RetiredUnits, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		var sum apd.Decimal
		err = math.Add(&sum, tradeable, retired)
		if err != nil {
			return nil, err
		}

		// subtract balance
		err = storeSafeSubDec(store, TradeableBalanceKey(sender, denom), &sum)
		if err != nil {
			return nil, err
		}

		// subtract retired from tradeable supply
		err = storeSafeSubDec(store, TradeableSupplyKey(denom), retired)
		if err != nil {
			return nil, err
		}

		// Add tradeable balance
		err = s.receiveTradeable(ctx, store, recipient, denom, tradeable)
		if err != nil {
			return nil, err
		}

		// Add retired balance
		err = s.retire(ctx, store, recipient, denom, retired)
		if err != nil {
			return nil, err
		}

		// Add retired supply
		err = storeAddDec(store, RetiredSupplyKey(denom), retired)
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgSendResponse{}, nil
}

func (s serverImpl) Retire(goCtx context.Context, req *ecocredit.MsgRetireRequest) (*ecocredit.MsgRetireResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	holder := req.Holder

	store := ctx.KVStore(s.storeKey)

	for _, credit := range req.Credits {
		denom := batchDenomT(credit.BatchDenom)

		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid credit denom", denom))
		}

		maxDecimalPlaces, err := storeGetUInt32(store, MaxDecimalPlacesKey(denom))
		if err != nil {
			return nil, err
		}

		toRetire, err := math.ParsePositiveFixedDecimal(credit.Units, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		// subtract tradeable balance
		err = storeSafeSubDec(store, TradeableBalanceKey(holder, denom), toRetire)
		if err != nil {
			return nil, err
		}

		// subtract tradeable supply
		err = storeSafeSubDec(store, TradeableSupplyKey(denom), toRetire)
		if err != nil {
			return nil, err
		}

		//  Add retired balance
		err = s.retire(ctx, store, holder, denom, toRetire)
		if err != nil {
			return nil, err
		}

		//  Add retired supply
		err = storeAddDec(store, RetiredSupplyKey(denom), toRetire)
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgRetireResponse{}, nil
}

func (s serverImpl) receiveTradeable(ctx sdk.Context, store sdk.KVStore, recipient string, batchDenom batchDenomT, tradeable *apd.Decimal) error {
	err := storeAddDec(store, TradeableBalanceKey(recipient, batchDenom), tradeable)
	if err != nil {
		return err
	}

	return ctx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
		Recipient:  recipient,
		BatchDenom: string(batchDenom),
		Units:      math.DecString(tradeable),
	})
}

func (s serverImpl) retire(ctx sdk.Context, store sdk.KVStore, recipient string, batchDenom batchDenomT, retired *apd.Decimal) error {
	err := storeAddDec(store, RetiredBalanceKey(recipient, batchDenom), retired)
	if err != nil {
		return err
	}

	return ctx.EventManager().EmitTypedEvent(&ecocredit.EventRetire{
		Retirer:    recipient,
		BatchDenom: string(batchDenom),
		Units:      math.DecString(retired),
	})
}

func (s serverImpl) SetPrecision(goCtx context.Context, request *ecocredit.MsgSetPrecisionRequest) (*ecocredit.MsgSetPrecisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)
	key := MaxDecimalPlacesKey(batchDenomT(request.BatchDenom))
	x, err := storeGetUInt32(store, key)
	if err != nil {
		return nil, err
	}

	if request.MaxDecimalPlaces <= x {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Maximum decimal can only be increased, it is currently %d, and %d was requested", x, request.MaxDecimalPlaces))
	}

	err = storeSetUInt32(store, key, x)
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgSetPrecisionResponse{}, nil
}
