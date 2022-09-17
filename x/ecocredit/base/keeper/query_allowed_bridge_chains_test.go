package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestQueryAllowedBridgeChains(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	chains := []string{"ethereum", "polygon", "solana"}
	for _, chain := range chains {
		err := s.stateStore.AllowedBridgeChainTable().Insert(s.ctx, &api.AllowedBridgeChain{ChainName: chain})
		require.NoError(t, err)
	}

	res, err := s.k.AllowedBridgeChains(s.ctx, &types.QueryAllowedBridgeChainsRequest{})
	require.NoError(t, err)

	for i, chain := range res.AllowedBridgeChains {
		require.Equal(t, chains[i], chain)
	}
}
