package testsuite

import (
	"context"
	"crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/data"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory testutil.FixtureFactory
	fixture        testutil.Fixture

	ctx         context.Context
	msgClient   data.MsgClient
	queryClient data.QueryClient
	addr1       sdk.AccAddress
	addr2       sdk.AccAddress
}

func NewIntegrationTestSuite(fixtureFactory testutil.FixtureFactory) *IntegrationTestSuite {
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

func (s *IntegrationTestSuite) TestGraphScenario() {
	testContent := []byte("xyzabc123")
	hash := crypto.BLAKE2b_256.New()
	_, err := hash.Write(testContent)
	s.Require().NoError(err)
	digest := hash.Sum(nil)
	graphHash := &data.ContentHash_Graph{
		Hash:                      digest,
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}
	contentHash := &data.ContentHash{Sum: &data.ContentHash_Graph_{Graph: graphHash}}

	//// anchor some data
	anchorRes, err := s.msgClient.AnchorData(s.ctx, &data.MsgAnchorData{
		Sender: s.addr1.String(),
		Hash:   contentHash,
	})
	s.Require().NoError(err)
	s.Require().NotNil(anchorRes)

	// anchoring same data twice is a no-op
	_, err = s.msgClient.AnchorData(s.ctx, &data.MsgAnchorData{
		Sender: s.addr1.String(),
		Hash:   contentHash,
	})
	s.Require().NoError(err)

	// can query data and get timestamp
	queryRes, err := s.queryClient.ByHash(s.ctx, &data.QueryByHashRequest{
		Hash: contentHash,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().NotNil(queryRes.Entry)
	ts := queryRes.Entry.Timestamp
	s.Require().NotNil(ts)
	s.Require().Empty(queryRes.Entry.Signers)
	iri, err := graphHash.ToIRI()
	s.Require().NoError(err)
	s.Require().Equal(iri, queryRes.Entry.Iri)

	// can sign data
	_, err = s.msgClient.SignData(s.ctx, &data.MsgSignData{
		Signers: []string{s.addr1.String()},
		Hash:    graphHash,
	})
	s.Require().NoError(err)

	// can retrieve signature, same timestamp
	// can query data and get timestamp
	queryRes, err = s.queryClient.ByHash(s.ctx, &data.QueryByHashRequest{Hash: contentHash})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(ts, queryRes.Entry.Timestamp)
	s.Require().Len(queryRes.Entry.Signers, 1)
	s.Require().Equal(s.addr1.String(), queryRes.Entry.Signers[0].Signer)

	// query data by signer
	bySignerRes, err := s.queryClient.BySigner(s.ctx, &data.QueryBySignerRequest{
		Signer: s.addr1.String(),
	})
	s.Require().NoError(err)
	s.Require().NotNil(bySignerRes)
	s.Require().Len(bySignerRes.Entries, 1)
	s.Require().Equal(queryRes.Entry, bySignerRes.Entries[0])

	// another signer can sign
	_, err = s.msgClient.SignData(s.ctx, &data.MsgSignData{
		Signers: []string{s.addr2.String()},
		Hash:    graphHash,
	})
	s.Require().NoError(err)

	// query data by signer
	bySignerRes, err = s.queryClient.BySigner(s.ctx, &data.QueryBySignerRequest{
		Signer: s.addr2.String(),
	})
	s.Require().NoError(err)
	s.Require().NotNil(bySignerRes)
	s.Require().Len(bySignerRes.Entries, 1)
	s.Require().Equal(contentHash, bySignerRes.Entries[0].Hash)

	// query and get both signatures
	queryRes, err = s.queryClient.ByHash(s.ctx, &data.QueryByHashRequest{Hash: contentHash})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(ts, queryRes.Entry.Timestamp)
	s.Require().Len(queryRes.Entry.Signers, 2)
	signers := make([]string, len(queryRes.Entry.Signers))
	for _, signer := range queryRes.Entry.Signers {
		signers = append(signers, signer.Signer)
	}
	s.Require().Contains(signers, s.addr1.String())
	s.Require().Contains(signers, s.addr2.String())
}

func (s *IntegrationTestSuite) TestRawDataScenario() {
	testContent := []byte("19sdgh23t7sdghasf98sf")
	hash := crypto.BLAKE2b_256.New()
	_, err := hash.Write(testContent)
	s.Require().NoError(err)
	digest := hash.Sum(nil)
	rawHash := &data.ContentHash_Raw{
		Hash:            digest,
		DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		MediaType:       data.MediaType_MEDIA_TYPE_UNSPECIFIED,
	}
	contentHash := &data.ContentHash{Sum: &data.ContentHash_Raw_{Raw: rawHash}}

	//// anchor some data
	anchorRes, err := s.msgClient.AnchorData(s.ctx, &data.MsgAnchorData{
		Sender: s.addr1.String(),
		Hash:   contentHash,
	})
	s.Require().NoError(err)
	s.Require().NotNil(anchorRes)

	// anchoring same data twice is a no-op
	_, err = s.msgClient.AnchorData(s.ctx, &data.MsgAnchorData{
		Sender: s.addr1.String(),
		Hash:   contentHash,
	})
	s.Require().NoError(err)

	// can query data and get timestamp
	queryRes, err := s.queryClient.ByHash(s.ctx, &data.QueryByHashRequest{
		Hash: contentHash,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().NotNil(queryRes.Entry)
	ts := queryRes.Entry.Timestamp
	s.Require().NotNil(ts)
	s.Require().Empty(queryRes.Entry.Signers)

	// can retrieve same timestamp, and data
	queryRes, err = s.queryClient.ByHash(s.ctx, &data.QueryByHashRequest{
		Hash: contentHash,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryRes)
	s.Require().Equal(ts, queryRes.Entry.Timestamp)
}
