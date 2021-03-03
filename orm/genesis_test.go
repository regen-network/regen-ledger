package orm_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/stretchr/testify/require"
)

func TestImportExportTableData(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const prefix = iota
	table := orm.NewTableBuilder(prefix, storeKey, &group.GroupInfo{}, orm.FixLengthIndexKeys(1), cdc).Build()

	ctx := orm.NewMockContext()

	groups := []*group.GroupInfo{
		{
			GroupId:     1,
			Version:     1,
			Admin:       "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du",
			TotalWeight: "1",
			Metadata:    []byte("1"),
		},
		{
			GroupId:     2,
			Version:     2,
			Admin:       "cosmos1qgpqyqszqgpqyqszqgpqyqszqgpqyqszrh8mx2",
			TotalWeight: "2",
			Metadata:    []byte("2"),
		},
	}

	err := orm.ImportTableData(ctx, table, groups, 0)
	require.NoError(t, err)

	for _, g := range groups {
		var loaded group.GroupInfo
		err := table.GetOne(ctx, orm.EncodeSequence(g.GroupId), &loaded)
		require.NoError(t, err)

		require.Equal(t, g, &loaded)
	}

	var exported []*group.GroupInfo
	_, err = orm.ExportTableData(ctx, table, &exported)
	require.NoError(t, err)

	for i, g := range exported {
		require.Equal(t, g, groups[i])
	}

	// require.Equal(t, seq, 0)
}
