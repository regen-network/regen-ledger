package math

import (
	"testing"

	"github.com/leanovate/gopter/gen"

	"github.com/leanovate/gopter/prop"

	"github.com/cockroachdb/apd/v2"
	"github.com/leanovate/gopter"
	"github.com/stretchr/testify/require"
)

// fuzz test to verify that malicious input cannot cause a panic
func TestParseMustNotPanic(t *testing.T) {
	props := gopter.NewProperties(nil)
	props.Property("ParseNonNegativeDecimal must not panic", prop.ForAll(
		func(x string) bool {
			_, _ = ParseNonNegativeDecimal(x)
			return true
		},
		gen.AnyString(),
	))
	props.Property("ParsePositiveDecimal must not panic", prop.ForAll(
		func(x string) bool {
			_, _ = ParsePositiveDecimal(x)
			return true
		},
		gen.AnyString(),
	))
	props.Property("ParseNonNegativeFixedDecimal must not panic", prop.ForAll(
		func(x string, d uint32) bool {
			_, _ = ParseNonNegativeFixedDecimal(x, d)
			return true
		},
		gen.AnyString(), gen.UInt32(),
	))
	props.Property("ParsePositiveFixedDecimal must not panic", prop.ForAll(
		func(x string, d uint32) bool {
			_, _ = ParsePositiveFixedDecimal(x, d)
			return true
		},
		gen.AnyString(), gen.UInt32(),
	))
	props.TestingRun(t)
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
			x, _, err := apd.NewFromString(tt.x)
			require.NoError(t, err)
			got := NumDecimalPlaces(x)
			require.Equal(t, uint32(tt.want), got)
			x, _ = x.Reduce(x)
			got = NumDecimalPlaces(x)
			require.Equal(t, uint32(tt.reduced), got)
		})
	}
}
