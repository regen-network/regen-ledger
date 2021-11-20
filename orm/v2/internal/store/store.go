package store

import (
	"google.golang.org/protobuf/proto"

	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
)

type Store interface {
	Create(kv KVStore, message proto.Message) error
	Has(kv KVStore, message proto.Message) bool
	Read(kv KVStore, message proto.Message) (found bool, err error)
	Save(kv KVStore, message proto.Message) error
	Delete(kv KVStore, message proto.Message) error
	List(kv KVStore, options *list.Options) list.Iterator
}
