package testsuite

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/client"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	val     *network.Validator

	addr1 sdk.AccAddress
	addr2 sdk.AccAddress

	iri1  string
	iri2  string
	hash1 *data.ContentHash
	hash2 *data.ContentHash

	resolverID uint64
	url        string
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.cfg.NumValidators = 2
	nw, err := network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)
	s.network = nw

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.val = s.network.Validators[0]

	info1, _, err := s.val.ClientCtx.Keyring.NewMnemonic("acc1", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	info2, _, err := s.val.ClientCtx.Keyring.NewMnemonic("acc2", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	pk, err := info1.GetPubKey()
	s.Require().NoError(err)
	s.addr1 = sdk.AccAddress(pk.Address())
	_, err = banktestutil.MsgSendExec(
		s.val.ClientCtx,
		s.val.Address,
		s.addr1,
		sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(2000))),
		commonFlags...,
	)
	s.Require().NoError(err)

	pk, err = info2.GetPubKey()
	s.Require().NoError(err)
	s.addr2 = sdk.AccAddress(pk.Address())
	_, err = banktestutil.MsgSendExec(
		s.val.ClientCtx,
		s.val.Address,
		s.addr2,
		sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(2000))),
		commonFlags...,
	)
	s.Require().NoError(err)

	s.iri1 = "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"
	s.iri2 = "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	s.hash1, err = data.ParseIRI(s.iri1)
	s.Require().NoError(err)

	s.hash2, err = data.ParseIRI(s.iri2)
	s.Require().NoError(err)

	iris := []string{s.iri1, s.iri2}

	for _, iri := range iris {
		_, err := cli.ExecTestCLICmd(s.val.ClientCtx, client.MsgAnchorCmd(),
			append(
				[]string{
					iri,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, s.addr1.String()),
				},
				commonFlags...,
			),
		)
		s.Require().NoError(err)

		_, err = cli.ExecTestCLICmd(s.val.ClientCtx, client.MsgAttestCmd(),
			append(
				[]string{
					iri,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, s.addr1.String()),
				},
				commonFlags...,
			),
		)
		s.Require().NoError(err)

		_, err = cli.ExecTestCLICmd(s.val.ClientCtx, client.MsgAttestCmd(),
			append(
				[]string{
					iri,
					fmt.Sprintf("--%s=%s", flags.FlagFrom, s.addr2.String()),
				},
				commonFlags...,
			),
		)
		s.Require().NoError(err)
	}

	s.url = "https://foo.bar"

	out, err := cli.ExecTestCLICmd(s.val.ClientCtx, client.MsgDefineResolverCmd(),
		append(
			[]string{
				s.url,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.addr1.String()),
			},
			commonFlags...,
		),
	)
	s.Require().NoError(err)

	var res sdk.TxResponse
	s.Require().NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

	id := strings.Trim(res.Logs[0].Events[1].Attributes[0].Value, "\"")
	s.resolverID, err = strconv.ParseUint(id, 10, 64)
	s.Require().NoError(err)

	chs := &data.ContentHashes{ContentHashes: []*data.ContentHash{s.hash1}}

	bz, err := s.val.ClientCtx.Codec.MarshalJSON(chs)
	s.Require().NoError(err)

	filePath := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

	_, err = cli.ExecTestCLICmd(s.val.ClientCtx, client.MsgRegisterResolverCmd(), append(
		[]string{
			fmt.Sprintf("%d", s.resolverID),
			filePath,
			fmt.Sprintf("--%s=%s", flags.FlagFrom, s.addr1.String()),
		},
		commonFlags...,
	))
	s.Require().NoError(err)

	out2, err := cli.ExecTestCLICmd(s.val.ClientCtx, client.MsgDefineResolverCmd(),
		append(
			[]string{
				"https://bar.baz",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, s.addr1.String()),
			},
			commonFlags...,
		),
	)
	s.Require().NoError(err)

	var res2 sdk.TxResponse
	s.Require().NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out2.Bytes(), &res2))

	id2 := strings.Trim(res2.Logs[0].Events[1].Attributes[0].Value, "\"")
	resolverID2, err := strconv.ParseUint(id2, 10, 64)
	s.Require().NoError(err)

	_, err = cli.ExecTestCLICmd(s.val.ClientCtx, client.MsgRegisterResolverCmd(), append(
		[]string{
			fmt.Sprintf("%d", resolverID2),
			filePath,
			fmt.Sprintf("--%s=%s", flags.FlagFrom, s.addr1.String()),
		},
		commonFlags...,
	))
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.Require().NoError(s.network.WaitForNextBlock())
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestTxAnchor() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.FromAddress = val.Address
	require := s.Require()

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, clientCtx.GetFromAddress().String()),
	}

	testCases := []struct {
		name   string
		iri    string
		expErr bool
		errMsg string
	}{
		{
			name:   "valid",
			iri:    "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
			expErr: false,
		},
		{
			name:   "no arg",
			expErr: true,
			errMsg: "iri cannot be empty",
		},
		{
			name:   "bad iri",
			iri:    "foo",
			expErr: true,
			errMsg: "invalid iri",
		},
	}

	cmd := client.MsgAnchorCmd()
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			args := []string{tc.iri}
			args = append(args, commonFlags...)
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxAttest() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.FromAddress = val.Address
	require := s.Require()

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, clientCtx.GetFromAddress().String()),
	}
	// first we anchor some data
	iri := "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
	cmd := client.MsgAnchorCmd()
	args := []string{iri}
	args = append(args, commonFlags...)
	_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	cmd = client.MsgAttestCmd()

	testCases := []struct {
		name   string
		iri    string
		expErr bool
		errMsg string
	}{
		{
			name:   "valid",
			iri:    iri,
			expErr: false,
		},
		{
			name:   "no arg",
			iri:    "",
			expErr: true,
			errMsg: "iri is required",
		},
		{
			name:   "invalid iri",
			iri:    "noooo",
			expErr: true,
			errMsg: "invalid iri",
		},
		{
			name:   "bad extension",
			iri:    "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.png",
			expErr: true,
			errMsg: "invalid iri: invalid extension .png for graph data, expected .rdf",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			args := []string{tc.iri}
			args = append(args, commonFlags...)
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestDefineResolverCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	require := s.Require()

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
	}

	testCases := []struct {
		name        string
		resolverURL string
		expErr      bool
		errMsg      string
	}{
		{
			"empty url",
			"",
			true,
			"empty url",
		},
		{
			"invalid url",
			"abcd",
			true,
			"invalid URI",
		},
		{
			"valid test",
			"https://foo.bar",
			false,
			"",
		},
	}

	cmd := client.MsgDefineResolverCmd()
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			args := []string{tc.resolverURL}
			args = append(args, commonFlags...)
			_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg, err.Error())
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestRegisterResolverCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	require := s.Require()

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
	}

	_, ch := s.createIRIAndGraphHash([]byte("xyzabc123"))

	chs := &data.ContentHashes{ContentHashes: []*data.ContentHash{ch}}
	bz, err := val.ClientCtx.Codec.MarshalJSON(chs)
	require.NoError(err)

	filePath := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		errInRes bool
		errMsg   string
	}{
		{
			"empty args",
			[]string{},
			true,
			false,
			"accepts 2 arg(s), received 0",
		},
		{
			"invalid file path",
			[]string{fmt.Sprintf("%d", s.resolverID), "test.json"},
			true,
			false,
			"no such file or directory",
		},
		{
			"invalid resolver id",
			[]string{fmt.Sprintf("%d", 12345), filePath},
			false,
			true,
			"not found",
		},
		{
			"valid test",
			[]string{fmt.Sprintf("%d", s.resolverID), filePath},
			false,
			false,
			"",
		},
	}

	cmd := client.MsgRegisterResolverCmd()
	for _, tc := range testCases {
		args := tc.args
		s.Run(tc.name, func() {
			args := append(args, commonFlags...)
			res, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg, err.Error())
			} else {
				if tc.errInRes {
					require.Contains(res.String(), tc.errMsg)
				} else {
					require.NoError(err)
				}
			}
		})
	}
}
