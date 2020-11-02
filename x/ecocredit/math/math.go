package math

import (
	"fmt"
	"github.com/cockroachdb/apd/v2"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func MustParseNonNegativeDecimal(x string) (*apd.Decimal, error) {
	res, _, err := apd.NewFromString(x)
	if err != nil {
		return nil, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a non-negative decimal, got %s", x))
	}

	if res.Sign() < 0 {
		return nil, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a non-negative decimal, got %s", x))
	}

	return res, nil
}

func MustParsePositiveDecimal(x string) (*apd.Decimal, error) {
	res, _, err := apd.NewFromString(x)
	if err != nil {
		return nil, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a positive decimal, got %s", x))
	}

	if res.Sign() <= 0 {
		return nil, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a positive decimal, got %s", x))
	}

	return res, nil
}

func DecString(x *apd.Decimal) string {
	return x.Text('f')
}

func NumDecimalPlaces(x *apd.Decimal) uint32 {
	if x.Exponent >= 0 {
		return 0
	}
	return uint32(-x.Exponent)
}

func MustParseNonNegativeFixedWidthDecimal(x string, maxDecimalPlaces uint32) (*apd.Decimal, error) {
	res, err := MustParseNonNegativeDecimal(x)
	if err != nil {
		return nil, err
	}

	err = requireMaxDecimals(res, maxDecimalPlaces)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func MustParsePositiveFixedWidthDecimal(x string, maxDecimalPlaces uint32) (*apd.Decimal, error) {
	res, err := MustParsePositiveDecimal(x)
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
