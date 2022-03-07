package orderbook_test

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog"

	"github.com/regen-network/regen-ledger/x/ecocredit/orderbook/fill"

	"github.com/regen-network/regen-ledger/x/ecocredit/orderbook"

	"github.com/regen-network/regen-ledger/x/ecocredit/orderbook/testutil"

	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

var testModuleSchema = ormdb.ModuleSchema{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: ecocreditv1beta1.File_regen_ecocredit_v1beta1_state_proto,
		2: marketplacev1beta1.File_regen_ecocredit_marketplace_v1beta1_state_proto,
		3: orderbookv1beta1.File_regen_ecocredit_orderbook_v1beta1_memory_proto,
	},
}

func Test1(t *testing.T) {
	db, err := ormdb.NewModuleDB(testModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)

	logger := zerolog.New(zerolog.NewConsoleWriter())
	fillMgr, err := fill.NewManager(db, testutil.NewTestTransferManager(logger), logger)
	assert.NilError(t, err)
	orderBook, err := orderbook.NewOrderBook(db, fillMgr, logger)
	assert.NilError(t, err)

	ctx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	bz, err := os.ReadFile("testdata/in/scenario1.json")
	assert.NilError(t, err)
	jsonSource, err := ormjson.NewRawMessageSource(bz)
	assert.NilError(t, err)
	assert.NilError(t, db.ImportJSON(ctx, jsonSource))

	assert.NilError(t, orderBook.Reload(ctx))

	exportJson(t, db, ctx, "testdata/out/scenario1.after_reload.json")

	assert.NilError(t, orderBook.ProcessBatch(ctx))

	exportJson(t, db, ctx, "testdata/out/scenario1.after_process.json")
}

func exportJson(t assert.TestingT, db ormdb.ModuleDB, ctx context.Context, file string) {
	jsonSink := ormjson.NewRawMessageTarget()
	assert.NilError(t, db.ExportJSON(ctx, jsonSink))
	bz, err := jsonSink.JSON()
	assert.NilError(t, err)
	assert.NilError(t, os.WriteFile(file, bz, 0644))
}
