package testsuite

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/testutil/server"
	"github.com/regen-network/regen-ledger/x/data"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory server.FixtureFactory
	fixture        server.Fixture

	ctx         context.Context
	msgClient   data.MsgClient
	queryClient data.QueryClient
	addr1       sdk.AccAddress
	addr2       sdk.AccAddress
}

func NewIntegrationTestSuite(fixtureFactory server.FixtureFactory) *IntegrationTestSuite {
	return &IntegrationTestSuite{fixtureFactory: fixtureFactory}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()
	s.ctx = s.fixture.Context()
	s.msgClient = data.NewMsgClient(s.fixture.TxConn())
	s.queryClient = data.NewQueryClient(s.fixture.QueryConn())
	s.Require().GreaterOrEqual(len(s.fixture.Signers()), 2)
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

func (s *IntegrationTestSuite) TestScenario() {
	//testContent := []byte("xyzabc123")
	//mh, err := multihash.Sum(testContent, multihash.SHA2_256, -1)
	//s.Require().NoError(err)
	//cid := gocid.NewCidV1(gocid.Raw, mh)
	//
	//// anchor some data
	//cidBz := cid.Bytes()
	//anchorRes, err := s.msgClient.AnchorData(s.ctx, &data.MsgAnchorDataRequest{
	//	Sender: s.addr1.String(),
	//	Cid:    cidBz,
	//})
	//s.Require().NoError(err)
	//s.Require().NotNil(anchorRes)
	//
	//// can't anchor same data twice
	//_, err = s.msgClient.AnchorData(s.ctx, &data.MsgAnchorDataRequest{
	//	Sender: s.addr1.String(),
	//	Cid:    cidBz,
	//})
	//s.Require().Error(err)
	//
	//// can query data and get timestamp
	//queryRes, err := s.queryClient.ByCid(s.ctx, &data.QueryByCidRequest{Cid: cidBz})
	//s.Require().NoError(err)
	//s.Require().NotNil(queryRes)
	//s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	//s.Require().Empty(queryRes.Signers)
	//s.Require().Empty(queryRes.Content)
	//
	//// can sign data
	//_, err = s.msgClient.SignData(s.ctx, &data.MsgSignDataRequest{
	//	Signers: []string{s.addr1.String()},
	//	Cid:     cidBz,
	//})
	//s.Require().NoError(err)
	//
	//// can retrieve signature, same timestamp
	//// can query data and get timestamp
	//queryRes, err = s.queryClient.ByCid(s.ctx, &data.QueryByCidRequest{Cid: cidBz})
	//s.Require().NoError(err)
	//s.Require().NotNil(queryRes)
	//s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	//s.Require().Equal([]string{s.addr1.String()}, queryRes.Signers)
	//s.Require().Empty(queryRes.Content)
	//
	//// query data by signer
	//bySignerRes, err := s.queryClient.BySigner(s.ctx, &data.QueryBySignerRequest{
	//	Signer: s.addr1.String(),
	//})
	//s.Require().NoError(err)
	//s.Require().NotNil(bySignerRes)
	//s.Require().Contains(bySignerRes.Cids, cidBz)
	//
	//// can't store bad data
	//_, err = s.msgClient.StoreData(s.ctx, &data.MsgStoreDataRequest{
	//	Sender:  s.addr1.String(),
	//	Cid:     cidBz,
	//	Content: []byte("sgkjhsgouiyh"),
	//})
	//s.Require().Error(err)
	//
	//// can store good data
	//_, err = s.msgClient.StoreData(s.ctx, &data.MsgStoreDataRequest{
	//	Sender:  s.addr1.String(),
	//	Cid:     cidBz,
	//	Content: testContent,
	//})
	//s.Require().NoError(err)
	//
	//// can retrieve signature, same timestamp, and data
	//queryRes, err = s.queryClient.ByCid(s.ctx, &data.QueryByCidRequest{Cid: cidBz})
	//s.Require().NoError(err)
	//s.Require().NotNil(queryRes)
	//s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	//s.Require().Equal([]string{s.addr1.String()}, queryRes.Signers)
	//s.Require().Equal(testContent, queryRes.Content)
	//
	//// another signer can sign
	//_, err = s.msgClient.SignData(s.ctx, &data.MsgSignDataRequest{
	//	Signers: []string{s.addr2.String()},
	//	Cid:     cidBz,
	//})
	//s.Require().NoError(err)
	//
	//// query data by signer
	//bySignerRes, err = s.queryClient.BySigner(s.ctx, &data.QueryBySignerRequest{
	//	Signer: s.addr2.String(),
	//})
	//s.Require().NoError(err)
	//s.Require().NotNil(bySignerRes)
	//s.Require().Contains(bySignerRes.Cids, cidBz)
	//
	//// query all data and both signatures
	//queryRes, err = s.queryClient.ByCid(s.ctx, &data.QueryByCidRequest{Cid: cidBz})
	//s.Require().NoError(err)
	//s.Require().NotNil(queryRes)
	//s.Require().Equal(anchorRes.Timestamp, queryRes.Timestamp)
	//s.Require().Len(queryRes.Signers, 2)
	//s.Require().Contains(queryRes.Signers, s.addr1.String())
	//s.Require().Contains(queryRes.Signers, s.addr2.String())
	//s.Require().Equal(testContent, queryRes.Content)
}
