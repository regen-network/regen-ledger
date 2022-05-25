package basket

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatBasketDenom(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		tname        string
		abbrev       string
		exponent     uint32
		denom        string
		displayDenom string
		err          bool
	}{
		{"wrong exponent",
			"X", 5, "", "", true},
		{"exponent-0",
			"X", 0, "eco.X.foo", "eco.X.foo", false},
		{"exponent-1`",
			"X", 1, "eco.dX.foo", "eco.X.foo", false},
		{"exponent-2",
			"X", 2, "eco.cX.foo", "eco.X.foo", false},
		{"exponent-6",
			"X", 6, "eco.uX.foo", "eco.X.foo", false},
	}
	require := require.New(t)
	for _, tc := range tcs {
		t.Run(tc.tname, func(t *testing.T) {
			t.Parallel()
			d, displayD, err := FormatBasketDenom("foo", tc.abbrev, tc.exponent)
			if tc.err {
				require.Error(err, tc.tname)
			} else {
				require.NoError(err, tc.tname)
				require.Equal(tc.denom, d, tc.tname)
				require.Equal(tc.displayDenom, displayD, tc.tname)
			}
		})
	}
}
