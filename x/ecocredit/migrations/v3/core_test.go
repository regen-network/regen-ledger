package v3_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regenorm "github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	v3 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v3"
)

func TestMigrations(t *testing.T) {
	admin1 := sdk.AccAddress("admin1")
	issuer1 := sdk.AccAddress("issuer1")
	issuer2 := sdk.AccAddress("issuer2")
	recipient1 := sdk.AccAddress("recipient1")
	recipient2 := sdk.AccAddress("recipient2")

	cdc := simapp.MakeTestEncodingConfig().Marshaler

	ecocreditKey := sdk.NewKVStoreKey("ecocredit")
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(ecocreditKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	store := sdkCtx.KVStore(ecocreditKey)

	classInfoTableBuilder, err := regenorm.NewPrimaryKeyTableBuilder(v3.ClassInfoTablePrefix, ecocreditKey, &v3.ClassInfo{}, cdc)
	require.NoError(t, err)

	classInfoTable := classInfoTableBuilder.Build()
	batchInfoTableBuilder, err := regenorm.NewPrimaryKeyTableBuilder(v3.BatchInfoTablePrefix, ecocreditKey, &v3.BatchInfo{}, cdc)
	require.NoError(t, err)

	batchInfoTable := batchInfoTableBuilder.Build()

	creditTypeSeqTableBuilder, err := regenorm.NewPrimaryKeyTableBuilder(v3.CreditTypeSeqTablePrefix, ecocreditKey, &v3.CreditTypeSeq{}, cdc)
	require.NoError(t, err)

	creditTypeSeqTable := creditTypeSeqTableBuilder.Build()

	err = creditTypeSeqTable.Create(sdkCtx, &v3.CreditTypeSeq{
		Abbreviation: "BIO",
		SeqNumber:    1,
	})
	require.NoError(t, err)

	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C01",
		Admin:      admin1.String(),
		Metadata:   []byte("metadata"),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{issuer1.String(), issuer2.String()},
		NumBatches: 2,
	})
	require.NoError(t, err)

	startDate := sdkCtx.BlockTime().UTC()
	endDate := startDate.AddDate(2, 0, 0)
	bd1, _ := ecocredit.FormatDenom("C01", 1, &startDate, &endDate)
	bd2, _ := ecocredit.FormatDenom("C01", 2, &startDate, &endDate)
	bd3, _ := ecocredit.FormatDenom("C01", 3, &startDate, &endDate)
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      bd1,
		Issuer:          issuer1.String(),
		TotalAmount:     "1000",
		Metadata:        []byte("metadata"),
		AmountCancelled: "100",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "AB-CDE FG1 345",
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      bd2,
		Issuer:          issuer2.String(),
		TotalAmount:     "1000",
		Metadata:        []byte("metadata"),
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "AB-CDE FG1 345",
	})
	require.NoError(t, err)
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      bd3,
		Issuer:          issuer2.String(),
		TotalAmount:     "1000",
		Metadata:        []byte("metadata"),
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "AB-CDE FG1 345",
	})
	require.NoError(t, err)

	err = creditTypeSeqTable.Create(sdkCtx, &v3.CreditTypeSeq{
		Abbreviation: "C",
		SeqNumber:    3,
	})
	require.NoError(t, err)

	tradableBKey1 := v3.TradableBalanceKey(recipient1, v3.BatchDenomT(bd1))
	retiredBKey1 := v3.RetiredBalanceKey(recipient1, v3.BatchDenomT(bd1))
	store.Set(tradableBKey1, []byte("550"))
	store.Set(retiredBKey1, []byte("350"))

	tradableBKey2 := v3.TradableBalanceKey(recipient2, v3.BatchDenomT(bd2))
	retiredBKey2 := v3.RetiredBalanceKey(recipient2, v3.BatchDenomT(bd2))

	store.Set(tradableBKey2, []byte("610"))
	store.Set(retiredBKey2, []byte("390"))

	tradableSKey1 := v3.TradableSupplyKey(v3.BatchDenomT(bd1))
	tradableSKey2 := v3.TradableSupplyKey(v3.BatchDenomT(bd2))
	retiredSKey1 := v3.RetiredSupplyKey(v3.BatchDenomT(bd1))
	retiredSKey2 := v3.RetiredSupplyKey(v3.BatchDenomT(bd2))
	store.Set(tradableSKey1, []byte("550"))
	store.Set(retiredSKey1, []byte("350"))

	store.Set(tradableSKey2, []byte("610"))
	store.Set(retiredSKey2, []byte("390"))

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)
	ss, err := api.NewStateStore(ormdb)
	require.Nil(t, err)

	err = v3.MigrateState(sdkCtx, ecocreditKey, cdc, ss)
	require.NoError(t, err)

	ctx := sdk.WrapSDKContext(sdkCtx)

	// verify credit class data
	res, err := ss.ClassInfoTable().GetById(ctx, "C01")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Admin, admin1.Bytes())
	require.Equal(t, res.CreditTypeAbbrev, "C")
	require.Equal(t, res.Metadata, "metadata")
	require.Equal(t, res.Id, "C01")

	// verify class issuers migration
	itr, err := ss.ClassIssuerTable().List(ctx, api.ClassIssuerClassKeyIssuerIndexKey{}.WithClassKey(1))
	require.NoError(t, err)
	require.NotNil(t, itr)

	issuers := [][]byte{issuer1.Bytes(), issuer2.Bytes()}
	for itr.Next() {
		val, err := itr.Value()
		require.NoError(t, err)
		require.Equal(t, val.ClassKey, uint64(1))
		require.Contains(t, issuers, val.Issuer)
	}
	itr.Close()

	// verify project migration
	res1, err := ss.ProjectInfoTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C0101")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.ProjectJurisdiction, "AB-CDE FG1 345")
	require.Equal(t, res1.ClassKey, uint64(1))
	require.NotNil(t, res1.Admin)

	// verify project sequence
	res2, err := ss.ProjectSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res2.ClassKey, uint64(1))
	require.Equal(t, res2.NextSequence, uint64(3))

	// verify class sequence table migration
	res3, err := ss.ClassSequenceTable().Get(ctx, "C")
	require.NoError(t, err)
	require.NotNil(t, res3)
	require.Equal(t, res3.CreditTypeAbbrev, "C")
	require.Equal(t, res3.NextSequence, uint64(3))

	// verify batch sequence table migration
	// project C0101 contains one credit batch ==> expected nextBatchId is 2
	res4, err := ss.BatchSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(1))
	require.Equal(t, res4.NextSequence, uint64(2))

	// projectC0102 contains two credit batches ==> expected nextBatchId is 3
	res4, err = ss.BatchSequenceTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(2))
	require.Equal(t, res4.NextSequence, uint64(3))

	// verify tradable and retired balance migration
	// recipient1 balance -> tradable: 550 , retired: 350
	// recipient2 balance -> tradable: 610 , retired: 390
	bb, err := ss.BatchBalanceTable().Get(ctx, recipient1.Bytes(), 1)
	require.NoError(t, err)
	require.Equal(t, bb.Tradable, "550")
	require.Equal(t, bb.Retired, "350")

	bb, err = ss.BatchBalanceTable().Get(ctx, recipient2.Bytes(), 2)
	require.NoError(t, err)
	require.Equal(t, bb.Tradable, "610")
	require.Equal(t, bb.Retired, "390")

	// verify tradable and retired supply migrations
	// Supply.b1 -> tradable: 550 , retired: 350, cancelled: 100
	// Supply.b2 -> tradable: 610 , retired: 390, cancelled: 0

	bs, err := ss.BatchSupplyTable().Get(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, bs.TradableAmount, "550")
	require.Equal(t, bs.RetiredAmount, "350")
	require.Equal(t, bs.CancelledAmount, "100")

	bs, err = ss.BatchSupplyTable().Get(ctx, 2)
	require.NoError(t, err)
	require.Equal(t, bs.TradableAmount, "610")
	require.Equal(t, bs.RetiredAmount, "390")
	require.Equal(t, bs.CancelledAmount, "0")

	// verify old state is deleted
	require.False(t, classInfoTable.Has(sdkCtx, regenorm.RowID("C01")))

	require.False(t, batchInfoTable.Has(sdkCtx, regenorm.RowID(bd1)))

	bz := store.Get(tradableBKey1)
	require.Nil(t, bz)

	bz = store.Get(tradableSKey1)
	require.Nil(t, bz)

}
