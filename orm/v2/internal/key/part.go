package key

import (
	"bytes"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type PartCodec interface {
	Decode(r *bytes.Reader) (protoreflect.Value, error)
	Encode(value protoreflect.Value, w io.Writer, partial bool) error
	Equal(v1, v2 protoreflect.Value) bool
}

func makePartCodec(field protoreflect.FieldDescriptor, nonTerminal bool) (PartCodec, error) {
	if field.IsMap() || field.IsList() {
		return nil, fmt.Errorf("map fields aren't supported in keys")
	}

	switch field.Kind() {
	case protoreflect.BytesKind:
		if nonTerminal {
			return bytesNT_PC{}, nil
		} else {
			return bytesPC{}, nil
		}
	case protoreflect.StringKind:
		if nonTerminal {
			return stringNT_PC{}, nil
		} else {
			return stringPC{}, nil
		}
	case protoreflect.Uint32Kind:
		return uint32PC{}, nil
	case protoreflect.Uint64Kind:
		return uint64PC{}, nil
	default:
		return nil, fmt.Errorf("unsupported index key kind %s", field.Kind())
	}
}
