package ormtestutil

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/gibson042/canonicaljson-go"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

// AssertGoldenDB does golden testing of a database's state using gotest.tools/v3/golden
// a JSON file on disk. By default, the JSON emitted by ModuleDB.ExportJSON
// isn't suitable for golden testing because it is non-deterministic. This
// method ensures that JSON state is exported deterministically before comparing.
// Note that this deterministic serialization may change
func AssertGoldenDB(t assert.TestingT, db ormdb.ModuleDB, ctx context.Context, goldenFile string) {
	target := ormjson.NewRawMessageTarget()
	assert.NilError(t, db.ExportJSON(ctx, target))
	bz, err := target.JSON()
	assert.NilError(t, err)
	assertGoldenJson(t, bz, goldenFile)
}

func assertGoldenJson(t assert.TestingT, bz []byte, goldenFile string) {
	var rawJson json.RawMessage
	err := json.Unmarshal(bz, &rawJson)
	assert.NilError(t, err)
	bz, err = canonicaljson.Marshal(rawJson)
	assert.NilError(t, err)
	buf := &bytes.Buffer{}
	assert.NilError(t, json.Indent(buf, bz, "", "  "))
	golden.Assert(t, buf.String(), goldenFile)
}
