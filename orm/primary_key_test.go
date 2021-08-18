package orm_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/orm/testdata"
)

func TestPrimaryKeyTablePrefixScan(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const (
		testTablePrefix = iota
	)

	tb := orm.NewPrimaryKeyTableBuilder(testTablePrefix, storeKey, &testdata.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc).
		Build()

	ctx := orm.NewMockContext()

	const anyWeight = 1
	m1 := testdata.GroupMember{
		Group:  []byte("group-a"),
		Member: []byte("member-one"),
		Weight: anyWeight,
	}
	m2 := testdata.GroupMember{
		Group:  []byte("group-a"),
		Member: []byte("member-two"),
		Weight: anyWeight,
	}
	m3 := testdata.GroupMember{
		Group:  []byte("group-b"),
		Member: []byte("member-two"),
		Weight: anyWeight,
	}
	for _, g := range []testdata.GroupMember{m1, m2, m3} {
		require.NoError(t, tb.Create(ctx, &g))
	}

	specs := map[string]struct {
		start, end []byte
		expResult  []testdata.GroupMember
		expRowIDs  []orm.RowID
		expError   *errors.Error
		method     func(ctx orm.HasKVStore, start, end []byte) (orm.Iterator, error)
	}{
		"exact match with a single result": {
			start: append(
				orm.AddLengthPrefix([]byte("group-a")),
				orm.AddLengthPrefix([]byte("member-one"))...,
			), // == orm.PrimaryKey(&m1)
			end: append(
				orm.AddLengthPrefix([]byte("group-a")),
				orm.AddLengthPrefix([]byte("member-two"))...,
			), // == orm.PrimaryKey(&m2)
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1)},
		},
		"one result by prefix": {
			start: orm.AddLengthPrefix([]byte("group-a")),
			end: append(
				orm.AddLengthPrefix([]byte("group-a")),
				orm.AddLengthPrefix([]byte("member-two"))...,
			), // == orm.PrimaryKey(&m2)
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1)},
		},
		"multi key elements by group prefix": {
			start:     orm.AddLengthPrefix([]byte("group-a")),
			end:       orm.AddLengthPrefix([]byte("group-b")),
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1), orm.PrimaryKey(&m2)},
		},
		"open end query with second group": {
			start:     orm.AddLengthPrefix([]byte("group-b")),
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m3},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m3)},
		},
		"open end query with all": {
			start:     orm.AddLengthPrefix([]byte("group-a")),
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2, m3},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1), orm.PrimaryKey(&m2), orm.PrimaryKey(&m3)},
		},
		"open start query": {
			start:     nil,
			end:       orm.AddLengthPrefix([]byte("group-b")),
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1), orm.PrimaryKey(&m2)},
		},
		"open start and end query": {
			start:     nil,
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2, m3},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1), orm.PrimaryKey(&m2), orm.PrimaryKey(&m3)},
		},
		"all matching prefix": {
			start:     orm.AddLengthPrefix([]byte("group")), // == LengthPrefix + "group"
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2, m3},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1), orm.PrimaryKey(&m2), orm.PrimaryKey(&m3)},
		},
		"non matching prefix": {
			start:     []byte("nobody"),
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{},
		},
		"start equals end": {
			start:    []byte("any"),
			end:      []byte("any"),
			method:   tb.PrefixScan,
			expError: orm.ErrArgument,
		},
		"start after end": {
			start:    []byte("b"),
			end:      []byte("a"),
			method:   tb.PrefixScan,
			expError: orm.ErrArgument,
		},
		"reverse: exact match with a single result": {
			start: append(
				orm.AddLengthPrefix([]byte("group-a")),
				orm.AddLengthPrefix([]byte("member-one"))...,
			), // == orm.PrimaryKey(&m1)
			end: append(
				orm.AddLengthPrefix([]byte("group-a")),
				orm.AddLengthPrefix([]byte("member-two"))...,
			), // == orm.PrimaryKey(&m2)
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1)},
		},
		"reverse: one result by prefix": {
			start: orm.AddLengthPrefix([]byte("group-a")),
			end: append(
				orm.AddLengthPrefix([]byte("group-a")),
				orm.AddLengthPrefix([]byte("member-two"))...,
			), // == orm.PrimaryKey(&m2)
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m1)},
		},
		"reverse: multi key elements by group prefix": {
			start:     orm.AddLengthPrefix([]byte("group-a")),
			end:       orm.AddLengthPrefix([]byte("group-b")),
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m2, m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m2), orm.PrimaryKey(&m1)},
		},
		"reverse: open end query with second group": {
			start:     orm.AddLengthPrefix([]byte("group-b")),
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m3},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m3)},
		},
		"reverse: open end query with all": {
			start:     orm.AddLengthPrefix([]byte("group-a")),
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m3, m2, m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m3), orm.PrimaryKey(&m2), orm.PrimaryKey(&m1)},
		},
		"reverse: open start query": {
			start:     nil,
			end:       orm.AddLengthPrefix([]byte("group-b")),
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m2, m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m2), orm.PrimaryKey(&m1)},
		},
		"reverse: open start and end query": {
			start:     nil,
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m3, m2, m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m3), orm.PrimaryKey(&m2), orm.PrimaryKey(&m1)},
		},
		"reverse: all matching prefix": {
			start:     orm.AddLengthPrefix([]byte("group")), // == LengthPrefix + "group"
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m3, m2, m1},
			expRowIDs: []orm.RowID{orm.PrimaryKey(&m3), orm.PrimaryKey(&m2), orm.PrimaryKey(&m1)},
		},
		"reverse: non matching prefix": {
			start:     []byte("nobody"),
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{},
		},
		"reverse: start equals end": {
			start:    []byte("any"),
			end:      []byte("any"),
			method:   tb.ReversePrefixScan,
			expError: orm.ErrArgument,
		},
		"reverse: start after end": {
			start:    []byte("b"),
			end:      []byte("a"),
			method:   tb.ReversePrefixScan,
			expError: orm.ErrArgument,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			it, err := spec.method(ctx, spec.start, spec.end)
			require.True(t, spec.expError.Is(err), "expected #+v but got #+v", spec.expError, err)
			if spec.expError != nil {
				return
			}
			var loaded []testdata.GroupMember
			rowIDs, err := orm.ReadAll(it, &loaded)
			require.NoError(t, err)
			assert.Equal(t, spec.expResult, loaded)
			assert.Equal(t, spec.expRowIDs, rowIDs)
		})
	}
}

