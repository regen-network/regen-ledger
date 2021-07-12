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
			name:           "can't parse credits with missing prefix",
			creditsListStr: "10ABC:123",
			expectErr:      true,
		},
		{
			name:           "can't parse credits with malformed batch denom",
			creditsListStr: "10eco:ABC123",
			expectErr:      true,
		},
		{
			name:           "can't parse credits with malformed amount",
			creditsListStr: "10!eco:ABC:123",
			expectErr:      true,
		},
		{
			name:           "can parse single credits with simple decimal",
			creditsListStr: "10eco:ABC:123",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "eco:ABC:123",
					amount:     "10",
				},
			},
		},
		{
			name:           "can parse single credits with multiple places",
			creditsListStr: "10.0000001eco:ABC:123",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "eco:ABC:123",
					amount:     "10.0000001",
				},
			},
		},
		{
			name:           "can parse single credits with no digit before decimal point",
			creditsListStr: ".0000001eco:ABC:123",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "eco:ABC:123",
					amount:     ".0000001",
				},
			},
		},
		{
			name:           "can parse multiple credits",
			creditsListStr: ".0000001eco:ABC:123,10eco:345:XYZ,10000.0001eco:IJK:LMN",
			expectErr:      false,
			expCreditsList: []credits{
				{
					batchDenom: "eco:ABC:123",
					amount:     ".0000001",
				},
				{
					batchDenom: "eco:345:XYZ",
					amount:     "10",
				},
				{
					batchDenom: "eco:IJK:LMN",
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
