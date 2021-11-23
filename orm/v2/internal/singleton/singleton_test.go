package singleton

import (
	"testing"

	"github.com/regen-network/regen-ledger/orm/v2/orm"

	"github.com/regen-network/regen-ledger/orm/v2/model/ormtable"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"
)

func TestSingleton(t *testing.T) {
	b1 := &testpb.B{X: "abc"}
	_, err := BuildStore(nil, &ormpb.SingletonDescriptor{Id: 0}, b1.ProtoReflect().Type())
	assert.ErrorContains(t, err, ormerrors.InvalidTableId.Error())

	store, err := BuildStore(nil, &ormpb.SingletonDescriptor{Id: 1}, b1.ProtoReflect().Type())
	assert.NilError(t, err)

	kv := mem.NewStore()

	// read empty
	found, err := store.Get(kv, nil, b1, nil)
	assert.Assert(t, !found)
	assert.NilError(t, err)

	// create
	err = store.Save(kv, b1, ormtable.SAVE_MODE_CREATE)
	assert.NilError(t, err)

	// read
	var b2 testpb.B
	found, err = store.Get(kv, nil, &b2, nil)
	assert.Assert(t, found)
	assert.NilError(t, err)
	assert.Equal(t, b1.X, b2.X)

	// create a second time works (singleton tables don't care)
	b1.X = "def"
	err = store.Save(kv, b1, ormtable.SAVE_MODE_CREATE)
	assert.NilError(t, err)

	// save succeeds
	err = store.Save(kv, b1, ormtable.SAVE_MODE_UPDATE)
	assert.NilError(t, err)

	// read
	found, err = store.Get(kv, nil, &b2, nil)
	assert.Assert(t, found)
	assert.NilError(t, err)
	assert.Equal(t, b1.X, b2.X)

	// iterator just returns one value always
	it := store.List(kv, &orm.ListOptions{})
	assert.Assert(t, it != nil)
	found, err = it.Next(&b2)
	assert.Assert(t, found)
	assert.NilError(t, err)
	assert.Equal(t, b1.X, b2.X)
	found, err = it.Next(&b2)
	assert.Assert(t, !found)
	assert.NilError(t, err)
	found, err = it.Next(&b2) // next always does the same thing
	assert.Assert(t, !found)
	assert.NilError(t, err)

	// delete
	err = store.Delete(kv, nil)
	assert.NilError(t, err)
	err = store.Delete(kv, nil)
	assert.NilError(t, err) // deleting twice is a no-op

	// can't read
	found, err = store.Get(kv, nil, b1, nil)
	assert.Assert(t, !found)
	assert.NilError(t, err)
}
