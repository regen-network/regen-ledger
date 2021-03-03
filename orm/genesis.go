package orm

import (
	"reflect"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// TableExportable
type TableExportable interface {
	// Table returns the table to export
	Table() Table
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
	it, err := t.Table().PrefixScan(ctx, nil, nil)
	if err != nil {
		return 0, errors.Wrap(err, "all rows prefix scan")
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

// ImportTableData initializes a table and and attached indexers from the given data interface{}.
// The seqValue is optional and only used with tables that implement the `SequenceExportable` interface.
// func ImportTableData(ctx HasKVStore, t TableExportable, seqValue uint64, createTableData func(HasKVStore, Table) error) error {
func ImportTableData(ctx HasKVStore, t TableExportable, data interface{}, seqValue uint64) error {
	table := t.Table()
	if err := clearAllInTable(ctx, table); err != nil {
		return errors.Wrap(err, "clear old entries")
	}

	if st, ok := t.(SequenceExportable); ok {
		if err := st.Sequence().InitVal(ctx, seqValue); err != nil {
			return errors.Wrap(err, "sequence")
		}
	}

	// Provided data must be a slice
	modelSlice := reflect.ValueOf(data)
	if modelSlice.Kind() != reflect.Slice {
		return errors.Wrap(ErrArgument, "data must be a slice")
	}

	// Create table entries
	for i := 0; i < modelSlice.Len(); i++ {
		obj, ok := modelSlice.Index(i).Interface().(NaturalKeyed)
		if !ok {
			return errors.Wrapf(ErrArgument, "unsupported type :%s", reflect.TypeOf(data).Elem().Elem())
		}
		err := table.Create(ctx, obj.NaturalKey(), obj)
		if err != nil {
			return err
		}
	}

	return nil
}

// clearAllInTable deletes all entries in a table with delete interceptors called
func clearAllInTable(ctx HasKVStore, table Table) error {
	store := prefix.NewStore(ctx.KVStore(table.storeKey), []byte{table.prefix})
	it := store.Iterator(nil, nil)
	defer it.Close()
	for ; it.Valid(); it.Next() {
		if err := table.Delete(ctx, it.Key()); err != nil {
			return err
		}
	}
	return nil
}
