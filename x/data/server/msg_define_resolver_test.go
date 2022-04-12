package server

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

type defineResolverSuite struct {
	*baseSuite
	alice sdk.AccAddress
	err   error
}

func TestDefineResolver(t *testing.T) {
	gocuke.NewRunner(t, &defineResolverSuite{}).Path("./features/define_resolver.feature").Run()
}

func (s *defineResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *defineResolverSuite) AliceIsTheManager() {
	s.alice = s.addrs[0]
}

func (s *defineResolverSuite) AliceHasDefinedAResolverWithUrl(a string) {
	_, err := s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.alice.String(),
		ResolverUrl: a,
	})
	require.NoError(s.t, err)
}

func (s *defineResolverSuite) AliceAttemptsToDefineAResolverWithUrl(a string) {
	_, s.err = s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.alice.String(),
		ResolverUrl: a,
	})
}

func (s *defineResolverSuite) TheResolverInfoEntryExistsWithUrl(a string) {
	dataResolver, err := s.server.stateStore.ResolverInfoTable().Get(s.ctx, 1)
	require.NoError(s.t, err)
	require.Equal(s.t, a, dataResolver.Url)
	require.Equal(s.t, s.alice.Bytes(), dataResolver.Manager)
}

func (s *defineResolverSuite) ExpectTheError(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
