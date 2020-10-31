package server

import (
	"context"
	"fmt"
	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/modules/incubator/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
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

	batchDenom := fmt.Sprintf("%s/%s/%x", s.denomPrefix, classId, batchId)

	err = s.batchInfoTable.Create(ctx, &ecocredit.BatchInfo{
		ClassId:    classId,
		BatchDenom: batchDenom,
		Issuer:     issuer,
	})
	if err != nil {
		return nil, err
	}

	tradeableSupply := apd.New(0, 0)
	retiredSupply := apd.New(0, 0)

	store := ctx.KVStore(s.storeKey)

	for _, issuance := range req.Issuance {
		tradeable, _, err := apd.NewFromString(issuance.TradeableUnits)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid integer", issuance.TradeableUnits))
		}

		retired, _, err := apd.NewFromString(issuance.RetiredUnits)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid integer", issuance.RetiredUnits))
		}

		recipient := issuance.Recipient

		if !tradeable.IsZero() {
			s.setDec(store, TradeableBalanceKey(recipient, batchDenom), tradeable)
			err = add(tradeableSupply, tradeableSupply, tradeable)
			if err != nil {
				return nil, err
			}
		}

		if !retired.IsZero() {
			s.setDec(store, RetiredBalanceKey(recipient, batchDenom), retired)
			err = add(retiredSupply, retiredSupply, retired)
			if err != nil {
				return nil, err
			}
		}
	}

	if !retiredSupply.IsZero() {
		s.setDec(store, RetiredSupplyKey(batchDenom), retiredSupply)
	}

	return &ecocredit.MsgCreateBatchResponse{BatchDenom: batchDenom}, nil
}

func (s serverImpl) setRetiredBalance(ctx sdk.Context, holder string, batchDenom string, retiredBalance sdk.Int) error {
	intProto := sdk.IntProto{Int: retiredBalance}
	bz, err := intProto.Marshal()
	if err != nil {
		return err
	}

	store := ctx.KVStore(s.storeKey)
	store.Set(RetiredBalanceKey(holder, batchDenom), bz)
	return nil
}

func (s serverImpl) Send(goCtx context.Context, req *ecocredit.MsgSendRequest) (*ecocredit.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender := req.Sender
	recipient := req.Recipient

	store := ctx.KVStore(s.storeKey)

	for _, credit := range req.Credits {
		tradeable, _, err := apd.NewFromString(credit.TradeableUnits)
		if err != nil {
			return nil, err
		}

		err = requirePositive(tradeable)
		if err != nil {
			return nil, err
		}

		retired, _, err := apd.NewFromString(credit.RetiredUnits)
		if err != nil {
			return nil, err
		}

		err = requirePositive(retired)
		if err != nil {
			return nil, err
		}

		var sum apd.Decimal
		err = add(&sum, tradeable, retired)
		if err != nil {
			return nil, err
		}

		denom := credit.BatchDenom

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
		err = s.addDec(store, TradeableBalanceKey(recipient, denom), tradeable)
		if err != nil {
			return nil, err
		}

		// add retired balance
		err = s.addDec(store, RetiredBalanceKey(recipient, denom), retired)
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
		denom := credit.BatchDenom

		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid credit denom", denom))
		}

		toRetire, _, err := apd.NewFromString(credit.Units)
		if err != nil {
			return nil, err
		}

		err = requirePositive(toRetire)
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
		err = s.addDec(store, RetiredBalanceKey(holder, denom), toRetire)
		if err != nil {
			return nil, err
		}

		//  add retired supply
		err = s.addDec(store, RetiredSupplyKey(holder), toRetire)
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgRetireResponse{}, nil
}
