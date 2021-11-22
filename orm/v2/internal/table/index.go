package table

import (
	"bytes"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Index struct {
	IndexFields []protoreflect.FieldDescriptor
	Prefix      []byte
	Codec       *ormkey.IndexKeyCodec
	FieldNames  string
}

var sentinel = []byte{0}

func (i *Index) onCreate(kv kv.KVStore, message protoreflect.Message) error {
	k, err := i.getKey(message)
	if err != nil {
		return err
	}

	kv.Set(k, sentinel)
	return nil
}

func (i *Index) getKey(message protoreflect.Message) ([]byte, error) {
	values := i.Codec.GetValues(message)
	buf := &bytes.Buffer{}
	err := i.Codec.EncodeWriter(values, buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (i *Index) onDelete(kv kv.KVStore, message protoreflect.Message) error {
	k, err := i.getKey(message)
	if err != nil {
		return err
	}

	kv.Delete(k)
	return nil
}

func (i *Index) onUpdate(kv kv.KVStore, new protoreflect.Message, existing protoreflect.Message) error {
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
