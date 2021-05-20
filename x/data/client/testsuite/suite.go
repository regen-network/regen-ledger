package testsuite

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/types"
	gocid "github.com/ipfs/go-cid"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	testContent        []byte
	storedCid          gocid.Cid
	storedCidTimestamp *types.Timestamp
	signer             sdk.AccAddress
}

// TODO call NewIntegrationTestSuite to set cfg field of IntegrationTestSuite
// remove cfg := network.DefaultConfig()
// replace cfg.NumValidators = 2 with s.cfg.NumValidators = 2
// replace s.network = network.New(s.T(), cfg) with s.network = network.New(s.T(), s.cfg)
// remove TestIntegrationTestSuite func
func (s *IntegrationTestSuite) SetupSuite() {
	//s.T().Log("setting up integration test suite")
	//
	//cfg := network.DefaultConfig()
	//cfg.NumValidators = 2
	//
	//s.cfg = cfg
	//s.network = network.New(s.T(), cfg)
	//
	//_, err := s.network.WaitForHeight(1)
	//s.Require().NoError(err)
	//
	//val := s.network.Validators[0]
	//clientCtx := val.ClientCtx
	//
	//testContent := []byte("xyzabc12345")
	//cid := s.getCid(testContent)
	//
	//commonFlags := []string{
	//	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	//	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	//	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	//}
	//
	//args := append(
	//	[]string{
	//		val.Address.String(),
	//		cid.String(),
	//	},
	//	commonFlags...,
	//)
	//
	//_, err = clitestutil.ExecTestCLICmd(clientCtx, dataclient.MsgAnchorDataCmd(), args)
	//
	//s.Require().NoError(s.network.WaitForNextBlock())
	//s.testContent = testContent
	//s.storedCid = cid
	//
	//// create a new account
	//info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewValidator", keyring.English, sdk.FullFundraiserPath, hd.Secp256k1)
	//s.Require().NoError(err)
	//
	//account := sdk.AccAddress(info.GetPubKey().Address())
	//_, err = banktestutil.MsgSendExec(
	//	val.ClientCtx,
	//	val.Address,
	//	account,
	//	sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(200))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	//	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	//	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	//)
	//
	//s.Require().NoError(err)
	//
	//// add a signer
	//args = append(
	//	[]string{
	//		account.String(),
	//		cid.String(),
	//	},
	//	commonFlags...,
	//)
	//_, err = clitestutil.ExecTestCLICmd(clientCtx, dataclient.MsgSignDataCmd(), args)
	//
	//s.signer = account
}

func (s *IntegrationTestSuite) TearDownSuite() {
	//s.T().Log("tearing down integration test suite")
	//s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestTxAnchorData() {
	//val := s.network.Validators[0]
	//clientCtx := val.ClientCtx
	//
	//testContent := []byte("xyzabc123")
	//cid := s.getCid(testContent)
	//
	//var commonFlags = []string{
	//	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	//	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	//	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	//}
	//
	//testCases := []struct {
	//	name         string
	//	args         []string
	//	expectErr    bool
	//	expectErrMsg string
	//	respType     proto.Message
	//	expectedCode uint32
	//}{
	//	{
	//		"wrong sender",
	//		append(
	//			[]string{
	//				"wrongSender",
	//				cid.String(),
	//			},
	//			commonFlags...,
	//		),
	//		true,
	//		"The specified item could not be found in the keyring",
	//		nil,
	//		0,
	//	},
	//	{
	//		"wrong cid",
	//		append(
	//			[]string{
	//				val.Address.String(),
	//				"wrongCid",
	//			},
	//			commonFlags...,
	//		),
	//		true,
	//		"selected encoding not supported",
	//		nil,
	//		0,
	//	},
	//	{
	//		"correct data",
	//		append(
	//			[]string{
	//				val.Address.String(),
	//				cid.String(),
	//			},
	//			commonFlags...,
	//		),
	//		false,
	//		"",
	//		&sdk.TxResponse{},
	//		0,
	//	},
	//	{
	//		"duplicate data",
	//		append(
	//			[]string{
	//				val.Address.String(),
	//				cid.String(),
	//			},
	//			commonFlags...,
	//		),
	//		false,
	//		"",
	//		&sdk.TxResponse{},
	//		18,
	//	},
	//}
	//
	//for _, tc := range testCases {
	//	tc := tc
	//
	//	s.Run(tc.name, func() {
	//		cmd := dataclient.MsgAnchorDataCmd()
	//
	//		out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
	//		if tc.expectErr {
	//			s.Require().Contains(out.String(), tc.expectErrMsg)
	//		} else {
	//			s.Require().NoError(err, out.String())
	//			s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
	//
	//			txResp := tc.respType.(*sdk.TxResponse)
	//			s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
	//		}
	//	})
	//}
}

