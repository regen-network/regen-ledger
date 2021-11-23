package ormkv

import (
	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"
	"github.com/regen-network/regen-ledger/orm/v2/model/ormschema"
	"github.com/regen-network/regen-ledger/orm/v2/model/ormtable"
	"github.com/regen-network/regen-ledger/orm/v2/orm"
	"google.golang.org/protobuf/proto"
)

type store struct {
	orm.UnimplementedStore
	schema *ormschema.Schema
	kv     kv.KVStore
}

func NewStore(schema *ormschema.Schema, kv kv.KVStore) orm.Store {
	return &store{schema: schema, kv: kv}
}

func (s store) Create(messages ...proto.Message) error {
	for _, message := range messages {
		st, err := s.schema.GetStoreForMessage(message)
		if err != nil {
			return err
		}
		err = st.Save(s.kv, message, ormtable.SAVE_MODE_CREATE)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s store) Has(messages ...proto.Message) bool {
	for _, message := range messages {
		st, err := s.schema.GetStoreForMessage(message)
		if err != nil {
			return false
		}
		found, _ := st.Has(s.kv, nil, nil)
		if !found {
			return false
		}
	}
	return true
}

func (s store) Get(messages ...proto.Message) (found bool, err error) {
	for _, message := range messages {
		st, err := s.schema.GetStoreForMessage(message)
		if err != nil {
			return false, err
		}
		found, err = st.Get(s.kv, nil, message, nil)
		if !found || err != nil {
			return false, err
		}
	}
	return true, err
}

func (s store) Save(messages ...proto.Message) error {
	for _, msg := range messages {
		st, err := s.schema.GetStoreForMessage(msg)
		if err != nil {
			return err
		}
		err = st.Save(s.kv, msg, ormtable.SAVE_MODE_DEFAULT)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s store) Delete(messages ...proto.Message) error {
	for _, msg := range messages {
		st, err := s.schema.GetStoreForMessage(msg)
		if err != nil {
			return err
		}
		values := st.PrimaryKey().GetValues(msg.ProtoReflect())
		err = st.Delete(s.kv, values)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s store) List(condition proto.Message, options *orm.ListOptions) orm.Iterator {
	st, err := s.schema.GetStoreForMessage(condition)
	if err != nil {
		return orm.ErrIterator{Err: err}
	}

	return st.List(s.kv, options)
}

var _ orm.Store = &store{}
