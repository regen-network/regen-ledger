package table

import (
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"

	"github.com/regen-network/regen-ledger/orm/v2/model/ormtable"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
)

type AutoIncStore struct {
	*TableModel
	SeqCodec *ormkey.SeqCodec
}

func (s *AutoIncStore) Save(kv kv.KVStore, message proto.Message, mode ormtable.SaveMode) error {
	f := s.PkCodec.Fields[0]
	mref := message.ProtoReflect()
	val := mref.Get(f).Uint()
	if val == 0 {
		if mode == ormtable.SAVE_MODE_UPDATE {
			return ormerrors.PrimaryKeyInvalidOnUpdate
		}

		mode = ormtable.SAVE_MODE_CREATE
		key, err := s.nextSeqValue(kv)
		if err != nil {
			return err
		}

		mref.Set(f, protoreflect.ValueOfUint64(key))
	} else {
		if mode == ormtable.SAVE_MODE_CREATE {
			return ormerrors.AutoIncrementKeyAlreadySet
		}

		mode = ormtable.SAVE_MODE_UPDATE
	}
	return s.TableModel.Save(kv, message, mode)
}

func (s *AutoIncStore) nextSeqValue(kv kv.KVStore) (uint64, error) {
	bz := kv.Get(s.SeqCodec.Prefix)
	seq, err := s.SeqCodec.DecodeValue(bz)
	if err != nil {
		return 0, err
	}

	seq++
	kv.Set(s.SeqCodec.Prefix, s.SeqCodec.EncodeValue(seq))
	return seq, nil
}
