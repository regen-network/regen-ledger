package testsuite

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	client "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (s *IntegrationTestSuite) TestQuerySellOrderCmd() {
	require := s.Require()

	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = "JSON"

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
			args: []string{fmt.Sprint(s.sellOrderId)},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySellOrderCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res marketplace.QuerySellOrderResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.SellOrder)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersCmd() {
	require := s.Require()

	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = "JSON"

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
			name: "valid with pagination",
			args: []string{
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySellOrdersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res marketplace.QuerySellOrdersResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.SellOrders)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.SellOrders, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersBySellerCmd() {
	require := s.Require()

	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = "JSON"

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
			args: []string{s.val.Address.String()},
		},
		{
			name: "valid with pagination",
			args: []string{
				s.val.Address.String(),
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySellOrdersBySellerCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res marketplace.QuerySellOrdersBySellerResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.SellOrders)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.SellOrders, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersByBatchCmd() {
	require := s.Require()

	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = "JSON"

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
			args: []string{s.batchDenom},
		},
		{
			name: "valid with pagination",
			args: []string{
				s.batchDenom,
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySellOrdersByBatchCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res marketplace.QuerySellOrdersByBatchResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.SellOrders)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.SellOrders, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAllowedDenomsCmd() {
	require := s.Require()

	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = "JSON"

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
			name: "valid with pagination",
			args: []string{
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryAllowedDenomsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res marketplace.QueryAllowedDenomsResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.AllowedDenoms)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.AllowedDenoms, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}
