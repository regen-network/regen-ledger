package testsuite

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket/client"
	types "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
)

func (s *IntegrationTestSuite) TestTxCreateBasketCmd() {
	require := s.Require()

	curator := s.addr1.String()

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
			name:      "missing required flags",
			args:      []string{"NCT"},
			expErr:    true,
			expErrMsg: "Error: required flag(s)",
		},
		{
			name: "valid",
			args: []string{
				"NCT1",
				fmt.Sprintf("--%s=%s", client.FlagAllowedClasses, s.classID),
				fmt.Sprintf("--%s=%s", client.FlagCreditTypeAbbrev, s.creditTypeAbbrev),
				fmt.Sprintf("--%s=%s", client.FlagBasketFee, s.basketFee),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, curator),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				"NCT2",
				fmt.Sprintf("--%s=%s", client.FlagAllowedClasses, s.classID),
				fmt.Sprintf("--%s=%s", client.FlagCreditTypeAbbrev, s.creditTypeAbbrev),
				fmt.Sprintf("--%s=%s", client.FlagBasketFee, s.basketFee),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				"NCT3",
				fmt.Sprintf("--%s=%s", client.FlagAllowedClasses, s.classID),
				fmt.Sprintf("--%s=%s", client.FlagCreditTypeAbbrev, s.creditTypeAbbrev),
				fmt.Sprintf("--%s=%s", client.FlagBasketFee, s.basketFee),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, curator),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxCreateBasketCmd()
			args = append(args, s.commonTxFlags()...)
			out, err := cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res sdk.TxResponse
				require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.Zero(res.Code, res.RawLog)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxPutInBasketCmd() {
	require := s.Require()

	owner := s.addr1.String()

	// using json package because array is not a proto message
	bz, err := json.Marshal([]*types.BasketCredit{
		{
			BatchDenom: s.batchDenom,
			Amount:     "10",
		},
	})
	require.NoError(err)

	validJSON := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()
	invalidJSON := testutil.WriteToNewTempFile(s.T(), `{foo:bar}`).Name()
	duplicateJSON := testutil.WriteToNewTempFile(s.T(), `{"foo":"bar","foo":"bar"`).Name()

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
			name:      "missing from flag",
			args:      []string{s.basketDenom, validJSON},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "invalid json file",
			args: []string{
				s.basketDenom,
				"foo.bar",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "invalid json format",
			args: []string{
				s.basketDenom,
				invalidJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				s.basketDenom,
				duplicateJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				s.basketDenom,
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.basketDenom,
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.basketDenom,
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxPutInBasketCmd()
			args = append(args, s.commonTxFlags()...)
			out, err := cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res sdk.TxResponse
				require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.Zero(res.Code, res.RawLog)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxTakeFromBasketCmd() {
	require := s.Require()

	owner := s.addr1.String()

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
			name:      "missing from flag",
			args:      []string{s.basketDenom, "10"},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				s.basketDenom,
				"10",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
				fmt.Sprintf("--%s=true", client.FlagRetireOnTake),
				fmt.Sprintf("--%s=AQ", client.FlagRetirementJurisdiction),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.basketDenom,
				"10",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
				fmt.Sprintf("--%s=true", client.FlagRetireOnTake),
				fmt.Sprintf("--%s=AQ", client.FlagRetirementJurisdiction),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.basketDenom,
				"10",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				fmt.Sprintf("--%s=true", client.FlagRetireOnTake),
				fmt.Sprintf("--%s=AQ", client.FlagRetirementJurisdiction),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxTakeFromBasketCmd()
			args = append(args, s.commonTxFlags()...)
			out, err := cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res sdk.TxResponse
				require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.Zero(res.Code, res.RawLog)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateBasketCuratorCmd() {
	require := s.Require()

	curator := s.addr1.String()
	newCurator := s.addr2.String()
	cmd := client.TxCreateBasketCmd()
	args := append([]string{
		"NCT2",
		fmt.Sprintf("--%s=%s", client.FlagAllowedClasses, s.classID),
		fmt.Sprintf("--%s=%s", client.FlagCreditTypeAbbrev, s.creditTypeAbbrev),
		fmt.Sprintf("--%s=%s", client.FlagBasketFee, s.basketFee),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, curator),
	}, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
	require.NoError(err)
	require.NotNil(out)

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
			expErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar", "foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name:      "missing required flags",
			args:      []string{"eco.uC.NCT2", curator},
			expErr:    true,
			expErrMsg: "Error: required flag(s)",
		},
		{
			name: "valid",
			args: []string{
				"eco.uC.NCT2",
				newCurator,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, curator),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				"eco.uC.NCT2",
				curator,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, newCurator),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxUpdateBasketCuratorCmd()
			args = append(args, s.commonTxFlags()...)
			out, err := cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res sdk.TxResponse
				require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.Zero(res.Code, res.RawLog)
			}
		})
	}
}
