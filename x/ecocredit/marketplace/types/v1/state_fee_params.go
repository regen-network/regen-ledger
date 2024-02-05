package v1

import "github.com/regen-network/regen-ledger/types/v2/math"

// Validate performs basic validation of the FeeParams state type.
func (m *FeeParams) Validate() error {
	_, err := math.NewNonNegativeDecFromString(m.BuyerPercentageFee)
	if err != nil {
		return err
	}

	_, err = math.NewNonNegativeDecFromString(m.SellerPercentageFee)
	if err != nil {
		return err
	}

	return nil
}
