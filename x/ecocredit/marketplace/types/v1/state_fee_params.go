package v1

import "github.com/regen-network/regen-ledger/types/v2/math"

// Validate performs basic validation of the FeeParams state type.
func (m *FeeParams) Validate() error {
	_, err := math.NewPositiveDecFromString(m.BuyerPercentageFee)
	if err != nil {
		return err
	}

	_, err = math.NewPositiveDecFromString(m.SellerPercentageFee)
	if err != nil {
		return err
	}

	return nil
}
