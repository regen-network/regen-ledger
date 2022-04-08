package server

import (
	"testing"

	"github.com/regen-network/gocuke"
	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

type registerResolverSuite struct {
	*baseSuite
	ch          *data.ContentHash
	manager     sdk.AccAddress
	user        sdk.AccAddress
	resolverUrl string
	resolverId  uint64
	err         error
}

func TestRegisterResolver(t *testing.T) {
	gocuke.NewRunner(t, &registerResolverSuite{}).Path("./features/register_resolver.feature").Run()
}

func (s *registerResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.manager = s.addrs[0]
}

func (s *registerResolverSuite) AValidContentHashThatHasBeenAnchored() {
	s.ch = &data.ContentHash{
		Graph: &data.ContentHash_Graph{
			Hash:                      make([]byte, 32),
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.manager.String(),
		Hash:   s.ch,
	})
}

func (s *registerResolverSuite) AValidContentHashThatHasNotBeenAnchored() {
	s.ch = &data.ContentHash{
		Graph: &data.ContentHash_Graph{
			Hash:                      make([]byte, 32),
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}
}

func (s *registerResolverSuite) AnInvalidContentHash() {
	s.ch = &data.ContentHash{
		Graph: &data.ContentHash_Graph{
			Hash:                      make([]byte, 16),
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}
}

func (s *registerResolverSuite) AResolverHasBeenDefinedByTheManager() {
	s.resolverUrl = "https://foo.bar"

	id, err := s.server.stateStore.ResolverInfoTable().InsertReturningID(s.ctx, &api.ResolverInfo{
		Url:     s.resolverUrl,
		Manager: s.manager,
	})
	require.NoError(s.t, err)

	s.resolverId = id
}

func (s *registerResolverSuite) TheManagerAttemptsToRegisterDataToTheResolver() {
	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:    s.manager.String(),
		ResolverId: s.resolverId,
		Data:       []*data.ContentHash{},
	})
}

func (s *registerResolverSuite) AnotherUserAttemptsToRegisterDataToTheResolver() {
	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:    s.user.String(),
		ResolverId: s.resolverId,
		Data:       []*data.ContentHash{},
	})
}

func (s *registerResolverSuite) TheDataIsRegisteredToTheResolver() {
	require.NoError(s.t, s.err)
}

func (s *registerResolverSuite) TheDataIsNotRegisteredToTheResolver() {
	require.Error(s.t, s.err)
}

func (s *registerResolverSuite) ADataResolverEntryIsCreated() {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	dataResolver, err := s.server.stateStore.DataResolverTable().Get(s.ctx, dataId.Id, s.resolverId)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataResolver)
}
