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
	id     []byte
}

func TestAnchor(t *testing.T) {
	gocuke.NewRunner(t, &anchorSuite{}).Path("./features/anchor.feature").Run()
}

func (s *anchorSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.sender = s.addrs[0]
}

func (s *anchorSuite) AValidContentHash() {
	s.ch = &data.ContentHash{
		Raw: &data.ContentHash_Raw{
			Hash:            make([]byte, 32),
			DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		},
	}
}

func (s *anchorSuite) AnInvalidContentHash() {
	s.ch = &data.ContentHash{
		Raw: &data.ContentHash_Raw{
			Hash:            make([]byte, 16),
			DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		},
	}
}

func (s *anchorSuite) AUserAttemptsToAnchorTheData() {
	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.sender.String(),
		Hash:   s.ch,
	})
}

func (s *anchorSuite) TheDataIsAnchored() {
	require.NoError(s.t, s.err)
}

func (s *anchorSuite) TheDataIsNotAnchored() {
	require.Error(s.t, s.err)
}

func (s *anchorSuite) ADataIdEntryIsCreated() {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	s.id = dataId.Id
}

func (s *anchorSuite) ADataAnchorEntryIsCreatedAndTheTimestampIsEqualToTheBlockTime() {
	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, s.id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataAnchor)
	require.Equal(s.t, s.sdkCtx.BlockTime(), dataAnchor.Timestamp.AsTime())
}
