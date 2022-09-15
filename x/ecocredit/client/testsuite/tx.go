package testsuite

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit/base/client"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func (s *IntegrationTestSuite) TestTxCreateClassCmd() {
	require := s.Require()

	admin := s.addr1.String()
	creditClassFee := s.creditClassFee.String()

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
			expErrMsg: "Error: accepts 3 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar", "baz", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 3 arg(s), received 4",
		},
		{
			name: "missing from flag",
			args: []string{
				admin,
				s.creditTypeAbbrev,
				"metadata",
				fmt.Sprintf("--%s=%s", client.FlagClassFee, creditClassFee),
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "invalid coin expression",
			args: []string{
				admin,
				s.creditTypeAbbrev,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", client.FlagClassFee, "foo"),
			},
			expErr:    true,
			expErrMsg: "invalid decimal coin expression",
		},
		{
			name: "valid",
			args: []string{
				admin,
				s.creditTypeAbbrev,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", client.FlagClassFee, creditClassFee),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				admin,
				s.creditTypeAbbrev,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
				fmt.Sprintf("--%s=%s", client.FlagClassFee, creditClassFee),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				admin,
				s.creditTypeAbbrev,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				fmt.Sprintf("--%s=%s", client.FlagClassFee, creditClassFee),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxCreateClassCmd()
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

func (s *IntegrationTestSuite) TestTxCreateProjectCmd() {
	require := s.Require()

	admin := s.addr1.String()

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "missing args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 3 arg(s), received 2",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar", "baz", "foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 3 arg(s), received 4",
		},
		{
			name: "missing from flag",
			args: []string{
				s.classID,
				"US-WA",
				"metadata",
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				s.classID,
				"US-WA",
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.classID,
				"US-WA",
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.classID,
				"US-WA",
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxCreateProjectCmd()
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

func (s *IntegrationTestSuite) TestTxCreateBatchCmd() {
	require := s.Require()

	issuer := s.addr1.String()
	recipient := s.addr2.String()

	startDate, err := regentypes.ParseDate("start date", "2020-01-01")
	require.NoError(err)

	endDate, err := regentypes.ParseDate("end date", "2021-01-01")
	require.NoError(err)

	bz, err := s.val.ClientCtx.Codec.MarshalJSON(&types.MsgCreateBatch{
		Issuer:    issuer,
		ProjectId: s.projectID,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:              recipient,
				TradableAmount:         "10",
				RetiredAmount:          "10",
				RetirementJurisdiction: "US-WA",
			},
			{
				Recipient:              recipient,
				TradableAmount:         "10",
				RetiredAmount:          "10",
				RetirementJurisdiction: "US-WA",
			},
		},
		Metadata:  "metadata",
		StartDate: &startDate,
		EndDate:   &endDate,
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
			name:      "missing from flag",
			args:      []string{validJSON},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "invalid json file",
			args: []string{
				"foo.bar",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
			},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "invalid json format",
			args: []string{
				invalidJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				duplicateJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxCreateBatchCmd()
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

func (s *IntegrationTestSuite) TestTxSendCmd() {
	require := s.Require()

	sender := s.addr1.String()
	amount := "10"
	retirementJurisdiction := "US-WA"

	recipient := s.addr2.String()

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "missing args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 3 arg(s), received 2",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar", "baz", "foobarbaz"},
			expErr:    true,
			expErrMsg: "Error: accepts 3 arg(s), received 4",
		},
		{
			name: "missing from flag",
			args: []string{
				amount,
				s.batchDenom,
				recipient,
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid tradeable",
			args: []string{
				amount,
				s.batchDenom,
				recipient,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
			},
		},
		{
			name: "valid retire",
			args: []string{
				amount,
				s.batchDenom,
				recipient,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
				fmt.Sprintf("--%s=%s", client.FlagRetirementJurisdiction, retirementJurisdiction),
			},
		},
	}
	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxSendCmd()
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

func (s *IntegrationTestSuite) TestTxSendBulkCmd() {
	require := s.Require()

	sender := s.addr1.String()
	recipient := s.addr2.String()

	// using json package because array is not a proto message
	bz, err := json.Marshal([]types.MsgSend_SendCredits{
		{
			BatchDenom:             s.batchDenom,
			TradableAmount:         "10",
			RetiredAmount:          "10",
			RetirementJurisdiction: "US-WA",
		},
		{
			BatchDenom:             s.batchDenom,
			TradableAmount:         "10",
			RetiredAmount:          "10",
			RetirementJurisdiction: "US-WA",
		},
	})
	require.NoError(err)

	validJSON := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()
	invalidJSON := testutil.WriteToNewTempFile(s.T(), "{foo:bar}").Name()
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
			name: "missing from flag",
			args: []string{
				recipient,
				validJSON,
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "invalid json file",
			args: []string{
				recipient,
				"foo.bar",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
			},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "invalid json format",
			args: []string{
				recipient,
				invalidJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				recipient,
				duplicateJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				recipient,
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				recipient,
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				recipient,
				validJSON,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxSendBulkCmd()
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

func (s *IntegrationTestSuite) TestTxRetire() {
	require := s.Require()

	owner := s.addr1.String()

	// using json package because array is not a proto message
	bz, err := json.Marshal([]types.Credits{
		{
			BatchDenom: s.batchDenom,
			Amount:     "10",
		},
		{
			BatchDenom: s.batchDenom,
			Amount:     "10",
		},
	})
	require.NoError(err)

	validJSON := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()
	invalidJSON := testutil.WriteToNewTempFile(s.T(), "{foo:bar}").Name()
	duplicateJSON := testutil.WriteToNewTempFile(s.T(), `{"foo":"bar","foo":"bar"`).Name()

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
			args:      []string{"foo", "bar", "baz"},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name:      "missing from flag",
			args:      []string{validJSON, "US-WA"},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "invalid json file",
			args: []string{
				"foo.bar",
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "invalid json format",
			args: []string{
				invalidJSON,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				duplicateJSON,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				validJSON,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				validJSON,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				validJSON,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxRetireCmd()
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

func (s *IntegrationTestSuite) TestTxCancel() {
	require := s.Require()

	owner := s.addr1.String()

	// using json package because array is not a proto message
	bz, err := json.Marshal([]types.Credits{
		{
			BatchDenom: s.batchDenom,
			Amount:     "10",
		},
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
			args:      []string{validJSON, "reason"},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "invalid json file",
			args: []string{
				"foo.bar",
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "invalid json format",
			args: []string{
				invalidJSON,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				duplicateJSON,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				validJSON,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				validJSON,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				validJSON,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxCancelCmd()
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

func (s *IntegrationTestSuite) TestTxUpdateClassAdmin() {
	require := s.Require()

	admin := s.addr1.String()
	newAdmin := s.addr2.String()

	// create new credit class to not interfere with other tests
	classID1 := s.createClass(s.val.ClientCtx, &types.MsgCreateClass{
		Admin:            admin,
		Issuers:          []string{admin},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              s.creditClassFee,
	})

	// create new credit class to not interfere with other tests
	classID2 := s.createClass(s.val.ClientCtx, &types.MsgCreateClass{
		Admin:            admin,
		Issuers:          []string{admin},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              s.creditClassFee,
	})

	// create new credit class to not interfere with other tests
	classID3 := s.createClass(s.val.ClientCtx, &types.MsgCreateClass{
		Admin:            admin,
		Issuers:          []string{admin},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              s.creditClassFee,
	})

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
			name: "missing from flag",
			args: []string{
				s.classID,
				newAdmin,
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				classID1,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				classID2,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				classID3,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxUpdateClassAdminCmd()
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

func (s *IntegrationTestSuite) TestTxUpdateIssuers() {
	require := s.Require()

	admin := s.addr1.String()
	issuer := s.addr2.String()

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
			name:      "missing from flag",
			args:      []string{s.classID},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "missing add or remove flag",
			args: []string{
				s.classID,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
			expErr:    true,
			expErrMsg: "must specify at least one of add_issuers or remove_issuers",
		},
		{
			name: "valid add issuer",
			args: []string{
				s.classID,
				fmt.Sprintf("--%s=%s", client.FlagAddIssuers, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid remove issuer",
			args: []string{
				s.classID,
				fmt.Sprintf("--%s=%s", client.FlagRemoveIssuers, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.classID,
				fmt.Sprintf("--%s=%s", client.FlagAddIssuers, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.classID,
				fmt.Sprintf("--%s=%s", client.FlagRemoveIssuers, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxUpdateClassIssuersCmd()
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

func (s *IntegrationTestSuite) TestTxUpdateClassMetadata() {
	require := s.Require()

	admin := s.addr1.String()

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
			name: "missing from flag",
			args: []string{
				s.classID,
				"metadata",
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				s.classID,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.classID,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.classID,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxUpdateClassMetadataCmd()
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

func (s *IntegrationTestSuite) TestUpdateProjectAdmin() {
	require := s.Require()

	admin := s.addr1.String()
	newAdmin := s.addr2.String()

	// create new project in order to not interfere with other tests
	projectID1 := s.createProject(s.val.ClientCtx, &types.MsgCreateProject{
		Admin:        admin,
		ClassId:      s.classID,
		Metadata:     "metadata",
		Jurisdiction: "US-WA",
		ReferenceId:  "VCS-002",
	})

	// create new project in order to not interfere with other tests
	projectID2 := s.createProject(s.val.ClientCtx, &types.MsgCreateProject{
		Admin:        admin,
		ClassId:      s.classID,
		Metadata:     "metadata",
		Jurisdiction: "US-WA",
		ReferenceId:  "VCS-003",
	})

	// create new project in order to not interfere with other tests
	projectID3 := s.createProject(s.val.ClientCtx, &types.MsgCreateProject{
		Admin:        admin,
		ClassId:      s.classID,
		Metadata:     "metadata",
		Jurisdiction: "US-WA",
		ReferenceId:  "VCS-004",
	})

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
			name: "missing from flag",
			args: []string{
				s.projectID,
				newAdmin,
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				projectID1,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				projectID2,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				projectID3,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxUpdateProjectAdminCmd()
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

func (s *IntegrationTestSuite) TestUpdateProjectMetadata() {
	require := s.Require()

	admin := s.addr1.String()

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
			name: "missing from flag",
			args: []string{
				s.projectID,
				"metadata",
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				s.projectID,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.projectID,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.projectID,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			cmd := client.TxUpdateProjectMetadataCmd()
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
