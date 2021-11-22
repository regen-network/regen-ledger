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
	FieldNames  string
}

var sentinel = []byte{0}

func (i *Indexer) onCreate(kv store.KVStore, message protoreflect.Message) error {
	k, err := i.getKey(message)
	if err != nil {
		return err
	}

	kv.Set(k, sentinel)
	return nil
}

func (i *Indexer) getKey(message protoreflect.Message) ([]byte, error) {
	values := i.Codec.GetValues(message)
	buf := &bytes.Buffer{}
	err := i.Codec.Encode(values, buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (i *Indexer) onDelete(kv store.KVStore, message protoreflect.Message) error {
	k, err := i.getKey(message)
	if err != nil {
		return err
	}

	kv.Delete(k)
	return nil
}

func (i *Indexer) onUpdate(kv store.KVStore, new protoreflect.Message, existing protoreflect.Message) error {
	newKey, err := i.getKey(new)
	if err != nil {
		return err
	}
	existingKey, err := i.getKey(existing)
	if err != nil {
		return err
	}

	if bytes.Equal(newKey, existingKey) {
		return nil
	}

	kv.Delete(existingKey)
	kv.Set(newKey, sentinel)
	return nil
}
