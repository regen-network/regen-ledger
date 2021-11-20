package table

import (
	"bytes"
	"fmt"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"

	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (s *Store) List(kv store.KVStore, opts *list.Options) list.Iterator {
	var cdc *key.Codec
	var idx *Indexer
	if opts.UseIndex != "" {
		var ok bool
		idx, ok = s.IndexersByFields[opts.UseIndex]
		if !ok {
			return list.ErrIterator{Err: fmt.Errorf("can't find indexer %s", opts.UseIndex)}
		}
		cdc = idx.Codec
	}

	var start, end []byte
	var err error
	if opts.Start != nil {
		_, start, err = cdc.EncodePartial(opts.Start.ProtoReflect())
		if err != nil {
			return list.ErrIterator{Err: err}
		}
	}

	if opts.End != nil {
		_, end, err = cdc.EncodePartial(opts.Start.ProtoReflect())
		if err != nil {
			return list.ErrIterator{Err: err}
		}
	}

	var iterator store.KVStoreIterator
	if !opts.Reverse {
		iterator = kv.Iterator(start, end)
	} else {
		iterator = kv.ReverseIterator(start, end)
	}

	if idx != nil {
		return &idxIterator{
			kv:        kv,
			store:     s,
			iterator:  iterator,
			start:     true,
			pkDecoder: idx.Codec.PKDecoder,
		}
	} else {
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

	k := t.iterator.Key()
	pkValues, err := t.store.PkCodec.Decode(bytes.NewReader(k))
	if err != nil {
		return false, err
	}

	// rehydrate primary key
	mref := message.ProtoReflect()
	t.store.PkCodec.SetValues(mref, pkValues)

	return true, nil
}

type idxIterator struct {
	kv        store.KVStore
	store     *Store
	iterator  store.KVStoreIterator
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

	k := t.iterator.Key()
	pkValues, err := t.pkDecoder(bytes.NewReader(k))
	if err != nil {
		return false, err
	}

	buf := &bytes.Buffer{}
	err = t.store.PkCodec.Encode(pkValues, buf)
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
	t.store.PkCodec.SetValues(message.ProtoReflect(), pkValues)

	return true, nil
}
