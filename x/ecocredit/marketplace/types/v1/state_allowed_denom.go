package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Validate performs basic validation of the AllowedDenom state type
func (m *AllowedDenom) Validate() error {
	if m.BankDenom == "" {
		return ecocredit.ErrParseFailure.Wrap("bank denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.BankDenom); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("bank denom: %s", err)
	}

	if m.DisplayDenom == "" {
		return ecocredit.ErrParseFailure.Wrap("display denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.DisplayDenom); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("display denom: %s", err)
	}

	if _, err := core.ExponentToPrefix(m.Exponent); err != nil {
		return err // returns parse error
	}

	return nil
}
