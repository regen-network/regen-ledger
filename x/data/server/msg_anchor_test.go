package server

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/data/v3"
)

type anchorSuite struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	ch    *data.ContentHash
	msg   *data.MsgAnchor
	err   error
}

func TestAnchor(t *testing.T) {
	gocuke.NewRunner(t, &anchorSuite{}).Path("./features/msg_anchor.feature").Run()
}

func (s *anchorSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
}

func (s *anchorSuite) TheContentHash(a gocuke.DocString) {
	s.ch = &data.ContentHash{}
	err := jsonpb.UnmarshalString(a.Content, s.ch)
	require.NoError(s.t, err)
}

func (s *anchorSuite) TheMessage(a gocuke.DocString) {
	s.msg = &data.MsgAnchor{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *anchorSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *anchorSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *anchorSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *anchorSuite) AliceHasAnchoredTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = s.sdkCtx.WithBlockTime(blockTime)

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.alice.String(),
		ContentHash: s.ch,
	})
}

func (s *anchorSuite) AliceAttemptsToAnchorTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = s.sdkCtx.WithBlockTime(blockTime)

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.alice.String(),
		ContentHash: s.ch,
	})
}

func (s *anchorSuite) BobAttemptsToAnchorTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = s.sdkCtx.WithBlockTime(blockTime)

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.bob.String(),
		ContentHash: s.ch,
	})
}

func (s *anchorSuite) TheAnchorEntryExistsWithTimestamp(a string) {
	anchorTime, err := types.ParseDate("anchor timestamp", a)
	require.NoError(s.t, err)

	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataID, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataID)

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataID.Id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataAnchor)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}

func (s *anchorSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event data.EventAnchor
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}
