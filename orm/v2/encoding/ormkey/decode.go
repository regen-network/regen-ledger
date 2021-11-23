package ormkey

import "github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"

type KVCodec interface {
	CodecI
	DecodeKV(k, v []byte) (ormdecode.Entry, error)
	EncodeKV(entry ormdecode.Entry) (k, v []byte, err error)
}
