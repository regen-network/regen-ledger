package marketplace

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestQuery_SellOrders(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom1 := "C01-20200101-20200201-001"
	batchDenom2 := "C02-20200101-20200201-001"

	id1, err := s.coreStore.BatchInfoTable().InsertReturningID(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom1,
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	})
	assert.NilError(t, err)
	id2, err := s.coreStore.BatchInfoTable().InsertReturningID(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom2,
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	})
	assert.NilError(t, err)

	_, _, otherAddr1 := testdata.KeyTestPubAddr()
	_, _, otherAddr2 := testdata.KeyTestPubAddr()

	// create 2 orders for s.addr with different batch id
	insertSellOrder(t, s, s.addr, id1)
	insertSellOrder(t, s, s.addr, id2)

	// create 2 orders for otherAddr1
	insertSellOrder(t, s, otherAddr1, id1)
	insertSellOrder(t, s, otherAddr1, id2)

	// create 2 orders for otherAddr2
	insertSellOrder(t, s, otherAddr2, id1)
	insertSellOrder(t, s, otherAddr2, id2)

	// can query with just batch denom
	queryAndAssertResponse(t, s, "", batchDenom1, 3)

	// can query with just addr
	queryAndAssertResponse(t, s, s.addr.String(), "", 2)

	// can query with both
	queryAndAssertResponse(t, s, s.addr.String(), batchDenom1, 1)

	// can query with none
	queryAndAssertResponse(t, s, "", "", 6)

	// check pagination
	res, err := s.k.SellOrders(s.ctx, &marketplace.QuerySellOrdersRequest{
		BatchDenom: "",
		Address:    "",
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     0,
			Limit:      3,
			CountTotal: false,
			Reverse:    false,
		},
	})
	assert.NilError(t, err)
	assert.Equal(t, 3, len(res.SellOrders))

	// check error when bad denom
	res, err = s.k.SellOrders(s.ctx, &marketplace.QuerySellOrdersRequest{
		BatchDenom: "foobar",
		Address:    "",
		Pagination: nil,
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func queryAndAssertResponse(t *testing.T, s *baseSuite, addr, batchDenom string, length int) {
	res, err := s.k.SellOrders(s.ctx, &marketplace.QuerySellOrdersRequest{
		BatchDenom: batchDenom,
		Address:    addr,
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, length, len(res.SellOrders))
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
