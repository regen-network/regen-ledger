package key

import (
	"bytes"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type bytesPC struct{}

func (b bytesPC) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	bz, err := io.ReadAll(r)
	return protoreflect.ValueOfBytes(bz), err
}

func (b bytesPC) Encode(value protoreflect.Value, w io.Writer, partial bool) error {
	_, err := w.Write(value.Bytes())
	return err
}

func (b bytesPC) Equal(v1, v2 protoreflect.Value) bool {
	return bytes.Equal(v1.Bytes(), v2.Bytes())
}

type bytesNT_PC struct{}

func (b bytesNT_PC) Equal(v1, v2 protoreflect.Value) bool {
	return bytes.Equal(v1.Bytes(), v2.Bytes())
}

func (b bytesNT_PC) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	n, err := r.ReadByte()
	if err != nil {
		return protoreflect.Value{}, err
	}

	if n == 0 {
		return protoreflect.ValueOfBytes([]byte{}), nil
	}

	bz := make([]byte, n)
	_, err = r.Read(bz)
	return protoreflect.ValueOfBytes(bz), err
}

func (b bytesNT_PC) Encode(value protoreflect.Value, w io.Writer, partial bool) error {
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
}
