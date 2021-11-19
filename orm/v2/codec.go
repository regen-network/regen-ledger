package v2

import (
	"bytes"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type keyCodec struct {
	numParts     int
	partEncoders []keyPartEncoder
	partDecoders []keyPartDecoder
	pkDecoder    func(r *bytes.Reader) ([]protoreflect.Value, error)
}

func makeKeyCodec(fieldDescs []protoreflect.FieldDescriptor, isPrimaryKey bool) (*keyCodec, error) {
	n := len(fieldDescs)
	var encoders []keyPartEncoder
	var decoders []keyPartDecoder
	for i := 0; i < n; i++ {
		nonTerminal := true
		if i == n-1 {
			nonTerminal = false
		}
		field := fieldDescs[i]
		if field.IsList() && isPrimaryKey {
			return nil, fmt.Errorf("repeated fields not allowed in primary key")
		}

		enc, err := makeKeyPartEncoder(field, nonTerminal)
		if err != nil {
			return nil, err
		}
		encoders = append(encoders, enc)

		dec, err := makeKeyPartDecoder(field, nonTerminal)
		if err != nil {
			return nil, err
		}
		decoders = append(decoders, dec)
	}

	return &keyCodec{
		partEncoders: encoders,
		partDecoders: decoders,
		numParts:     n,
	}, nil
}

func (cdc *keyCodec) encode(values []protoreflect.Value, w io.Writer, partial bool) error {
	for i := 0; i < cdc.numParts; i++ {
		err := cdc.partEncoders[i](values[i], w, partial)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cdc *keyCodec) decode(r *bytes.Reader) ([]protoreflect.Value, error) {
	values := make([]protoreflect.Value, cdc.numParts)
	for i := 0; i < cdc.numParts; i++ {
		value, err := cdc.partDecoders[i](r)
		values = append(values, value)
		if err == io.EOF {
			if i == cdc.numParts-1 {
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
