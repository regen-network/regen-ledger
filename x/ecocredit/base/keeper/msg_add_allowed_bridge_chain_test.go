package keeper

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

type addAllowedBridgeChainSuite struct {
	*baseSuite
	err error
}

func TestAddAllowedBridgeChain(t *testing.T) {
	gocuke.NewRunner(t, &addAllowedBridgeChainSuite{}).Path("./features/msg_add_allowed_bridge_chain.feature").Run()
}

func (s *addAllowedBridgeChainSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *addAllowedBridgeChainSuite) ExpectChainNameToExist(a string) {
	found, err := s.stateStore.AllowedBridgeChainTable().Has(s.ctx, a)
	require.NoError(s.t, err)
	require.True(s.t, found)
}

func (s *addAllowedBridgeChainSuite) AllowedChainName(a string) {
	_, err := s.k.AddAllowedBridgeChain(s.ctx, &types.MsgAddAllowedBridgeChain{
		Authority: s.authority.String(),
		ChainName: a,
	})
	require.NoError(s.t, err)
}

func (s *addAllowedBridgeChainSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *addAllowedBridgeChainSuite) TheAuthorityAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.authority = addr
}

func (s *addAllowedBridgeChainSuite) AliceAttemptsToAddAllowedBridgeChainWithProperties(a gocuke.DocString) {
	var msg *types.MsgAddAllowedBridgeChain
	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.AddAllowedBridgeChain(s.ctx, msg)
}

func (s *addAllowedBridgeChainSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *addAllowedBridgeChainSuite) ExpectTheErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *addAllowedBridgeChainSuite) AliceAttemptsToAddChainName(a string) {
	_, s.err = s.k.AddAllowedBridgeChain(s.ctx, &types.MsgAddAllowedBridgeChain{
		Authority: s.authority.String(),
		ChainName: a,
	})
}
