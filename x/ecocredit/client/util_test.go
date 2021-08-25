package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCredits(t *testing.T) {
	specs := []struct {
		name           string
		creditsListStr string
		expectErr      bool
		expCreditsList []credits
	}{
		{
			name:           "missing space",
			creditsListStr: "10C01-20200101-20210101-001",
			expectErr:      true,
		},
		{
			name:           "malformed batch denom",
			creditsListStr: "10 ABC123",
			expectErr:      true,
		},
		{
			name:           "malformed amount",
			creditsListStr: "10! C01-20200101-20210101-001",
			expectErr:      true,
		},
		{
			name:           "single credits with simple decimal",
			creditsListStr: "10 C01-20200101-20210101-001",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C01-20200101-20210101-001",
					amount:     "10",
				},
			},
		},
		{
			name:           "single credits with multiple places",
			creditsListStr: "10.0000001 C01-20200101-20210101-001",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C01-20200101-20210101-001",
					amount:     "10.0000001",
				},
			},
		},
		{
			name:           "single credits with no digit before decimal point",
			creditsListStr: ".0000001 C01-20200101-20210101-001",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C01-20200101-20210101-001",
					amount:     ".0000001",
				},
			},
		},
		{
			name:           "single credits overflowing padding",
			creditsListStr: ".0000001 C123-20200101-20210101-1234",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C123-20200101-20210101-1234",
					amount:     ".0000001",
				},
			},
		},
		{
			name:           "multiple credits",
			creditsListStr: ".0000001 C01-20200101-20210101-001,10 C01-20200101-20210101-002, 10000.0001 C01-20200101-20210101-003",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C01-20200101-20210101-001",
					amount:     ".0000001",
				},
				{
					batchDenom: "C01-20200101-20210101-002",
					amount:     "10",
				},
				{
					batchDenom: "C01-20200101-20210101-003",
					amount:     "10000.0001",
				},
			},
		},
	}

	for _, spec := range specs {
		t.Run(spec.name, func(t *testing.T) {
			creditsList, err := parseCreditsList(spec.creditsListStr)
			if spec.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, spec.expCreditsList, creditsList)
			}

			// Also check that the tests pass successfully when wrapping with CancelCredits
			_, err = parseCancelCreditsList(spec.creditsListStr)
			if spec.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
