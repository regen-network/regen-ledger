package testsuite

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
)

func (s *IntegrationTestSuite) TestQueryClasses() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	testCases := []struct {
		name            string
		args            []string
		expectErr       bool
		expectedErrMsg  string
		expectedClasses []string
	}{
		{
			name:           "too many args",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		{
			name:            "no pagination flags",
			args:            []string{},
			expectErr:       false,
			expectedClasses: []string{"C01", "C02", "C03", "C04"},
		},
		{
			name: "limit 2",
			args: []string{
				fmt.Sprintf("--%s=2", flags.FlagLimit),
			},
			expectErr:       false,
			expectedClasses: []string{"C01", "C02"},
		},
		{
			name: "limit 2, offset 2",
			args: []string{
				fmt.Sprintf("--%s=2", flags.FlagLimit),
				fmt.Sprintf("--%s=2", flags.FlagOffset),
			},
			expectErr:       false,
			expectedClasses: []string{"C03", "C04"},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryClassesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryClassesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				classIDs := make([]string, len(res.Classes))
				for i, class := range res.Classes {
					classIDs[i] = class.ClassId
				}

				s.Require().Equal(tc.expectedClasses, classIDs)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryClassInfo() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		expectedErrMsg    string
		expectedClassInfo *ecocredit.ClassInfo
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:           "invalid credit class",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "not found: invalid request",
		},
		{
			name:      "valid credit class",
			args:      []string{s.classInfo.ClassId},
			expectErr: false,
			expectedClassInfo: &ecocredit.ClassInfo{
				ClassId:    s.classInfo.ClassId,
				Admin:      s.classInfo.Admin,
				Issuers:    s.classInfo.Issuers,
				Metadata:   s.classInfo.Metadata,
				CreditType: s.classInfo.CreditType,
				NumBatches: 4,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryClassInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryClassInfoResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expectedClassInfo, res.Info)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatches() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrMsg      string
		expectedBatchDenoms []string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:           "invalid project id",
			args:           []string{"abcd-e"},
			expectErr:      true,
			expectedErrMsg: "invalid project id",
		},
		{
			name:                "existing project no batches",
			args:                []string{"P02"},
			expectErr:           false,
			expectedBatchDenoms: []string{},
		},
		{
			name:      "no pagination flags",
			args:      []string{"P01"},
			expectErr: false,
			expectedBatchDenoms: []string{
				"C01-20210101-20210201-001",
				"C01-20210101-20210201-002",
				"C01-20210101-20210201-003",
				"C01-20210101-20210201-004",
			},
		},
		{
			name: "limit 2",
			args: []string{
				"P01",
				fmt.Sprintf("--%s=2", flags.FlagLimit),
			},
			expectErr: false,
			expectedBatchDenoms: []string{
				"C01-20210101-20210201-001",
				"C01-20210101-20210201-002",
			},
		},
		{
			name: "limit 2, offset 2",
			args: []string{
				"P01",
				fmt.Sprintf("--%s=2", flags.FlagLimit),
				fmt.Sprintf("--%s=2", flags.FlagOffset),
			},
			expectErr: false,
			expectedBatchDenoms: []string{
				"C01-20210101-20210201-003",
				"C01-20210101-20210201-004",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryBatchesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				batchDenoms := make([]string, len(res.Batches))
				for i, batch := range res.Batches {
					batchDenoms[i] = batch.BatchDenom
				}

				s.Require().Equal(tc.expectedBatchDenoms, batchDenoms)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchInfo() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		expectedErrMsg    string
		expectedBatchInfo *ecocredit.BatchInfo
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:           "malformed batch denom",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "invalid denom",
		},
		{
			name:           "non-existent credit batch",
			args:           []string{"A00-00000000-00000000-000"},
			expectErr:      true,
			expectedErrMsg: "not found",
		},
		{
			name:              "valid credit batch",
			args:              []string{s.batchInfo.BatchDenom},
			expectErr:         false,
			expectedBatchInfo: s.batchInfo,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryBatchInfoResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expectedBatchInfo, res.Info)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBalance() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	testCases := []struct {
		name                   string
		args                   []string
		expectErr              bool
		expectedErrMsg         string
		expectedTradableAmount string
		expectedRetiredAmount  string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name:           "invalid credit batch",
			args:           []string{"abcde", s.network.Validators[0].Address.String()},
			expectErr:      true,
			expectedErrMsg: "invalid denom",
		},
		{
			name:                   "valid credit batch and invalid account",
			args:                   []string{s.batchInfo.BatchDenom, "abcde"},
			expectErr:              true,
			expectedTradableAmount: "0",
			expectedRetiredAmount:  "0",
		},
		{
			name:                   "valid credit batch and account with no funds",
			args:                   []string{s.batchInfo.BatchDenom, s.network.Validators[2].Address.String()},
			expectErr:              false,
			expectedTradableAmount: "0",
			expectedRetiredAmount:  "0",
		},
		{
			name:                   "valid credit batch and account with enough funds",
			args:                   []string{s.batchInfo.BatchDenom, s.network.Validators[0].Address.String()},
			expectErr:              false,
			expectedTradableAmount: "100",
			expectedRetiredAmount:  "0.000001",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBalanceCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryBalanceResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expectedTradableAmount, res.TradableAmount)
				s.Require().Equal(tc.expectedRetiredAmount, res.RetiredAmount)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySupply() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	testCases := []struct {
		name                   string
		args                   []string
		expectErr              bool
		expectedErrMsg         string
		expectedTradableSupply string
		expectedRetiredSupply  string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:           "invalid credit batch",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "invalid denom",
		},
		{
			name:                   "valid credit batch",
			args:                   []string{s.batchInfo.BatchDenom},
			expectErr:              false,
			expectedTradableSupply: "100",
			expectedRetiredSupply:  "0.000001",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySupplyCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QuerySupplyResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expectedTradableSupply, res.TradableSupply)
				s.Require().Equal(tc.expectedRetiredSupply, res.RetiredSupply)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCreditTypes() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name               string
		args               []string
		expectErr          bool
		expectedErrMsg     string
		expectedCreditType []*ecocredit.CreditType
	}{
		{
			name:               "should give credit type",
			args:               []string{},
			expectErr:          false,
			expectedErrMsg:     "",
			expectedCreditType: []*ecocredit.CreditType{s.classInfo.CreditType},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryCreditTypesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryCreditTypesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expectedCreditType, res.CreditTypes)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryParams() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	require := s.Require()

	cmd := client.QueryParamsCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})
	require.NoError(err)

	var params ecocredit.QueryParamsResponse
	err = json.Unmarshal(out.Bytes(), &params)
	require.NoError(err)

	require.Equal(ecocredit.DefaultParams(), *params.Params)
}

