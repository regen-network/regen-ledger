package core

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

func (k Keeper) UpdateParams(ctx context.Context, msg *v1beta1.MsgUpdateParams) (*v1beta1.MsgUpdateParamsResponse, error) {
	root, err := sdk.AccAddressFromBech32(msg.RootAddress)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address %s", msg.RootAddress)
	}
	if err = k.onlyGovernance(root); err != nil {
		return nil, err
	}

	if msg.UpdateClassAllowlist != nil {
		// update
	}
	if msg.UpdateBasketFees != nil {
		// update
	}
	if msg.UpdateClassFees != nil {
		// update
	}
	if msg.ToggleAllowlist != nil {
		// update
	}
}

func (k Keeper) updateClassAllowlist(msg v1beta1.MsgUpdateAllowedClassCreators) error {
	return nil
}

func (k Keeper) onlyGovernance(addr sdk.AccAddress) error {
	if !k.ak.GetModuleAddress(govtypes.ModuleName).Equals(addr) {
		return sdkerrors.ErrUnauthorized.Wrapf("params can only be updated by governance")
	}
	return nil
}
