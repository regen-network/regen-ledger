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

// MintBatchCredits issues additional credits from an open batch.
func (k Keeper) MintBatchCredits(ctx context.Context, req *core.MsgMintBatchCredits) (*core.MsgMintBatchCreditsResponse, error) {
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err
	}

	batch, err := k.stateStore.BatchTable().GetByDenom(ctx, req.BatchDenom)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get batch with denom %s: %s", req.BatchDenom, err)
	}

	if err = k.assertCanMintBatch(issuer, batch); err != nil {
		return nil, err
	}

	project, err := k.stateStore.ProjectTable().Get(ctx, batch.ProjectKey)
	if err != nil {
		return nil, err
	}

	if err = k.stateStore.OriginTxIndexTable().Insert(ctx, &api.OriginTxIndex{
		ClassKey: project.ClassKey,
		Id:       req.OriginTx.Id,
		Source:   req.OriginTx.Source,
	}); err != nil {
		if ormerrors.AlreadyExists.Is(err) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("credits already issued with tx id: %s", req.OriginTx.Id)
		}
		return nil, err
	}

	ct, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, batch.Denom)
	if err != nil {
		return nil, err
	}
	precision := ct.Precision
	moduleAddrString := k.moduleAddress.String()
	for _, iss := range req.Issuance {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		recipient, err := sdk.AccAddressFromBech32(iss.Recipient)
		if err != nil {
			return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid recipient address %s: %s", iss.Recipient, err.Error())
		}
		decs, err := utils.GetNonNegativeFixedDecs(precision, iss.TradableAmount, iss.RetiredAmount)
		if err != nil {
			return nil, err
		}
		tradable, retired := decs[0], decs[1]

		balance, err := utils.GetBalance(ctx, k.stateStore.BatchBalanceTable(), recipient, batch.Key)
		if err != nil {
			return nil, err
		}
		supply, err := k.stateStore.BatchSupplyTable().Get(ctx, batch.Key)
		if err != nil {
			return nil, err
		}
		decs, err = utils.GetNonNegativeFixedDecs(precision, balance.TradableAmount, balance.RetiredAmount, supply.TradableAmount, supply.RetiredAmount)
		if err != nil {
			return nil, err
		}

		balanceTradable, balanceRetired := decs[0], decs[1]
		supplyTradable, supplyRetired := decs[2], decs[3]

		if !retired.IsZero() {
			balanceRetired, err = balanceRetired.Add(retired)
			if err != nil {
				return nil, err
			}
			supplyRetired, err = supplyRetired.Add(retired)
			if err != nil {
				return nil, err
			}
			if err := sdkCtx.EventManager().EmitTypedEvent(&core.EventRetire{
				Owner:        iss.Recipient,
				BatchDenom:   req.BatchDenom,
				Amount:       iss.RetiredAmount,
				Jurisdiction: iss.RetirementJurisdiction,
			}); err != nil {
				return nil, err
			}
			balance.RetiredAmount = balanceRetired.String()
			supply.RetiredAmount = supplyRetired.String()
		}
		if !tradable.IsZero() {
			balanceTradable, err = balanceTradable.Add(tradable)
			if err != nil {
				return nil, err
			}
			supplyTradable, err = supplyTradable.Add(tradable)
			if err != nil {
				return nil, err
			}
			if err := sdkCtx.EventManager().EmitTypedEvent(&core.EventTransfer{
				Sender:         moduleAddrString, // ecocredit module
				Recipient:      iss.Recipient,
				BatchDenom:     req.BatchDenom,
				TradableAmount: iss.TradableAmount,
				RetiredAmount:  iss.RetiredAmount,
			}); err != nil {
				return nil, err
			}
			balance.TradableAmount = balanceTradable.String()
			supply.TradableAmount = supplyTradable.String()
		}
		if err := k.stateStore.BatchBalanceTable().Save(ctx, balance); err != nil {
			return nil, err
		}
		if err := k.stateStore.BatchSupplyTable().Update(ctx, supply); err != nil {
			return nil, err
		}

		if err := sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&core.EventMint{
			BatchDenom:     batch.Denom,
			TradableAmount: tradable.String(),
			RetiredAmount:  retired.String(),
		}); err != nil {
			return nil, err
		}
	}

	if err := sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&core.EventMintBatchCredits{
		BatchDenom: batch.Denom,
		OriginTx:   req.OriginTx,
	}); err != nil {
		return nil, err
	}

	return &core.MsgMintBatchCreditsResponse{}, nil
}

// asserts that the batch is open for minting and that the requester address matches the batch issuer address.
func (k Keeper) assertCanMintBatch(issuer sdk.AccAddress, batch *api.Batch) error {
	if !batch.Open {
		return sdkerrors.ErrInvalidRequest.Wrap("credits cannot be minted in a closed batch")
	}
	if !sdk.AccAddress(batch.Issuer).Equals(issuer) {
		return sdkerrors.ErrUnauthorized.Wrapf("only the account that issued the batch can mint additional credits")
	}
	return nil
}
