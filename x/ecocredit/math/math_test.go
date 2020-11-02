package math

import (
	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

const (
	X = 10
	Y = 1000000010
)

func TestDec(t *testing.T) {
	t.Logf("sdk.Dec A: %s", sdk.NewDec(X).QuoInt64(Y).MulInt64(Y))
	t.Logf("sdk.Dec B: %s", sdk.NewDec(X).MulInt64(Y).QuoInt64(Y))

	rx := big.NewRat(X, 1)
	ry := big.NewRat(Y, 1)
	rz := rx.Quo(rx, ry)
	rz = rz.Mul(rz, ry)
	t.Logf("sdk.Rat A: %s", rz.FloatString(18))
	rz = rx.Mul(rx, ry)
	rz = rz.Quo(rz, ry)
	t.Logf("sdk.Rat B: %s", rz.FloatString(18))

	dx, _, err := apd.NewFromString("10")
	require.NoError(t, err)
	dy, _, err := apd.NewFromString("1000000010")
	require.NoError(t, err)
	var dz apd.Decimal
	_, err = StrictDecimal128Context.Quo(&dz, dx, dy)
	require.NoError(t, err)
	_, err = StrictDecimal128Context.Mul(&dz, &dz, dy)
	require.NoError(t, err)
	t.Logf("dec128 A: %s", &dz)

	_, err = StrictDecimal128Context.Mul(&dz, dx, dy)
	require.NoError(t, err)
	_, err = StrictDecimal128Context.Quo(&dz, &dz, dy)
	require.NoError(t, err)
	t.Logf("dec128 B: %s", &dz)

	dec64Ctx := apd.Context{
		Precision:   16,
		MaxExponent: 384,
		MinExponent: -383,
		Traps:       apd.DefaultTraps,
		Rounding:    apd.RoundHalfEven,
	}

	_, err = dec64Ctx.Quo(&dz, dx, dy)
	require.NoError(t, err)
	_, err = dec64Ctx.Mul(&dz, &dz, dy)
	require.NoError(t, err)
	t.Logf("dec64 A: %s", &dz)

	_, err = dec64Ctx.Mul(&dz, dx, dy)
	require.NoError(t, err)
	_, err = dec64Ctx.Quo(&dz, &dz, dy)
	require.NoError(t, err)
	t.Logf("dec64 B: %s", &dz)

	dec32Ctx := apd.Context{
		Precision:   7,
		MaxExponent: 96,
		MinExponent: -95,
		Traps:       apd.DefaultTraps,
		Rounding:    apd.RoundHalfEven,
	}

	_, err = dec32Ctx.Quo(&dz, dx, dy)
	require.NoError(t, err)
	_, err = dec32Ctx.Mul(&dz, &dz, dy)
	require.NoError(t, err)
	t.Logf("dec32 A: %s", &dz)

	_, err = dec32Ctx.Mul(&dz, dx, dy)
	require.NoError(t, err)
	_, err = dec32Ctx.Quo(&dz, &dz, dy)
	require.NoError(t, err)
	t.Logf("dec32 B: %s", &dz)

	doublex := float64(X)
	doubley := float64(Y)
	t.Logf("float64 A: %f", (doublex/doubley)*doubley)
	t.Logf("float64 B: %f", (doublex*doubley)/doubley)

	floatx := float32(X)
	floaty := float32(Y)
	t.Logf("float32 A: %f", (floatx/floaty)*floaty)
	t.Logf("float32 B: %f", (floatx*floaty)/floaty)
}
