package ecocredit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateJurisdiction(t *testing.T) {
	specs := []struct {
		name         string
		jurisdiction string
		expectErr    bool
	}{
		{
			name:         "valid country code",
			jurisdiction: "AF",
			expectErr:    false,
		},
		{
			name:         "valid country code 2",
			jurisdiction: "BF",
			expectErr:    false,
		},
		{
			name:         "invalid country code",
			jurisdiction: "ZZZ",
			expectErr:    true,
		},
		{
			name:         "invalid country code 2",
			jurisdiction: "Z!adflksdfZ",
			expectErr:    true,
		},
		{
			name:         "valid region code",
			jurisdiction: "AF-BDS",
			expectErr:    false,
		},
		{
			name:         "valid region code 2",
			jurisdiction: "BF-B12",
			expectErr:    false,
		},
		{
			name:         "invalid region code",
			jurisdiction: "BF-ZZZZ",
			expectErr:    true,
		},
		{
			name:         "invalid region code 2",
			jurisdiction: "BF-AB!",
			expectErr:    true,
		},
		{
			name:         "valid postal code",
			jurisdiction: "BF-BAL 1",
			expectErr:    false,
		},
		{
			name:         "valid postal code 2",
			jurisdiction: "BF-B12 0123456789",
			expectErr:    false,
		},
		{
			name:         "invalid postal code",
			jurisdiction: "BF-BAL 0123456789012345678901234567890123456789012345678901234567890123456789",
			expectErr:    true,
		},
	}

	for _, spec := range specs {
		t.Run(spec.name, func(t *testing.T) {
			if err := ValidateJurisdiction(spec.jurisdiction); spec.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})

	}

}
