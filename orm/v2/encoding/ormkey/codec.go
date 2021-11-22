package ormkey

import (
	"bytes"
	"io"
	"strings"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormvalue"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Codec struct {
	prefix      []byte
	Fields      []protoreflect.FieldDescriptor
	ValueCodecs []ormvalue.Codec
}

func MakeCodec(prefix []byte, fieldDescs []protoreflect.FieldDescriptor) (*Codec, error) {
	n := len(fieldDescs)
	var partCodecs []ormvalue.Codec
	for i := 0; i < n; i++ {
		nonTerminal := true
		if i == n-1 {
			nonTerminal = false
		}
		field := fieldDescs[i]
		enc, err := ormvalue.MakeCodec(field, nonTerminal)
		if err != nil {
			return nil, err
		}
		partCodecs = append(partCodecs, enc)
	}

	return &Codec{
		ValueCodecs: partCodecs,
		Fields:      fieldDescs,
		prefix:      prefix,
	}, nil
}

func (cdc *Codec) EncodeWriter(values []protoreflect.Value, w io.Writer) error {
	_, err := w.Write(cdc.prefix)
	if err != nil {
		return err
	}

	for i := 0; i < len(values); i++ {
		err = cdc.ValueCodecs[i].Encode(values[i], w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cdc *Codec) Encode(values []protoreflect.Value) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := cdc.EncodeWriter(values, buf)
	return buf.Bytes(), err
}

// EncodePartial encodes the key up to the presence of any empty values in the
// list of key values.
func (cdc *Codec) EncodePartial(message protoreflect.Message) ([]protoreflect.Value, []byte, error) {
	lastNonEmpty := 0
	n := len(cdc.ValueCodecs)
	values := make([]protoreflect.Value, n)
	for i := 0; i < n; i++ {
		f := cdc.Fields[i]
		if message.Has(f) {
			lastNonEmpty = i + 1
		}
		values[i] = message.Get(f)
	}

	var b bytes.Buffer
	err := cdc.EncodeWriter(values[:lastNonEmpty], &b)
	if err != nil {
		return nil, nil, err
	}
	return values, b.Bytes(), nil
}

func (cdc *Codec) GetValues(mref protoreflect.Message) []protoreflect.Value {
	var res []protoreflect.Value
	for _, f := range cdc.Fields {
		res = append(res, mref.Get(f))
	}
	return res
}

func (cdc *Codec) ClearKey(mref protoreflect.Message) {
	for _, f := range cdc.Fields {
		mref.Clear(f)
	}
}

func (cdc *Codec) SetValues(mref protoreflect.Message, values []protoreflect.Value) {
	for i, f := range cdc.Fields {
		mref.Set(f, values[i])
	}
}

func SkipPrefix(r *bytes.Reader, prefix []byte) error {
	n := len(prefix)
	if n > 0 {
		// we skip checking the prefix for performance reasons because we assume
		// that it was checked by the caller
		_, err := r.Seek(int64(n), io.SeekCurrent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cdc *Codec) Decode(r *bytes.Reader) ([]protoreflect.Value, error) {
	err := SkipPrefix(r, cdc.prefix)
	if err != nil {
		return nil, err
	}

	n := len(cdc.ValueCodecs)
	values := make([]protoreflect.Value, n)
	for i := 0; i < n; i++ {
		value, err := cdc.ValueCodecs[i].Decode(r)
		values[i] = value
		if err == io.EOF {
			if i == n-1 {
				return values, nil
			} else {
				return nil, io.ErrUnexpectedEOF
			}
		} else if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func (cdc *Codec) EncodeFromMessage(message protoreflect.Message) ([]protoreflect.Value, []byte, error) {
	var b bytes.Buffer
	values := cdc.GetValues(message)
	err := cdc.EncodeWriter(values, &b)
	if err != nil {
		return nil, nil, err
	}
	return values, b.Bytes(), nil
}

// IsFullyOrdered returns true if all parts are also ordered
func (cdc *Codec) IsFullyOrdered() bool {
	for _, p := range cdc.ValueCodecs {
		if !p.IsOrdered() {
			return false
		}
	}
	return true
}

func (cdc *Codec) CompareValues(values1, values2 []protoreflect.Value) int {
	n := len(values1)
	if n != len(values2) {
		panic("expected arrays of the same length")
	}
	if n > len(cdc.ValueCodecs) {
		panic("array is too long")
	}

	var cmp int
	for i := 0; i < n; i++ {
		cmp = cdc.ValueCodecs[i].Compare(values1[i], values2[i])
		// any non-equal parts determine our ordering
		if cmp != 0 {
			break
		}
	}

	return cmp
}

func GetFieldDescriptors(desc protoreflect.MessageDescriptor, fields string) ([]protoreflect.FieldDescriptor, error) {
	if len(fields) == 0 {
		return nil, ormerrors.InvalidKeyFields.Wrapf("got fields %q for table %q", fields, desc.FullName())
	}

	fieldNames := strings.Split(fields, ",")

	have := map[string]bool{}

	var fieldDescs []protoreflect.FieldDescriptor
	for _, fname := range fieldNames {
		if have[fname] {
			return nil, ormerrors.DuplicateKeyField.Wrapf("field %q in %q", fname, desc.FullName())
		}

		have[fname] = true
		fieldDesc := GetFieldDescriptor(desc, fname)
		if fieldDesc == nil {
			return nil, ormerrors.FieldNotFound.Wrapf("field %q in %q", fname, desc.FullName())
		}

		fieldDescs = append(fieldDescs, fieldDesc)
	}
	return fieldDescs, nil
}

func GetFieldDescriptor(desc protoreflect.MessageDescriptor, fname string) protoreflect.FieldDescriptor {
	if desc == nil {
		return nil
	}

	return desc.Fields().ByName(protoreflect.Name(fname))
}

func (cdc Codec) Size(values []protoreflect.Value) (int, error) {
	panic("TODO")
}

func (cdc Codec) Prefix() []byte {
	return cdc.prefix
}

type CodecI interface {
	EncodeWriter(values []protoreflect.Value, w io.Writer) error
	EncodePartial(message protoreflect.Message) ([]protoreflect.Value, []byte, error)
	GetValues(mref protoreflect.Message) []protoreflect.Value
	ClearKey(mref protoreflect.Message)
	SetValues(mref protoreflect.Message, values []protoreflect.Value)
	Decode(r *bytes.Reader) ([]protoreflect.Value, error)
	EncodeFromMessage(message protoreflect.Message) ([]protoreflect.Value, []byte, error)
	IsFullyOrdered() bool
	CompareValues(values1, values2 []protoreflect.Value) int
	Size(values []protoreflect.Value) (int, error)
	Prefix() []byte
}
