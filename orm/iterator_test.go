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

func TestReadAll(t *testing.T) {
	specs := map[string]struct {
		srcIT     orm.Iterator
		destSlice func() orm.ModelSlicePtr
		expErr    *errors.Error
		expIDs    []orm.RowID
		expResult orm.ModelSlicePtr
	}{
		"all good with object slice": {
			srcIT: mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			destSlice: func() orm.ModelSlicePtr {
				x := make([]testdata.GroupInfo, 1)
				return &x
			},
			expIDs:    []orm.RowID{orm.EncodeSequence(1)},
			expResult: &[]testdata.GroupInfo{{Description: "test"}},
		},
		"all good with pointer slice": {
			srcIT: mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			destSlice: func() orm.ModelSlicePtr {
				x := make([]*testdata.GroupInfo, 1)
				return &x
			},
			expIDs:    []orm.RowID{orm.EncodeSequence(1)},
			expResult: &[]*testdata.GroupInfo{{Description: "test"}},
		},
		"dest slice empty": {
			srcIT: mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{}),
			destSlice: func() orm.ModelSlicePtr {
				x := make([]testdata.GroupInfo, 0)
				return &x
			},
			expIDs:    []orm.RowID{orm.EncodeSequence(1)},
			expResult: &[]testdata.GroupInfo{{}},
		},
		"dest pointer with nil value": {
			srcIT: mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{}),
			destSlice: func() orm.ModelSlicePtr {
				return (*[]testdata.GroupInfo)(nil)
			},
			expErr: orm.ErrArgument,
		},
		"iterator is nil": {
			srcIT:     nil,
			destSlice: func() orm.ModelSlicePtr { return new([]testdata.GroupInfo) },
			expErr:    orm.ErrArgument,
		},
		"dest slice is nil": {
			srcIT:     noopIter(),
			destSlice: func() orm.ModelSlicePtr { return nil },
			expErr:    orm.ErrArgument,
		},
		"dest slice is not a pointer": {
			srcIT:     orm.IteratorFunc(nil),
			destSlice: func() orm.ModelSlicePtr { return make([]testdata.GroupInfo, 1) },
			expErr:    orm.ErrArgument,
		},
		"error on loadNext is returned": {
			srcIT: orm.NewInvalidIterator(),
			destSlice: func() orm.ModelSlicePtr {
				x := make([]testdata.GroupInfo, 1)
				return &x
			},
			expErr: orm.ErrIteratorInvalid,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			loaded := spec.destSlice()
			ids, err := orm.ReadAll(spec.srcIT, loaded)
			require.True(t, spec.expErr.Is(err), "expected %s but got %s", spec.expErr, err)
			assert.Equal(t, spec.expIDs, ids)
			if err == nil {
				assert.Equal(t, spec.expResult, loaded)
			}
		})
	}
}

func TestLimitedIterator(t *testing.T) {
	specs := map[string]struct {
		parent      orm.Iterator
		max         int
		expectErr   bool
		expectedErr string
		exp         []testdata.GroupInfo
	}{
		"nil parent": {
			parent:      nil,
			max:         0,
			expectErr:   true,
			expectedErr: "parent iterator must not be nil",
		},
		"negative max": {
			parent:      mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			max:         -1,
			expectErr:   true,
			expectedErr: "quantity must not be negative",
		},
		"all from range with max > length": {
			parent: mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			max:    2,
			exp:    []testdata.GroupInfo{{Description: "test"}},
		},
		"up to max": {
			parent: mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			max:    1,
			exp:    []testdata.GroupInfo{{Description: "test"}},
		},
		"none when max = 0": {
			parent: mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			max:    0,
			exp:    []testdata.GroupInfo{},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			src, err := orm.LimitIterator(spec.parent, spec.max)
			if spec.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), spec.expectedErr)
			} else {
				require.NoError(t, err)
				var loaded []testdata.GroupInfo
				_, err := orm.ReadAll(src, &loaded)
				require.NoError(t, err)
				assert.EqualValues(t, spec.exp, loaded)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	testCases := []struct {
		name          string
		iterator      orm.Iterator
		dest          codec.ProtoMarshaler
		expectErr     bool
		expectedErr   string
		expectedRowID orm.RowID
		expectedDest  codec.ProtoMarshaler
	}{
		{
			name:        "nil iterator",
			iterator:    nil,
			dest:        &testdata.GroupInfo{},
			expectErr:   true,
			expectedErr: "iterator must not be nil",
		},
		{
			name:        "nil dest",
			iterator:    mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			dest:        nil,
			expectErr:   true,
			expectedErr: "destination object must not be nil",
		},
		{
			name:          "all not nil",
			iterator:      mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			dest:          &testdata.GroupInfo{},
			expectErr:     false,
			expectedRowID: orm.EncodeSequence(1),
			expectedDest:  &testdata.GroupInfo{Description: "test"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rowID, err := orm.First(tc.iterator, tc.dest)
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedRowID, rowID)
				require.Equal(t, tc.expectedDest, tc.dest)
			}
		})
	}
}

