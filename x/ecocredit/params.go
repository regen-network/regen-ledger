package ecocredit

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/util"
	"strings"
)

var (
	// TODO: Decide a sensible default value
	DefaultCreditClassFeeTokens = sdk.NewInt(10000)
	KeyCreditClassFee           = []byte("CreditClassFee")
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
	creditTypes, ok := i.([]*CreditType)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// ensure no duplicate credit types and that all precisions conform to hardcoded PRECISION above
	seenTypes := make(map[string]bool)
	for _, creditType := range creditTypes {
		T := strings.ToLower(creditType.Name)
		T = util.FastRemoveWhitespace(T)
		if T != creditType.Name {
			return sdkerrors.ErrInvalidRequest.Wrapf("credit type should be normalized: got %s, should be %s", creditType.Name, T)
		}
		if seenTypes[T] == true {
			return sdkerrors.ErrInvalidRequest.Wrapf("duplicate credit types in request: %s", T)
		}
		// TODO: remove after we open governance changes for precision
		if creditType.Precision != PRECISION {
			return sdkerrors.ErrInvalidRequest.Wrapf("invalid precision %d: precision is currently locked to %d", creditType.Precision, PRECISION)
		}
		if creditType.Name == "" {
			return sdkerrors.ErrInvalidRequest.Wrap("empty credit type name")
		}
		if creditType.Units == "" {
			return sdkerrors.ErrInvalidRequest.Wrap("empty credit type units")
		}
		seenTypes[T] = true
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
	return NewParams(sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultCreditClassFeeTokens)), []*CreditType{{Name: "carbon", Units: "tons", Precision: PRECISION}})
}
