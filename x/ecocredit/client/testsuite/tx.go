package testsuite

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	gogotypes "github.com/gogo/protobuf/types"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/rand"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
	marketplaceclient "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

const (
	validCreditTypeAbbrev = "C"
	validMetadata         = "hi"
)

// Write a MsgCreateBatch to a new temporary file and return the filename
func (s *IntegrationTestSuite) writeMsgCreateBatchJSON(msg *core.MsgCreateBatch) string {
	bytes, err := s.network.Validators[0].ClientCtx.Codec.MarshalJSON(msg)
	s.Require().NoError(err)

	return testutil.WriteToNewTempFile(s.T(), string(bytes)).Name()
}

func (s *IntegrationTestSuite) fundAccount(clientCtx client.Context, from, to sdk.AccAddress, coins sdk.Coins) {
	_, err := banktestutil.MsgSendExec(
		clientCtx,
		from,
		to,
		coins, fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) commonTxFlags() []string {
	return []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}
}

var flagOutputJSON = fmt.Sprintf("--%s=json", tmcli.OutputFlag)

func makeFlagFrom(from string) string {
	return fmt.Sprintf("--%s=%s", flags.FlagFrom, from)
}

func (s *IntegrationTestSuite) TestTxCreateClass() {
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx
	fee := core.DefaultParams().CreditClassFee[0]
	feeStr := fee.String()

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
		respCode       uint32
		expectedClass  *core.ClassInfo
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "accepts 4 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde", "abcde", "abcde", "abce"},
			expectErr:      true,
			expectedErrMsg: "accepts 4 arg(s), received 5",
		},
		{
			name:           "missing from flag",
			args:           makeCreateClassArgs([]string{val0.Address.String()}, validCreditTypeAbbrev, validMetadata, feeStr, s.commonTxFlags()...),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name:      "single issuer",
			args:      makeCreateClassArgs([]string{val0.Address.String()}, validCreditTypeAbbrev, validMetadata, feeStr, append(s.commonTxFlags(), makeFlagFrom(val0.Address.String()))...),
			expectErr: false,
			expectedClass: &core.ClassInfo{
				Admin:            val0.Address.String(),
				Metadata:         validMetadata,
				CreditTypeAbbrev: validCreditTypeAbbrev,
			},
		},
		{
			name:      "single issuer with from key-name",
			args:      makeCreateClassArgs([]string{val0.Address.String()}, validCreditTypeAbbrev, validMetadata, feeStr, append(s.commonTxFlags(), makeFlagFrom("node0"))...),
			expectErr: false,
			expectedClass: &core.ClassInfo{
				Admin:            val0.Address.String(),
				Metadata:         validMetadata,
				CreditTypeAbbrev: validCreditTypeAbbrev,
			},
		},
		{
			name: "with amino-json",
			args: makeCreateClassArgs([]string{val0.Address.String()}, validCreditTypeAbbrev, validMetadata, feeStr,
				append(s.commonTxFlags(), makeFlagFrom(val0.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON))...),
			expectErr: false,
			expectedClass: &core.ClassInfo{
				Admin:            val0.Address.String(),
				Metadata:         validMetadata,
				CreditTypeAbbrev: validCreditTypeAbbrev,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := coreclient.TxCreateClassCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.respCode, res.Code, "got %d wanted %d\nresponse: %v", res.Code, tc.respCode, res)
				if tc.respCode == 0 {
					classIdFound := false
					for _, e := range res.Logs[0].Events {
						if e.Type == proto.MessageName(&core.EventCreateClass{}) {
							for _, attr := range e.Attributes {
								if attr.Key == "class_id" {
									classIdFound = true
									classId := strings.Trim(attr.Value, "\"")
									queryCmd := coreclient.QueryClassCmd()
									queryArgs := []string{classId, flagOutputJSON}
									queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
									s.Require().NoError(err, queryOut.String())
									var queryRes core.QueryClassResponse
									s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))

									s.Require().Equal(tc.expectedClass.Admin, queryRes.Class.Admin)
									s.Require().Equal(tc.expectedClass.Metadata, queryRes.Class.Metadata)
									s.Require().Equal(tc.expectedClass.CreditTypeAbbrev, queryRes.Class.CreditTypeAbbrev)
								}
							}
						}
					}
					s.Require().True(classIdFound)
				} else {
					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxCreateBatch() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	fee := core.DefaultParams().CreditClassFee[0]
	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String()},
		Metadata:         validMetadata,
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &fee,
	})
	s.Require().NoError(err)
	projectId, err := s.createProject(clientCtx, &core.MsgCreateProject{
		Issuer:       val.Address.String(),
		ClassId:      classId,
		Metadata:     validMetadata,
		Jurisdiction: "US-OR",
	})
	s.Require().NoError(err)

	// Write some invalid JSON to a file
	invalidJsonFile := testutil.WriteToNewTempFile(s.T(), "{asdljdfklfklksdflk}")

	// Create a valid MsgCreateBatch
	startDate, err := types.ParseDate("start date", "2021-01-01")
	s.Require().NoError(err)
	endDate, err := types.ParseDate("end date", "2021-02-01")
	s.Require().NoError(err)
	msgCreateBatch := core.MsgCreateBatch{
		Issuer:    val.Address.String(),
		ProjectId: projectId,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:              s.network.Validators[1].Address.String(),
				TradableAmount:         "100",
				RetiredAmount:          "0.000001",
				RetirementJurisdiction: "AB",
			},
		},
		Metadata:  validMetadata,
		StartDate: &startDate,
		EndDate:   &endDate,
	}

	validBatchJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	testCases := []struct {
		name            string
		args            []string
		expectErr       bool
		errInTxResponse bool
		expectedErrMsg  string
		expectedBatch   *core.BatchInfo
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"r", "e", "g", "e", "n"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 5",
		},
		{
			name: "invalid json",
			args: append(
				[]string{
					invalidJsonFile.Name(),
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid character",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
					validBatchJson,
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "valid batch",
			args: append(
				[]string{
					validBatchJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedBatch: &core.BatchInfo{
				Issuer: val.Address.String(),
			},
		},
		{
			name: "valid batch with from key-name",
			args: append(
				[]string{
					validBatchJson,
					makeFlagFrom("node0"),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedBatch: &core.BatchInfo{
				Issuer: val.Address.String(),
			},
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					validBatchJson,
					makeFlagFrom(val.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedBatch: &core.BatchInfo{
				Issuer: val.Address.String(),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := coreclient.TxCreateBatchCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				if tc.errInTxResponse {
					var res sdk.TxResponse
					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					s.Require().NotEqual(res.Code, 0)
					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
				} else {
					s.Require().Error(err)
					s.Require().Contains(out.String(), tc.expectedErrMsg)
				}
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				batchDenomFound := false
				for _, e := range res.Logs[0].Events {
					if e.Type == proto.MessageName(&core.EventCreateBatch{}) {
						for _, attr := range e.Attributes {
							if attr.Key == "batch_denom" {
								batchDenomFound = true
								batchDenom := strings.Trim(attr.Value, "\"")

								queryCmd := coreclient.QueryBatchCmd()
								queryArgs := []string{batchDenom, flagOutputJSON}
								queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
								s.Require().NoError(err, queryOut.String())
								var queryRes core.QueryBatchResponse
								s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))
								s.Require().Equal(tc.expectedBatch.Issuer, queryRes.Batch.Issuer)

							}
						}
					}
				}
				s.Require().True(batchDenomFound)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxSend() {
	val0 := s.network.Validators[0]
	val1 := s.network.Validators[1]
	clientCtx := val0.ClientCtx
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, val0.Address.String())

	validCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"4\", retired_amount: \"1\", retirement_jurisdiction: \"AB-CD\"}]", batchDenom)

	testCases := []struct {
		name            string
		args            []string
		expectErr       bool
		errInTxResponse bool
		expectedErrMsg  string
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
			name: "missing from flag",
			args: append(
				[]string{
					val1.Address.String(),
					validCredits,
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "valid credits",
			args: append(
				[]string{
					val1.Address.String(),
					validCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					val1.Address.String(),
					validCredits,
					makeFlagFrom(val0.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := coreclient.TxSendCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				if tc.errInTxResponse {
					var res sdk.TxResponse
					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					s.Require().NotEqual(uint32(0), res.Code)
					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
				} else {
					s.Require().Error(err)
					s.Require().Contains(out.String(), tc.expectedErrMsg)
				}
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxRetire() {
	val0 := s.network.Validators[0]
	valAddrStr := val0.Address.String()
	clientCtx := val0.ClientCtx
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, valAddrStr)

	validCredits := fmt.Sprintf("[{batch_denom: \"%s\", amount: \"5\"}]", batchDenom)

	testCases := []struct {
		name            string
		args            []string
		expectErr       bool
		errInTxResponse bool
		expectedErrMsg  string
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
			name: "missing from flag",
			args: append(
				[]string{
					validCredits,
					"AB-CD 12345",
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "valid credits",
			args: append(
				[]string{
					validCredits,
					"AB-CD 12345",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					validCredits,
					"AB-CD 12345",
					makeFlagFrom(val0.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := coreclient.TxRetireCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				if tc.errInTxResponse {
					var res sdk.TxResponse
					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					s.Require().NotEqual(uint32(0), res.Code)
					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
				} else {
					s.Require().Error(err)
					s.Require().Contains(out.String(), tc.expectedErrMsg)
				}
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxCancel() {
	val0 := s.network.Validators[0]
	valAddrStr := val0.Address.String()
	clientCtx := val0.ClientCtx
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, valAddrStr)

	validCredits := fmt.Sprintf("5 %s", batchDenom)

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
			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"foo", "bar", "bar1"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
					validCredits,
					"reason",
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "valid credits",
			args: append(
				[]string{
					validCredits,
					"reason",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
		{
			name: "with amino-json",
			args: append(
				[]string{
					validCredits,
					"reason",
					makeFlagFrom(val0.Address.String()),
					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Commands may panic, so we need to recover and check the error messages
			defer func() {
				if r := recover(); r != nil {
					s.Require().True(tc.expectErr)
					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
				}
			}()

			cmd := coreclient.TxCancelCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateClassAdmin() {
	// use this classId as to not corrupt other tests
	_, _, a1 := testdata.KeyTestPubAddr()
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx

	fee := core.DefaultParams().CreditClassFee[0]
	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val0.Address.String(),
		Issuers:          []string{val0.Address.String()},
		Metadata:         validMetadata,
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &fee,
	})
	s.Require().NoError(err)

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "invalid request: not enough args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "accepts 2 arg(s), received 0",
		},
		{
			name:      "invalid request: no id",
			args:      []string{"", a1.String()},
			expErr:    true,
			expErrMsg: "class-id is required",
		},
		{
			name:      "invalid request: no admin address",
			args:      append([]string{classId, "", makeFlagFrom(a1.String())}, s.commonTxFlags()...),
			expErr:    true,
			expErrMsg: "new admin address is required",
		},
		{
			name:   "valid request",
			args:   append([]string{classId, a1.String(), makeFlagFrom(val0.Address.String())}, s.commonTxFlags()...),
			expErr: false,
		},
		{
			name:   "valid test: from key-name",
			args:   append([]string{classId, a1.String(), makeFlagFrom("node0")}, s.commonTxFlags()...),
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxUpdateClassAdminCmd()
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query the class info
				query := coreclient.QueryClassCmd()
				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
				s.Require().NoError(err, out.String())
				var res core.QueryClassResponse
				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
				s.Require().NoError(err)

				// check the admin has been changed
				s.Require().Equal(res.Class.Admin, tc.args[1])
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateClassMetadata() {
	newMetaData := "hello"
	_, _, a1 := testdata.KeyTestPubAddr()
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx
	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val0.Address.String(),
		Issuers:          []string{val0.Address.String()},
		Metadata:         validMetadata,
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &core.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "invalid request: not enough args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "accepts 2 arg(s), received 0",
		},
		{
			name:      "invalid request: bad id",
			args:      []string{"", a1.String()},
			expErr:    true,
			expErrMsg: "class-id is required",
		},
		{
			name:      "invalid request: no metadata",
			args:      append([]string{classId, "", makeFlagFrom(a1.String())}, s.commonTxFlags()...),
			expErr:    true,
			expErrMsg: "metadata is required",
		},
		{
			name:   "valid request",
			args:   append([]string{classId, newMetaData, makeFlagFrom(val0.Address.String())}, s.commonTxFlags()...),
			expErr: false,
		},
		{
			name:   "valid test: from key-name",
			args:   append([]string{classId, newMetaData, makeFlagFrom("node0")}, s.commonTxFlags()...),
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxUpdateClassMetadataCmd()
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query the credit class info
				query := coreclient.QueryClassCmd()
				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
				s.Require().NoError(err, out.String())
				var res core.QueryClassResponse
				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
				s.Require().NoError(err)

				// check metadata changed
				s.Require().NoError(err)
				s.Require().Equal(res.Class.Metadata, tc.args[1])
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateIssuers() {
	_, _, a2 := testdata.KeyTestPubAddr()
	_, _, a3 := testdata.KeyTestPubAddr()
	newIssuers := []string{a3.String(), a2.String()}
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx
	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val0.Address.String(),
		Issuers:          []string{val0.Address.String()},
		Metadata:         validMetadata,
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &core.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)

	makeArgs := func(add, remove []string, classId, from string) []string {
		args := []string{classId}
		if len(add) > 0 {
			args = append(args, fmt.Sprintf("--%s=%s", coreclient.FlagAddIssuers, strings.Join(add, ",")))
		}
		if len(remove) > 0 {
			args = append(args, fmt.Sprintf("--%s=%s", coreclient.FlagRemoveIssuers, strings.Join(remove, ",")))
		}
		args = append(args, makeFlagFrom(from))
		return append(args, s.commonTxFlags()...)
	}

	testCases := []struct {
		name       string
		args       []string
		expErr     bool
		expErrMsg  string
		expIssuers []string
	}{
		{
			name:      "invalid request: no id",
			args:      makeArgs(nil, nil, "", val0.Address.String()),
			expErr:    true,
			expErrMsg: "class-id is required",
		},
		{
			name:       "valid add request",
			args:       makeArgs(newIssuers, nil, classId, val0.Address.String()),
			expErr:     false,
			expIssuers: newIssuers,
		},
		{
			name:       "valid remove request",
			args:       makeArgs(nil, newIssuers, classId, val0.Address.String()),
			expErr:     false,
			expIssuers: []string{val0.Address.String()},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxUpdateClassIssuersCmd()
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)

				// query the credit class info
				query := coreclient.QueryClassCmd()
				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
				s.Require().NoError(err, out.String())
				var res core.QueryClassResponse
				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
				s.Require().NoError(err)

				// verify issuers list was changed
				query = coreclient.QueryClassIssuersCmd()
				out, err = cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
				s.Require().NoError(err, out.String())
				var res1 core.QueryClassIssuersResponse
				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res1)
				s.Require().NoError(err)
				s.Require().Subset(res1.Issuers, tc.expIssuers)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxSell() {
	val0 := s.network.Validators[0]
	valAddrStr := val0.Address.String()
	clientCtx := val0.ClientCtx
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, valAddrStr)
	expiration, err := types.ParseDate("expiration", "2024-01-01")
	s.Require().NoError(err)

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
			expErrMsg: "accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "accepts 1 arg(s), received 2",
		},
		{
			name: "valid",
			args: append(
				[]string{
					fmt.Sprintf("[{batch_denom: \"%s\", quantity: \"5\", ask_price: \"100stake\", disable_auto_retire: false}]", batchDenom),
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expErr: false,
			expOrder: &marketplace.SellOrder{
				Seller:            val0.Address,
				Quantity:          "5",
				AskAmount:         "100",
				DisableAutoRetire: false,
				Expiration:        &gogotypes.Timestamp{},
			},
		},
		{
			name: "valid with expiration",
			args: append(
				[]string{
					fmt.Sprintf("[{batch_denom: \"%s\", quantity: \"5\", ask_price: \"100stake\", disable_auto_retire: false, expiration: \"2024-01-01\"}]", batchDenom),
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expErr: false,
			expOrder: &marketplace.SellOrder{
				Seller:            val0.Address,
				Quantity:          "5",
				AskAmount:         "100",
				DisableAutoRetire: false,
				Expiration:        types.ProtobufToGogoTimestamp(timestamppb.New(expiration)),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := marketplaceclient.TxSellCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)
				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().True(len(res.Logs) > 0)

				found := false
				for _, e := range res.Logs[0].Events {
					if e.Type == proto.MessageName(&marketplace.EventSell{}) {
						for _, attr := range e.Attributes {
							if attr.Key == "sell_order_id" {
								found = true
								orderIdStr := strings.Trim(attr.Value, "\"")
								_, err := strconv.ParseUint(orderIdStr, 10, 64)
								s.Require().NoError(err)
								queryCmd := marketplaceclient.QuerySellOrderCmd()
								queryArgs := []string{orderIdStr, flagOutputJSON}
								queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
								s.Require().NoError(err, queryOut.String())
								var queryRes marketplace.QuerySellOrderResponse
								s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))
								s.Require().Equal(queryRes.SellOrder.Quantity, tc.expOrder.Quantity)
								s.Require().Equal(tc.expOrder.DisableAutoRetire, queryRes.SellOrder.DisableAutoRetire)
								s.Require().Equal(tc.expOrder.Expiration, queryRes.SellOrder.Expiration)
								break
							}
							if found {
								break
							}
						}
					}
					if found {
						break
					}
				}
				s.Require().True(found)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateSellOrders() {
	val0 := s.network.Validators[0]
	valAddrStr := val0.Address.String()
	clientCtx := val0.ClientCtx
	askCoin := sdk.NewInt64Coin(sdk.DefaultBondDenom, 10)
	expiration, err := types.ParseDate("expiration", "3020-04-15")
	s.Require().NoError(err)
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, valAddrStr)
	orderIds, err := s.createSellOrder(clientCtx, &marketplace.MsgSell{
		Seller: valAddrStr,
		Orders: []*marketplace.MsgSell_Order{
			{batchDenom, "10", &askCoin, true, &expiration},
		},
	})
	s.Require().NoError(err)
	orderId := orderIds[0]

	makeArgs := func(msg *marketplace.MsgUpdateSellOrders) []string {
		updates := make([]string, len(msg.Updates))
		for i, u := range msg.Updates {
			updates[i] = fmt.Sprintf(`{sell_order_id: %d, new_quantity: %s, new_ask_price: %v, disable_auto_retire: %t, new_expiration: %s}`, u.SellOrderId, u.NewQuantity, u.NewAskPrice, u.DisableAutoRetire, formatTime(u.NewExpiration))
		}
		updatesStr := strings.Join(updates, ",")
		updateArg := fmt.Sprintf(`[%s]`, updatesStr)
		args := []string{updateArg, makeFlagFrom(msg.Seller)}
		return append(args, s.commonTxFlags()...)
	}

	newAsk := sdk.NewInt64Coin(askCoin.Denom, 3)
	newExpiration, err := types.ParseDate("newExpiration", "2049-07-15")
	s.Require().NoError(err)

	gogoNewExpiration, err := gogotypes.TimestampProto(newExpiration)
	s.Require().NoError(err)
	s.Require().NoError(err)
	testCases := []struct {
		name        string
		args        []string
		sellOrderId string
		expErr      bool
		expErrMsg   string
		expOrder    *marketplace.SellOrder
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "accepts 1 arg(s), received 2",
		},
		{
			name: "valid",
			args: makeArgs(&marketplace.MsgUpdateSellOrders{
				Seller: valAddrStr,
				Updates: []*marketplace.MsgUpdateSellOrders_Update{
					{SellOrderId: orderId, NewQuantity: "9.99", NewAskPrice: &newAsk, DisableAutoRetire: false, NewExpiration: &newExpiration},
				},
			}),
			sellOrderId: fmt.Sprintf("%d", orderId),
			expErr:      false,
			expOrder: &marketplace.SellOrder{
				Id:                orderId,
				Seller:            val0.Address,
				Quantity:          "9.99",
				AskAmount:         "3",
				DisableAutoRetire: false,
				Expiration:        gogoNewExpiration,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := marketplaceclient.TxUpdateSellOrdersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code, res)

				// query sell order
				queryCmd := marketplaceclient.QuerySellOrderCmd()
				queryArgs := []string{tc.sellOrderId, flagOutputJSON}
				queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
				s.Require().NoError(err, queryOut.String())
				var queryRes marketplace.QuerySellOrderResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCreateProject() {
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx
	require := s.Require()
	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val0.Address.String(),
		Issuers:          []string{val0.Address.String()},
		Metadata:         validMetadata,
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &core.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)

	makeArgs := func(msg *core.MsgCreateProject) []string {
		args := []string{msg.ClassId, msg.Jurisdiction, msg.Metadata}
		args = append(args, makeFlagFrom(msg.Issuer))
		return append(args, s.commonTxFlags()...)
	}

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			"minimum args",
			[]string{},
			true,
			"accepts 3 arg(s), received 0",
		},
		{
			"too many args",
			[]string{"C01", "foo", "bar", "baz"},
			true,
			"accepts 3 arg(s), received 4",
		},
		{
			"valid tx without project id",
			makeArgs(&core.MsgCreateProject{
				Issuer:       val0.Address.String(),
				ClassId:      classId,
				Metadata:     validMetadata,
				Jurisdiction: "US-OR",
			}),
			false,
			"",
		},
		{
			"valid tx with project id",
			makeArgs(&core.MsgCreateProject{
				Issuer:       val0.Address.String(),
				ClassId:      classId,
				Metadata:     validMetadata,
				Jurisdiction: "US-OR",
			}),
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.TxCreateProject()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(err.Error(), tc.expErrMsg)
			} else {
				require.NoError(err)
				var res sdk.TxResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.Equal(uint32(0), res.Code)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxBuyDirect() {
	val0 := s.network.Validators[0]
	valAddrStr := val0.Address.String()
	clientCtx := val0.ClientCtx

	buyerAcc, _, err := val0.ClientCtx.Keyring.NewMnemonic("buyDirectAcc", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)
	buyerAddr := buyerAcc.GetAddress()

	validAskDenom := sdk.DefaultBondDenom
	askCoin := sdk.NewInt64Coin(validAskDenom, 10)

	s.fundAccount(clientCtx, val0.Address, buyerAddr, sdk.Coins{sdk.NewInt64Coin(validAskDenom, 500)})

	expiration, err := types.ParseDate("expiration", "3020-04-15")
	s.Require().NoError(err)
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, valAddrStr)
	orderIds, err := s.createSellOrder(clientCtx, &marketplace.MsgSell{
		Seller: valAddrStr,
		Orders: []*marketplace.MsgSell_Order{
			{batchDenom, "10", &askCoin, true, &expiration},
			{batchDenom, "10", &askCoin, false, &expiration},
		},
	})
	s.Require().NoError(err)
	orderId := orderIds[0]
	orderIdRetired := orderIds[1]

	makeArgs := func(sellOrderId uint64, qty, bidPrice string, disableAutoRetire bool, from sdk.Address) []string {
		args := []string{
			strconv.FormatUint(sellOrderId, 10),
			qty, bidPrice, fmt.Sprintf("%t", disableAutoRetire),
		}
		if !disableAutoRetire {
			args = append(args, fmt.Sprintf(`--%s=US-OR`, marketplaceclient.FlagRetirementJurisdiction))
		}
		args = append(args, makeFlagFrom(from.String()))
		return append(args, s.commonTxFlags()...)
	}
	type fields struct {
		orderId           uint64
		qty               string
		askPrice          sdk.Coin
		disableAutoRetire bool
		buyerAddr         sdk.AccAddress
	}
	testCases := []struct {
		name      string
		fields    fields
		expErr    bool
		expErrMsg string
	}{
		{
			"valid tx purchase tradable",
			fields{
				orderId:           orderId,
				qty:               "10",
				askPrice:          askCoin,
				disableAutoRetire: true,
				buyerAddr:         buyerAddr,
			},
			false,
			"",
		},
		{
			"valid tx purchase retired",
			fields{
				orderId:           orderIdRetired,
				qty:               "10",
				askPrice:          askCoin,
				disableAutoRetire: false,
				buyerAddr:         buyerAddr,
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			fields := tc.fields
			args := makeArgs(fields.orderId, fields.qty, fields.askPrice.String(), fields.disableAutoRetire, fields.buyerAddr)
			if tc.expErr {
				cmd := marketplaceclient.TxBuyDirect()
				_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				qtyDec, err := math.NewDecFromString(fields.qty)
				s.Require().NoError(err)
				costDec, err := math.NewDecFromString(askCoin.Amount.String())
				s.Require().NoError(err)
				totalCostDec, err := qtyDec.Mul(costDec)
				s.Require().NoError(err)
				totalCostInt := totalCostDec.SdkIntTrim()
				totalCost := sdk.NewCoin(askCoin.Denom, totalCostInt)
				sellerAccBefore := s.getAccountInfo(clientCtx, val0.Address, askCoin.Denom, batchDenom)
				buyerAccBefore := s.getAccountInfo(clientCtx, buyerAddr, askCoin.Denom, batchDenom)

				cmd := marketplaceclient.TxBuyDirect()
				out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
				s.Require().NoError(err)

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code)

				sellerAccAfter := s.getAccountInfo(clientCtx, val0.Address, askCoin.Denom, batchDenom)
				buyerAccAfter := s.getAccountInfo(clientCtx, buyerAddr, askCoin.Denom, batchDenom)
				s.assertMarketBalancesUpdated(sellerAccBefore, sellerAccAfter, buyerAccBefore, buyerAccAfter, math.NewDecFromInt64(10), totalCost, !fields.disableAutoRetire)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxBuyDirectBatch() {
	val0 := s.network.Validators[0]
	valAddrStr := val0.Address.String()
	clientCtx := val0.ClientCtx
	cmd := marketplaceclient.TxBuyDirectBatch()

	validAskDenom := sdk.DefaultBondDenom
	askCoin := sdk.NewInt64Coin(validAskDenom, 10)

	buyerAcc := s.addr
	s.fundAccount(clientCtx, val0.Address, buyerAcc, sdk.Coins{sdk.NewInt64Coin(validAskDenom, 500)})

	expiration, err := types.ParseDate("expiration", "3020-04-15")
	s.Require().NoError(err)
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, valAddrStr)
	orderIds, err := s.createSellOrder(clientCtx, &marketplace.MsgSell{
		Seller: valAddrStr,
		Orders: []*marketplace.MsgSell_Order{
			{batchDenom, "10", &askCoin, true, &expiration},
			{batchDenom, "10", &askCoin, false, &expiration},
		},
	})

	buyOrders := []*marketplace.MsgBuyDirect_Order{
		{SellOrderId: orderIds[0], Quantity: "10", BidPrice: &askCoin, DisableAutoRetire: true},
		{SellOrderId: orderIds[1], Quantity: "10", BidPrice: &askCoin, RetirementJurisdiction: "US-OR"},
	}
	ordersBz, err := json.Marshal(buyOrders)
	s.Require().NoError(err)
	jsonFile := testutil.WriteToNewTempFile(s.T(), string(ordersBz))

	makeArgs := func(fileName, from string) []string {
		args := []string{fileName, makeFlagFrom(from)}
		return append(args, s.commonTxFlags()...)
	}

	testCases := []struct {
		name   string
		args   []string
		errMsg string
	}{
		{
			name:   "too many args",
			args:   []string{"foo", "bar"},
			errMsg: "accepts 1 arg(s), received 2",
		},
		{
			name:   "invalid: file does not exist",
			args:   []string{"monkey.jpeg"},
			errMsg: "no such file or directory",
		},
		{
			name: "valid order",
			args: makeArgs(jsonFile.Name(), buyerAcc.String()),
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			if len(tc.errMsg) != 0 {
				_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
				s.Require().ErrorContains(err, tc.errMsg)
			} else {
				sellerAccBefore := s.getAccountInfo(clientCtx, val0.Address, askCoin.Denom, batchDenom)
				buyerAccBefore := s.getAccountInfo(clientCtx, buyerAcc, askCoin.Denom, batchDenom)

				out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
				s.Require().NoError(err)

				sellerAccAfter := s.getAccountInfo(clientCtx, val0.Address, askCoin.Denom, batchDenom)
				buyerAccAfter := s.getAccountInfo(clientCtx, buyerAcc, askCoin.Denom, batchDenom)

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(uint32(0), res.Code)
				s.assertMarketBalanceBatchUpdated(sellerAccBefore, sellerAccAfter, buyerAccBefore, buyerAccAfter, buyOrders)
			}
		})
	}
}

// assertMarketBalanceBatchUpdated asserts that all accounts involved in a marketplace transaction are updated properly.
// it assumes that both seller/buyer accounts used the same denom for amount sold/bought.
func (s *IntegrationTestSuite) assertMarketBalanceBatchUpdated(sb, sa, bb, ba accountInfo, orders []*marketplace.MsgBuyDirect_Order) {
	tradSold, retSold := math.NewDecFromInt64(0), math.NewDecFromInt64(0)
	cost := sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)
	for _, order := range orders {
		qty, err := math.NewDecFromString(order.Quantity)
		s.Require().NoError(err)
		if order.DisableAutoRetire {
			tradSold, err = tradSold.Add(qty)
			s.Require().NoError(err)
		} else {
			retSold, err = retSold.Add(qty)
			s.Require().NoError(err)
		}
		costDec, err := math.NewDecFromString(order.BidPrice.Amount.String())
		s.Require().NoError(err)
		totalCostDec, err := qty.Mul(costDec)
		s.Require().NoError(err)
		c := sdk.NewCoin(order.BidPrice.Denom, totalCostDec.SdkIntTrim())
		cost = cost.Add(c)
	}

	totalSold, err := tradSold.Add(retSold)
	s.Require().NoError(err)

	// check sellers coins
	expectedSellerGain := sb.coinBal.Add(cost)
	s.Require().Equal(expectedSellerGain, sa.coinBal)

	// check buyers coins
	// we use LT in the buyer case, as some coins go towards fees, so their balance will be LOWER than before - total cost.
	expectedBuyerCost := bb.coinBal.Sub(cost)
	s.Require().True(ba.coinBal.IsLT(expectedBuyerCost))

	// check sellers credits
	expectedEscrowed, err := sb.escrowed.Sub(totalSold)
	s.Require().NoError(err)
	s.Require().Equal(expectedEscrowed.String(), sa.escrowed.String())
	s.Require().Equal(sb.tradable.String(), sa.tradable.String())
	s.Require().Equal(sb.retired.String(), sa.retired.String())

	expectedRetired, err := bb.retired.Add(retSold)
	s.Require().NoError(err)
	expectedTradable, err := bb.tradable.Add(tradSold)
	s.Require().NoError(err)

	s.Require().Equal(bb.escrowed.String(), bb.escrowed.String())
	s.Require().Equal(ba.tradable.String(), expectedTradable.String())
	s.Require().Equal(ba.retired.String(), expectedRetired.String())
}

func (s *IntegrationTestSuite) TestUpdateProjectMetadata() {
	admin := s.network.Validators[0]
	valAddrStr := admin.Address.String()
	clientCtx := admin.ClientCtx
	cmd := coreclient.TxUpdateProjectMetadataCmd()
	_, projectId := s.createClassProject(clientCtx, valAddrStr)

	unauthAddr := s.addr
	s.fundAccount(clientCtx, admin.Address, unauthAddr, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 50)})

	makeArgs := func(projectId, metadata, from string) []string {
		args := []string{projectId, metadata, makeFlagFrom(from)}
		return append(args, s.commonTxFlags()...)
	}

	testCases := []struct {
		name   string
		args   []string
		errMsg string
		errLog string
	}{
		{
			name:   "not enough args",
			errMsg: "accepts 2 arg(s), received 0",
		},
		{
			name:   "too many args",
			args:   []string{"foo", "bar", "baz"},
			errMsg: "accepts 2 arg(s), received 3",
		},
		{
			name:   "invalid: unauthorized",
			args:   makeArgs(projectId, rand.Str(5), unauthAddr.String()),
			errLog: sdkerrors.ErrUnauthorized.Error(),
		},
		{
			name: "valid",
			args: makeArgs(projectId, rand.Str(5), valAddrStr),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if len(tc.errMsg) != 0 {
				s.Require().ErrorContains(err, tc.errMsg)
			} else {
				s.Require().NoError(err)
				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				if len(tc.errLog) != 0 {
					s.Require().NotEqual(uint32(0), res.Code)
					s.Require().Contains(res.RawLog, tc.errLog)
				} else {
					s.Require().Equal(uint32(0), res.Code)
					pId, expectedMetadata := tc.args[0], tc.args[1]
					project := s.getProject(clientCtx, pId)
					s.Require().Equal(expectedMetadata, project.Metadata)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateProjectAdmin() {
	admin := s.network.Validators[0]
	newAdmin := s.network.Validators[1].Address.String()
	clientCtx := admin.ClientCtx
	_, projectId := s.createClassProject(clientCtx, admin.Address.String())
	cmd := coreclient.TxUpdateProjectAdminCmd()
	makeArgs := func(projectId, newAdminAddr, from string) []string {
		args := make([]string, 0, 6)
		args = append(args, projectId, newAdminAddr, makeFlagFrom(from))
		return append(args, s.commonTxFlags()...)
	}

	unauthAddr := s.addr
	s.fundAccount(clientCtx, admin.Address, unauthAddr, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 50)})

	testCases := []struct {
		name          string
		args          []string
		errMsg        string
		errLog        string
		expectedAdmin string
	}{
		{
			name:   "min args",
			args:   []string{},
			errMsg: "accepts 2 arg(s), received 0",
		},
		{
			name:   "max args",
			args:   []string{"foo", "bar", "baz"},
			errMsg: "accepts 2 arg(s), received 3",
		},
		{
			name:   "invalid: unauthorized",
			args:   makeArgs(projectId, admin.Address.String(), unauthAddr.String()),
			errLog: sdkerrors.ErrUnauthorized.Error(),
		},
		{
			name:          "valid update",
			args:          makeArgs(projectId, newAdmin, admin.Address.String()),
			expectedAdmin: newAdmin,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if len(tc.errMsg) != 0 {
				s.Require().ErrorContains(err, tc.errMsg)
			} else {
				s.Require().NoError(err)
				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				if len(tc.errLog) != 0 {
					s.Require().NotEqual(uint32(0), res.Code)
					s.Require().Contains(res.RawLog, tc.errMsg)
				} else {
					s.Require().Equal(uint32(0), res.Code)
					gotProject := s.getProject(clientCtx, projectId)
					s.Require().Equal(tc.expectedAdmin, gotProject.Admin)
				}
			}

		})
	}
}

func (s *IntegrationTestSuite) getProject(ctx client.Context, projectId string) *core.ProjectInfo {
	cmd := coreclient.QueryProjectCmd()
	out, err := cli.ExecTestCLICmd(ctx, cmd, []string{projectId, flagOutputJSON})
	s.Require().NoError(err)
	var res core.QueryProjectResponse
	s.Require().NoError(ctx.Codec.UnmarshalJSON(out.Bytes(), &res))
	return res.Project
}

func (s *IntegrationTestSuite) assertMarketBalancesUpdated(sb, sa, bb, ba accountInfo, qtySold math.Dec, totalCost sdk.Coin, retired bool) {
	// check sellers coins
	expectedSellerGain := sb.coinBal.Add(totalCost)
	s.Require().True(expectedSellerGain.Equal(sa.coinBal))

	// check buyers coins
	// we use LT in the buyer case, as some coins go towards fees, so their balance will be LOWER than before - total cost.
	expectedBuyerCost := bb.coinBal.Sub(totalCost)
	s.Require().True(ba.coinBal.IsLT(expectedBuyerCost))

	// check sellers credits
	expectedEscrowed, err := sb.escrowed.Sub(qtySold)
	s.Require().NoError(err)
	s.Require().True(expectedEscrowed.Equal(sa.escrowed))
	s.Require().True(sb.tradable.Equal(sa.tradable))
	s.Require().True(sb.retired.Equal(sa.retired))

	// check buyers credits
	if retired {
		expectedRetired, err := bb.retired.Add(qtySold)
		s.Require().NoError(err)
		s.Require().True(expectedRetired.Equal(ba.retired))
		s.Require().True(bb.tradable.Equal(bb.tradable))
	} else {
		expectedTradable, err := bb.tradable.Add(qtySold)
		s.Require().NoError(err)
		s.Require().True(expectedTradable.Equal(ba.tradable))
		s.Require().True(bb.retired.Equal(bb.retired))
	}
	s.Require().True(bb.escrowed.Equal(bb.escrowed))
}

type accountInfo struct {
	coinBal                     sdk.Coin
	tradable, retired, escrowed math.Dec
}

func (s *IntegrationTestSuite) getAccountInfo(clientCtx client.Context, addr sdk.AccAddress, bankDenom, batchDenom string) accountInfo {
	a := accountInfo{}
	a.coinBal = s.getBankBalance(clientCtx, addr, bankDenom)
	batchBal := s.getBalance(clientCtx, addr, batchDenom)
	decs, err := utils.GetNonNegativeFixedDecs(6, batchBal.TradableAmount, batchBal.RetiredAmount, batchBal.EscrowedAmount)
	s.Require().NoError(err)
	a.tradable, a.retired, a.escrowed = decs[0], decs[1], decs[2]
	return a
}

func (s *IntegrationTestSuite) getBankBalance(clientCtx client.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	coins := s.getBankBalances(clientCtx, addr)
	return sdk.Coin{
		Denom:  denom,
		Amount: coins.AmountOf(denom),
	}
}

func (s *IntegrationTestSuite) getBankBalances(clientCtx client.Context, addr sdk.AccAddress) sdk.Coins {
	out, err := banktestutil.QueryBalancesExec(clientCtx, addr)
	s.Require().NoError(err)
	var res banktypes.QueryAllBalancesResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	return res.Balances
}

func (s *IntegrationTestSuite) getBalance(clientCtx client.Context, addr sdk.AccAddress, batchDenom string) *core.BatchBalanceInfo {
	cmd := coreclient.QueryBalanceCmd()
	args := []string{batchDenom, addr.String(), flagOutputJSON}
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	var res core.QueryBalanceResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	return res.Balance
}

func (s *IntegrationTestSuite) createClass(clientCtx client.Context, msg *core.MsgCreateClass) (string, error) {
	args := makeCreateClassArgs(msg.Issuers, msg.CreditTypeAbbrev, msg.Metadata, msg.Fee.String(), append(s.commonTxFlags(), makeFlagFrom(msg.Admin))...)
	cmd := coreclient.TxCreateClassCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	var res sdk.TxResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateClass{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "class_id" {
					return strings.Trim(attr.Value, "\""), nil
				}
			}
		}
	}
	return "", fmt.Errorf("class_id not found")
}

func (s *IntegrationTestSuite) createProject(clientCtx client.Context, msg *core.MsgCreateProject) (string, error) {
	cmd := coreclient.TxCreateProject()
	makeCreateProjectArgs := func(msg *core.MsgCreateProject, flags ...string) []string {
		args := []string{msg.ClassId, msg.Jurisdiction, msg.Metadata}
		return append(args, flags...)
	}

	referenceIdFlag := fmt.Sprintf("--reference-id=%s", msg.ReferenceId)
	flags := append(s.commonTxFlags(), makeFlagFrom(msg.Issuer), referenceIdFlag)
	args := makeCreateProjectArgs(msg, flags...)

	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	var res sdk.TxResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateProject{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "project_id" {
					return strings.Trim(attr.Value, "\""), nil
				}
			}
		}
	}
	return "", fmt.Errorf("project_id not found")
}

func (s *IntegrationTestSuite) createBatch(clientCtx client.Context, msg *core.MsgCreateBatch) (string, error) {
	batchJson := s.writeMsgCreateBatchJSON(msg)
	args := []string{batchJson, makeFlagFrom(msg.Issuer)}
	args = append(args, s.commonTxFlags()...)
	cmd := coreclient.TxCreateBatchCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	var res sdk.TxResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateBatch{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "batch_denom" {
					return strings.Trim(attr.Value, "\""), nil
				}
			}
		}
	}
	return "", fmt.Errorf("could not find batch_denom")
}

func (s *IntegrationTestSuite) createSellOrder(clientCtx client.Context, msg *marketplace.MsgSell) ([]uint64, error) {
	cmd := marketplaceclient.TxSellCmd()

	// order format closure
	formatOrder := func(o *marketplace.MsgSell_Order) string {
		return fmt.Sprintf(`{batch_denom: %s, quantity: %s, ask_price: %v, disable_auto_retire: %t, expiration: %s}`,
			o.BatchDenom, o.Quantity, o.AskPrice, o.DisableAutoRetire, formatTime(o.Expiration))
	}

	// go through all orders and format them
	orders := make([]string, len(msg.Orders))
	for i, o := range msg.Orders {
		orders[i] = formatOrder(o)
	}

	// merge args
	ordersStr := strings.Join(orders, ",")
	orderArg := fmt.Sprintf(`[%s]`, ordersStr)
	args := []string{orderArg, makeFlagFrom(msg.Seller)}
	args = append(args, s.commonTxFlags()...)

	// execute command
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)

	// extract order id's via response output
	var res sdk.TxResponse
	s.Require().Equal(uint32(0), res.Code)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	s.Require().True(len(res.Logs) > 0)
	orderIds := make([]uint64, 0, len(msg.Orders))
	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&marketplace.EventSell{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "sell_order_id" {
					orderId, err := strconv.ParseUint(strings.Trim(attr.Value, "\""), 10, 64)
					s.Require().NoError(err)
					orderIds = append(orderIds, orderId)
				}
			}
		}
	}
	if len(orderIds) == 0 {
		return nil, fmt.Errorf("no order ids found")
	}
	return orderIds, nil
}

func formatTime(t *time.Time) string {
	var monthStr string
	m := t.Month()
	if m < 10 {
		monthStr = fmt.Sprintf("0%d", m)
	} else {
		monthStr = fmt.Sprintf("%d", m)
	}
	return fmt.Sprintf("%d-%s-%d", t.Year(), monthStr, t.Day())
}

// createClassProject creates a class and project, returning their IDs in that order.
func (s *IntegrationTestSuite) createClassProject(clientCtx client.Context, addr string) (classId, projectId string) {
	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            addr,
		Issuers:          []string{addr},
		Metadata:         validMetadata,
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &core.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)

	projectId, err = s.createProject(clientCtx, &core.MsgCreateProject{
		Issuer:       addr,
		ClassId:      classId,
		Metadata:     validMetadata,
		Jurisdiction: "US-OR",
		ReferenceId:  s.projectReferenceId,
	})
	s.Require().NoError(err)

	return classId, projectId
}

// createClassProjectBatch creates a class, project, and batch, returning their IDs in that order.
func (s *IntegrationTestSuite) createClassProjectBatch(clientCtx client.Context, addr string) (classId, projectId, batchDenom string) {
	classId, projectId = s.createClassProject(clientCtx, addr)
	start, end := time.Now(), time.Now()
	var err error
	batchDenom, err = s.createBatch(clientCtx, &core.MsgCreateBatch{
		Issuer:    addr,
		ProjectId: projectId,
		Issuance: []*core.BatchIssuance{
			{Recipient: addr, TradableAmount: "999999999999999999", RetiredAmount: "100000000000", RetirementJurisdiction: "US-OR"},
		},
		Metadata:  validMetadata,
		StartDate: &start,
		EndDate:   &end,
		Open:      false,
		OriginTx:  nil,
		Note:      "",
	})
	s.Require().NoError(err)
	return
}

func makeCreateClassArgs(issuers []string, ctAbbrev, metadata, fee string, flags ...string) []string {
	var issuersStr string
	if len(issuers) == 1 {
		issuersStr = issuers[0]
	} else if len(issuers) > 1 {
		issuersStr = strings.Join(
			issuers,
			",",
		)
	}
	args := []string{
		issuersStr,
		ctAbbrev,
		metadata,
		fee,
	}
	args = append(args, flags...)
	return args
}
