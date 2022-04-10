package server

import (
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
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
	gocuke.NewRunner(t, &registerResolverSuite{}).Path("./features/register_resolver.feature").Run()
}

func (s *registerResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
}

func (s *registerResolverSuite) AValidContentHash() {
	s.ch = &data.ContentHash{
		Graph: &data.ContentHash_Graph{
			Hash:                      make([]byte, 32),
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}
}

func (s *registerResolverSuite) AliceHasAnchoredTheData() {
	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.alice.String(),
		Hash:   s.ch,
	})
}

func (s *registerResolverSuite) AliceHasDefinedAResolver() {
	id, err := s.server.stateStore.ResolverInfoTable().InsertReturningID(s.ctx, &api.ResolverInfo{
		Url:     "https://foo.bar",
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

func (s *registerResolverSuite) NoErrorIsReturned() {
	require.NoError(s.t, s.err)
}

func (s *registerResolverSuite) AnErrorIsReturned() {
	require.Error(s.t, s.err)
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

func (s *registerResolverSuite) AResolverWithIdDoesNotExist(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	found, err := s.server.stateStore.ResolverInfoTable().Has(s.ctx, id)
	require.NoError(s.t, err)
	require.False(s.t, found)
}
