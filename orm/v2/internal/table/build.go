package table

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"github.com/regen-network/regen-ledger/orm/v2/types"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	SchemaSpacePrefix   byte = 0
	SequenceSpacePrefix      = 2
	PrimaryKeyPrefix         = 0
)

func BuildStore(nsPrefix []byte, tableDesc *types.TableDescriptor, desc protoreflect.MessageDescriptor) (store.Store, error) {
	tableId := tableDesc.Id
	if tableId == 0 {
		return nil, fmt.Errorf("0 is not a valid id for table %s", desc.FullName())
	}

	pkFields := getFieldDescriptors(desc, tableDesc.PrimaryKey.Fields)
	var seqPrefix []byte
	if tableDesc.PrimaryKey.AutoIncrement {
		if len(pkFields) != 1 && pkFields[0].Kind() != protoreflect.Uint64Kind {
			return nil, fmt.Errorf("only a single uint64 field is supported for primary keys, got %s", pkFields)
		}

		buf := &bytes.Buffer{}
		buf.Write(nsPrefix)
		buf.WriteByte(SchemaSpacePrefix)
		buf.WriteByte(SequenceSpacePrefix)
		seqPrefix = key.MakeUint32Prefix(buf.Bytes(), tableId)
	}

	pkCodec, err := key.MakeCodec(pkFields, true)
	if err != nil {
		return nil, err
	}

	numPrimaryKeyFields := len(pkFields)

	prefix := key.MakeUint32Prefix(nsPrefix, tableDesc.Id)
	pkPrefix := make([]byte, len(prefix))
	copy(pkPrefix, prefix)
	pkPrefix = append(pkPrefix, PrimaryKeyPrefix) // primary key table always prefixed with 0

	st := &Store{
		NumPrimaryKeyFields: numPrimaryKeyFields,
		Prefix:              prefix,
		PkPrefix:            pkPrefix,
		PkFields:            pkFields,
		PkCodec:             pkCodec,
		IndexerMap:          map[string]*Indexer{},
		SeqPrefix:           seqPrefix,
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

		idxFields := getFieldDescriptors(desc, idxDesc.Fields)
		cdc, err := key.MakeIndexKeyCodec(idxFields, pkFields)
		if err != nil {
			return nil, err
		}
		lenPrefix := len(prefix)
		idxPrefix := make([]byte, lenPrefix+binary.MaxVarintLen32)
		copy(idxPrefix, prefix)
		n := binary.PutUvarint(idxPrefix[lenPrefix:], uint64(id))
		idx := &Indexer{
			IndexFields: idxFields,
			Prefix:      idxPrefix[:lenPrefix+n],
			Codec:       cdc,
		}
		st.Indexers = append(st.Indexers, idx)
		st.IndexerMap[idxDesc.Fields] = idx
	}

	return st, nil
}

func getFieldDescriptors(desc protoreflect.MessageDescriptor, fields string) []protoreflect.FieldDescriptor {
	fieldNames := strings.Split(fields, ",")
	var fieldDescs []protoreflect.FieldDescriptor
	for _, fname := range fieldNames {
		fieldDesc := desc.Fields().ByName(protoreflect.Name(strings.TrimSpace(fname)))
		fieldDescs = append(fieldDescs, fieldDesc)
	}
	return fieldDescs
}
