package singleton

import (
	"encoding/json"
	io "io"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"
	"github.com/regen-network/regen-ledger/orm/v2/model/ormtable"
	"github.com/regen-network/regen-ledger/orm/v2/orm"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"
)

func BuildStore(nsPrefix []byte, singletonDescriptor *ormpb.SingletonDescriptor, messageType protoreflect.MessageType) (ormtable.Model, error) {
	id := singletonDescriptor.Id
	if id == 0 {
		return nil, ormerrors.InvalidTableId.Wrapf("singleton %s", messageType.Descriptor().FullName())
	}

	prefix := ormkey.MakeUint32Prefix(nsPrefix, id)
	key, err := ormkey.NewSingletonKey(prefix, messageType)
	if err != nil {
		return nil, err
	}
	s := &Store{key}
	return s, nil
}

type Store struct {
	*ormkey.SingletonKey
}

func (s *Store) PrimaryKey() ormkey.KVCodec {
	return s.SingletonKey
}

func (s *Store) isStore() {}

func (s *Store) Has(kv kv.ReadKVStore, key []protoreflect.Value, opts *ormtable.GetOptions) (bool, error) {
	return kv.Has(s.Prefix()), nil
}

func (s *Store) Get(kv kv.ReadKVStore, key []protoreflect.Value, message proto.Message, opts *ormtable.GetOptions) (found bool, err error) {
	bz := kv.Get(s.Prefix())
	if bz == nil {
		return false, nil
	}

	err = proto.Unmarshal(bz, message)
	return true, err
}

func (s *Store) Save(kv kv.KVStore, message proto.Message, _ ormtable.SaveMode) error {
	bz, err := proto.MarshalOptions{Deterministic: true}.Marshal(message)
	if err != nil {
		return err
	}
	kv.Set(s.Prefix(), bz)
	return nil
}

func (s *Store) Delete(kv kv.KVStore, key []protoreflect.Value) error {
	kv.Delete(s.Prefix())
	return nil
}

func (s *Store) List(kv kv.ReadKVStore, options *orm.ListOptions) orm.Iterator {
	return &singletonIterator{store: s, kv: kv}
}

func (s *Store) Decode(k []byte, v []byte) (ormdecode.Entry, error) {
	return s.DecodeKV(k, v)
}

func (s *Store) DefaultJSON() json.RawMessage {
	msg := s.MsgType.New().Interface()
	bz, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return json.RawMessage("{}")
	}
	return bz
}

func (s *Store) ValidateJSON(reader io.Reader) error {
	panic("implement me")
}

func (s *Store) ImportJSON(kvStore kv.KVStore, reader io.Reader) error {
	bz, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	msg := s.MsgType.New().Interface()
	err = protojson.Unmarshal(bz, msg)
	if err != nil {
		return err
	}

	return s.Save(kvStore, msg, ormtable.SAVE_MODE_DEFAULT)
}

func (s *Store) ExportJSON(kvStore kv.ReadKVStore, writer io.Writer) error {
	msg := s.MsgType.New().Interface()
	found, err := s.Get(kvStore, nil, msg, nil)
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
	orm.UnimplementedIterator

	store *Store
	kv    kv.ReadKVStore
	done  bool
}

func (s *singletonIterator) Next(message proto.Message) (bool, error) {
	if s.done {
		return false, nil
	}

	s.done = true
	return s.store.Get(s.kv, nil, message, nil)
}

func (s *singletonIterator) Close() {}
