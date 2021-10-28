package orm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var _ Indexable = &AutoUInt64TableBuilder{}

// NewAutoUInt64TableBuilder creates a builder to setup a AutoUInt64Table object.
func NewAutoUInt64TableBuilder(prefixData byte, prefixSeq byte, storeKey sdk.StoreKey, model codec.ProtoMarshaler, cdc codec.Codec) (*AutoUInt64TableBuilder, error) {
	if prefixData == prefixSeq {
		return nil, ErrUniqueConstraint.Wrap("prefixData and prefixSeq must be unique")
	}

	tableBuilder, err := newTableBuilder(prefixData, storeKey, model, cdc)
	if err != nil {
		return nil, err
	}
	return &AutoUInt64TableBuilder{
		tableBuilder: tableBuilder,
		seq:          NewSequence(storeKey, prefixSeq),
	}, nil
}

type AutoUInt64TableBuilder struct {
	*tableBuilder
	seq Sequence
}

// Build create the AutoUInt64Table object.
func (a AutoUInt64TableBuilder) Build() AutoUInt64Table {
	return AutoUInt64Table{
		table: a.tableBuilder.Build(),
		seq:   a.seq,
	}
}

var _ TableExportable = &AutoUInt64Table{}

// AutoUInt64Table is the table type which an auto incrementing ID.
type AutoUInt64Table struct {
	table table
	seq   Sequence
}

// Create a new persistent object with an auto generated uint64 primary key. The
// key is returned.
//
// Create iterates through the registered callbacks that may add secondary index
// keys.
func (a AutoUInt64Table) Create(ctx HasKVStore, obj codec.ProtoMarshaler) (uint64, error) {
	autoIncID := a.seq.NextVal(ctx)
	err := a.table.Create(ctx, EncodeSequence(autoIncID), obj)
	if err != nil {
		return 0, err
	}
	return autoIncID, nil
}

// Update updates the given object under the rowID key. It expects the key to
// exists already and fails with an `ErrNotFound` otherwise. Any caller must
// therefore make sure that this contract is fulfilled. Parameters must not be
// nil.
//
// Update iterates through the registered callbacks that may add or remove
// secondary index keys.
func (a AutoUInt64Table) Update(ctx HasKVStore, rowID uint64, newValue codec.ProtoMarshaler) error {
	return a.table.Update(ctx, EncodeSequence(rowID), newValue)
}

// Set persists the given object under the rowID key. It does not check if the
// key already exists and overwrites the value if it does.
//
// Set iterates through the registered callbacks that may add secondary index
// keys.
func (a AutoUInt64Table) Set(ctx HasKVStore, rowID uint64, newValue codec.ProtoMarshaler) error {
	return a.table.Set(ctx, EncodeSequence(rowID), newValue)
}

// Delete removes the object under the rowID key. It expects the key to exists
// already and fails with a `ErrNotFound` otherwise. Any caller must therefore
// make sure that this contract is fulfilled.
//
// Delete iterates through the registered callbacks that remove secondary index
// keys.
func (a AutoUInt64Table) Delete(ctx HasKVStore, rowID uint64) error {
	return a.table.Delete(ctx, EncodeSequence(rowID))
}

// Has checks if a rowID exists.
func (a AutoUInt64Table) Has(ctx HasKVStore, rowID uint64) bool {
	return a.table.Has(ctx, EncodeSequence(rowID))
}

// GetOne load the object persisted for the given RowID into the dest parameter.
// If none exists `ErrNotFound` is returned instead. Parameters must not be nil.
func (a AutoUInt64Table) GetOne(ctx HasKVStore, rowID uint64, dest codec.ProtoMarshaler) (RowID, error) {
	rawRowID := EncodeSequence(rowID)
	if err := a.table.GetOne(ctx, rawRowID, dest); err != nil {
		return nil, err
	}
	return rawRowID, nil
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
func (a AutoUInt64Table) PrefixScan(ctx HasKVStore, start, end uint64) (Iterator, error) {
	return a.table.PrefixScan(ctx, EncodeSequence(start), EncodeSequence(end))
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
func (a AutoUInt64Table) ReversePrefixScan(ctx HasKVStore, start uint64, end uint64) (Iterator, error) {
	return a.table.ReversePrefixScan(ctx, EncodeSequence(start), EncodeSequence(end))
}

// Sequence returns the sequence used by this table
func (a AutoUInt64Table) Sequence() Sequence {
	return a.seq
}

// Export stores all the values in the table in the passed ModelSlicePtr and
// returns the current value of the associated sequence.
func (a AutoUInt64Table) Export(ctx HasKVStore, dest ModelSlicePtr) (uint64, error) {
	_, err := a.table.Export(ctx, dest)
	if err != nil {
		return 0, err
	}
	return a.seq.CurVal(ctx), nil
}

// Import clears the table and initializes it from the given data interface{}.
// data should be a slice of structs that implement PrimaryKeyed.
func (a AutoUInt64Table) Import(ctx HasKVStore, data interface{}, seqValue uint64) error {
	if err := a.seq.InitVal(ctx, seqValue); err != nil {
		return errors.Wrap(err, "sequence")
	}
	return a.table.Import(ctx, data, seqValue)
}
