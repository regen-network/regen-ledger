package keeper

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type removeAllowedBridgeChainSuite struct {
	*baseSuite
	err error
}

func TestRemoveAllowedBridgeChain(t *testing.T) {
	gocuke.NewRunner(t, &removeAllowedBridgeChainSuite{}).Path("./features/msg_remove_allowed_bridge_chain.feature").Run()
}

func (s *removeAllowedBridgeChainSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *removeAllowedBridgeChainSuite) ExpectChainNameToNotExist(a string) {
	chainName := a
	found, err := s.stateStore.AllowedBridgeChainTable().Has(s.ctx, chainName)
	require.NoError(s.t, err)
	require.False(s.t, found)
}

func (s *removeAllowedBridgeChainSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *removeAllowedBridgeChainSuite) TheAuthorityAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.authority = addr
}

func (s *removeAllowedBridgeChainSuite) AliceAttemptsToRemoveChainName(a string) {
	_, s.err = s.k.RemoveAllowedBridgeChain(s.ctx, &types.MsgRemoveAllowedBridgeChain{
		Authority: s.authority.String(),
		ChainName: a,
	})
}

func (s *removeAllowedBridgeChainSuite) AliceAttemptsToRemoveAllowedBridgeChainWithProperties(a gocuke.DocString) {
	var msg *types.MsgRemoveAllowedBridgeChain
	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.RemoveAllowedBridgeChain(s.ctx, msg)
}

func (s *removeAllowedBridgeChainSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *removeAllowedBridgeChainSuite) ExpectTheErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *removeAllowedBridgeChainSuite) TheChainName(a string) {
	_, err := s.k.AddAllowedBridgeChain(s.ctx, &types.MsgAddAllowedBridgeChain{
		Authority: s.authority.String(),
		ChainName: a,
	})
	require.NoError(s.t, err)
}
