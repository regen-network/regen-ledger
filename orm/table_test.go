package orm_test

import (
	"fmt"
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

func TestBuilder(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")

	testCases := []struct {
		name        string
		storeKey    sdk.StoreKey
		model       codec.ProtoMarshaler
		idxKeyCodec orm.IndexKeyCodec
		expectErr   bool
		expectedErr string
	}{
		{
			name:        "nil storeKey",
			storeKey:    nil,
			model:       &testdata.GroupInfo{},
			idxKeyCodec: orm.Max255DynamicLengthIndexKeyCodec{},
			expectErr:   true,
			expectedErr: "StoreKey must not be nil",
		},
		{
			name:        "nil model",
			storeKey:    storeKey,
			model:       nil,
			idxKeyCodec: orm.Max255DynamicLengthIndexKeyCodec{},
			expectErr:   true,
			expectedErr: "Model must not be nil",
		},
		{
			name:        "nil idxKeyCodec",
			storeKey:    storeKey,
			model:       &testdata.GroupInfo{},
			idxKeyCodec: nil,
			expectErr:   true,
			expectedErr: "IndexKeyCodec must not be nil",
		},
		{
			name:        "all not nil",
			storeKey:    storeKey,
			model:       &testdata.GroupInfo{},
			idxKeyCodec: orm.Max255DynamicLengthIndexKeyCodec{},
			expectErr:   false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			builder, err := orm.TestTableBuilder(0x1, tc.storeKey, tc.model, tc.idxKeyCodec, cdc)
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
			} else {
				require.NoError(t, err)
				require.NotNil(t, builder)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	specs := map[string]struct {
		rowID  orm.RowID
		src    codec.ProtoMarshaler
		expErr *errors.Error
	}{
		"empty rowID": {
			rowID: []byte{},
			src: &testdata.GroupInfo{
				Description: "my group",
				Admin:       sdk.AccAddress([]byte("my-admin-address")),
			},
			expErr: orm.ErrEmptyKey,
		},
		"happy path": {
			rowID: []byte("my-id"),
			src: &testdata.GroupInfo{
				Description: "my group",
				Admin:       sdk.AccAddress([]byte("my-admin-address")),
			},
		},
		"wrong type": {
			rowID: []byte("my-id"),
			src: &testdata.GroupMember{
				Group:  sdk.AccAddress(orm.EncodeSequence(1)),
				Member: sdk.AccAddress([]byte("member-address")),
				Weight: 10,
			},
			expErr: orm.ErrType,
		},
		"model validation fails": {
			rowID:  []byte("my-id"),
			src:    &testdata.GroupInfo{Description: "invalid"},
			expErr: testdata.ErrTest,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			interfaceRegistry := types.NewInterfaceRegistry()
			cdc := codec.NewProtoCodec(interfaceRegistry)

			storeKey := sdk.NewKVStoreKey("test")
			const anyPrefix = 0x10
			tableBuilder, err := orm.TestTableBuilder(anyPrefix, storeKey, &testdata.GroupInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
			require.NoError(t, err)
			myTable := tableBuilder.Build()

			ctx := orm.NewMockContext()
			err = myTable.Create(ctx, spec.rowID, spec.src)

			require.True(t, spec.expErr.Is(err), err)
			shouldExists := spec.expErr == nil
			assert.Equal(t, shouldExists, myTable.Has(ctx, spec.rowID), fmt.Sprintf("expected %v", shouldExists))

			// then
			var loaded testdata.GroupInfo
			err = myTable.GetOne(ctx, spec.rowID, &loaded)
			if spec.expErr != nil {
				require.True(t, orm.ErrNotFound.Is(err))
				return
			}
			require.NoError(t, err)
			assert.Equal(t, spec.src, &loaded)
		})
	}

}
func TestUpdate(t *testing.T) {
	specs := map[string]struct {
		src    codec.ProtoMarshaler
		expErr *errors.Error
	}{
		"happy path": {
			src: &testdata.GroupInfo{
				Description: "my group",
				Admin:       sdk.AccAddress([]byte("my-admin-address")),
			},
		},
		"wrong type": {
			src: &testdata.GroupMember{
				Group:  sdk.AccAddress(orm.EncodeSequence(1)),
				Member: sdk.AccAddress([]byte("member-address")),
				Weight: 9999,
			},
			expErr: orm.ErrType,
		},
		"model validation fails": {
			src:    &testdata.GroupInfo{Description: "invalid"},
			expErr: testdata.ErrTest,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			interfaceRegistry := types.NewInterfaceRegistry()
			cdc := codec.NewProtoCodec(interfaceRegistry)

			storeKey := sdk.NewKVStoreKey("test")
			const anyPrefix = 0x10
			tableBuilder, err := orm.TestTableBuilder(anyPrefix, storeKey, &testdata.GroupInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
			require.NoError(t, err)
			myTable := tableBuilder.Build()

			initValue := testdata.GroupInfo{
				Description: "my old group description",
				Admin:       sdk.AccAddress([]byte("my-old-admin-address")),
			}
			ctx := orm.NewMockContext()
			err = myTable.Create(ctx, []byte("my-id"), &initValue)
			require.NoError(t, err)

			// when
			err = myTable.Save(ctx, []byte("my-id"), spec.src)
			require.True(t, spec.expErr.Is(err), "got ", err)

			// then
			var loaded testdata.GroupInfo
			require.NoError(t, myTable.GetOne(ctx, []byte("my-id"), &loaded))
			if spec.expErr == nil {
				assert.Equal(t, spec.src, &loaded)
			} else {
				assert.Equal(t, initValue, loaded)
			}
		})
	}

}
