package orm

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/jsonpb"
)

// Model defines the IO structure for table imports and exports
type Model struct {
	Key   []byte          `json:"key" yaml:"key"`
	Value json.RawMessage `json:"value" yaml:"value"`
}

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

// ExportTableData returns a json encoded `[]Model` slice of all the data persisted in the table.
// When the given table implements the `SequenceExportable` interface then it's current value
// is returned as well or otherwise defaults to 0.
func ExportTableData(ctx HasKVStore, t TableExportable) (json.RawMessage, uint64, error) {
	enc := jsonpb.Marshaler{}
	var r []Model
	err := forEachInTable(ctx, t.Table(), func(rowID RowID, obj codec.ProtoMarshaler) error {
		var buf bytes.Buffer
		err := enc.Marshal(&buf, obj)
		if err != nil {
			return errors.Wrap(err, "json encoding")
		}
		r = append(r, Model{Key: rowID, Value: buf.Bytes()})
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	var seqValue uint64
	if st, ok := t.(SequenceExportable); ok {
		seqValue = st.Sequence().CurVal(ctx)
	}
	b, err := json.Marshal(r)
	return b, seqValue, err
}

// ImportTableData initializes a table and attached indexers from the given json encoded `[]Model`s.
// The seqValue is optional and only used with tables that implement the `SequenceExportable` interface.
func ImportTableData(ctx HasKVStore, t TableExportable, src json.RawMessage, seqValue uint64) error {
	dec := json.NewDecoder(bytes.NewReader(src))
	if _, err := dec.Token(); err != nil {
		return errors.Wrap(err, "open bracket")
	}
	table := t.Table()
	if err := clearAllInTable(ctx, table); err != nil {
		return errors.Wrap(err, "clear old entries")
	}
	for dec.More() {
		var m Model
		if err := dec.Decode(&m); err != nil {
			return errors.Wrap(err, "decode")
		}
		if err := putIntoTable(ctx, table, m); err != nil {
			return errors.Wrap(err, "insert from genesis model")
		}
	}
	if _, err := dec.Token(); err != nil {
		return errors.Wrap(err, "closing bracket")
	}
	if st, ok := t.(SequenceExportable); ok {
		if err := st.Sequence().InitVal(ctx, seqValue); err != nil {
			return errors.Wrap(err, "sequence")
		}
	}
	return nil
}

// forEachInTable iterates through all entries in the given table and calls the callback function.
// Aborts on first error.
func forEachInTable(ctx HasKVStore, table Table, f func(RowID, codec.ProtoMarshaler) error) error {
	it, err := table.PrefixScan(ctx, nil, nil)
	if err != nil {
		return errors.Wrap(err, "all rows prefix scan")
	}
	defer it.Close()
	for {
		obj := reflect.New(table.model).Interface().(codec.ProtoMarshaler)
		switch rowID, err := it.LoadNext(obj); {
		case ErrIteratorDone.Is(err):
			return nil
		case err != nil:
			return errors.Wrap(err, "load next")
		default:
			if err := f(rowID, obj); err != nil {
				return err
			}
		}
	}
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

// putIntoTable inserts the model into the table with all save interceptors called
func putIntoTable(ctx HasKVStore, table Table, m Model) error {
	obj := reflect.New(table.model).Interface().(codec.ProtoMarshaler)

	if err := table.cdc.UnmarshalJSON(m.Value, obj); err != nil {
		return errors.Wrapf(err, "can not unmarshal %s into %T", string(m.Value), obj)
	}
	return table.Create(ctx, m.Key, obj)
}
