package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// UpdateClassIssuers updates a class's issuers by either adding more issuers, or removing issuers from the class issuer store.
func (k Keeper) UpdateClassIssuers(ctx context.Context, req *core.MsgUpdateClassIssuers) (*core.MsgUpdateClassIssuersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	class, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"could not get credit class with id %s: %s", req.ClassId, err,
		)
	}

	admin := sdk.AccAddress(class.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"%s is not the admin of credit class %s", req.Admin, req.ClassId,
		)
	}

	// remove issuers
	for _, issuer := range req.RemoveIssuers {
		issuerAcc, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return nil, err
		}
		if err = k.stateStore.ClassIssuerTable().Delete(ctx, &api.ClassIssuer{
			ClassKey: class.Key,
			Issuer:   issuerAcc,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgUpdateClassIssuers issuer iteration")
	}

	// add the new issuers
	for _, issuer := range req.AddIssuers {
		issuerAcc, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return nil, err
		}
		if err = k.stateStore.ClassIssuerTable().Insert(ctx, &api.ClassIssuer{
			ClassKey: class.Key,
			Issuer:   issuerAcc,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgUpdateClassIssuers issuer iteration")
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventUpdateClassIssuers{
		ClassId: req.ClassId,
	}); err != nil {
		return nil, err
	}

	return &core.MsgUpdateClassIssuersResponse{}, nil
}
