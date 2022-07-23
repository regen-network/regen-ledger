package server

import (
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/regen-network/regen-ledger/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
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

	dataId := s.getDataId()

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataId)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataAnchor)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}

func (s *registerResolverSuite) TheDataResolverEntryExists() {
	dataId := s.getDataId()

	dataResolver, err := s.server.stateStore.DataResolverTable().Get(s.ctx, dataId, s.id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataResolver)
}

func (s *registerResolverSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *registerResolverSuite) getDataId() []byte {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	return dataId.Id
}
