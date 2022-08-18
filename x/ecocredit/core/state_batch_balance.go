package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/math"
)

// Validate performs basic validation of the BatchBalance state type
func (b BatchBalance) Validate() error {
	if b.BatchKey == 0 {
		return fmt.Errorf("batch key cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(b.Address).String()); err != nil {
		return fmt.Errorf("address: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(b.TradableAmount); err != nil {
		return fmt.Errorf("tradable amount: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(b.RetiredAmount); err != nil {
		return fmt.Errorf("retired amount: %s", err)
	}

	if _, err := math.NewNonNegativeDecFromString(b.EscrowedAmount); err != nil {
		return fmt.Errorf("escrowed amount: %s", err)
	}

	return nil
}
