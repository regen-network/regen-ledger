package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

func (k Keeper) MintBatchCredits(ctx context.Context, req *core.MsgMintBatchCredits) (*core.MsgMintBatchCreditsResponse, error) {
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err.Error())
	}

	if !batch.Open {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("batch credits cannot be minted in a closed batch")
	}

	if !sdk.AccAddress(batch.Issuer).Equals(issuer) {
		return nil, sdkerrors.ErrUnauthorized.Wrap("only the batch issuer can mint more credits")
	}

	if err = k.stateStore.BatchOrigTxTable().Insert(ctx, &api.BatchOrigTx{
		TxId:       req.OriginTx.Id,
		Typ:        req.OriginTx.Typ,
		Note:       req.Note,
		BatchDenom: req.BatchDenom,
	}); err != nil {
		return nil, err
	}

	ct, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, k.paramsKeeper, batch.Denom)
	if err != nil {
		return nil, err
	}
	for _, iss := range req.Issuance {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		recipient, err := sdk.AccAddressFromBech32(iss.Recipient)
		if err != nil {
			return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid recipient address %s: %s", iss.Recipient, err.Error())
		}
		decs, err := utils.GetNonNegativeFixedDecs(ct.Precision, iss.TradableAmount, iss.RetiredAmount)
		if err != nil {
			return nil, err
		}
		tradable, retired := decs[0], decs[1]
		balance, err := k.stateStore.BatchBalanceTable().Get(ctx, recipient, batch.Key)
		if err != nil {
			if ormerrors.IsNotFound(err) {
				balance = &api.BatchBalance{
					BatchKey: batch.Key,
					Address:  recipient,
				}
			} else {
				return nil, err
			}
		}
		supply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
		if err != nil {
			return nil, err
		}
		sDecs, err := utils.GetNonNegativeFixedDecs(ct.Precision, balance.Tradable, balance.Retired, supply.TradableAmount, supply.RetiredAmount)
		if err != nil {
			return nil, err
		}

		balT, balR := sDecs[0], sDecs[1]
		supT, supR := sDecs[2], sDecs[3]

		if !retired.IsZero() {
			balR, err = balR.Add(retired)
			if err != nil {
				return nil, err
			}
			supR, err = supR.Add(retired)
			if err != nil {
				return nil, err
			}
			if err := sdkCtx.EventManager().EmitTypedEvent(&core.EventRetire{
				Retirer:      iss.Recipient,
				BatchDenom:   req.BatchDenom,
				Amount:       iss.RetiredAmount,
				Jurisdiction: iss.RetirementJurisdiction,
			}); err != nil {
				return nil, err
			}

		}
		if !tradable.IsZero() {
			balT, err = balT.Add(tradable)
			if err != nil {
				return nil, err
			}
			supT, err = supT.Add(tradable)
			if err != nil {
				return nil, err
			}
			if err := sdkCtx.EventManager().EmitTypedEvent(&core.EventReceive{
				Sender:         req.Issuer,
				Recipient:      iss.Recipient,
				BatchDenom:     req.BatchDenom,
				TradableAmount: iss.TradableAmount,
				RetiredAmount:  iss.RetiredAmount,
				BasketDenom:    "",
			}); err != nil {
				return nil, err
			}
		}
		balance.Tradable, balance.Retired = balT.String(), balR.String()
		if err := k.stateStore.BatchBalanceTable().Save(ctx, balance); err != nil {
			return nil, err
		}

		supply.TradableAmount, supply.RetiredAmount = supT.String(), supR.String()
		if err := k.stateStore.BatchSupplyTable().Update(ctx, supply); err != nil {
			return nil, err
		}
	}
	return &core.MsgMintBatchCreditsResponse{}, nil
}
