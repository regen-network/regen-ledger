package testsuite

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit/base/client"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/genesis"
)

const outputFormat = "JSON"

func (s *IntegrationTestSuite) TestQueryClassesCmd() {
	val := s.network.Validators[0]
	val2 := s.network.Validators[1]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat
	classID := s.createClass(clientCtx, &types.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String()},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              genesis.DefaultClassFee().Fee,
	})

	classID2 := s.createClass(clientCtx, &types.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String(), val2.Address.String()},
		Metadata:         "metadata2",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              genesis.DefaultClassFee().Fee,
	})

	classIDs := [2]string{classID, classID2}

	testCases := []struct {
		name               string
		args               []string
		expectErr          bool
		expectedErrMsg     string
		expectedAmtClasses int
	}{
		{
			name:           "too many args",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		{
			name:               "no pagination flags",
			args:               []string{},
			expectErr:          false,
			expectedAmtClasses: -1,
		},
		{
			name: "limit 1",
			args: []string{
				fmt.Sprintf("--%s=1", flags.FlagLimit),
			},
			expectErr:          false,
			expectedAmtClasses: 1,
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

				var res types.QueryClassesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				if tc.expectedAmtClasses > 0 {
					s.Require().Len(res.Classes, tc.expectedAmtClasses)
				} else {
					resClassIds := make([]string, len(res.Classes))
					for i, cls := range res.Classes {
						resClassIds[i] = cls.Id
					}
					for _, id := range classIDs {
						s.Require().Contains(resClassIds, id)
					}
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryClassCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat
	class := &types.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String()},
		Metadata:         "hi",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              genesis.DefaultClassFee().Fee,
	}

	classID := s.createClass(clientCtx, class)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
		expectedClass  *types.ClassInfo
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
			name:      "valid credit class",
			args:      []string{classID},
			expectErr: false,
			expectedClass: &types.ClassInfo{
				Id:               classID,
				Admin:            val.Address.String(),
				Metadata:         class.Metadata,
				CreditTypeAbbrev: class.CreditTypeAbbrev,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryClassCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryClassResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expectedClass, res.Class)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesCmd() {
	ctx := s.val.ClientCtx
	ctx.OutputFormat = outputFormat

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "too many args",
			args:           []string{"foo"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		{
			name: "valid with pagination",
			args: []string{
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchesCmd()
			out, err := cli.ExecTestCLICmd(ctx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryBatchesResponse
				s.Require().NoError(ctx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().True(len(res.Batches) > 0)
				s.Require().NotNil(res.Pagination)
				s.Require().True(res.Pagination.Total > 0)
				denoms := make([]string, len(res.Batches))
				for i, batch := range res.Batches {
					denoms[i] = batch.Denom
				}
				s.Require().Contains(denoms, s.batchDenom)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByIssuerCmd() {
	ctx := s.val.ClientCtx
	ctx.OutputFormat = outputFormat

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"foo", "bar"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name: "valid with pagination",
			args: []string{
				s.addr1.String(),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchesByIssuerCmd()
			out, err := cli.ExecTestCLICmd(ctx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryBatchesByIssuerResponse
				s.Require().NoError(ctx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().True(len(res.Batches) > 0)
				s.Require().NotNil(res.Pagination)
				s.Require().True(res.Pagination.Total > 0)
				denoms := make([]string, len(res.Batches))
				for i, batch := range res.Batches {
					denoms[i] = batch.Denom
				}
				s.Require().Contains(denoms, s.batchDenom)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByClassCmd() {
	ctx := s.val.ClientCtx
	ctx.OutputFormat = outputFormat

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"foo", "bar"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name: "valid with pagination",
			args: []string{
				s.classID,
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchesByClassCmd()
			out, err := cli.ExecTestCLICmd(ctx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryBatchesByClassResponse
				s.Require().NoError(ctx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().True(len(res.Batches) > 0)
				s.Require().NotNil(res.Pagination)
				s.Require().True(res.Pagination.Total > 0)
				denoms := make([]string, len(res.Batches))
				for i, batch := range res.Batches {
					denoms[i] = batch.Denom
				}
				s.Require().Contains(denoms, s.batchDenom)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByProjectCmd() {
	ctx := s.val.ClientCtx
	ctx.OutputFormat = outputFormat

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"foo", "bar"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name: "valid with pagination",
			args: []string{
				s.projectID,
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchesByProjectCmd()
			out, err := cli.ExecTestCLICmd(ctx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryBatchesByProjectResponse
				s.Require().NoError(ctx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().True(len(res.Batches) > 0)
				s.Require().NotNil(res.Pagination)
				s.Require().True(res.Pagination.Total > 0)
				denoms := make([]string, len(res.Batches))
				for i, batch := range res.Batches {
					denoms[i] = batch.Denom
				}
				s.Require().Contains(denoms, s.batchDenom)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchCmd() {
	ctx := s.val.ClientCtx
	ctx.OutputFormat = outputFormat

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
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
			name:      "valid credit batch",
			args:      []string{s.batchDenom},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchCmd()
			out, err := cli.ExecTestCLICmd(ctx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryBatchResponse
				s.Require().NoError(ctx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(res.Batch.Denom, s.batchDenom)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBalanceCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

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
			name:                   "valid",
			args:                   []string{s.batchDenom, val.Address.String()},
			expectErr:              false,
			expectedTradableAmount: "100",
			expectedRetiredAmount:  "0.000001",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchBalanceCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryBalanceResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(res.Balance.Address, val.Address.String())
				s.Require().NotEmpty(res.Balance.TradableAmount)
				s.Require().NotEmpty(res.Balance.RetiredAmount)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySupplyCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
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
			name:      "valid credit batch",
			args:      []string{s.batchDenom},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBatchSupplyCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QuerySupplyResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().NotEmpty(res.TradableAmount)
				s.Require().NotEmpty(res.RetiredAmount)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCreditTypesCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat
	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "valid",
			args:           []string{},
			expectErr:      false,
			expectedErrMsg: "",
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

				var res types.QueryCreditTypesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Greater(len(res.CreditTypes), 0)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryParamsCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	require := s.Require()

	cmd := client.QueryParamsCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})
	require.NoError(err)

	var params types.QueryParamsResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &params))
	require.NoError(err)

	require.Equal(genesis.DefaultBasketFee().Fee.Amount, params.Params.BasketFee[0].Amount)
	require.Equal(genesis.DefaultBasketFee().Fee.Denom, params.Params.BasketFee[0].Denom)

	require.Equal(genesis.DefaultClassFee().Fee.Amount, params.Params.CreditClassFee[0].Amount)
	require.Equal(genesis.DefaultClassFee().Fee.Denom, params.Params.CreditClassFee[0].Denom)
	require.False(params.Params.AllowlistEnabled)
	require.Equal([]string{sdk.AccAddress("issuer1").String(), sdk.AccAddress("issuer2").String()}, params.Params.AllowedClassCreators)
}

func (s *IntegrationTestSuite) TestQueryProjectsCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "too many args",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		{
			name: "valid",
			args: []string{},
		},
		{
			name: "valid within pagination",
			args: []string{
				fmt.Sprintf("--%s", flags.FlagCountTotal),
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryProjectsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res types.QueryProjectsResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Projects)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Projects, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjectsByClassCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
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
			name: "valid",
			args: []string{s.classID},
		},
		{
			name: "valid with pagination",
			args: []string{
				s.classID,
				fmt.Sprintf("--%s", flags.FlagCountTotal),
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryProjectsByClassCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res types.QueryProjectsByClassResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Projects)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Projects, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjectCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
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
			name: "valid query",
			args: []string{s.projectID},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryProjectCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res types.QueryProjectResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Project)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryClassIssuersCmd() {
	val := s.network.Validators[0]
	val2 := s.network.Validators[1]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat
	require := s.Require()

	classID := s.createClass(clientCtx, &types.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String(), val2.Address.String()},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              genesis.DefaultClassFee().Fee,
	})

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
		numItems       int
	}{
		{
			name:           "no pagination flags",
			args:           []string{classID},
			expectErr:      false,
			expectedErrMsg: "",
			numItems:       -1,
		},
		{
			name:           "pagination limit 1",
			args:           []string{classID, "--limit=1"},
			expectErr:      false,
			expectedErrMsg: "",
			numItems:       1,
		},
		{
			name:           "class not found",
			args:           []string{"Z100"},
			expectErr:      true,
			expectedErrMsg: "not found",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryClassIssuersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				require.Error(err)
				require.Contains(out.String(), tc.expectedErrMsg)
			} else {
				require.NoError(err, out.String())

				var res types.QueryClassIssuersResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				if tc.numItems > 0 {
					require.Len(res.Issuers, tc.numItems)
				} else {
					require.GreaterOrEqual(len(res.Issuers), 1)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCreditTypeCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
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
			name:      "valid credit type",
			args:      []string{"C"},
			expectErr: false,
		},
		{
			name:           "unknown credit type",
			args:           []string{"CD"},
			expectErr:      true,
			expectedErrMsg: "not found",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryCreditTypeCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryCreditTypeResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(res.CreditType.Abbreviation, "C")
				s.Require().Equal(res.CreditType.Precision, uint32(6))
				s.Require().Equal(res.CreditType.Name, "carbon")
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAllowedClassCreatorsCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	cmd := client.QueryAllowedClassCreatorsCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})

	s.Require().NoError(err, out.String())

	var res types.QueryAllowedClassCreatorsResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	s.Require().Len(res.ClassCreators, 2)
	s.Require().Equal(res.ClassCreators[0], sdk.AccAddress("issuer1").String())
	s.Require().Equal(res.ClassCreators[1], sdk.AccAddress("issuer2").String())
}

func (s *IntegrationTestSuite) TestQueryCreditClassAllowlistEnableCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	cmd := client.QueryClassCreatorAllowlistCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})

	s.Require().NoError(err, out.String())

	var res types.QueryClassCreatorAllowlistResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	s.Require().False(res.Enabled)
}

func (s *IntegrationTestSuite) TestQueryClassFeeCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	cmd := client.QueryClassFeeCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})

	s.Require().NoError(err, out.String())

	var res types.QueryClassFeeResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

	s.Require().NotEmpty(res.Fee)
	s.Require().Equal(res.Fee.Denom, sdk.DefaultBondDenom)
	s.Require().Equal(res.Fee.Amount, types.DefaultClassFee)
}

func (s *IntegrationTestSuite) TestQueryAllBalancesCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	cmd := client.QueryAllBalancesCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{fmt.Sprintf("--%s", flags.FlagCountTotal)})

	s.Require().NoError(err, out.String())

	var res types.QueryAllBalancesResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	s.Require().Greater(len(res.Balances), 0)
	s.Require().Greater(res.Pagination.Total, uint64(0))
}

func (s *IntegrationTestSuite) TestQueryBalancesByBatchCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
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
			name:      "valid",
			args:      []string{s.batchDenom},
			expectErr: false,
		},
		{
			name: "valid with pagination",
			args: []string{
				s.batchDenom,
				fmt.Sprintf("--%s", flags.FlagCountTotal),
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBalancesByBatchCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res types.QueryBalancesByBatchResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().NotEmpty(res.Balances)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAllowedBridgeChainsCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	cmd := client.QueryAllowedBridgeChainsCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})

	s.Require().NoError(err, out.String())

	var res types.QueryAllowedBridgeChainsResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	s.Require().Len(res.AllowedBridgeChains, 1)
	s.Require().Equal(res.AllowedBridgeChains[0], s.bridgeChain)
}
