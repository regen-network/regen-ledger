package ormkey

import (
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type SingletonKey struct {
	*Codec
	MsgType protoreflect.MessageType
}

func (s SingletonKey) DecodeKV(k, v []byte) (ormdecode.Entry, error) {
	msg := s.MsgType.New().Interface()
	err := proto.Unmarshal(v, msg)
	return ormdecode.PrimaryKeyEntry{Value: msg}, err
}

func (s SingletonKey) EncodeKV(entry ormdecode.Entry) (k, v []byte, err error) {
	pEntry, ok := entry.(ormdecode.PrimaryKeyEntry)
	if !ok {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	if len(pEntry.Key) != 0 {
		return nil, nil, ormerrors.BadDecodeEntry.Wrap("singleton entry shouldn't have non-empty a key")
	}

	bz, err := proto.Marshal(pEntry.Value)
	if err != nil {
		return nil, nil, err
	}

	return s.Prefix(), bz, nil
}

func NewSingletonKey(prefix []byte, msgType protoreflect.MessageType) (*SingletonKey, error) {
	cdc, err := MakeCodec(prefix, nil)
	if err != nil {
		return nil, err
	}
	return &SingletonKey{Codec: cdc, MsgType: msgType}, nil
}
