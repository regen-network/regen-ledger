package table

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"
)

const (
	SchemaSpacePrefix   = 0
	SequenceSpacePrefix = 2
	PrimaryKeyPrefix    = 0
)

func BuildStore(nsPrefix []byte, tableDesc *ormpb.TableDescriptor, messageType protoreflect.MessageType) (store.Store, error) {
	desc := messageType.Descriptor()
	tableId := tableDesc.Id
	if tableId == 0 {
		return nil, ormerrors.InvalidTableId.Wrapf("table %s", messageType.Descriptor().FullName())
	}

	if tableDesc.PrimaryKey == nil {
		return nil, ormerrors.MissingPrimaryKey.Wrap(string(messageType.Descriptor().FullName()))
	}

	pkFields, err := key.GetFieldDescriptors(desc, tableDesc.PrimaryKey.Fields)
	if err != nil {
		return nil, err
	}

	var seqPrefix []byte
	if tableDesc.PrimaryKey.AutoIncrement {
		if len(pkFields) != 1 && pkFields[0].Kind() != protoreflect.Uint64Kind {
			return nil, ormerrors.InvalidAutoIncrementKey.Wrapf("got %s for %s", tableDesc.PrimaryKey.Fields, desc.FullName())
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
		MsgType:             messageType,
	}

	idxIds := map[uint32]bool{}
	for _, idxDesc := range tableDesc.Index {
		id := idxDesc.Id
		if id == 0 {
			return nil, ormerrors.InvalidIndexId.Wrapf("index on table %s with fields %s", desc.FullName(), idxDesc.Fields)
		}

		if idxIds[id] {
			return nil, ormerrors.DuplicateIndexId.Wrapf("id %d on table %s", id, desc.FullName())
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
			FieldNames:  idxDesc.Fields,
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
