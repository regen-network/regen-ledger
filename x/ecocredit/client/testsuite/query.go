package testsuite

import (
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
			name:      "valid credit class",
			args:      []string{s.classInfo.ClassId},
			expectErr: false,
			expectedClassInfo: &ecocredit.ClassInfo{
				ClassId:    s.classInfo.ClassId,
				Designer:   s.classInfo.Designer,
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
			name:           "missing class id",
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
			name:           "invalid class id",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "class ID didn't match the format",
		},
		{
			name:                "existing class no batches",
			args:                []string{"C02"},
			expectErr:           false,
			expectedBatchDenoms: []string{},
		},
		{
			name:      "no pagination flags",
			args:      []string{"C01"},
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
				"C01",
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
				"C01",
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
			name:           "malformed batch denom",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "denomination didn't match the format",
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
			args:                   []string{"abcde", s.network.Validators[0].Address.String()},
			expectErr:              false,
			expectedTradableAmount: "0",
			expectedRetiredAmount:  "0",
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
			args:                   []string{"abcde"},
			expectErr:              false,
			expectedTradableSupply: "0",
			expectedRetiredSupply:  "0",
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
