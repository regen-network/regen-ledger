package key

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type stringPC struct{}

func (s stringPC) IsOrdered() bool {
	return true
}

func (s stringPC) Compare(v1, v2 protoreflect.Value) int {
	return strings.Compare(v1.String(), v2.String())
}

func (s stringPC) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	bz, err := io.ReadAll(r)
	return protoreflect.ValueOfString(string(bz)), err
}

func (s stringPC) Encode(value protoreflect.Value, w io.Writer) error {
	_, err := w.Write([]byte(value.String()))
	return err
}

func (s stringPC) IsEmpty(value protoreflect.Value) bool {
	return value.String() == ""
}

type stringNT_PC struct{}

func (s stringNT_PC) IsOrdered() bool {
	return true
}

func (s stringNT_PC) IsEmpty(value protoreflect.Value) bool {
	return value.String() == ""
}

func (s stringNT_PC) Compare(v1, v2 protoreflect.Value) int {
	return strings.Compare(v1.String(), v2.String())
}

func (s stringNT_PC) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	var bz []byte
	for {
		b, err := r.ReadByte()
		if b == 0 || err == io.EOF {
			return protoreflect.ValueOfString(string(bz)), err
		}
		bz = append(bz, b)
	}
}

func (s stringNT_PC) Encode(value protoreflect.Value, w io.Writer) error {
	str := value.String()
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
