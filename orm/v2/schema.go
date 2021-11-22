package orm

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/regen-network/regen-ledger/orm/v2/types/kvlayout"

	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/internal/singleton"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"github.com/regen-network/regen-ledger/orm/v2/internal/table"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"
)

type Schema struct {
	prefix     []byte
	stores     map[protoreflect.FullName]store.Store
	storesById map[uint32]map[uint32]store.Store
	fileDescs  map[uint32]protoreflect.FileDescriptor
}

func (s Schema) getStoreForMessage(message proto.Message) (store.Store, error) {
	name := message.ProtoReflect().Descriptor().FullName()
	if ts, ok := s.stores[name]; ok {
		return ts, nil
	}

	return nil, fmt.Errorf("can't find table store for message %s", name)
}

func (s Schema) buildStore(nsPrefix []byte, fdId uint32, desc protoreflect.MessageDescriptor) error {
	msgName := desc.FullName()
	if _, ok := s.stores[msgName]; ok {
		return fmt.Errorf("store already registered for %s", msgName)
	}

	tableDesc := proto.GetExtension(desc.Options(), ormpb.E_Table).(*ormpb.TableDescriptor)
	singDesc := proto.GetExtension(desc.Options(), ormpb.E_Singleton).(*ormpb.SingletonDescriptor)

	var id uint32
	var st store.Store
	var err error
	if tableDesc != nil {
		if singDesc != nil {
			return fmt.Errorf("message %s cannot be declared as both a table and a singleton", msgName)
		}

		id = tableDesc.Id
		st, err = table.BuildStore(nsPrefix, tableDesc, getMessageType(desc))
		if err != nil {
			return err
		}

	} else if singDesc != nil {
		id = singDesc.Id
		st, err = singleton.BuildStore(nsPrefix, singDesc, getMessageType(desc))
		if err != nil {
			return err
		}
	}

	if st != nil {
		if _, ok := s.storesById[fdId][id]; ok {
			return fmt.Errorf("duplicate ID %d in file", id)
		}

		s.stores[msgName] = st
		s.storesById[fdId][id] = st
	}

	return nil
}

func getMessageType(descriptor protoreflect.MessageDescriptor) protoreflect.MessageType {
	typ, err := protoregistry.GlobalTypes.FindMessageByName(descriptor.FullName())
	if err != nil {
		return dynamicpb.NewMessageType(descriptor)
	}
	return typ
}

func gatherListOptions(opts []ListOption) *list.Options {
	res := &list.Options{}
	for _, opt := range opts {
		opt.applyListOption(res)
	}
	return res
}

type SchemaOption interface {
	applySchemaOption(*Schema) error
}

type schemaOpt func(*Schema) error

func (s schemaOpt) applySchemaOption(sch *Schema) error {
	return s(sch)
}

func BuildSchema(opts ...SchemaOption) (*Schema, error) {
	sch := &Schema{
		stores:     map[protoreflect.FullName]store.Store{},
		storesById: map[uint32]map[uint32]store.Store{},
		fileDescs:  map[uint32]protoreflect.FileDescriptor{},
	}

	for _, opt := range opts {
		err := opt.applySchemaOption(sch)
		if err != nil {
			return nil, err
		}
	}

	return sch, nil
}

func FileDescriptor(id uint32, descriptor protoreflect.FileDescriptor) SchemaOption {
	return schemaOpt(func(schema *Schema) error {
		if fd, ok := schema.fileDescs[id]; ok {
			return fmt.Errorf("file descriptor for package %d already "+
				"registered with prefix %s, trying to register package %s", id, fd.Package(), descriptor.Package())
		}

		schema.fileDescs[id] = descriptor
		schema.storesById[id] = map[uint32]store.Store{}

		prefix := key.MakeUint32Prefix(schema.prefix, id)
		msgs := descriptor.Messages()
		n := msgs.Len()
		for i := 0; i < n; i++ {
			err := schema.buildStore(prefix, id, msgs.Get(i))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func Prefix(prefix []byte) SchemaOption {
	return schemaOpt(func(schema *Schema) error {
		schema.prefix = prefix
		return nil
	})
}

func (s Schema) Decode(k, v []byte) (kvlayout.Entry, error) {
	r := bytes.NewReader(k)
	// we assume the prefix has been checked by the caller for performance
	err := key.SkipPrefix(r, s.prefix)
	if err != nil {
		return nil, err
	}

	fdId, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, err
	}

	tableId, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, err
	}

	fdStores, ok := s.storesById[uint32(fdId)]
	if !ok {
		return nil, fmt.Errorf("can't resolve file descriptor ID %d", fdId)
	}

	st, ok := fdStores[uint32(tableId)]
	if !ok {
		return nil, fmt.Errorf("can't resolve table or singleton with ID %d", fdId)
	}

	return st.Decode(k, v)
}
