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
	MsgType             protoreflect.MessageType
	PkCodec             *ormkey.PrimaryKeyCodec
	Indexes             []*Index
	IndexesByFields     map[string]*Index
	IndexesByFieldNames map[protoreflect.Name]*Index
	IndexesById         map[uint32]*Index

	Prefix   []byte
	PkPrefix []byte
}

func (s TableModel) PrimaryKey() ormkey.KVCodec {
	return s.PkCodec
}

func (s TableModel) primaryKey(message protoreflect.Message) ([]protoreflect.Value, []byte, error) {
	return s.PkCodec.EncodeFromMessage(message)
}

func (s TableModel) primaryKeyValues(message protoreflect.Message) []protoreflect.Value {
	return s.PkCodec.GetValues(message)
}

func (s TableModel) get(kv kv.ReadKVStore, key []protoreflect.Value, message proto.Message, opts *ormtable.GetOptions) (bool, []byte, error) {
	if opts != nil && opts.UseUniqueIndex != "" {
		return false, nil, ormerrors.UnsupportedOperation
	}

	pk, err := s.PkCodec.Encode(key)
	if err != nil {
		return false, nil, err
	}

	bz := kv.Get(pk)
	if bz == nil {
		return false, pk, nil
	}

	err = s.PkCodec.Unmarshal(key, bz, message)
	return true, pk, err
}

func (s TableModel) Get(kv kv.ReadKVStore, key []protoreflect.Value, message proto.Message, opts *ormtable.GetOptions) (bool, error) {
	found, _, err := s.get(kv, key, message, opts)
	return found, err
}

func (s TableModel) Has(kv kv.ReadKVStore, key []protoreflect.Value, opts *ormtable.GetOptions) (bool, error) {
	if opts != nil && opts.UseUniqueIndex != "" {
		return false, ormerrors.UnsupportedOperation
	}

	pk, err := s.PkCodec.Encode(key)
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

	existing := mref.New().Interface()
	haveExisting, _, err := s.get(kv, pkValues, existing, &ormtable.GetOptions{})
	if err != nil {
		return err
	}

	if haveExisting {
		if mode == ormtable.SAVE_MODE_CREATE {
			return ormerrors.PrimaryKeyConstraintViolation.Wrapf("%q", mref.Descriptor().FullName())
		}
	} else {
		if mode == ormtable.SAVE_MODE_UPDATE {
			return ormerrors.NotFoundOnUpdate.Wrapf("%q", mref.Descriptor().FullName())
		}
	}

	// temporarily clear primary key
	s.PkCodec.Clear(mref)

	// store object
	bz, err := proto.MarshalOptions{Deterministic: true}.Marshal(message)
	kv.Set(pk, bz)

	// set primary key again
	s.PkCodec.SetValues(mref, pkValues)

	// set indexes
	for _, idx := range s.Indexes {
		if !haveExisting {
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

func (s TableModel) Delete(kv kv.KVStore, key []protoreflect.Value) error {
	msg := s.MsgType.New().Interface()
	found, pk, err := s.get(kv, key, msg, &ormtable.GetOptions{})
	if err != nil {
		return err
	}

	if !found {
		return nil
	}

	// delete object
	kv.Delete(pk)

	// clear indexes
	mref := msg.ProtoReflect()
	for _, idx := range s.Indexes {
		err := idx.onDelete(kv, mref)
		if err != nil {
			return err
		}
	}

	return nil
}
