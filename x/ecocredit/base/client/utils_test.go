package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/testutil"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func TestParseMsgCreateBatch(t *testing.T) {
	clientCtx := client.Context{}.WithCodec(&codec.ProtoCodec{})

	invalidJSON := testutil.WriteToNewTempFile(t, `{foo:bar}`).Name()
	duplicateJSON := testutil.WriteToNewTempFile(t, `{"foo":"bar","foo":"baz"}`).Name()
	validJSON := testutil.WriteToNewTempFile(t, `{
		"issuer": "regen1",
		"project_id": "C01-001",
		"issuance": [
			{
				"recipient": "regen2",
				"tradable_amount": "10",
				"retired_amount": "2.5",
				"retirement_jurisdiction": "US-WA"
			}
		],
		"metadata": "metadata",
		"start_date": "2020-01-01T00:00:00Z",
		"end_date": "2021-01-01T00:00:00Z"
	}`).Name()

	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name      string
		file      string
		expErr    bool
		expErrMsg string
		expRes    *types.MsgCreateBatch
	}{
		{
			name:      "empty file path",
			file:      "",
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name:      "invalid json format",
			file:      invalidJSON,
			expErr:    true,
			expErrMsg: "invalid character",
		},
		{
			name:      "duplicate json keys",
			file:      duplicateJSON,
			expErr:    true,
			expErrMsg: "duplicate key",
		},
		{
			name: "valid test",
			file: validJSON,
			expRes: &types.MsgCreateBatch{
				Issuer:    "regen1",
				ProjectId: "C01-001",
				Issuance: []*types.BatchIssuance{
					{
						Recipient:              "regen2",
						TradableAmount:         "10",
						RetiredAmount:          "2.5",
						RetirementJurisdiction: "US-WA",
					},
				},
				Metadata:  "metadata",
				StartDate: &startDate,
				EndDate:   &endDate,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseMsgCreateBatch(clientCtx, tc.file)
			if tc.expErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.expErrMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRes, res)
			}
		})
	}
}

func TestParseSendCredits(t *testing.T) {
	emptyJSON := testutil.WriteToNewTempFile(t, `{}`).Name()
	invalidJSON := testutil.WriteToNewTempFile(t, `{foo:bar}`).Name()
	duplicateJSON := testutil.WriteToNewTempFile(t, `{"foo":"bar","foo":"baz"}`).Name()
	validJSON := testutil.WriteToNewTempFile(t, `[
		{
			"batch_denom": "C01-001-20210101-20210101-001",
			"tradable_amount": "10"
		},
		{
			"batch_denom": "C01-001-20210101-20210101-002",
			"retired_amount": "2.5",
			"retirement_jurisdiction": "US-WA"
		}
	]`).Name()

	testCases := []struct {
		name      string
		file      string
		expErr    bool
		expErrMsg string
		expRes    []*types.MsgSend_SendCredits
	}{
		{
			name:      "empty file path",
			file:      "",
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name:      "empty json object",
			file:      emptyJSON,
			expErr:    true,
			expErrMsg: "cannot unmarshal object",
		},
		{
			name:      "invalid json format",
			file:      invalidJSON,
			expErr:    true,
			expErrMsg: "invalid character",
		},
		{
			name:      "duplicate json keys",
			file:      duplicateJSON,
			expErr:    true,
			expErrMsg: "duplicate key",
		},
		{
			name: "valid test",
			file: validJSON,
			expRes: []*types.MsgSend_SendCredits{
				{
					BatchDenom:     "C01-001-20210101-20210101-001",
					TradableAmount: "10",
				},
				{
					BatchDenom:             "C01-001-20210101-20210101-002",
					RetiredAmount:          "2.5",
					RetirementJurisdiction: "US-WA",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseSendCredits(tc.file)
			if tc.expErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.expErrMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRes, res)
			}
		})
	}
}

func TestParseCredits(t *testing.T) {
	emptyJSON := testutil.WriteToNewTempFile(t, `{}`).Name()
	invalidJSON := testutil.WriteToNewTempFile(t, `{foo:bar}`).Name()
	duplicateJSON := testutil.WriteToNewTempFile(t, `{"foo":"bar","foo":"baz"}`).Name()
	validJSON := testutil.WriteToNewTempFile(t, `[
		{
			"batch_denom": "C01-001-20210101-20210101-001",
			"amount": "10"
		},
		{
			"batch_denom": "C01-001-20210101-20210101-002",
			"amount": "2.5"
		}
	]`).Name()

	testCases := []struct {
		name      string
		file      string
		expErr    bool
		expErrMsg string
		expRes    []*types.Credits
	}{
		{
			name:      "empty file path",
			file:      "",
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name:      "empty json object",
			file:      emptyJSON,
			expErr:    true,
			expErrMsg: "cannot unmarshal object",
		},
		{
			name:      "invalid file format",
			file:      invalidJSON,
			expErr:    true,
			expErrMsg: "invalid character",
		},
		{
			name:      "duplicate json keys",
			file:      duplicateJSON,
			expErr:    true,
			expErrMsg: "duplicate key",
		},
		{
			name: "valid test",
			file: validJSON,
			expRes: []*types.Credits{
				{
					BatchDenom: "C01-001-20210101-20210101-001",
					Amount:     "10",
				},
				{
					BatchDenom: "C01-001-20210101-20210101-002",
					Amount:     "2.5",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseCredits(tc.file)
			if tc.expErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.expErrMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRes, res)
			}
		})
	}
}
