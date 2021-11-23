package orm

import (
	"context"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"

	"google.golang.org/protobuf/proto"
)

type ReadStore interface {
	mustEmbedUnimplementedReadStore()

	Has(...proto.Message) bool
	Get(...proto.Message) (found bool, err error)
	List(msgType proto.Message, options *ListOptions) Iterator
}

type Store interface {
	mustEmbedUnimplementedStore()

	ReadStore

	Create(...proto.Message) error
	Save(...proto.Message) error
	Delete(...proto.Message) error
}

type Iterator interface {
	mustEmbedUnimplementedIterator()

	Next(proto.Message) (bool, error)
	Cursor() Cursor
	Close()
}

type Cursor []byte

type ListOptions struct {
	Reverse bool
	// Index defines an index by field names. If it is empty the primary key
	// will be used as the index.
	Index string

	// Cursor specifies a cursor returned by Iterator.Cursor() to restart iteration from.
	Cursor Cursor

	// Prefix defines an iteration prefix using values corresponding the the key
	// being used. Not all of the values in the key need to be specified and
	// they do not be sortable unlike start and end. Prefix or Start/End are
	// mutually exclusive and shouldn't be specified together.
	Prefix []protoreflect.Value

	// Start defines a start position using a set of values corresponding to the
	// index or primary key being used. Each of the values must match the type
	// of the key at that position and also be a sortable value. Not all of
	// the values in the key need to be specified.
	Start []protoreflect.Value

	// End defines an end position using a set of values correspond to the key
	// being used. Not all of the values in the key need to be specified.
	End []protoreflect.Value
}

type ReadStoreConn interface {
	OpenRead(context.Context) (ReadStore, error)
}

type StoreConn interface {
	ReadStoreConn
	Open(context.Context) (Store, error)
}

type UnimplementedReadStore struct{}

func (UnimplementedReadStore) mustEmbedUnimplementedReadStore() {}

func (u UnimplementedReadStore) Has(...proto.Message) bool {
	return false
}

func (u UnimplementedReadStore) Get(...proto.Message) (found bool, err error) {
	return false, ormerrors.UnsupportedOperation
}

func (u UnimplementedReadStore) List(proto.Message, *ListOptions) Iterator {
	return ErrIterator{Err: ormerrors.UnsupportedOperation}
}

var _ ReadStore = UnimplementedReadStore{}

type UnimplementedStore struct {
	UnimplementedReadStore
}

func (u UnimplementedStore) mustEmbedUnimplementedStore() {}

func (u UnimplementedStore) Create(...proto.Message) error {
	return ormerrors.UnsupportedOperation
}

func (u UnimplementedStore) Save(...proto.Message) error {
	return ormerrors.UnsupportedOperation
}

func (u UnimplementedStore) Delete(...proto.Message) error {
	return ormerrors.UnsupportedOperation
}

var _ Store = UnimplementedStore{}

type UnimplementedIterator struct{}

func (u UnimplementedIterator) mustEmbedUnimplementedIterator() {}

func (u UnimplementedIterator) Next(proto.Message) (bool, error) {
	return false, ormerrors.UnsupportedOperation
}

func (u UnimplementedIterator) Cursor() Cursor { return nil }

func (u UnimplementedIterator) Close() {
	panic("implement me")
}

var _ Iterator = UnimplementedIterator{}

type ErrIterator struct {
	UnimplementedIterator
	Err error
}

func (e ErrIterator) Cursor() Cursor { return nil }

func (e ErrIterator) Next(proto.Message) (bool, error) { return false, e.Err }

func (e ErrIterator) Close() {}

var _ Iterator = ErrIterator{}
