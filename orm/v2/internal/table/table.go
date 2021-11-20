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
	PkFields            []protoreflect.FieldDescriptor
	Prefix              []byte
	PkPrefix            []byte
	PkCodec             *key.Codec
	Indexers            []*Indexer
	IndexerMap          map[string]*Indexer
	SeqPrefix           []byte
}

func (s Store) isStore() {}

func (s Store) primaryKey(message proto.Message) ([]protoreflect.Value, []byte, error) {
	pkValues := s.primaryKeyValues(message)

	pkBuf := &bytes.Buffer{}
	pkBuf.Write(s.PkPrefix)
	err := s.PkCodec.Encode(pkValues, pkBuf, false)
	if err != nil {
		return nil, nil, err
	}

	return pkValues, pkBuf.Bytes(), nil
}

func (s Store) primaryKeyValues(message proto.Message) []protoreflect.Value {
	refm := message.ProtoReflect()
	// encode primary key
	pkValues := make([]protoreflect.Value, s.NumPrimaryKeyFields)
	for i, f := range s.PkFields {
		pkValues[i] = refm.Get(f)
	}

	return pkValues
}

func (s Store) Create(kv store.KVStore, message proto.Message) error {
	_, err := s.save(kv, message, true)
	return err
}

func (s Store) Read(kv store.KVStore, message proto.Message) (bool, error) {
	pkValues, pk, err := s.primaryKey(message)
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

	refm := message.ProtoReflect()

	// rehydrate primary key
	for i, f := range s.PkFields {
		refm.Set(f, pkValues[i])
	}

	return true, nil
}

func (s Store) Has(kv store.KVStore, message proto.Message) bool {
	_, pk, err := s.primaryKey(message)
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
	_, pk, err := s.primaryKey(message)
	if err != nil {
		return err
	}

	// clear indexes
	for _, idx := range s.Indexers {
		err := idx.onCreate(kv, message.ProtoReflect())
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
	refm := message.ProtoReflect()

	// handle auto-incrementing primary keys
	if create && s.SeqPrefix != nil {
		id := refm.Get(s.PkFields[0]).Uint()
		if id != 0 {
			return false, fmt.Errorf("trying generate an auto-incremented primary key, but the key is already set")
		}

		var err error
		id, err = s.nextSeqValue(kv)
		if err != nil {
			return false, err
		}

		refm.Set(s.PkFields[0], protoreflect.ValueOfUint64(id))
	}

	pkValues, pk, err := s.primaryKey(message)
	if err != nil {
		return false, err
	}

	bz := kv.Get(pk)
	var existing proto.Message
	if bz != nil {
		if create {
			return true, fmt.Errorf("object of type %T with primary key %s already exists, can't create", message, pkValues)
		}

		existing = refm.New().Interface()
		err = proto.Unmarshal(bz, existing)
		if err != nil {
			return true, err
		}
	}

	// temporarily clear primary key
	for _, f := range s.PkFields {
		refm.Clear(f)
	}

	// store object
	bz, err = proto.Marshal(message)
	kv.Set(pk, bz)

	// set primary key again
	for i, f := range s.PkFields {
		refm.Set(f, pkValues[i])
	}

	created := existing == nil

	// set indexes
	if !created {
		existingRef := existing.ProtoReflect()
		for _, idx := range s.Indexers {
			if existing == nil {
				err = idx.onCreate(kv, refm)
			} else {
				err = idx.onUpdate(kv, refm, existingRef)
			}
			if err != nil {
				return created, err
			}
		}
	}

	return created, nil
}
