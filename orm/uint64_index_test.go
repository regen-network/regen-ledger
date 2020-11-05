package orm

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/testutil/testdata"
)

func TestUInt64Index(t *testing.T) {
	storeKey := sdk.NewKVStoreKey("test")

	const anyPrefix = 0x10
	tableBuilder := NewNaturalKeyTableBuilder(anyPrefix, storeKey, &testdata.GroupMember{}, Max255DynamicLengthIndexKeyCodec{})
	myIndex := NewUInt64Index(tableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]uint64, error) {
		return []uint64{uint64(val.(*testdata.GroupMember).Member[0])}, nil
	})
	myTable := tableBuilder.Build()

	ctx := NewMockContext()

	m := testdata.GroupMember{
		Group:  sdk.AccAddress(EncodeSequence(1)),
		Member: sdk.AccAddress([]byte("member-address")),
		Weight: 10,
	}
	err := myTable.Create(ctx, &m)
	require.NoError(t, err)

	indexedKey := uint64('m')

	// Has
	assert.True(t, myIndex.Has(ctx, indexedKey))

	// Get
	it, err := myIndex.Get(ctx, indexedKey)
	require.NoError(t, err)
	var loaded testdata.GroupMember
	rowID, err := it.LoadNext(&loaded)
	require.NoError(t, err)
	require.Equal(t, uint64(1), DecodeSequence(rowID))
	require.Equal(t, m, loaded)

	// PrefixScan match
	it, err = myIndex.PrefixScan(ctx, 0, 255)
	require.NoError(t, err)
	rowID, err = it.LoadNext(&loaded)
	require.NoError(t, err)
	require.Equal(t, uint64(1), DecodeSequence(rowID))
	require.Equal(t, m, loaded)

	// PrefixScan no match
	it, err = myIndex.PrefixScan(ctx, indexedKey+1, 255)
	require.NoError(t, err)
	rowID, err = it.LoadNext(&loaded)
	require.Error(t, ErrIteratorDone, err)

	// ReversePrefixScan match
	it, err = myIndex.ReversePrefixScan(ctx, 0, 255)
	require.NoError(t, err)
	rowID, err = it.LoadNext(&loaded)
	require.NoError(t, err)
	require.Equal(t, uint64(1), DecodeSequence(rowID))
	require.Equal(t, m, loaded)

	// ReversePrefixScan no match
	it, err = myIndex.ReversePrefixScan(ctx, indexedKey+1, 255)
	require.NoError(t, err)
	rowID, err = it.LoadNext(&loaded)
	require.Error(t, ErrIteratorDone, err)
}

func TestUInt64MultiKeyAdapter(t *testing.T) {
	specs := map[string]struct {
		srcFunc UInt64IndexerFunc
		exp     []RowID
		expErr  error
	}{
		"single key": {
			srcFunc: func(value interface{}) ([]uint64, error) {
				return []uint64{1}, nil
			},
			exp: []RowID{{0, 0, 0, 0, 0, 0, 0, 1}},
		},
		"multi key": {
			srcFunc: func(value interface{}) ([]uint64, error) {
				return []uint64{1, 1 << 56}, nil
			},
			exp: []RowID{{0, 0, 0, 0, 0, 0, 0, 1}, {1, 0, 0, 0, 0, 0, 0, 0}},
		},
		"empty key": {
			srcFunc: func(value interface{}) ([]uint64, error) {
				return []uint64{}, nil
			},
			exp: []RowID{},
		},
		"nil key": {
			srcFunc: func(value interface{}) ([]uint64, error) {
				return nil, nil
			},
			exp: []RowID{},
		},
		"error case": {
			srcFunc: func(value interface{}) ([]uint64, error) {
				return nil, errors.New("test")
			},
			expErr: errors.New("test"),
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			fn := UInt64MultiKeyAdapter(spec.srcFunc)
			r, err := fn(nil)
			if spec.expErr != nil {
				require.Equal(t, spec.expErr, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, spec.exp, r)
		})
	}
}
