package marketplace

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestQueryBuyOrder(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	assert.NilError(t, s.marketStore.BuyOrderTable().Insert(s.ctx, &api.BuyOrder{Buyer: s.addr}))

	res, err := s.k.BuyOrder(s.ctx, &marketplace.QueryBuyOrderRequest{BuyOrderId: 1})
	assert.NilError(t, err)
	assert.Check(t, s.addr.Equals(sdk.AccAddress(res.BuyOrder.Buyer)))

	_, err = s.k.BuyOrder(s.ctx, &marketplace.QueryBuyOrderRequest{BuyOrderId: 35})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
