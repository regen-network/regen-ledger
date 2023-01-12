package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"

	v1 "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

func TestQueryInterchainAccount(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	owner := s.addrs[0]

	portID, err := types.NewControllerPortID(owner.String())
	assert.NilError(t, err)

	connectionID := "ch-5"

	icaAddr := "foo"
	s.ica.
		EXPECT().
		GetInterchainAccountAddress(s.ctx, connectionID, portID).
		Times(1).
		Return(icaAddr, true)

	res, err := s.k.InterchainAccount(s.ctx, &v1.QueryInterchainAccountRequest{
		Owner:        owner.String(),
		ConnectionId: connectionID,
	})
	assert.NilError(t, err)

	assert.Equal(t, icaAddr, res.InterchainAccountAddress)
}
