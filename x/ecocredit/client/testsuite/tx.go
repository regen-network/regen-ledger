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

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
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
	validMetadata         = "hi"
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
