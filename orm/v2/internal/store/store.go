package store

import (
	"google.golang.org/protobuf/proto"

	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
)

type Store interface {
	Has(kv KVStore, message proto.Message) bool
	Read(kv KVStore, message proto.Message) (found bool, err error)
	Save(kv KVStore, message proto.Message, mode SaveMode) error
	Delete(kv KVStore, message proto.Message) error
	List(kv KVStore, options *list.Options) list.Iterator
}

type SaveMode int

const (
	SAVE_MODE_DEFAULT SaveMode = iota
	SAVE_MODE_CREATE
	SAVE_MODE_UPDATE
)
