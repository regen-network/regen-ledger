package orm

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/testutil/testdata"
)

func TestTypeSafeRowGetter(t *testing.T) {
	storeKey := sdk.NewKVStoreKey("test")
	ctx := NewMockContext()
	const prefixKey = 0x2
	store := prefix.NewStore(ctx.KVStore(storeKey), []byte{prefixKey})
	md := testdata.GroupMetadata{Description: "foo"}
	bz, err := md.Marshal()
	require.NoError(t, err)
	store.Set(EncodeSequence(1), bz)

	specs := map[string]struct {
		srcRowID     RowID
		srcModelType reflect.Type
		mutateDest   func(*testdata.GroupMetadata) Persistent
		expObj       interface{}
		expErr       *errors.Error
	}{
		"happy path": {
			srcRowID:     EncodeSequence(1),
			srcModelType: reflect.TypeOf(testdata.GroupMetadata{}),
			expObj:       testdata.GroupMetadata{Description: "foo"},
		},
		"unknown rowID should return ErrNotFound": {
			srcRowID:     EncodeSequence(999),
			srcModelType: reflect.TypeOf(testdata.GroupMetadata{}),
			expErr:       ErrNotFound,
		},
		"wrong type should cause ErrType": {
			srcRowID:     EncodeSequence(1),
			srcModelType: reflect.TypeOf(testdata.GroupMember{}),
			expErr:       ErrType,
		},
		"empty rowID not allowed": {
			srcRowID:     []byte{},
			srcModelType: reflect.TypeOf(testdata.GroupMetadata{}),
			expErr:       ErrArgument,
		},
		"nil rowID not allowed": {
			srcModelType: reflect.TypeOf(testdata.GroupMetadata{}),
			expErr:       ErrArgument,
		},
		"target not a pointer": {
			srcRowID:     EncodeSequence(1),
			srcModelType: reflect.TypeOf(alwaysPanicPersistenceTarget{}),
			mutateDest: func(m *testdata.GroupMetadata) Persistent {
				return alwaysPanicPersistenceTarget{}
			},
			expErr: ErrType,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			getter := NewTypeSafeRowGetter(storeKey, prefixKey, spec.srcModelType)
			var loadedObj testdata.GroupMetadata
			var dest Persistent
			if spec.mutateDest != nil {
				dest = spec.mutateDest(&loadedObj)
			} else {
				dest = &loadedObj
			}
			err := getter(ctx, spec.srcRowID, dest)
			if spec.expErr != nil {
				require.True(t, spec.expErr.Is(err), err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, spec.expObj, loadedObj)
		})
	}
}

type alwaysPanicPersistenceTarget struct{}

func (n alwaysPanicPersistenceTarget) Marshal() ([]byte, error) {
	panic("implement me")
}

func (n alwaysPanicPersistenceTarget) Unmarshal([]byte) error {
	panic("implement me")
}

func (n alwaysPanicPersistenceTarget) ValidateBasic() error {
	panic("implement me")
}
