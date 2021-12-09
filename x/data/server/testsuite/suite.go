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
	require := s.Require()
	require.NoError(err)
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
	require.NoError(err)
	require.NotNil(anchorRes)

	// anchoring same data twice is a no-op
	_, err = s.msgClient.AnchorData(s.ctx, &data.MsgAnchorData{
		Sender: s.addr1.String(),
		Hash:   contentHash,
	})
	require.NoError(err)

	// can query data and get timestamp
	iri, err := contentHash.ToIRI()
	require.NoError(err)
	queryRes, err := s.queryClient.ByIRI(s.ctx, &data.QueryByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(queryRes)
	require.NotNil(queryRes.Entry)
	ts := queryRes.Entry.Timestamp
	require.NotNil(ts)

	signerRes, err := s.queryClient.Signers(s.ctx, &data.QuerySignersRequest{Iri: queryRes.Entry.Iri, Pagination: nil})
	require.NoError(err)
	require.Empty(signerRes.Signers)
	iri, err = graphHash.ToIRI()
	require.NoError(err)
	require.Equal(iri, queryRes.Entry.Iri)

	// can sign data
	_, err = s.msgClient.SignData(s.ctx, &data.MsgSignData{
		Signers: []string{s.addr1.String()},
		Hash:    graphHash,
	})
	require.NoError(err)

	// can retrieve signature, same timestamp
	// can query data and get timestamp
	iri, err = contentHash.ToIRI()
	require.NoError(err)
	queryRes, err = s.queryClient.ByIRI(s.ctx, &data.QueryByIRIRequest{Iri: iri})
	require.NoError(err)
	require.NotNil(queryRes)
	require.Equal(ts, queryRes.Entry.Timestamp) // ensure timestamp is equal to the original
	signerRes, err = s.queryClient.Signers(s.ctx, &data.QuerySignersRequest{Iri: iri, Pagination: nil})
	require.NoError(err)
	require.Len(signerRes.Signers, 1)
	require.Equal(s.addr1.String(), signerRes.Signers[0])

	// query data by signer
	bySignerRes, err := s.queryClient.BySigner(s.ctx, &data.QueryBySignerRequest{
		Signer: s.addr1.String(),
	})
	require.NoError(err)
	require.NotNil(bySignerRes)
	require.Len(bySignerRes.Entries, 1)
	require.Equal(queryRes.Entry, bySignerRes.Entries[0])

	// another signer can sign
	_, err = s.msgClient.SignData(s.ctx, &data.MsgSignData{
		Signers: []string{s.addr2.String()},
		Hash:    graphHash,
	})
	require.NoError(err)

	// query data by signer
	bySignerRes, err = s.queryClient.BySigner(s.ctx, &data.QueryBySignerRequest{
		Signer: s.addr2.String(),
	})
	require.NoError(err)
	require.NotNil(bySignerRes)
	require.Len(bySignerRes.Entries, 1)
	require.Equal(contentHash, bySignerRes.Entries[0].Hash)

	// query and get both signatures
	iri, err = contentHash.ToIRI()
	require.NoError(err)
	queryRes, err = s.queryClient.ByIRI(s.ctx, &data.QueryByIRIRequest{Iri: iri})
	require.NoError(err)
	require.NotNil(queryRes)
	require.Equal(ts, queryRes.Entry.Timestamp)

	iri2, err := contentHash.ToIRI()
	require.NoError(err)
	signerRes, err = s.queryClient.Signers(s.ctx, &data.QuerySignersRequest{Iri: iri2, Pagination: nil})
	require.NoError(err)
	require.Len(signerRes.Signers, 2)
	signers := make([]string, len(signerRes.Signers))
	for _, signer := range signerRes.Signers {
		signers = append(signers, signer)
	}
	require.Contains(signers, s.addr1.String())
	require.Contains(signers, s.addr2.String())
}

func (s *IntegrationTestSuite) TestRawDataScenario() {
	testContent := []byte("19sdgh23t7sdghasf98sf")
	hash := crypto.BLAKE2b_256.New()
	_, err := hash.Write(testContent)
	require := s.Require()
	require.NoError(err)
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
	require.NoError(err)
	require.NotNil(anchorRes)

	// anchoring same data twice is a no-op
	_, err = s.msgClient.AnchorData(s.ctx, &data.MsgAnchorData{
		Sender: s.addr1.String(),
		Hash:   contentHash,
	})
	require.NoError(err)

	// can query data and get timestamp
	iri, err := contentHash.ToIRI()
	require.NoError(err)
	queryRes, err := s.queryClient.ByIRI(s.ctx, &data.QueryByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(queryRes)
	require.NotNil(queryRes.Entry)
	ts := queryRes.Entry.Timestamp
	require.NotNil(ts)

	signerRes, err := s.queryClient.Signers(s.ctx, &data.QuerySignersRequest{Iri: queryRes.Entry.Iri, Pagination: nil})
	require.Empty(signerRes.Signers)

	// can retrieve same timestamp, and data
	iri, err = contentHash.ToIRI()
	require.NoError(err)
	queryRes, err = s.queryClient.ByIRI(s.ctx, &data.QueryByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(queryRes)
	require.Equal(ts, queryRes.Entry.Timestamp)
}
