package server

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

type defineResolverSuite struct {
	*baseSuite
	manager     sdk.AccAddress
	resolverUrl string
	err         error
}

func TestDefineResolver(t *testing.T) {
	gocuke.NewRunner(t, &defineResolverSuite{}).Path("./features/define_resolver.feature").Run()
}

func (s *defineResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.manager = s.addrs[0]
}

func (s *defineResolverSuite) AValidResolverUrl() {
	s.resolverUrl = "https://foo.bar"
}

func (s *defineResolverSuite) AResolverEntryWithTheSameUrlAlreadyExists() {
	err := s.server.stateStore.ResolverInfoTable().Insert(s.ctx, &api.ResolverInfo{
		Url: s.resolverUrl,
	})
	require.NoError(s.t, err)
}

func (s *defineResolverSuite) AUserAttemptsToDefineAResolver() {
	_, s.err = s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.manager.String(),
		ResolverUrl: s.resolverUrl,
	})
}

func (s *defineResolverSuite) TheResolverIsDefined() {
	require.NoError(s.t, s.err)
}

func (s *defineResolverSuite) TheResolverIsNotDefined() {
	require.Error(s.t, s.err)
}

func (s *defineResolverSuite) AResolverInfoEntryIsCreatedAndTheManagerIsEqualToTheUserAddress() {
	dataResolver, err := s.server.stateStore.ResolverInfoTable().Get(s.ctx, 1)
	require.NoError(s.t, err)
	require.Equal(s.t, s.manager.Bytes(), dataResolver.Manager)
}
