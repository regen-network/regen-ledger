package v1

import (
	"cosmossdk.io/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
)

// Validate performs basic validation of the BasketClass state type
func (m *BasketClass) Validate() error {
	if m.BasketId == 0 {
		return ecocredit.ErrParseFailure.Wrapf("basket id cannot be zero")
	}

	if err := base.ValidateClassID(m.ClassId); err != nil {
		return errors.Wrap(err, "class id") // returns parse error
	}

	return nil
}
