package ecocredit

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// TODO: Decide a sensible default value
	DefaultCreditClassFeeTokens = sdk.NewInt(10000)
	KeyCreditClassFee           = []byte("CreditClassFee")
	KeyCreditTypes              = []byte("CreditTypes")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCreditClassFee, &p.CreditClassFee, validateCreditClassFee),
		paramtypes.NewParamSetPair(KeyCreditTypes, &p.CreditTypes, validateCreditTypes),
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

func validateCreditTypes(i interface{}) error {
	v, ok := i.([]*CreditType)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// ensure no duplicate credit types
	seenTypes := make(map[string]bool)
	for _, ct := range v {
		t := strings.ToLower(ct.Type)
		if seenTypes[t] == true {
			return fmt.Errorf("duplicate credit type: %s", t)
		}
		seenTypes[t] = true
	}

	return nil
}

func NewParams(creditClassFee sdk.Coins, creditTypes []*CreditType) Params {
	return Params{
		CreditClassFee: creditClassFee,
		CreditTypes:    creditTypes,
	}
}

func DefaultParams() Params {
	return NewParams(sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultCreditClassFeeTokens)), []*CreditType{})
}
