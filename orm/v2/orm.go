package v2

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DynamicStore struct {
	stores map[protoreflect.FullName]store
}

type KVStore interface {
	Get(key []byte) []byte
	Set(key, value []byte)
	Has(key []byte) bool
	Delete(key []byte)
}

type store interface {
	Create(kv KVStore, message proto.Message) error
	Read(kv KVStore, message proto.Message) (bool, error)
	Update(kv KVStore, message proto.Message) error
	Delete(kv KVStore, message proto.Message) error
	List(message proto.Message, options ...ListOption) Iterator
}

func (s DynamicStore) Create(kv KVStore, message proto.Message) error {
	panic("TODO")
}

func getFieldDescriptors(desc protoreflect.MessageDescriptor, fields string) []protoreflect.FieldDescriptor {
	fieldNames := strings.Split(fields, ",")
	var fieldDescs []protoreflect.FieldDescriptor
	for _, fname := range fieldNames {
		fieldDesc := desc.Fields().ByName(protoreflect.Name(fname))
		fieldDescs = append(fieldDescs, fieldDesc)
	}
	return fieldDescs
}

type keyEncoder func(values []protoreflect.Value, w io.Writer) error

func makeKeyEncoder(fieldDescs []protoreflect.FieldDescriptor) (keyEncoder, error) {
	n := len(fieldDescs)
	var encoders []indexPartEncoder
	for i := 0; i < n; i++ {
		nonTerminal := true
		if i == n-1 {
			nonTerminal = false
		}
		field := fieldDescs[i]
		enc, err := makeIndexPartEncoder(field, nonTerminal)
		if err != nil {
			return nil, err
		}
		encoders = append(encoders, enc)
	}

	return func(values []protoreflect.Value, w io.Writer) error {
		for i := 0; i < n; i++ {
			err := encoders[i](values[i], w)
			if err != nil {
				return err
			}
		}
		return nil
	}, nil
}

type indexPartEncoder func(value protoreflect.Value, w io.Writer) error

var zeroBuf = []byte{0}
var nullTerminator = zeroBuf

func makeIndexPartEncoder(field protoreflect.FieldDescriptor, nonTerminal bool) (indexPartEncoder, error) {
	if field.IsList() || field.IsMap() {
		return nil, fmt.Errorf("repeated and map fields aren't supported in index keys")
	}

	switch field.Kind() {
	case protoreflect.BytesKind:
		if nonTerminal {
			return func(value protoreflect.Value, w io.Writer) error {
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
			}, nil
		} else {
			return func(value protoreflect.Value, w io.Writer) error {
				_, err := w.Write(value.Bytes())
				return err
			}, nil
		}
	case protoreflect.StringKind:
		if nonTerminal {
			return func(value protoreflect.Value, w io.Writer) error {
				_, err := w.Write([]byte(value.String()))
				if err != nil {
					return err
				}
				_, err = w.Write(nullTerminator)
				return err
			}, nil
		} else {
			return func(value protoreflect.Value, w io.Writer) error {
				_, err := w.Write([]byte(value.String()))
				return err
			}, nil
		}
	case protoreflect.Uint32Kind:
		return func(value protoreflect.Value, w io.Writer) error {
			return binary.Write(w, binary.BigEndian, uint32(value.Uint()))
		}, nil
	case protoreflect.Uint64Kind:
		return func(value protoreflect.Value, w io.Writer) error {
			return binary.Write(w, binary.BigEndian, value.Uint())
		}, nil
	default:
		return nil, fmt.Errorf("unsupported index key kind %s", field.Kind())
	}
}

func (s DynamicStore) Read(message proto.Message) error {
	panic("TODO")
}

func (s DynamicStore) Update(message proto.Message) error {
	panic("TODO")
}

func (s DynamicStore) Delete(message proto.Message) error {
	panic("TODO")
}

func (s DynamicStore) List(message proto.Message, options ...ListOption) Iterator {
	panic("TODO")
}

