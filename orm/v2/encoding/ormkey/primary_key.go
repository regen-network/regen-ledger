package ormkey

import (
	"bytes"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"

	"google.golang.org/protobuf/proto"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type PrimaryKeyCodec struct {
	*Codec
	Type protoreflect.MessageType
}

func NewPrimaryKeyCodec(
	prefix []byte,
	messageType protoreflect.MessageType,
	tableDesc *ormpb.TableDescriptor,
) (*PrimaryKeyCodec, error) {
	tableId := tableDesc.Id
	if tableId == 0 {
		return nil, ormerrors.InvalidTableId.Wrapf("table %s", messageType.Descriptor().FullName())
	}

	primaryKeyDescriptor := tableDesc.PrimaryKey
	if primaryKeyDescriptor == nil {
		return nil, ormerrors.MissingPrimaryKey.Wrap(string(messageType.Descriptor().FullName()))
	}

	desc := messageType.Descriptor()
	pkFields, err := GetFieldDescriptors(desc, primaryKeyDescriptor.Fields)
	if err != nil {
		return nil, err
	}

	if primaryKeyDescriptor.AutoIncrement {
		if len(pkFields) != 1 && pkFields[0].Kind() != protoreflect.Uint64Kind {
			return nil, ormerrors.InvalidAutoIncrementKey.Wrapf("got %s for %s", primaryKeyDescriptor.Fields, desc.FullName())
		}
	}

	pkPrefix := MakeUint32Prefix(prefix, tableDesc.Id)
	pkPrefix = MakeUint32Prefix(pkPrefix, 0)

	cdc, err := MakeCodec(pkPrefix, pkFields)

	return &PrimaryKeyCodec{
		Codec: cdc,
		Type:  messageType,
	}, nil
}

func (p PrimaryKeyCodec) DecodeKV(k, v []byte) (ormdecode.Entry, error) {
	vals, err := p.Decode(bytes.NewReader(k))
	if err != nil {
		return nil, err
	}

	msg := p.Type.New().Interface()
	err = proto.Unmarshal(v, msg)
	if err != nil {
		return nil, err
	}

	return ormdecode.PrimaryKeyEntry{
		Key:   vals,
		Value: msg,
	}, nil
}

func (p PrimaryKeyCodec) EncodeKV(entry ormdecode.Entry) (k, v []byte, err error) {
	pkEntry, ok := entry.(ormdecode.PrimaryKeyEntry)
	if !ok {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	if pkEntry.Value.ProtoReflect().Descriptor().FullName() != p.Type.Descriptor().FullName() {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	kbuf := &bytes.Buffer{}
	err = p.Codec.EncodeWriter(pkEntry.Key, kbuf)
	if err != nil {
		return nil, nil, err
	}

	v, err = proto.MarshalOptions{Deterministic: true}.Marshal(pkEntry.Value)
	if err != nil {
		return nil, nil, err
	}

	return kbuf.Bytes(), v, nil
}
