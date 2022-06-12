package testsuite

import (
	"context"
	"crypto"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
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

	// set block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// convert block time to expected format for anchor response
	anchorBlockTime, err := gogotypes.TimestampProto(s.sdkCtx.BlockTime())
	require.NoError(err)

	// anchor some data
	anchorRes1, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.addr1.String(),
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.NotNil(anchorRes1)
	require.Equal(iri, anchorRes1.Iri)
	require.Equal(anchorBlockTime, anchorRes1.Timestamp)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
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
	require.NotNil(iri, anchorByIRI.Anchor.Iri)
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

	// can convert iri to hash
	iriToHash, err := s.queryClient.ConvertIRIToHash(s.ctx, &data.ConvertIRIToHashRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.NotNil(iriToHash)
	require.Equal(s.hash1, iriToHash.ContentHash)

	// can query attestor entries by iri (no attestors)
	attestationsByIri, err := s.queryClient.AttestationsByIRI(s.ctx, &data.QueryAttestationsByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.Empty(attestationsByIri.Attestations)

	// can query attestor entries by hash (no attestors)
	attestationsByHash, err := s.queryClient.AttestationsByHash(s.ctx, &data.QueryAttestationsByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.Empty(attestationsByHash.Attestations)

	// can attest to data
	attestRes1, err := s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr1.String(),
		ContentHashes: []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)

	// attesting to the same data twice is a no-op
	attestRes2, err := s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr1.String(),
		ContentHashes: []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)
	require.Nil(attestRes2.Iris)
	require.Equal(attestRes1.Timestamp, attestRes2.Timestamp)

	// can query attestor entries by iri (one attestor)
	attestationsByIri, err = s.queryClient.AttestationsByIRI(s.ctx, &data.QueryAttestationsByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.Len(attestationsByIri.Attestations, 1)
	require.Equal(s.addr1.String(), attestationsByIri.Attestations[0].Attestor)
	require.Equal(attestRes1.Timestamp, attestationsByIri.Attestations[0].Timestamp)

	// can query attestor entries by hash (one attestor)
	attestationsByHash, err = s.queryClient.AttestationsByHash(s.ctx, &data.QueryAttestationsByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.Len(attestationsByHash.Attestations, 1)
	require.Equal(s.addr1.String(), attestationsByHash.Attestations[0].Attestor)
	require.Equal(attestRes1.Timestamp, attestationsByHash.Attestations[0].Timestamp)

	// can query anchored data entries by attestor
	attestationsByAttestor, err := s.queryClient.AttestationsByAttestor(s.ctx, &data.QueryAttestationsByAttestorRequest{
		Attestor: s.addr1.String(),
	})
	require.NoError(err)
	require.NotNil(attestationsByAttestor)
	require.Len(attestationsByAttestor.Attestations, 1)
	require.Equal(iri, attestationsByAttestor.Attestations[0].Iri)
	require.Equal(anchorRes1.Timestamp, attestationsByAttestor.Attestations[0].Timestamp)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// another attestor can attest to the same anchored data entry
	attestRes3, err := s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr2.String(),
		ContentHashes: []*data.ContentHash_Graph{graphHash},
	})
	require.NoError(err)
	require.Len(attestRes3.Iris, 1)
	require.Equal(iri, attestRes3.Iris[0])
	require.NotEqual(attestRes2.Timestamp, attestRes3.Timestamp)

	// can query attestor entries by IRI (two attestors)
	attestationsByIri, err = s.queryClient.AttestationsByIRI(s.ctx, &data.QueryAttestationsByIRIRequest{
		Iri: iri,
	})
	require.NoError(err)
	require.Len(attestationsByIri.Attestations, 2)

	// order may vary from query response
	if s.addr1.String() == attestationsByIri.Attestations[0].Attestor {
		require.Equal(attestRes1.Timestamp, attestationsByIri.Attestations[0].Timestamp)
		require.Equal(s.addr2.String(), attestationsByIri.Attestations[1].Attestor)
		require.Equal(attestRes3.Timestamp, attestationsByIri.Attestations[1].Timestamp)
	} else {
		require.Equal(s.addr2.String(), attestationsByIri.Attestations[0].Attestor)
		require.Equal(attestRes3.Timestamp, attestationsByIri.Attestations[0].Timestamp)
		require.Equal(s.addr1.String(), attestationsByIri.Attestations[1].Attestor)
		require.Equal(attestRes1.Timestamp, attestationsByIri.Attestations[1].Timestamp)
	}

	// can query attestor entries by hash (two attestors)
	attestationsByHash, err = s.queryClient.AttestationsByHash(s.ctx, &data.QueryAttestationsByHashRequest{
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.Len(attestationsByHash.Attestations, 2)

	// order may vary from query response
	if s.addr1.String() == attestationsByHash.Attestations[0].Attestor {
		require.Equal(attestRes1.Timestamp, attestationsByHash.Attestations[0].Timestamp)
		require.Equal(s.addr2.String(), attestationsByHash.Attestations[1].Attestor)
		require.Equal(attestRes3.Timestamp, attestationsByHash.Attestations[1].Timestamp)
	} else {
		require.Equal(s.addr2.String(), attestationsByHash.Attestations[0].Attestor)
		require.Equal(attestRes3.Timestamp, attestationsByHash.Attestations[0].Timestamp)
		require.Equal(s.addr1.String(), attestationsByHash.Attestations[1].Attestor)
		require.Equal(attestRes1.Timestamp, attestationsByHash.Attestations[1].Timestamp)
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
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
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