func (s *IntegrationTestSuite) TestGetAnchorDataByCID() {
	//val := s.network.Validators[0]
	//clientCtx := val.ClientCtx
	//
	//testCases := []struct {
	//	name         string
	//	args         []string
	//	expectErr    bool
	//	expectErrMsg string
	//	resp         datatypes.QueryByCidResponse
	//}{
	//	{
	//		"with non existed cid",
	//		[]string{
	//			s.getCid([]byte("xyzabc")).String(),
	//			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	//		},
	//		true,
	//		"CID not found",
	//		datatypes.QueryByCidResponse{},
	//	},
	//	{
	//		"with correct data",
	//		[]string{
	//			s.storedCid.String(),
	//			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	//		},
	//		false,
	//		"",
	//		datatypes.QueryByCidResponse{},
	//	},
	//}
	//
	//for _, tc := range testCases {
	//	s.Run(tc.name, func() {
	//		cmd := dataclient.QueryByCidCmd()
	//
	//		out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
	//		if tc.expectErr {
	//			s.Require().Contains(out.String(), tc.expectErrMsg)
	//
	//		} else {
	//			s.Require().NoError(err, out.String())
	//			s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &tc.resp), out.String())
	//
	//			txResp := tc.resp
	//			s.Require().NotNil(txResp)
	//
	//			s.Require().Equal([]string{s.signer.String()}, txResp.Signers)
	//			s.Require().Empty(txResp.Content)
	//		}
	//	})
	//}
}

func (s *IntegrationTestSuite) TestTxSignData() {
	//val := s.network.Validators[0]
	//clientCtx := val.ClientCtx
	//
	//cid := s.storedCid
	//
	//var commonFlags = []string{
	//	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	//	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	//	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	//}
	//
	//testCases := []struct {
	//	name         string
	//	args         []string
	//	expectErr    bool
	//	expectErrMsg string
	//	respType     proto.Message
	//	expectedCode uint32
	//}{
	//	{
	//		"wrong signer address",
	//		append(
	//			[]string{
	//				"wrongSigner",
	//				cid.String(),
	//			},
	//			commonFlags...,
	//		),
	//		true,
	//		"The specified item could not be found in the keyring",
	//		nil,
	//		0,
	//	},
	//	{
	//		"wrong cid",
	//		append(
	//			[]string{
	//				val.Address.String(),
	//				"wrongCid",
	//			},
	//			commonFlags...,
	//		),
	//		true,
	//		"selected encoding not supported",
	//		nil,
	//		0,
	//	},
	//	{
	//		"correct data",
	//		append(
	//			[]string{
	//				val.Address.String(),
	//				cid.String(),
	//			},
	//			commonFlags...,
	//		),
	//		false,
	//		"",
	//		&sdk.TxResponse{},
	//		0,
	//	},
	//}
	//
	//for _, tc := range testCases {
	//	tc := tc
	//
	//	s.Run(tc.name, func() {
	//		cmd := dataclient.MsgSignDataCmd()
	//
	//		out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
	//		if tc.expectErr {
	//			s.Require().Contains(out.String(), tc.expectErrMsg)
	//		} else {
	//			s.Require().NoError(err, out.String())
	//			s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
	//
	//			txResp := tc.respType.(*sdk.TxResponse)
	//			s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
	//		}
	//	})
	//}
}

func (s *IntegrationTestSuite) TestTxStoreData() {
	//val := s.network.Validators[0]
	//val2 := s.network.Validators[1]
	//clientCtx := val.ClientCtx
	//
	//cid := s.storedCid
	//base64Encoded := base64.StdEncoding.EncodeToString(s.testContent)
	//
	//var commonFlags = []string{
	//	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	//	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	//	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	//}
	//
	//testCases := []struct {
	//	name         string
	//	args         []string
	//	expectErr    bool
	//	expectErrMsg string
	//	respType     proto.Message
	//	expectedCode uint32
	//}{
	//	{
	//		"store bad data",
	//		append(
	//			[]string{
	//				val.Address.String(),
	//				cid.String(),
	//				"wrongBase64",
	//			},
	//			commonFlags...,
	//		),
	//		true,
	//		"illegal base64 data",
	//		nil,
	//		0,
	//	},
	//	{
	//		"wrong signer",
	//		append(
	//			[]string{
	//				val2.Address.String(),
	//				cid.String(),
	//				base64Encoded,
	//			},
	//			commonFlags...,
	//		),
	//		true,
	//		"The specified item could not be found in the keyring",
	//		nil,
	//		0,
	//	},
	//	{
	//		"correct data",
	//		append(
	//			[]string{
	//				val.Address.String(),
	//				cid.String(),
	//				base64Encoded,
	//			},
	//			commonFlags...,
	//		),
	//		false,
	//		"",
	//		&sdk.TxResponse{},
	//		0,
	//	},
	//}
	//
	//for _, tc := range testCases {
	//	tc := tc
	//
	//	s.Run(tc.name, func() {
	//		cmd := dataclient.MsgStoreDataCmd()
	//
	//		out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
	//		if tc.expectErr {
	//			s.Require().Contains(out.String(), tc.expectErrMsg)
	//		} else {
	//			s.Require().NoError(err, out.String())
	//			s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
	//
	//			txResp := tc.respType.(*sdk.TxResponse)
	//			s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
	//		}
	//	})
	//}
}

//func (s *IntegrationTestSuite) getCid(content []byte) gocid.Cid {
//	mh, err := multihash.Sum(content, multihash.SHA2_256, -1)
//	s.Require().NoError(err)
//	return gocid.NewCidV1(gocid.Raw, mh)
//}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
