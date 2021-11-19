package key

import (
	"bytes"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Codec struct {
	NumParts     int
	PartEncoders []keyPartEncoder
	PartDecoders []keyPartDecoder
	PKDecoder    func(r *bytes.Reader) ([]protoreflect.Value, error)
}

func MakeCodec(fieldDescs []protoreflect.FieldDescriptor, isPrimaryKey bool) (*Codec, error) {
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

	return &Codec{
		PartEncoders: encoders,
		PartDecoders: decoders,
		NumParts:     n,
	}, nil
}

func (cdc *Codec) Encode(values []protoreflect.Value, w io.Writer, partial bool) error {
	for i := 0; i < cdc.NumParts; i++ {
		err := cdc.PartEncoders[i](values[i], w, partial)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cdc *Codec) Decode(r *bytes.Reader) ([]protoreflect.Value, error) {
	values := make([]protoreflect.Value, cdc.NumParts)
	for i := 0; i < cdc.NumParts; i++ {
		value, err := cdc.PartDecoders[i](r)
		values = append(values, value)
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
