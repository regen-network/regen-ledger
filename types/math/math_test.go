package math

import (
	"testing"

	"github.com/cockroachdb/apd/v2"
	"github.com/stretchr/testify/require"
)

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
