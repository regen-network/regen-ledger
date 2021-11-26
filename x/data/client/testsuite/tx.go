package testsuite

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/data/client"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
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

	val := s.network.Validators[0]

	// create a new account
	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewValidator", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
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
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestTxAnchorData() {
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

	cmd := client.MsgAnchorDataCmd()
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

func (s *IntegrationTestSuite) TestTxSignData() {
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
	cmd := client.MsgAnchorDataCmd()
	args := []string{iri}
	args = append(args, commonFlags...)
	_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	cmd = client.MsgSignDataCmd()

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
			errMsg: "invalid iri: expected extension .rdf for graph data, got .png",
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
