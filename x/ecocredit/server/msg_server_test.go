package server

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
			name:      "can validate valid country code",
			location:  "AF",
			expectErr: false,
		},
		{
			name:      "can validate valid country code 2",
			location:  "BF",
			expectErr: false,
		},
		{
			name:      "can't validate invalid country code",
			location:  "ZZ",
			expectErr: true,
		},
		{
			name:      "can't validate invalid country code 2",
			location:  "Z!adflksdfZ",
			expectErr: true,
		},
		{
			name:      "can validate valid region code",
			location:  "AF-BDS",
			expectErr: false,
		},
		{
			name:      "can validate valid region code 2",
			location:  "BF-BAL",
			expectErr: false,
		},
		{
			name:      "can't validate invalid region code",
			location:  "BF-ZZ",
			expectErr: true,
		},
		{
			name:      "can validate valid postal code",
			location:  "BF-BAL-1",
			expectErr: false,
		},
		{
			name:      "can validate valid postal code 2",
			location:  "BF-BAL-0123456789",
			expectErr: false,
		},
		{
			name:      "can't validate invalid postal code",
			location:  "BF-BAL-01234567890",
			expectErr: true,
		},
	}

	for _, spec := range specs {
		t.Run(spec.name, func (t *testing.T) {
			if err := validateLocation(spec.location); spec.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})

	}

}
