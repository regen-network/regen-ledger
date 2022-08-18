package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate performs basic validation of the Batch state type
func (b Batch) Validate() error {
	if err := ValidateBatchDenom(b.Denom); err != nil {
		return err
	}

	if b.ProjectKey == 0 {
		return fmt.Errorf("project key cannot be zero")
	}

	if b.StartDate == nil {
		return fmt.Errorf("must provide a start date for the credit batch")
	}

	if b.EndDate == nil {
		return fmt.Errorf("must provide an end date for the credit batch")
	}

	if b.EndDate.Compare(*b.StartDate) != 1 {
		return fmt.Errorf(
			"the batch end date (%s) must be the same as or after the batch start date (%s)",
			b.EndDate.String(),
			b.StartDate.String(),
		)
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(b.Issuer).String()); err != nil {
		return fmt.Errorf("issuer: %s", err)
	}

	return nil
}
