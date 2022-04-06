package testsuite

import (
	"context"
	"crypto"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/data"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory testutil.FixtureFactory
	fixture        testutil.Fixture

	ctx         context.Context
	sdkCtx      sdk.Context
	msgClient   data.MsgClient
	queryClient data.QueryClient
	addr1       sdk.AccAddress
	addr2       sdk.AccAddress
	hash1       *data.ContentHash
	hash2       *data.ContentHash
}

func NewIntegrationTestSuite(fixtureFactory testutil.FixtureFactory) *IntegrationTestSuite {
	return &IntegrationTestSuite{fixtureFactory: fixtureFactory}
}

func (s *IntegrationTestSuite) SetupSuite() {
	require := s.Require()

	blockTime, err := time.Parse("2006-01-02", "2022-01-01")
	require.NoError(err)

	s.fixture = s.fixtureFactory.Setup()
	s.ctx = s.fixture.Context()
	s.sdkCtx = s.ctx.(types.Context).WithBlockTime(blockTime)
	s.msgClient = data.NewMsgClient(s.fixture.TxConn())
	s.queryClient = data.NewQueryClient(s.fixture.QueryConn())
	require.GreaterOrEqual(len(s.fixture.Signers()), 2)
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]

	content := []byte("xyzabc123")
	hash := crypto.BLAKE2b_256.New()
	_, err = hash.Write(content)
	require.NoError(err)
	digest := hash.Sum(nil)

	graphHash := &data.ContentHash_Graph{
		Hash:                      digest,
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}
	s.hash1 = &data.ContentHash{Graph: graphHash}

	rawHash := &data.ContentHash_Raw{
		Hash:            digest,
		DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		MediaType:       data.RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
	}
	s.hash2 = &data.ContentHash{Raw: rawHash}
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

