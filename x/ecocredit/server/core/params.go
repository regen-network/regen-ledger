package core

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// AddCreditType is a governance only function that allows the addition of credit types to the credit types chain parameter
func (k Keeper) AddCreditType(ctx context.Context, req *core.MsgAddCreditType) (*core.MsgAddCreditTypeResponse, error) {
	if err := ecocredit.AssertGovernance(req.RootAddress, k.accountKeeper); err != nil {
		return nil, err
	}

	store := k.stateStore.CreditTypeTable()
	for _, ct := range req.CreditTypes {
		if ct.Precision != ecocredit.PRECISION {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("invalid precision: credit type precision is currently locked to %d", ecocredit.PRECISION)
		}
		if err := store.Insert(ctx, &api.CreditType{
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

// ToggleAllowList is a governance only function that toggles the allowlist, enabling it if the request contains `True`,
// and disabling it if the request contains `false`
func (k Keeper) ToggleAllowList(ctx context.Context, req *core.MsgToggleAllowList) (*core.MsgToggleAllowListResponse, error) {
	if err := ecocredit.AssertGovernance(req.RootAddress, k.accountKeeper); err != nil {
		return nil, err
	}
	return &core.MsgToggleAllowListResponse{}, k.stateStore.AllowlistEnabledTable().Save(ctx, &api.AllowlistEnabled{Enabled: req.Toggle})
}

// UpdateAllowedCreditClassCreators is a governance only function that allows for the removal and addition of addresses
// to the credit class creator list
// NOTE: this list is only effective when the governance controlled AllowlistEnabled parameter is true.
func (k Keeper) UpdateAllowedCreditClassCreators(ctx context.Context, req *core.MsgUpdateAllowedCreditClassCreators) (*core.MsgUpdateAllowedCreditClassCreatorsResponse, error) {
	if err := ecocredit.AssertGovernance(req.RootAddress, k.accountKeeper); err != nil {
		return nil, err
	}

	store := k.stateStore.AllowedClassCreatorTable()

	for _, addr := range req.RemoveCreators {
		acc, err := types.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		if err = store.Delete(ctx, &api.AllowedClassCreator{Address: acc}); err != nil {
			return nil, err
		}
	}

	for _, addr := range req.AddCreators {
		acc, err := types.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		if err = store.Insert(ctx, &api.AllowedClassCreator{Address: acc}); err != nil {
			return nil, err
		}
	}

	return &core.MsgUpdateAllowedCreditClassCreatorsResponse{}, nil
}

// UpdateCreditClassFee is a governance only function that allows for the removal and addition of fees one can pay to create a class
func (k Keeper) UpdateCreditClassFee(ctx context.Context, req *core.MsgUpdateCreditClassFee) (*core.MsgUpdateCreditClassFeeResponse, error) {
	if err := ecocredit.AssertGovernance(req.RootAddress, k.accountKeeper); err != nil {
		return nil, err
	}

	store := k.stateStore.CreditClassFeeTable()

	for _, denom := range req.RemoveDenoms {
		if err := store.Delete(ctx, &api.CreditClassFee{
			Denom: denom,
		}); err != nil {
			return nil, err
		}
	}

	for _, fee := range req.AddFees {
		if err := store.Insert(ctx, &api.CreditClassFee{
			Denom:  fee.Denom,
			Amount: fee.Amount.String(),
		}); err != nil {
			return nil, err
		}
	}

	return &core.MsgUpdateCreditClassFeeResponse{}, nil
}
