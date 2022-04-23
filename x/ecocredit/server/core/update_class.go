package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// UpdateClassAdmin updates the admin address for a class.
// WARNING: this method will forfeit control of the entire class to the provided address.
// double check your inputs to ensure you do not lose control of the class.
func (k Keeper) UpdateClassAdmin(ctx context.Context, req *core.MsgUpdateClassAdmin) (*core.MsgUpdateClassAdminResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}
	newAdmin, err := sdk.AccAddressFromBech32(req.NewAdmin)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("class %s not found", req.ClassId)
	}

	classAdmin := sdk.AccAddress(classInfo.Admin)
	if !classAdmin.Equals(reqAddr) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}
	classInfo.Admin = newAdmin
	if err = k.stateStore.ClassTable().Update(ctx, classInfo); err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&api.EventClassAdminUpdated{
		ClassId:  req.ClassId,
		OldAdmin: reqAddr.String(),
		NewAdmin: newAdmin.String(),
	}); err != nil {
		return nil, err
	}

	return &core.MsgUpdateClassAdminResponse{}, err
}

// UpdateClassIssuers updates a class's issuers by either adding more issuers, or removing issuers from the class issuer store.
func (k Keeper) UpdateClassIssuers(ctx context.Context, req *core.MsgUpdateClassIssuers) (*core.MsgUpdateClassIssuersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	class, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("class %s not found", req.ClassId)
	}

	admin := sdk.AccAddress(class.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", class.Admin, req.Admin)
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

	if err = sdkCtx.EventManager().EmitTypedEvent(&api.EventClassIssuersUpdated{
		ClassId:        req.ClassId,
		AddedIssuers:   req.AddIssuers,
		RemovedIssuers: req.RemoveIssuers,
	}); err != nil {
		return nil, err
	}

	return &core.MsgUpdateClassIssuersResponse{}, nil
}

// UpdateClassMetadata updates the metadata for the class.
func (k Keeper) UpdateClassMetadata(ctx context.Context, req *core.MsgUpdateClassMetadata) (*core.MsgUpdateClassMetadataResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("class %s not found", req.ClassId)
	}

	admin := sdk.AccAddress(classInfo.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}

	oldMetadata := classInfo.Metadata
	classInfo.Metadata = req.Metadata
	if err = k.stateStore.ClassTable().Update(ctx, classInfo); err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&api.EventClassMetadataUpdated{
		ClassId:     req.ClassId,
		OldMetadata: oldMetadata,
		NewMetadata: req.Metadata,
	}); err != nil {
		return nil, err
	}

	return &core.MsgUpdateClassMetadataResponse{}, err
}
