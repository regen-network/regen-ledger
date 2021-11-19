package singleton

import (
	"github.com/regen-network/regen-ledger/orm/v2/internal/key"

	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"github.com/regen-network/regen-ledger/orm/v2/types"
	"google.golang.org/protobuf/proto"
)

func BuildStore(nsPrefix []byte, descriptor *types.SingletonDescriptor) (store.Store, error) {
	prefix := key.MakePrefix(nsPrefix, descriptor.Id)
	s := &Store{prefix: prefix}
	return s, nil
}

type Store struct {
	prefix []byte
}

func (s *Store) isStore() {}

func (s *Store) Create(kv store.KVStore, message proto.Message) error {
	return s.Save(kv, message)
}

func (s *Store) Has(kv store.KVStore, _ proto.Message) bool {
	return kv.Has(s.prefix)
}

func (s *Store) Read(kv store.KVStore, message proto.Message) (found bool, err error) {
	bz := kv.Get(s.prefix)
	if bz == nil {
		return false, nil
	}

	err = proto.Unmarshal(bz, message)
	return true, err
}

func (s *Store) Save(kv store.KVStore, message proto.Message) error {
	bz, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	kv.Set(s.prefix, bz)
	return nil
}

func (s *Store) Delete(kv store.KVStore, _ proto.Message) error {
	kv.Delete(s.prefix)
	return nil
}

func (s *Store) List(store.KVStore, proto.Message, *list.Options) list.Iterator {
	return &singletonIterator{store: s}
}

type singletonIterator struct {
	store *Store
	kv    store.KVStore
	done  bool
}

func (s singletonIterator) isIterator() {}

func (s singletonIterator) Next(message proto.Message) (bool, error) {
	if s.done {
		return false, nil
	}

	s.done = true
	return s.store.Read(s.kv, message)
}
