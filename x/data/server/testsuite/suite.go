package testsuite

import (
	"bytes"
	"context"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/testutil/fixture"
	"github.com/regen-network/regen-ledger/x/data"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory fixture.Factory
	fixture        fixture.Fixture

	ctx         context.Context
	sdkCtx      sdk.Context
	msgClient   data.MsgClient
	queryClient data.QueryClient

	addr1 sdk.AccAddress
	addr2 sdk.AccAddress
	hash1 *data.ContentHash
	hash2 *data.ContentHash

	graphHash *data.ContentHash_Graph // hash1
	rawHash   *data.ContentHash_Raw   // hash2
}

func NewIntegrationTestSuite(fixtureFactory fixture.Factory) *IntegrationTestSuite {
	return &IntegrationTestSuite{fixtureFactory: fixtureFactory}
}

func (s *IntegrationTestSuite) SetupSuite() {
	require := s.Require()

	s.fixture = s.fixtureFactory.Setup()
	s.ctx = s.fixture.Context()
	s.sdkCtx = sdk.UnwrapSDKContext(s.ctx)
	s.msgClient = data.NewMsgClient(s.fixture.TxConn())
	s.queryClient = data.NewQueryClient(s.fixture.QueryConn())
	require.GreaterOrEqual(len(s.fixture.Signers()), 2)
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]

	s.graphHash = &data.ContentHash_Graph{
		Hash:                      bytes.Repeat([]byte{0}, 32),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}
	s.hash1 = &data.ContentHash{Graph: s.graphHash}

	s.rawHash = &data.ContentHash_Raw{
		Hash:            bytes.Repeat([]byte{0}, 32),
		DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		MediaType:       data.RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
	}
	s.hash2 = &data.ContentHash{Raw: s.rawHash}
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

func (s *IntegrationTestSuite) TestGraphScenario() {
	require := s.Require()

	iri, err := s.graphHash.ToIRI()
	require.NoError(err)

	// set block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// convert block time to expected format for anchor response
	startingBlockTime, err := gogotypes.TimestampProto(s.sdkCtx.BlockTime())
	require.NoError(err)

	// can anchor data
	anchorRes1, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.addr1.String(),
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.Equal(startingBlockTime, anchorRes1.Timestamp)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// anchoring same data twice is a no-op
	anchorRes2, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.addr1.String(),
		ContentHash: s.hash1,
	})
	require.NoError(err)
	require.Equal(anchorRes1.Timestamp, anchorRes2.Timestamp)

	// can attest to data
	attestRes1, err := s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr1.String(),
		ContentHashes: []*data.ContentHash_Graph{s.graphHash},
	})
	require.NoError(err)
	require.NotEqual(anchorRes1.Timestamp, attestRes1.Timestamp)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// attesting to the same data twice is a no-op
	attestRes2, err := s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr1.String(),
		ContentHashes: []*data.ContentHash_Graph{s.graphHash},
	})
	require.NoError(err)
	require.Len(attestRes2.Iris, 0)
	require.NotContains(attestRes2.Iris, iri)
	require.NotEqual(attestRes1.Timestamp, attestRes2.Timestamp)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// another attestor can attest to the same data
	attestRes3, err := s.msgClient.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.addr2.String(),
		ContentHashes: []*data.ContentHash_Graph{s.graphHash},
	})
	require.NoError(err)
	require.Len(attestRes3.Iris, 1)
	require.Contains(attestRes3.Iris, iri)
	require.NotEqual(attestRes2.Timestamp, attestRes3.Timestamp)
}

func (s *IntegrationTestSuite) TestRawDataScenario() {
	require := s.Require()

	iri, err := s.hash2.ToIRI()
	require.NoError(err)
	require.NotEmpty(iri)

	// can anchor data
	anchorRes1, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.addr1.String(),
		ContentHash: s.hash2,
	})
	require.NoError(err)

	// update block time
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now().UTC())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// anchoring same data twice is a no-op
	anchorRes2, err := s.msgClient.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.addr1.String(),
		ContentHash: s.hash2,
	})
	require.NoError(err)
	require.Equal(anchorRes1.Timestamp, anchorRes2.Timestamp)
}

func (s *IntegrationTestSuite) TestResolver() {
	require := s.Require()
	testURL := "https://foo.bar"
	hashes := []*data.ContentHash{s.hash1, s.hash2}

	// can define a resolver
	defineResolver, err := s.msgClient.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.addr1.String(),
		ResolverUrl: testURL,
	})
	require.NoError(err)

	// can register content to a resolver
	_, err = s.msgClient.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:       s.addr1.String(),
		ResolverId:    defineResolver.ResolverId,
		ContentHashes: hashes,
	})
	require.NoError(err)

	// registering same data twice is a no-op
	_, err = s.msgClient.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:       s.addr1.String(),
		ResolverId:    defineResolver.ResolverId,
		ContentHashes: hashes,
	})
	require.NoError(err)
}
