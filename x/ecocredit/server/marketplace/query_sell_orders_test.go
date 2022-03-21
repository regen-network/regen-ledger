package marketplace

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

var (
	classId    = "C01"
	batchDenom = "C01-20200101-20200201-001"
	start, end = timestamppb.Now(), timestamppb.Now()
	ask        = sdk.NewInt64Coin("ufoo", 10)
	creditType = ecocredit.CreditType{Name: "carbon", Abbreviation: "C", Unit: "tonnes", Precision: 6}
)

func TestSellOrders(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	_, _, addr2 := testdata.KeyTestPubAddr()

	insertSellOrder(t, s, s.addr, 1)
	insertSellOrder(t, s, addr2, 1)

	res, err := s.k.SellOrders(s.ctx, &marketplace.QuerySellOrdersRequest{
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.SellOrders))
}

func TestSellOrdersByDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	// make another batch
	otherDenom := "C01-19990101-20290101-001"
	assert.NilError(t, s.coreStore.BatchInfoTable().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: otherDenom,
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	insertSellOrder(t, s, s.addr, 1)
	insertSellOrder(t, s, s.addr, 2)

	// query the first denom
	res, err := s.k.SellOrdersByBatchDenom(s.ctx, &marketplace.QuerySellOrdersByBatchDenomRequest{
		BatchDenom: batchDenom,
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))

	// query the second denom
	res, err = s.k.SellOrdersByBatchDenom(s.ctx, &marketplace.QuerySellOrdersByBatchDenomRequest{
		BatchDenom: otherDenom,
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))

	// bad denom should error
	res, err = s.k.SellOrdersByBatchDenom(s.ctx, &marketplace.QuerySellOrdersByBatchDenomRequest{
		BatchDenom: "yikes!",
		Pagination: nil,
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func TestSellOrdersByAddress(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	_, _, otherAddr := testdata.KeyTestPubAddr()
	_, _, noOrdersAddr := testdata.KeyTestPubAddr()

	insertSellOrder(t, s, s.addr, 1)
	insertSellOrder(t, s, otherAddr, 1)

	res, err := s.k.SellOrdersByAddress(s.ctx, &marketplace.QuerySellOrdersByAddressRequest{
		Address:    s.addr.String(),
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))

	res, err = s.k.SellOrdersByAddress(s.ctx, &marketplace.QuerySellOrdersByAddressRequest{
		Address:    otherAddr.String(),
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))

	// addr with no sell orders should just return empty slice
	res, err = s.k.SellOrdersByAddress(s.ctx, &marketplace.QuerySellOrdersByAddressRequest{
		Address:    noOrdersAddr.String(),
		Pagination: nil,
	})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.SellOrders))

	// bad address should fail
	res, err = s.k.SellOrdersByAddress(s.ctx, &marketplace.QuerySellOrdersByAddressRequest{
		Address:    "foobar1vlk23jrkl",
		Pagination: nil,
	})
	assert.ErrorContains(t, err, "decoding bech32 failed")
}

func insertSellOrder(t *testing.T, s *baseSuite, addr sdk.AccAddress, batchId uint64) {
	sellOrder := &api.SellOrder{
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
