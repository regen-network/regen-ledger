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
	s.resolverUrl = "https://foo.bar"
}

func (s *defineResolverSuite) AUserAttemptsToDefineAResolver() {
	_, s.err = s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.manager.String(),
		ResolverUrl: s.resolverUrl,
	})
}

func (s *defineResolverSuite) TheResolverIsCreated() {
	require.NoError(s.t, s.err)
}
