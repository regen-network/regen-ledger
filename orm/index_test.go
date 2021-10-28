package orm_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/orm/testdata"
)

var _, _ orm.Indexable = &nilStoreKeyBuilder{}, &nilRowGetterBuilder{}

type nilStoreKeyBuilder struct{}

func (b *nilStoreKeyBuilder) StoreKey() sdk.StoreKey { return nil }
func (b *nilStoreKeyBuilder) RowGetter() orm.RowGetter {
	return func(a orm.HasKVStore, b orm.RowID, c codec.ProtoMarshaler) error { return nil }
}
func (b *nilStoreKeyBuilder) AddAfterSetInterceptor(orm.AfterSetInterceptor)       {}
func (b *nilStoreKeyBuilder) AddAfterDeleteInterceptor(orm.AfterDeleteInterceptor) {}

type nilRowGetterBuilder struct{}

func (b *nilRowGetterBuilder) StoreKey() sdk.StoreKey {
	return sdk.NewKVStoreKey("test")
}
func (b *nilRowGetterBuilder) RowGetter() orm.RowGetter {
	return nil
}
func (b *nilRowGetterBuilder) AddAfterSetInterceptor(orm.AfterSetInterceptor)       {}
func (b *nilRowGetterBuilder) AddAfterDeleteInterceptor(orm.AfterDeleteInterceptor) {}

func TestNewIndex(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	storeKey := sdk.NewKVStoreKey("test")
	const (
		testTablePrefix = iota
		testTableSeqPrefix
	)
	tBuilder, err := orm.NewAutoUInt64TableBuilder(testTablePrefix, testTableSeqPrefix, storeKey, &testdata.GroupInfo{}, cdc)
	require.NoError(t, err)
	indexer := func(val interface{}) ([]interface{}, error) {
		return []interface{}{[]byte(val.(*testdata.GroupInfo).Admin)}, nil
	}

	testCases := []struct {
		name        string
		builder     orm.Indexable
		expectErr   bool
		expectedErr string
		indexKey    interface{}
	}{
		{
			name:        "nil storeKey",
			builder:     &nilStoreKeyBuilder{},
			expectErr:   true,
			expectedErr: "StoreKey must not be nil",
			indexKey:    []byte{},
		},
		{
			name:        "nil rowGetter",
			builder:     &nilRowGetterBuilder{},
			expectErr:   true,
			expectedErr: "RowGetter must not be nil",
			indexKey:    []byte{},
		},
		{
			name:      "all not nil",
			builder:   tBuilder,
			expectErr: false,
			indexKey:  []byte{},
		},
		{
			name:      "index key type not allowed",
			builder:   tBuilder,
			expectErr: true,
			indexKey:  1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index, err := orm.NewIndex(tc.builder, 0x1, indexer, tc.indexKey)
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, index)
			}
		})
	}
}

