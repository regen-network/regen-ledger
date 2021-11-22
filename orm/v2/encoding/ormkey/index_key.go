package ormkey

import (
	"bytes"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormpb"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormdecode"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type IndexKeyCodec struct {
	*Codec
	pkFieldOrder    []int
	tableName       protoreflect.FullName
	indexFieldNames []protoreflect.Name
	indexFields     []protoreflect.FieldDescriptor
}

func (cdc IndexKeyCodec) DecodeKV(k, _ []byte) (ormdecode.Entry, error) {
	values, err := cdc.Decode(bytes.NewReader(k))
	if err != nil {
		return nil, err
	}

	numPkFields := len(cdc.pkFieldOrder)
	pkValues := make([]protoreflect.Value, numPkFields)

	for i := 0; i < numPkFields; i++ {
		pkValues[i] = values[cdc.pkFieldOrder[i]]
	}

	numIndexFields := len(cdc.indexFields)
	return &ormdecode.IndexKeyEntry{
		TableName:       cdc.tableName,
		IndexFieldNames: cdc.indexFieldNames,
		FullKey:         values,
		IndexKey:        values[:numIndexFields],
		PrimaryKeyRest:  values[numIndexFields:],
		PrimaryKey:      pkValues,
	}, nil
}

func (i IndexKeyCodec) EncodeKV(entry ormdecode.Entry) (k, v []byte, err error) {
	indexEntry, ok := entry.(*ormdecode.IndexKeyEntry)
	if !ok {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	if indexEntry.TableName != i.tableName {
		return nil, nil, ormerrors.BadDecodeEntry
	}

	kbuf := &bytes.Buffer{}
	err = i.Codec.EncodeWriter(indexEntry.FullKey, kbuf)
	if err != nil {
		return nil, nil, err
	}

	return kbuf.Bytes(), sentinel, nil
}

var sentinel = []byte{0}

func (cdc IndexKeyCodec) ReadPrimaryKey(reader *bytes.Reader) ([]protoreflect.Value, error) {
	values, err := cdc.Decode(reader)
	if err != nil {
		return nil, err
	}

	return cdc.extractPrimaryKey(values), nil
}

func (cdc IndexKeyCodec) extractPrimaryKey(values []protoreflect.Value) []protoreflect.Value {
	numPkFields := len(cdc.pkFieldOrder)
	pkValues := make([]protoreflect.Value, numPkFields)

	for i := 0; i < numPkFields; i++ {
		pkValues[i] = values[cdc.pkFieldOrder[i]]
	}

	return pkValues
}

var _ KVCodec = &IndexKeyCodec{}

func MakeIndexKeyCodec(prefix []byte, indexFields []protoreflect.FieldDescriptor, primaryKeyFields []protoreflect.FieldDescriptor) (*IndexKeyCodec, error) {
	indexFieldMap := map[protoreflect.Name]int{}

	var keyFields []protoreflect.FieldDescriptor
	for i, f := range indexFields {
		indexFieldMap[f.Name()] = i
		keyFields = append(keyFields, f)
	}

	numIndexFields := len(indexFields)
	numPrimaryKeyFields := len(primaryKeyFields)
	pkFieldOrder := make([]int, numPrimaryKeyFields)
	k := 0
	for j, f := range primaryKeyFields {
		if i, ok := indexFieldMap[f.Name()]; ok {
			pkFieldOrder[j] = i
			continue
		}
		keyFields = append(keyFields, f)
		pkFieldOrder[j] = numIndexFields + k
		k++
	}

	cdc, err := MakeCodec(prefix, keyFields)
	if err != nil {
		return nil, err
	}
	return &IndexKeyCodec{
		Codec:        cdc,
		pkFieldOrder: pkFieldOrder,
		indexFields:  indexFields,
	}, nil
}

var _ KVCodec = &IndexKeyCodec{}

func NewIndexKeyCodec(
	prefix []byte,
	messageDescriptor protoreflect.MessageDescriptor,
	tableDescriptor *ormpb.TableDescriptor,
	indexId int,
) {

}
