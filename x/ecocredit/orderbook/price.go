package orderbook

import (
	"fmt"
	"math"

	"github.com/cockroachdb/apd/v3"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IntPriceToUInt32(price sdk.Int, precisionModifier uint32) (uint32, error) {
	p, _, err := apd.NewFromString(price.String())
	if err != nil {
		return 0, err
	}

	div := apd.New(10, int32(precisionModifier))

	var res apd.Decimal
	_, err = toUint32Ctx.QuoInteger(&res, p, div)
	if err != nil {
		return 0, err
	}

	i64, err := res.Int64()
	if err != nil {
		return 0, err
	}

	if i64 > math.MaxUint32 || i64 < 0 {
		return 0, fmt.Errorf("%d out of bounds for converting to a uint32", i64)
	}

	return uint32(i64), nil
}

var toUint32Ctx = apd.Context{
	Precision:   18,
	MaxExponent: apd.MaxExponent,
	MinExponent: apd.MinExponent,
	Traps:       apd.DefaultTraps,
	Rounding:    apd.RoundHalfUp,
}
