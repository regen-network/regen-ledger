package key

import (
	"bytes"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type stringPC struct{}

func (s stringPC) Equal(v1, v2 protoreflect.Value) bool {
	return v1.String() == v2.String()
}

func (s stringPC) decode(r *bytes.Reader) (protoreflect.Value, error) {
	bz, err := io.ReadAll(r)
	return protoreflect.ValueOfString(string(bz)), err
}

func (s stringPC) encode(value protoreflect.Value, w io.Writer, partial bool) error {
	_, err := w.Write([]byte(value.String()))
	return err
}

type stringNT_PC struct{}

func (s stringNT_PC) Equal(v1, v2 protoreflect.Value) bool {
	return v1.String() == v2.String()
}

func (s stringNT_PC) decode(r *bytes.Reader) (protoreflect.Value, error) {
	var bz []byte
	for {
		b, err := r.ReadByte()
		if b == 0 || err == io.EOF {
			return protoreflect.ValueOfString(string(bz)), err
		}
		bz = append(bz, b)
	}
}

func (s stringNT_PC) encode(value protoreflect.Value, w io.Writer, partial bool) error {
	str := value.String()
	if str == "" && partial {
		return io.EOF
	}

	bz := []byte(str)
	for _, b := range bz {
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
}

var nullTerminator = []byte{0}
