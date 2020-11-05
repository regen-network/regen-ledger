package orm

import (
	"testing"

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
			srcIT: mockIter(EncodeSequence(1), &testdata.GroupMetadata{Description: "test"}),
			destSlice: func() ModelSlicePtr {
				x := make([]testdata.GroupMetadata, 1)
				return &x
			},
			expIDs:    []RowID{EncodeSequence(1)},
			expResult: &[]testdata.GroupMetadata{{Description: "test"}},
		},
		"all good with pointer slice": {
			srcIT: mockIter(EncodeSequence(1), &testdata.GroupMetadata{Description: "test"}),
			destSlice: func() ModelSlicePtr {
				x := make([]*testdata.GroupMetadata, 1)
				return &x
			},
			expIDs:    []RowID{EncodeSequence(1)},
			expResult: &[]*testdata.GroupMetadata{{Description: "test"}},
		},
		"dest slice empty": {
			srcIT: mockIter(EncodeSequence(1), &testdata.GroupMetadata{}),
			destSlice: func() ModelSlicePtr {
				x := make([]testdata.GroupMetadata, 0)
				return &x
			},
			expIDs:    []RowID{EncodeSequence(1)},
			expResult: &[]testdata.GroupMetadata{{}},
		},
		"dest pointer with nil value": {
			srcIT: mockIter(EncodeSequence(1), &testdata.GroupMetadata{}),
			destSlice: func() ModelSlicePtr {
				return (*[]testdata.GroupMetadata)(nil)
			},
			expErr: ErrArgument,
		},
		"iterator is nil": {
			srcIT:     nil,
			destSlice: func() ModelSlicePtr { return new([]testdata.GroupMetadata) },
			expErr:    ErrArgument,
		},
		"dest slice is nil": {
			srcIT:     noopIter(),
			destSlice: func() ModelSlicePtr { return nil },
			expErr:    ErrArgument,
		},
		"dest slice is not a pointer": {
			srcIT:     IteratorFunc(nil),
			destSlice: func() ModelSlicePtr { return make([]testdata.GroupMetadata, 1) },
			expErr:    ErrArgument,
		},
		"error on loadNext is returned": {
			srcIT: NewInvalidIterator(),
			destSlice: func() ModelSlicePtr {
				x := make([]testdata.GroupMetadata, 1)
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
	sliceIter := func(s ...string) Iterator {
		var pos int
		return IteratorFunc(func(dest Persistent) (RowID, error) {
			if pos == len(s) {
				return nil, ErrIteratorDone
			}
			v := s[pos]

			*dest.(*persistentString) = persistentString(v)
			pos++
			return []byte(v), nil
		})
	}
	specs := map[string]struct {
		src Iterator
		exp []persistentString
	}{
		"all from range with max > length": {
			src: LimitIterator(sliceIter("a", "b", "c"), 4),
			exp: []persistentString{"a", "b", "c"},
		},
		"up to max": {
			src: LimitIterator(sliceIter("a", "b", "c"), 2),
			exp: []persistentString{"a", "b"},
		},
		"none when max = 0": {
			src: LimitIterator(sliceIter("a", "b", "c"), 0),
			exp: []persistentString{},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			var loaded []persistentString
			_, err := ReadAll(spec.src, &loaded)
			require.NoError(t, err)
			assert.EqualValues(t, spec.exp, loaded)
		})
	}
}

// mockIter amino encodes + decodes value object.
func mockIter(rowID RowID, val Persistent) Iterator {
	b, err := val.Marshal()
	if err != nil {
		panic(err)
	}
	return NewSingleValueIterator(rowID, b)
}

func noopIter() Iterator {
	return IteratorFunc(func(dest Persistent) (RowID, error) {
		return nil, nil
	})
}

type persistentString string

func (p persistentString) Marshal() ([]byte, error) {
	return []byte(p), nil
}

func (p *persistentString) Unmarshal(b []byte) error {
	s := persistentString(string(b))
	p = &s
	return nil
}

func (p persistentString) ValidateBasic() error {
	return nil
}
