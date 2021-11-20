package table

import (
	"bytes"
	"encoding/binary"
	"fmt"

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
	IndexerMap          map[string]*Indexer
	SeqPrefix           []byte
}

func (s Store) isStore() {}

func (s Store) primaryKey(message protoreflect.Message) ([]protoreflect.Value, []byte, error) {
	return s.PkCodec.EncodeFromMessage(message)
}

func (s Store) primaryKeyValues(message protoreflect.Message) []protoreflect.Value {
	return s.PkCodec.GetValues(message)
}

func (s Store) Create(kv store.KVStore, message proto.Message) error {
	_, err := s.save(kv, message, true)
	return err
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

func (s Store) Save(kv store.KVStore, message proto.Message) error {
	_, err := s.save(kv, message, false)
	return err
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

func (s Store) nextSeqValue(kv store.KVStore) (uint64, error) {
	bz := kv.Get(s.SeqPrefix)
	seq := uint64(1)
	if bz != nil {
		x, err := binary.ReadUvarint(bytes.NewReader(bz))
		if err != nil {
			return 0, err
		}
		seq = x + 1
	}
	bz = make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(bz, seq)
	kv.Set(s.SeqPrefix, bz[:n])
	return seq, nil
}

func (s Store) save(kv store.KVStore, message proto.Message, create bool) (bool, error) {
	mref := message.ProtoReflect()

	// handle auto-incrementing primary keys
	if create && s.SeqPrefix != nil {
		pkField := s.PkCodec.Fields[0]
		id := mref.Get(pkField).Uint()
		if id != 0 {
			return false, fmt.Errorf("trying generate an auto-incremented primary key, but the key is already set")
		}

		var err error
		id, err = s.nextSeqValue(kv)
		if err != nil {
			return false, err
		}

		mref.Set(pkField, protoreflect.ValueOfUint64(id))
	}

	pkValues, pk, err := s.primaryKey(mref)
	if err != nil {
		return false, err
	}

	bz := kv.Get(pk)
	var existing proto.Message
	if bz != nil {
		if create {
			return true, fmt.Errorf("object of type %T with primary key %s already exists, can't create", message, pkValues)
		}

		existing = mref.New().Interface()
		err = proto.Unmarshal(bz, existing)
		if err != nil {
			return true, err
		}
	}

	// temporarily clear primary key
	s.PkCodec.ClearKey(mref)

	// store object
	bz, err = proto.MarshalOptions{Deterministic: true}.Marshal(message)
	kv.Set(pk, bz)

	// set primary key again
	s.PkCodec.SetValues(mref, pkValues)

	created := existing == nil

	// set indexes
	if !created {
		existingRef := existing.ProtoReflect()
		for _, idx := range s.Indexers {
			if existing == nil {
				err = idx.onCreate(kv, mref)
			} else {
				err = idx.onUpdate(kv, mref, existingRef)
			}
			if err != nil {
				return created, err
			}
		}
	}

	return created, nil
}
