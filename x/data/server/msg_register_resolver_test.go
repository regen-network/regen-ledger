package server

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
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
	runner := gocuke.NewRunner(t, &registerResolverSuite{}).Path("./features/register_resolver.feature")
	runner.Step(`a content hash of "((?:[^\"]|\")*)"`, (*registerResolverSuite).AContentHashOf)
	runner.Run()
}

func (s *registerResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *registerResolverSuite) AliceIsTheManager() {
	s.alice = s.addrs[0]
}

func (s *registerResolverSuite) BobIsNotTheManager() {
	s.bob = s.addrs[1]
}

func (s *registerResolverSuite) AContentHashOf(a gocuke.DocString) {
	err := json.Unmarshal([]byte(a.Content), &s.ch)
	require.NoError(s.t, err)
}

func (s *registerResolverSuite) AliceHasAnchoredTheData() {
	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.alice.String(),
		Hash:   s.ch,
	})
}

func (s *registerResolverSuite) AliceHasDefinedAResolverWithUrl(a string) {
	id, err := s.server.stateStore.ResolverInfoTable().InsertReturningID(s.ctx, &api.ResolverInfo{
		Url:     a,
		Manager: s.alice,
	})
	require.NoError(s.t, err)

	s.id = id
}

func (s *registerResolverSuite) AliceAttemptsToRegisterTheDataToTheResolver() {
	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:    s.alice.String(),
		ResolverId: s.id,
		Data:       []*data.ContentHash{s.ch},
	})
}

func (s *registerResolverSuite) AliceAttemptsToRegisterTheDataToAResolverWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:    s.alice.String(),
		ResolverId: id,
		Data:       []*data.ContentHash{s.ch},
	})
}

func (s *registerResolverSuite) BobAttemptsToRegisterDataToTheResolver() {
	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:    s.bob.String(),
		ResolverId: s.id,
		Data:       []*data.ContentHash{s.ch},
	})
}

func (s *registerResolverSuite) AnErrorOf(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}

func (s *registerResolverSuite) TheDataAnchorEntryExists() {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataId.Id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataAnchor)
}

func (s *registerResolverSuite) TheDataResolverEntryExists() {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	dataResolver, err := s.server.stateStore.DataResolverTable().Get(s.ctx, dataId.Id, s.id)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataResolver)
}