func TestIndexPrefixScan(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	storeKey := sdk.NewKVStoreKey("test")
	const (
		testTablePrefix = iota
		testTableSeqPrefix
	)
	tBuilder, err := orm.NewAutoUInt64TableBuilder(testTablePrefix, testTableSeqPrefix, storeKey, &testdata.GroupInfo{}, cdc)
	require.NoError(t, err)
	idx, err := orm.NewIndex(tBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]interface{}, error) {
		i := []interface{}{val.(*testdata.GroupInfo).Admin.Bytes()}
		return i, nil
	}, testdata.GroupInfo{}.Admin.Bytes())
	require.NoError(t, err)
	strIdx, err := orm.NewIndex(tBuilder, GroupByDescriptionIndexPrefix, func(val interface{}) ([]interface{}, error) {
		i := []interface{}{val.(*testdata.GroupInfo).Description}
		return i, nil
	}, testdata.GroupInfo{}.Description)
	require.NoError(t, err)

	tb := tBuilder.Build()
	ctx := orm.NewMockContext()

	g1 := testdata.GroupInfo{
		Description: "my test 1",
		Admin:       sdk.AccAddress([]byte("admin-address-a")),
	}
	g2 := testdata.GroupInfo{
		Description: "my test 2",
		Admin:       sdk.AccAddress([]byte("admin-address-b")),
	}
	g3 := testdata.GroupInfo{
		Description: "my test 3",
		Admin:       sdk.AccAddress([]byte("admin-address-b")),
	}
	for _, g := range []testdata.GroupInfo{g1, g2, g3} {
		_, err := tb.Create(ctx, &g)
		require.NoError(t, err)
	}

	specs := map[string]struct {
		start, end interface{}
		expResult  []testdata.GroupInfo
		expRowIDs  []orm.RowID
		expError   *errors.Error
		method     func(ctx orm.HasKVStore, start, end interface{}) (orm.Iterator, error)
	}{
		"exact match with a single result": {
			start:     []byte("admin-address-a"),
			end:       []byte("admin-address-b"),
			method:    idx.PrefixScan,
			expResult: []testdata.GroupInfo{g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1)},
		},
		"one result by prefix": {
			start:     []byte("admin-address"),
			end:       []byte("admin-address-b"),
			method:    idx.PrefixScan,
			expResult: []testdata.GroupInfo{g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1)},
		},
		"multi key elements by exact match": {
			start:     []byte("admin-address-b"),
			end:       []byte("admin-address-c"),
			method:    idx.PrefixScan,
			expResult: []testdata.GroupInfo{g2, g3},
			expRowIDs: []orm.RowID{orm.EncodeSequence(2), orm.EncodeSequence(3)},
		},
		"open end query": {
			start:     []byte("admin-address-b"),
			end:       nil,
			method:    idx.PrefixScan,
			expResult: []testdata.GroupInfo{g2, g3},
			expRowIDs: []orm.RowID{orm.EncodeSequence(2), orm.EncodeSequence(3)},
		},
		"open start query": {
			start:     nil,
			end:       []byte("admin-address-b"),
			method:    idx.PrefixScan,
			expResult: []testdata.GroupInfo{g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1)},
		},
		"open start and end query": {
			start:     nil,
			end:       nil,
			method:    idx.PrefixScan,
			expResult: []testdata.GroupInfo{g1, g2, g3},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1), orm.EncodeSequence(2), orm.EncodeSequence(3)},
		},
		"all matching prefix": {
			start:     []byte("admin"),
			end:       nil,
			method:    idx.PrefixScan,
			expResult: []testdata.GroupInfo{g1, g2, g3},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1), orm.EncodeSequence(2), orm.EncodeSequence(3)},
		},
		"non matching prefix": {
			start:     []byte("admin-address-c"),
			end:       nil,
			method:    idx.PrefixScan,
			expResult: []testdata.GroupInfo{},
		},
		"start equals end": {
			start:    []byte("any"),
			end:      []byte("any"),
			method:   idx.PrefixScan,
			expError: orm.ErrArgument,
		},
		"start after end": {
			start:    []byte("b"),
			end:      []byte("a"),
			method:   idx.PrefixScan,
			expError: orm.ErrArgument,
		},
		"reverse: exact match with a single result": {
			start:     []byte("admin-address-a"),
			end:       []byte("admin-address-b"),
			method:    idx.ReversePrefixScan,
			expResult: []testdata.GroupInfo{g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1)},
		},
		"reverse: one result by prefix": {
			start:     []byte("admin-address"),
			end:       []byte("admin-address-b"),
			method:    idx.ReversePrefixScan,
			expResult: []testdata.GroupInfo{g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1)},
		},
		"reverse: multi key elements by exact match": {
			start:     []byte("admin-address-b"),
			end:       []byte("admin-address-c"),
			method:    idx.ReversePrefixScan,
			expResult: []testdata.GroupInfo{g3, g2},
			expRowIDs: []orm.RowID{orm.EncodeSequence(3), orm.EncodeSequence(2)},
		},
		"reverse: open end query": {
			start:     []byte("admin-address-b"),
			end:       nil,
			method:    idx.ReversePrefixScan,
			expResult: []testdata.GroupInfo{g3, g2},
			expRowIDs: []orm.RowID{orm.EncodeSequence(3), orm.EncodeSequence(2)},
		},
		"reverse: open start query": {
			start:     nil,
			end:       []byte("admin-address-b"),
			method:    idx.ReversePrefixScan,
			expResult: []testdata.GroupInfo{g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1)},
		},
		"reverse: open start and end query": {
			start:     nil,
			end:       nil,
			method:    idx.ReversePrefixScan,
			expResult: []testdata.GroupInfo{g3, g2, g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(3), orm.EncodeSequence(2), orm.EncodeSequence(1)},
		},
		"reverse: all matching prefix": {
			start:     []byte("admin"),
			end:       nil,
			method:    idx.ReversePrefixScan,
			expResult: []testdata.GroupInfo{g3, g2, g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(3), orm.EncodeSequence(2), orm.EncodeSequence(1)},
		},
		"reverse: non matching prefix": {
			start:     []byte("admin-address-c"),
			end:       nil,
			method:    idx.ReversePrefixScan,
			expResult: []testdata.GroupInfo{},
		},
		"reverse: start equals end": {
			start:    []byte("any"),
			end:      []byte("any"),
			method:   idx.ReversePrefixScan,
			expError: orm.ErrArgument,
		},
		"reverse: start after end": {
			start:    []byte("b"),
			end:      []byte("a"),
			method:   idx.ReversePrefixScan,
			expError: orm.ErrArgument,
		},
		"exact match with a single result using string based index": {
			start:     "my test 1",
			end:       "my test 2",
			method:    strIdx.PrefixScan,
			expResult: []testdata.GroupInfo{g1},
			expRowIDs: []orm.RowID{orm.EncodeSequence(1)},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			it, err := spec.method(ctx, spec.start, spec.end)
			require.True(t, spec.expError.Is(err), "expected #+v but got #+v", spec.expError, err)
			if spec.expError != nil {
				return
			}
			var loaded []testdata.GroupInfo
			rowIDs, err := orm.ReadAll(it, &loaded)
			require.NoError(t, err)
			assert.Equal(t, spec.expResult, loaded)
			assert.Equal(t, spec.expRowIDs, rowIDs)
		})
	}
}

