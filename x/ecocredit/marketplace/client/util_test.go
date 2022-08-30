package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

func TestParseSellOrders(t *testing.T) {
	emptyJSON := testutil.WriteToNewTempFile(t, `{}`).Name()
	invalidJSON := testutil.WriteToNewTempFile(t, `{foo:bar}`).Name()
	duplicateJSON := testutil.WriteToNewTempFile(t, `{"foo":"bar","foo":"baz"}`).Name()
	validJSON := testutil.WriteToNewTempFile(t, `[
		{
			"batch_denom": "C01-001-20200101-20210101-001",
			"quantity": "10",
			"ask_price": {
				"denom": "regen",
				"amount": "10"
			},
			"disable_auto_retire": true
		},
		{
			"batch_denom": "C01-001-20200101-20210101-002",
			"quantity": "20",
			"ask_price": {
				"denom": "regen",
				"amount": "20"
			},
			"expiration": "2022-01-01T00:00:00Z"
		}
	]`).Name()

	expiration := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name      string
		file      string
		expErr    bool
		expErrMsg string
		expRes    []*types.MsgSell_Order
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
			name: "valid",
			file: validJSON,
			expRes: []*types.MsgSell_Order{
				{
					BatchDenom:        "C01-001-20200101-20210101-001",
					Quantity:          "10",
					AskPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(10)},
					DisableAutoRetire: true,
				},
				{
					BatchDenom: "C01-001-20200101-20210101-002",
					Quantity:   "20",
					AskPrice:   &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(20)},
					Expiration: &expiration,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseSellOrders(tc.file)
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

func TestParseSellUpdates(t *testing.T) {
	emptyJSON := testutil.WriteToNewTempFile(t, `{}`).Name()
	invalidJSON := testutil.WriteToNewTempFile(t, `{foo:bar}`).Name()
	duplicateJSON := testutil.WriteToNewTempFile(t, `{"foo":"bar","foo":"baz"}`).Name()
	validJSON := testutil.WriteToNewTempFile(t, `[
		{
			"sell_order_id": 1,
			"new_quantity": "10",
			"new_ask_price": {
				"denom": "regen",
				"amount": "10"
			},
			"disable_auto_retire": true
		},
		{
			"sell_order_id": 2,
			"new_quantity": "20",
			"new_ask_price": {
				"denom": "regen",
				"amount": "20"
			},
			"new_expiration": "2022-01-01T00:00:00Z"
		}
	]`).Name()

	expiration := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name      string
		file      string
		expErr    bool
		expErrMsg string
		expRes    []*types.MsgUpdateSellOrders_Update
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
			name: "valid",
			file: validJSON,
			expRes: []*types.MsgUpdateSellOrders_Update{
				{
					SellOrderId:       1,
					NewQuantity:       "10",
					NewAskPrice:       &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(10)},
					DisableAutoRetire: true,
				},
				{
					SellOrderId:   2,
					NewQuantity:   "20",
					NewAskPrice:   &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(20)},
					NewExpiration: &expiration,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseSellUpdates(tc.file)
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

func TestParseBuyOrders(t *testing.T) {
	emptyJSON := testutil.WriteToNewTempFile(t, `{}`).Name()
	invalidJSON := testutil.WriteToNewTempFile(t, `{foo:bar}`).Name()
	duplicateJSON := testutil.WriteToNewTempFile(t, `{"foo":"bar","foo":"baz"}`).Name()
	validJSON := testutil.WriteToNewTempFile(t, `[
		{
			"sell_order_id": 1,
			"quantity": "10",
			"bid_price": {
				"denom": "regen",
				"amount": "10"
			},
			"disable_auto_retire": true
		},
		{
			"sell_order_id": 2,
			"quantity": "20",
			"bid_price": {
				"denom": "regen",
				"amount": "20"
			},
			"retirement_jurisdiction": "US-WA"
		}
	]`).Name()

	testCases := []struct {
		name      string
		file      string
		expErr    bool
		expErrMsg string
		expRes    []*types.MsgBuyDirect_Order
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
			name: "valid",
			file: validJSON,
			expRes: []*types.MsgBuyDirect_Order{
				{
					SellOrderId:       1,
					Quantity:          "10",
					BidPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(10)},
					DisableAutoRetire: true,
				},
				{
					SellOrderId:            2,
					Quantity:               "20",
					BidPrice:               &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(20)},
					RetirementJurisdiction: "US-WA",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseBuyOrders(tc.file)
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
