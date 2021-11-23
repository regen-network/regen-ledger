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
	prefix    []byte
	fixedSize int
	varSizers []struct {
		cdc ormvalue.Codec
		i   int
	}
	Fields      []protoreflect.FieldDescriptor
	ValueCodecs []ormvalue.Codec
}

func MakeCodec(prefix []byte, fieldDescs []protoreflect.FieldDescriptor) (*Codec, error) {
	n := len(fieldDescs)
	var valueCodecs []ormvalue.Codec
	var varSizers []struct {
		cdc ormvalue.Codec
		i   int
	}
	fixedSize := 0
	for i := 0; i < n; i++ {
		nonTerminal := true
		if i == n-1 {
			nonTerminal = false
		}
		field := fieldDescs[i]
		cdc, err := ormvalue.MakeCodec(field, nonTerminal)
		if err != nil {
			return nil, err
		}
		if x := cdc.FixedSize(); x > 0 {
			fixedSize += x
		} else {
			varSizers = append(varSizers, struct {
				cdc ormvalue.Codec
				i   int
			}{cdc, i})
		}
		valueCodecs = append(valueCodecs, cdc)
	}

	return &Codec{
		ValueCodecs: valueCodecs,
		Fields:      fieldDescs,
		prefix:      prefix,
		fixedSize:   fixedSize,
		varSizers:   varSizers,
	}, nil
}

func (cdc *Codec) Encode(values []protoreflect.Value) ([]byte, error) {
	sz, err := cdc.Size(values)
	if err != nil {
		return nil, err
	}

	w := bytes.NewBuffer(make([]byte, 0, sz))
	_, err = w.Write(cdc.prefix)
	if err != nil {
		return nil, err
	}

	n := len(values)
	if n > len(cdc.ValueCodecs) {
		return nil, ormerrors.IndexOutOfBounds
	}

	for i := 0; i < n; i++ {
		err = cdc.ValueCodecs[i].Encode(values[i], w)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}

func (cdc *Codec) GetValues(mref protoreflect.Message) []protoreflect.Value {
	var res []protoreflect.Value
	for _, f := range cdc.Fields {
		res = append(res, mref.Get(f))
	}
	return res
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
	values := cdc.GetValues(message)
	bz, err := cdc.Encode(values)
	return values, bz, err
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
	size := cdc.fixedSize
	n := len(values)
	for _, sz := range cdc.varSizers {
		if sz.i >= n {
			return size, nil
		}
		x, err := sz.cdc.Size(values[sz.i])
		if err != nil {
			return 0, err
		}
		size += x
	}
	return size, nil
}

func (cdc *Codec) SetValues(mref protoreflect.Message, values []protoreflect.Value) {
	for i, f := range cdc.Fields {
		mref.Set(f, values[i])
	}
}

func (cdc Codec) Prefix() []byte {
	return cdc.prefix
}

type CodecI interface {
	Encode(values []protoreflect.Value) ([]byte, error)
	GetValues(mref protoreflect.Message) []protoreflect.Value
	Decode(r *bytes.Reader) ([]protoreflect.Value, error)
	EncodeFromMessage(message protoreflect.Message) ([]protoreflect.Value, []byte, error)
	IsFullyOrdered() bool
	CompareValues(values1, values2 []protoreflect.Value) int
	Size(values []protoreflect.Value) (int, error)
	Prefix() []byte
}
