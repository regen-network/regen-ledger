package client

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"time"
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

func TestJSONDupliteKeys(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		expErr     bool
		errMessage string
	}{
		{
			"invalid json",
			`{abcd}`,
			true,
			"invalid character",
		},
		{
			"valid json simple",
			`{"class_id": "C01", "end_date": "2022-09-08T00:00:00Z", "project_location": "AB-CDE FG1 345"}`,
			false,
			"",
		},
		{
			"valid json nested",
			`{
				"class_id": "C01",
				"issuance": [
					{
						"recipient": "regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38",
						"tradable_amount": "1000",
						"retired_amount": "5",
						"retirement_location": "ST-UVW XY Z12"
					}
				],
				"metadata": "Y2FyYm9uCg==",
				"start_date": "2021-09-08T00:00:00Z",
				"end_date": "2022-09-08T00:00:00Z",
				"project_location": "AB-CDE FG1 345"
			}`,
			false,
			"",
		},
		{
			"invalid json duplicate keys",
			`{"class_id": "C01", "end_date": "2022-09-08T00:00:00Z", "class_id": "C01"}`,
			true,
			"duplicate key class_id",
		},
		{
			"invalid nested json duplicate keys",
			`{"class_id": "C01", "end_date": "2022-09-08T00:00:00Z", "issuance": [
				{
					"recipient": "regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38",
					"recipient": "regen1r9pl9gvr56kmclgkpjg3ynh4rm5am66f2a6y38",
					"tradable_amount": "1000",
					"retired_amount": "5",
					"retirement_location": "ST-UVW XY Z12"
				}
			]}`,
			true,
			"duplicate key recipient",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkDuplicateKey(json.NewDecoder(strings.NewReader(tc.input)), nil)
			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errMessage)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

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
			err := parseAndSetDate(&tm, tc.date, tc.date)
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
