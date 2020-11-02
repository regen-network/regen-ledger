package client_test

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gocid "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	from        sdk.AccAddress
	msgClient   data.MsgClient
	queryClient data.QueryClient
	ctx         context.Context
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	cfg.NumValidators = 1

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

type ClientTxConn struct {
	ClientContext client.Context
	FlagSet       *pflag.FlagSet
}

func (c ClientTxConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	req, ok := args.(sdk.MsgRequest)
	if !ok {
		return fmt.Errorf("%T should implement %T", args, (*sdk.MsgRequest)(nil))
	}

	err := req.ValidateBasic()
	if err != nil {
		return err
	}

	msg := sdk.ServiceMsg{
		MethodName: method,
		Request:    req,
	}

	return clienttx.GenerateOrBroadcastTxCLI(c.ClientContext, c.FlagSet, msg)
}

func (c ClientTxConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("unsupported")
}

var _ grpc.ClientConnInterface = ClientTxConn{}

func (s *IntegrationTestSuite) TestAnchorData() {
	testContent := []byte("xyzabc123")
	mh, err := multihash.Sum(testContent, multihash.SHA2_256, -1)
	s.Require().NoError(err)
	cid := gocid.NewCidV1(gocid.Raw, mh)

	// anchor some data
	cidBz := cid.Bytes()
	anchorRes, err := s.msgClient.AnchorData(s.ctx, &data.MsgAnchorDataRequest{
		Sender: s.from.String(),
		Cid:    cidBz,
	})
	s.Require().NoError(err)
	s.Require().NotNil(anchorRes)

	// can't anchor same data twice
	_, err = s.msgClient.AnchorData(s.ctx, &data.MsgAnchorDataRequest{
		Sender: s.from.String(),
		Cid:    cidBz,
	})
	s.Require().Error(err)

	// can query data and get timestamp
	queryRes, err := s.queryClient.Data(s.ctx, &data.QueryDataRequest{Cid: cidBz})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	s.Require().Empty(queryRes.Signers)
	s.Require().Empty(queryRes.Content)

	// can sign data
	_, err = s.msgClient.SignData(s.ctx, &data.MsgSignDataRequest{
		Signers: []string{s.from.String()},
		Cid:     cidBz,
	})
	s.Require().NoError(err)

	// can retrieve signature, same timestamp
	// can query data and get timestamp
	queryRes, err = s.queryClient.Data(s.ctx, &data.QueryDataRequest{Cid: cidBz})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	s.Require().Equal([]string{s.from.String()}, queryRes.Signers)
	s.Require().Empty(queryRes.Content)

	// can't store bad data
	_, err = s.msgClient.StoreData(s.ctx, &data.MsgStoreDataRequest{
		Sender:  s.from.String(),
		Cid:     cidBz,
		Content: []byte("sgkjhsgouiyh"),
	})
	s.Require().Error(err)

	// can store good data
	_, err = s.msgClient.StoreData(s.ctx, &data.MsgStoreDataRequest{
		Sender:  s.from.String(),
		Cid:     cidBz,
		Content: testContent,
	})
	s.Require().NoError(err)

	// can retrieve signature, same timestamp, and data
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	s.Require().Equal([]string{s.from.String()}, queryRes.Signers)
	s.Require().Equal(testContent, queryRes.Content)
}
