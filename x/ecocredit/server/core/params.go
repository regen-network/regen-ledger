package core

import (
	"context"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	"github.com/cosmos/cosmos-sdk/types"
)

// AddCreditType is a governance only function that allows the addition of credit types to the credit types chain parameter
func (k Keeper) AddCreditType(ctx context.Context, req *core.MsgAddCreditType) (*core.MsgAddCreditTypeResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = ecocredit.AssertGovernance(govAddr, k.accountKeeper); err != nil {
		return nil, err
	}

	// TODO: validate fields? should abbreviation be all uppercase? no numbers? no special characters?

	store := k.stateStore.CreditTypeTable()
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

// ToggleAllowList is a governance only function that toggles the allowlist, enabling it if the request contains `True`,
// and disabling it if the request contains `false`
func (k Keeper) ToggleAllowList(ctx context.Context, req *core.MsgToggleAllowListRequest) (*core.MsgToggleAllowListResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = ecocredit.AssertGovernance(govAddr, k.accountKeeper); err != nil {
		return nil, err
	}
	return &core.MsgToggleAllowListResponse{}, k.stateStore.AllowlistEnabledTable().Save(ctx, &ecocreditv1.AllowlistEnabled{Enabled: req.Toggle})
}

// UpdateAllowedCreditClassCreators is a governance only function that allows for the removal and addition of addresses
// to the credit class creator list
// NOTE: this list is only effective when the governance controlled AllowlistEnabled parameter is true.
func (k Keeper) UpdateAllowedCreditClassCreators(ctx context.Context, req *core.MsgUpdateAllowedCreditClassCreatorsRequest) (*core.MsgUpdateAllowedCreditClassCreatorsResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = ecocredit.AssertGovernance(govAddr, k.accountKeeper); err != nil {
		return nil, err
	}

	store := k.stateStore.AllowedClassCreatorsTable()

	for _, addr := range req.RemoveCreators {
		acc, err := types.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		if err = store.Delete(ctx, &ecocreditv1.AllowedClassCreators{Address: acc}); err != nil {
			return nil, err
		}
	}

	for _, addr := range req.AddCreators {
		acc, err := types.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		if err = store.Insert(ctx, &ecocreditv1.AllowedClassCreators{Address: acc}); err != nil {
			return nil, err
		}
	}

	return &core.MsgUpdateAllowedCreditClassCreatorsResponse{}, nil
}

// UpdateCreditClassFee is a governance only function that allows for the removal and addition of fees one can pay to create a class
func (k Keeper) UpdateCreditClassFee(ctx context.Context, req *core.MsgUpdateCreditClassFeeRequest) (*core.MsgUpdateCreditClassFeeResponse, error) {
	govAddr, err := types.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = ecocredit.AssertGovernance(govAddr, k.accountKeeper); err != nil {
		return nil, err
	}

	store := k.stateStore.CreditClassFeeTable()

	for _, fee := range req.RemoveFees {
		if err = store.Delete(ctx, &ecocreditv1.CreditClassFee{
			Denom: fee.Denom,
		}); err != nil {
			return nil, err
		}
	}

	for _, fee := range req.AddFees {
		if err = store.Insert(ctx, &ecocreditv1.CreditClassFee{
			Denom:  fee.Denom,
			Amount: fee.Amount,
		}); err != nil {
			return nil, err
		}
	}

	return &core.MsgUpdateCreditClassFeeResponse{}, nil
}
