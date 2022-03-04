package core

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

// UpdateClassAdmin updates the admin address for a class.
// WARNING: this method will forfeit control of the entire class to the provided address.
// double check your inputs to ensure you do not lose control of the class.
func (k Keeper) UpdateClassAdmin(ctx context.Context, req *v1beta1.MsgUpdateClassAdmin) (*v1beta1.MsgUpdateClassAdminResponse, error) {
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}
	newAdmin, err := sdk.AccAddressFromBech32(req.NewAdmin)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("class %s not found", req.ClassId)
	}

	classAdmin := sdk.AccAddress(classInfo.Admin)
	if !classAdmin.Equals(reqAddr) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}
	classInfo.Admin = newAdmin
	if err = k.stateStore.ClassInfoStore().Update(ctx, classInfo); err != nil {
		return nil, err
	}
	return &v1beta1.MsgUpdateClassAdminResponse{}, err
}

// UpdateClassIssuers updates a class's issuers by either adding more issuers, or removing issuers from the class issuer store.
func (k Keeper) UpdateClassIssuers(ctx context.Context, req *v1beta1.MsgUpdateClassIssuers) (*v1beta1.MsgUpdateClassIssuersResponse, error) {
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	class, err := k.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
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
		if err = k.stateStore.ClassIssuerStore().Delete(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: class.Id,
			Issuer:  issuerAcc,
		}); err != nil {
			return nil, err
		}
	}

	// add the new issuers
	for _, issuer := range req.AddIssuers {
		issuerAcc, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return nil, err
		}
		if err = k.stateStore.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: class.Id,
			Issuer:  issuerAcc,
		}); err != nil {
			return nil, err
		}
	}
	return &v1beta1.MsgUpdateClassIssuersResponse{}, nil
}

// UpdateClassMetadata updates the metadata for the class.
func (k Keeper) UpdateClassMetadata(ctx context.Context, req *v1beta1.MsgUpdateClassMetadata) (*v1beta1.MsgUpdateClassMetadataResponse, error) {
	reqAddr, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrapf("class %s not found", req.ClassId)
	}

	admin := sdk.AccAddress(classInfo.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}

	classInfo.Metadata = req.Metadata
	if err = k.stateStore.ClassInfoStore().Update(ctx, classInfo); err != nil {
		return nil, err
	}

	return &v1beta1.MsgUpdateClassMetadataResponse{}, err
}
