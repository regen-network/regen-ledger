package core

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// This is a value of 20 REGEN
	DefaultCreditClassFee   = math.NewInt(2e7)
	DefaultBasketFee        = math.NewInt(2e7)
	KeyCreditClassFee       = []byte("CreditClassFee")
	KeyAllowedClassCreators = []byte("AllowedClassCreators")
	KeyAllowlistEnabled     = []byte("AllowlistEnabled")
	KeyBasketFee            = []byte("BasketFee")
)

// TODO: remove after we allow standard SI units for precision

const (
	PRECISION uint32 = 6
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCreditClassFee, &p.CreditClassFee, validateCreditClassFee),
		paramtypes.NewParamSetPair(KeyAllowedClassCreators, &p.AllowedClassCreators, validateAllowedClassCreators),
		paramtypes.NewParamSetPair(KeyAllowlistEnabled, &p.AllowlistEnabled, validateAllowlistEnabled),
		paramtypes.NewParamSetPair(KeyBasketFee, &p.BasketFee, validateBasketFee),
	}
}

// Validate will run each param field's validate method
func (p Params) Validate() error {
	if err := validateAllowedClassCreators(p.AllowedClassCreators); err != nil {
		return err
	}

	if err := validateAllowlistEnabled(p.AllowlistEnabled); err != nil {
		return err
	}

	if err := validateCreditClassFee(p.CreditClassFee); err != nil {
		return err
	}

	if err := validateBasketFee(p.BasketFee); err != nil {
		return err
	}

	return nil
}

func validateCreditClassFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return sdkerrors.ErrInvalidType.Wrapf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return err
	}

	return nil
}

func validateAllowedClassCreators(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return sdkerrors.ErrInvalidType.Wrapf("invalid parameter type: %T", i)
	}
	for _, sAddr := range v {
		_, err := sdk.AccAddressFromBech32(sAddr)
		if err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("invalid creator address: %s", err.Error())
		}
	}
	return nil
}

func validateAllowlistEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return sdkerrors.ErrInvalidType.Wrapf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBasketFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return sdkerrors.ErrInvalidType.Wrapf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return err
	}

	return nil
}

// NewParams creates a new Params object.
func NewParams(creditClassFee, basketFee sdk.Coins, allowlist []string, allowlistEnabled bool) Params {
	return Params{
		CreditClassFee:       creditClassFee,
		AllowedClassCreators: allowlist,
		AllowlistEnabled:     allowlistEnabled,
		BasketFee:            basketFee,
	}
}
