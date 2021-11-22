package ormtable

import (
	"encoding/json"
	"io"

	"github.com/regen-network/regen-ledger/orm/v2/orm"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"google.golang.org/protobuf/proto"
)

type Model interface {
	Has(kv kv.ReadKVStore, message proto.Message) bool
	Get(kv kv.ReadKVStore, message proto.Message) (found bool, err error)
	Save(kv kv.KVStore, message proto.Message, mode SaveMode) error
	Delete(kv kv.KVStore, message proto.Message) error
	List(kv kv.ReadKVStore, condition proto.Message, options *orm.ListOptions) orm.Iterator
	Decode(k []byte, v []byte) (ormdecode.Entry, error)
	DefaultJSON() json.RawMessage
	ValidateJSON(io.Reader) error
	ImportJSON(kv.KVStore, io.Reader) error
	ExportJSON(kv.ReadKVStore, io.Writer) error
}

type SaveMode int

const (
	SAVE_MODE_DEFAULT SaveMode = iota
	SAVE_MODE_CREATE
	SAVE_MODE_UPDATE
)
