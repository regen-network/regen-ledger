package v2

import (
	"encoding/binary"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	sdkstore "github.com/cosmos/cosmos-sdk/store/types"
)

type Store interface {
	Create(kv sdkstore.KVStore, message proto.Message) error
	Has(kv sdkstore.KVStore, message proto.Message) bool
	Read(kv sdkstore.KVStore, message proto.Message) (found bool, err error)
	Save(kv sdkstore.KVStore, message proto.Message) (created bool, err error)
	Delete(kv sdkstore.KVStore, message proto.Message) error
	List(kv sdkstore.KVStore, message proto.Message, options ...ListOption) Iterator

	isStore()
}

type orm struct {
	stores map[protoreflect.FullName]Store
	idMap  map[uint32]bool
}

func (s orm) Create(kv sdkstore.KVStore, message proto.Message) error {
	store, err := s.getStoreForMessage(message)
	if err != nil {
		return err
	}
	return store.Create(kv, message)
}

func (s orm) Has(kv sdkstore.KVStore, message proto.Message) bool {
	store, err := s.getStoreForMessage(message)
	if err != nil {
		return false
	}
	return store.Has(kv, message)
}

func (s orm) Read(kv sdkstore.KVStore, message proto.Message) (found bool, err error) {
	store, err := s.getStoreForMessage(message)
	if err != nil {
		return false, err
	}
	return store.Read(kv, message)
}

func (s orm) Save(kv sdkstore.KVStore, message proto.Message) (created bool, err error) {
	store, err := s.getStoreForMessage(message)
	if err != nil {
		return false, err
	}
	return store.Save(kv, message)
}

func (s orm) Delete(kv sdkstore.KVStore, message proto.Message) error {
	store, err := s.getStoreForMessage(message)
	if err != nil {
		return err
	}
	return store.Delete(kv, message)
}

func (s orm) List(kv sdkstore.KVStore, message proto.Message, options ...ListOption) Iterator {
	store, err := s.getStoreForMessage(message)
	if err != nil {
		return errIterator{err: err}
	}
	return store.List(kv, message, options...)
}

func (s orm) isStore() {}

var _ Store = &orm{}

func getFieldDescriptors(desc protoreflect.MessageDescriptor, fields string) []protoreflect.FieldDescriptor {
	fieldNames := strings.Split(fields, ",")
	var fieldDescs []protoreflect.FieldDescriptor
	for _, fname := range fieldNames {
		fieldDesc := desc.Fields().ByName(protoreflect.Name(strings.TrimSpace(fname)))
		fieldDescs = append(fieldDescs, fieldDesc)
	}
	return fieldDescs
}

func (s orm) getStoreForMessage(message proto.Message) (Store, error) {
	desc := message.ProtoReflect().Descriptor()
	if ts, ok := s.stores[desc.FullName()]; ok {
		return ts, nil
	}

	tableDesc := proto.GetExtension(desc.Options(), E_Table).(*TableDescriptor)
	singDesc := proto.GetExtension(desc.Options(), E_Singleton).(*SingletonDescriptor)

	if tableDesc != nil {
		if singDesc != nil {
			return nil, fmt.Errorf("message %s cannot be declared as both a table and a singleton", desc.FullName())
		}

		tableId := tableDesc.Id
		if s.idMap[tableId] {
			return nil, fmt.Errorf("duplicate ID %d in ORM", tableId)
		}

		pkFields := getFieldDescriptors(desc, tableDesc.PrimaryKey.Fields)
		pkCodec, err := makeKeyCodec(pkFields, true)
		if err != nil {
			return nil, err
		}

		numPrimaryKeyFields := len(pkFields)
		prefix := make([]byte, binary.MaxVarintLen32)
		n := binary.PutUvarint(prefix, uint64(tableId))
		prefix = prefix[:n]

		store := &tableStore{
			numPrimaryKeyFields: numPrimaryKeyFields,
			prefix:              prefix,
			pkFields:            pkFields,
			pkCodec:             pkCodec,
			indexerMap:          map[string]*indexer{},
		}

		for _, idxDesc := range tableDesc.Index {
			idxFields := getFieldDescriptors(desc, idxDesc.Fields)
			cdc, err := makeIndexKeyCodec(idxFields, pkFields)
			if err != nil {
				return nil, err
			}
			lenPrefix := len(prefix)
			idxPrefix := make([]byte, lenPrefix+1+binary.MaxVarintLen32)
			copy(idxPrefix, prefix)
			idxPrefix[lenPrefix] = 1 // indexes all prefixed with 1
			n = binary.PutUvarint(prefix[lenPrefix+1:], uint64(idxDesc.Id))
			idx := &indexer{
				indexFields: idxFields,
				prefix:      idxPrefix[:lenPrefix+1+n],
				codec:       cdc,
			}
			store.indexers = append(store.indexers, idx)
			store.indexerMap[idxDesc.Fields] = idx
		}

		s.stores[desc.FullName()] = store
		s.idMap[tableId] = true
		return store, nil
	} else if singDesc != nil {
		id := singDesc.Id
		if s.idMap[id] {
			return nil, fmt.Errorf("duplicate ID %d in ORM", id)
		}

		prefix := make([]byte, binary.MaxVarintLen32)
		n := binary.PutUvarint(prefix, uint64(singDesc.Id))
		prefix = prefix[:n]
		store := &singletonStore{prefix: prefix}
		s.stores[desc.FullName()] = store
		s.idMap[id] = true
		return store, nil
	} else {
		return nil, fmt.Errorf("proto message %s not configured with orm annotations", desc.FullName())
	}
}

type Iterator interface {
	isIterator()
	Next(proto.Message) (bool, error)
}

type ListOption interface {
	applyListOption(*listOptions)
}

func Reverse() ListOption {
	return listOption(func(options *listOptions) {
		options.reverse = true
	})
}

func IndexHint(fields string) ListOption {
	return listOption(func(options *listOptions) {
		options.indexHint = fields
	})
}

type listOption func(*listOptions)

func (listOption) applyListOption(*listOptions) {}

type listOptions struct {
	reverse   bool
	indexHint string
}

func gatherListOptions(opts []ListOption) *listOptions {
	res := &listOptions{}
	for _, opt := range opts {
		opt.applyListOption(res)
	}
	return res
}

type errIterator struct {
	err error
}

func (e errIterator) isIterator() {}

func (e errIterator) Next(proto.Message) (bool, error) { return false, e.err }
