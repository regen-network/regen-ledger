package table

import (
	"bytes"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"
	"github.com/regen-network/regen-ledger/orm/v2/orm"

	"google.golang.org/protobuf/proto"
)

func (s *TableModel) List2(kvStore kv.ReadKVStore, keyPrefix []protoreflect.Value, opts *orm.ListOptions) orm.Iterator {
	panic("TODO")
}

func (s *TableModel) List(kvStore kv.ReadKVStore, condition proto.Message, opts *orm.ListOptions) orm.Iterator {
	if opts == nil {
		opts = &orm.ListOptions{}
	}

	var cdc ormkey.CodecI
	var idx *Index
	if opts.UseIndex != "" {
		var ok bool
		idx, ok = s.IndexesByFields[opts.UseIndex]
		if !ok {
			return orm.ErrIterator{Err: fmt.Errorf("can't find indexer %s", opts.UseIndex)}
		}
		cdc = idx.Codec.Codec
	} else {
		cdc = s.PkCodec
	}

	var start, end []byte
	var err error

	_, prefix, err := cdc.EncodePartial(condition.ProtoReflect())
	if err != nil {
		return orm.ErrIterator{Err: err}
	}

	if opts.Cursor != nil && !opts.Reverse {
		start = opts.Cursor
	} else {
		start = prefix
	}

	if opts.Cursor != nil && opts.Reverse {
		start = opts.Cursor
	} else {
		end = storetypes.PrefixEndBytes(prefix)
	}

	var iterator kv.KVStoreIterator
	if !opts.Reverse {
		iterator = kvStore.Iterator(start, end)
	} else {
		iterator = kvStore.ReverseIterator(start, end)
	}

	if idx != nil {
		return &idxIterator{
			kv:       kvStore,
			store:    s,
			iterator: iterator,
			start:    true,
			cdc:      idx.Codec,
		}
	} else {
		return &pkIterator{
			kv:       kvStore,
			store:    s,
			iterator: iterator,
			start:    true,
		}
	}
}

type pkIterator struct {
	orm.UnimplementedIterator

	kv       kv.ReadKVStore
	store    *TableModel
	iterator kv.KVStoreIterator
	start    bool
}

func (t *pkIterator) Close() {
	_ = t.iterator.Close()
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

func (t pkIterator) Cursor() orm.Cursor {
	return t.iterator.Key()
}

type idxIterator struct {
	orm.UnimplementedIterator

	kv       kv.ReadKVStore
	store    *TableModel
	iterator kv.KVStoreIterator
	start    bool
	cdc      *ormkey.IndexKeyCodec
}

func (t *idxIterator) Close() {
	_ = t.iterator.Close()
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
	pkValues, err := t.cdc.ReadPrimaryKey(bytes.NewReader(k))
	if err != nil {
		return false, err
	}

	buf := &bytes.Buffer{}
	err = t.store.PkCodec.EncodeWriter(pkValues, buf)
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

func (t idxIterator) Cursor() orm.Cursor {
	return t.iterator.Key()
}
