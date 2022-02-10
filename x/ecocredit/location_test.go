package ecocredit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateLocation(t *testing.T) {
	specs := []struct {
		name      string
		location  string
		expectErr bool
	}{
		{
			name:      "valid country code",
			location:  "AF",
			expectErr: false,
		},
		{
			name:      "valid country code 2",
			location:  "BF",
			expectErr: false,
		},
		{
			name:      "invalid country code",
			location:  "ZZZ",
			expectErr: true,
		},
		{
			name:      "invalid country code 2",
			location:  "Z!adflksdfZ",
			expectErr: true,
		},
		{
			name:      "valid region code",
			location:  "AF-BDS",
			expectErr: false,
		},
		{
			name:      "valid region code 2",
			location:  "BF-B12",
			expectErr: false,
		},
		{
			name:      "invalid region code",
			location:  "BF-ZZZZ",
			expectErr: true,
		},
		{
			name:      "invalid region code 2",
			location:  "BF-AB!",
			expectErr: true,
		},
		{
			name:      "valid postal code",
			location:  "BF-BAL 1",
			expectErr: false,
		},
		{
			name:      "valid postal code 2",
			location:  "BF-B12 0123456789",
			expectErr: false,
		},
		{
			name:      "invalid postal code",
			location:  "BF-BAL 0123456789012345678901234567890123456789012345678901234567890123456789",
			expectErr: true,
		},
	}

	for _, spec := range specs {
		t.Run(spec.name, func(t *testing.T) {
			if err := ValidateLocation(spec.location); spec.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})

	}

}
