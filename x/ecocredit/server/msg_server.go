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

	store := ctx.KVStore(s.storeKey)

	for _, issuance := range req.Issuance {
		tradeable, err := math.MustParseNonNegativeDecimal(issuance.TradeableUnits)
		if err != nil {
			return nil, err
		}

		retired, err := math.MustParseNonNegativeDecimal(issuance.RetiredUnits)
		if err != nil {
			return nil, err
		}

		recipient := issuance.Recipient

		if !tradeable.IsZero() {
			err = add(tradeableSupply, tradeableSupply, tradeable)
			if err != nil {
				return nil, err
			}

			err = s.receiveTradeable(ctx, store, recipient, batchDenom, tradeable)
			if err != nil {
				return nil, err
			}
		}

		if !retired.IsZero() {
			err = add(retiredSupply, retiredSupply, retired)
			if err != nil {
				return nil, err
			}

			err = s.retire(ctx, store, recipient, batchDenom, retired)
			if err != nil {
				return nil, err
			}
		}
	}

	s.setDec(store, TradeableSupplyKey(batchDenom), tradeableSupply)
	s.setDec(store, RetiredSupplyKey(batchDenom), retiredSupply)

	var totalSupply apd.Decimal
	err = add(&totalSupply, tradeableSupply, retiredSupply)
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
		tradeable, err := math.MustParseNonNegativeDecimal(credit.TradeableUnits)
		if err != nil {
			return nil, err
		}

		retired, err := math.MustParseNonNegativeDecimal(credit.RetiredUnits)
		if err != nil {
			return nil, err
		}

		var sum apd.Decimal
		err = add(&sum, tradeable, retired)
		if err != nil {
			return nil, err
		}

		denom := batchDenomT(credit.BatchDenom)

		// subtract balance
		err = s.safeSubDec(store, TradeableBalanceKey(sender, denom), &sum)
		if err != nil {
			return nil, err
		}

		// subtract retired from tradeable supply
		err = s.safeSubDec(store, TradeableSupplyKey(denom), retired)
		if err != nil {
			return nil, err
		}

		// add tradeable balance
		err = s.receiveTradeable(ctx, store, recipient, denom, tradeable)
		if err != nil {
			return nil, err
		}

		// add retired balance
		err = s.retire(ctx, store, recipient, denom, retired)
		if err != nil {
			return nil, err
		}

		// add retired supply
		err = s.addDec(store, RetiredSupplyKey(denom), retired)
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

		toRetire, err := math.MustParsePositiveDecimal(credit.Units)
		if err != nil {
			return nil, err
		}

		// subtract tradeable balance
		err = s.safeSubDec(store, TradeableBalanceKey(holder, denom), toRetire)
		if err != nil {
			return nil, err
		}

		// subtract tradeable supply
		err = s.safeSubDec(store, TradeableSupplyKey(denom), toRetire)
		if err != nil {
			return nil, err
		}

		//  add retired balance
		err = s.retire(ctx, store, holder, denom, toRetire)
		if err != nil {
			return nil, err
		}

		//  add retired supply
		err = s.addDec(store, RetiredSupplyKey(denom), toRetire)
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgRetireResponse{}, nil
}

func (s serverImpl) receiveTradeable(ctx sdk.Context, store sdk.KVStore, recipient string, batchDenom batchDenomT, tradeable *apd.Decimal) error {
	err := s.addDec(store, TradeableBalanceKey(recipient, batchDenom), tradeable)
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
	err := s.addDec(store, RetiredBalanceKey(recipient, batchDenom), retired)
	if err != nil {
		return err
	}

	return ctx.EventManager().EmitTypedEvent(&ecocredit.EventRetire{
		Retirer:    recipient,
		BatchDenom: string(batchDenom),
		Units:      math.DecString(retired),
	})
}

func (s serverImpl) setDec(store sdk.KVStore, key []byte, value *apd.Decimal) {
	// always remove all trailing zeros for canonical representation
	value, _ = value.Reduce(value)
	// use scientific notation here always for canonical representation
	str := value.Text('e')
	store.Set(key, []byte(str))
}

func (s serverImpl) addDec(store sdk.KVStore, key []byte, x *apd.Decimal) error {
	value, err := s.getDec(store, key)
	if err != nil {
		return err
	}

	err = add(value, value, x)
	if err != nil {
		return err
	}

	s.setDec(store, key, value)
	return nil
}

func (s serverImpl) safeSubDec(store sdk.KVStore, key []byte, x *apd.Decimal) error {
	value, err := s.getDec(store, key)
	if err != nil {
		return err
	}

	_, err = math.StrictDecimal128Context.Sub(value, value, x)
	if err != nil {
		return sdkerrors.Wrap(err, "decimal subtraction error")
	}

	if math.IsNegative(value) {
		return sdkerrors.ErrInsufficientFunds
	}

	s.setDec(store, key, value)
	return nil
}

func add(res, x, y *apd.Decimal) error {
	_, err := math.StrictDecimal128Context.Add(res, x, y)
	if err != nil {
		return sdkerrors.Wrap(err, "decimal addition error")
	}
	return nil
}
