package v1

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the BasketFee state type
func (m *BasketFee) Validate() error {
	if m.Fee != nil {
		if m.Fee.Denom == "" {
			return ecocredit.ErrParseFailure.Wrap("fee: denom cannot be empty")
		}

		if m.Fee.Amount.IsNil() {
			return ecocredit.ErrParseFailure.Wrap("fee: amount cannot be empty")
		}

		if err := m.Fee.Validate(); err != nil {
			return ecocredit.ErrParseFailure.Wrapf("fee: %s", err)
		}
	}

	return nil
}
