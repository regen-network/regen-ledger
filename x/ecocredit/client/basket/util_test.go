package basketclient

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestParseCredits(t *testing.T) {
	invalidContent := testutil.WriteToNewTempFile(t, `{}`).Name()
	validCredits := testutil.WriteToNewTempFile(t, `[
		{
			"batch_denom": "C01-20210101-20220101-001",
			"amount": "10"
		},
		{
			"batch_denom": "C01-20210101-20220101-001",
			"amount": "10.555"
		}
	]`).Name()

	testCases := []struct {
		name     string
		filePath string
		expErr   bool
		result   []*basket.BasketCredit
		errMsg   string
	}{
		{
			"empty file path",
			"",
			true,
			nil,
			"file path is empty",
		},
		{
			"invalid file content",
			invalidContent,
			true,
			nil,
			"cannot unmarshal object",
		},
		{
			"valid test",
			validCredits,
			false,
			[]*basket.BasketCredit{
				{
					BatchDenom: "C01-20210101-20220101-001",
					Amount:     "10",
				},
				{
					BatchDenom: "C01-20210101-20220101-001",
					Amount:     "10.555",
				},
			},
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseBasketCredits(tc.filePath)
			if tc.expErr {
				require.Error(t, err, err.Error())
				require.Contains(t, err.Error(), tc.errMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, res, tc.result)
			}
		})
	}
}
