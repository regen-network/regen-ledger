package singleton

import (
	"fmt"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/ormpb"

	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"google.golang.org/protobuf/proto"
)

func BuildStore(nsPrefix []byte, descriptor *ormpb.SingletonDescriptor) (store.Store, error) {
	id := descriptor.Id
	if id == 0 {
		return nil, fmt.Errorf("singleton must have non-zero id")
	}

	prefix := key.MakeUint32Prefix(nsPrefix, id)
	s := &Store{prefix: prefix}
	return s, nil
}

type Store struct {
	prefix []byte
}

func (s *Store) isStore() {}

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

func (s *Store) Save(kv store.KVStore, message proto.Message, _ store.SaveMode) error {
	bz, err := proto.MarshalOptions{Deterministic: true}.Marshal(message)
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

func (s *Store) List(kv store.KVStore, _ *list.Options) list.Iterator {
	return &singletonIterator{store: s, kv: kv}
}

type singletonIterator struct {
	store *Store
	kv    store.KVStore
	done  bool
}

func (s *singletonIterator) isIterator() {}

func (s *singletonIterator) Next(message proto.Message) (bool, error) {
	if s.done {
		return false, nil
	}

	s.done = true
	return s.store.Read(s.kv, message)
}
