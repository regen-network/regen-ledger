// Package math provides helper functions for doing mathematical calculations and parsing for the ecocredit module.
package math

import (
	"fmt"

	"github.com/cockroachdb/apd/v2"

	"github.com/cosmos/cosmos-sdk/types/errors"
)

// Deprecated: ParseNonNegativeDecimal parses a non-negative decimal or returns an error.
func ParseNonNegativeDecimal(x string) (*apd.Decimal, error) {
	res, _, err := apd.NewFromString(x)
	if err != nil || res.Sign() < 0 {
		return nil, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a non-negative decimal, got %s", x))
	}

	return res, nil
}

// Deprecated: ParsePositiveDecimal parses a positive decimal or returns an error.
func ParsePositiveDecimal(x string) (*apd.Decimal, error) {
	res, _, err := apd.NewFromString(x)
	if err != nil || res.Sign() <= 0 {
		return nil, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a positive decimal, got %s", x))
	}

	return res, nil
}

// Deprecated: DecimalString prints x as a floating point string.
func DecimalString(x *apd.Decimal) string {
	return x.Text('f')
}

// Deprecated: NumDecimalPlaces returns the number of decimal places in x.
func NumDecimalPlaces(x *apd.Decimal) uint32 {
	if x.Exponent >= 0 {
		return 0
	}
	return uint32(-x.Exponent)
}

// Deprecated: ParseNonNegativeDecimal parses a non-negative decimal with a fixed maxDecimalPlaces or returns an error.
func ParseNonNegativeFixedDecimal(x string, maxDecimalPlaces uint32) (*apd.Decimal, error) {
	res, err := ParseNonNegativeDecimal(x)
	if err != nil {
		return nil, err
	}

	err = requireMaxDecimals(res, maxDecimalPlaces)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Deprecated: ParsePositiveFixedDecimal parses a positive decimal with a fixed maxDecimalPlaces or returns an error.
func ParsePositiveFixedDecimal(x string, maxDecimalPlaces uint32) (*apd.Decimal, error) {
	res, err := ParsePositiveDecimal(x)
	if err != nil {
		return nil, err
	}

	err = requireMaxDecimals(res, maxDecimalPlaces)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func requireMaxDecimals(x *apd.Decimal, maxDecimalPlaces uint32) error {
	n := NumDecimalPlaces(x)
	if n > maxDecimalPlaces {
		return errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected no more than %d decimal places in %s, got %d", maxDecimalPlaces, x, n))
	}
	return nil
}

var exactContext = apd.Context{
	Precision:   0,
	MaxExponent: apd.MaxExponent,
	MinExponent: apd.MinExponent,
	Traps:       apd.DefaultTraps | apd.Inexact | apd.Rounded,
}

// Deprecated: Add adds x and y and stores the result in res with arbitrary precision or returns an error.
func Add(res, x, y *apd.Decimal) error {
	_, err := exactContext.Add(res, x, y)
	if err != nil {
		return errors.Wrap(err, "decimal addition error")
	}
	return nil
}

// Deprecated: SafeSub subtracts the value of x from y and stores the result in res with arbitrary precision only
// if the result will be non-negative. An insufficient funds error is returned if the result would be negative.
func SafeSub(res, x, y *apd.Decimal) error {
	_, err := exactContext.Sub(res, x, y)
	if err != nil {
		return errors.Wrap(err, "decimal subtraction error")
	}

	if res.Sign() < 0 {
		return errors.ErrInsufficientFunds
	}

	return nil
}

// SafeSubBalance subtracts the value of y from x and returns the result with arbitrary precision.
// Returns with ErrInsufficientFunds error if the result is negative.
func SafeSubBalance(x Dec, y Dec) (Dec, error) {
	var z Dec
	_, err := exactContext.Sub(&z.dec, &x.dec, &y.dec)
	if err != nil {
		return z, errors.Wrap(err, "decimal subtraction error")
	}

	if z.IsNegative() {
		return z, errors.ErrInsufficientFunds
	}

	return z, nil
}

// SafeAddBalance adds the value of x+y and returns the result with arbitrary precision.
// Returns with ErrInvalidRequest error if either x or y is negative.
func SafeAddBalance(x Dec, y Dec) (Dec, error) {
	var z Dec

	if x.IsNegative() || y.IsNegative() {
		return z, errors.Wrap(
			errors.ErrInvalidRequest,
			fmt.Sprintf("AddBalance() requires two non-negative Dec parameters, but received %s and %s", x, y))
	}

	_, err := exactContext.Add(&z.dec, &x.dec, &y.dec)
	if err != nil {
		return z, errors.Wrap(err, "decimal subtraction error")
	}

	return z, nil
}
