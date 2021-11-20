package key

import (
	"bytes"
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type uint32PC struct{}

func (u uint32PC) decode(r *bytes.Reader) (protoreflect.Value, error) {
	var x uint32
	err := binary.Read(r, binary.BigEndian, &x)
	return protoreflect.ValueOfUint32(x), err
}

func (u uint32PC) encode(value protoreflect.Value, w io.Writer, partial bool) error {
	return binary.Write(w, binary.BigEndian, uint32(value.Uint()))
}
