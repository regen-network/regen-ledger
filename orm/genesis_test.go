package orm

import (
	"bytes"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/testutil/testdata"
)

func TestExportTableData(t *testing.T) {
	storeKey := sdk.NewKVStoreKey("test")
	const prefix = iota
	table := NewTableBuilder(prefix, storeKey, &testdata.GroupMetadata{}, FixLengthIndexKeys(1)).Build()

	ctx := NewMockContext()
	testRecordsNum := 2
	testRecords := make([]testdata.GroupMetadata, testRecordsNum)
	for i := 1; i <= testRecordsNum; i++ {
		myAddr := sdk.AccAddress(bytes.Repeat([]byte{byte(i)}, sdk.AddrLen))
		g := testdata.GroupMetadata{
			Description: fmt.Sprintf("my test %d", i),
			Admin:       myAddr,
		}
		err := table.Create(ctx, []byte{byte(i)}, &g)
		require.NoError(t, err)
		testRecords[i-1] = g
	}

	jsonModels, _, err := ExportTableData(ctx, table)
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
	storeKey := sdk.NewKVStoreKey("test")
	const prefix = iota
	table := NewTableBuilder(prefix, storeKey, &testdata.GroupMetadata{}, FixLengthIndexKeys(1)).Build()

	ctx := NewMockContext()

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
	err := ImportTableData(ctx, table, []byte(jsonModels), 0)
	require.NoError(t, err)

	// then
	for i := 1; i < 3; i++ {
		var loaded testdata.GroupMetadata
		err := table.GetOne(ctx, []byte{byte(i)}, &loaded)
		require.NoError(t, err)

		exp := testdata.GroupMetadata{
			Description: fmt.Sprintf("my test %d", i),
			Admin:       sdk.AccAddress(bytes.Repeat([]byte{byte(i)}, sdk.AddrLen)),
		}
		require.Equal(t, exp, loaded)
	}

}
