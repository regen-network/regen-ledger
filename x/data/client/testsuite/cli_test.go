package testsuite

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	gocid "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	clitestutil "github.com/regen-network/regen-ledger/testutil/cli"
	"github.com/regen-network/regen-ledger/testutil/network"
	datatypes "github.com/regen-network/regen-ledger/x/data"
	dataclient "github.com/regen-network/regen-ledger/x/data/client"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	testContent        []byte
	storedCid          gocid.Cid
	storedCidTimestamp *types.Timestamp
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
	clientCtx := val.ClientCtx

	testContent := []byte("xyzabc12345")
	mh, err := multihash.Sum(testContent, multihash.SHA2_256, -1)
	s.Require().NoError(err)
	cid := gocid.NewCidV1(gocid.Raw, mh)

	args := []string{
		val.Address.String(),
		cid.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	out, err := clitestutil.ExecTestCLICmd(clientCtx, dataclient.MsgAnchorDataCmd(), args)
	s.Require().NoError(err, out.String())

	txRes := &sdk.TxResponse{}
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), txRes), out.String())
	s.Require().Equal(uint32(0), txRes.Code)

	s.Require().NoError(s.network.WaitForNextBlock())
	s.testContent = testContent
	s.storedCid = cid
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestTxAnchorData() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testContent := []byte("xyzabc123")
	cid := s.getCid(testContent)

	var commonFlags = []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
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
			"wrong sender",
			append(
				[]string{
					"wrongSender",
					cid.String(),
				},
				commonFlags...,
			),
			true,
			"The specified item could not be found in the keyring",
			nil,
			0,
		},
		{
			"wrong cid",
			append(
				[]string{
					val.Address.String(),
					"wrongCid",
				},
				commonFlags...,
			),
			true,
			"selected encoding not supported",
			nil,
			0,
		},
		{
			"correct data",
			append(
				[]string{
					val.Address.String(),
					cid.String(),
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
		{
			"duplicate data",
			append(
				[]string{
					val.Address.String(),
					cid.String(),
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			18,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := dataclient.MsgAnchorDataCmd()

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
				s.Require().Error(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetAnchorDataByCID() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cid := s.storedCid

	testCases := []struct {
		name      string
		malleate  func()
		args      []string
		expectErr bool
		respType  proto.Message
	}{
		{
			"with non existed cid",
			func() {
				cid = s.getCid([]byte("xyzabc"))
			},
			[]string{
				cid.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			true,
			nil,
		},
		{
			"with correct data",
			func() {
				cid = s.storedCid
			},
			[]string{
				cid.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			false,
			&datatypes.QueryByCidResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			tc.malleate()
			cmd := dataclient.QueryByCidCmd()

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*datatypes.QueryByCidResponse)
				s.Require().NotNil(txResp)

				s.Require().Empty(txResp.Signers)
				s.Require().Empty(txResp.Content)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxSignData() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cid := s.storedCid

	var commonFlags = []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
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
			"wrong signer address",
			append(
				[]string{
					"wrongSigner",
					cid.String(),
				},
				commonFlags...,
			),
			true,
			"The specified item could not be found in the keyring",
			nil,
			0,
		},
		{
			"wrong cid",
			append(
				[]string{
					val.Address.String(),
					"wrongCid",
				},
				commonFlags...,
			),
			true,
			"selected encoding not supported",
			nil,
			0,
		},
		{
			"correct data",
			append(
				[]string{
					val.Address.String(),
					cid.String(),
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := dataclient.MsgSignDataCmd()

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Contains(out.String(), tc.expectErrMsg)
				s.Require().Error(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestTxStoreData() {
	val := s.network.Validators[0]
	val2 := s.network.Validators[1]
	clientCtx := val.ClientCtx

	cid := s.storedCid
	base64Encoded := base64.StdEncoding.EncodeToString(s.testContent)

	var commonFlags = []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
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
			"store bad data",
			append(
				[]string{
					val.Address.String(),
					cid.String(),
					"wrongBase64",
				},
				commonFlags...,
			),
			true,
			"illegal base64 data at input byte 8",
			nil,
			0,
		},
		{
			"wrong signer",
			append(
				[]string{
					val2.Address.String(),
					cid.String(),
					base64Encoded,
				},
				commonFlags...,
			),
			true,
			"The specified item could not be found in the keyring",
			nil,
			0,
		},
		{
			"correct data",
			append(
				[]string{
					val.Address.String(),
					cid.String(),
					base64Encoded,
				},
				commonFlags...,
			),
			false,
			"",
			&sdk.TxResponse{},
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := dataclient.MsgStoreDataCmd()

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) getCid(content []byte) gocid.Cid {
	mh, err := multihash.Sum(content, multihash.SHA2_256, -1)
	s.Require().NoError(err)
	return gocid.NewCidV1(gocid.Raw, mh)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
