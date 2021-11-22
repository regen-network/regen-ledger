package key

import (
	"bytes"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func MakeIndexKeyCodec(prefix []byte, indexFields []protoreflect.FieldDescriptor, primaryKeyFields []protoreflect.FieldDescriptor) (*Codec, error) {
	indexFieldMap := map[protoreflect.Name]int{}
	pkFieldOrderMap := map[int]int{}

	var keyFields []protoreflect.FieldDescriptor
	for i, f := range indexFields {
		indexFieldMap[f.Name()] = i
		keyFields = append(keyFields, f)
	}

	for j, f := range primaryKeyFields {
		if i, ok := indexFieldMap[f.Name()]; ok {
			pkFieldOrderMap[j] = i
			continue
		}
		keyFields = append(keyFields, f)
		pkFieldOrderMap[j] = j
	}

	cdc, err := MakeCodec(prefix, keyFields)
	if err != nil {
		return nil, err
	}

	numPrimaryKeyFields := len(primaryKeyFields)
	cdc.NumIndexParts = len(indexFields)
	cdc.PKDecoder = func(r *bytes.Reader) ([]protoreflect.Value, error) {
		fields, err := cdc.Decode(r)
		if err != nil {
			return nil, err
		}

		pkValues := make([]protoreflect.Value, numPrimaryKeyFields)

		for i := 0; i < numPrimaryKeyFields; i++ {
			pkValues[i] = fields[pkFieldOrderMap[i]]
		}

		return pkValues, nil
	}
	cdc.SplitIndexPK = func(fields []protoreflect.Value) (idxKey, pk []protoreflect.Value, err error) {
		pkValues := make([]protoreflect.Value, numPrimaryKeyFields)

		for i := 0; i < numPrimaryKeyFields; i++ {
			pkValues[i] = fields[pkFieldOrderMap[i]]
		}

		return pkValues, fields[:cdc.NumIndexParts], nil
	}

	return cdc, nil
}
