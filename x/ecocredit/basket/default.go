package basket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// DefaultBasketFees returns a default basket creation fees.
func DefaultBasketFees() BasketFees {
	return BasketFees{
		Fees: []*sdk.Coin{
			{
				Denom:  sdk.DefaultBondDenom,
				Amount: core.DefaultBasketFee,
			},
		},
	}
}
