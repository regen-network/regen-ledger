package ecocredit

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// TODO: Decide a sensible default value
	DefaultCreditClassFeeTokens = sdk.NewInt(10000)
	KeyCreditClassFee           = []byte("CreditClassFee")
	KeyAllowedClassDesigners    = []byte("AllowedClassDesigners")
	KeyAllowlistEnabled         = []byte("AllowlistEnabled")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCreditClassFee, &p.CreditClassFee, validateCreditClassFee),
		paramtypes.NewParamSetPair(KeyAllowedClassDesigners, &p.AllowedClassDesigners, validateAllowlistCreditDesigners),
		paramtypes.NewParamSetPair(KeyAllowlistEnabled, &p.AllowlistEnabled, validateAllowlistEnabled),
	}
}

func validateCreditClassFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return err
	}

	return nil
}

func validateAllowlistCreditDesigners(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, sAddr := range v {
		_, err := sdk.AccAddressFromBech32(sAddr)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateAllowlistEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func NewParams(creditClassFee sdk.Coins, allowlist []string, allowlistEnabled bool) Params {
	return Params{
		CreditClassFee:        creditClassFee,
		AllowedClassDesigners: allowlist,
		AllowlistEnabled:      allowlistEnabled,
	}
}

func DefaultParams() Params {
	return NewParams(sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultCreditClassFeeTokens)), []string{}, false)
}
