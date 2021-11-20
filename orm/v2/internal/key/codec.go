package key

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Codec struct {
	NumParts   int
	PartCodecs []PartCodec
	PKDecoder  func(r *bytes.Reader) ([]protoreflect.Value, error)
	Fields     []protoreflect.FieldDescriptor
}

func MakeCodec(fieldDescs []protoreflect.FieldDescriptor) (*Codec, error) {
	n := len(fieldDescs)
	var partCodecs []PartCodec
	for i := 0; i < n; i++ {
		nonTerminal := true
		if i == n-1 {
			nonTerminal = false
		}
		field := fieldDescs[i]
		enc, err := makePartCodec(field, nonTerminal)
		if err != nil {
			return nil, err
		}
		partCodecs = append(partCodecs, enc)
	}

	return &Codec{
		PartCodecs: partCodecs,
		NumParts:   n,
		Fields:     fieldDescs,
	}, nil
}

func (cdc *Codec) Encode(values []protoreflect.Value, w io.Writer, partial bool) error {
	for i := 0; i < cdc.NumParts; i++ {
		err := cdc.PartCodecs[i].Encode(values[i], w, partial)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cdc *Codec) GetValues(mref protoreflect.Message) []protoreflect.Value {
	var res []protoreflect.Value
	for _, f := range cdc.Fields {
		res = append(res, mref.Get(f))
	}
	return res
}

func (cdc *Codec) SetValues(mref protoreflect.Message, values []protoreflect.Value) {
	for i, f := range cdc.Fields {
		mref.Set(f, values[i])
	}
}

func (cdc *Codec) Decode(r *bytes.Reader) ([]protoreflect.Value, error) {
	values := make([]protoreflect.Value, cdc.NumParts)
	for i := 0; i < cdc.NumParts; i++ {
		value, err := cdc.PartCodecs[i].Decode(r)
		values[i] = value
		if err == io.EOF {
			if i == cdc.NumParts-1 {
				return values, nil
			} else {
				return nil, io.ErrUnexpectedEOF
			}
		} else if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func GetFieldDescriptors(desc protoreflect.MessageDescriptor, fields string) ([]protoreflect.FieldDescriptor, error) {
	fieldNames := strings.Split(fields, ",")
	if len(fieldNames) == 0 {
		return nil, fmt.Errorf("must specify non-empty list of fields, got %s", fields)
	}

	have := map[string]bool{}

	var fieldDescs []protoreflect.FieldDescriptor
	for _, fname := range fieldNames {
		if have[fname] {
			return nil, fmt.Errorf("duplicate field in key %s", fname)
		}

		have[fname] = true
		fieldDesc := GetFieldDescriptor(desc, fname)
		if fieldDesc == nil {
			return nil, fmt.Errorf("unable to resolve field %s", fname)
		}

		fieldDescs = append(fieldDescs, fieldDesc)
	}
	return fieldDescs, nil
}

func GetFieldDescriptor(desc protoreflect.MessageDescriptor, fname string) protoreflect.FieldDescriptor {
	if desc == nil {
		return nil
	}

	return desc.Fields().ByName(protoreflect.Name(fname))
}
