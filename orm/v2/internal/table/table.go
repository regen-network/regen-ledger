package table

import (
	"bytes"
	"fmt"

	"github.com/regen-network/regen-ledger/orm/v2/internal/dynamicmsg"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
)

type Store struct {
	NumPrimaryKeyFields int
	Prefix              []byte
	PkPrefix            []byte
	PkCodec             *key.Codec
	Indexers            []*Indexer
	IndexersByFields    map[string]*Indexer
	IndexersById        map[uint32]*Indexer
	Descriptor          protoreflect.MessageDescriptor
}

func (s Store) isStore() {}

func (s Store) primaryKey(message protoreflect.Message) ([]protoreflect.Value, []byte, error) {
	return s.PkCodec.EncodeFromMessage(message)
}

func (s Store) primaryKeyValues(message protoreflect.Message) []protoreflect.Value {
	return s.PkCodec.GetValues(message)
}

func (s Store) Read(kv store.KVStore, message proto.Message) (bool, error) {
	refm := message.ProtoReflect()
	pkValues, pk, err := s.primaryKey(refm)
	if err != nil {
		return false, err
	}

	bz := kv.Get(pk)
	if bz == nil {
		return false, nil
	}

	err = proto.Unmarshal(bz, message)
	if err != nil {
		return true, err
	}

	// rehydrate primary key
	s.PkCodec.SetValues(refm, pkValues)

	return true, nil
}

func (s Store) Has(kv store.KVStore, message proto.Message) bool {
	_, pk, err := s.primaryKey(message.ProtoReflect())
	if err != nil {
		return false
	}

	return kv.Has(pk)
}

func (s Store) Save(kv store.KVStore, message proto.Message, mode store.SaveMode) error {
	mref := message.ProtoReflect()

	pkValues, pk, err := s.primaryKey(mref)
	if err != nil {
		return err
	}

	bz := kv.Get(pk)
	var existing proto.Message
	if bz != nil {
		if mode == store.SAVE_MODE_CREATE {
			return fmt.Errorf("object of type %T with primary key %s already exists, can't create", message, pkValues)
		}

		existing = mref.New().Interface()
		err = proto.Unmarshal(bz, existing)
		if err != nil {
			return err
		}
	} else {
		if mode == store.SAVE_MODE_UPDATE {
			return fmt.Errorf("object of type %T with primary key %s wasn't saved before, can't update", message, pkValues)
		}
	}

	// temporarily clear primary key
	s.PkCodec.ClearKey(mref)

	// store object
	bz, err = proto.MarshalOptions{Deterministic: true}.Marshal(message)
	kv.Set(pk, bz)

	// set primary key again
	s.PkCodec.SetValues(mref, pkValues)

	// set indexes
	if existing != nil {
		existingRef := existing.ProtoReflect()
		for _, idx := range s.Indexers {
			if existing == nil {
				err = idx.onCreate(kv, mref)
			} else {
				err = idx.onUpdate(kv, mref, existingRef)
			}
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s Store) Delete(kv store.KVStore, message proto.Message) error {
	mref := message.ProtoReflect()
	_, pk, err := s.primaryKey(mref)
	if err != nil {
		return err
	}

	// clear indexes
	for _, idx := range s.Indexers {
		err := idx.onCreate(kv, mref)
		if err != nil {
			return err
		}
	}

	// delete object
	kv.Delete(pk)

	return nil
}

func (s Store) Decode(k []byte, v []byte) (proto.Message, error) {
	if bytes.HasPrefix(k, s.PkPrefix) {
		pkValues, err := s.PkCodec.Decode(bytes.NewReader(k))
		if err != nil {
			return nil, err
		}

		msg, err := dynamicmsg.Unmarshal(s.Descriptor, v)
		if err != nil {
			return nil, err
		}

		// rehydrate pk
		s.PkCodec.SetValues(msg.ProtoReflect(), pkValues)

		return msg, nil
	}
	return nil, nil
}
