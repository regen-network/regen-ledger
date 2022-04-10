package server

import (
	"testing"
	"time"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

type anchorSuite struct {
	*baseSuite
	alice sdk.AccAddress
	ch    *data.ContentHash
	err   error
	id    []byte
}

func TestAnchor(t *testing.T) {
	gocuke.NewRunner(t, &anchorSuite{}).Path("./features/anchor.feature").Run()
}

func (s *anchorSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
}

func (s *anchorSuite) AValidContentHash() {
	s.ch = &data.ContentHash{
		Raw: &data.ContentHash_Raw{
			Hash:            make([]byte, 32),
			DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		},
	}
}

func (s *anchorSuite) AliceHasAnchoredTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.alice.String(),
		Hash:   s.ch,
	})
}

func (s *anchorSuite) AliceAttemptsToAnchorTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.alice.String(),
		Hash:   s.ch,
	})
}

func (s *anchorSuite) NoErrorIsReturned() {
	require.NoError(s.t, s.err)
}

func (s *anchorSuite) TheDataIdEntryExists() {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	s.id = dataId.Id
}

func (s *anchorSuite) TheDataAnchorEntryExistsAndTheTimestampIsEqualTo(a string) {
	anchorTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, s.id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataAnchor)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}
