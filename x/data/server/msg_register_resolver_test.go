//nolint:revive,stylecheck
package server

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

type registerResolverSuite struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	ch    *data.ContentHash
	id    uint64
	err   error
}

func TestRegisterResolver(t *testing.T) {
	gocuke.NewRunner(t, &registerResolverSuite{}).Path("./features/msg_register_resolver.feature").Run()
}

func (s *registerResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
}

func (s *registerResolverSuite) TheContentHash(a gocuke.DocString) {
	s.ch = &data.ContentHash{}
	err := jsonpb.UnmarshalString(a.Content, s.ch)
	require.NoError(s.t, err)
}

func (s *registerResolverSuite) AliceHasDefinedTheResolverWithUrl(a string) {
	res, err := s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.alice.String(),
		ResolverUrl: a,
	})
	require.NoError(s.t, err)

	s.id = res.ResolverId
}

func (s *registerResolverSuite) AliceHasAnchoredTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.alice.String(),
		ContentHash: s.ch,
	})
}

func (s *registerResolverSuite) AliceHasRegisteredTheDataToTheResolver() {
	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:       s.alice.String(),
		ResolverId:    s.id,
		ContentHashes: []*data.ContentHash{s.ch},
	})
}

func (s *registerResolverSuite) AliceAttemptsToRegisterTheDataToTheResolver() {
	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:       s.alice.String(),
		ResolverId:    s.id,
		ContentHashes: []*data.ContentHash{s.ch},
	})
}

func (s *registerResolverSuite) AliceAttemptsToRegisterTheDataToAResolverWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:       s.alice.String(),
		ResolverId:    id,
		ContentHashes: []*data.ContentHash{s.ch},
	})
}

func (s *registerResolverSuite) AliceAttemptsToRegisterTheDataToTheResolverAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:       s.alice.String(),
		ResolverId:    s.id,
		ContentHashes: []*data.ContentHash{s.ch},
	})
}

func (s *registerResolverSuite) BobAttemptsToRegisterTheDataToTheResolver() {
	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:       s.bob.String(),
		ResolverId:    s.id,
		ContentHashes: []*data.ContentHash{s.ch},
	})
}

func (s *registerResolverSuite) TheAnchorEntryExistsWithTimestamp(a string) {
	anchorTime, err := types.ParseDate("anchor timestamp", a)
	require.NoError(s.t, err)

	dataID := s.getDataID()

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataID)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataAnchor)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}

func (s *registerResolverSuite) TheDataResolverEntryExists() {
	dataID := s.getDataID()

	dataResolver, err := s.server.stateStore.DataResolverTable().Get(s.ctx, dataID, s.id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataResolver)
}

func (s *registerResolverSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *registerResolverSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event data.EventRegisterResolver
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *registerResolverSuite) getDataID() []byte {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataID, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataID)

	return dataID.Id
}
