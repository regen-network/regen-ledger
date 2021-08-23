package ecocredit

import (
	"fmt"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// TODO: Decide a sensible default value
	DefaultCreditClassFeeTokens = sdk.NewInt(10000)
	KeyCreditClassFee           = []byte("CreditClassFee")
	KeyAllowedClassDesigners    = []byte("AllowedClassDesigners")
	KeyAllowlistEnabled         = []byte("AllowlistEnabled")
	KeyCreditTypes              = []byte("CreditTypes")
)

// TODO: remove after we open governance changes for precision
const (
	PRECISION uint32 = 6
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

func validateCreditTypes(i interface{}) error {
	creditTypes, ok := i.([]*CreditType)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// ensure no duplicate credit types or abbreviations and that all
	// precisions conform to hardcoded PRECISION above
	seenTypes := make(map[string]bool)
	seenAbbrs := make(map[string]bool)
	for _, creditType := range creditTypes {
		// Validate name
		T := NormalizeCreditTypeName(creditType.Name)
		if T != creditType.Name {
			return sdkerrors.ErrInvalidRequest.Wrapf("credit type name should be normalized: got %s, should be %s", creditType.Name, T)
		}
		if creditType.Name == "" {
			return sdkerrors.ErrInvalidRequest.Wrap("empty credit type name")
		}
		if seenTypes[T] {
			return sdkerrors.ErrInvalidRequest.Wrapf("duplicate credit types in request: %s", T)
		}

		// Validate abbreviation
		abbr := creditType.Abbreviation
		err := validateCreditTypeAbbreviation(abbr)
		if err != nil {
			return err
		}
		if seenAbbrs[abbr] {
			return sdkerrors.ErrInvalidRequest.Wrapf("duplicate credit type abbreviation: %s", abbr)
		}

		// Validate precision
		// TODO: remove after we open governance changes for precision
		if creditType.Precision != PRECISION {
			return sdkerrors.ErrInvalidRequest.Wrapf("invalid precision %d: precision is currently locked to %d", creditType.Precision, PRECISION)
		}

		// Validate units
		if creditType.Unit == "" {
			return sdkerrors.ErrInvalidRequest.Wrap("empty credit type unit")
		}

		// Mark type and abbr as seen
		seenTypes[T] = true
		seenAbbrs[abbr] = true
	}

	return nil
}

// Check that CreditType abbreviation is valid, i.e. it consists of 1-3
// uppercase letters
func validateCreditTypeAbbreviation(abbr string) error {
	reAbbr := regexp.MustCompile(`^[A-Z]{1,3}$`)
	matches := reAbbr.FindStringSubmatch(abbr)
	if matches == nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit type abbreviation must be 1-3 uppercase letters: got %s", abbr)
	}
	return nil
}

func NewParams(creditClassFee sdk.Coins, allowlist []string, allowlistEnabled bool, creditTypes []*CreditType) Params {
	return Params{
		CreditClassFee:        creditClassFee,
		AllowedClassDesigners: allowlist,
		AllowlistEnabled:      allowlistEnabled,
		CreditTypes:           creditTypes,
	}
}

func DefaultParams() Params {
	return NewParams(
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultCreditClassFeeTokens)),
		[]string{},
		false,
		[]*CreditType{
			{
				Name:         "carbon",
				Abbreviation: "C",
				Unit:         "metric ton CO2 equivalent",
				Precision:    PRECISION,
			},
		},
	)
}
