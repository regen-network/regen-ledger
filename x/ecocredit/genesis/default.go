package genesis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

// DefaultParams returns a default set of parameters.
func DefaultParams() basetypes.Params {
	return basetypes.NewParams(
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, basetypes.DefaultCreditClassFee)),
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, basetypes.DefaultBasketFee)),
		[]string{},
		false,
	)
}

// DefaultCreditTypes returns a default set of credit basetypes.
func DefaultCreditTypes() []basetypes.CreditType {
	return []basetypes.CreditType{
		{
			Name:         "carbon",
			Abbreviation: "C",
			Unit:         "metric ton CO2 equivalent",
			Precision:    basetypes.PRECISION,
		},
	}
}

// DefaultCreditClassFees returns a default credit class fees.
func DefaultCreditClassFees() basetypes.ClassFees {
	return basetypes.ClassFees{
		Fees: []*sdk.Coin{
			{
				Denom:  sdk.DefaultBondDenom,
				Amount: basetypes.DefaultCreditClassFee,
			},
		},
	}
}

// DefaultBasketFee returns a default basket creation fees.
func DefaultBasketFee() baskettypes.BasketFee {
	return baskettypes.BasketFee{
		Fee: &sdk.Coin{
			Denom:  sdk.DefaultBondDenom,
			Amount: basetypes.DefaultBasketFee,
		},
	}
}

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
