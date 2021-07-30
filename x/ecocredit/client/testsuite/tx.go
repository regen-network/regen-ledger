package testsuite

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	proto "github.com/gogo/protobuf/proto"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	classInfo *ecocredit.ClassInfo
	batchInfo *ecocredit.BatchInfo
}

const (
	validMetadata = "AQ=="
	classId       = "18AV53K"
	batchId       = "1Lb4WV1"
)

var validMetadataBytes = []byte{0x1}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

// Write a MsgCreateBatch to a new temporary file and return the filename
func (s *IntegrationTestSuite) writeMsgCreateBatchJSON(msg *ecocredit.MsgCreateBatch) string {
	bytes, err := s.network.Validators[0].ClientCtx.Codec.MarshalJSON(msg)
	s.Require().NoError(err)

	return testutil.WriteToNewTempFile(s.T(), string(bytes)).Name()
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.network = network.New(s.T(), s.cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	val := s.network.Validators[0]

	// create an account for val
	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewValidator0", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	account := sdk.AccAddress(info.GetPubKey().Address())
	_, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		account,
		sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(2000))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	// Create a credit class
	out, err := cli.ExecTestCLICmd(val.ClientCtx, client.TxCreateClassCmd(),
		append(
			[]string{
				val.Address.String(),
				val.Address.String(),
				validMetadata,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			commonFlags...,
		),
	)

	s.Require().NoError(err, out.String())
	var txResp = sdk.TxResponse{}
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
	s.Require().Equal(uint32(0), txResp.Code, out.String())

	s.classInfo = &ecocredit.ClassInfo{
		ClassId:  classId,
		Designer: val.Address.String(),
		Issuers:  []string{val.Address.String()},
		Metadata: validMetadataBytes,
	}

	startDate, err := client.ParseDate("start date", "2021-01-01")
	s.Require().NoError(err)
	endDate, err := client.ParseDate("end date", "2021-02-01")
	s.Require().NoError(err)

	msgCreateBatch := ecocredit.MsgCreateBatch{
		ClassId: classId,
		Issuance: []*ecocredit.MsgCreateBatch_BatchIssuance{
			{
				Recipient:          val.Address.String(),
				TradableAmount:     "100",
				RetiredAmount:      "0.000001",
				RetirementLocation: "AB",
			},
		},
		Metadata:        validMetadataBytes,
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "GB",
	}

	// Write MsgCreateBatch to a temporary file
	batchFile := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Create a credit batch
	out, err = cli.ExecTestCLICmd(val.ClientCtx, client.TxCreateBatchCmd(),
		append(
			[]string{
				batchFile,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			commonFlags...,
		),
	)

	s.Require().NoError(err, out.String())
	txResp = sdk.TxResponse{}
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
	s.Require().Equal(uint32(0), txResp.Code, out.String())

	s.batchInfo = &ecocredit.BatchInfo{
		ClassId:         classId,
		BatchDenom:      fmt.Sprintf("%s/%s", classId, batchId),
		Issuer:          val.Address.String(),
		TotalAmount:     "100.000001",
		Metadata:        []byte{0x01},
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "GB",
	}
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

func (s *IntegrationTestSuite) TestTxCreateClass() {
	val0 := s.network.Validators[0]
	val1 := s.network.Validators[1]
	clientCtx := val0.ClientCtx

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		expectedErrMsg    string
		expectedClassInfo *ecocredit.ClassInfo
	}{
		{
			name:           "missing designer",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 3 arg(s), received 0",
		},
		{
			name:           "missing issuer",
			args:           []string{val0.Address.String()},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 3 arg(s), received 1",
		},
		{
			name:           "missing metadata",
			args:           []string{val0.Address.String(), val0.Address.String()},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 3 arg(s), received 2",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde", "abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 3 arg(s), received 4",
		},
		{
			name: "invalid designer",
			args: append(
				[]string{
					"abcde",
					val0.Address.String(),
					validMetadata,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
		},
		{
			name: "invalid issuer",
			args: append(
				[]string{
					val0.Address.String(),
					"abcde",
					validMetadata,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
		},
		{
			name: "invalid metadata",
			args: append(
				[]string{
					val0.Address.String(),
					val0.Address.String(),
					"=",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "metadata is malformed, proper base64 string is required",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
					val0.Address.String(),
					val0.Address.String(),
					validMetadata,
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "required flag(s) \"from\" not set",
		},
		{
			name: "single issuer",
			args: append(
				[]string{
					val0.Address.String(),
					val0.Address.String(),
					validMetadata,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedClassInfo: &ecocredit.ClassInfo{
				Designer: val0.Address.String(),
				Issuers:  []string{val0.Address.String()},
				Metadata: []byte{0x1},
			},
		},
		{
			name: "multiple issuers",
			args: append(
				[]string{
					val0.Address.String(),
					strings.Join(
						[]string{
							val0.Address.String(),
							val1.Address.String(),
						},
						",",
					),
					validMetadata,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr: false,
			expectedClassInfo: &ecocredit.ClassInfo{
				Designer: val0.Address.String(),
				Issuers:  []string{val0.Address.String(), val1.Address.String()},
				Metadata: []byte{0x1},
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

			cmd := client.TxCreateClassCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res sdk.TxResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				classIdFound := false
				for _, e := range res.Logs[0].Events {
					if e.Type == proto.MessageName(&ecocredit.EventCreateClass{}) {
						for _, attr := range e.Attributes {
							if attr.Key == "class_id" {
								classIdFound = true
								classId := strings.Trim(attr.Value, "\"")

								queryCmd := client.QueryClassInfoCmd()
								queryArgs := []string{classId, flagOutputJSON}
								queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
								s.Require().NoError(err, queryOut.String())
								var queryRes ecocredit.QueryClassInfoResponse
								s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))

								s.Require().Equal(tc.expectedClassInfo.Designer, queryRes.Info.Designer)
								s.Require().Equal(tc.expectedClassInfo.Issuers, queryRes.Info.Issuers)
								s.Require().Equal(tc.expectedClassInfo.Metadata, queryRes.Info.Metadata)
							}
						}
					}
				}
				s.Require().True(classIdFound)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxCreateBatch() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	// Write some invalid JSON to a file
	invalidJsonFile := testutil.WriteToNewTempFile(s.T(), "{asdljdfklfklksdflk}")

	// Create a valid MsgCreateBatch
	startDate, err := client.ParseDate("start date", "2021-01-01")
	s.Require().NoError(err)
	endDate, err := client.ParseDate("end date", "2021-02-01")
	s.Require().NoError(err)

	msgCreateBatch := ecocredit.MsgCreateBatch{
		ClassId: s.classInfo.ClassId,
		Issuance: []*ecocredit.MsgCreateBatch_BatchIssuance{
			{
				Recipient:          s.network.Validators[1].Address.String(),
				TradableAmount:     "100",
				RetiredAmount:      "0.000001",
				RetirementLocation: "AB",
			},
		},
		Metadata:        validMetadataBytes,
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "GB",
	}

	validBatchJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid class id
	msgCreateBatch.ClassId = "abcde"
	invalidClassIdJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with missing start date
	msgCreateBatch.ClassId = s.classInfo.ClassId
	msgCreateBatch.StartDate = nil
	missingStartDateJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with missing end date
	msgCreateBatch.StartDate = &startDate
	msgCreateBatch.EndDate = nil
	missingEndDateJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with missing project location
	msgCreateBatch.EndDate = &endDate
	msgCreateBatch.ProjectLocation = ""
	missingProjectLocationJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid issuance recipient
	msgCreateBatch.ProjectLocation = "AB"
	msgCreateBatch.Issuance[0].Recipient = "abcde"
	invalidRecipientJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid issuance tradable amount
	msgCreateBatch.Issuance[0].Recipient = s.network.Validators[1].Address.String()
	msgCreateBatch.Issuance[0].TradableAmount = "abcde"
	invalidTradableAmountJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid issuance retired amount
	msgCreateBatch.Issuance[0].TradableAmount = "100"
	msgCreateBatch.Issuance[0].RetiredAmount = "abcde"
	invalidRetiredAmountJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	// Write batch with invalid issuance retirement location
	msgCreateBatch.Issuance[0].RetiredAmount = "0.000001"
	msgCreateBatch.Issuance[0].RetirementLocation = "abcde"
	invalidRetirementLocationJson := s.writeMsgCreateBatchJSON(&msgCreateBatch)

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		errInTxResponse   bool
		expectedErrMsg    string
		expectedBatchInfo *ecocredit.BatchInfo
	}{
		{
			name:           "missing filename",
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
			name: "invalid json",
			args: append(
				[]string{
					invalidJsonFile.Name(),
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Error: parsing batch JSON",
		},
		{
			name: "invalid class id",
			args: append(
				[]string{
					invalidClassIdJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:       true,
			errInTxResponse: true,
			expectedErrMsg:  "not found",
		},
		{
			name: "missing start date",
			args: append(
				[]string{
					missingStartDateJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Must provide a start date for the credit batch: invalid request",
		},
		{
			name: "missing end date",
			args: append(
				[]string{
					missingEndDateJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Must provide an end date for the credit batch: invalid request",
		},
		{
			name: "missing project location",
			args: append(
				[]string{
					missingProjectLocationJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Invalid retirement location",
		},
		{
			name: "invalid issuance recipient",
			args: append(
				[]string{
					invalidRecipientJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
		},
		{
			name: "invalid issuance tradable amount",
			args: append(
				[]string{
					invalidTradableAmountJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "expected a non-negative decimal, got abcde",
		},
		{
			name: "invalid issuance retired amount",
			args: append(
				[]string{
					invalidRetiredAmountJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "expected a non-negative decimal, got abcde",
		},
		{
			name: "invalid issuance retirement location",
			args: append(
				[]string{
					invalidRetirementLocationJson,
					makeFlagFrom(val.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Invalid retirement location: abcde",
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
			expectedBatchInfo: &ecocredit.BatchInfo{
				ClassId:         s.classInfo.ClassId,
				Issuer:          val.Address.String(),
				TotalAmount:     "100.000001",
				Metadata:        []byte{0x1},
				AmountCancelled: "0",
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

			cmd := client.TxCreateBatchCmd()
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
					if e.Type == proto.MessageName(&ecocredit.EventCreateBatch{}) {
						for _, attr := range e.Attributes {
							if attr.Key == "batch_denom" {
								batchDenomFound = true
								batchDenom := strings.Trim(attr.Value, "\"")

								queryCmd := client.QueryBatchInfoCmd()
								queryArgs := []string{batchDenom, flagOutputJSON}
								queryOut, err := cli.ExecTestCLICmd(clientCtx, queryCmd, queryArgs)
								s.Require().NoError(err, queryOut.String())
								var queryRes ecocredit.QueryBatchInfoResponse
								s.Require().NoError(clientCtx.Codec.UnmarshalJSON(queryOut.Bytes(), &queryRes))

								s.Require().Equal(tc.expectedBatchInfo.ClassId, queryRes.Info.ClassId)
								s.Require().Equal(tc.expectedBatchInfo.Issuer, queryRes.Info.Issuer)
								s.Require().Equal(tc.expectedBatchInfo.TotalAmount, queryRes.Info.TotalAmount)
								s.Require().Equal(tc.expectedBatchInfo.Metadata, queryRes.Info.Metadata)
								s.Require().Equal(tc.expectedBatchInfo.AmountCancelled, queryRes.Info.AmountCancelled)
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

	validCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"88\", retired_amount: \"2\", retirement_location: \"AB-CD\"}]", s.batchInfo.BatchDenom)
	invalidBatchDenomCredits := fmt.Sprintf("[{batch_denom: abcde, tradable_amount: \"88\", retired_amount: \"2\", retirement_location: \"AB-CD\"}]")
	invalidTradableAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"abcde\", retired_amount: \"2\", retirement_location: \"AB-CD\"}]", s.batchInfo.BatchDenom)
	invalidRetiredAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"88\", retired_amount: \"abcde\", retirement_location: \"AB-CD\"}]", s.batchInfo.BatchDenom)
	invalidRetirementLocationCredits := fmt.Sprintf("[{batch_denom: \"%s\", tradable_amount: \"88\", retired_amount: \"2\", retirement_location: \"abcde\"}]", s.batchInfo.BatchDenom)

	testCases := []struct {
		name            string
		args            []string
		expectErr       bool
		errInTxResponse bool
		expectedErrMsg  string
	}{
		{
			name:           "missing recipient",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:           "missing credits",
			args:           []string{val1.Address.String()},
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
			name: "invalid recipient",
			args: append(
				[]string{
					"abcde",
					validCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "decoding bech32 failed: invalid bech32 string length 5",
		},
		{
			name: "invalid batch denom",
			args: append(
				[]string{
					val1.Address.String(),
					invalidBatchDenomCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:       true,
			errInTxResponse: true,
			expectedErrMsg:  "abcde is not a valid credit batch denom",
		},
		{
			name: "invalid tradable amount",
			args: append(
				[]string{
					val1.Address.String(),
					invalidTradableAmountCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "expected a non-negative decimal, got abcde",
		},
		{
			name: "invalid retired amount",
			args: append(
				[]string{
					val1.Address.String(),
					invalidRetiredAmountCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "expected a non-negative decimal, got abcde",
		},
		{
			name: "invalid retirement location",
			args: append(
				[]string{
					val1.Address.String(),
					invalidRetirementLocationCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Invalid retirement location: abcde",
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

			cmd := client.TxSendCmd()
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
	clientCtx := val0.ClientCtx

	validCredits := fmt.Sprintf("[{batch_denom: \"%s\", amount: \"5\"}]", s.batchInfo.BatchDenom)
	invalidBatchDenomCredits := fmt.Sprintf("[{batch_denom: abcde, amount: \"5\"}]")
	invalidAmountCredits := fmt.Sprintf("[{batch_denom: \"%s\", amount: \"abcde\"}]", s.batchInfo.BatchDenom)

	testCases := []struct {
		name            string
		args            []string
		expectErr       bool
		errInTxResponse bool
		expectedErrMsg  string
	}{
		{
			name:           "missing credits",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:           "missing retirement location",
			args:           []string{validCredits},
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
			name: "invalid batch denom",
			args: append(
				[]string{
					invalidBatchDenomCredits,
					"AB-CD 12345",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:       true,
			errInTxResponse: true,
			expectedErrMsg:  "abcde is not a valid credit batch denom",
		},
		{
			name: "invalid amount",
			args: append(
				[]string{
					invalidAmountCredits,
					"AB-CD 12345",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "expected a positive decimal, got abcde",
		},
		{
			name: "invalid retirement location",
			args: append(
				[]string{
					validCredits,
					"abcde",
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "Invalid retirement location: abcde",
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

			cmd := client.TxRetireCmd()
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
	clientCtx := val0.ClientCtx

	validCredits := fmt.Sprintf("5:%s", s.batchInfo.BatchDenom)
	invalidBatchDenomCredits := "5:abcde"
	invalidAmountCredits := fmt.Sprintf("abcde:%s", s.batchInfo.BatchDenom)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "missing credits",
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
			name: "invalid batch denom",
			args: append(
				[]string{
					invalidBatchDenomCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid credit expression",
		},
		{
			name: "invalid amount",
			args: append(
				[]string{
					invalidAmountCredits,
					makeFlagFrom(val0.Address.String()),
				},
				s.commonTxFlags()...,
			),
			expectErr:      true,
			expectedErrMsg: "invalid credit expression",
		},
		{
			name: "missing from flag",
			args: append(
				[]string{
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
					validCredits,
					makeFlagFrom(val0.Address.String()),
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

			cmd := client.TxCancelCmd()
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
