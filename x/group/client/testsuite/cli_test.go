package testsuite

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/testutil/cli"
	"github.com/regen-network/regen-ledger/testutil/network"

	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/client"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	group         *group.GroupInfo
	groupAccount  *group.GroupAccountInfo
	groupAccount2 *group.GroupAccountInfo
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	cfg.NumValidators = 2

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	val := s.network.Validators[0]

	// create a new account
	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewValidator", keyring.English, sdk.FullFundraiserPath, hd.Secp256k1)
	s.Require().NoError(err)

	account := sdk.AccAddress(info.GetPubKey().Address())
	_, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		account,
		sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(200))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	// create a group
	validMembers := fmt.Sprintf(`[{
	  "address": "%s",
		"weight": "1",
		"metadata": "AQ=="
	}]`, val.Address.String())
	validMembersFile := testutil.WriteToNewTempFile(s.T(), validMembers)
	out, err := cli.ExecTestCLICmd(val.ClientCtx, client.MsgCreateGroupCmd(),
		append(
			[]string{
				val.Address.String(),
				"AQ==",
				validMembersFile.Name(),
			},
			commonFlags...,
		),
	)

	s.Require().NoError(err, out.String())
	var txResp = sdk.TxResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &txResp), out.String())
	s.Require().Equal(uint32(0), txResp.Code, out.String())

	s.group = &group.GroupInfo{GroupId: group.ID(1), Admin: val.Address.String(), Metadata: []byte{1}, TotalWeight: "1", Version: 1}

	// create 2 group accounts
	for i := 0; i < 2; i++ {
		out, err = cli.ExecTestCLICmd(val.ClientCtx, client.MsgCreateGroupAccountCmd(),
			append(
				[]string{
					val.Address.String(),
					"1",
					"AQ==",
					"{\"@type\":\"/regen.group.v1alpha1.ThresholdDecisionPolicy\", \"threshold\":\"1\", \"timeout\":\"1s\"}",
				},
				commonFlags...,
			),
		)
		s.Require().NoError(err, out.String())
		s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &txResp), out.String())
		s.Require().Equal(uint32(0), txResp.Code, out.String())

		out, err = cli.ExecTestCLICmd(val.ClientCtx, client.QueryGroupAccountsByGroupCmd(), []string{"1", fmt.Sprintf("--%s=json", tmcli.OutputFlag)})
		s.Require().NoError(err, out.String())
	}

	var res group.QueryGroupAccountsByGroupResponse
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &res))
	s.Require().Equal(len(res.GroupAccounts), 2)
	s.groupAccount = res.GroupAccounts[0]
	s.groupAccount2 = res.GroupAccounts[1]
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestQueryGroupInfo() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		expectedCode uint32
	}{
		{
			"group not found",
			[]string{"12345", fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			true,
			"not found: invalid request",
			0,
		},
		{
			"group found",
			[]string{strconv.FormatUint(s.group.GroupId.Uint64(), 10), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.QueryGroupInfoCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var g group.GroupInfo
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &g))
				s.Require().Equal(s.group.GroupId, g.GroupId)
				s.Require().Equal(s.group.Admin, g.Admin)
				s.Require().Equal(s.group.TotalWeight, g.TotalWeight)
				s.Require().Equal(s.group.Metadata, g.Metadata)
				s.Require().Equal(s.group.Version, g.Version)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryGroupMembers() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name          string
		args          []string
		expectErr     bool
		expectErrMsg  string
		expectedCode  uint32
		expectMembers []*group.GroupMember
	}{
		{
			"no group",
			[]string{"12345", fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
			[]*group.GroupMember{},
		},
		{
			"members found",
			[]string{strconv.FormatUint(s.group.GroupId.Uint64(), 10), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
			[]*group.GroupMember{
				{
					GroupId: s.group.GroupId,
					Member: &group.Member{
						Address:  val.Address.String(),
						Weight:   "1",
						Metadata: []byte{1},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.QueryGroupMembersCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res group.QueryGroupMembersResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(len(res.Members), len(tc.expectMembers))
				for i := range res.Members {
					s.Require().Equal(res.Members[i].GroupId, tc.expectMembers[i].GroupId)
					s.Require().Equal(res.Members[i].Member.Address, tc.expectMembers[i].Member.Address)
					s.Require().Equal(res.Members[i].Member.Metadata, tc.expectMembers[i].Member.Metadata)
					s.Require().Equal(res.Members[i].Member.Weight, tc.expectMembers[i].Member.Weight)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryGroupsByAdmin() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		expectedCode uint32
		expectGroups []*group.GroupInfo
	}{
		{
			"invalid admin address",
			[]string{"invalid"},
			true,
			"decoding bech32 failed: invalid bech32 string",
			0,
			[]*group.GroupInfo{},
		},
		{
			"no group",
			[]string{s.network.Validators[1].Address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
			[]*group.GroupInfo{},
		},
		{
			"found groups",
			[]string{val.Address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
			[]*group.GroupInfo{
				s.group,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.QueryGroupsByAdminCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res group.QueryGroupsByAdminResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(len(res.Groups), len(tc.expectGroups))
				for i := range res.Groups {
					s.Require().Equal(res.Groups[i].GroupId, tc.expectGroups[i].GroupId)
					s.Require().Equal(res.Groups[i].Metadata, tc.expectGroups[i].Metadata)
					s.Require().Equal(res.Groups[i].Version, tc.expectGroups[i].Version)
					s.Require().Equal(res.Groups[i].TotalWeight, tc.expectGroups[i].TotalWeight)
					s.Require().Equal(res.Groups[i].Admin, tc.expectGroups[i].Admin)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryGroupAccountInfo() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		expectedCode uint32
	}{
		{
			"invalid account address",
			[]string{"invalid", fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			true,
			"decoding bech32 failed: invalid bech32",
			0,
		},
		{
			"group account not found",
			[]string{val.Address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			true,
			"not found: invalid request",
			0,
		},
		{
			"group account found",
			[]string{s.groupAccount.GroupAccount, fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.QueryGroupAccountInfoCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var g group.GroupAccountInfo
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &g))
				s.Require().Equal(s.groupAccount.GroupId, g.GroupId)
				s.Require().Equal(s.groupAccount.GroupAccount, g.GroupAccount)
				s.Require().Equal(s.groupAccount.Admin, g.Admin)
				s.Require().Equal(s.groupAccount.Metadata, g.Metadata)
				s.Require().Equal(s.groupAccount.Version, g.Version)
				s.Require().Equal(s.groupAccount.GetDecisionPolicy(), g.GetDecisionPolicy())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryGroupAccountsByGroup() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectErrMsg        string
		expectedCode        uint32
		expectGroupAccounts []*group.GroupAccountInfo
	}{
		{
			"invalid group id",
			[]string{""},
			true,
			"strconv.ParseUint: parsing \"\": invalid syntax",
			0,
			[]*group.GroupAccountInfo{},
		},
		{
			"no group account",
			[]string{"12345", fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
			[]*group.GroupAccountInfo{},
		},
		{
			"found group accounts",
			[]string{strconv.FormatUint(s.group.GroupId.Uint64(), 10), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
			[]*group.GroupAccountInfo{
				s.groupAccount,
				s.groupAccount2,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.QueryGroupAccountsByGroupCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res group.QueryGroupAccountsByGroupResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(len(res.GroupAccounts), len(tc.expectGroupAccounts))
				for i := range res.GroupAccounts {
					s.Require().Equal(res.GroupAccounts[i].GroupId, tc.expectGroupAccounts[i].GroupId)
					s.Require().Equal(res.GroupAccounts[i].Metadata, tc.expectGroupAccounts[i].Metadata)
					s.Require().Equal(res.GroupAccounts[i].Version, tc.expectGroupAccounts[i].Version)
					s.Require().Equal(res.GroupAccounts[i].Admin, tc.expectGroupAccounts[i].Admin)
					s.Require().Equal(res.GroupAccounts[i].GetDecisionPolicy(), tc.expectGroupAccounts[i].GetDecisionPolicy())
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryGroupAccountsByAdmin() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectErrMsg        string
		expectedCode        uint32
		expectGroupAccounts []*group.GroupAccountInfo
	}{
		{
			"invalid admin address",
			[]string{"invalid"},
			true,
			"decoding bech32 failed: invalid bech32 string",
			0,
			[]*group.GroupAccountInfo{},
		},
		{
			"no group account",
			[]string{s.network.Validators[1].Address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
			[]*group.GroupAccountInfo{},
		},
		{
			"found group accounts",
			[]string{val.Address.String(), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
			"",
			0,
			[]*group.GroupAccountInfo{
				s.groupAccount,
				s.groupAccount2,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.QueryGroupAccountsByAdminCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res group.QueryGroupAccountsByAdminResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(len(res.GroupAccounts), len(tc.expectGroupAccounts))
				for i := range res.GroupAccounts {
					s.Require().Equal(res.GroupAccounts[i].GroupId, tc.expectGroupAccounts[i].GroupId)
					s.Require().Equal(res.GroupAccounts[i].Metadata, tc.expectGroupAccounts[i].Metadata)
					s.Require().Equal(res.GroupAccounts[i].Version, tc.expectGroupAccounts[i].Version)
					s.Require().Equal(res.GroupAccounts[i].Admin, tc.expectGroupAccounts[i].Admin)
					s.Require().Equal(res.GroupAccounts[i].GetDecisionPolicy(), tc.expectGroupAccounts[i].GetDecisionPolicy())
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxCreateGroup() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	invalidMembersAddress := `[{
	"address": "",
	"weight": "1"
}]`
	invalidMembersAddressFile := testutil.WriteToNewTempFile(s.T(), invalidMembersAddress)

	invalidMembersWeight := fmt.Sprintf(`[{
	  "address": "%s",
		"weight": "0"
	}]`, val.Address.String())
	invalidMembersWeightFile := testutil.WriteToNewTempFile(s.T(), invalidMembersWeight)

	invalidMembersMetadata := fmt.Sprintf(`[{
	  "address": "%s",
		"weight": "1",
		"metadata": "AQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQ=="
	}]`, val.Address.String())
	invalidMembersMetadataFile := testutil.WriteToNewTempFile(s.T(), invalidMembersMetadata)

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"correct data",
			append(
				[]string{
					val.Address.String(),
					"",
					"",
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
		{
			"group metadata too long",
			append(
				[]string{
					val.Address.String(),
					"AQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQ==",
					"",
				},
				commonFlags...,
			),
			true,
			"",
			nil,
			204,
		},
		{
			"invalid members address",
			append(
				[]string{
					val.Address.String(),
					"null",
					invalidMembersAddressFile.Name(),
				},
				commonFlags...,
			),
			true,
			"message validation failed: members: address: empty address string is not allowed",
			nil,
			0,
		},
		{
			"invalid members weight",
			append(
				[]string{
					val.Address.String(),
					"null",
					invalidMembersWeightFile.Name(),
				},
				commonFlags...,
			),
			true,
			"message validation failed: member weight: expected a positive decimal, got 0",
			nil,
			0,
		},
		{
			"members metadata too long",
			append(
				[]string{
					val.Address.String(),
					"null",
					invalidMembersMetadataFile.Name(),
				},
				commonFlags...,
			),
			true,
			"",
			nil,
			204,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.MsgCreateGroupCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateGroupAdmin() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"correct data",
			append(
				[]string{
					val.Address.String(),
					"2",
					s.network.Validators[1].Address.String(),
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
		{
			"group doesn't exist",
			append(
				[]string{
					val.Address.String(),
					"12345",
					s.network.Validators[1].Address.String(),
				},
				commonFlags...,
			),
			true,
			"",
			nil,
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.MsgUpdateGroupAdminCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateGroupMetadata() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"correct data",
			append(
				[]string{
					val.Address.String(),
					strconv.FormatUint(s.group.GroupId.Uint64(), 10),
					"AQ==",
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
		{
			"group metadata too long",
			append(
				[]string{
					val.Address.String(),
					strconv.FormatUint(s.group.GroupId.Uint64(), 10),
					"AQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQ==",
				},
				commonFlags...,
			),
			true,
			"group metadata: limit exceeded",
			nil,
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.MsgUpdateGroupMetadataCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateGroupMembers() {
	// TODO
}

func (s *IntegrationTestSuite) TestTxCreateGroupAccount() {
	val := s.network.Validators[0]
	wrongAdmin := s.network.Validators[1].Address
	clientCtx := val.ClientCtx

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	validMembers := fmt.Sprintf(`[{
	  "address": "%s",
		"weight": "1",
		"metadata": "AQ=="
	}]`, val.Address.String())
	validMembersFile := testutil.WriteToNewTempFile(s.T(), validMembers)
	out, err := cli.ExecTestCLICmd(val.ClientCtx, client.MsgCreateGroupCmd(),
		append(
			[]string{
				val.Address.String(),
				"AQ==",
				validMembersFile.Name(),
			},
			commonFlags...,
		),
	)

	s.Require().NoError(err, out.String())
	var txResp = sdk.TxResponse{}
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &txResp), out.String())
	s.Require().Equal(uint32(0), txResp.Code, out.String())

	groupID := group.ID(2)

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"correct data",
			append(
				[]string{
					val.Address.String(),
					fmt.Sprintf("%v", groupID),
					"AQ==",
					"{\"@type\":\"/regen.group.v1alpha1.ThresholdDecisionPolicy\", \"threshold\":\"1\", \"timeout\":\"1s\"}",
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
		{
			"wrong admin",
			append(
				[]string{
					wrongAdmin.String(),
					fmt.Sprintf("%v", groupID),
					"AQ==",
					"{\"@type\":\"/regen.group.v1alpha1.ThresholdDecisionPolicy\", \"threshold\":\"1\", \"timeout\":\"1s\"}",
				},
				commonFlags...,
			),
			true,
			"The specified item could not be found in the keyring",
			&sdk.TxResponse{},
			0,
		},
		{
			"metadata too long",
			append(
				[]string{
					val.Address.String(),
					fmt.Sprintf("%v", groupID),
					strings.Repeat("a", 500),
					"{\"@type\":\"/regen.group.v1alpha1.ThresholdDecisionPolicy\", \"threshold\":\"1\", \"timeout\":\"1s\"}",
				},
				commonFlags...,
			),
			true,
			"group account metadata: limit exceeded",
			&sdk.TxResponse{},
			0,
		},
		{
			"wrong group id",
			append(
				[]string{
					val.Address.String(),
					"10",
					"AQ==",
					"{\"@type\":\"/regen.group.v1alpha1.ThresholdDecisionPolicy\", \"threshold\":\"1\", \"timeout\":\"1s\"}",
				},
				commonFlags...,
			),
			true,
			"not found",
			&sdk.TxResponse{},
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.MsgCreateGroupAccountCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateGroupAccountAdmin() {
	val := s.network.Validators[0]
	newAdmin := s.network.Validators[1].Address
	clientCtx := val.ClientCtx
	groupAccount := s.groupAccount2

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"correct data",
			append(
				[]string{
					groupAccount.Admin,
					groupAccount.GroupAccount,
					newAdmin.String(),
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
		{
			"wrong admin",
			append(
				[]string{
					newAdmin.String(),
					groupAccount.GroupAccount,
					newAdmin.String(),
				},
				commonFlags...,
			),
			true,
			"The specified item could not be found in the keyring",
			&sdk.TxResponse{},
			0,
		},
		{
			"wrong group account",
			append(
				[]string{
					groupAccount.Admin,
					newAdmin.String(),
					newAdmin.String(),
				},
				commonFlags...,
			),
			true,
			"load group account: not found",
			&sdk.TxResponse{},
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.MsgUpdateGroupAccountAdminCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateGroupAccountDecisionPolicy() {
	val := s.network.Validators[0]
	newAdmin := s.network.Validators[1].Address
	clientCtx := val.ClientCtx
	groupAccount := s.groupAccount

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"correct data",
			append(
				[]string{
					groupAccount.Admin,
					groupAccount.GroupAccount,
					"{\"@type\":\"/regen.group.v1alpha1.ThresholdDecisionPolicy\", \"threshold\":\"2\", \"timeout\":\"2s\"}",
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
		{
			"wrong admin",
			append(
				[]string{
					newAdmin.String(),
					groupAccount.GroupAccount,
					"{\"@type\":\"/regen.group.v1alpha1.ThresholdDecisionPolicy\", \"threshold\":\"1\", \"timeout\":\"1s\"}",
				},
				commonFlags...,
			),
			true,
			"The specified item could not be found in the keyring",
			&sdk.TxResponse{},
			0,
		},
		{
			"wrong group account",
			append(
				[]string{
					groupAccount.Admin,
					newAdmin.String(),
					"{\"@type\":\"/regen.group.v1alpha1.ThresholdDecisionPolicy\", \"threshold\":\"1\", \"timeout\":\"1s\"}",
				},
				commonFlags...,
			),
			true,
			"load group account: not found",
			&sdk.TxResponse{},
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.MsgUpdateGroupAccountDecisionPolicyCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxUpdateGroupAccountMetadata() {
	val := s.network.Validators[0]
	newAdmin := s.network.Validators[1].Address
	clientCtx := val.ClientCtx
	groupAccount := s.groupAccount

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		expectErrMsg string
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"correct data",
			append(
				[]string{
					groupAccount.Admin,
					groupAccount.GroupAccount,
					"AQ==",
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
		{
			"long metadata",
			append(
				[]string{
					groupAccount.Admin,
					groupAccount.GroupAccount,
					strings.Repeat("a", 500),
				},
				commonFlags...,
			),
			true,
			"group account metadata: limit exceeded",
			&sdk.TxResponse{},
			0,
		},
		{
			"wrong admin",
			append(
				[]string{
					newAdmin.String(),
					groupAccount.GroupAccount,
					"AQ==",
				},
				commonFlags...,
			),
			true,
			"The specified item could not be found in the keyring",
			&sdk.TxResponse{},
			0,
		},
		{
			"wrong group account",
			append(
				[]string{
					groupAccount.Admin,
					newAdmin.String(),
					"AQ==",
				},
				commonFlags...,
			),
			true,
			"load group account: not found",
			&sdk.TxResponse{},
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := client.MsgUpdateGroupAccountMetadataCmd()

			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
