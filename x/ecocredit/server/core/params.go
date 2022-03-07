package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

func (k Keeper) NewCreditType(ctx context.Context, req *v1beta1.MsgNewCreditTypeRequest) (*v1beta1.MsgNewCreditTypeResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}

	store := k.stateStore.CreditTypeStore()
	for _, ct := range req.CreditTypes {
		if err = store.Insert(ctx, &ecocreditv1beta1.CreditType{
			Abbreviation: ct.Abbreviation,
			Name:         ct.Name,
			Unit:         ct.Unit,
			Precision:    ct.Precision,
		}); err != nil {
			return nil, err
		}
	}

	return &v1beta1.MsgNewCreditTypeResponse{}, nil
}

func (k Keeper) ToggleAllowList(ctx context.Context, req *v1beta1.MsgToggleAllowListRequest) (*v1beta1.MsgToggleAllowListResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}
	return &v1beta1.MsgToggleAllowListResponse{}, k.stateStore.AllowlistEnabledStore().Save(ctx, &ecocreditv1beta1.AllowlistEnabled{Enabled: req.Toggle})
}

func (k Keeper) UpdateAllowedCreditClassCreators(ctx context.Context, req *v1beta1.MsgUpdateAllowedCreditClassCreatorsRequest) (*v1beta1.MsgUpdateAllowedCreditClassCreatorsResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}

	store := k.stateStore.AllowedClassCreatorsStore()
	for _, addr := range req.AddCreators {
		acc, err := types.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		if err = store.Insert(ctx, &ecocreditv1beta1.AllowedClassCreators{Address: acc}); err != nil {
			return nil, err
		}
	}

	for _, addr := range req.RemoveCreators {
		acc, err := types.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		if err = store.Delete(ctx, &ecocreditv1beta1.AllowedClassCreators{Address: acc}); err != nil {
			return nil, err
		}
	}

	return &v1beta1.MsgUpdateAllowedCreditClassCreatorsResponse{}, nil
}

func (k Keeper) UpdateCreditClassFee(ctx context.Context, req *v1beta1.MsgUpdateCreditClassFeeRequest) (*v1beta1.MsgUpdateCreditClassFeeResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}

	store := k.stateStore.CreditClassFeeStore()
	for _, fee := range req.AddFees {
		if err = store.Insert(ctx, &ecocreditv1beta1.CreditClassFee{
			Denom:  fee.Denom,
			Amount: fee.Amount,
		}); err != nil {
			return nil, err
		}
	}

	for _, fee := range req.RemoveFees {
		if err = store.Delete(ctx, &ecocreditv1beta1.CreditClassFee{
			Denom: fee.Denom,
		}); err != nil {
			return nil, err
		}
	}

	return &v1beta1.MsgUpdateCreditClassFeeResponse{}, nil
}

func (k Keeper) UpdateBasketFee(ctx context.Context, req *v1beta1.MsgUpdateBasketFeeRequest) (*v1beta1.MsgUpdateBasketFeeResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}

	store := k.stateStore.BasketFeeStore()
	for _, fee := range req.AddFees {
		if err = store.Insert(ctx, &ecocreditv1beta1.BasketFee{
			Denom:  fee.Denom,
			Amount: fee.Amount,
		}); err != nil {
			return nil, err
		}
	}

	for _, fee := range req.RemoveFees {
		if err = store.Delete(ctx, &ecocreditv1beta1.BasketFee{
			Denom: fee.Denom,
		}); err != nil {
			return nil, err
		}
	}

	return &v1beta1.MsgUpdateBasketFeeResponse{}, nil
}

func (k Keeper) assertGovernance(addr types.AccAddress) error {
	if !k.ak.GetModuleAddress(govtypes.ModuleName).Equals(addr) {
		return sdkerrors.ErrUnauthorized.Wrapf("params can only be updated via governance")
	}
	return nil
}
