package testsuite

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/stretchr/testify/suite"
)

type AllowListEnabledTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewAllowListEnabledTestSuite(cfg network.Config) *AllowListEnabledTestSuite {
	return &AllowListEnabledTestSuite{cfg: cfg}
}

func (s *AllowListEnabledTestSuite) SetupSuite() {
	s.network = network.New(s.T(), s.cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *AllowListEnabledTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *AllowListEnabledTestSuite) TestCreateClass() {
	val := s.network.Validators[0]

	out, err := cli.ExecTestCLICmd(val.ClientCtx, client.TxCreateClassCmd(),
		[]string{
			val.Address.String(),
			"carbon",
			"AQ==",
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
		},
	)

	s.Require().NoError(err, out.String())

	// tx should fail with error unauthorized since we enabled `AllowlistEnabled` param & set `AllowedClassCreators` as `[]`.
	var res sdk.TxResponse
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	s.Require().Contains(res.RawLog, fmt.Sprintf("%s is not allowed to create credit classes", val.Address.String()))
}
