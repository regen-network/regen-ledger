package key

import (
	"bytes"
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type uint32PC struct{}

func (u uint32PC) IsOrdered() bool {
	return true
}

func (u uint32PC) IsEmpty(value protoreflect.Value) bool {
	return value.Uint() == 0
}

func (u uint32PC) Compare(v1, v2 protoreflect.Value) int {
	return compareUint(v1, v2)
}

func (u uint32PC) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	var x uint32
	err := binary.Read(r, binary.BigEndian, &x)
	return protoreflect.ValueOfUint32(x), err
}

func (u uint32PC) Encode(value protoreflect.Value, w io.Writer) error {
	return binary.Write(w, binary.BigEndian, uint32(value.Uint()))
}
