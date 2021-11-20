package key

import (
	"bytes"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type PartCodec interface {
	decode(r *bytes.Reader) (protoreflect.Value, error)
	encode(value protoreflect.Value, w io.Writer, partial bool) error
}

func makePartCodec(field protoreflect.FieldDescriptor, nonTerminal bool) (PartCodec, error) {
	if field.IsMap() {
		return nil, fmt.Errorf("map fields aren't supported in index keys")
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
