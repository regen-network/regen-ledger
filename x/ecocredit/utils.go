package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// AssertGovernance asserts the address is equal to the governance module address
func AssertGovernance(addr sdk.AccAddress, k AccountKeeper) error {
	if !k.GetModuleAddress(govtypes.ModuleName).Equals(addr) {
		return sdkerrors.ErrInvalidRequest.Wrapf("params can only be updated via governance")
	}
	return nil
}