func TestUniqueIndex(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")

	tableBuilder, err := orm.NewPrimaryKeyTableBuilder(GroupMemberTablePrefix, storeKey, &testdata.GroupMember{}, cdc)
	require.NoError(t, err)
	uniqueIdx, err := orm.NewUniqueIndex(tableBuilder, 0x10, func(val interface{}) (interface{}, error) {
		return []byte{val.(*testdata.GroupMember).Member[0]}, nil
	}, []byte{})
	require.NoError(t, err)
	myTable := tableBuilder.Build()

	ctx := orm.NewMockContext()

	m := testdata.GroupMember{
		Group:  sdk.AccAddress(orm.EncodeSequence(1)),
		Member: sdk.AccAddress([]byte("member-address")),
		Weight: 10,
	}
	err = myTable.Create(ctx, &m)
	require.NoError(t, err)

	indexedKey := []byte{'m'}

	// Has
	exists, err := uniqueIdx.Has(ctx, indexedKey)
	require.NoError(t, err)
	assert.True(t, exists)

	// Get
	it, err := uniqueIdx.Get(ctx, indexedKey)
	require.NoError(t, err)
	var loaded testdata.GroupMember
	rowID, err := it.LoadNext(&loaded)
	require.NoError(t, err)
	require.Equal(t, orm.RowID(orm.PrimaryKey(&m)), rowID)
	require.Equal(t, m, loaded)

	// GetPaginated
	cases := map[string]struct {
		pageReq *query.PageRequest
		expErr  bool
	}{
		"nil key": {
			pageReq: &query.PageRequest{Key: nil},
			expErr:  false,
		},
		"after indexed key": {
			pageReq: &query.PageRequest{Key: indexedKey},
			expErr:  true,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			it, err := uniqueIdx.GetPaginated(ctx, indexedKey, tc.pageReq)
			require.NoError(t, err)
			rowID, err := it.LoadNext(&loaded)
			if tc.expErr { // iterator done
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, orm.RowID(orm.PrimaryKey(&m)), rowID)
				require.Equal(t, m, loaded)
			}
		})
	}

	// PrefixScan match
	it, err = uniqueIdx.PrefixScan(ctx, indexedKey, nil)
	require.NoError(t, err)
	rowID, err = it.LoadNext(&loaded)
	require.NoError(t, err)
	require.Equal(t, orm.RowID(orm.PrimaryKey(&m)), rowID)
	require.Equal(t, m, loaded)

	// PrefixScan no match
	it, err = uniqueIdx.PrefixScan(ctx, []byte{byte('n')}, nil)
	require.NoError(t, err)
	_, err = it.LoadNext(&loaded)
	require.Error(t, orm.ErrIteratorDone, err)

	// ReversePrefixScan match
	it, err = uniqueIdx.ReversePrefixScan(ctx, indexedKey, nil)
	require.NoError(t, err)
	rowID, err = it.LoadNext(&loaded)
	require.NoError(t, err)
	require.Equal(t, orm.RowID(orm.PrimaryKey(&m)), rowID)
	require.Equal(t, m, loaded)

	// ReversePrefixScan no match
	it, err = uniqueIdx.ReversePrefixScan(ctx, []byte{byte('l')}, nil)
	require.NoError(t, err)
	_, err = it.LoadNext(&loaded)
	require.Error(t, orm.ErrIteratorDone, err)
	// create with same index key should fail
	new := testdata.GroupMember{
		Group:  sdk.AccAddress(orm.EncodeSequence(1)),
		Member: sdk.AccAddress([]byte("my-other")),
		Weight: 10,
	}
	err = myTable.Create(ctx, &new)
	require.Error(t, orm.ErrUniqueConstraint, err)

	// and when delete
	err = myTable.Delete(ctx, &m)
	require.NoError(t, err)

	// then no persistent element
	exists, err = uniqueIdx.Has(ctx, indexedKey)
	require.NoError(t, err)
	assert.False(t, exists)
}

func TestPrefixRange(t *testing.T) {
	cases := map[string]struct {
		src      []byte
		expStart []byte
		expEnd   []byte
		expPanic bool
	}{
		"normal":                 {src: []byte{1, 3, 4}, expStart: []byte{1, 3, 4}, expEnd: []byte{1, 3, 5}},
		"normal short":           {src: []byte{79}, expStart: []byte{79}, expEnd: []byte{80}},
		"empty case":             {src: []byte{}},
		"roll-over example 1":    {src: []byte{17, 28, 255}, expStart: []byte{17, 28, 255}, expEnd: []byte{17, 29, 0}},
		"roll-over example 2":    {src: []byte{15, 42, 255, 255}, expStart: []byte{15, 42, 255, 255}, expEnd: []byte{15, 43, 0, 0}},
		"pathological roll-over": {src: []byte{255, 255, 255, 255}, expStart: []byte{255, 255, 255, 255}},
		"nil prohibited":         {expPanic: true},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if tc.expPanic {
				require.Panics(t, func() {
					orm.PrefixRange(tc.src)
				})
				return
			}
			start, end := orm.PrefixRange(tc.src)
			assert.Equal(t, tc.expStart, start)
			assert.Equal(t, tc.expEnd, end)
		})
	}
}
