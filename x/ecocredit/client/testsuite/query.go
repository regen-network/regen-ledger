package testsuite

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
	marketplaceclient "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
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
			cmd := coreclient.QueryClassesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryClassesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				classIDs := make([]string, len(res.Classes))
				for i, class := range res.Classes {
					classIDs[i] = class.Name
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
		expectedClassInfo *core.ClassInfo
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
			args:      []string{s.classInfo.Name},
			expectErr: false,
			expectedClassInfo: &core.ClassInfo{
				Id:         s.classInfo.Id,
				Name:       s.classInfo.Name,
				Admin:      s.classInfo.Admin,
				Metadata:   s.classInfo.Metadata,
				CreditType: s.classInfo.CreditType,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryClassInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryClassInfoResponse
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
			cmd := coreclient.QueryBatchesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryBatchesResponse
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
		expectedBatchInfo *core.BatchInfo
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
			cmd := coreclient.QueryBatchInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryBatchInfoResponse
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
			cmd := coreclient.QueryBalanceCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryBalanceResponse
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
			cmd := coreclient.QuerySupplyCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QuerySupplyResponse
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
		expectedCreditType []*core.CreditType
	}{
		{
			name:               "should give credit type",
			args:               []string{},
			expectErr:          false,
			expectedErrMsg:     "",
			expectedCreditType: []*core.CreditType{&core.CreditType{Abbreviation: s.classInfo.CreditType}},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryCreditTypesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryCreditTypesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expectedCreditType, res.CreditTypes)
			}
		})
	}
}

// TODO: migrate Params to ORM #854
// func (s *IntegrationTestSuite) TestQueryParams() {
// 	val := s.network.Validators[0]
// 	clientCtx := val.ClientCtx
// 	clientCtx.OutputFormat = "JSON"
// 	require := s.Require()

// 	cmd := coreclient.QueryParamsCmd()
// 	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})
// 	require.NoError(err)

// 	var params core.QueryParamsResponse
// 	err = json.Unmarshal(out.Bytes(), &params)
// 	require.NoError(err)

// 	require.Equal(core.DefaultParams(), *params.Params)
// }

func (s *IntegrationTestSuite) TestQuerySellOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrder  *marketplace.SellOrder
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
			cmd := marketplaceclient.QuerySellOrderCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QuerySellOrderResponse
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
		expOrders []*marketplace.SellOrder
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
			cmd := marketplaceclient.QuerySellOrdersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QuerySellOrdersResponse
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
		expOrders []*marketplace.SellOrder
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
			cmd := marketplaceclient.QuerySellOrdersByAddressCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QuerySellOrdersByAddressResponse
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
		expOrders []*marketplace.SellOrder
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
			cmd := marketplaceclient.QuerySellOrdersByBatchDenomCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QuerySellOrdersByBatchDenomResponse
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
		expOrder  *marketplace.BuyOrder
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
			cmd := marketplaceclient.QueryBuyOrderCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QueryBuyOrderResponse
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
		expOrders []*marketplace.BuyOrder
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
			cmd := marketplaceclient.QueryBuyOrdersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QueryBuyOrdersResponse
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
		expOrders []*marketplace.BuyOrder
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
			cmd := marketplaceclient.QueryBuyOrdersByAddressCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QueryBuyOrdersByAddressResponse
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
		expDenoms []*marketplace.AllowedDenom
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
			cmd := marketplaceclient.QueryAllowedDenomsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QueryAllowedDenomsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expDenoms, res.AllowedDenoms)
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
			cmd := coreclient.QueryProjectsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryProjectsResponse
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

	cmd := coreclient.QueryProjectsCmd()
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
			cmd := coreclient.QueryProjectInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err, out.String())

				var res core.QueryProjectInfoResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.Equal(project, res.Info)
			}
		})
	}

}
