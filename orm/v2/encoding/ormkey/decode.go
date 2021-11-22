package ormkey

import "github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"

type KVCodec interface {
	DecodeKV(k, v []byte) (ormdecode.Entry, error)
}
