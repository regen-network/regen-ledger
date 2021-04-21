package orm_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/stretchr/testify/require"
)

func TestImportExportTableData(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const prefix = iota
	table := orm.NewAutoUInt64TableBuilder(prefix, 0x1, storeKey, &orm.GroupInfo{}, cdc).Build()

	ctx := orm.NewMockContext()

	groups := []*orm.GroupInfo{
		{
			GroupId: 1,
			Admin:   sdk.AccAddress([]byte("admin1-address")),
		},
		{
			GroupId: 2,
			Admin:   sdk.AccAddress([]byte("admin2-address")),
		},
	}

	err := orm.ImportTableData(ctx, table, groups, 2)
	require.NoError(t, err)

	for _, g := range groups {
		var loaded orm.GroupInfo
		_, err := table.GetOne(ctx, g.GroupId, &loaded)
		require.NoError(t, err)

		require.Equal(t, g, &loaded)
	}

	var exported []*orm.GroupInfo
	seq, err := orm.ExportTableData(ctx, table, &exported)
	require.NoError(t, err)
	require.Equal(t, seq, uint64(2))

	for i, g := range exported {
		require.Equal(t, g, groups[i])
	}
}
