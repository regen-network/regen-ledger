package table

import (
	"bytes"
	"fmt"
	"io"

	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (s *Store) List(kv store.KVStore, message proto.Message, opts *list.Options) list.Iterator {
	if opts.IndexHint != "" {
		idx, ok := s.IndexerMap[opts.IndexHint]
		if !ok {
			return list.ErrIterator{Err: fmt.Errorf("can't find indexer %s", opts.IndexHint)}
		}

		refm := message.ProtoReflect()
		var values []protoreflect.Value
		for _, f := range idx.IndexFields {
			values = append(values, refm.Get(f))
		}
		buf := &bytes.Buffer{}
		buf.Write(idx.Prefix)
		err := idx.Codec.Encode(values, buf, true)
		if err != nil && err != io.EOF {
			return list.ErrIterator{Err: err}
		}
		prefix := buf.Bytes()

		var iterator store.KVStoreIterator
		if !opts.Reverse {
			iterator = kv.Iterator(prefix, nil)
		} else {
			iterator = kv.ReverseIterator(prefix, nil)
		}
		return &idxIterator{
			kv:        kv,
			store:     s,
			iterator:  iterator,
			start:     true,
			pkDecoder: idx.Codec.PKDecoder,
			prefix:    idx.Prefix,
		}
	} else {
		// first make prefix store for pk table
		buf := &bytes.Buffer{}

		pkValues := s.primaryKeyValues(message)
		buf = &bytes.Buffer{}
		buf.Write(s.PkPrefix)
		err := s.PkCodec.Encode(pkValues, buf, true)
		if err != nil && err != io.EOF {
			return list.ErrIterator{Err: err}
		}

		prefix := buf.Bytes()
		var iterator store.KVStoreIterator
		if !opts.Reverse {
			iterator = kv.Iterator(prefix, nil)
		} else {
			iterator = kv.ReverseIterator(prefix, nil)
		}
		return &pkIterator{
			kv:       kv,
			store:    s,
			iterator: iterator,
			start:    true,
		}
	}
}

type pkIterator struct {
	kv       store.KVStore
	store    *Store
	iterator store.KVStoreIterator
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

	key := bytes.NewReader(t.iterator.Key()[len(t.store.PkPrefix):])
	pkValues, err := t.store.PkCodec.Decode(key)
	if err != nil {
		return false, err
	}

	// rehydrate primary key
	refm := message.ProtoReflect()
	for i := 0; i < t.store.NumPrimaryKeyFields; i++ {
		field := t.store.PkFields[i]
		refm.Set(field, pkValues[i])
	}

	return true, nil
}

type idxIterator struct {
	kv        store.KVStore
	store     *Store
	iterator  store.KVStoreIterator
	start     bool
	pkDecoder func(r *bytes.Reader) ([]protoreflect.Value, error)
	prefix    []byte
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

	key := t.iterator.Key()[len(t.prefix):]
	pkValues, err := t.pkDecoder(bytes.NewReader(key))
	if err != nil {
		return false, err
	}

	buf := &bytes.Buffer{}
	buf.Write(t.store.Prefix)
	buf.WriteByte(0)
	err = t.store.PkCodec.Encode(pkValues, buf, false)
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
	for i := 0; i < t.store.NumPrimaryKeyFields; i++ {
		field := t.store.PkFields[i]
		refm.Set(field, pkValues[i])
	}

	return true, nil
}