func (s DynamicStore) getTableStore(message proto.Message) (store, error) {
	desc := message.ProtoReflect().Descriptor()
	if ts, ok := s.stores[desc.FullName()]; ok {
		return ts, nil
	}

	if tableDesc := proto.GetExtension(desc.Options(), E_Table).(*TableDescriptor); tableDesc != nil {
		pkFields := getFieldDescriptors(desc, tableDesc.PrimaryKey.Fields)
		pkEncoder, err := makeKeyEncoder(pkFields)
		if err != nil {
			return nil, err
		}

		numPrimaryKeyFields := len(pkFields)
		tableId := tableDesc.Id
		prefix := make([]byte, binary.MaxVarintLen32)
		n := binary.PutUvarint(prefix, uint64(tableId))
		prefix = prefix[:n]
		pkPrefix := make([]byte, n+1)
		copy(pkPrefix, prefix)
		pkPrefix[n] = 0

		store := &dynamicTableStore{
			numPrimaryKeyFields: numPrimaryKeyFields,
			prefix:              prefix,
			pkFields:            pkFields,
			pkEncoder:           pkEncoder,
		}

		s.stores[desc.FullName()] = store
		return store, nil
	}

	panic("TODO")
}

type Iterator interface {
	isIterator()
	Next(proto.Message) bool
}

type ListOption interface {
	isListOption()
}

func Reverse() ListOption { panic("TODO") }

func IndexHint(fields string) ListOption { panic("TODO") }

type dynamicTableStore struct {
	numPrimaryKeyFields int
	pkFields            []protoreflect.FieldDescriptor
	prefix              []byte
	pkEncoder           keyEncoder
}

func (d dynamicTableStore) primaryKey(message proto.Message) ([]protoreflect.Value, []byte, error) {
	refm := message.ProtoReflect()
	// encode primary key
	pkValues := make([]protoreflect.Value, d.numPrimaryKeyFields)
	for i, f := range d.pkFields {
		pkValues[i] = refm.Get(f)
	}

	pkBuf := &bytes.Buffer{}
	pkBuf.Write(d.prefix)
	pkBuf.WriteByte(0)
	err := d.pkEncoder(pkValues, pkBuf)
	if err != nil {
		return nil, nil, err
	}

	return pkValues, pkBuf.Bytes(), nil
}

func (d dynamicTableStore) Create(kv KVStore, message proto.Message) error {
	pkValues, pk, err := d.primaryKey(message)
	if err != nil {
		return err
	}

	refm := message.ProtoReflect()

	// temporarily clear primary key
	for _, f := range d.pkFields {
		refm.Clear(f)
	}

	// store object
	bz, err := proto.Marshal(message)
	kv.Set(pk, bz)

	// set primary key again
	for i, f := range d.pkFields {
		refm.Set(f, pkValues[i])
	}

	// TODO: build indexes

	return nil
}

func (d dynamicTableStore) Read(kv KVStore, message proto.Message) (bool, error) {
	pkValues, pk, err := d.primaryKey(message)
	if err != nil {
		return false, err
	}

	// store object
	bz := kv.Get(pk)
	if bz == nil {
		return false, nil
	}

	err = proto.Unmarshal(bz, message)
	if err != nil {
		return true, err
	}

	refm := message.ProtoReflect()

	// rehydrate primary key
	for i, f := range d.pkFields {
		refm.Set(f, pkValues[i])
	}

	return true, nil
}

func (d dynamicTableStore) Update(kv KVStore, message proto.Message) error {
	panic("implement me")
}

func (d dynamicTableStore) Delete(kv KVStore, message proto.Message) error {
	_, pk, err := d.primaryKey(message)
	if err != nil {
		return err
	}

	// TODO: clear indexes

	// delete object
	kv.Delete(pk)

	return nil
}

func (d dynamicTableStore) List(message proto.Message, options ...ListOption) Iterator {
	panic("implement me")
}
