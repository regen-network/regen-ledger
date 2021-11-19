package key

import (
	"encoding/binary"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type keyPartEncoder func(value protoreflect.Value, w io.Writer, partial bool) error

var nullTerminator = []byte{0}

func makeKeyPartEncoder(field protoreflect.FieldDescriptor, nonTerminal bool) (keyPartEncoder, error) {
	if field.IsMap() {
		return nil, fmt.Errorf("map fields aren't supported in index keys")
	}

	switch field.Kind() {
	case protoreflect.BytesKind:
		if nonTerminal {
			return func(value protoreflect.Value, w io.Writer, partial bool) error {
				bz := value.Bytes()
				n := len(bz)
				if n == 0 && partial {
					return io.EOF
				}

				if n > 255 {
					return fmt.Errorf("can't encode a byte array longer than 255 bytes as an index part")
				}
				_, err := w.Write([]byte{byte(n)})
				if err != nil {
					return err
				}
				_, err = w.Write(bz)
				return err
			}, nil
		} else {
			return func(value protoreflect.Value, w io.Writer, partial bool) error {
				_, err := w.Write(value.Bytes())
				return err
			}, nil
		}
	case protoreflect.StringKind:
		if nonTerminal {
			return func(value protoreflect.Value, w io.Writer, partial bool) error {
				str := value.String()
				if str == "" && partial {
					return io.EOF
				}

				for b := range []byte(str) {
					if b == 0 {
						return fmt.Errorf("illegal null terminator found in index string: %s", str)
					}
				}
				_, err := w.Write([]byte(str))
				if err != nil {
					return err
				}
				_, err = w.Write(nullTerminator)
				return err
			}, nil
		} else {
			return func(value protoreflect.Value, w io.Writer, partial bool) error {
				_, err := w.Write([]byte(value.String()))
				return err
			}, nil
		}
	case protoreflect.Uint32Kind:
		return func(value protoreflect.Value, w io.Writer, partial bool) error {
			return binary.Write(w, binary.BigEndian, uint32(value.Uint()))
		}, nil
	case protoreflect.Uint64Kind:
		return func(value protoreflect.Value, w io.Writer, partial bool) error {
			return binary.Write(w, binary.BigEndian, value.Uint())
		}, nil
	default:
		return nil, fmt.Errorf("unsupported index key kind %s", field.Kind())
	}
}
