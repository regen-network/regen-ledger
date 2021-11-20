package table

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"google.golang.org/protobuf/proto"
)

type AutoIncStore struct {
	*Store
	SeqPrefix []byte
}

func (s *AutoIncStore) Save(kv store.KVStore, message proto.Message, mode store.SaveMode) error {
	f := s.PkCodec.Fields[0]
	mref := message.ProtoReflect()
	val := mref.Get(f).Uint()
	if val == 0 {
		if mode == store.SAVE_MODE_UPDATE {
			return fmt.Errorf("can't update when no primary key is set")
		}

		mode = store.SAVE_MODE_CREATE
		key, err := s.nextSeqValue(kv)
		if err != nil {
			return err
		}

		mref.Set(f, protoreflect.ValueOfUint64(key))
	} else {
		if mode == store.SAVE_MODE_CREATE {
			return fmt.Errorf("can't create with auto-increment primary key already set")
		}

		mode = store.SAVE_MODE_UPDATE
	}
	return s.Store.Save(kv, message, mode)
}

func (s *AutoIncStore) nextSeqValue(kv store.KVStore) (uint64, error) {
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
