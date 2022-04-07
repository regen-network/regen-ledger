package server

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

type anchorSuite struct {
	*baseSuite
	ch     *data.ContentHash
	sender sdk.AccAddress
	err    error
}

func TestAnchor(t *testing.T) {
	gocuke.NewRunner(t, &anchorSuite{}).Path("./features/anchor.feature").Run()
}

func (s *anchorSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.sender = s.addrs[0]
}

func (s *anchorSuite) ARawDataContentHash() {
	s.ch = &data.ContentHash{
		Raw: &data.ContentHash_Raw{
			Hash:            make([]byte, 32),
			DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		},
	}
}

func (s *anchorSuite) AGraphDataContentHash() {
	s.ch = &data.ContentHash{
		Graph: &data.ContentHash_Graph{
			Hash:                      make([]byte, 32),
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}
}

func (s *anchorSuite) AUserAttemptsToAnchorTheData() {
	res, err := s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.sender.String(),
		Hash:   s.ch,
	})
	require.NoError(s.t, err)
	require.NotNil(s.t, res)
}

func (s *anchorSuite) TheDataIsAnchored() {
	require.NoError(s.t, s.err)
}

func (s *anchorSuite) TheDataIsNotAnchored() {
	require.Error(s.t, s.err)
}

func (s *anchorSuite) TheAnchoredDataIsEqualToTheDataSubmitted() {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataId.Id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataAnchor)

	require.Equal(s.t, dataId.Id, dataAnchor.Id)
	require.Equal(s.t, s.sdkCtx.BlockTime(), dataAnchor.Timestamp.AsTime())
}
