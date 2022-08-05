package testsuite

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
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
			expErrMsg: "Error: accepts 4 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar", "baz", "bar", "foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 4 arg(s), received 5",
		},
		{
			name: "missing from flag",
			args: []string{
				admin,
				s.creditTypeAbbrev,
				"metadata",
				creditClassFee,
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
				"foo",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
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
				creditClassFee,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				admin,
				s.creditTypeAbbrev,
				"metadata",
				creditClassFee,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				admin,
				s.creditTypeAbbrev,
				"metadata",
				creditClassFee,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxCreateClassCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
				s.classId,
				"US-WA",
				"metadata",
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				s.classId,
				"US-WA",
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.classId,
				"US-WA",
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.classId,
				"US-WA",
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxCreateProjectCmd()
			args := append(tc.args, s.commonTxFlags()...)
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

	startDate, err := types.ParseDate("start date", "2020-01-01")
	require.NoError(err)

	endDate, err := types.ParseDate("end date", "2021-01-01")
	require.NoError(err)

	bz, err := s.val.ClientCtx.Codec.MarshalJSON(&core.MsgCreateBatch{
		Issuer:    issuer,
		ProjectId: s.projectId,
		Issuance: []*core.BatchIssuance{
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

	validJson := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()
	invalidJson := testutil.WriteToNewTempFile(s.T(), `{foo:bar}`).Name()
	duplicateJson := testutil.WriteToNewTempFile(s.T(), `{"foo":"bar","foo":"bar"`).Name()

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
			args:      []string{validJson},
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
				invalidJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				duplicateJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				validJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				validJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				validJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxCreateBatchCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
				fmt.Sprintf("--%s=%s", coreclient.FlagRetirementJurisdiction, retirementJurisdiction),
			},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxSendCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
	bz, err := json.Marshal([]core.MsgSend_SendCredits{
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

	validJson := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()
	invalidJson := testutil.WriteToNewTempFile(s.T(), "{foo:bar}").Name()
	duplicateJson := testutil.WriteToNewTempFile(s.T(), `{"foo":"bar","foo":"bar"`).Name()

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
				validJson,
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
				invalidJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				recipient,
				duplicateJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				recipient,
				validJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				recipient,
				validJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				recipient,
				validJson,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, sender),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxSendBulkCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
	bz, err := json.Marshal([]core.Credits{
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

	validJson := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()
	invalidJson := testutil.WriteToNewTempFile(s.T(), "{foo:bar}").Name()
	duplicateJson := testutil.WriteToNewTempFile(s.T(), `{"foo":"bar","foo":"bar"`).Name()

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
			args:      []string{validJson, "US-WA"},
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
				invalidJson,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				duplicateJson,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				validJson,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				validJson,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				validJson,
				"US-WA",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxRetireCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
	bz, err := json.Marshal([]core.Credits{
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

	validJson := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()
	invalidJson := testutil.WriteToNewTempFile(s.T(), `{foo:bar}`).Name()
	duplicateJson := testutil.WriteToNewTempFile(s.T(), `{"foo":"bar","foo":"bar"`).Name()

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
			args:      []string{validJson, "reason"},
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
				invalidJson,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: invalid character",
		},
		{
			name: "duplicate json key",
			args: []string{
				duplicateJson,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
			expErr:    true,
			expErrMsg: "failed to parse json: duplicate key",
		},
		{
			name: "valid",
			args: []string{
				validJson,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				validJson,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				validJson,
				"reason",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, owner),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxCancelCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
	classId1 := s.createClass(s.val.ClientCtx, &core.MsgCreateClass{
		Admin:            admin,
		Issuers:          []string{admin},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              &s.creditClassFee[0],
	})

	// create new credit class to not interfere with other tests
	classId2 := s.createClass(s.val.ClientCtx, &core.MsgCreateClass{
		Admin:            admin,
		Issuers:          []string{admin},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              &s.creditClassFee[0],
	})

	// create new credit class to not interfere with other tests
	classId3 := s.createClass(s.val.ClientCtx, &core.MsgCreateClass{
		Admin:            admin,
		Issuers:          []string{admin},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              &s.creditClassFee[0],
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
				s.classId,
				newAdmin,
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				classId1,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				classId2,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				classId3,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxUpdateClassAdminCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
			args:      []string{s.classId},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "missing add or remove flag",
			args: []string{
				s.classId,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
			expErr:    true,
			expErrMsg: "must specify at least one of add_issuers or remove_issuers",
		},
		{
			name: "valid add issuer",
			args: []string{
				s.classId,
				fmt.Sprintf("--%s=%s", coreclient.FlagAddIssuers, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid remove issuer",
			args: []string{
				s.classId,
				fmt.Sprintf("--%s=%s", coreclient.FlagRemoveIssuers, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.classId,
				fmt.Sprintf("--%s=%s", coreclient.FlagAddIssuers, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.classId,
				fmt.Sprintf("--%s=%s", coreclient.FlagRemoveIssuers, issuer),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxUpdateClassIssuersCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
				s.classId,
				"metadata",
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				s.classId,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.classId,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.classId,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxUpdateClassMetadataCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
	projectId1 := s.createProject(s.val.ClientCtx, &core.MsgCreateProject{
		Admin:        admin,
		ClassId:      s.classId,
		Metadata:     "metadata",
		Jurisdiction: "US-WA",
		ReferenceId:  "VCS-002",
	})

	// create new project in order to not interfere with other tests
	projectId2 := s.createProject(s.val.ClientCtx, &core.MsgCreateProject{
		Admin:        admin,
		ClassId:      s.classId,
		Metadata:     "metadata",
		Jurisdiction: "US-WA",
		ReferenceId:  "VCS-003",
	})

	// create new project in order to not interfere with other tests
	projectId3 := s.createProject(s.val.ClientCtx, &core.MsgCreateProject{
		Admin:        admin,
		ClassId:      s.classId,
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
				s.projectId,
				newAdmin,
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				projectId1,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				projectId2,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				projectId3,
				newAdmin,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxUpdateProjectAdminCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
				s.projectId,
				"metadata",
			},
			expErr:    true,
			expErrMsg: "Error: required flag(s) \"from\" not set",
		},
		{
			name: "valid",
			args: []string{
				s.projectId,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
			},
		},
		{
			name: "valid from key-name",
			args: []string{
				s.projectId,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.val.Moniker),
			},
		},
		{
			name: "valid with amino-json",
			args: []string{
				s.projectId,
				"metadata",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, admin),
				fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxUpdateProjectMetadataCmd()
			args := append(tc.args, s.commonTxFlags()...)
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
