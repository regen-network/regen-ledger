package math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDec(t *testing.T) {
	zero := Dec{}
	one := NewDecFromInt64(1)
	two := NewDecFromInt64(2)
	three := NewDecFromInt64(3)
	four := NewDecFromInt64(4)
	five := NewDecFromInt64(5)
	minusOne := NewDecFromInt64(-1)

	onePointOneFive, err := NewDecFromString("1.15")
	require.NoError(t, err)
	twoPointThreeFour, err := NewDecFromString("2.34")
	require.NoError(t, err)
	threePointFourNine, err := NewDecFromString("3.49")
	require.NoError(t, err)
	onePointFourNine, err := NewDecFromString("1.49")
	require.NoError(t, err)
	minusFivePointZero, err := NewDecFromString("-5.0")
	require.NoError(t, err)

	res, err := two.Add(zero)
	require.NoError(t, err)
	require.True(t, res.IsEqual(two))

	res, err = five.Sub(two)
	require.NoError(t, err)
	require.True(t, res.IsEqual(three))

	res, err = SafeSubBalance(five, two)
	require.NoError(t, err)
	require.True(t, res.IsEqual(three))

	res, err = SafeSubBalance(two, five)
	require.Error(t, err, "Expected insufficient funds error")

	res, err = SafeAddBalance(three, two)
	require.NoError(t, err)
	require.True(t, res.IsEqual(five))

	res, err = SafeAddBalance(minusFivePointZero, five)
	require.Error(t, err, "Expected ErrInvalidRequest")

	res, err = four.Quo(two)
	require.NoError(t, err)
	require.True(t, res.IsEqual(two))

	res, err = five.QuoInteger(two)
	require.NoError(t, err)
	require.True(t, res.IsEqual(two))

	res, err = five.Rem(two)
	require.NoError(t, err)
	require.True(t, res.IsEqual(one))

	x, err := four.Int64()
	require.NoError(t, err)
	require.Equal(t, int64(4), x)

	require.Equal(t, "5", five.String())

	res, err = onePointOneFive.Add(twoPointThreeFour)
	require.NoError(t, err)
	require.True(t, res.IsEqual(threePointFourNine))

	res, err = threePointFourNine.Sub(two)
	require.NoError(t, err)
	require.True(t, res.IsEqual(onePointFourNine))

	res, err = minusOne.Sub(four)
	require.NoError(t, err)
	require.True(t, res.IsEqual(minusFivePointZero))

	require.True(t, zero.IsZero())
	require.False(t, zero.IsPositive())
	require.False(t, zero.IsNegative())

	require.False(t, one.IsZero())
	require.True(t, one.IsPositive())
	require.False(t, one.IsNegative())

	require.False(t, minusOne.IsZero())
	require.False(t, minusOne.IsPositive())
	require.True(t, minusOne.IsNegative())


	/*

	TODO: Add Property Tests
	========================

	// Arithmetic tests

    Mul(QuoInteger(a,b),b) + Rem(a,b) == a

	Quo(Mul(a,b),b) == a

	Quo(Mul(a,b),a) == b

	Add(b,Sub(a,b)) == a

	Sub(Add(a,b),b) == a

	// Comparison tests
	Cmp(a,b) IFF -1*Cmp(b,a) for varying precision b & a values
	IsEqual(a,b) IFF IsEqual(b,a) for varying precision b & a values

	*/
}
