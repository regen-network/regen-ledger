package marketplace

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
	"testing"
)

func TestQuery_SellOrders(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom1 := "C01-20200101-20200201-001"
	batchDenom2 := "C02-20200101-20200201-001"
	
	assert.NilError(t, s.coreStore.BatchInfoTable().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom1,
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.coreStore.BatchInfoTable().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom2,
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	_, _, otherAddr1 := testdata.KeyTestPubAddr()
	_, _, otherAddr2 := testdata.KeyTestPubAddr()

	// create 2 orders for s.addr with different batch id
	insertSellOrder(t, s, s.addr, 1)
	insertSellOrder(t, s, s.addr, 2)

	// create 2 orders for otherAddr1
	insertSellOrder(t, s, otherAddr1, 1)
	insertSellOrder(t, s, otherAddr1, 2)

	// create 2 orders for otherAddr2
	insertSellOrder(t, s, otherAddr2, 1)
	insertSellOrder(t, s, otherAddr2, 2)

	res, err := s.k.SellOrders(s.ctx, &marketplace.QuerySellOrdersRequest{
		BatchDenom: batchDenom1,
		Address:    "",
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 3, len(res.SellOrders))
	
	res, err = s.k.SellOrders(s.ctx, &marketplace.QuerySellOrdersRequest{
		BatchDenom: "",
		Address:    s.addr.String(),
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.SellOrders))
}

func insertSellOrder(t *testing.T, s *baseSuite, addr sdk.AccAddress, batchId uint64) {
	sellOrder :=  &marketplacev1.SellOrder{
		Seller:            addr,
		BatchId:           batchId,
		Quantity:          "10",
		MarketId:          1,
		AskPrice:          "10",
		DisableAutoRetire: false,
		Expiration:        timestamppb.Now(),
		Maker:             false,
	}
	assert.NilError(t, s.marketStore.SellOrderTable().Insert(s.ctx, sellOrder))
}
