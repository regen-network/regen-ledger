package orm

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/testing/protocmp"

	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/orm/v2/types/kvlayout"

	"github.com/cosmos/cosmos-sdk/store/listenkv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

type TestDecoder struct {
	schema *Schema
	ops    []Op
}

func (t *TestDecoder) OnWrite(_ storetypes.StoreKey, key []byte, value []byte, delete bool) error {
	entry, err := t.schema.Decode(key, value)
	t.ops = append(t.ops, Op{
		Err:    err,
		Entry:  entry,
		Delete: delete,
	})
	return nil
}

func (t *TestDecoder) ConsumeOps() []Op {
	ops := t.ops
	t.ops = nil
	return ops
}

type Op struct {
	Err    error
	Entry  kvlayout.Entry
	Delete bool
}

func (o Op) String() string {
	str := ""
	if o.Delete {
		str += "!"
	}
	if o.Entry != nil {
		str += fmt.Sprintf("%s", o.Entry)
	}
	if o.Err != nil {
		str += fmt.Sprintf("(ERR:%v)", o.Err)
	}
	str += ""
	return str
}

func TestClient(t *testing.T) {
	schema, err := BuildSchema(FileDescriptor(0, testpb.File__1_proto))
	assert.NilError(t, err)
	clientConn := &ClientConn{schema}
	decoder := &TestDecoder{schema: schema}
	kv := listenkv.NewStore(mem.NewStore(), nil, []storetypes.WriteListener{decoder})
	client := clientConn.Open(kv)

	assert.NilError(t, client.Save(
		&testpb.A{
			UINT32: 4,
			UINT64: 10,
			STRING: "abc",
		},
	))
	t.Logf("%+v", decoder.ops)
	assert.DeepEqual(t, []Op{
		{
			Entry: kvlayout.PrimaryEntry{
				Key: []protoreflect.Value{
					protoreflect.ValueOfUint32(4),
					protoreflect.ValueOfUint64(10),
					protoreflect.ValueOfString("abc"),
				},
				Value: &testpb.A{},
			},
		},
	}, decoder.ConsumeOps(), protocmp.Transform())

	data := []proto.Message{
		&testpb.A{
			UINT32: 4,
			UINT64: 10,
			STRING: "foo",
		},
		&testpb.A{
			UINT32: 4,
			UINT64: 11,
			STRING: "abc",
		},
		&testpb.A{
			UINT32: 5,
			UINT64: 3,
			STRING: "foo",
		},
	}

	assert.NilError(t, client.Save(data...))

	for i, x := range data {
		assert.Assert(t, client.Has(x), "data", i)
	}

	it := client.List(&testpb.A{})
	defer it.Close()
	require.NotNil(t, it)
	var a1 testpb.A
	have, err := it.Next(&a1)
	assert.Assert(t, have)
	assert.NilError(t, err)
	have, err = it.Next(&a1)
	assert.Assert(t, have)
	assert.NilError(t, err)
	have, err = it.Next(&a1)
	assert.Assert(t, have)
	assert.NilError(t, err)
	have, err = it.Next(&a1)
	assert.Assert(t, have)
	assert.NilError(t, err)
	// no more elements
	have, err = it.Next(&a1)
	require.False(t, have)
	assert.NilError(t, err)
}
