package server

import (
	"testing"
	"time"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

type attestSuite struct {
	*baseSuite
	alice sdk.AccAddress
	ch    *data.ContentHash
	err   error
	id    []byte
}

func TestAttest(t *testing.T) {
	gocuke.NewRunner(t, &attestSuite{}).Path("./features/attest.feature").Run()
}

func (s *attestSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
}

func (s *attestSuite) AValidContentHash() {
	s.ch = &data.ContentHash{
		Graph: &data.ContentHash_Graph{
			Hash:                      make([]byte, 32),
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}
}

func (s *attestSuite) AliceHasAnchoredTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.alice.String(),
		Hash:   s.ch,
	})
}

func (s *attestSuite) AliceHasAttestedToTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.alice.String(),
		Hashes:   []*data.ContentHash_Graph{s.ch.GetGraph()},
	})
}

func (s *attestSuite) AliceAttemptsToAttestToTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.alice.String(),
		Hashes:   []*data.ContentHash_Graph{s.ch.GetGraph()},
	})
}

func (s *attestSuite) AnErrorOf(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}

func (s *attestSuite) TheDataIdEntryExists() {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	s.id = dataId.Id
}

func (s *attestSuite) TheDataAnchorEntryExistsAndTheTimestampIsEqualTo(a string) {
	anchorTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, s.id)
	require.NoError(s.t, err)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}

func (s *attestSuite) TheDataAttestorEntryExistsAndTheTimestampIsEqualTo(a string) {
	attestTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	dataAttestor, err := s.server.stateStore.DataAttestorTable().Get(s.ctx, s.id, s.alice)
	require.NoError(s.t, err)
	require.Equal(s.t, attestTime, dataAttestor.Timestamp.AsTime())
}
