package key

import (
	"bytes"
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type uint64PC struct{}

func (u uint64PC) decode(r *bytes.Reader) (protoreflect.Value, error) {
	var x uint64
	err := binary.Read(r, binary.BigEndian, &x)
	return protoreflect.ValueOfUint64(x), err
}

func (u uint64PC) encode(value protoreflect.Value, w io.Writer, partial bool) error {
	return binary.Write(w, binary.BigEndian, value.Uint())
}
