package singleton

import (
	"encoding/json"
	io "io"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"
)

func BuildStore(nsPrefix []byte, singletonDescriptor *ormpb.SingletonDescriptor, messageType protoreflect.MessageType) (store.Store, error) {
	id := singletonDescriptor.Id
	if id == 0 {
		return nil, ormerrors.InvalidTableId.Wrapf("singleton %s", messageType.Descriptor().FullName())
	}

	prefix := ormkey.MakeUint32Prefix(nsPrefix, id)
	s := &Store{prefix: prefix, msgType: messageType}
	return s, nil
}

type Store struct {
	prefix  []byte
	msgType protoreflect.MessageType
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

func (s *Store) Decode(k []byte, v []byte) (ormdecode.Entry, error) {
	msg := s.msgType.New().Interface()
	err := proto.Unmarshal(v, msg)
	return ormdecode.PrimaryKeyEntry{Value: msg}, err
}

func (s *Store) DefaultJSON() json.RawMessage {
	msg := s.msgType.New().Interface()
	bz, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return json.RawMessage("{}")
	}
	return bz
}

func (s *Store) ValidateJSON(reader io.Reader) error {
	panic("implement me")
}

func (s *Store) ImportJSON(kvStore store.KVStore, reader io.Reader) error {
	bz, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	msg := s.msgType.New().Interface()
	err = protojson.Unmarshal(bz, msg)
	if err != nil {
		return err
	}

	return s.Save(kvStore, msg, store.SAVE_MODE_DEFAULT)
}

func (s *Store) ExportJSON(kvStore store.KVStore, writer io.Writer) error {
	msg := s.msgType.New().Interface()
	found, err := s.Read(kvStore, msg)
	if err != nil {
		return err
	}

	var bz []byte
	if !found {
		bz = s.DefaultJSON()
	} else {
		bz, err = protojson.Marshal(msg)
		if err != nil {
			return err
		}
	}

	_, err = writer.Write(bz)
	return err
}

type singletonIterator struct {
	store *Store
	kv    store.KVStore
	done  bool
}

func (s *singletonIterator) Next(message proto.Message) (bool, error) {
	if s.done {
		return false, nil
	}

	s.done = true
	return s.store.Read(s.kv, message)
}

func (s *singletonIterator) Close() {}
