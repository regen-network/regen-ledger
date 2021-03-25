// Package math provides helper functions for doing mathematical calculations and parsing for the ecocredit module.
package math

import (
	"fmt"

	"github.com/cockroachdb/apd/v2"

	"github.com/cosmos/cosmos-sdk/types/errors"
)

type Dec struct {
	dec apd.Decimal
}

var dec128Context = apd.Context{
	Precision:   34,
	MaxExponent: apd.MaxExponent,
	MinExponent: apd.MinExponent,
	Traps:       apd.DefaultTraps,
}

var exactContext = apd.Context{
	Precision:   0,
	MaxExponent: apd.MaxExponent,
	MinExponent: apd.MinExponent,
	Traps:       apd.DefaultTraps | apd.Inexact | apd.Rounded,
}

func NewDecFromString(s string) (Dec, error) {
	d, _, err := apd.NewFromString(s)
	if err != nil {
		return Dec{}, err
	}
	return Dec{*d}, nil
}

func NewDecFromInt64(x int64) Dec {
	var res Dec
	res.dec.SetInt64(x)
	return res
}

func (x Dec) Add(y Dec) (Dec, error) {
	var z Dec
	_, err := apd.BaseContext.Add(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal addition error")
}

func (x Dec) Sub(y Dec) (Dec, error) {
	var z Dec
	_, err := apd.BaseContext.Sub(&z.dec, &x.dec, &y.dec)
	return z, err
}

// SafeSub subtracts the value of y from x and stores the result in res with arbitrary precision only
// if the result will be non-negative. An insufficient funds error is returned if the result would be negative.
func (x Dec) SafeSub(y Dec) error {
	var z Dec
	_, err := exactContext.Sub(&z.dec, &x.dec, &y.dec)
	if err != nil {
		return errors.Wrap(err, "decimal subtraction error")
	}

	if z.IsNegative() {
		return errors.ErrInsufficientFunds
	}

	return nil
}

func (x Dec) Quo(y Dec) (Dec, error) {
	var z Dec
	_, err := dec128Context.Quo(&z.dec, &x.dec, &y.dec)
	return z, err
}

func (x Dec) QuoInteger(y Dec) (Dec, error) {
	var z Dec
	_, err := dec128Context.QuoInteger(&z.dec, &x.dec, &y.dec)
	return z, err
}

func (x Dec) Rem(y Dec) (Dec, error) {
	var z Dec
	_, err := dec128Context.Rem(&z.dec, &x.dec, &y.dec)
	return z, err
}

func (x Dec) Mul(y Dec) (Dec, error) {
	var z Dec
	_, err := dec128Context.Mul(&z.dec, &x.dec, &y.dec)
	return z, err
}

func (x Dec) Int64() (int64, error) {
	return x.dec.Int64()
}

func (x Dec) String() string {
	return x.dec.Text('f')
}

func (x Dec) Cmp(y Dec) int {
	return x.dec.Cmp(&y.dec)
}

func (x Dec) IsEqual(y Dec) bool {
	return x.dec.Cmp(&y.dec) == 0
}

func (x Dec) IsZero() bool {
	return x.dec.IsZero()
}

func (x Dec) IsNegative() bool {
	return x.dec.Negative && !x.dec.IsZero()
}

func (x Dec) IsPositive() bool {
	return !x.dec.Negative && !x.dec.IsZero()
}

// NumDecimalPlaces returns the number of decimal places in x.
func (x Dec) NumDecimalPlaces() uint32 {
	exp := x.dec.Exponent
	if exp >= 0 {
		return 0
	}
	return uint32(-exp)
}

func (x Dec) Reduce() (Dec, int) {
	y := Dec{}
	_, n := y.dec.Reduce(&x.dec)
	return y, n
}

// ParseNonNegativeDecimal parses a non-negative decimal or returns an error.
func ParseNonNegativeDecimal(x string) (Dec, error) {
	res, err := NewDecFromString(x)
	if err != nil || res.IsNegative() {
		return Dec{}, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a non-negative decimal, got %s", x))
	}

	return res, nil
}

// ParsePositiveDecimal parses a positive decimal or returns an error.
func ParsePositiveDecimal(x string) (Dec, error) {
	res, err := NewDecFromString(x)
	if err != nil || !res.IsPositive() {
		return Dec{}, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a positive decimal, got %s", x))
	}

	return res, nil
}

// ParseNonNegativeDecimal parses a non-negative decimal with a fixed maxDecimalPlaces or returns an error.
func ParseNonNegativeFixedDecimal(x string, maxDecimalPlaces uint32) (Dec, error) {
	res, err := ParseNonNegativeDecimal(x)
	if err != nil {
		return Dec{}, err
	}

	err = requireMaxDecimals(res, maxDecimalPlaces)
	if err != nil {
		return Dec{}, err
	}

	return res, nil
}

// ParsePositiveFixedDecimal parses a positive decimal with a fixed maxDecimalPlaces or returns an error.
func ParsePositiveFixedDecimal(x string, maxDecimalPlaces uint32) (Dec, error) {
	res, err := ParsePositiveDecimal(x)
	if err != nil {
		return Dec{}, err
	}

	err = requireMaxDecimals(res, maxDecimalPlaces)
	if err != nil {
		return Dec{}, err
	}

	return res, nil
}

func requireMaxDecimals(x Dec, maxDecimalPlaces uint32) error {
	n := x.NumDecimalPlaces()
	if n > maxDecimalPlaces {
		return errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected no more than %d decimal places in %s, got %d", maxDecimalPlaces, x, n))
	}
	return nil
}
