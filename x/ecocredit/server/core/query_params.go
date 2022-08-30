package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Params queries the ecocredit module parameters.
// Deprecated: This rpc method is deprecated and will be removed in next version.
// Use individual param query instead.
func (k Keeper) Params(ctx context.Context, _ *core.QueryParamsRequest) (*core.QueryParamsResponse, error) {

	allowlistEnabled, err := k.stateStore.AllowListEnabledTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	itr, err := k.stateStore.AllowedClassCreatorTable().List(ctx, ecocreditv1.AllowedClassCreatorPrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer itr.Close()

	creators := []string{}
	for itr.Next() {
		val, err := itr.Value()
		if err != nil {
			return nil, err
		}

		creators = append(creators, sdk.AccAddress(val.Address).String())

	}

	classFees, err := k.stateStore.ClassFeesTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	classFees1, ok := types.ProtoCoinsToCoins(classFees.Fees)
	if !ok {
		return nil, sdkerrors.ErrInvalidCoins.Wrap("class fees")
	}

	basketFees, err := k.basketStore.BasketFeesTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	basketFees1, ok := types.ProtoCoinsToCoins(basketFees.Fees)
	if !ok {
		return nil, sdkerrors.ErrInvalidCoins.Wrap("basket fees")
	}

	allowedDenomsItr, err := k.marketStore.AllowedDenomTable().List(ctx, marketplacev1.AllowedDenomPrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer allowedDenomsItr.Close()

	var allowedDenoms []*core.AllowedDenomInfo
	for allowedDenomsItr.Next() {
		val, err := allowedDenomsItr.Value()
		if err != nil {
			return nil, err
		}

		allowedDenoms = append(allowedDenoms, &core.AllowedDenomInfo{
			BankDenom:    val.BankDenom,
			DisplayDenom: val.DisplayDenom,
			Exponent:     val.Exponent,
		})
	}

	return &core.QueryParamsResponse{
		Params: &core.Params{
			AllowedClassCreators: creators,
			AllowlistEnabled:     allowlistEnabled.Enabled,
			CreditClassFee:       classFees1,
			BasketFee:            basketFees1,
		},
		AllowedDenoms: allowedDenoms,
	}, nil
}
