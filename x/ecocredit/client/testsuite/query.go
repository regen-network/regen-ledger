package testsuite

import (
	"fmt"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	tmcli "github.com/tendermint/tendermint/libs/cli"
)

func (s *IntegrationTestSuite) TestQueryClassInfo() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		expectedErrMsg    string
		expectedClassInfo *ecocredit.ClassInfo
	}{
		{
			name:           "missing credit class",
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
			name:              "valid credit class",
			args:              []string{s.classInfo.ClassId, fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr:         false,
			expectedClassInfo: s.classInfo,
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

func (s *IntegrationTestSuite) TestQueryBatchInfo() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		expectedErrMsg    string
		expectedBatchInfo *ecocredit.BatchInfo
	}{
		{
			name:           "missing credit batch",
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
			expectedErrMsg: "not found: invalid request",
		},
		{
			name:              "valid credit batch",
			args:              []string{s.batchInfo.ClassId, fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
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

	testCases := []struct {
		name                   string
		args                   []string
		expectErr              bool
		expectedErrMsg         string
		expectedTradableAmount string
		expectedRetiredAmount  string
	}{
		{
			name:           "missing credit batch",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:           "missing address",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 1",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name:                   "invalid credit batch",
			args:                   []string{"abcde", s.network.Validators[0].Address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr:              false,
			expectedTradableAmount: "0",
			expectedRetiredAmount:  "0",
		},
		{
			name:                   "valid credit batch and invalid account",
			args:                   []string{s.batchInfo.BatchDenom, "abcde", fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr:              false,
			expectedTradableAmount: "0",
			expectedRetiredAmount:  "0",
		},
		{
			name:                   "valid credit batch and account with no funds",
			args:                   []string{s.batchInfo.BatchDenom, s.network.Validators[2].Address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr:              false,
			expectedTradableAmount: "0",
			expectedRetiredAmount:  "0",
		},
		{
			name:                   "valid credit batch and account with enough funds",
			args:                   []string{s.batchInfo.BatchDenom, s.network.Validators[0].Address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
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

	testCases := []struct {
		name                   string
		args                   []string
		expectErr              bool
		expectedErrMsg         string
		expectedTradableSupply string
		expectedRetiredSupply  string
	}{
		{
			name:           "missing credit batch",
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
			name:                   "invalid credit batch",
			args:                   []string{"abcde", fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr:              false,
			expectedTradableSupply: "0",
			expectedRetiredSupply:  "0",
		},
		{
			name:                   "valid credit batch",
			args:                   []string{s.batchInfo.BatchDenom, fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
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
