package v2

import (
	"bytes"
	"fmt"
	"io"

	prefixstore "github.com/cosmos/cosmos-sdk/store/prefix"

	sdkstore "github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/types"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type tableStore struct {
	numPrimaryKeyFields int
	pkFields            []protoreflect.FieldDescriptor
	prefix              []byte
	pkCodec             *keyCodec
	indexers            []*indexer
	indexerMap          map[string]*indexer
}

func (d tableStore) isStore() {}

func (d tableStore) primaryKey(message proto.Message) ([]protoreflect.Value, []byte, error) {
	pkValues := d.primaryKeyValues(message)

	pkBuf := &bytes.Buffer{}
	pkBuf.Write(d.prefix)
	pkBuf.WriteByte(0) // primary key table always prefixed with 0
	err := d.pkCodec.encode(pkValues, pkBuf, false)
	if err != nil {
		return nil, nil, err
	}

	return pkValues, pkBuf.Bytes(), nil
}

func (d tableStore) primaryKeyValues(message proto.Message) []protoreflect.Value {
	refm := message.ProtoReflect()
	// encode primary key
	pkValues := make([]protoreflect.Value, d.numPrimaryKeyFields)
	for i, f := range d.pkFields {
		pkValues[i] = refm.Get(f)
	}

	return pkValues
}

func (d tableStore) Create(kv sdkstore.KVStore, message proto.Message) error {
	_, err := d.save(kv, message, true)
	return err
}

func (d tableStore) Read(kv sdkstore.KVStore, message proto.Message) (bool, error) {
	pkValues, pk, err := d.primaryKey(message)
	if err != nil {
		return false, err
	}

	bz := kv.Get(pk)
	if bz == nil {
		return false, nil
	}

	err = proto.Unmarshal(bz, message)
	if err != nil {
		return true, err
	}

	refm := message.ProtoReflect()

	// rehydrate primary key
	for i, f := range d.pkFields {
		refm.Set(f, pkValues[i])
	}

	return true, nil
}

func (d tableStore) Has(kv sdkstore.KVStore, message proto.Message) bool {
	_, pk, err := d.primaryKey(message)
	if err != nil {
		return false
	}

	return kv.Has(pk)
}

func (d tableStore) Save(kv sdkstore.KVStore, message proto.Message) (bool, error) {
	return d.save(kv, message, false)
}

func (d tableStore) Delete(kv sdkstore.KVStore, message proto.Message) error {
	_, pk, err := d.primaryKey(message)
	if err != nil {
		return err
	}

	// clear indexes
	for _, idx := range d.indexers {
		err := idx.onCreate(kv, message.ProtoReflect())
		if err != nil {
			return err
		}
	}

	// delete object
	kv.Delete(pk)

	return nil
}

func (d tableStore) save(kv sdkstore.KVStore, message proto.Message, create bool) (bool, error) {
	pkValues, pk, err := d.primaryKey(message)
	if err != nil {
		return false, err
	}

	refm := message.ProtoReflect()
	bz := kv.Get(pk)
	var existing proto.Message
	if bz != nil {
		if create {
			return true, fmt.Errorf("object of type %T with primary key %s already exists, can't create", message, pkValues)
		}

		existing = refm.New().Interface()
		err = proto.Unmarshal(bz, existing)
		if err != nil {
			return true, err
		}
	}

	// temporarily clear primary key
	for _, f := range d.pkFields {
		refm.Clear(f)
	}

	// store object
	bz, err = proto.Marshal(message)
	kv.Set(pk, bz)

	// set primary key again
	for i, f := range d.pkFields {
		refm.Set(f, pkValues[i])
	}

	created := existing != nil

	// set indexes
	existingRef := existing.ProtoReflect()
	for _, idx := range d.indexers {
		if existing == nil {
			err = idx.onCreate(kv, refm)
		} else {
			err = idx.onUpdate(kv, refm, existingRef)
		}
		if err != nil {
			return created, err
		}
	}

	return created, nil
}

func (d *tableStore) List(kv types.KVStore, message proto.Message, options ...ListOption) Iterator {
	opts := gatherListOptions(options)
	if opts.indexHint != "" {
		idx, ok := d.indexerMap[opts.indexHint]
		if !ok {
			return errIterator{err: fmt.Errorf("can't find indexer %s", opts.indexHint)}
		}

		prefixStore := prefixstore.NewStore(kv, idx.prefix)

		refm := message.ProtoReflect()
		var values []protoreflect.Value
		for _, f := range idx.indexFields {
			values = append(values, refm.Get(f))
		}
		buf := &bytes.Buffer{}
		err := idx.codec.encode(values, buf, true)
		if err != nil && err != io.EOF {
			return errIterator{err: err}
		}
		prefix := buf.Bytes()

		var iterator types.Iterator
		if !opts.reverse {
			iterator = prefixStore.Iterator(prefix, nil)
		} else {
			iterator = prefixStore.ReverseIterator(prefix, nil)
		}
		return &idxIterator{
			kv:        kv,
			store:     d,
			iterator:  iterator,
			start:     true,
			pkDecoder: idx.codec.pkDecoder,
		}
	} else {
		// first make prefix store for pk table
		buf := &bytes.Buffer{}
		buf.Write(d.prefix)
		buf.WriteByte(0)
		prefixStore := prefixstore.NewStore(kv, buf.Bytes())

		pkValues := d.primaryKeyValues(message)
		buf = &bytes.Buffer{}
		err := d.pkCodec.encode(pkValues, buf, true)
		if err != nil && err != io.EOF {
			return errIterator{err: err}
		}

		prefix := buf.Bytes()
		var iterator types.Iterator
		if !opts.reverse {
			iterator = prefixStore.Iterator(prefix, nil)
		} else {
			iterator = prefixStore.ReverseIterator(prefix, nil)
		}
		return &pkIterator{
			kv:       kv,
			store:    d,
			iterator: iterator,
			start:    true,
		}
	}
}

type pkIterator struct {
	kv       types.KVStore
	store    *tableStore
	iterator types.Iterator
	start    bool
}

func (t *pkIterator) isIterator() {}

func (t *pkIterator) Next(message proto.Message) (bool, error) {
	if t.start {
		t.start = false
	} else {
		t.iterator.Next()
	}

	if !t.iterator.Valid() {
		return false, nil
	}

	bz := t.iterator.Value()
	err := proto.Unmarshal(bz, message)
	if err != nil {
		return false, err
	}

	pkValues, err := t.store.pkCodec.decode(bytes.NewReader(t.iterator.Key()))
	if err != nil {
		return false, err
	}

	// rehydrate primary key
	refm := message.ProtoReflect()
	for i := 0; i < t.store.numPrimaryKeyFields; i++ {
		field := t.store.pkFields[i]
		refm.Set(field, pkValues[i])
	}

	return true, nil
}

type idxIterator struct {
	kv        types.KVStore
	store     *tableStore
	iterator  types.Iterator
	start     bool
	pkDecoder func(r *bytes.Reader) ([]protoreflect.Value, error)
}

func (t *idxIterator) isIterator() {}

func (t *idxIterator) Next(message proto.Message) (bool, error) {
	if t.start {
		t.start = false
	} else {
		t.iterator.Next()
	}

	if !t.iterator.Valid() {
		return false, nil
	}

	pkValues, err := t.pkDecoder(bytes.NewReader(t.iterator.Key()))
	if err != nil {
		return false, err
	}

	buf := &bytes.Buffer{}
	buf.Write(t.store.prefix)
	buf.WriteByte(0)
	err = t.store.pkCodec.encode(pkValues, buf, false)
	if err != nil {
		return false, err
	}
	pk := buf.Bytes()

	bz := t.kv.Get(pk)
	err = proto.Unmarshal(bz, message)
	if err != nil {
		return false, err
	}

	// rehydrate primary key
	refm := message.ProtoReflect()
	for i := 0; i < t.store.numPrimaryKeyFields; i++ {
		field := t.store.pkFields[i]
		refm.Set(field, pkValues[i])
	}

	return true, nil
}
