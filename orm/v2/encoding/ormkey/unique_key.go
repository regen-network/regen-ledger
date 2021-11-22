package ormkey

import (
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type UniqueKeyCodec struct {
	*Codec
}

func NewUniqueKeyCodec(
	messageDescriptor protoreflect.MessageDescriptor,
	descriptor *ormpb.TableDescriptor,
	id int,
) {

}

func (cdc UniqueKeyCodec) DecodeKV(k, v []byte) (ormdecode.Entry, error) {
	panic("TODO")
}
