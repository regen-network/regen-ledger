package orm_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/testutil/testdata"
	grouptypes "github.com/regen-network/regen-ledger/x/group"
)

func TestExportTableData(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const prefix = iota
	table := orm.NewTableBuilder(prefix, storeKey, &testdata.GroupInfo{}, orm.FixLengthIndexKeys(1), cdc).Build()

	ctx := orm.NewMockContext()
	testRecordsNum := 2
	testRecords := make([]testdata.GroupInfo, testRecordsNum)
	for i := 1; i <= testRecordsNum; i++ {
		myAddr := sdk.AccAddress(bytes.Repeat([]byte{byte(i)}, sdk.AddrLen))
		g := testdata.GroupInfo{
			Description: fmt.Sprintf("my test %d", i),
			Admin:       myAddr,
		}
		err := table.Create(ctx, []byte{byte(i)}, &g)
		require.NoError(t, err)
		testRecords[i-1] = g
	}

	jsonModels, _, err := orm.ExportTableData(ctx, table)
	require.NoError(t, err)
	exp := `[
	{
	"key" : "AQ==",
	"value": {"admin":"cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du", "description":"my test 1"}
	},
	{
	"key":"Ag==", 
	"value": {"admin":"cosmos1qgpqyqszqgpqyqszqgpqyqszqgpqyqszrh8mx2", "description":"my test 2"}
	}
]`
	assert.JSONEq(t, exp, string(jsonModels))
}

func TestImportTableData(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const prefix = iota
	table := orm.NewTableBuilder(prefix, storeKey, &testdata.GroupInfo{}, orm.FixLengthIndexKeys(1), cdc).Build()

	ctx := orm.NewMockContext()

	jsonModels := `[
	{
	"key" : "AQ==",
	"value": {"admin":"cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du", "description":"my test 1"}
	},
	{
	"key":"Ag==", 
	"value": {"admin":"cosmos1qgpqyqszqgpqyqszqgpqyqszqgpqyqszrh8mx2", "description":"my test 2"}
	}
]`
	// when
	err := orm.ImportTableData(ctx, table, []byte(jsonModels), 0)
	require.NoError(t, err)

	// then
	for i := 1; i < 3; i++ {
		var loaded testdata.GroupInfo
		err := table.GetOne(ctx, []byte{byte(i)}, &loaded)
		require.NoError(t, err)

		exp := testdata.GroupInfo{
			Description: fmt.Sprintf("my test %d", i),
			Admin:       sdk.AccAddress(bytes.Repeat([]byte{byte(i)}, sdk.AddrLen)),
		}
		require.Equal(t, exp, loaded)
	}

}

func TestImportTableDataAny(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	grouptypes.RegisterTypes(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const prefix = iota
	table := orm.NewTableBuilder(prefix, storeKey, &grouptypes.GroupAccountInfo{}, orm.FixLengthIndexKeys(1), cdc).Build()

	ctx := orm.NewMockContext()

	jsonModels := `[
	{
	"key" : "AQ==",
	"value": {"admin":"cosmos1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnrql8a", "metadata": "AQ==", "decisionPolicy":{"@type":"/regen.group.v1alpha1.ThresholdDecisionPolicy", "threshold":"1", "timeout":"1s"}, "groupId":"1", "groupAccount":"cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du", "version":"1"}
	},
	{
	"key" : "Ag==",
	"value": {"admin":"cosmos1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnrql8a", "metadata": "AQ==", "decisionPolicy":{"@type":"/regen.group.v1alpha1.ThresholdDecisionPolicy", "threshold":"1", "timeout":"2s"}, "groupId":"2", "groupAccount":"cosmos1qgpqyqszqgpqyqszqgpqyqszqgpqyqszrh8mx2", "version":"2"}
	}
]`
	// when
	err := orm.ImportTableData(ctx, table, []byte(jsonModels), 0)
	require.NoError(t, err)

	// then
	for i := 1; i < 3; i++ {
		var loaded grouptypes.GroupAccountInfo
		err := table.GetOne(ctx, []byte{byte(i)}, &loaded)
		require.NoError(t, err)

		exp, err := grouptypes.NewGroupAccountInfo(
			sdk.AccAddress(bytes.Repeat([]byte{byte(i)}, sdk.AddrLen)),
			grouptypes.ID(i),
			sdk.AccAddress(bytes.Repeat([]byte{byte(0)}, sdk.AddrLen)),
			[]byte{1},
			uint64(i),
			&grouptypes.ThresholdDecisionPolicy{
				Threshold: "1",
				Timeout:   proto.Duration{Seconds: int64(i)},
			},
		)
		require.NoError(t, err)
		require.Equal(t, exp, loaded)
	}

}

func TestExportTableDataAny(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	grouptypes.RegisterTypes(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	storeKey := sdk.NewKVStoreKey("test")
	const prefix = iota
	table := orm.NewTableBuilder(prefix, storeKey, &grouptypes.GroupAccountInfo{}, orm.FixLengthIndexKeys(1), cdc).Build()

	ctx := orm.NewMockContext()
	testRecordsNum := 2
	testRecords := make([]grouptypes.GroupAccountInfo, testRecordsNum)
	adminAddr := sdk.AccAddress(bytes.Repeat([]byte{byte(0)}, sdk.AddrLen))
	for i := 1; i <= testRecordsNum; i++ {
		g, err := grouptypes.NewGroupAccountInfo(
			sdk.AccAddress(bytes.Repeat([]byte{byte(i)}, sdk.AddrLen)),
			grouptypes.ID(i),
			adminAddr,
			[]byte{1},
			uint64(i),
			&grouptypes.ThresholdDecisionPolicy{
				Threshold: "1",
				Timeout:   proto.Duration{Seconds: int64(i)},
			},
		)
		require.NoError(t, err)

		err = table.Create(ctx, []byte{byte(i)}, &g)
		require.NoError(t, err)
		testRecords[i-1] = g
	}

	jsonModels, _, err := orm.ExportTableData(ctx, table)
	require.NoError(t, err)
	exp := `[
	{
	"key" : "AQ==",
	"value": {"admin":"cosmos1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnrql8a", "metadata": "AQ==", "decisionPolicy":{"@type":"/regen.group.v1alpha1.ThresholdDecisionPolicy", "threshold":"1", "timeout":"1s"}, "groupId":"1", "groupAccount":"cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du", "version":"1"}
	},
	{
	"key" : "Ag==",
	"value": {"admin":"cosmos1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnrql8a", "metadata": "AQ==", "decisionPolicy":{"@type":"/regen.group.v1alpha1.ThresholdDecisionPolicy", "threshold":"1", "timeout":"2s"}, "groupId":"2", "groupAccount":"cosmos1qgpqyqszqgpqyqszqgpqyqszqgpqyqszrh8mx2", "version":"2"}
	}
]`
	assert.JSONEq(t, exp, string(jsonModels))
}
