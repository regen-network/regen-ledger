package table

import (
	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"
	"github.com/regen-network/regen-ledger/orm/v2/model/ormtable"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
)

type TableModel struct {
	MsgType         protoreflect.MessageType
	PkCodec         *ormkey.PrimaryKeyCodec
	Indexes         []*Index
	IndexesByFields map[string]*Index
	IndexesById     map[uint32]*Index

	Prefix   []byte
	PkPrefix []byte
}

func (s TableModel) primaryKey(message protoreflect.Message) ([]protoreflect.Value, []byte, error) {
	return s.PkCodec.EncodeFromMessage(message)
}

func (s TableModel) primaryKeyValues(message protoreflect.Message) []protoreflect.Value {
	return s.PkCodec.GetValues(message)
}

func (s TableModel) Get2(kv kv.ReadKVStore, primaryKey []protoreflect.Value, message proto.Message) (bool, error) {
	pk, err := s.PkCodec.Encode(primaryKey)
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
	s.PkCodec.SetValues(message.ProtoReflect(), primaryKey)

	return true, nil
}

func (s TableModel) Get(kv kv.ReadKVStore, message proto.Message) (bool, error) {
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

func (s TableModel) Has(kv kv.ReadKVStore, message proto.Message) bool {
	_, pk, err := s.primaryKey(message.ProtoReflect())
	if err != nil {
		return false
	}

	return kv.Has(pk)
}

func (s TableModel) Has2(kv kv.ReadKVStore, primaryKey []protoreflect.Value) (bool, error) {
	pk, err := s.PkCodec.Encode(primaryKey)
	if err != nil {
		return false, err
	}

	return kv.Has(pk), nil
}

func (s TableModel) Save(kv kv.KVStore, message proto.Message, mode ormtable.SaveMode) error {
	mref := message.ProtoReflect()

	pkValues, pk, err := s.primaryKey(mref)
	if err != nil {
		return err
	}

	bz := kv.Get(pk)
	var existing proto.Message
	if bz != nil {
		if mode == ormtable.SAVE_MODE_CREATE {
			return ormerrors.PrimaryKeyConstraintViolation.Wrapf("%q", mref.Descriptor().FullName())
		}

		existing = mref.New().Interface()
		err = proto.Unmarshal(bz, existing)
		if err != nil {
			return err
		}
	} else {
		if mode == ormtable.SAVE_MODE_UPDATE {
			return ormerrors.NotFoundOnUpdate.Wrapf("%q", mref.Descriptor().FullName())
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
	for _, idx := range s.Indexes {
		if existing == nil {
			err = idx.onCreate(kv, mref)
		} else {
			err = idx.onUpdate(kv, mref, existing.ProtoReflect())
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TableModel) Delete(kv kv.KVStore, message proto.Message) error {
	mref := message.ProtoReflect()
	_, pk, err := s.primaryKey(mref)
	if err != nil {
		return err
	}

	// delete object
	kv.Delete(pk)

	// clear indexes
	for _, idx := range s.Indexes {
		err := idx.onDelete(kv, mref)
		if err != nil {
			return err
		}
	}

	return nil
}
