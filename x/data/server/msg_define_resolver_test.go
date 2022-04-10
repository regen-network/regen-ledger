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
	alice sdk.AccAddress
	url   string
	err   error
}

func TestDefineResolver(t *testing.T) {
	gocuke.NewRunner(t, &defineResolverSuite{}).Path("./features/define_resolver.feature").Run()
}

func (s *defineResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
}

func (s *defineResolverSuite) AValidResolverUrl() {
	s.url = "https://foo.bar"
}

func (s *defineResolverSuite) AResolverEntryWithTheSameUrlAlreadyExists() {
	err := s.server.stateStore.ResolverInfoTable().Insert(s.ctx, &api.ResolverInfo{
		Url: s.url,
	})
	require.NoError(s.t, err)
}

func (s *defineResolverSuite) AliceHasDefinedTheResolver() {
	_, err := s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.alice.String(),
		ResolverUrl: s.url,
	})
	require.NoError(s.t, err)
}

func (s *defineResolverSuite) AliceAttemptsToDefineTheResolver() {
	_, s.err = s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.alice.String(),
		ResolverUrl: s.url,
	})
}

func (s *defineResolverSuite) NoErrorIsReturned() {
	require.NoError(s.t, s.err)
}

func (s *defineResolverSuite) AnErrorIsReturned() {
	require.Error(s.t, s.err)
}

func (s *defineResolverSuite) TheResolverInfoEntryExistsAndAliceIsTheManager() {
	dataResolver, err := s.server.stateStore.ResolverInfoTable().Get(s.ctx, 1)
	require.NoError(s.t, err)
	require.Equal(s.t, s.alice.Bytes(), dataResolver.Manager)
}
