package orm

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// TableExportable
type TableExportable interface {
	// ExportIterator returns an iterator over the values to export
	ExportIterator(HasKVStore) (Iterator, error)

	// ImportSlice clears the table and initialises it with the data in the
	// interface{}. The interface{} should be a slice of values which
	// implement the PrimaryKeyed interface, but this is checked at runtime.
	ImportSlice(HasKVStore, interface{}) error
}

// SequenceExportable
type SequenceExportable interface {
	// Sequence returns the sequence to export
	Sequence() Sequence
}

// ExportTableData iterates over the given table entries and stores them at the passed ModelSlicePtr.
// When the given table implements the `SequenceExportable` interface then it's current value
// is returned as well or otherwise defaults to 0.
func ExportTableData(ctx HasKVStore, t TableExportable, dest ModelSlicePtr) (uint64, error) {
	it, err := t.ExportIterator(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "table ExportIterator failure when exporting table data")
	}
	_, err = ReadAll(it, dest)
	if err != nil {
		return 0, err
	}
	var seqValue uint64
	if st, ok := t.(SequenceExportable); ok {
		seqValue = st.Sequence().CurVal(ctx)
	}
	return seqValue, err
}

// ImportTableData initializes a table and attaches indexers from the given data interface{}.
// data should be a slice of structs that implement PrimaryKeyed (eg []*GroupInfo).
// The seqValue is optional and only used with tables that implement the `SequenceExportable` interface.
func ImportTableData(ctx HasKVStore, t TableExportable, data interface{}, seqValue uint64) error {
	// Import sequence if table is SequenceExportable
	if st, ok := t.(SequenceExportable); ok {
		if err := st.Sequence().InitVal(ctx, seqValue); err != nil {
			return errors.Wrap(err, "sequence")
		}
	}

	// Create table entries
	err := t.ImportSlice(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
