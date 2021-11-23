package ormkv_test

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/testing/protocmp"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/regen-network/regen-ledger/orm/v2/orm"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"

	"github.com/regen-network/regen-ledger/orm/v2/backend/ormkv"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm/v2/model/ormschema"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/store/listenkv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

type TestDecoder struct {
	schema *ormschema.Schema
	ops    []Op
}

func (t *TestDecoder) OnWrite(_ storetypes.StoreKey, key []byte, value []byte, delete bool) error {
	entry, err := t.schema.Decode(key, value)
	op := Op{
		Err:    err,
		Entry:  entry,
		Delete: delete,
	}
	fmt.Printf("  WRITE %s\n", op)
	t.ops = append(t.ops, op)
	return nil
}

func (t *TestDecoder) ConsumeOps() []Op {
	ops := t.ops
	t.ops = nil
	return ops
}

type Op struct {
	Err    error
	Entry  ormdecode.Entry
	Delete bool
}

func (o Op) String() string {
	str := ""
	if o.Delete {
		str += "-"
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

type kvReadListener struct {
	kv.KVStore
	schema *ormschema.Schema
}

var _ kv.KVStore = &kvReadListener{}

func fmtEmtry(entry ormdecode.Entry, err error) string {
	if err != nil {
		return fmt.Sprintf("ERROR %v", err)
	} else {
		return fmt.Sprintf("%s", entry)
	}
}

func (r kvReadListener) Get(key []byte) []byte {
	value := r.KVStore.Get(key)
	fmt.Printf("  READ %s\n", fmtEmtry(r.schema.Decode(key, value)))
	return value
}

func (r kvReadListener) Has(key []byte) bool {
	value := r.KVStore.Get(key)
	fmt.Printf("  READ HAS %s\n", fmtEmtry(r.schema.Decode(key, value)))
	return r.KVStore.Has(key)
}

func (r kvReadListener) Iterator(start, end []byte) kv.KVStoreIterator {
	fmt.Printf("  ITERATE %x -> %x\n", start, end)
	it := r.KVStore.Iterator(start, end)
	return &kvIteratorListener{
		it:     it,
		schema: r.schema,
	}
}

func (r kvReadListener) ReverseIterator(start, end []byte) kv.KVStoreIterator {
	fmt.Printf("  ITERATE %x <- %x\n", start, end)
	it := r.KVStore.ReverseIterator(start, end)
	return &kvIteratorListener{
		it:     it,
		schema: r.schema,
	}
}

type kvIteratorListener struct {
	it     kv.KVStoreIterator
	schema *ormschema.Schema
}

func (k kvIteratorListener) Domain() (start []byte, end []byte) {
	return k.it.Domain()
}

func (k kvIteratorListener) Valid() bool {
	valid := k.it.Valid()
	if valid {
		fmt.Printf("    VALID %s\n", fmtEmtry(k.schema.Decode(k.it.Key(), k.it.Value())))
	} else {
		fmt.Printf("    INVALID\n")
	}
	return valid
}

func (k kvIteratorListener) Next() {
	fmt.Printf("    NEXT \n")
	k.it.Next()
}

func (k kvIteratorListener) Key() (key []byte) {
	return k.it.Key()
}

func (k kvIteratorListener) Value() (value []byte) {
	return k.it.Value()
}

func (k kvIteratorListener) Error() error {
	return k.it.Error()
}

func (k kvIteratorListener) Close() error {
	return k.it.Close()
}

type storeListener struct {
	orm.UnimplementedStore
	store orm.Store
}

func marshalJsonMessages(messages ...proto.Message) string {
	var str string
	for _, msg := range messages {
		fullName := msg.ProtoReflect().Descriptor().FullName()
		bz, err := protojson.Marshal(msg)
		if err != nil {
			str += fmt.Sprintf("%s(%s) ", fullName, msg)
		} else {
			str += fmt.Sprintf("%s%s ", fullName, bz)
		}
	}
	return str
}

func (o storeListener) Has(messages ...proto.Message) bool {
	fmt.Printf("has %s\n", marshalJsonMessages(messages...))
	return o.store.Has(messages...)
}

func (o storeListener) Get(messages ...proto.Message) (found bool, err error) {
	fmt.Printf("get %s\n", marshalJsonMessages(messages...))
	return o.store.Get(messages...)
}

func (o storeListener) List(message proto.Message, options *orm.ListOptions) orm.Iterator {
	fmt.Printf("list %s %+v\n", marshalJsonMessages(message), options)
	return listenStoreIterator{it: o.store.List(message, options)}
}

type listenStoreIterator struct {
	orm.UnimplementedIterator
	it orm.Iterator
}

func (l listenStoreIterator) Next(message proto.Message) (bool, error) {
	found, err := l.it.Next(message)
	if err != nil {
		fmt.Printf("    iterator ERROR %v\n", err)
	} else {
		if found {
			fmt.Printf("    next %s\n", marshalJsonMessages(message))
		} else {
			fmt.Printf("    last\n")
		}
	}
	return found, err
}

func (l listenStoreIterator) Cursor() orm.Cursor {
	return l.it.Cursor()
}

func (l listenStoreIterator) Close() {
	l.it.Close()
}

func (o storeListener) Create(messages ...proto.Message) error {
	fmt.Printf("create %s\n", marshalJsonMessages(messages...))
	return o.store.Create(messages...)
}

func (o storeListener) Save(messages ...proto.Message) error {
	fmt.Printf("save %s\n", marshalJsonMessages(messages...))
	return o.store.Save(messages...)
}

func (o storeListener) Delete(messages ...proto.Message) error {
	fmt.Printf("delete %s\n", marshalJsonMessages(messages...))
	return o.store.Delete(messages...)
}

var _ orm.Store = storeListener{}

func TestClient(t *testing.T) {
	schema, err := ormschema.BuildSchema(ormschema.FileDescriptor(0, testpb.File__1_proto))
	assert.NilError(t, err)
	decoder := &TestDecoder{schema: schema}

	var kvStore types.KVStore
	kvStore = mem.NewStore()
	kvStore = listenkv.NewStore(kvStore, nil, []storetypes.WriteListener{decoder})
	store := ormkv.NewStore(schema, &kvReadListener{
		KVStore: kvStore,
		schema:  schema,
	})
	store = &storeListener{store: store}

	data := []proto.Message{
		&testpb.A{
			UINT32: 4,
			UINT64: 10,
			STRING: "abc",
			BYTES:  []byte{0, 1, 2},
		},
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

	assert.NilError(t, store.Save(data...))

	for i, x := range data {
		assert.Assert(t, store.Has(x), "data", i)
	}

	// forward iteration
	it := store.List(&testpb.A{}, nil)
	defer it.Close()
	require.NotNil(t, it)
	var acopy testpb.A
	for i := 0; i < len(data); i++ {
		have, err := it.Next(&acopy)
		assert.Assert(t, have)
		assert.NilError(t, err)
		AssertProtoEqual(t, data[i], &acopy)
	}
	// no more elements
	have, err := it.Next(&acopy)
	require.False(t, have)
	assert.NilError(t, err)

	// reverse iteration
	it = store.List(&testpb.A{}, &orm.ListOptions{
		Reverse: true,
	})
	defer it.Close()
	require.NotNil(t, it)
	for i := len(data) - 1; i >= 0; i-- {
		have, err := it.Next(&acopy)
		assert.Assert(t, have)
		assert.NilError(t, err)
		AssertProtoEqual(t, data[i], &acopy)
	}
	// no more elements
	have, err = it.Next(&acopy)
	require.False(t, have)
	assert.NilError(t, err)

	// condition
	it = store.List(&testpb.A{}, &orm.ListOptions{
		Prefix: []protoreflect.Value{protoreflect.ValueOfUint32(4)},
	})
	defer it.Close()
	require.NotNil(t, it)
	for i := 0; i < 3; i++ {
		have, err := it.Next(&acopy)
		assert.Assert(t, have)
		assert.NilError(t, err)
		AssertProtoEqual(t, data[i], &acopy)
	}
	// no more elements
	have, err = it.Next(&acopy)
	require.False(t, have)
	assert.NilError(t, err)

	// use index
	it = store.List(&testpb.A{}, &orm.ListOptions{
		Index: "UINT64,STRING",
	})
	defer it.Close()
	require.NotNil(t, it)
	for _, i := range []int{3, 0, 1, 2} {
		have, err := it.Next(&acopy)
		assert.Assert(t, have)
		assert.NilError(t, err)
		AssertProtoEqual(t, data[i], &acopy)
	}
	// no more elements
	have, err = it.Next(&acopy)
	require.False(t, have)
	assert.NilError(t, err)

	// use index and prefix
	it = store.List(&testpb.A{}, &orm.ListOptions{
		Index:  "UINT64,STRING",
		Prefix: []protoreflect.Value{protoreflect.ValueOfUint64(10)},
	})
	defer it.Close()
	require.NotNil(t, it)
	for _, i := range []int{0, 1} {
		have, err := it.Next(&acopy)
		assert.Assert(t, have)
		assert.NilError(t, err)
		AssertProtoEqual(t, data[i], &acopy)
	}
	// no more elements
	have, err = it.Next(&acopy)
	require.False(t, have)
	assert.NilError(t, err)

	// use index and start and end
	it = store.List(&testpb.A{}, &orm.ListOptions{
		Index:   "STRING,UINT32",
		Reverse: true,
		Start: []protoreflect.Value{
			protoreflect.ValueOfString("abc"),
			protoreflect.ValueOfUint32(4),
		},
		End: []protoreflect.Value{
			protoreflect.ValueOfString("foo"),
			protoreflect.ValueOfUint32(4),
		},
	})
	defer it.Close()
	require.NotNil(t, it)
	for _, i := range []int{1, 2, 0} {
		have, err := it.Next(&acopy)
		assert.Assert(t, have)
		assert.NilError(t, err)
		AssertProtoEqual(t, data[i], &acopy)
	}
	// no more elements
	have, err = it.Next(&acopy)
	require.False(t, have)
	assert.NilError(t, err)
}

func AssertProtoEqual(t assert.TestingT, x, y proto.Message) {
	assert.DeepEqual(t, x, y, protocmp.Transform())
}
