package server

import (
	"context"
	"fmt"
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

	retiredSupply := sdk.ZeroInt()

	for _, issuance := range req.Issuance {
		tradeable, ok := sdk.NewIntFromString(issuance.TradeableUnits)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid integer", issuance.TradeableUnits))
		}

		retired, ok := sdk.NewIntFromString(issuance.RetiredUnits)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid integer", issuance.RetiredUnits))
		}

		recipient := issuance.Recipient

		if !tradeable.IsZero() {
			coins := []sdk.Coin{
				sdk.NewCoin(batchDenom, tradeable),
			}
			err = s.bankKeeper.MintCoins(ctx, ecocredit.ModuleName, coins)
			if err != nil {
				return nil, err
			}

			addr, err := sdk.AccAddressFromBech32(recipient)
			if err != nil {
				return nil, err
			}

			err = s.bankKeeper.SendCoinsFromModuleToAccount(ctx, ecocredit.ModuleName, addr, coins)
			if err != nil {
				return nil, err
			}
		}

		if !retired.IsZero() {
			retiredSupply = retiredSupply.Add(retired)

			err = s.setRetiredBalance(ctx, recipient, batchDenom, retired)
			if err != nil {
				return nil, err
			}
		}
	}

	if !retiredSupply.IsZero() {
		err = s.setRetiredSupply(ctx, batchDenom, retiredSupply)
		if err != nil {
			return nil, err
		}
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

func (s serverImpl) Retire(goCtx context.Context, req *ecocredit.MsgRetireRequest) (*ecocredit.MsgRetireResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	holder := req.Holder
	addr, err := sdk.AccAddressFromBech32(holder)
	if err != nil {
		return nil, err
	}

	var coins sdk.Coins

	for _, credit := range req.Credits {
		denom := credit.Denom

		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid credit denom", denom))
		}

		coins = append(coins, *credit)

		// Handle Balance
		retiredBalance, err := s.getRetiredBalance(ctx, holder, denom)
		if err != nil {
			return nil, err
		}

		amount := credit.Amount
		retiredBalance = retiredBalance.Add(amount)

		err = s.setRetiredBalance(ctx, holder, denom, retiredBalance)
		if err != nil {
			return nil, err
		}

		// Handle Supply
		retiredSupply, err := s.getRetiredSupply(ctx, denom)
		if err != nil {
			return nil, err
		}

		retiredSupply = retiredSupply.Add(amount)

		err = s.setRetiredSupply(ctx, denom, retiredSupply)
	}

	err = s.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, ecocredit.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = s.bankKeeper.BurnCoins(ctx, ecocredit.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgRetireResponse{}, nil
}

func (s serverImpl) setRetiredSupply(ctx sdk.Context, batchDenom string, supply sdk.Int) error {
	intProto := sdk.IntProto{Int: supply}

	bz, err := intProto.Marshal()
	if err != nil {
		return err
	}

	store := ctx.KVStore(s.storeKey)
	store.Set(RetiredSupplyKey(batchDenom), bz)

	return nil
}
