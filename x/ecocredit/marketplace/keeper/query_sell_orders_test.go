package keeper

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

var (
	classID       = "C01"
	batchDenom    = "C01-001-20200101-20200201-001"
	start, end    = timestamppb.Now(), timestamppb.Now()
	validAskDenom = sdk.DefaultBondDenom
	ask           = sdk.NewInt64Coin(validAskDenom, 10)
	creditType    = basetypes.CreditType{Name: "carbon", Abbreviation: "C", Unit: "tonnes", Precision: 6}
)

func TestSellOrders(t *testing.T) {
	t.Parallel()
	s := setupBase(t, 2)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classID, start, end, creditType)

	// nil request
	_, err := s.k.SellOrders(s.ctx, nil)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid argument")

	order1 := insertSellOrder(t, s, s.addrs[0], 1)
	order2 := insertSellOrder(t, s, s.addrs[1], 1)

	res, err := s.k.SellOrders(s.ctx, &types.QuerySellOrdersRequest{
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	if res.SellOrders[0].Id == order1.Id {
		assertOrderEqual(s.ctx, t, s.k, res.SellOrders[0], order1)
	} else {
		assertOrderEqual(s.ctx, t, s.k, res.SellOrders[0], order2)
	}
	assert.Equal(t, uint64(2), res.Pagination.Total)
}

func insertSellOrder(t *testing.T, s *baseSuite, addr sdk.AccAddress, batchKey uint64) *api.SellOrder {
	sellOrder := &api.SellOrder{
		Seller:            addr,
		BatchKey:          batchKey,
		Quantity:          "10",
		MarketId:          1,
		AskAmount:         "10",
		DisableAutoRetire: false,
		Expiration:        timestamppb.Now(),
		Maker:             false,
	}
	assert.NilError(t, s.marketStore.SellOrderTable().Insert(s.ctx, sellOrder))

	return sellOrder
}

func assertOrderEqual(ctx context.Context, t *testing.T, k Keeper, received *types.SellOrderInfo, order *api.SellOrder) {
	seller := sdk.AccAddress(order.Seller)

	batch, err := k.baseStore.BatchTable().Get(ctx, order.BatchKey)
	assert.NilError(t, err)

	market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
	assert.NilError(t, err)

	info := types.SellOrderInfo{
		Id:                order.Id,
		Seller:            seller.String(),
		BatchDenom:        batch.Denom,
		Quantity:          order.Quantity,
		AskDenom:          market.BankDenom,
		AskAmount:         order.AskAmount,
		DisableAutoRetire: order.DisableAutoRetire,
		Expiration:        regentypes.ProtobufToGogoTimestamp(order.Expiration),
	}

	assert.DeepEqual(t, info, *received)
}
