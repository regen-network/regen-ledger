package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// Params queries the ecocredit module parameters.
// Deprecated: This rpc method is deprecated and will be removed in next version.
// Use individual param query instead.
func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {

	allowlistEnabled, err := k.stateStore.ClassCreatorAllowlistTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	itr, err := k.stateStore.AllowedClassCreatorTable().List(ctx, api.AllowedClassCreatorPrimaryKey{})
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

	classFee, err := k.stateStore.ClassFeeTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	classFeeCoin, ok := regentypes.ProtoCoinToCoin(classFee.Fee)
	if !ok {
		return nil, sdkerrors.ErrInvalidCoins.Wrap("class fees")
	}

	basketFee, err := k.basketStore.BasketFeeTable().Get(ctx)
	if err != nil {
		return nil, err
	}

	basketFeeCoin, ok := regentypes.ProtoCoinToCoin(basketFee.Fee)
	if !ok {
		return nil, sdkerrors.ErrInvalidCoins.Wrap("basket fees")
	}

	allowedDenomsItr, err := k.marketStore.AllowedDenomTable().List(ctx, marketplacev1.AllowedDenomPrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer allowedDenomsItr.Close()

	var allowedDenoms []*types.AllowedDenomInfo
	for allowedDenomsItr.Next() {
		val, err := allowedDenomsItr.Value()
		if err != nil {
			return nil, err
		}

		allowedDenoms = append(allowedDenoms, &types.AllowedDenomInfo{
			BankDenom:    val.BankDenom,
			DisplayDenom: val.DisplayDenom,
			Exponent:     val.Exponent,
		})
	}

	var allowedBridgeChains []string
	it, err := k.stateStore.AllowedBridgeChainTable().List(ctx, api.AllowedBridgeChainPrimaryKey{})
	if err != nil {
		return nil, err
	}
	defer it.Close()

	for it.Next() {
		entry, err := it.Value()
		if err != nil {
			return nil, err
		}
		allowedBridgeChains = append(allowedBridgeChains, entry.ChainName)
	}

	return &types.QueryParamsResponse{
		Params: &types.Params{
			AllowedClassCreators: creators,
			AllowlistEnabled:     allowlistEnabled.Enabled,
			CreditClassFee:       sdk.Coins{classFeeCoin},
			BasketFee:            sdk.Coins{basketFeeCoin},
		},
		AllowedDenoms:       allowedDenoms,
		AllowedBridgeChains: allowedBridgeChains,
	}, nil
}
