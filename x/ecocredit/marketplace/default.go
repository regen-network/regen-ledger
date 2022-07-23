package marketplace

import sdk "github.com/cosmos/cosmos-sdk/types"

// DefaultAllowedDenoms returns a default set of allowed denoms.
func DefaultAllowedDenoms() []AllowedDenom {
	return []AllowedDenom{
		{
			BankDenom:    sdk.DefaultBondDenom,
			DisplayDenom: sdk.DefaultBondDenom,
			Exponent:     6,
		},
	}
}
