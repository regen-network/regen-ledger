package genesis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
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
