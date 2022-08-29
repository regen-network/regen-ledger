package testsuite

import (
	"fmt"
	"strings"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket/client"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (s *IntegrationTestSuite) TestQueryBasketCmd() {
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
			args: []string{s.basketDenom},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBasketCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res types.QueryBasketResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Basket) // deprecated
				require.NotEmpty(res.BasketInfo)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBasketsCmd() {
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
			name: "valid with pagination",
			args: []string{
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBasketsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res types.QueryBasketsResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Baskets) // deprecated
				require.NotEmpty(res.BasketsInfo)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Baskets, 1) // deprecated
					require.Len(res.BasketsInfo, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBasketBalanceCmd() {
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
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 1",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar", "baz"},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name: "valid",
			args: []string{s.basketDenom, s.batchDenom},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBasketBalanceCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res types.QueryBasketBalanceResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Balance)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBasketBalancesCmd() {
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
			args: []string{s.basketDenom},
		},
		{
			name: "valid with pagination",
			args: []string{
				s.basketDenom,
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBasketBalancesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res types.QueryBasketBalancesResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Balances) // deprecated
				require.NotEmpty(res.BalancesInfo)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Balances, 1) // deprecated
					require.Len(res.BalancesInfo, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBasketFees() {
	require := s.Require()

	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	cmd := client.QueryBasketFeesCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})
	require.NoError(err)

	var res types.QueryBasketFeesResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.NotEmpty(res.Fees)
	require.Equal(res.Fees.AmountOf(sdk.DefaultBondDenom), math.NewInt(10))
}
