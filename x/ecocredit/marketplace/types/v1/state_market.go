package v1

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
)

// Validate performs basic validation of the Market state type
func (m *Market) Validate() error {
	if m.Id == 0 {
		return ecocredit.ErrParseFailure.Wrapf("id cannot be zero")
	}

	if err := base.ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return errors.Wrap(err, "credit type abbrev") // returns parse error
	}

	if m.BankDenom == "" {
		return ecocredit.ErrParseFailure.Wrap("bank denom cannot be empty")
	}

	if err := sdk.ValidateDenom(m.BankDenom); err != nil {
		return ecocredit.ErrParseFailure.Wrapf("bank denom: %s", err)
	}

	if m.PrecisionModifier != 0 {
		return ecocredit.ErrParseFailure.Wrapf("precision modifier must be zero")
	}

	return nil
}
