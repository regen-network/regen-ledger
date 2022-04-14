package testsuite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
	marketplaceclient "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	addr sdk.AccAddress
}

const (
	validCreditTypeAbbrev = "C"
	validMetadata         = "metadata"
)

func RunCLITests(t *testing.T, cfg network.Config) {
	suite.Run(t, NewIntegrationTestSuite(cfg))
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

// Write a MsgCreateBatch to a new temporary file and return the filename
func (s *IntegrationTestSuite) writeMsgCreateBatchJSON(msg *core.MsgCreateBatch) string {
	bytes, err := s.network.Validators[0].ClientCtx.Codec.MarshalJSON(msg)
	s.Require().NoError(err)

	return testutil.WriteToNewTempFile(s.T(), string(bytes)).Name()
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	val := s.network.Validators[0]

	// create an account for val
	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewValidator0", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	_, a1pub, a1 := testdata.KeyTestPubAddr()
	_, err = val.ClientCtx.Keyring.SavePubKey("throwaway", a1pub, hd.Secp256k1Type)
	s.Require().NoError(err)

	// fund the test account
	account := sdk.AccAddress(info.GetPubKey().Address())
	for _, acc := range []sdk.AccAddress{account, a1} {
		_, err = banktestutil.MsgSendExec(
			val.ClientCtx,
			val.Address,
			acc,
			sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(20000000000000000))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		)
		s.Require().NoError(err)
	}

	s.addr = account
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
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

func makeCreateProjectArgs(msg *core.MsgCreateProject, flags ...string) []string {
	args := []string{msg.ClassId, msg.ProjectLocation, msg.Metadata, fmt.Sprintf("--%s=%s", coreclient.FlagProjectId, msg.ProjectId)}
	return append(args, flags...)
}

//func (s *IntegrationTestSuite) TestTxCreateClass() {
//	val0 := s.network.Validators[0]
//	val1 := s.network.Validators[1]
//	clientCtx := val0.ClientCtx
//	fee := core.DefaultParams().CreditClassFee[0]
//	feeStr := fee.String()
//
//	makeArgs := func(issuers []string, ctAbbrev, metadata, fee string, flags ...string) []string {
//		var issuersStr string
//		if len(issuers) == 1 {
//			issuersStr = issuers[0]
//		} else if len(issuers) > 1 {
//			issuersStr = strings.Join(
//				issuers,
//				",",
//			)
//		}
//		args := []string{
//			issuersStr,
//			ctAbbrev,
//			metadata,
//			fee,
//		}
//		args = append(args, flags...)
//		fmt.Println(args)
//		return args
//	}
//
//	testCases := []struct {
//		name              string
//		args              []string
//		expectErr         bool
//		expectedErrMsg    string
//		respCode          uint32
//		expectedClassInfo *core.ClassInfo
//	}{
//		//{
//		//	name:           "missing args",
//		//	args:           []string{},
//		//	expectErr:      true,
//		//	expectedErrMsg: "accepts 4 arg(s), received 0",
//		//},
//		//{
//		//	name:           "too many args",
//		//	args:           []string{"abcde", "abcde", "abcde", "abcde", "abce"},
//		//	expectErr:      true,
//		//	expectedErrMsg: "accepts 4 arg(s), received 5",
//		//},
//		//{
//		//	name: "invalid issuer",
//		//	args: append(
//		//		[]string{
//		//			"abcde",
//		//			validCreditTypeAbbrev,
//		//			validMetadata,
//		//			feeStr,
//		//			makeFlagFrom(val0.Address.String()),
//		//		},
//		//		s.commonTxFlags()...,
//		//	),
//		//	expectErr:      true,
//		//	expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
//		//},
//		//{
//		//	name:           "missing from flag",
//		//	args:           makeArgs([]string{val0.Address.String()}, validCreditTypeAbbrev, validMetadata, feeStr, s.commonTxFlags()...),
//		//	expectErr:      true,
//		//	expectedErrMsg: "required flag(s) \"from\" not set",
//		//},
//		//{
//		//	name:           "invalid credit type",
//		//	args:           makeArgs([]string{val0.Address.String()}, "caarbon", validMetadata, feeStr, append(s.commonTxFlags(), makeFlagFrom(val0.Address.String()))...),
//		//	expectErr:      false,
//		//	expectedErrMsg: "caarbon is not a valid credit type",
//		//	respCode:       29,
//		//},
//		{
//			name:      "single issuer",
//			args:      makeArgs([]string{val0.Address.String()}, validCreditTypeAbbrev, validMetadata, feeStr, append(s.commonTxFlags(), makeFlagFrom(val0.Address.String()))...),
//			expectErr: false,
//			expectedClassInfo: &core.ClassInfo{
//				Admin:      val0.Address,
//				Metadata:   validMetadata,
//				CreditType: validCreditTypeAbbrev,
//			},
//		},
//		{
//			name:      "single issuer with from key-name",
//			args:      makeArgs([]string{val0.Address.String()}, validCreditTypeAbbrev, validMetadata, feeStr, append(s.commonTxFlags(), makeFlagFrom("node0"))...),
//			expectErr: false,
//			expectedClassInfo: &core.ClassInfo{
//				Admin:      val0.Address,
//				Metadata:   validMetadata,
//				CreditType: validCreditTypeAbbrev,
//			},
//		},
//		{
//			name: "multiple issuers",
//			args: makeArgs([]string{val0.Address.String(), val1.Address.String()}, validCreditTypeAbbrev,
//				validMetadata, feeStr, append(s.commonTxFlags(), makeFlagFrom(val0.Address.String()))...),
//			expectErr: false,
//			expectedClassInfo: &core.ClassInfo{
//				Admin:      val0.Address,
//				Metadata:   validMetadata,
//				CreditType: validCreditTypeAbbrev,
//			},
//		},
//		{
//			name: "with amino-json",
//			args: makeArgs([]string{val0.Address.String()}, validCreditTypeAbbrev, validMetadata, feeStr,
//				append(s.commonTxFlags(), makeFlagFrom(val0.Address.String()),
//					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON))...),
//			expectErr: false,
//			expectedClassInfo: &core.ClassInfo{
//				Admin:      val0.Address,
//				Metadata:   validMetadata,
//				CreditType: validCreditTypeAbbrev,
//			},
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			// Commands may panic, so we need to recover and check the error messages
//			defer func() {
//				if r := recover(); r != nil {
//					s.Require().True(tc.expectErr)
//					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
//				}
//			}()
//
//			cmd := coreclient.TxCreateClassCmd()
//			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expectErr {
//				s.Require().Error(err)
//				s.Require().Contains(out.String(), tc.expectedErrMsg)
//			} else {
//				s.Require().NoError(err, out.String())
//
//				var res sdk.TxResponse
//				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//				s.Require().Equal(tc.respCode, res.Code, "got %d wanted %d\nresponse: %v", res.Code, tc.respCode, res)
//				if tc.respCode == 0 {
//					classIdFound := false
//					for _, e := range res.Logs[0].Events {
//						if e.Type == proto.MessageName(&core.EventCreateClass{}) {
//							for _, attr := range e.Attributes {
//								if attr.Key == "class_id" {
//									classIdFound = true
//									classId := strings.Trim(attr.Value, "\"")
//									queryCmd := coreclient.QueryClassInfoCmd()
//									queryArgs := []string{classId, flagOutputJSON}
//									queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
//									s.Require().NoError(err, queryOut.String())
//									var queryRes core.QueryClassInfoResponse
//									s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))
//
//									s.Require().Equal(tc.expectedClassInfo.Admin, queryRes.Info.Admin)
//									s.Require().Equal(tc.expectedClassInfo.Metadata, queryRes.Info.Metadata)
//									s.Require().Equal(tc.expectedClassInfo.CreditType, queryRes.Info.CreditType)
//								}
//							}
//						}
//					}
//					s.Require().True(classIdFound)
//				} else {
//					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
//				}
//			}
//		})
//	}
//}

//func (s *IntegrationTestSuite) TestTxCreateBatch() {
//	val := s.network.Validators[0]
//	clientCtx := val.ClientCtx
//	fee := core.DefaultParams().CreditClassFee[0]
//	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
//		Admin:            val.Address.String(),
//		Issuers:          []string{val.Address.String()},
//		Metadata:         "META",
//		CreditTypeAbbrev: "C",
//		Fee:              &fee,
//	})
//	s.Require().NoError(err)
//	projectId, err := s.createProject(clientCtx, &core.MsgCreateProject{
//		Issuer:          val.Address.String(),
//		ClassId:         classId,
//		Metadata:        "META2",
//		ProjectLocation: "US-OR",
//		ProjectId:       "FBI",
//	})
//	s.Require().NoError(err)
//
//	// Write some invalid JSON to a file
//	invalidJsonFile := testutil.WriteToNewTempFile(s.T(), "{asdljdfklfklksdflk}")
//
//	// Create a valid MsgCreateBatch
//	startDate, err := types.ParseDate("start date", "2021-01-01")
//	s.Require().NoError(err)
//	endDate, err := types.ParseDate("end date", "2021-02-01")
//	s.Require().NoError(err)
//	msgCreateBatch := core.MsgCreateBatch{
//		Issuer:    val.Address.String(),
//		ProjectId: projectId,
//		Issuance: []*core.BatchIssuance{
//			{
//				Recipient:          s.network.Validators[1].Address.String(),
//				TradableAmount:     "100",
//				RetiredAmount:      "0.000001",
//				RetirementLocation: "AB",
//			},
//		},
//		Metadata:  validMetadata,
//		StartDate: &startDate,
//		EndDate:   &endDate,
//	}
//
//	validBatchJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)
//
//	// Write batch with invalid project
//	msgCreateBatch.ProjectId = "abcde-"
//	invalidProjectIdJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)
//
//	// Write batch with missing start date
//	msgCreateBatch.ProjectId = projectId
//	msgCreateBatch.StartDate = nil
//	missingStartDateJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)
//
//	// Write batch with missing end date
//	msgCreateBatch.StartDate = &startDate
//	msgCreateBatch.EndDate = nil
//	missingEndDateJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)
//
//	// Write batch with invalid issuance recipient
//	msgCreateBatch.Issuance[0].Recipient = "abcde"
//	msgCreateBatch.EndDate = &endDate
//	invalidRecipientJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)
//
//	// Write batch with invalid issuance tradable amount
//	msgCreateBatch.Issuance[0].Recipient = s.network.Validators[1].Address.String()
//	msgCreateBatch.Issuance[0].TradableAmount = "abcde"
//	invalidTradableAmountJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)
//
//	// Write batch with invalid issuance retired amount
//	msgCreateBatch.Issuance[0].TradableAmount = "100"
//	msgCreateBatch.Issuance[0].RetiredAmount = "abcde"
//	invalidRetiredAmountJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)
//
//	// Write batch with invalid issuance retirement location
//	msgCreateBatch.Issuance[0].RetiredAmount = "0.000001"
//	msgCreateBatch.Issuance[0].RetirementLocation = "abcde"
//	invalidRetirementLocationJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)
//
//	testCases := []struct {
//		name              string
//		args              []string
//		expectErr         bool
//		errInTxResponse   bool
//		expectedErrMsg    string
//		expectedBatchInfo *core.BatchInfo
//	}{
//		{
//			name:           "missing args",
//			args:           []string{},
//			expectErr:      true,
//			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
//		},
//		{
//			name:           "too many args",
//			args:           []string{"r", "e", "g", "e", "n"},
//			expectErr:      true,
//			expectedErrMsg: "Error: accepts 1 arg(s), received 5",
//		},
//		{
//			name: "invalid json",
//			args: append(
//				[]string{
//					invalidJsonFile.Name(),
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid character",
//		},
//		{
//			name: "invalid project id",
//			args: append(
//				[]string{
//					invalidProjectIdJson,
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:       true,
//			errInTxResponse: false,
//			expectedErrMsg:  "invalid project id",
//		},
//		{
//			name: "missing start date",
//			args: append(
//				[]string{
//					missingStartDateJson,
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "must provide a start date for the credit batch: invalid request",
//		},
//		{
//			name: "missing end date",
//			args: append(
//				[]string{
//					missingEndDateJson,
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "must provide an end date for the credit batch: invalid request",
//		},
//		{
//			name: "invalid issuance recipient",
//			args: append(
//				[]string{
//					invalidRecipientJson,
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: sdkerrors.ErrInvalidAddress.Error(),
//		},
//		{
//			name: "invalid issuance tradable amount",
//			args: append(
//				[]string{
//					invalidTradableAmountJson,
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid decimal string",
//		},
//		{
//			name: "invalid issuance retired amount",
//			args: append(
//				[]string{
//					invalidRetiredAmountJson,
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid decimal string",
//		},
//		{
//			name: "invalid issuance retirement location",
//			args: append(
//				[]string{
//					invalidRetirementLocationJson,
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "Invalid location: abcde",
//		},
//		{
//			name: "missing from flag",
//			args: append(
//				[]string{
//					validBatchJson,
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "required flag(s) \"from\" not set",
//		},
//		{
//			name: "valid batch",
//			args: append(
//				[]string{
//					validBatchJson,
//					makeFlagFrom(val.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//			expectedBatchInfo: &core.BatchInfo{
//				Issuer: val.Address,
//			},
//		},
//		{
//			name: "valid batch with from key-name",
//			args: append(
//				[]string{
//					validBatchJson,
//					makeFlagFrom("node0"),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//			expectedBatchInfo: &core.BatchInfo{
//				Issuer: val.Address,
//			},
//		},
//		{
//			name: "with amino-json",
//			args: append(
//				[]string{
//					validBatchJson,
//					makeFlagFrom(val.Address.String()),
//					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//			expectedBatchInfo: &core.BatchInfo{
//				Issuer: val.Address,
//			},
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			// Commands may panic, so we need to recover and check the error messages
//			defer func() {
//				if r := recover(); r != nil {
//					s.Require().True(tc.expectErr)
//					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
//				}
//			}()
//
//			cmd := coreclient.TxCreateBatchCmd()
//			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expectErr {
//				if tc.errInTxResponse {
//					var res sdk.TxResponse
//					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//					s.Require().NotEqual(res.Code, 0)
//					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
//				} else {
//					s.Require().Error(err)
//					s.Require().Contains(out.String(), tc.expectedErrMsg)
//				}
//			} else {
//				s.Require().NoError(err, out.String())
//
//				var res sdk.TxResponse
//				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//
//				batchDenomFound := false
//				for _, e := range res.Logs[0].Events {
//					if e.Type == proto.MessageName(&core.EventCreateBatch{}) {
//						for _, attr := range e.Attributes {
//							if attr.Key == "batch_denom" {
//								batchDenomFound = true
//								batchDenom := strings.Trim(attr.Value, "\"")
//
//								queryCmd := coreclient.QueryBatchInfoCmd()
//								queryArgs := []string{batchDenom, flagOutputJSON}
//								queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
//								s.Require().NoError(err, queryOut.String())
//								var queryRes core.QueryBatchInfoResponse
//								s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))
//								s.Require().Equal(tc.expectedBatchInfo.Issuer, queryRes.Info.Issuer)
//
//							}
//						}
//					}
//				}
//				s.Require().True(batchDenomFound)
//			}
//		})
//	}
//}

//func (s *IntegrationTestSuite) TestTxSend() {
//	val0 := s.network.Validators[0]
//	val1 := s.network.Validators[1]
//	clientCtx := val0.ClientCtx
//	fee := core.DefaultParams().CreditClassFee[0]
//	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
//		Admin:            val0.Address.String(),
//		Issuers:          []string{val0.Address.String()},
//		Metadata:         "META",
//		CreditTypeAbbrev: validCreditTypeAbbrev,
//		Fee:              &fee,
//	})
//	s.Require().NoError(err)
//	projectId, err := s.createProject(clientCtx, &core.MsgCreateProject{
//		Issuer:          val0.Address.String(),
//		ClassId:         classId,
//		Metadata:        "META",
//		ProjectLocation: "US-OR",
//		ProjectId:       rand.Str(3),
//	})
//	s.Require().NoError(err)
//	start := time.Now()
//	end := time.Now()
//	batchDenom, err := s.createBatch(clientCtx, &core.MsgCreateBatch{
//		Issuer:    val0.Address.String(),
//		ProjectId: projectId,
//		Issuance: []*core.BatchIssuance{
//			{Recipient: val0.Address.String(), TradableAmount: "900000000000000000"},
//		},
//		Metadata:  "",
//		StartDate: &start,
//		EndDate:   &end,
//		Open:      false,
//		OriginTx:  nil,
//		Note:      "",
//	})
//	s.Require().NoError(err)
//
//	validCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"4\", retired_amount: \"1\", retirement_location: \"AB-CD\"}]", batchDenom)
//	invalidBatchDenomCredits := "[{batch_denom: abcde, tradable_amount: \"4\", retired_amount: \"1\", retirement_location: \"AB-CD\"}]"
//	invalidTradableAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"abcde\", retired_amount: \"1\", retirement_location: \"AB-CD\"}]", batchDenom)
//	invalidRetiredAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"4\", retired_amount: \"abcde\", retirement_location: \"AB-CD\"}]", batchDenom)
//	invalidRetirementLocationCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"4\", retired_amount: \"1\", retirement_location: \"abcde\"}]", batchDenom)
//
//	testCases := []struct {
//		name            string
//		args            []string
//		expectErr       bool
//		errInTxResponse bool
//		expectedErrMsg  string
//	}{
//		{
//			name:           "missing args",
//			args:           []string{},
//			expectErr:      true,
//			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
//		},
//		{
//			name:           "too many args",
//			args:           []string{"abcde", "abcde", "abcde"},
//			expectErr:      true,
//			expectedErrMsg: "Error: accepts 2 arg(s), received 3",
//		},
//		{
//			name: "invalid recipient",
//			args: append(
//				[]string{
//					"abcde",
//					validCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
//		},
//		{
//			name: "invalid batch denom",
//			args: append(
//				[]string{
//					val1.Address.String(),
//					invalidBatchDenomCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid denom",
//		},
//		{
//			name: "invalid tradable amount",
//			args: append(
//				[]string{
//					val1.Address.String(),
//					invalidTradableAmountCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid decimal string",
//		},
//		{
//			name: "invalid retired amount",
//			args: append(
//				[]string{
//					val1.Address.String(),
//					invalidRetiredAmountCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid decimal string",
//		},
//		{
//			name: "invalid retirement location",
//			args: append(
//				[]string{
//					val1.Address.String(),
//					invalidRetirementLocationCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "Invalid location: abcde",
//		},
//		{
//			name: "missing from flag",
//			args: append(
//				[]string{
//					val1.Address.String(),
//					validCredits,
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "required flag(s) \"from\" not set",
//		},
//		{
//			name: "valid credits",
//			args: append(
//				[]string{
//					val1.Address.String(),
//					validCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//		},
//		{
//			name: "with amino-json",
//			args: append(
//				[]string{
//					val1.Address.String(),
//					validCredits,
//					makeFlagFrom(val0.Address.String()),
//					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			// Commands may panic, so we need to recover and check the error messages
//			defer func() {
//				if r := recover(); r != nil {
//					s.Require().True(tc.expectErr)
//					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
//				}
//			}()
//
//			cmd := coreclient.TxSendCmd()
//			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expectErr {
//				if tc.errInTxResponse {
//					var res sdk.TxResponse
//					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//					s.Require().NotEqual(uint32(0), res.Code)
//					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
//				} else {
//					s.Require().Error(err)
//					s.Require().Contains(out.String(), tc.expectedErrMsg)
//				}
//			} else {
//				s.Require().NoError(err, out.String())
//
//				var res sdk.TxResponse
//				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//				s.Require().Equal(uint32(0), res.Code)
//			}
//		})
//	}
//}

//func (s *IntegrationTestSuite) TestTxRetire() {
//	val0 := s.network.Validators[0]
//	valAddrStr := val0.Address.String()
//	clientCtx := val0.ClientCtx
//	fee := core.DefaultParams().CreditClassFee[0]
//	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
//		Admin:            valAddrStr,
//		Issuers:          []string{valAddrStr},
//		Metadata:         "META",
//		CreditTypeAbbrev: validCreditTypeAbbrev,
//		Fee:              &fee,
//	})
//	s.Require().NoError(err)
//	projectId, err := s.createProject(clientCtx, &core.MsgCreateProject{
//		Issuer:          valAddrStr,
//		ClassId:         classId,
//		Metadata:        "META",
//		ProjectLocation: "US-OR",
//		ProjectId:       rand.Str(3),
//	})
//	s.Require().NoError(err)
//	start, end := time.Now(), time.Now()
//	batchDenom, err := s.createBatch(clientCtx, &core.MsgCreateBatch{
//		Issuer:    valAddrStr,
//		ProjectId: projectId,
//		Issuance: []*core.BatchIssuance{
//			{Recipient: valAddrStr, TradableAmount: "90000000000000000"},
//		},
//		Metadata:  "META",
//		StartDate: &start,
//		EndDate:   &end,
//	})
//	s.Require().NoError(err)
//
//	validCredits := fmt.Sprintf("[{batch_denom: \"%s\", amount: \"5\"}]", batchDenom)
//	invalidBatchDenomCredits := "[{batch_denom: abcde, amount: \"5\"}]"
//	invalidAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", amount: \"abcde\"}]", batchDenom)
//
//	testCases := []struct {
//		name            string
//		args            []string
//		expectErr       bool
//		errInTxResponse bool
//		expectedErrMsg  string
//	}{
//		{
//			name:           "missing args",
//			args:           []string{},
//			expectErr:      true,
//			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
//		},
//		{
//			name:           "too many args",
//			args:           []string{"abcde", "abcde", "abcde"},
//			expectErr:      true,
//			expectedErrMsg: "Error: accepts 2 arg(s), received 3",
//		},
//		{
//			name: "invalid batch denom",
//			args: append(
//				[]string{
//					invalidBatchDenomCredits,
//					"AB-CD 12345",
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid denom",
//		},
//		{
//			name: "invalid amount",
//			args: append(
//				[]string{
//					invalidAmountCredits,
//					"AB-CD 12345",
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid decimal string",
//		},
//		{
//			name: "invalid retirement location",
//			args: append(
//				[]string{
//					validCredits,
//					"abcde",
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "Invalid location: abcde",
//		},
//		{
//			name: "missing from flag",
//			args: append(
//				[]string{
//					validCredits,
//					"AB-CD 12345",
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "required flag(s) \"from\" not set",
//		},
//		{
//			name: "valid credits",
//			args: append(
//				[]string{
//					validCredits,
//					"AB-CD 12345",
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//		},
//		{
//			name: "with amino-json",
//			args: append(
//				[]string{
//					validCredits,
//					"AB-CD 12345",
//					makeFlagFrom(val0.Address.String()),
//					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			// Commands may panic, so we need to recover and check the error messages
//			defer func() {
//				if r := recover(); r != nil {
//					s.Require().True(tc.expectErr)
//					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
//				}
//			}()
//
//			cmd := coreclient.TxRetireCmd()
//			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expectErr {
//				if tc.errInTxResponse {
//					var res sdk.TxResponse
//					s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//					s.Require().NotEqual(uint32(0), res.Code)
//					s.Require().Contains(res.RawLog, tc.expectedErrMsg)
//				} else {
//					s.Require().Error(err)
//					s.Require().Contains(out.String(), tc.expectedErrMsg)
//				}
//			} else {
//				s.Require().NoError(err, out.String())
//
//				var res sdk.TxResponse
//				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//				s.Require().Equal(uint32(0), res.Code)
//			}
//		})
//	}
//}

//func (s *IntegrationTestSuite) TestTxCancel() {
//	val0 := s.network.Validators[0]
//	valAddrStr := val0.Address.String()
//	clientCtx := val0.ClientCtx
//	fee := core.DefaultParams().CreditClassFee[0]
//	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
//		Admin:            valAddrStr,
//		Issuers:          []string{valAddrStr},
//		Metadata:         "META",
//		CreditTypeAbbrev: validCreditTypeAbbrev,
//		Fee:              &fee,
//	})
//	s.Require().NoError(err)
//	projectId, err := s.createProject(clientCtx, &core.MsgCreateProject{
//		Issuer:          valAddrStr,
//		ClassId:         classId,
//		Metadata:        "META",
//		ProjectLocation: "US-OR",
//		ProjectId:       rand.Str(3),
//	})
//	s.Require().NoError(err)
//	start, end := time.Now(), time.Now()
//	batchDenom, err := s.createBatch(clientCtx, &core.MsgCreateBatch{
//		Issuer:    valAddrStr,
//		ProjectId: projectId,
//		Issuance: []*core.BatchIssuance{
//			{Recipient: valAddrStr, TradableAmount: "90000000000000000"},
//		},
//		Metadata:  "META",
//		StartDate: &start,
//		EndDate:   &end,
//	})
//	s.Require().NoError(err)
//
//	validCredits := fmt.Sprintf("5 %s", batchDenom)
//	invalidBatchDenomCredits := "5 abcde"
//	invalidAmountCredits := fmt.Sprintf("abcde %s", batchDenom)
//
//	testCases := []struct {
//		name           string
//		args           []string
//		expectErr      bool
//		expectedErrMsg string
//	}{
//		{
//			name:           "missing args",
//			args:           []string{},
//			expectErr:      true,
//			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
//		},
//		{
//			name: "invalid batch denom",
//			args: append(
//				[]string{
//					invalidBatchDenomCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid credit expression",
//		},
//		{
//			name: "invalid amount",
//			args: append(
//				[]string{
//					invalidAmountCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "invalid credit expression",
//		},
//		{
//			name: "missing from flag",
//			args: append(
//				[]string{
//					validCredits,
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr:      true,
//			expectedErrMsg: "required flag(s) \"from\" not set",
//		},
//		{
//			name: "valid credits",
//			args: append(
//				[]string{
//					validCredits,
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//		},
//		{
//			name: "with amino-json",
//			args: append(
//				[]string{
//					validCredits,
//					makeFlagFrom(val0.Address.String()),
//					fmt.Sprintf("--%s=%s", flags.FlagSignMode, flags.SignModeLegacyAminoJSON),
//				},
//				s.commonTxFlags()...,
//			),
//			expectErr: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			// Commands may panic, so we need to recover and check the error messages
//			defer func() {
//				if r := recover(); r != nil {
//					s.Require().True(tc.expectErr)
//					s.Require().Contains(r.(error).Error(), tc.expectedErrMsg)
//				}
//			}()
//
//			cmd := coreclient.TxCancelCmd()
//			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expectErr {
//				s.Require().Error(err)
//				s.Require().Contains(out.String(), tc.expectedErrMsg)
//			} else {
//				s.Require().NoError(err, out.String())
//
//				var res sdk.TxResponse
//				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//				s.Require().Equal(uint32(0), res.Code)
//			}
//		})
//	}
//}

//func (s *IntegrationTestSuite) TestTxUpdateClassAdmin() {
//	// use this classId as to not corrupt other tests
//	_, _, a1 := testdata.KeyTestPubAddr()
//	val0 := s.network.Validators[0]
//	clientCtx := val0.ClientCtx
//
//	fee := core.DefaultParams().CreditClassFee[0]
//	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
//		Admin:            val0.Address.String(),
//		Issuers:          []string{val0.Address.String()},
//		Metadata:         "META",
//		CreditTypeAbbrev: validCreditTypeAbbrev,
//		Fee:              &fee,
//	})
//	s.Require().NoError(err)
//
//	testCases := []struct {
//		name      string
//		args      []string
//		expErr    bool
//		expErrMsg string
//	}{
//		{
//			name:      "invalid request: not enough args",
//			args:      []string{},
//			expErr:    true,
//			expErrMsg: "accepts 2 arg(s), received 0",
//		},
//		{
//			name:      "invalid request: no id",
//			args:      []string{"", a1.String()},
//			expErr:    true,
//			expErrMsg: "class-id is required",
//		},
//		{
//			name:      "invalid request: no admin address",
//			args:      append([]string{classId, "", makeFlagFrom(a1.String())}, s.commonTxFlags()...),
//			expErr:    true,
//			expErrMsg: "new admin address is required",
//		},
//		{
//			name:   "valid request",
//			args:   append([]string{classId, a1.String(), makeFlagFrom(val0.Address.String())}, s.commonTxFlags()...),
//			expErr: false,
//		},
//		{
//			name:   "valid test: from key-name",
//			args:   append([]string{classId, a1.String(), makeFlagFrom("node0")}, s.commonTxFlags()...),
//			expErr: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			cmd := coreclient.TxUpdateClassAdminCmd()
//			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expErr {
//				s.Require().Error(err)
//			} else {
//				s.Require().NoError(err)
//
//				// query the class info
//				query := coreclient.QueryClassInfoCmd()
//				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
//				s.Require().NoError(err, out.String())
//				var res core.QueryClassInfoResponse
//				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
//				s.Require().NoError(err)
//
//				// check the admin has been changed
//				s.Require().Equal(sdk.AccAddress(res.Info.Admin).String(), tc.args[1])
//			}
//		})
//	}
//}

//func (s *IntegrationTestSuite) TestTxUpdateClassMetadata() {
//	newMetaData := "hello"
//	_, _, a1 := testdata.KeyTestPubAddr()
//	val0 := s.network.Validators[0]
//	clientCtx := val0.ClientCtx
//	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
//		Admin:            val0.Address.String(),
//		Issuers:          []string{val0.Address.String()},
//		Metadata:         "META",
//		CreditTypeAbbrev: validCreditTypeAbbrev,
//		Fee:              &core.DefaultParams().CreditClassFee[0],
//	})
//	s.Require().NoError(err)
//
//	testCases := []struct {
//		name      string
//		args      []string
//		expErr    bool
//		expErrMsg string
//	}{
//		{
//			name:      "invalid request: not enough args",
//			args:      []string{},
//			expErr:    true,
//			expErrMsg: "accepts 2 arg(s), received 0",
//		},
//		{
//			name:      "invalid request: bad id",
//			args:      []string{"", a1.String()},
//			expErr:    true,
//			expErrMsg: "class-id is required",
//		},
//		{
//			name:      "invalid request: no metadata",
//			args:      append([]string{classId, "", makeFlagFrom(a1.String())}, s.commonTxFlags()...),
//			expErr:    true,
//			expErrMsg: "base64_metadata is required",
//		},
//		{
//			name:   "valid request",
//			args:   append([]string{classId, newMetaData, makeFlagFrom(val0.Address.String())}, s.commonTxFlags()...),
//			expErr: false,
//		},
//		{
//			name:   "valid test: from key-name",
//			args:   append([]string{classId, newMetaData, makeFlagFrom("node0")}, s.commonTxFlags()...),
//			expErr: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			cmd := coreclient.TxUpdateClassMetadataCmd()
//			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expErr {
//				s.Require().Error(err)
//			} else {
//				s.Require().NoError(err)
//
//				// query the credit class info
//				query := coreclient.QueryClassInfoCmd()
//				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
//				s.Require().NoError(err, out.String())
//				var res core.QueryClassInfoResponse
//				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
//				s.Require().NoError(err)
//
//				// check metadata changed
//				s.Require().NoError(err)
//				s.Require().Equal(res.Info.Metadata, tc.args[1])
//			}
//		})
//	}
//}

//func (s *IntegrationTestSuite) TestTxUpdateIssuers() {
//	_, _, a2 := testdata.KeyTestPubAddr()
//	_, _, a3 := testdata.KeyTestPubAddr()
//	newIssuers := []string{a3.String(), a2.String()}
//	val0 := s.network.Validators[0]
//	clientCtx := val0.ClientCtx
//	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
//		Admin:            val0.Address.String(),
//		Issuers:          []string{val0.Address.String()},
//		Metadata:         "META",
//		CreditTypeAbbrev: validCreditTypeAbbrev,
//		Fee:              &core.DefaultParams().CreditClassFee[0],
//	})
//	s.Require().NoError(err)
//
//	makeArgs := func(add, remove []string, classId, from string) []string {
//		args := []string{classId}
//		if len(add) > 0 {
//			args = append(args, fmt.Sprintf("--%s=%s", coreclient.FlagAddIssuers, strings.Join(add, ",")))
//		}
//		if len(remove) > 0 {
//			args = append(args, fmt.Sprintf("--%s=%s", coreclient.FlagRemoveIssuers, strings.Join(remove, ",")))
//		}
//		args = append(args, makeFlagFrom(from))
//		return append(args, s.commonTxFlags()...)
//	}
//
//	testCases := []struct {
//		name       string
//		args       []string
//		expErr     bool
//		expErrMsg  string
//		expIssuers []string
//	}{
//		{
//			name:      "invalid request: no id",
//			args:      makeArgs(nil, nil, "", val0.Address.String()),
//			expErr:    true,
//			expErrMsg: "class-id is required",
//		},
//		{
//			name:       "valid add request",
//			args:       makeArgs(newIssuers, nil, classId, val0.Address.String()),
//			expErr:     false,
//			expIssuers: newIssuers,
//		},
//		{
//			name:       "valid remove request",
//			args:       makeArgs(nil, newIssuers, classId, val0.Address.String()),
//			expErr:     false,
//			expIssuers: []string{val0.Address.String()},
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			cmd := coreclient.TxUpdateClassIssuersCmd()
//			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expErr {
//				s.Require().Error(err)
//				s.Require().Contains(err.Error(), tc.expErrMsg)
//			} else {
//				s.Require().NoError(err)
//
//				// query the credit class info
//				query := coreclient.QueryClassInfoCmd()
//				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{classId, flagOutputJSON})
//				s.Require().NoError(err, out.String())
//				var res core.QueryClassInfoResponse
//				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
//				s.Require().NoError(err)
//
//				// TODO: check issuers list was changed https://github.com/regen-network/regen-ledger/issues/1025
//			}
//		})
//	}
//}

//func (s *IntegrationTestSuite) TestTxSell() {
//	val0 := s.network.Validators[0]
//	valAddrStr := val0.Address.String()
//	clientCtx := val0.ClientCtx
//	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
//		Admin:            valAddrStr,
//		Issuers:          []string{valAddrStr},
//		Metadata:         "meta",
//		CreditTypeAbbrev: validCreditTypeAbbrev,
//		Fee:              &core.DefaultParams().CreditClassFee[0],
//	})
//	s.Require().NoError(err)
//	projectId, err := s.createProject(clientCtx, &core.MsgCreateProject{
//		Issuer:          valAddrStr,
//		ClassId:         classId,
//		Metadata:        "meta",
//		ProjectLocation: "US-OR",
//		ProjectId:       rand.Str(3),
//	})
//	s.Require().NoError(err)
//	start, end := time.Now(), time.Now()
//	batchDenom, err := s.createBatch(clientCtx, &core.MsgCreateBatch{
//		Issuer:    valAddrStr,
//		ProjectId: projectId,
//		Issuance: []*core.BatchIssuance{
//			{Recipient: valAddrStr, TradableAmount: "999999999999999999"},
//		},
//		Metadata:  "meta",
//		StartDate: &start,
//		EndDate:   &end,
//		Open:      false,
//		OriginTx:  nil,
//		Note:      "",
//	})
//
//	expiration, err := types.ParseDate("expiration", "2024-01-01")
//	s.Require().NoError(err)
//
//	testCases := []struct {
//		name      string
//		args      []string
//		expErr    bool
//		expErrMsg string
//		expOrder  *marketplace.SellOrder
//	}{
//		{
//			name:      "missing args",
//			args:      []string{},
//			expErr:    true,
//			expErrMsg: "accepts 1 arg(s), received 0",
//		},
//		{
//			name:      "too many args",
//			args:      []string{"foo", "bar"},
//			expErr:    true,
//			expErrMsg: "accepts 1 arg(s), received 2",
//		},
//		{
//			name: "valid",
//			args: append(
//				[]string{
//					fmt.Sprintf("[{batch_denom: \"%s\", quantity: \"5\", ask_price: \"100uregen\", disable_auto_retire: false}]", batchDenom),
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expErr: false,
//			expOrder: &marketplace.SellOrder{
//				Seller:            val0.Address,
//				Quantity:          "5",
//				AskPrice:          "100",
//				DisableAutoRetire: false,
//			},
//		},
//		{
//			name: "valid with expiration",
//			args: append(
//				[]string{
//					fmt.Sprintf("[{batch_denom: \"%s\", quantity: \"5\", ask_price: \"100uregen\", disable_auto_retire: false, expiration: \"2024-01-01\"}]", batchDenom),
//					makeFlagFrom(val0.Address.String()),
//				},
//				s.commonTxFlags()...,
//			),
//			expErr: false,
//			expOrder: &marketplace.SellOrder{
//				Seller:            val0.Address,
//				Quantity:          "5",
//				AskPrice:          "100",
//				DisableAutoRetire: false,
//				Expiration:        types.ProtobufToGogoTimestamp(timestamppb.New(expiration)),
//			},
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			cmd := marketplaceclient.TxSellCmd()
//			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expErr {
//				s.Require().Error(err)
//				s.Require().Contains(err.Error(), tc.expErrMsg)
//			} else {
//				s.Require().NoError(err)
//				var res sdk.TxResponse
//				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//				s.Require().True(len(res.Logs) > 0)
//
//				found := false
//				for _, e := range res.Logs[0].Events {
//					if e.Type == proto.MessageName(&marketplace.EventSell{}) {
//						for _, attr := range e.Attributes {
//							if attr.Key == "order_id" {
//								found = true
//								orderIdStr := strings.Trim(attr.Value, "\"")
//								_, err := strconv.ParseUint(orderIdStr, 10, 64)
//								s.Require().NoError(err)
//								queryCmd := marketplaceclient.QuerySellOrderCmd()
//								queryArgs := []string{orderIdStr, flagOutputJSON}
//								queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
//								s.Require().NoError(err, queryOut.String())
//								var queryRes marketplace.QuerySellOrderResponse
//								s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))
//								s.Require().Equal(queryRes.SellOrder.Quantity, tc.expOrder.Quantity)
//								s.Require().Equal(tc.expOrder.DisableAutoRetire, queryRes.SellOrder.DisableAutoRetire)
//								s.Require().True(tc.expOrder.Expiration.Equal(queryRes.SellOrder.Expiration))
//								break
//							}
//							if found {
//								break
//							}
//						}
//					}
//					if found {
//						break
//					}
//				}
//				s.Require().True(found)
//			}
//		})
//	}
//}

func (s *IntegrationTestSuite) TestTxUpdateSellOrders() {
	val0 := s.network.Validators[0]
	clientCtx := val0.ClientCtx

	expiration, err := types.ParseDate("expiration", "2026-01-01")
	s.Require().NoError(err)

	testCases := []struct {
		name        string
		args        []string
		sellOrderId string
		expErr      bool
		expErrMsg   string
		expOrder    *ecocredit.SellOrder
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
			name: "missing sell order",
			args: append(
				[]string{
					"[{new_quantity: \"5\", new_ask_price: \"200regen\", disable_auto_retire: false}]",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expErr:    true,
			expErrMsg: "invalid sell order",
		},
		{
			name: "invalid sell order",
			args: append(
				[]string{
					"[{sell_order_id: \"foo\", new_quantity: \"5\", new_ask_price: \"200regen\", disable_auto_retire: false}]",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expErr:    true,
			expErrMsg: "invalid sell order",
		},
		{
			name: "missing new quantity",
			args: append(
				[]string{
					"[{sell_order_id: \"4\", new_ask_price: \"200regen\", disable_auto_retire: false}]",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expErr:    true,
			expErrMsg: "quantity must be positive decimal",
		},
		{
			name: "invalid new quantity",
			args: append(
				[]string{
					"[{sell_order_id: \"4\", new_quantity: \"foo\", new_ask_price: \"200regen\", disable_auto_retire: false}]",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expErr:    true,
			expErrMsg: "quantity must be positive decimal",
		},
		{
			name: "missing new ask price",
			args: append(
				[]string{
					"[{sell_order_id: \"4\", new_quantity: \"5\", disable_auto_retire: false}]",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expErr:    true,
			expErrMsg: "invalid decimal coin expression",
		},
		{
			name: "invalid new ask price",
			args: append(
				[]string{
					"[{sell_order_id: \"4\", new_quantity: \"5\", new_ask_price: \"foo\", disable_auto_retire: false}]",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expErr:    true,
			expErrMsg: "invalid decimal coin expression",
		},
		{
			name: "valid",
			args: append(
				[]string{
					"[{sell_order_id: \"4\", new_quantity: \"5\", new_ask_price: \"200regen\", disable_auto_retire: false}]",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			sellOrderId: "4",
			expErr:      false,
			expOrder: &ecocredit.SellOrder{
				OrderId:           4,
				Owner:             val0.Address.String(),
				BatchDenom:        batchDenom,
				Quantity:          "5",
				AskPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(200)},
				DisableAutoRetire: false,
			},
		},
		{
			name: "valid with expiration",
			args: append(
				[]string{
					"[{sell_order_id: \"5\", new_quantity: \"5\", new_ask_price: \"200regen\", disable_auto_retire: false, new_expiration: \"2026-01-01\"}]",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			sellOrderId: "5",
			expErr:      false,
			expOrder: &ecocredit.SellOrder{
				OrderId:           5,
				Owner:             val0.Address.String(),
				BatchDenom:        batchDenom,
				Quantity:          "5",
				AskPrice:          &sdk.Coin{Denom: "regen", Amount: sdk.NewInt(200)},
				DisableAutoRetire: false,
				Expiration:        &expiration,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := marketplaceclient.TxUpdateSellOrdersCmd()
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)

				// query sell order
				query := marketplaceclient.QuerySellOrderCmd()
				out, err := cli.ExecTestCLICmd(clientCtx, query, []string{
					tc.sellOrderId,
					flagOutputJSON,
				})
				s.Require().NoError(err, out.String())

				// unmarshal query response
				var res ecocredit.QuerySellOrderResponse
				err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
				s.Require().NoError(err)

				// verify expected order
				s.Require().Equal(tc.expOrder, res.SellOrder)
			}
		})
	}
}

//func (s *IntegrationTestSuite) TestCreateProject() {
//	val0 := s.network.Validators[0]
//	clientCtx := val0.ClientCtx
//	require := s.Require()
//
//	query := coreclient.QueryClassesCmd()
//	out, err := cli.ExecTestCLICmd(clientCtx, query, []string{flagOutputJSON})
//	require.NoError(err)
//
//	// unmarshal query response
//	var res ecocredit.QueryClassesResponse
//	err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
//	require.NoError(err)
//	require.Greater(len(res.Classes), 0)
//
//	testCases := []struct {
//		name      string
//		args      []string
//		expErr    bool
//		expErrMsg string
//	}{
//		{
//			"minimum args",
//			[]string{},
//			true,
//			"accepts 3 arg(s), received 0",
//		},
//		{
//			"missing project-location",
//			[]string{"C01"},
//			true,
//			"accepts 3 arg(s), received 1",
//		},
//		{
//			"missing metadata",
//			[]string{"C01", "AQ"},
//			true,
//			"accepts 3 arg(s), received 2",
//		},
//		{
//			"invalid metadata",
//			[]string{"C01", "AQ", "invalid-metadata", makeFlagFrom(val0.Address.String())},
//			true,
//			"metadata is malformed",
//		},
//		{
//			"invalid project location",
//			[]string{"C01", "abcde", "AQ==", makeFlagFrom(val0.Address.String())},
//			true,
//			"Invalid location: abcde",
//		},
//		{
//			"valid tx without project id",
//			append(
//				[]string{res.Classes[0].ClassId, "AQ", "AQ==", makeFlagFrom(val0.Address.String())},
//				s.commonTxFlags()...,
//			),
//			false,
//			"",
//		},
//		{
//			"valid tx with project id",
//			append(
//				[]string{res.Classes[0].ClassId, "AQ", "AQ==", makeFlagFrom(val0.Address.String()),
//					fmt.Sprintf("--project-id=%s", "C01P01"),
//				},
//				s.commonTxFlags()...,
//			),
//			false,
//			"",
//		},
//		{
//			"invalid project id format",
//			append(
//				[]string{res.Classes[0].ClassId, "AQ", "AQ==", makeFlagFrom(val0.Address.String()),
//					fmt.Sprintf("--project-id=%s", "C@a"),
//				},
//				s.commonTxFlags()...,
//			),
//			true,
//			"invalid project id",
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			cmd := coreclient.TxCreateProject()
//			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expErr {
//				require.Error(err)
//				require.Contains(err.Error(), tc.expErrMsg)
//			} else {
//				require.NoError(err)
//
//				var res sdk.TxResponse
//				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
//				require.Equal(uint32(0), res.Code)
//			}
//		})
//	}
//}

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
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, makeCreateProjectArgs(msg, append(s.commonTxFlags(), makeFlagFrom(msg.Issuer))...))
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
