package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (k Keeper) AddCreditType(ctx context.Context, req *core.MsgAddCreditType) (*core.MsgAddCreditTypeResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}

	store := k.stateStore.CreditTypeStore()
	for _, ct := range req.CreditTypes {
		if err = store.Insert(ctx, &ecocreditv1.CreditType{
			Abbreviation: ct.Abbreviation,
			Name:         ct.Name,
			Unit:         ct.Unit,
			Precision:    ct.Precision,
		}); err != nil {
			return nil, err
		}
	}

	return &core.MsgAddCreditTypeResponse{}, nil
}

func (k Keeper) ToggleAllowList(ctx context.Context, req *core.MsgToggleAllowListRequest) (*core.MsgToggleAllowListResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}
	return &core.MsgToggleAllowListResponse{}, k.stateStore.AllowlistEnabledStore().Save(ctx, &ecocreditv1.AllowlistEnabled{Enabled: req.Toggle})
}

func (k Keeper) UpdateAllowedCreditClassCreators(ctx context.Context, req *core.MsgUpdateAllowedCreditClassCreatorsRequest) (*core.MsgUpdateAllowedCreditClassCreatorsResponse, error) {
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
		if err = store.Insert(ctx, &ecocreditv1.AllowedClassCreators{Address: acc}); err != nil {
			return nil, err
		}
	}

	for _, addr := range req.RemoveCreators {
		acc, err := types.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		if err = store.Delete(ctx, &ecocreditv1.AllowedClassCreators{Address: acc}); err != nil {
			return nil, err
		}
	}

	return &core.MsgUpdateAllowedCreditClassCreatorsResponse{}, nil
}

func (k Keeper) UpdateCreditClassFee(ctx context.Context, req *core.MsgUpdateCreditClassFeeRequest) (*core.MsgUpdateCreditClassFeeResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}

	store := k.stateStore.CreditClassFeeStore()
	for _, fee := range req.AddFees {
		if err = store.Insert(ctx, &ecocreditv1.CreditClassFee{
			Denom:  fee.Denom,
			Amount: fee.Amount,
		}); err != nil {
			return nil, err
		}
	}

	for _, fee := range req.RemoveFees {
		if err = store.Delete(ctx, &ecocreditv1.CreditClassFee{
			Denom: fee.Denom,
		}); err != nil {
			return nil, err
		}
	}

	return &core.MsgUpdateCreditClassFeeResponse{}, nil
}

func (k Keeper) UpdateBasketFee(ctx context.Context, req *core.MsgUpdateBasketFeeRequest) (*core.MsgUpdateBasketFeeResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = k.assertGovernance(govAddr); err != nil {
		return nil, err
	}

	store := k.stateStore.BasketFeeStore()
	for _, fee := range req.AddFees {
		if err = store.Insert(ctx, &ecocreditv1.BasketFee{
			Denom:  fee.Denom,
			Amount: fee.Amount,
		}); err != nil {
			return nil, err
		}
	}

	for _, fee := range req.RemoveFees {
		if err = store.Delete(ctx, &ecocreditv1.BasketFee{
			Denom: fee.Denom,
		}); err != nil {
			return nil, err
		}
	}

	return &core.MsgUpdateBasketFeeResponse{}, nil
}

func (k Keeper) assertGovernance(addr types.AccAddress) error {
	if !k.accountKeeper.GetModuleAddress(govtypes.ModuleName).Equals(addr) {
		return sdkerrors.ErrUnauthorized.Wrapf("params can only be updated via governance")
	}
	return nil
}
