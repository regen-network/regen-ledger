package v3_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	orm "github.com/regen-network/regen-ledger/orm"
	v3 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v3"
	"github.com/stretchr/testify/require"
)

func TestMigrations(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	ecocreditKey := sdk.NewKVStoreKey("ecocredit")
	ctx := testutil.DefaultContext(ecocreditKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(ecocreditKey)

	classInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(v3.ClassInfoTablePrefix, ecocreditKey, &v3.ClassInfo{}, cdc)
	require.NoError(t, err)

	classInfoTable := classInfoTableBuilder.Build()
	batchInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(v3.BatchInfoTablePrefix, ecocreditKey, &v3.BatchInfo{}, cdc)
	require.NoError(t, err)

	batchInfoTable := batchInfoTableBuilder.Build()

	creditTypeSeqTableBuilder, err := orm.NewPrimaryKeyTableBuilder(v3.CreditTypeSeqTablePrefix, ecocreditKey, &v3.CreditTypeSeq{}, cdc)
	require.NoError(t, err)

	creditTypeSeqTable := creditTypeSeqTableBuilder.Build()


}
