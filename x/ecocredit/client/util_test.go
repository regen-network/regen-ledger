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
			name:           "can't parse credits with missing colon",
			creditsListStr: "10ABC/123",
			expectErr:      true,
		},
		{
			name:           "can't parse credits with malformed batch denom",
			creditsListStr: "10:ABC123",
			expectErr:      true,
		},
		{
			name:           "can't parse credits with malformed amount",
			creditsListStr: "10!:ABC/123",
			expectErr:      true,
		},
		{
			name:           "can parse single credits with simple decimal",
			creditsListStr: "10:ABC/123",
			expectErr:      false,
			expCreditsList: []credits{
				credits{
					batchDenom: "ABC/123",
					amount:     "10",
				},
			},
		},
		{
			name:           "can parse single credits with multiple places",
			creditsListStr: "10.0000001:ABC/123",
			expectErr:      false,
			expCreditsList: []credits{
				credits{
					batchDenom: "ABC/123",
					amount:     "10.0000001",
				},
			},
		},
		{
			name:           "can parse single credits with no digit before decimal point",
			creditsListStr: ".0000001:ABC/123",
			expectErr:      false,
			expCreditsList: []credits{
				credits{
					batchDenom: "ABC/123",
					amount:     ".0000001",
				},
			},
		},
		{
			name:           "can parse multiple credits",
			creditsListStr: ".0000001:ABC/123,10:345/XYZ,10000.0001:IJK/LMN",
			expectErr:      false,
			expCreditsList: []credits{
				credits{
					batchDenom: "ABC/123",
					amount:     ".0000001",
				},
				credits{
					batchDenom: "345/XYZ",
					amount:     "10",
				},
				credits{
					batchDenom: "IJK/LMN",
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
