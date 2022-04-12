package server

import (
	"encoding/json"
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
}

func TestAnchor(t *testing.T) {
	runner := gocuke.NewRunner(t, &anchorSuite{}).Path("./features/anchor.feature")
	runner.Step(`the content hash "((?:[^\"]|\")*)"`, (*anchorSuite).TheContentHash)
	runner.Run()
}

func (s *anchorSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *anchorSuite) AliceIsTheSender() {
	s.alice = s.addrs[0]
}

func (s *anchorSuite) TheContentHash(a gocuke.DocString) {
	err := json.Unmarshal([]byte(a.Content), &s.ch)
	require.NoError(s.t, err)
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

func (s *anchorSuite) TheDataAnchorEntryExistsAndTheTimestampIsEqualTo(a string) {
	anchorTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataId.Id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataAnchor)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}
