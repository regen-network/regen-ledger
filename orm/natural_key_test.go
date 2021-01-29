package orm

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/testutil/testdata"
)

func TestNaturalKeyTablePrefixScan(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const (
		testTablePrefix = iota
	)

	tb := NewNaturalKeyTableBuilder(testTablePrefix, storeKey, &testdata.GroupMember{}, Max255DynamicLengthIndexKeyCodec{}, cdc).
		Build()

	ctx := NewMockContext()

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
		expRowIDs  []RowID
		expError   *errors.Error
		method     func(ctx HasKVStore, start, end []byte) (Iterator, error)
	}{
		"exact match with a single result": {
			start:     []byte("group-amember-one"), // == m1.NaturalKey()
			end:       []byte("group-amember-two"), // == m2.NaturalKey()
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1},
			expRowIDs: []RowID{m1.NaturalKey()},
		},
		"one result by prefix": {
			start:     []byte("group-a"),
			end:       []byte("group-amember-two"), // == m2.NaturalKey()
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1},
			expRowIDs: []RowID{m1.NaturalKey()},
		},
		"multi key elements by group prefix": {
			start:     []byte("group-a"),
			end:       []byte("group-b"),
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2},
			expRowIDs: []RowID{m1.NaturalKey(), m2.NaturalKey()},
		},
		"open end query with second group": {
			start:     []byte("group-b"),
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m3},
			expRowIDs: []RowID{m3.NaturalKey()},
		},
		"open end query with all": {
			start:     []byte("group-a"),
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2, m3},
			expRowIDs: []RowID{m1.NaturalKey(), m2.NaturalKey(), m3.NaturalKey()},
		},
		"open start query": {
			start:     nil,
			end:       []byte("group-b"),
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2},
			expRowIDs: []RowID{m1.NaturalKey(), m2.NaturalKey()},
		},
		"open start and end query": {
			start:     nil,
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2, m3},
			expRowIDs: []RowID{m1.NaturalKey(), m2.NaturalKey(), m3.NaturalKey()},
		},
		"all matching prefix": {
			start:     []byte("group"),
			end:       nil,
			method:    tb.PrefixScan,
			expResult: []testdata.GroupMember{m1, m2, m3},
			expRowIDs: []RowID{m1.NaturalKey(), m2.NaturalKey(), m3.NaturalKey()},
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
			expError: ErrArgument,
		},
		"start after end": {
			start:    []byte("b"),
			end:      []byte("a"),
			method:   tb.PrefixScan,
			expError: ErrArgument,
		},
		"reverse: exact match with a single result": {
			start:     []byte("group-amember-one"), // == m1.NaturalKey()
			end:       []byte("group-amember-two"), // == m2.NaturalKey()
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m1},
			expRowIDs: []RowID{m1.NaturalKey()},
		},
		"reverse: one result by prefix": {
			start:     []byte("group-a"),
			end:       []byte("group-amember-two"), // == m2.NaturalKey()
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m1},
			expRowIDs: []RowID{m1.NaturalKey()},
		},
		"reverse: multi key elements by group prefix": {
			start:     []byte("group-a"),
			end:       []byte("group-b"),
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m2, m1},
			expRowIDs: []RowID{m2.NaturalKey(), m1.NaturalKey()},
		},
		"reverse: open end query with second group": {
			start:     []byte("group-b"),
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m3},
			expRowIDs: []RowID{m3.NaturalKey()},
		},
		"reverse: open end query with all": {
			start:     []byte("group-a"),
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m3, m2, m1},
			expRowIDs: []RowID{m3.NaturalKey(), m2.NaturalKey(), m1.NaturalKey()},
		},
		"reverse: open start query": {
			start:     nil,
			end:       []byte("group-b"),
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m2, m1},
			expRowIDs: []RowID{m2.NaturalKey(), m1.NaturalKey()},
		},
		"reverse: open start and end query": {
			start:     nil,
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m3, m2, m1},
			expRowIDs: []RowID{m3.NaturalKey(), m2.NaturalKey(), m1.NaturalKey()},
		},
		"reverse: all matching prefix": {
			start:     []byte("group"),
			end:       nil,
			method:    tb.ReversePrefixScan,
			expResult: []testdata.GroupMember{m3, m2, m1},
			expRowIDs: []RowID{m3.NaturalKey(), m2.NaturalKey(), m1.NaturalKey()},
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
			expError: ErrArgument,
		},
		"reverse: start after end": {
			start:    []byte("b"),
			end:      []byte("a"),
			method:   tb.ReversePrefixScan,
			expError: ErrArgument,
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
			rowIDs, err := ReadAll(it, &loaded)
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

	tb := NewNaturalKeyTableBuilder(testTablePrefix, storeKey, &testdata.GroupMember{}, Max255DynamicLengthIndexKeyCodec{}, cdc).
		Build()

	ctx := NewMockContext()

	myPersistentObj := testdata.GroupMember{
		Group:  []byte("group-a"),
		Member: []byte("member-one"),
		Weight: 1,
	}
	err := tb.Create(ctx, &myPersistentObj)
	require.NoError(t, err)

	specs := map[string]struct {
		src NaturalKeyed
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
		"different natural key": {
			src: &testdata.GroupMember{
				Group:  []byte("another group"),
				Member: []byte("member-one"),
				Weight: 1,
			},
			exp: false,
		},
		"different type, same key": {
			src: mockNaturalKeyed{&myPersistentObj},
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

type mockNaturalKeyed struct {
	*testdata.GroupMember
}
