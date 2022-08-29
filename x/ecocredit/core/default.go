package core

import sdk "github.com/cosmos/cosmos-sdk/types"

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultCreditClassFee)),
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultBasketFee)),
		[]string{},
		false,
	)
}

// DefaultCreditTypes returns a default credit class fees.
func DefaultCreditClassFees() ClassFees {
	return ClassFees{
		Fees: []*sdk.Coin{
			{
				Denom:  sdk.DefaultBondDenom,
				Amount: DefaultCreditClassFee,
			},
		},
	}
}

// DefaultCreditTypes returns a default set of credit types.
func DefaultCreditTypes() []CreditType {
	return []CreditType{
		{
			Name:         "carbon",
			Abbreviation: "C",
			Unit:         "metric ton CO2 equivalent",
			Precision:    PRECISION,
		},
	}
}
