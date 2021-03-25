package math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMath(t *testing.T) {
	zero := Dec{}
	one := NewDecFromInt64(1)
	two := NewDecFromInt64(2)
	three := NewDecFromInt64(3)
	four := NewDecFromInt64(4)
	five := NewDecFromInt64(5)
	minusOne := NewDecFromInt64(-1)

	res, err := two.Add(zero)
	require.NoError(t, err)
	require.True(t, res.IsEqual(two))

	res, err = five.Sub(two)
	require.NoError(t, err)
	require.True(t, res.IsEqual(three))

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

	require.True(t, zero.IsZero())
	require.False(t, zero.IsPositive())
	require.False(t, zero.IsNegative())

	require.False(t, one.IsZero())
	require.True(t, one.IsPositive())
	require.False(t, one.IsNegative())

	require.False(t, minusOne.IsZero())
	require.False(t, minusOne.IsPositive())
	require.True(t, minusOne.IsNegative())
}

func Test_CountDecPlaces(t *testing.T) {
	tests := []struct {
		x       string
		want    int32
		reduced int32
	}{
		{"0", 0, 0},
		{"1.0", 1, 0},
		{"1.0000", 4, 0},
		{"1.23", 2, 2},
		{"100", 0, 0},
		{"0.0003", 4, 4},
	}
	for _, tt := range tests {
		t.Run(tt.x, func(t *testing.T) {
			x, err := NewDecFromString(tt.x)
			require.NoError(t, err)
			got := x.NumDecimalPlaces()
			require.Equal(t, uint32(tt.want), got)
			x, _ = x.Reduce()
			got = x.NumDecimalPlaces()
			require.Equal(t, uint32(tt.reduced), got)
		})
	}
}
