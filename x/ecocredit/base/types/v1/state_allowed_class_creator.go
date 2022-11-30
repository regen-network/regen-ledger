package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
)

// Validate performs basic validation of the AllowedClassCreator state type
func (m *AllowedClassCreator) Validate() error {
	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Address).String()); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("address: %s", err)
	}

	return nil
}
