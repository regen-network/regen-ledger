package orm

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
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
	md := testdata.GroupInfo{Description: "foo"}
	bz, err := md.Marshal()
	require.NoError(t, err)
	store.Set(EncodeSequence(1), bz)

	specs := map[string]struct {
		srcRowID     RowID
		srcModelType reflect.Type
		expObj       interface{}
		expErr       *errors.Error
	}{
		"happy path": {
			srcRowID:     EncodeSequence(1),
			srcModelType: reflect.TypeOf(testdata.GroupInfo{}),
			expObj:       testdata.GroupInfo{Description: "foo"},
		},
		"unknown rowID should return ErrNotFound": {
			srcRowID:     EncodeSequence(999),
			srcModelType: reflect.TypeOf(testdata.GroupInfo{}),
			expErr:       ErrNotFound,
		},
		"wrong type should cause ErrType": {
			srcRowID:     EncodeSequence(1),
			srcModelType: reflect.TypeOf(testdata.GroupMember{}),
			expErr:       ErrType,
		},
		"empty rowID not allowed": {
			srcRowID:     []byte{},
			srcModelType: reflect.TypeOf(testdata.GroupInfo{}),
			expErr:       ErrArgument,
		},
		"nil rowID not allowed": {
			srcModelType: reflect.TypeOf(testdata.GroupInfo{}),
			expErr:       ErrArgument,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			interfaceRegistry := types.NewInterfaceRegistry()
			cdc := codec.NewProtoCodec(interfaceRegistry)

			getter := NewTypeSafeRowGetter(storeKey, prefixKey, spec.srcModelType, cdc)
			var loadedObj testdata.GroupInfo

			err := getter(ctx, spec.srcRowID, &loadedObj)
			if spec.expErr != nil {
				require.True(t, spec.expErr.Is(err), err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, spec.expObj, loadedObj)
		})
	}
}
