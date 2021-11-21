package key

import (
	"bytes"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type bytesPC struct{}

func (b bytesPC) IsOrdered() bool {
	return false
}

func (b bytesPC) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	bz, err := io.ReadAll(r)
	return protoreflect.ValueOfBytes(bz), err
}

func (b bytesPC) Encode(value protoreflect.Value, w io.Writer) error {
	_, err := w.Write(value.Bytes())
	return err
}

func (b bytesPC) Compare(v1, v2 protoreflect.Value) int {
	return bytes.Compare(v1.Bytes(), v2.Bytes())
}

func (b bytesPC) IsEmpty(value protoreflect.Value) bool {
	return len(value.Bytes()) == 0
}

type bytesNT_PC struct{}

func (b bytesNT_PC) IsOrdered() bool {
	return false
}

func (b bytesNT_PC) IsEmpty(value protoreflect.Value) bool {
	return len(value.Bytes()) == 0
}

func (b bytesNT_PC) Compare(v1, v2 protoreflect.Value) int {
	return bytes.Compare(v1.Bytes(), v2.Bytes())
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

func (b bytesNT_PC) Encode(value protoreflect.Value, w io.Writer) error {
	bz := value.Bytes()
	n := len(bz)
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
