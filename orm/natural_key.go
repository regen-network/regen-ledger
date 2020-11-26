package orm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ Indexable = &NaturalKeyTableBuilder{}

// NewNaturalKeyTableBuilder creates a builder to setup a NaturalKeyTable object.
func NewNaturalKeyTableBuilder(prefixData byte, storeKey sdk.StoreKey, model NaturalKeyed, codec IndexKeyCodec, cdc codec.Marshaler) *NaturalKeyTableBuilder {
	return &NaturalKeyTableBuilder{
		TableBuilder: NewTableBuilder(prefixData, storeKey, model, codec, cdc),
	}
}

type NaturalKeyTableBuilder struct {
	*TableBuilder
}

func (a NaturalKeyTableBuilder) Build() NaturalKeyTable {
	return NaturalKeyTable{table: a.TableBuilder.Build()}

}

// NaturalKeyed defines an object type that is aware of it's immutable natural key.
type NaturalKeyed interface {
	// NaturalKey returns the immutable and serialized natural key of this object. The natural key has to be unique within
	// it's domain so that not two with same value can exist in the same table.
	//
	// The `IndexKeyCodec` used with the `NaturalKeyTable` may add certain constraints to the byte representation as
	// max length = 255 in `Max255DynamicLengthIndexKeyCodec` or a fix length in `FixLengthIndexKeyCodec` for example.
	NaturalKey() []byte
	codec.ProtoMarshaler
}

var _ TableExportable = &NaturalKeyTable{}

// NaturalKeyTable provides simpler object style orm methods without passing database RowIDs.
// Entries are persisted and loaded with a reference to their unique natural key.
type NaturalKeyTable struct {
	table Table
}

// Create persists the given object under their natural key. It checks if the
// key already exists and may return an `ErrUniqueConstraint`.
// Create iterates though the registered callbacks and may add secondary index keys by them.
func (a NaturalKeyTable) Create(ctx HasKVStore, obj NaturalKeyed) error {
	rowID := obj.NaturalKey()
	if a.table.Has(ctx, rowID) {
		return ErrUniqueConstraint
	}
	return a.table.Create(ctx, rowID, obj)
}

// Save updates the given object under the natural key. It expects the key to exists already
// and fails with an `ErrNotFound` otherwise. Any caller must therefore make sure that this contract
// is fulfilled. Parameters must not be nil.
//
// Save iterates though the registered callbacks and may add or remove secondary index keys by them.
func (a NaturalKeyTable) Save(ctx HasKVStore, newValue NaturalKeyed) error {
	return a.table.Save(ctx, newValue.NaturalKey(), newValue)
}

// Delete removes the object. It expects the natural key to exists already
// and fails with a `ErrNotFound` otherwise. Any caller must therefore make sure that this contract
// is fulfilled.
//
// Delete iterates though the registered callbacks and removes secondary index keys by them.
func (a NaturalKeyTable) Delete(ctx HasKVStore, obj NaturalKeyed) error {
	return a.table.Delete(ctx, obj.NaturalKey())
}

// Has checks if a key exists. Panics on nil key.
func (a NaturalKeyTable) Has(ctx HasKVStore, naturalKey RowID) bool {
	return a.table.Has(ctx, naturalKey)
}

// Contains returns true when an object with same type and natural key is persisted in this table.
func (a NaturalKeyTable) Contains(ctx HasKVStore, obj NaturalKeyed) bool {
	if err := assertCorrectType(a.table.model, obj); err != nil {
		return false
	}
	return a.table.Has(ctx, obj.NaturalKey())
}

// GetOne load the object persisted for the given primary Key into the dest parameter.
// If none exists `ErrNotFound` is returned instead. Parameters must not be nil.
func (a NaturalKeyTable) GetOne(ctx HasKVStore, primKey RowID, dest codec.ProtoMarshaler) error {
	return a.table.GetOne(ctx, primKey, dest)
}

// PrefixScan returns an Iterator over a domain of keys in ascending order. End is exclusive.
// Start is an MultiKeyIndex key or prefix. It must be less than end, or the Iterator is invalid and error is returned.
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
func (a NaturalKeyTable) PrefixScan(ctx HasKVStore, start, end []byte) (Iterator, error) {
	return a.table.PrefixScan(ctx, start, end)
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
func (a NaturalKeyTable) ReversePrefixScan(ctx HasKVStore, start, end []byte) (Iterator, error) {
	return a.table.ReversePrefixScan(ctx, start, end)
}

// Table satisfies the TableExportable interface and must not be used otherwise.
func (a NaturalKeyTable) Table() Table {
	return a.table
}
