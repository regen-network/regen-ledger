package testsuite

import (
	"context"

	gocid "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/testutil/server"
	"github.com/regen-network/regen-ledger/x/data"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixture server.Fixture

	ctx         context.Context
	msgClient   data.MsgClient
	queryClient data.QueryClient
	from        sdk.AccAddress
}

func NewIntegrationTestSuite(fixture server.Fixture) *IntegrationTestSuite {
	return &IntegrationTestSuite{fixture: fixture}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture.Setup()
	s.ctx = s.fixture.Context()
	s.msgClient = data.NewMsgClient(s.fixture.TxConn())
	s.queryClient = data.NewQueryClient(s.fixture.QueryConn())
	s.from = s.fixture.Signers()[0]
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

func (s *IntegrationTestSuite) TestScenario() {
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
