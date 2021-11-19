package table

import (
	"bytes"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
)

type Indexer struct {
	IndexFields []protoreflect.FieldDescriptor
	Prefix      []byte
	Codec       *key.Codec
}

var sentinel = []byte{0}

func (i *Indexer) onCreate(kv store.KVStore, message protoreflect.Message) error {
	return i.iterateFields(func(key []byte) {
		kv.Set(key, sentinel)
	}, message, nil, i.IndexFields)
}

func (i *Indexer) onDelete(kv store.KVStore, message protoreflect.Message) error {
	return i.iterateFields(func(key []byte) {
		kv.Delete(key)
	}, message, nil, i.IndexFields)
}

func (i *Indexer) onUpdate(kv store.KVStore, message protoreflect.Message, existing protoreflect.Message) error {
	keep := map[string]bool{}
	err := i.iterateFields(func(key []byte) {
		keep[string(key)] = true
		kv.Set(key, sentinel)
	}, message, nil, i.IndexFields)
	if err != nil {
		return err
	}

	return i.iterateFields(func(key []byte) {
		if !keep[string(key)] {
			kv.Delete(key)
		}
	}, existing, nil, i.IndexFields)
}

func (idx *Indexer) iterateFields(f func(key []byte), message protoreflect.Message, values []protoreflect.Value, nextFields []protoreflect.FieldDescriptor) error {
	if len(nextFields) == 0 {
		buf := &bytes.Buffer{}
		err := idx.Codec.Encode(values, buf, false)
		if err != nil {
			return err
		}

		f(buf.Bytes())
		return nil
	} else {
		field := nextFields[0]
		val := message.Get(field)
		if field.IsList() {
			list := val.List()
			n := list.Len()
			for i := 0; i < n; i++ {
				elem := list.Get(i)
				valuesCopy := make([]protoreflect.Value, len(values))
				copy(valuesCopy, values)
				err := idx.iterateFields(f, message, append(valuesCopy, elem), nextFields[1:])
				if err != nil {
					return err
				}
			}
		} else {
			err := idx.iterateFields(f, message, append(values, val), nextFields[1:])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
