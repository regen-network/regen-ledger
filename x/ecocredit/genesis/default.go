package genesis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

// DefaultAllowedDenoms returns a default set of allowed denoms.
func DefaultAllowedDenoms() []types.AllowedDenom {
	return []types.AllowedDenom{
		{
			BankDenom:    sdk.DefaultBondDenom,
			DisplayDenom: sdk.DefaultBondDenom,
			Exponent:     6,
		},
	}
}