func TestContains(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const testTablePrefix = iota

	tb := orm.NewPrimaryKeyTableBuilder(testTablePrefix, storeKey, &testdata.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc).
		Build()

	ctx := orm.NewMockContext()

	myPersistentObj := testdata.GroupMember{
		Group:  []byte("group-a"),
		Member: []byte("member-one"),
		Weight: 1,
	}
	err := tb.Create(ctx, &myPersistentObj)
	require.NoError(t, err)

	specs := map[string]struct {
		src orm.PrimaryKeyed
		exp bool
	}{

		"same object": {src: &myPersistentObj, exp: true},
		"clone": {
			src: &testdata.GroupMember{
				Group:  []byte("group-a"),
				Member: []byte("member-one"),
				Weight: 1,
			},
			exp: true,
		},
		"different primary key": {
			src: &testdata.GroupMember{
				Group:  []byte("another group"),
				Member: []byte("member-one"),
				Weight: 1,
			},
			exp: false,
		},
		"different type, same key": {
			src: mockPrimaryKeyed{&myPersistentObj},
			exp: false,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			got := tb.Contains(ctx, spec.src)
			assert.Equal(t, spec.exp, got)
		})
	}
}

func TestAddLengthPrefix(t *testing.T) {
	tcs := []struct {
		name     string
		in       []byte
		expected []byte
	}{
		{"empty", []byte{}, []byte{0}},
		{"nil", nil, []byte{0}},
		{"some data", []byte{0, 1, 100, 200}, []byte{4, 0, 1, 100, 200}},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			out := orm.AddLengthPrefix(tc.in)
			require.Equal(t, tc.expected, out)
		})
	}

	require.Panics(t, func() {
		orm.AddLengthPrefix(make([]byte, 300))
	})
}

func TestNullTerminatedBytes(t *testing.T) {
	tcs := []struct {
		name     string
		in       string
		expected []byte
	}{
		{"empty", "", []byte{0}},
		{"some data", "abc", []byte{0x61, 0x62, 0x63, 0}},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			out := orm.NullTerminatedBytes(tc.in)
			require.Equal(t, tc.expected, out)
		})
	}
}

type mockPrimaryKeyed struct {
	*testdata.GroupMember
}
