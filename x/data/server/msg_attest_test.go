package server

import (
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

type attestSuite struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	ch    *data.ContentHash
	err   error
}

func TestAttest(t *testing.T) {
	gocuke.NewRunner(t, &attestSuite{}).Path("./features/msg_attest.feature").Run()
}

func (s *attestSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
}

func (s *attestSuite) TheContentHash(a gocuke.DocString) {
	s.ch = &data.ContentHash{}
	err := jsonpb.UnmarshalString(a.Content, s.ch)
	require.NoError(s.t, err)
}

func (s *attestSuite) AlicesAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.alice = addr
}

func (s *attestSuite) AliceHasAnchoredTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.alice.String(),
		ContentHash: s.ch,
	})
}

func (s *attestSuite) AliceHasAttestedToTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.alice.String(),
		ContentHashes: []*data.ContentHash_Graph{s.ch.Graph},
	})
}

func (s *attestSuite) AliceAttemptsToAttestToTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.alice.String(),
		ContentHashes: []*data.ContentHash_Graph{s.ch.Graph},
	})
}

func (s *attestSuite) BobAttemptsToAttestToTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.bob.String(),
		ContentHashes: []*data.ContentHash_Graph{s.ch.Graph},
	})
}

func (s *attestSuite) TheAnchorEntryExistsWithTimestamp(a string) {
	anchorTime, err := types.ParseDate("anchor timestamp", a)
	require.NoError(s.t, err)

	dataID := s.getDataID()

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataID)
	require.NoError(s.t, err)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}

func (s *attestSuite) TheAttestorEntryForAliceExistsWithTimestamp(a string) {
	attestTime, err := types.ParseDate("attest timestamp", a)
	require.NoError(s.t, err)

	dataID := s.getDataID()

	dataAttestor, err := s.server.stateStore.DataAttestorTable().Get(s.ctx, dataID, s.alice)
	require.NoError(s.t, err)
	require.Equal(s.t, attestTime, dataAttestor.Timestamp.AsTime())
}

func (s *attestSuite) TheAttestorEntryForBobExistsWithTimestamp(a string) {
	attestTime, err := types.ParseDate("attest timestamp", a)
	require.NoError(s.t, err)

	dataID := s.getDataID()

	dataAttestor, err := s.server.stateStore.DataAttestorTable().Get(s.ctx, dataID, s.bob)
	require.NoError(s.t, err)
	require.Equal(s.t, attestTime, dataAttestor.Timestamp.AsTime())
}

func (s *attestSuite) EventIsEmittedWithProperties(a gocuke.DocString) {
	var event data.EventAttest
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *attestSuite) getDataID() []byte {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataID, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataID)

	return dataID.Id
}