func (s *IntegrationTestSuite) TestGraphScenario() {
	require := s.Require()

	iri, err := s.hash1.ToIRI()
	require.NoError(err)
	require.NotNil(iri)

	graphHash := s.hash1.GetGraph()

	// anchor some data
	anchorRes1, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.addr1.String(),
		Hash:   s.hash1,
	})
	require.NoError(err)
	require.NotNil(anchorRes1)
	require.Equal(iri, anchorRes1.Iri)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// anchoring same data twice is a no-op
	anchorRes2, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.addr1.String(),
		Hash:   s.hash1,
	})
	require.NoError(err)
	require.NotNil(anchorRes2)
	require.Equal(iri, anchorRes2.Iri)
	require.Equal(anchorRes1.Timestamp, anchorRes2.Timestamp)

	// can query data by iri
	dataByIRI, err := s.queryClient.ByIRI(s.ctx, &data.QueryByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(dataByIRI)
	require.NotNil(dataByIRI.Entry)
	require.Equal(anchorRes1.Timestamp, dataByIRI.Entry.Timestamp)

	// can query data by hash
	dataByHash, err := s.queryClient.ByHash(s.ctx, &data.QueryByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(dataByHash)
	require.NotNil(dataByHash.Entry)
	require.Equal(anchorRes1.Timestamp, dataByHash.Entry.Timestamp)

	// can query iri by hash
	iriByHash, err := s.queryClient.IRIByHash(s.ctx, &data.QueryIRIByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(iriByHash)
	require.Equal(iri, iriByHash.Iri)

	// can query hash by iri
	hashByIri, err := s.queryClient.HashByIRI(s.ctx, &data.QueryHashByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(hashByIri)
	require.Equal(s.hash1, hashByIri.ContentHash)

	// can query attestors by iri
	attestorsByIri, err := s.queryClient.AttestorsByIRI(s.ctx, &data.QueryAttestorsByIRIRequest{
		Iri: dataByIRI.Entry.Iri,
	})
	require.NoError(err)
	require.Empty(attestorsByIri.Attestors)

	// can query attestors by hash
	attestorsByHash, err := s.queryClient.AttestorsByHash(s.ctx, &data.QueryAttestorsByHashRequest{
		ContentHash: dataByIRI.Entry.Hash,
	})
	require.NoError(err)
	require.Empty(attestorsByHash.Attestors)

	// can attest to data
	_, err = s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.addr1.String(),
		Hashes:   []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)

	// attesting to the same data twice is a no-op
	attestRes, err := s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.addr1.String(),
		Hashes:   []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)
	require.Nil(attestRes.NewEntries)

	// can query attestors by iri
	attestorsByIri, err = s.queryClient.AttestorsByIRI(s.ctx, &data.QueryAttestorsByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.Len(attestorsByIri.Attestors, 1)
	require.Equal(s.addr1.String(), attestorsByIri.Attestors[0])

	// can query attestors by hash
	attestorsByHash, err = s.queryClient.AttestorsByHash(s.ctx, &data.QueryAttestorsByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.Len(attestorsByHash.Attestors, 1)
	require.Equal(s.addr1.String(), attestorsByHash.Attestors[0])

	// can query data by attestor
	byAttestors, err := s.queryClient.ByAttestor(s.ctx, &data.QueryByAttestorRequest{
		Attestor: s.addr1.String(),
	})
	require.NoError(err)
	require.NotNil(byAttestors)
	require.Len(byAttestors.Entries, 1)
	require.Equal(dataByIRI.Entry, byAttestors.Entries[0])

	// another attestor can attest
	_, err = s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.addr2.String(),
		Hashes:   []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)

	// can query attestors and get both attestations
	attestorsByIri, err = s.queryClient.AttestorsByIRI(s.ctx, &data.QueryAttestorsByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.Len(attestorsByIri.Attestors, 2)

	// loop through attestors as the order can vary
	attestors := make([]string, len(attestorsByIri.Attestors))
	for _, attestor := range attestorsByIri.Attestors {
		attestors = append(attestors, attestor)
	}
	require.Contains(attestors, s.addr1.String())
	require.Contains(attestors, s.addr2.String())
}

func (s *IntegrationTestSuite) TestRawDataScenario() {
	require := s.Require()

	iri, err := s.hash2.ToIRI()
	require.NoError(err)
	require.NotNil(iri)

	// anchor some data
	anchorRes1, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.addr1.String(),
		Hash:   s.hash2,
	})
	require.NoError(err)
	require.NotNil(anchorRes1)
	require.Equal(iri, anchorRes1.Iri)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// anchoring same data twice is a no-op
	anchorRes2, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.addr1.String(),
		Hash:   s.hash2,
	})
	require.NoError(err)
	require.NotNil(anchorRes2)
	require.Equal(iri, anchorRes2.Iri)
	require.Equal(anchorRes1.Timestamp, anchorRes2.Timestamp)

	// can query data by iri
	dataByIRI, err := s.queryClient.ByIRI(s.ctx, &data.QueryByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(dataByIRI)
	require.NotNil(dataByIRI.Entry)
	require.Equal(anchorRes1.Timestamp, dataByIRI.Entry.Timestamp)

	// can query data by hash
	dataByHash, err := s.queryClient.ByHash(s.ctx, &data.QueryByHashRequest{
		ContentHash: s.hash2,
	})
	require.NoError(err)
	require.NotNil(dataByHash)
	require.NotNil(dataByHash.Entry)
	require.Equal(anchorRes1.Timestamp, dataByHash.Entry.Timestamp)

	// can query iri by hash
	iriByHash, err := s.queryClient.IRIByHash(s.ctx, &data.QueryIRIByHashRequest{
		ContentHash: s.hash2,
	})
	require.NoError(err)
	require.NotNil(iriByHash)
	require.Equal(iri, iriByHash.Iri)

	// can query hash by iri
	hashByIri, err := s.queryClient.HashByIRI(s.ctx, &data.QueryHashByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(hashByIri)
	require.Equal(s.hash2, hashByIri.ContentHash)
}

func (s *IntegrationTestSuite) TestResolver() {
	require := s.Require()
	testUrl := "https://foo.bar"
	testData := []*data.ContentHash{s.hash1}

	iri, err := s.hash1.ToIRI()
	require.NoError(err)
	require.NotNil(iri)

	// can define a resolver
	res1, err := s.msgClient.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.addr1.String(),
		ResolverUrl: testUrl,
	})
	require.NoError(err)
	require.NotNil(res1)

	// can register content to a resolver
	res2, err := s.msgClient.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:    s.addr1.String(),
		ResolverId: res1.ResolverId,
		Data:       testData,
	})
	require.NoError(err)
	require.NotNil(res2)

	// can query resolvers by iri
	res3, err := s.queryClient.ResolversByIRI(s.ctx, &data.QueryResolversByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(res3)
	require.Equal([]string{testUrl}, res3.ResolverUrls)

	// can query resolvers by hash
	res4, err := s.queryClient.ResolversByHash(s.ctx, &data.QueryResolversByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(res4)
	require.Equal([]string{testUrl}, res4.ResolverUrls)

	// can query resolver info
	res5, err := s.queryClient.ResolverInfo(s.ctx, &data.QueryResolverInfoRequest{
		Url: testUrl,
	})
	require.NoError(err)
	require.NotNil(res5)
	require.Equal(s.addr1.String(), res5.Manager)
}
