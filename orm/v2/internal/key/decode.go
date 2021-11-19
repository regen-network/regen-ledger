package key

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type keyPartDecoder func(r *bytes.Reader) (protoreflect.Value, error)

func makeKeyPartDecoder(field protoreflect.FieldDescriptor, nonTerminal bool) (keyPartDecoder, error) {
	if field.IsMap() {
		return nil, fmt.Errorf("map fields aren't supported in index keys")
	}

	switch field.Kind() {
	case protoreflect.BytesKind:
		if nonTerminal {
			return func(r *bytes.Reader) (protoreflect.Value, error) {
				n, err := r.ReadByte()
				if err != nil {
					return protoreflect.Value{}, err
				}

				bz := make([]byte, n)
				_, err = r.Read(bz)
				return protoreflect.ValueOfBytes(bz), err
			}, nil
		} else {
			return func(r *bytes.Reader) (protoreflect.Value, error) {
				bz, err := io.ReadAll(r)
				return protoreflect.ValueOfBytes(bz), err
			}, nil
		}
	case protoreflect.StringKind:
		if nonTerminal {
			return func(r *bytes.Reader) (protoreflect.Value, error) {
				var bz []byte
				for {
					b, err := r.ReadByte()
					if b == 0 || err == io.EOF {
						return protoreflect.ValueOfString(string(bz)), err
					}
					bz = append(bz, b)
				}
			}, nil
		} else {
			return func(r *bytes.Reader) (protoreflect.Value, error) {
				bz, err := io.ReadAll(r)
				return protoreflect.ValueOfString(string(bz)), err
			}, nil
		}
	case protoreflect.Uint32Kind:
		return func(r *bytes.Reader) (protoreflect.Value, error) {
			var x uint32
			err := binary.Read(r, binary.BigEndian, &x)
			return protoreflect.ValueOfUint32(x), err
		}, nil
	case protoreflect.Uint64Kind:
		return func(r *bytes.Reader) (protoreflect.Value, error) {
			var x uint64
			err := binary.Read(r, binary.BigEndian, &x)
			return protoreflect.ValueOfUint64(x), err
		}, nil
	default:
		return nil, fmt.Errorf("unsupported index key kind %s", field.Kind())
	}
}
