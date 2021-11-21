package key

import (
	"bytes"
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type uint64PC struct{}

func (u uint64PC) IsOrdered() bool {
	return true
}

func (u uint64PC) IsEmpty(value protoreflect.Value) bool {
	return value.Uint() == 0
}

func (u uint64PC) Compare(v1, v2 protoreflect.Value) int {
	return compareUint(v1, v2)
}

func (u uint64PC) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	var x uint64
	err := binary.Read(r, binary.BigEndian, &x)
	return protoreflect.ValueOfUint64(x), err
}

func (u uint64PC) Encode(value protoreflect.Value, w io.Writer) error {
	return binary.Write(w, binary.BigEndian, value.Uint())
}

func compareUint(v1, v2 protoreflect.Value) int {
	x := v1.Uint()
	y := v2.Uint()
	if x == y {
		return 0
	} else if x < y {
		return -1
	} else {
		return 1
	}
}
