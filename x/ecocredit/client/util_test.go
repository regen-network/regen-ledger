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
			creditsListStr: "10C001-20200101-20210101-01",
			expectErr:      true,
		},
		{
			name:           "malformed batch denom",
			creditsListStr: "10 ABC123",
			expectErr:      true,
		},
		{
			name:           "malformed amount",
			creditsListStr: "10! C001-20200101-20210101-01",
			expectErr:      true,
		},
		{
			name:           "single credits with simple decimal",
			creditsListStr: "10 C001-20200101-20210101-01",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C001-20200101-20210101-01",
					amount:     "10",
				},
			},
		},
		{
			name:           "single credits with multiple places",
			creditsListStr: "10.0000001 C001-20200101-20210101-01",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C001-20200101-20210101-01",
					amount:     "10.0000001",
				},
			},
		},
		{
			name:           "single credits with no digit before decimal point",
			creditsListStr: ".0000001 C001-20200101-20210101-01",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C001-20200101-20210101-01",
					amount:     ".0000001",
				},
			},
		},
		{
			name:           "multiple credits",
			creditsListStr: ".0000001 C001-20200101-20210101-01,10 C001-20200101-20210101-02, 10000.0001 C001-20200101-20210101-03",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "C001-20200101-20210101-01",
					amount:     ".0000001",
				},
				{
					batchDenom: "C001-20200101-20210101-02",
					amount:     "10",
				},
				{
					batchDenom: "C001-20200101-20210101-03",
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