func TestPaginate(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const (
		testTablePrefix = iota
		testTableSeqPrefix
	)
	tBuilder, err := orm.NewAutoUInt64TableBuilder(testTablePrefix, testTableSeqPrefix, storeKey, &testdata.GroupInfo{}, cdc)
	require.NoError(t, err)
	idx, err := orm.NewIndex(tBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		return []orm.RowID{[]byte(val.(*testdata.GroupInfo).Admin)}, nil
	})
	require.NoError(t, err)
	tb := tBuilder.Build()
	ctx := orm.NewMockContext()

	admin := sdk.AccAddress([]byte("admin-address"))
	g1 := testdata.GroupInfo{
		Description: "my test 1",
		Admin:       admin,
	}
	g2 := testdata.GroupInfo{
		Description: "my test 2",
		Admin:       admin,
	}
	g3 := testdata.GroupInfo{
		Description: "my test 3",
		Admin:       sdk.AccAddress([]byte("other-admin-address")),
	}
	g4 := testdata.GroupInfo{
		Description: "my test 4",
		Admin:       admin,
	}
	g5 := testdata.GroupInfo{
		Description: "my test 5",
		Admin:       sdk.AccAddress([]byte("other-admin-address")),
	}

	for _, g := range []testdata.GroupInfo{g1, g2, g3, g4, g5} {
		_, err := tb.Create(ctx, &g)
		require.NoError(t, err)
	}

	specs := map[string]struct {
		pageReq    *query.PageRequest
		expPageRes *query.PageResponse
		exp        []testdata.GroupInfo
		key        []byte
		expErr     bool
	}{
		"one item": {
			pageReq:    &query.PageRequest{Key: nil, Limit: 1},
			exp:        []testdata.GroupInfo{g1},
			expPageRes: &query.PageResponse{Total: 0, NextKey: orm.EncodeSequence(2)},
			key:        admin,
		},
		"with both key and offset": {
			pageReq: &query.PageRequest{Key: orm.EncodeSequence(2), Offset: 1},
			expErr:  true,
			key:     admin,
		},
		"up to max": {
			pageReq:    &query.PageRequest{Key: nil, Limit: 3, CountTotal: true},
			exp:        []testdata.GroupInfo{g1, g2, g4},
			expPageRes: &query.PageResponse{Total: 3, NextKey: nil},
			key:        admin,
		},
		"no results": {
			pageReq:    &query.PageRequest{Key: nil, Limit: 2, CountTotal: true},
			exp:        []testdata.GroupInfo{},
			expPageRes: &query.PageResponse{Total: 0, NextKey: nil},
			key:        sdk.AccAddress([]byte("no-group-address")),
		},
		"with offset and count total": {
			pageReq:    &query.PageRequest{Key: nil, Offset: 1, Limit: 2, CountTotal: true},
			exp:        []testdata.GroupInfo{g2, g4},
			expPageRes: &query.PageResponse{Total: 3, NextKey: nil},
			key:        admin,
		},
		"nil/default page req (limit = 100 > number of items)": {
			pageReq:    nil,
			exp:        []testdata.GroupInfo{g1, g2, g4},
			expPageRes: &query.PageResponse{Total: 3, NextKey: nil},
			key:        admin,
		},
		"with key and limit < number of elem (count total is ignored in this case)": {
			pageReq:    &query.PageRequest{Key: orm.EncodeSequence(2), Limit: 1, CountTotal: true},
			exp:        []testdata.GroupInfo{g2},
			expPageRes: &query.PageResponse{Total: 0, NextKey: orm.EncodeSequence(4)},
			key:        admin,
		},
		"with key and limit >= number of elem": {
			pageReq:    &query.PageRequest{Key: orm.EncodeSequence(2), Limit: 2},
			exp:        []testdata.GroupInfo{g2, g4},
			expPageRes: &query.PageResponse{Total: 0, NextKey: nil},
			key:        admin,
		},
		"with nothing left to iterate from key": {
			pageReq:    &query.PageRequest{Key: orm.EncodeSequence(5)},
			exp:        []testdata.GroupInfo{},
			expPageRes: &query.PageResponse{Total: 0, NextKey: nil},
			key:        admin,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			var loaded []testdata.GroupInfo

			it, err := idx.GetPaginated(ctx, spec.key, spec.pageReq)
			require.NoError(t, err)

			res, err := orm.Paginate(it, spec.pageReq, &loaded)
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.EqualValues(t, spec.exp, loaded)
				assert.EqualValues(t, spec.expPageRes.Total, res.Total)
				assert.EqualValues(t, spec.expPageRes.NextKey, res.NextKey)
			}

		})
	}

	t.Run("nil iterator", func(t *testing.T) {
		var loaded []testdata.GroupInfo
		res, err := orm.Paginate(nil, &query.PageRequest{}, &loaded)
		require.Error(t, err)
		require.Contains(t, err.Error(), "iterator must not be nil")
		require.Nil(t, res)
	})

	t.Run("non-slice destination", func(t *testing.T) {
		var loaded testdata.GroupInfo
		res, err := orm.Paginate(
			mockIter(orm.EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			&query.PageRequest{},
			&loaded,
		)
		require.Error(t, err)
		require.Contains(t, err.Error(), "destination must point to a slice")
		require.Nil(t, res)
	})
}

// mockIter amino encodes + decodes value object.
func mockIter(rowID orm.RowID, val codec.ProtoMarshaler) orm.Iterator {
	b, err := val.Marshal()
	if err != nil {
		panic(err)
	}
	return orm.NewSingleValueIterator(rowID, b)
}

func noopIter() orm.Iterator {
	return orm.IteratorFunc(func(dest codec.ProtoMarshaler) (orm.RowID, error) {
		return nil, nil
	})
}
