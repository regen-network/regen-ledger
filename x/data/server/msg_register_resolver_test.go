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
	manager    sdk.AccAddress
	resolverId uint64
	err        error
}

func TestRegisterResolver(t *testing.T) {
	gocuke.NewRunner(t, &registerResolverSuite{}).Path("./features/register_resolver.feature").Run()
}

func (s *registerResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.manager = s.addrs[0]

	id, err := s.server.stateStore.ResolverInfoTable().InsertReturningID(s.ctx, &api.ResolverInfo{
		Url:     "https://foo.bar",
		Manager: s.manager,
	})
	require.NoError(s.t, err)

	s.resolverId = id
}

func (s *registerResolverSuite) AUserAttemptsToRegisterDataToAResolver() {
	_, s.err = s.server.RegisterResolver(s.ctx, &data.MsgRegisterResolver{
		Manager:    s.manager.String(),
		ResolverId: s.resolverId,
		Data:       []*data.ContentHash{},
	})
}

func (s *registerResolverSuite) TheDataIsRegisteredToTheResolver() {
	require.NoError(s.t, s.err)
}
