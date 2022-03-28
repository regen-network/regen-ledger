package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// AssertGovernance asserts the address is equal to the governance module address
func AssertGovernance(addrStr string, k AccountKeeper) error {
	addr, err := sdk.AccAddressFromBech32(addrStr)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if !k.GetModuleAddress(govtypes.ModuleName).Equals(addr) {
		return sdkerrors.ErrInvalidRequest.Wrapf("params can only be updated via governance")
	}
	return nil
}
