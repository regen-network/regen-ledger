package server

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/x/data"
)

type attestSuite struct {
	*baseSuite
	ch       *data.ContentHash
	attestor sdk.AccAddress
	err      error
	id       []byte
}

func TestAttest(t *testing.T) {
	gocuke.NewRunner(t, &attestSuite{}).Path("./features/attest.feature").Run()
}

func (s *attestSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.attestor = s.addrs[0]
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

func (s *attestSuite) AnInvalidContentHash() {
	s.ch = &data.ContentHash{
		Graph: &data.ContentHash_Graph{
			Hash:                      make([]byte, 16),
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}
}

func (s *attestSuite) TheDataHasBeenAnchoredAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.attestor.String(),
		Hash:   s.ch,
	})
}

func (s *attestSuite) TheDataHasNotBeenAnchored() {
	// no-op
}

func (s *attestSuite) AUserAttemptsToAttestToTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.attestor.String(),
		Hashes:   []*data.ContentHash_Graph{s.ch.GetGraph()},
	})
}

func (s *attestSuite) TheDataIsAttestedTo() {
	require.NoError(s.t, s.err)
}

func (s *attestSuite) TheDataIsNotAttestedTo() {
	require.Error(s.t, s.err)
}

func (s *attestSuite) ADataIdEntryIsCreated() {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	s.id = dataId.Id
}

func (s *attestSuite) ADataAnchorEntryIsCreatedAndTheTimestampIsEqualTo(a string) {
	anchorTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, s.id)
	require.NoError(s.t, err)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}

func (s *attestSuite) ADataAttestorEntryIsCreatedAndTheTimestampIsEqualTo(a string) {
	attestTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	dataAttestor, err := s.server.stateStore.DataAttestorTable().Get(s.ctx, s.id, s.attestor)
	require.NoError(s.t, err)
	require.Equal(s.t, attestTime, dataAttestor.Timestamp.AsTime())
}
