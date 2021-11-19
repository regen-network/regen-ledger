package v2

import (
	"fmt"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/internal/singleton"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"github.com/regen-network/regen-ledger/orm/v2/internal/table"
	"github.com/regen-network/regen-ledger/orm/v2/types"
)

type Schema struct {
	stores    map[protoreflect.FullName]store.Store
	idMap     map[uint32]bool
	fileDescs map[uint32]protoreflect.FileDescriptor
}

func (s Schema) getStoreForMessage(message proto.Message) (store.Store, error) {
	name := message.ProtoReflect().Descriptor().FullName()
	if ts, ok := s.stores[name]; ok {
		return ts, nil
	}

	return nil, fmt.Errorf("can't find table store for message %s", name)
}

func (s Schema) buildStore(nsPrefix []byte, desc protoreflect.MessageDescriptor) error {
	msgName := desc.FullName()
	if _, ok := s.stores[msgName]; ok {
		return fmt.Errorf("store already registered for %s", msgName)
	}

	tableDesc := proto.GetExtension(desc.Options(), types.E_Table).(*types.TableDescriptor)
	singDesc := proto.GetExtension(desc.Options(), types.E_Singleton).(*types.SingletonDescriptor)

	if tableDesc != nil {
		if singDesc != nil {
			return fmt.Errorf("message %s cannot be declared as both a table and a singleton", msgName)
		}

		tableId := tableDesc.Id
		if s.idMap[tableId] {
			return fmt.Errorf("duplicate ID %d in ORM", tableId)
		}

		st, err := table.BuildStore(nsPrefix, tableDesc, desc)
		if err != nil {
			return err
		}

		s.stores[msgName] = st
		s.idMap[tableId] = true
		return nil
	} else if singDesc != nil {
		id := singDesc.Id
		if s.idMap[id] {
			return fmt.Errorf("duplicate ID %d in ORM", id)
		}

		st, err := singleton.BuildStore(nsPrefix, singDesc)
		if err != nil {
			return err
		}

		s.stores[msgName] = st
		s.idMap[id] = true
		return nil
	} else {
		return nil
	}
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
		stores:    map[protoreflect.FullName]store.Store{},
		idMap:     map[uint32]bool{},
		fileDescs: map[uint32]protoreflect.FileDescriptor{},
	}

	for _, opt := range opts {
		err := opt.applySchemaOption(sch)
		if err != nil {
			return nil, err
		}
	}

	return sch, nil
}

func FileDescriptor(prefix uint32, descriptor protoreflect.FileDescriptor) SchemaOption {
	return schemaOpt(func(schema *Schema) error {
		if fd, ok := schema.fileDescs[prefix]; ok {
			return fmt.Errorf("file descriptor for package %d already "+
				"registered with prefix %s, trying to register package %s", prefix, fd.Package(), descriptor.Package())
		}

		schema.fileDescs[prefix] = descriptor

		prefix := key.MakeUint32Prefix(nil, prefix)

		msgs := descriptor.Messages()
		n := msgs.Len()
		for i := 0; i < n; i++ {
			err := schema.buildStore(prefix, msgs.Get(i))
			if err != nil {
				return err
			}
		}
		return nil
	})
}
