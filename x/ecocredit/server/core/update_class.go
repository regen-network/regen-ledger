package core

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

func (k Keeper) UpdateClassAdmin(ctx context.Context, req *v1beta1.MsgUpdateClassAdmin) (*v1beta1.MsgUpdateClassAdminResponse, error) {
	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}
	reqAddr, _ := sdk.AccAddressFromBech32(req.Admin)
	classAdmin := sdk.AccAddress(classInfo.Admin)
	if !classAdmin.Equals(reqAddr) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}
	classInfo.Admin = reqAddr
	if err = k.stateStore.ClassInfoStore().Update(ctx, classInfo); err != nil {
		return nil, err
	}
	return &v1beta1.MsgUpdateClassAdminResponse{}, err
}

func (k Keeper) UpdateClassIssuers(ctx context.Context, req *v1beta1.MsgUpdateClassIssuers) (*v1beta1.MsgUpdateClassIssuersResponse, error) {
	class, err := k.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}
	reqAddr, _ := sdk.AccAddressFromBech32(req.Admin)
	admin := sdk.AccAddress(class.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", class.Admin, req.Admin)
	}

	for _, issuer := range req.RemoveIssuers {
		if err = k.stateStore.ClassIssuerStore().Delete(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: class.Id,
			Issuer:  issuer,
		}); err != nil {
			return nil, err
		}
	}

	// add the new issuers
	for _, issuer := range req.AddIssuers {
		if err = k.stateStore.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: class.Id,
			Issuer:  issuer,
		}); err != nil {
			return nil, err
		}
	}
	return &v1beta1.MsgUpdateClassIssuersResponse{}, nil
}

func (k Keeper) UpdateClassMetadata(ctx context.Context, req *v1beta1.MsgUpdateClassMetadata) (*v1beta1.MsgUpdateClassMetadataResponse, error) {
	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}
	reqAddr, _ := sdk.AccAddressFromBech32(req.Admin)
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
