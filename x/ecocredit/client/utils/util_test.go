package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseAndSetDate(t *testing.T) {
	tcs := []struct {
		name   string
		date   string
		hasErr bool
	}{
		{"good", "2022-01-20", false},
		{"bad", "01-2021-20", true},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			var tm *time.Time
			err := ParseAndSetDate(&tm, tc.date, tc.date)
			if tc.hasErr {
				require.Error(err)
				require.Nil(tm)
			} else {
				require.NoError(err)
				require.NotNil(tm)
			}
		})
	}
}
