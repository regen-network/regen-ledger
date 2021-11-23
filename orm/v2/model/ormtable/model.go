package ormtable

import (
	"encoding/json"
	"io"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/orm"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"google.golang.org/protobuf/proto"
)

type Model interface {
	Has(kv kv.ReadKVStore, key []protoreflect.Value, opts *GetOptions) (bool, error)
	Get(kv kv.ReadKVStore, key []protoreflect.Value, message proto.Message, opts *GetOptions) (found bool, err error)
	Save(kv kv.KVStore, message proto.Message, mode SaveMode) error
	Delete(kv kv.KVStore, key []protoreflect.Value) error
	List(kv kv.ReadKVStore, options *orm.ListOptions) orm.Iterator

	PrimaryKey() ormkey.KVCodec
	Decode(k []byte, v []byte) (ormdecode.Entry, error)

	DefaultJSON() json.RawMessage
	ValidateJSON(io.Reader) error
	ImportJSON(kv.KVStore, io.Reader) error
	ExportJSON(kv.ReadKVStore, io.Writer) error
}

type GetOptions struct {
	UseUniqueIndex string
}

type ListOptions struct {
	Reverse bool
	// Index defines an index by field names. If it is empty the primary key
	// will be used as the index.
	Index string

	// Cursor specifies a cursor returned by Iterator.Cursor() to restart iteration from.
	Cursor orm.Cursor

	// Prefix defines an iteration prefix using values corresponding the the key
	// being used. Not all of the values in the key need to be specified and
	// they do not be sortable unlike start and end. Prefix or Start/End are
	// mutually exclusive and shouldn't be specified together.
	Prefix []protoreflect.Value

	// Start defines a start position using a set of values corresponding to the
	// index or primary key being used. Each of the values must match the type
	// of the key at that position and also be a sortable value. Not all of
	// the values in the key need to be specified.
	Start []protoreflect.Value

	// End defines an end position using a set of values correspond to the key
	// being used. Not all of the values in the key need to be specified.
	End []protoreflect.Value
}

type SaveMode int

const (
	SAVE_MODE_DEFAULT SaveMode = iota
	SAVE_MODE_CREATE
	SAVE_MODE_UPDATE
)
