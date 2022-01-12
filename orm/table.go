package orm

import (
	"bytes"
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var _ Indexable = &tableBuilder{}

type tableBuilder struct {
	model       reflect.Type
	prefixData  byte
	storeKey    sdk.StoreKey
	afterSet    []AfterSetInterceptor
	afterDelete []AfterDeleteInterceptor
	cdc         codec.Codec
}

// newTableBuilder creates a builder to setup a table object.
func newTableBuilder(prefixData byte, storeKey sdk.StoreKey, model codec.ProtoMarshaler, cdc codec.Codec) (*tableBuilder, error) {
	if model == nil {
		return nil, ErrArgument.Wrap("Model must not be nil")
	}
	if storeKey == nil {
		return nil, ErrArgument.Wrap("StoreKey must not be nil")
	}
	tp := reflect.TypeOf(model)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	return &tableBuilder{
		prefixData: prefixData,
		storeKey:   storeKey,
		model:      tp,
		cdc:        cdc,
	}, nil
}

// TestTableBuilder exposes the private tableBuilder type for testing purposes.
// It is not safe to use this outside of test code.
func TestTableBuilder(prefixData byte, storeKey sdk.StoreKey, model codec.ProtoMarshaler, cdc codec.Codec) (*tableBuilder, error) {
	return newTableBuilder(prefixData, storeKey, model, cdc)
}

// RowGetter returns a type safe RowGetter.
func (a tableBuilder) RowGetter() RowGetter {
	return NewTypeSafeRowGetter(a.storeKey, a.prefixData, a.model, a.cdc)
}

func (a tableBuilder) StoreKey() sdk.StoreKey {
	return a.storeKey
}

// Build creates a new table object.
func (a tableBuilder) Build() table {
	return table{
		model:       a.model,
		prefix:      a.prefixData,
		storeKey:    a.storeKey,
		afterSet:    a.afterSet,
		afterDelete: a.afterDelete,
		cdc:         a.cdc,
	}
}

// AddAfterSetInterceptor can be used to register a callback function that is executed after an object is created and/or updated.
func (a *tableBuilder) AddAfterSetInterceptor(interceptor AfterSetInterceptor) {
	a.afterSet = append(a.afterSet, interceptor)
}

// AddAfterDeleteInterceptor can be used to register a callback function that is executed after an object is deleted.
func (a *tableBuilder) AddAfterDeleteInterceptor(interceptor AfterDeleteInterceptor) {
	a.afterDelete = append(a.afterDelete, interceptor)
}

var _ TableExportable = &table{}

// table is the high level object to storage mapper functionality. Persistent
// entities are stored by an unique identifier called `RowID`. The table struct
// does not:
// - enforce uniqueness of the `RowID`
// - enforce prefix uniqueness of keys, i.e. not allowing one key to be a prefix
// of another
// - optimize Gas usage conditions
// The caller must ensure that these things are handled. The table struct is
// private, so that we only have custom tables built on top of table, that do satisfy
// these requirements.
type table struct {
	model       reflect.Type
	prefix      byte
	storeKey    sdk.StoreKey
	afterSet    []AfterSetInterceptor
	afterDelete []AfterDeleteInterceptor
	cdc         codec.Codec
}

// Create persists the given object under the rowID key, returning an
// ErrUniqueConstraint if a value already exists at that key.
//
// Create iterates through the registered callbacks that may add secondary index
// keys.
func (a table) Create(ctx HasKVStore, rowID RowID, obj codec.ProtoMarshaler) error {
	if a.Has(ctx, rowID) {
		return ErrUniqueConstraint
	}

	return a.Set(ctx, rowID, obj)
}

// Update updates the given object under the rowID key. It expects the key to
// exists already and fails with an `ErrNotFound` otherwise. Any caller must
// therefore make sure that this contract is fulfilled. Parameters must not be
// nil.
//
// Update iterates through the registered callbacks that may add or remove
// secondary index keys.
func (a table) Update(ctx HasKVStore, rowID RowID, newValue codec.ProtoMarshaler) error {
	if !a.Has(ctx, rowID) {
		return ErrNotFound
	}

	return a.Set(ctx, rowID, newValue)
}

// Set persists the given object under the rowID key. It does not check if the
// key already exists and overwrites the value if it does.
//
// Set iterates through the registered callbacks that may add secondary index
// keys.
func (a table) Set(ctx HasKVStore, rowID RowID, newValue codec.ProtoMarshaler) error {
	if len(rowID) == 0 {
		return ErrEmptyKey
	}
	if err := assertCorrectType(a.model, newValue); err != nil {
		return err
	}
	if err := assertValid(newValue); err != nil {
		return err
	}

	store := prefix.NewStore(ctx.KVStore(a.storeKey), []byte{a.prefix})

	var oldValue codec.ProtoMarshaler
	if a.Has(ctx, rowID) {
		oldValue = reflect.New(a.model).Interface().(codec.ProtoMarshaler)
		a.GetOne(ctx, rowID, oldValue)
	}

	newValueEncoded, err := a.cdc.Marshal(newValue)
	if err != nil {
		return errors.Wrapf(err, "failed to serialize %T", newValue)
	}

	store.Set(rowID, newValueEncoded)
	for i, itc := range a.afterSet {
		if err := itc(ctx, rowID, newValue, oldValue); err != nil {
			return errors.Wrapf(err, "interceptor %d failed", i)
		}
	}
	return nil
}

func assertValid(obj codec.ProtoMarshaler) error {
	if v, ok := obj.(Validateable); ok {
		if err := v.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}

// Delete removes the object under the rowID key. It expects the key to exists
// already and fails with a `ErrNotFound` otherwise. Any caller must therefore
// make sure that this contract is fulfilled.
//
// Delete iterates through the registered callbacks that remove secondary index
// keys.
func (a table) Delete(ctx HasKVStore, rowID RowID) error {
	store := prefix.NewStore(ctx.KVStore(a.storeKey), []byte{a.prefix})

	var oldValue = reflect.New(a.model).Interface().(codec.ProtoMarshaler)
	if err := a.GetOne(ctx, rowID, oldValue); err != nil {
		return errors.Wrap(err, "load old value")
	}
	store.Delete(rowID)

	for i, itc := range a.afterDelete {
		if err := itc(ctx, rowID, oldValue); err != nil {
			return errors.Wrapf(err, "delete interceptor %d failed", i)
		}
	}
	return nil
}

// Has checks if a key exists. Returns false when the key is empty or nil
// because we don't allow creation of values without a key.
func (a table) Has(ctx HasKVStore, key RowID) bool {
	if len(key) == 0 {
		return false
	}
	store := prefix.NewStore(ctx.KVStore(a.storeKey), []byte{a.prefix})
	it := store.Iterator(PrefixRange(key))
	defer it.Close()
	return it.Valid()
}

// GetOne load the object persisted for the given RowID into the dest parameter.
// If none exists or `rowID==nil` then `ErrNotFound` is returned instead.
// Parameters must not be nil - we don't allow creation of values with empty keys.
func (a table) GetOne(ctx HasKVStore, rowID RowID, dest codec.ProtoMarshaler) error {
	if len(rowID) == 0 {
		return ErrNotFound
	}
	x := NewTypeSafeRowGetter(a.storeKey, a.prefix, a.model, a.cdc)
	return x(ctx, rowID, dest)
}

// PrefixScan returns an Iterator over a domain of keys in ascending order. End is exclusive.
// Start is an MultiKeyIndex key or prefix. It must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// To iterate over entire domain, use PrefixScan(nil, nil)
//
// WARNING: The use of a PrefixScan can be very expensive in terms of Gas. Please make sure you do not expose
// this as an endpoint to the public without further limits.
// Example:
//			it, err := idx.PrefixScan(ctx, start, end)
//			if err !=nil {
//				return err
//			}
//			const defaultLimit = 20
//			it = LimitIterator(it, defaultLimit)
//
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
func (a table) PrefixScan(ctx HasKVStore, start, end RowID) (Iterator, error) {
	if start != nil && end != nil && bytes.Compare(start, end) >= 0 {
		return NewInvalidIterator(), errors.Wrap(ErrArgument, "start must be before end")
	}
	store := prefix.NewStore(ctx.KVStore(a.storeKey), []byte{a.prefix})
	return &typeSafeIterator{
		ctx:       ctx,
		rowGetter: NewTypeSafeRowGetter(a.storeKey, a.prefix, a.model, a.cdc),
		it:        store.Iterator(start, end),
	}, nil
}

// ReversePrefixScan returns an Iterator over a domain of keys in descending order. End is exclusive.
// Start is an MultiKeyIndex key or prefix. It must be less than end, or the Iterator is invalid  and error is returned.
// Iterator must be closed by caller.
// To iterate over entire domain, use PrefixScan(nil, nil)
//
// WARNING: The use of a ReversePrefixScan can be very expensive in terms of Gas. Please make sure you do not expose
// this as an endpoint to the public without further limits. See `LimitIterator`
//
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
func (a table) ReversePrefixScan(ctx HasKVStore, start, end RowID) (Iterator, error) {
	if start != nil && end != nil && bytes.Compare(start, end) >= 0 {
		return NewInvalidIterator(), errors.Wrap(ErrArgument, "start must be before end")
	}
	store := prefix.NewStore(ctx.KVStore(a.storeKey), []byte{a.prefix})
	return &typeSafeIterator{
		ctx:       ctx,
		rowGetter: NewTypeSafeRowGetter(a.storeKey, a.prefix, a.model, a.cdc),
		it:        store.ReverseIterator(start, end),
	}, nil
}

// Export stores all the values in the table in the passed ModelSlicePtr.
func (a table) Export(ctx HasKVStore, dest ModelSlicePtr) (uint64, error) {
	it, err := a.PrefixScan(ctx, nil, nil)
	if err != nil {
		return 0, errors.Wrap(err, "table Export failure when exporting table data")
	}
	_, err = ReadAll(it, dest)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

// Import clears the table and initializes it from the given data interface{}.
// data should be a slice of structs that implement PrimaryKeyed.
func (a table) Import(ctx HasKVStore, data interface{}, _ uint64) error {
	// Clear all data
	store := prefix.NewStore(ctx.KVStore(a.storeKey), []byte{a.prefix})
	it := store.Iterator(nil, nil)
	defer it.Close()
	for ; it.Valid(); it.Next() {
		if err := a.Delete(ctx, it.Key()); err != nil {
			return err
		}
	}

	// Provided data must be a slice
	modelSlice := reflect.ValueOf(data)
	if modelSlice.Kind() != reflect.Slice {
		return errors.Wrap(ErrArgument, "data must be a slice")
	}

	// Import values from slice
	for i := 0; i < modelSlice.Len(); i++ {
		obj, ok := modelSlice.Index(i).Interface().(PrimaryKeyed)
		if !ok {
			return errors.Wrapf(ErrArgument, "unsupported type :%s", reflect.TypeOf(data).Elem().Elem())
		}
		err := a.Create(ctx, PrimaryKey(obj), obj)
		if err != nil {
			return err
		}
	}

	return nil
}

// typeSafeIterator is initialized with a type safe RowGetter only.
type typeSafeIterator struct {
	ctx       HasKVStore
	rowGetter RowGetter
	it        types.Iterator
}

func (i typeSafeIterator) LoadNext(dest codec.ProtoMarshaler) (RowID, error) {
	if !i.it.Valid() {
		return nil, ErrIteratorDone
	}
	rowID := i.it.Key()
	i.it.Next()
	return rowID, i.rowGetter(i.ctx, rowID, dest)
}

func (i typeSafeIterator) Close() error {
	i.it.Close()
	return nil
}
