package testsuite

import (
	"context"
	"crypto"
	"encoding/base64"
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
		Sender:      s.addr1.String(),
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(anchorRes1)
	require.Equal(iri, anchorRes1.Iri)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// anchoring same data twice is a no-op
	anchorRes2, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.addr1.String(),
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(anchorRes2)
	require.Equal(iri, anchorRes2.Iri)
	require.Equal(anchorRes1.Timestamp, anchorRes2.Timestamp)

	// can query anchored data entry by iri
	anchorByIRI, err := s.queryClient.AnchorByIRI(s.ctx, &data.QueryAnchorByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(anchorByIRI)
	require.NotNil(anchorByIRI.Anchor)
	require.Equal(anchorRes1.Timestamp, anchorByIRI.Anchor.Timestamp)

	// can query anchored data entry by hash
	anchorByHash, err := s.queryClient.AnchorByHash(s.ctx, &data.QueryAnchorByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(anchorByHash)
	require.NotNil(anchorByHash.Anchor)
	require.Equal(anchorRes1.Timestamp, anchorByHash.Anchor.Timestamp)

	// can convert hash to iri
	hashToIri, err := s.queryClient.ConvertHashToIRI(s.ctx, &data.ConvertHashToIRIRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(hashToIri)
	require.Equal(iri, hashToIri.Iri)

	// can convert graph hash properties to iri
	graphHashToIri, err := s.queryClient.ConvertGraphHashToIRI(s.ctx, &data.ConvertGraphHashToIRIRequest{
		Hash:                      base64.StdEncoding.EncodeToString(s.hash1.Graph.Hash),
		DigestAlgorithm:           s.hash1.Graph.DigestAlgorithm,
		CanonicalizationAlgorithm: s.hash1.Graph.CanonicalizationAlgorithm,
		MerkleTree:                s.hash1.Graph.MerkleTree,
	})
	require.NoError(err)
	require.NotNil(graphHashToIri)
	require.Equal(iri, graphHashToIri.Iri)

	// can convert iri to hash
	iriToHash, err := s.queryClient.ConvertIRIToHash(s.ctx, &data.ConvertIRIToHashRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(iriToHash)
	require.Equal(s.hash1, iriToHash.ContentHash)

	// can query attestor entries by iri
	attestorsByIri, err := s.queryClient.AttestorsByIRI(s.ctx, &data.QueryAttestorsByIRIRequest{
		Iri: anchorByIRI.Anchor.Iri,
	})
	require.NoError(err)
	require.Empty(attestorsByIri.Attestors)

	// can query attestor entries by hash
	attestorsByHash, err := s.queryClient.AttestorsByHash(s.ctx, &data.QueryAttestorsByHashRequest{
		ContentHash: anchorByIRI.Anchor.ContentHash,
	})
	require.NoError(err)
	require.Empty(attestorsByHash.Attestors)

	// can attest to data
	_, err = s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr1.String(),
		ContentHashes: []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)

	// attesting to the same data twice is a no-op
	attestRes, err := s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr1.String(),
		ContentHashes: []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)
	require.Nil(attestRes.NewEntries)

	// can query attestor entries by iri
	attestorsByIri, err = s.queryClient.AttestorsByIRI(s.ctx, &data.QueryAttestorsByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.Len(attestorsByIri.Attestors, 1)
	require.Equal(s.addr1.String(), attestorsByIri.Attestors[0].Attestor)

	// can query attestor entries by hash
	attestorsByHash, err = s.queryClient.AttestorsByHash(s.ctx, &data.QueryAttestorsByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.Len(attestorsByHash.Attestors, 1)
	require.Equal(s.addr1.String(), attestorsByHash.Attestors[0].Attestor)

	// can query anchored data entries by attestor
	anchorsByAttestor, err := s.queryClient.AnchorsByAttestor(s.ctx, &data.QueryAnchorsByAttestorRequest{
		Attestor: s.addr1.String(),
	})
	require.NoError(err)
	require.NotNil(anchorsByAttestor)
	require.Len(anchorsByAttestor.Anchors, 1)
	require.Equal(anchorByIRI.Anchor, anchorsByAttestor.Anchors[0])

	// another attestor can attest
	_, err = s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr2.String(),
		ContentHashes: []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)

	// can query attestor entries and get both attestations
	attestorsByIri, err = s.queryClient.AttestorsByIRI(s.ctx, &data.QueryAttestorsByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.Len(attestorsByIri.Attestors, 2)

	// order may vary from query response
	if attestorsByIri.Attestors[0].Attestor == s.addr1.String() {
		require.Equal(attestorsByIri.Attestors[0].Attestor, s.addr1.String())
		require.Equal(attestorsByIri.Attestors[1].Attestor, s.addr2.String())
	} else {
		require.Equal(attestorsByIri.Attestors[0].Attestor, s.addr2.String())
		require.Equal(attestorsByIri.Attestors[1].Attestor, s.addr1.String())
	}
}

func (s *IntegrationTestSuite) TestRawDataScenario() {
	require := s.Require()

	iri, err := s.hash2.ToIRI()
	require.NoError(err)
	require.NotNil(iri)

	// anchor some data
	anchorRes1, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.addr1.String(),
		ContentHash: s.hash2,
	})
	require.NoError(err)
	require.NotNil(anchorRes1)
	require.Equal(iri, anchorRes1.Iri)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// anchoring same data twice is a no-op
	anchorRes2, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.addr1.String(),
		ContentHash: s.hash2,
	})
	require.NoError(err)
	require.NotNil(anchorRes2)
	require.Equal(iri, anchorRes2.Iri)
	require.Equal(anchorRes1.Timestamp, anchorRes2.Timestamp)

	// can query anchored data entry by iri
	anchorByIRI, err := s.queryClient.AnchorByIRI(s.ctx, &data.QueryAnchorByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(anchorByIRI)
	require.NotNil(anchorByIRI.Anchor)
	require.Equal(anchorRes1.Timestamp, anchorByIRI.Anchor.Timestamp)

	// can query anchored data entry by hash
	anchorByHash, err := s.queryClient.AnchorByHash(s.ctx, &data.QueryAnchorByHashRequest{
		ContentHash: s.hash2,
	})
	require.NoError(err)
	require.NotNil(anchorByHash)
	require.NotNil(anchorByHash.Anchor)
	require.Equal(anchorRes1.Timestamp, anchorByHash.Anchor.Timestamp)

	// can convert hash to iri
	hashToIri, err := s.queryClient.ConvertHashToIRI(s.ctx, &data.ConvertHashToIRIRequest{
		ContentHash: s.hash2,
	})
	require.NoError(err)
	require.NotNil(hashToIri)
	require.Equal(iri, hashToIri.Iri)

	// can convert raw hash properties to iri
	rawHashToIri, err := s.queryClient.ConvertRawHashToIRI(s.ctx, &data.ConvertRawHashToIRIRequest{
		Hash:            base64.StdEncoding.EncodeToString(s.hash2.Raw.Hash),
		DigestAlgorithm: s.hash2.Raw.DigestAlgorithm,
		MediaType:       s.hash2.Raw.MediaType,
	})
	require.NoError(err)
	require.NotNil(rawHashToIri)
	require.Equal(iri, rawHashToIri.Iri)

	// can convert iri to hash
	iriToHash, err := s.queryClient.ConvertIRIToHash(s.ctx, &data.ConvertIRIToHashRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(iriToHash)
	require.Equal(s.hash2, iriToHash.ContentHash)
}

func (s *IntegrationTestSuite) TestResolver() {
	require := s.Require()
	testUrl := "https://foo.bar"
	hashes := []*data.ContentHash{s.hash1}

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
		Manager:       s.addr1.String(),
		ResolverId:    res1.ResolverId,
		ContentHashes: hashes,
	})
	require.NoError(err)
	require.NotNil(res2)

	// can query resolver
	res3, err := s.queryClient.Resolver(s.ctx, &data.QueryResolverRequest{
		Id: res1.ResolverId,
	})
	require.NoError(err)
	require.NotNil(res3)
	require.Equal(s.addr1.String(), res3.Resolver.Manager)
	require.Equal(testUrl, res3.Resolver.Url)

	// can query resolvers by iri
	res4, err := s.queryClient.ResolversByIRI(s.ctx, &data.QueryResolversByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(res4)
	require.Equal(s.addr1.String(), res4.Resolvers[0].Manager)
	require.Equal(testUrl, res4.Resolvers[0].Url)

	// can query resolvers by hash
	res5, err := s.queryClient.ResolversByHash(s.ctx, &data.QueryResolversByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(res5)
	require.Equal(s.addr1.String(), res5.Resolvers[0].Manager)
	require.Equal(testUrl, res5.Resolvers[0].Url)

	// can query resolvers by url
	res6, err := s.queryClient.ResolversByURL(s.ctx, &data.QueryResolversByURLRequest{
		Url: testUrl,
	})
	require.NoError(err)
	require.NotNil(res6)
	require.Equal(s.addr1.String(), res6.Resolvers[0].Manager)
	require.Equal(testUrl, res6.Resolvers[0].Url)
}
