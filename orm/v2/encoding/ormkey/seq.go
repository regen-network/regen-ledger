package ormkey

import (
	"bytes"
	"encoding/binary"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
)

type SeqCodec struct {
	TableName protoreflect.FullName
	Prefix    []byte
}

func (s SeqCodec) EncodeValue(seq uint64) (v []byte) {
	bz := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(bz, seq)
	return bz[:n]
}

func (s SeqCodec) DecodeKV(k, v []byte) (ormdecode.Entry, error) {
	if !bytes.Equal(k, s.Prefix) {
		return nil, ormerrors.UnexpectedDecodePrefix
	}

	x, err := s.DecodeValue(v)
	if err != nil {
		return nil, err
	}

	return ormdecode.SeqEntry{
		TableName: s.TableName,
		Value:     x,
	}, nil
}

func (s SeqCodec) EncodeKV(entry ormdecode.Entry) (k, v []byte, err error) {
	seqEntry, ok := entry.(ormdecode.SeqEntry)
	if !ok {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	if seqEntry.TableName != s.TableName {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	return s.Prefix, s.EncodeValue(seqEntry.Value), nil
}

func (s SeqCodec) DecodeValue(v []byte) (uint64, error) {
	if len(v) == 0 {
		return 0, nil
	}
	return binary.ReadUvarint(bytes.NewReader(v))
}
