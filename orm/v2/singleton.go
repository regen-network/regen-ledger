package v2

import (
	"github.com/cosmos/cosmos-sdk/store/types"
	"google.golang.org/protobuf/proto"
)

type singletonStore struct {
	prefix []byte
}

func (s *singletonStore) isStore() {}

func (s *singletonStore) Create(kv types.KVStore, message proto.Message) error {
	_, err := s.Save(kv, message)
	return err
}

func (s *singletonStore) Has(kv types.KVStore, _ proto.Message) bool {
	return kv.Has(s.prefix)
}

func (s *singletonStore) Read(kv types.KVStore, message proto.Message) (found bool, err error) {
	bz := kv.Get(s.prefix)
	if bz == nil {
		return false, nil
	}

	err = proto.Unmarshal(bz, message)
	return true, err
}

func (s *singletonStore) Save(kv types.KVStore, message proto.Message) (created bool, err error) {
	created = kv.Has(s.prefix)
	bz, err := proto.Marshal(message)
	kv.Set(s.prefix, bz)
	return created, nil
}

func (s *singletonStore) Delete(kv types.KVStore, _ proto.Message) error {
	kv.Delete(s.prefix)
	return nil
}

func (s *singletonStore) List(types.KVStore, proto.Message, ...ListOption) Iterator {
	return &singletonIterator{store: s}
}

type singletonIterator struct {
	store *singletonStore
	kv    types.KVStore
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
