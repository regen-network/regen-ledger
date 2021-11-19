package v2

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/store/types"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func makeIndexKeyCodec(indexFields []protoreflect.FieldDescriptor, primaryKeyFields []protoreflect.FieldDescriptor) (*keyCodec, error) {
	indexFieldMap := map[protoreflect.Name]int{}
	pkFieldOrderMap := map[int]int{}

	var keyFields []protoreflect.FieldDescriptor
	for i, f := range indexFields {
		indexFieldMap[f.Name()] = i
		keyFields = append(keyFields, f)
	}

	for j, f := range primaryKeyFields {
		if i, ok := indexFieldMap[f.Name()]; ok {
			pkFieldOrderMap[j] = i
			continue
		}
		keyFields = append(keyFields, f)
		pkFieldOrderMap[j] = j
	}

	cdc, err := makeKeyCodec(keyFields, false)
	if err != nil {
		return nil, err
	}

	numPrimaryKeyFields := len(primaryKeyFields)
	cdc.pkDecoder = func(r *bytes.Reader) ([]protoreflect.Value, error) {
		fields, err := cdc.decode(r)
		if err != nil {
			return nil, err
		}

		pkValues := make([]protoreflect.Value, numPrimaryKeyFields)

		for i := 0; i < numPrimaryKeyFields; i++ {
			pkValues[i] = fields[pkFieldOrderMap[i]]
		}

		return pkValues, nil
	}

	return cdc, nil
}

type indexer struct {
	indexFields []protoreflect.FieldDescriptor
	prefix      []byte
	codec       *keyCodec
}

var sentinel = []byte{0}

func (i *indexer) onCreate(kv types.KVStore, message protoreflect.Message) error {
	return i.iterateFields(func(key []byte) {
		kv.Set(key, sentinel)
	}, message, nil, i.indexFields)
}

func (i *indexer) onDelete(kv types.KVStore, message protoreflect.Message) error {
	return i.iterateFields(func(key []byte) {
		kv.Delete(key)
	}, message, nil, i.indexFields)
}

func (i *indexer) onUpdate(kv types.KVStore, message protoreflect.Message, existing protoreflect.Message) error {
	keep := map[string]bool{}
	err := i.iterateFields(func(key []byte) {
		keep[string(key)] = true
		kv.Set(key, sentinel)
	}, message, nil, i.indexFields)
	if err != nil {
		return err
	}

	return i.iterateFields(func(key []byte) {
		if !keep[string(key)] {
			kv.Delete(key)
		}
	}, existing, nil, i.indexFields)
}

func (idx *indexer) iterateFields(f func(key []byte), message protoreflect.Message, values []protoreflect.Value, nextFields []protoreflect.FieldDescriptor) error {
	if len(nextFields) == 0 {
		buf := &bytes.Buffer{}
		err := idx.codec.encode(values, buf, false)
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
