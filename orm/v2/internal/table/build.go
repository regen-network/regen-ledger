package table

import (
	"fmt"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"github.com/regen-network/regen-ledger/orm/v2/ormpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	SchemaSpacePrefix   = 0
	SequenceSpacePrefix = 2
	PrimaryKeyPrefix    = 0
)

func BuildStore(nsPrefix []byte, tableDesc *ormpb.TableDescriptor, desc protoreflect.MessageDescriptor) (store.Store, error) {
	tableId := tableDesc.Id
	if tableId == 0 {
		return nil, fmt.Errorf("0 is not a valid id for table %s", desc.FullName())
	}

	if tableDesc.Id == 0 {
		return nil, fmt.Errorf("table id must be non-zero")
	}

	if tableDesc.PrimaryKey == nil {
		return nil, fmt.Errorf("no primary key defined")
	}

	pkFields, err := key.GetFieldDescriptors(desc, tableDesc.PrimaryKey.Fields)
	if err != nil {
		return nil, err
	}

	var seqPrefix []byte
	if tableDesc.PrimaryKey.AutoIncrement {
		if len(pkFields) != 1 && pkFields[0].Kind() != protoreflect.Uint64Kind {
			return nil, fmt.Errorf("only a single uint64 field is supported for primary keys, got %s", pkFields)
		}

		seqPrefix = key.MakeUint32Prefix(nsPrefix, SchemaSpacePrefix)
		seqPrefix = key.MakeUint32Prefix(seqPrefix, SequenceSpacePrefix)
		seqPrefix = key.MakeUint32Prefix(seqPrefix, tableId)
	}

	prefix := key.MakeUint32Prefix(nsPrefix, tableDesc.Id)
	pkPrefix := key.MakeUint32Prefix(prefix, PrimaryKeyPrefix)

	pkCodec, err := key.MakeCodec(pkPrefix, pkFields)
	if err != nil {
		return nil, err
	}

	numPrimaryKeyFields := len(pkFields)

	st := &Store{
		NumPrimaryKeyFields: numPrimaryKeyFields,
		Prefix:              prefix,
		PkPrefix:            pkPrefix,
		PkCodec:             pkCodec,
		IndexersByFields:    map[string]*Indexer{},
		IndexersById:        map[uint32]*Indexer{},
		Descriptor:          desc,
	}

	idxIds := map[uint32]bool{}
	for _, idxDesc := range tableDesc.Index {
		id := idxDesc.Id
		if id == 0 {
			return nil, fmt.Errorf("0 is not a valid id for index on table %s with fields %s", desc.FullName(), idxDesc.Fields)
		}

		if idxIds[id] {
			return nil, fmt.Errorf("duplicate index on table %s with id %d", desc.FullName(), id)
		}

		idxIds[id] = true

		idxFields, err := key.GetFieldDescriptors(desc, idxDesc.Fields)
		if err != nil {
			return nil, err
		}

		idxPrefix := key.MakeUint32Prefix(prefix, id)
		cdc, err := key.MakeIndexKeyCodec(idxPrefix, idxFields, pkFields)
		if err != nil {
			return nil, err
		}
		idx := &Indexer{
			IndexFields: idxFields,
			Prefix:      idxPrefix,
			Codec:       cdc,
		}
		st.Indexers = append(st.Indexers, idx)
		st.IndexersByFields[idxDesc.Fields] = idx
		st.IndexersById[id] = idx
	}

	if len(seqPrefix) != 0 {
		return &AutoIncStore{Store: st, SeqPrefix: seqPrefix}, nil
	} else {
		return st, nil
	}
}
