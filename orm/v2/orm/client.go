package orm

import (
	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"google.golang.org/protobuf/proto"
)

type Client interface {
	Create(message proto.Message) error
	Has(message proto.Message) bool
	Read(message proto.Message) (found bool, err error)
	Save(message proto.Message) error
	Delete(message proto.Message) error
	List(message proto.Message, options ...ListOption) list.Iterator
}

type ListOption interface {
	applyListOption(*list.Options)
}

func Reverse() ListOption {
	return listOption(func(options *list.Options) {
		options.Reverse = true
	})
}

func UseIndex(fields string) ListOption {
	return listOption(func(options *list.Options) {
		options.UseIndex = fields
	})
}

type ClientConn struct {
	Schema *Schema
}

func (c *ClientConn) Open(kvStore store.KVStore) Client {
	return &client{
		schema: c.Schema,
		kv:     kvStore,
	}
}

type listOption func(*list.Options)

func (listOption) applyListOption(*list.Options) {}

type client struct {
	schema *Schema
	kv     store.KVStore
}

func (s client) Create(message proto.Message) error {
	st, err := s.schema.getStoreForMessage(message)
	if err != nil {
		return err
	}
	return st.Save(s.kv, message, store.SAVE_MODE_CREATE)
}

func (s client) Has(message proto.Message) bool {
	st, err := s.schema.getStoreForMessage(message)
	if err != nil {
		return false
	}
	return st.Has(s.kv, message)
}

func (s client) Read(message proto.Message) (found bool, err error) {
	st, err := s.schema.getStoreForMessage(message)
	if err != nil {
		return false, err
	}
	return st.Read(s.kv, message)
}

func (s client) Save(message proto.Message) error {
	st, err := s.schema.getStoreForMessage(message)
	if err != nil {
		return err
	}
	return st.Save(s.kv, message, store.SAVE_MODE_DEFAULT)
}

func (s client) Delete(message proto.Message) error {
	st, err := s.schema.getStoreForMessage(message)
	if err != nil {
		return err
	}
	return st.Delete(s.kv, message)
}

func (s client) List(message proto.Message, options ...ListOption) list.Iterator {
	st, err := s.schema.getStoreForMessage(message)
	if err != nil {
		return list.ErrIterator{Err: err}
	}
	return st.List(s.kv, gatherListOptions(options))
}

var _ Client = &client{}