func (s *IntegrationTestSuite) TestQuerySellOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrder  *ecocredit.SellOrder
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "invalid sell order",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "invalid sell order",
		},
		{
			name:      "valid",
			args:      []string{"1"},
			expErr:    false,
			expErrMsg: "",
			expOrder:  s.sellOrders[0],
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySellOrderCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QuerySellOrderResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expOrder, res.SellOrder)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrders() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrders []*ecocredit.SellOrder
	}{
		{
			name:      "too many args",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		{
			name:      "valid",
			args:      []string{},
			expErr:    false,
			expErrMsg: "",
			expOrders: s.sellOrders,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySellOrdersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QuerySellOrdersResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expOrders, res.SellOrders)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersByAddress() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrders []*ecocredit.SellOrder
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "invalid address",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "invalid request",
		},
		{
			name:      "valid",
			args:      []string{val.Address.String()},
			expErr:    false,
			expErrMsg: "",
			expOrders: s.sellOrders,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySellOrdersByAddressCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QuerySellOrdersByAddressResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expOrders, res.SellOrders)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersByBatchDenom() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrders []*ecocredit.SellOrder
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "invalid denom",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "invalid request",
		},
		{
			name:      "valid",
			args:      []string{batchDenom},
			expErr:    false,
			expErrMsg: "",
			expOrders: s.sellOrders,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySellOrdersByBatchDenomCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QuerySellOrdersByBatchDenomResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expOrders, res.SellOrders)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBuyOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrder  *ecocredit.BuyOrder
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "invalid buy order",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "invalid buy order",
		},
		// TODO: filtered buy orders required #623
		//{
		//	name:      "valid",
		//	args:      []string{"1"},
		//	expErr:    false,
		//	expErrMsg: "",
		//	expOrder: &ecocredit.BuyOrder{
		//		BuyOrderId:        1,
		//		Buyer:             val.Address.String(),
		//		Quantity:          "1",
		//		BidPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(100)},
		//		DisableAutoRetire: false,
		//	},
		//},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBuyOrderCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryBuyOrderResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expOrder, res.BuyOrder)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBuyOrders() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrders []*ecocredit.BuyOrder
	}{
		{
			name:      "too many args",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		// TODO: filtered buy orders required #623
		//{
		//	name:      "valid",
		//	args:      []string{},
		//	expErr:    false,
		//	expErrMsg: "",
		//	expOrders: []*ecocredit.BuyOrder{
		//		{
		//			BuyOrderId:        1,
		//			Buyer:             val.Address.String(),
		//			Quantity:          "1",
		//			BidPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(100)},
		//			DisableAutoRetire: false,
		//		},
		//		{
		//			BuyOrderId:        2,
		//			Buyer:             val.Address.String(),
		//			Quantity:          "1",
		//			BidPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(100)},
		//			DisableAutoRetire: false,
		//		},
		//		{
		//			BuyOrderId:        3,
		//			Buyer:             val.Address.String(),
		//			Quantity:          "1",
		//			BidPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(100)},
		//			DisableAutoRetire: false,
		//		},
		//	},
		//},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBuyOrdersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryBuyOrdersResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expOrders, res.BuyOrders)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBuyOrdersByAddress() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrders []*ecocredit.BuyOrder
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "invalid address",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "invalid request",
		},
		// TODO: filtered buy orders required #623
		//{
		//	name:      "valid",
		//	args:      []string{val.Address.String()},
		//	expErr:    false,
		//	expErrMsg: "",
		//	expOrders: []*ecocredit.BuyOrder{
		//		{
		//			BuyOrderId:        1,
		//			Buyer:             val.Address.String(),
		//			Quantity:          "1",
		//			BidPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(100)},
		//			DisableAutoRetire: false,
		//		},
		//		{
		//			BuyOrderId:        2,
		//			Buyer:             val.Address.String(),
		//			Quantity:          "1",
		//			BidPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(100)},
		//			DisableAutoRetire: false,
		//		},
		//		{
		//			BuyOrderId:        3,
		//			Buyer:             val.Address.String(),
		//			Quantity:          "1",
		//			BidPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(100)},
		//			DisableAutoRetire: false,
		//		},
		//	},
		//},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBuyOrdersByAddressCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryBuyOrdersByAddressResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expOrders, res.BuyOrders)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAllowedAskDenoms() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expDenoms []*ecocredit.AskDenom
	}{
		{
			name:      "too many args",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		// TODO: AllowAskDenom not yet implemented #624
		//{
		//	name:      "valid",
		//	args:      []string{},
		//	expErr:    false,
		//	expErrMsg: "",
		//	expDenoms: []*ecocredit.AskDenom{
		//		{
		//			Denom:        "regen",
		//			DisplayDenom: "uregen",
		//			Exponent:     6,
		//		},
		//	},
		//},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryAllowedAskDenomsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryAllowedAskDenomsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expDenoms, res.AskDenoms)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjects() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expLen    int
	}{
		{
			name:      "no args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "no projects found ",
			args:      []string{"CA10"},
			expErr:    false,
			expErrMsg: "",
		},
		{
			name:      "valid query",
			args:      []string{"C01"},
			expErr:    false,
			expErrMsg: "",
			expLen:    3,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryProjectsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res ecocredit.QueryProjectsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Len(res.Projects, tc.expLen)
			}
		})
	}

}

func (s *IntegrationTestSuite) TestQueryProjectInfo() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	require := s.Require()

	cmd := client.QueryProjectsCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{"C01"})
	require.NoError(err)
	var res ecocredit.QueryProjectsResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.GreaterOrEqual(len(res.Projects), 1)
	project := res.Projects[0]

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "no args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "invalid project id ",
			args:      []string{"A@a@"},
			expErr:    true,
			expErrMsg: "invalid project id",
		},
		{
			name:      "not found",
			args:      []string{"P100"},
			expErr:    true,
			expErrMsg: "not found",
		},
		{
			name:      "valid query",
			args:      []string{project.ProjectId},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryProjectInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err, out.String())

				var res ecocredit.QueryProjectInfoResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.Equal(project, res.Info)
			}
		})
	}

}
