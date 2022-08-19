package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the ClassIssuer state type
func (c ClassIssuer) Validate() error {
	if c.ClassKey == 0 {
		return ecocredit.ErrParseFailure.Wrap("class key cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(c.Issuer).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("issuer: %s", err)
	}

	return nil
}
