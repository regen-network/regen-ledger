package ordermatch_test

import (
	"context"
	"os"
	"testing"

	ormtestutil "github.com/regen-network/regen-ledger/types/testutil/orm"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/ordermatch"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
)

var testModuleSchema = ormdb.ModuleSchema{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: ecocreditv1beta1.File_regen_ecocredit_v1beta1_state_proto,
		2: marketplacev1beta1.File_regen_ecocredit_marketplace_v1beta1_state_proto,
		3: orderbookv1beta1.File_regen_ecocredit_orderbook_v1beta1_memory_proto,
	},
}

func TestMatcher(t *testing.T) {
	db, err := ormdb.NewModuleDB(testModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)

	matcher, err := ordermatch.NewMatcher(db)
	assert.NilError(t, err)

	ctx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	bz, err := os.ReadFile("testdata/in/scenario1.json")
	assert.NilError(t, err)
	jsonSource, err := ormjson.NewRawMessageSource(bz)
	assert.NilError(t, err)
	assert.NilError(t, db.ImportJSON(ctx, jsonSource))

	assert.NilError(t, matcher.Reload(ctx))

	ormtestutil.AssertGoldenDB(t, db, ctx, "out/scenario1.after_reload.json")

	assertOrderBookOrder(t, db, ctx,
		pair{2, 1},
		pair{2, 2},
		pair{1, 1},
		pair{3, 3},
		pair{3, 1},
	)

	// TODO test project locations
	// TODO test dates
}

type pair struct {
	buyOrderId  uint64
	sellOrderId uint64
}

func assertOrderBookOrder(t *testing.T, db ormdb.ModuleDB, ctx context.Context, pairs ...pair) {
	store, err := orderbookv1beta1.NewBuyOrderSellOrderMatchStore(db)
	assert.NilError(t, err)
	it, err := store.List(ctx, orderbookv1beta1.BuyOrderSellOrderMatchMarketIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey{})
	assert.NilError(t, err)
	for _, p := range pairs {
		assert.Assert(t, it.Next())
		val, err := it.Value()
		assert.NilError(t, err)
		assert.Equal(t, p.buyOrderId, val.BuyOrderId)
		assert.Equal(t, p.sellOrderId, val.SellOrderId)
	}
	assert.Assert(t, !it.Next(), "should have reached end of order matches")
}
