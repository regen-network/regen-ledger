package genesis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

// DefaultAllowedDenoms returns a default set of allowed denoms.
func DefaultAllowedDenoms() []markettypes.AllowedDenom {
	return []markettypes.AllowedDenom{
		{
			BankDenom:    sdk.DefaultBondDenom,
			DisplayDenom: sdk.DefaultBondDenom,
			Exponent:     6,
		},
	}
}

// DefaultBasketFees returns a default basket creation fees.
func DefaultBasketFees() baskettypes.BasketFees {
	return baskettypes.BasketFees{
		Fees: []*sdk.Coin{
			{
				Denom:  sdk.DefaultBondDenom,
				Amount: core.DefaultBasketFee,
			},
		},
	}
}
