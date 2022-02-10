package math

import (
	"fmt"

	"github.com/cockroachdb/apd/v2"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// Dec is a wrapper struct around apd.Decimal that does no mutation of apd.Decimal's when performing
// arithmetic, instead creating a new apd.Decimal for every operation ensuring usage is safe.
//
// Using apd.Decimal directly can be unsafe because apd operations mutate the underlying Decimal,
// but when copying the big.Int structure can be shared between Decimal instances causing corruption.
// This was originally discovered in regen0-network/mainnet#15.
type Dec struct {
	dec apd.Decimal
}

const mathCodespace = "math"

var ErrInvalidDecString = errors.Register(mathCodespace, 1, "invalid decimal string")

// In cosmos-sdk#7773, decimal128 (with 34 digits of precision) was suggested for performing
// Quo/Mult arithmetic generically across the SDK. Even though the SDK
// has yet to support a GDA with decimal128 (34 digits), we choose to utilize it here.
// https://github.com/cosmos/cosmos-sdk/issues/7773#issuecomment-725006142
var dec128Context = apd.Context{
	Precision:   34,
	MaxExponent: apd.MaxExponent,
	MinExponent: apd.MinExponent,
	Traps:       apd.DefaultTraps,
}

func NewDecFromString(s string) (Dec, error) {
	d, _, err := apd.NewFromString(s)
	if err != nil {
		return Dec{}, ErrInvalidDecString.Wrap(err.Error())
	}
	return Dec{*d}, nil
}

func NewNonNegativeDecFromString(s string) (Dec, error) {
	d, err := NewDecFromString(s)
	if err != nil {
		return Dec{}, ErrInvalidDecString.Wrap(err.Error())
	}
	if d.IsNegative() {
		return Dec{}, ErrInvalidDecString.Wrapf("expected a non-negative decimal, got %s", s)
	}
	return d, nil
}

func NewNonNegativeFixedDecFromString(s string, max uint32) (Dec, error) {
	d, err := NewNonNegativeDecFromString(s)
	if err != nil {
		return Dec{}, err
	}
	if d.NumDecimalPlaces() > max {
		return Dec{}, fmt.Errorf("%s exceeds maximum decimal places: %d", s, max)
	}
	return d, nil
}

func NewPositiveDecFromString(s string) (Dec, error) {
	d, err := NewDecFromString(s)
	if err != nil {
		return Dec{}, ErrInvalidDecString.Wrap(err.Error())
	}
	if !d.IsPositive() {
		return Dec{}, ErrInvalidDecString.Wrapf("expected a positive decimal, got %s", s)
	}
	return d, nil
}

func NewPositiveFixedDecFromString(s string, max uint32) (Dec, error) {
	d, err := NewPositiveDecFromString(s)
	if err != nil {
		return Dec{}, err
	}
	if d.NumDecimalPlaces() > max {
		return Dec{}, fmt.Errorf("%s exceeds maximum decimal places: %d", s, max)
	}
	return d, nil
}

func NewDecFromInt64(x int64) Dec {
	var res Dec
	res.dec.SetInt64(x)
	return res
}

func NewFinite(coeff int64, exp int32) Dec {
	var res Dec
	res.dec.SetFinite(coeff, exp)
	return res
}

// Add returns a new Dec with value `x+y` without mutating any argument and error if
// there is an overflow.
func (x Dec) Add(y Dec) (Dec, error) {
	var z Dec
	_, err := apd.BaseContext.Add(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal addition error")
}

// Sub returns a new Dec with value `x-y` without mutating any argument and error if
// there is an overflow.
func (x Dec) Sub(y Dec) (Dec, error) {
	var z Dec
	_, err := apd.BaseContext.Sub(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal subtraction error")
}

// Quo returns a new Dec with value `x/y` (formatted as decimal128, 34 digit precision) without mutating any
// argument and error if there is an overflow.
func (x Dec) Quo(y Dec) (Dec, error) {
	var z Dec
	_, err := dec128Context.Quo(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal quotient error")
}

type RoundingMode string

const (
	// RoundDown rounds toward 0; truncate.
	RoundDown RoundingMode = apd.RoundDown
	// RoundHalfUp rounds up if the digits are >= 0.5.
	RoundHalfUp RoundingMode = apd.RoundHalfUp
	// RoundHalfEven rounds up if the digits are > 0.5. If the digits are equal
	// to 0.5, it rounds up if the previous digit is odd, always producing an
	// even digit.
	RoundHalfEven RoundingMode = apd.RoundHalfEven
	// RoundCeiling towards +Inf: rounds up if digits are > 0 and the number
	// is positive.
	RoundCeiling RoundingMode = apd.RoundCeiling
	// RoundFloor towards -Inf: rounds up if digits are > 0 and the number
	// is negative.
	RoundFloor RoundingMode = apd.RoundFloor
	// RoundHalfDown rounds up if the digits are > 0.5.
	RoundHalfDown RoundingMode = apd.RoundHalfDown
	// RoundUp rounds away from 0.
	RoundUp RoundingMode = apd.RoundUp
	// Round05Up rounds zero or five away from 0; same as round-up, except that
	// rounding up only occurs if the digit to be rounded up is 0 or 5.
	Round05Up RoundingMode = apd.Round05Up
)

func (x Dec) QuoPrec(y Dec, precision uint32, mode RoundingMode) (Dec, error) {
	ctx := apd.Context{
		Precision:   precision,
		MaxExponent: apd.MaxExponent,
		MinExponent: apd.MinExponent,
		Traps:       apd.DefaultTraps,
		Rounding:    string(mode),
	}
	var z Dec
	_, err := ctx.Quo(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal quotient error")
}

// QuoInteger returns a new integral Dec with value `x/y` (formatted as decimal128, with 34 digit precision)
// without mutating any argument and error if there is an overflow.
func (x Dec) QuoInteger(y Dec) (Dec, error) {
	var z Dec
	_, err := dec128Context.QuoInteger(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal quotient error")
}

// Rem returns the integral remainder from `x/y` (formatted as decimal128, with 34 digit precision) without
// mutating any argument and error if the integer part of x/y cannot fit in 34 digit precision
func (x Dec) Rem(y Dec) (Dec, error) {
	var z Dec
	_, err := dec128Context.Rem(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal remainder error")
}

// Mul returns a new Dec with value `x*y` (formatted as decimal128, with 34 digit precision) without
// mutating any argument and error if there is an overflow.
func (x Dec) Mul(y Dec) (Dec, error) {
	var z Dec
	_, err := dec128Context.Mul(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal multiplication error")
}

func (x Dec) MulPrec(y Dec, precision uint32, mode RoundingMode) (Dec, error) {
	ctx := apd.Context{
		Precision:   precision,
		MaxExponent: apd.MaxExponent,
		MinExponent: apd.MinExponent,
		Traps:       apd.DefaultTraps,
		Rounding:    string(mode),
	}
	var z Dec
	_, err := ctx.Mul(&z.dec, &x.dec, &y.dec)
	return z, errors.Wrap(err, "decimal quotient error")
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
