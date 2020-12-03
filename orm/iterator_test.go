package orm

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/testutil/testdata"
)

func TestReadAll(t *testing.T) {
	specs := map[string]struct {
		srcIT     Iterator
		destSlice func() ModelSlicePtr
		expErr    *errors.Error
		expIDs    []RowID
		expResult ModelSlicePtr
	}{
		"all good with object slice": {
			srcIT: mockIter(EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			destSlice: func() ModelSlicePtr {
				x := make([]testdata.GroupInfo, 1)
				return &x
			},
			expIDs:    []RowID{EncodeSequence(1)},
			expResult: &[]testdata.GroupInfo{{Description: "test"}},
		},
		"all good with pointer slice": {
			srcIT: mockIter(EncodeSequence(1), &testdata.GroupInfo{Description: "test"}),
			destSlice: func() ModelSlicePtr {
				x := make([]*testdata.GroupInfo, 1)
				return &x
			},
			expIDs:    []RowID{EncodeSequence(1)},
			expResult: &[]*testdata.GroupInfo{{Description: "test"}},
		},
		"dest slice empty": {
			srcIT: mockIter(EncodeSequence(1), &testdata.GroupInfo{}),
			destSlice: func() ModelSlicePtr {
				x := make([]testdata.GroupInfo, 0)
				return &x
			},
			expIDs:    []RowID{EncodeSequence(1)},
			expResult: &[]testdata.GroupInfo{{}},
		},
		"dest pointer with nil value": {
			srcIT: mockIter(EncodeSequence(1), &testdata.GroupInfo{}),
			destSlice: func() ModelSlicePtr {
				return (*[]testdata.GroupInfo)(nil)
			},
			expErr: ErrArgument,
		},
		"iterator is nil": {
			srcIT:     nil,
			destSlice: func() ModelSlicePtr { return new([]testdata.GroupInfo) },
			expErr:    ErrArgument,
		},
		"dest slice is nil": {
			srcIT:     noopIter(),
			destSlice: func() ModelSlicePtr { return nil },
			expErr:    ErrArgument,
		},
		"dest slice is not a pointer": {
			srcIT:     IteratorFunc(nil),
			destSlice: func() ModelSlicePtr { return make([]testdata.GroupInfo, 1) },
			expErr:    ErrArgument,
		},
		"error on loadNext is returned": {
			srcIT: NewInvalidIterator(),
			destSlice: func() ModelSlicePtr {
				x := make([]testdata.GroupInfo, 1)
				return &x
			},
			expErr: ErrIteratorInvalid,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			loaded := spec.destSlice()
			ids, err := ReadAll(spec.srcIT, loaded)
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
		src Iterator
		exp []testdata.GroupInfo
	}{
		"all from range with max > length": {
			src: LimitIterator(mockIter(EncodeSequence(1), &testdata.GroupInfo{Description: "test"}), 2),
			exp: []testdata.GroupInfo{testdata.GroupInfo{Description: "test"}},
		},
		"up to max": {
			src: LimitIterator(mockIter(EncodeSequence(1), &testdata.GroupInfo{Description: "test"}), 1),
			exp: []testdata.GroupInfo{testdata.GroupInfo{Description: "test"}},
		},
		"none when max = 0": {
			src: LimitIterator(mockIter(EncodeSequence(1), &testdata.GroupInfo{Description: "test"}), 0),
			exp: []testdata.GroupInfo{},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			var loaded []testdata.GroupInfo
			_, err := ReadAll(spec.src, &loaded)
			require.NoError(t, err)
			assert.EqualValues(t, spec.exp, loaded)
		})
	}
}

// mockIter amino encodes + decodes value object.
func mockIter(rowID RowID, val codec.ProtoMarshaler) Iterator {
	b, err := val.Marshal()
	if err != nil {
		panic(err)
	}
	return NewSingleValueIterator(rowID, b)
}

func noopIter() Iterator {
	return IteratorFunc(func(dest codec.ProtoMarshaler) (RowID, error) {
		return nil, nil
	})
}
