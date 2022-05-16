package v3_test

import (
	"fmt"
	"testing"
	"time"

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
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	regenorm "github.com/regen-network/regen-ledger/orm"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	v3 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v3"
)

func TestMigrations(t *testing.T) {
	admin1 := sdk.AccAddress("admin1")
	issuer1 := sdk.AccAddress("issuer1")
	issuer2 := sdk.AccAddress("issuer2")
	recipient1 := sdk.AccAddress("recipient1")
	recipient2 := sdk.AccAddress("recipient2")

	ecocreditKey := sdk.NewKVStoreKey("ecocredit")
	tecocreditKey := sdk.NewTransientStoreKey("transient_test")
	encCfg := simapp.MakeTestEncodingConfig()
	paramStore := paramtypes.NewSubspace(encCfg.Marshaler, encCfg.Amino, ecocreditKey, tecocreditKey, ecocredit.ModuleName)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(ecocreditKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tecocreditKey, sdk.StoreTypeTransient, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	store := sdkCtx.KVStore(ecocreditKey)

	paramStore.WithKeyTable(v3.ParamKeyTable())

	ctypes := []*v3.CreditType{
		{
			Name:         "carbon",
			Abbreviation: "C",
			Unit:         "metric ton CO2 equivalent",
			Precision:    6,
		},
		{
			Name:         "biodiversity",
			Abbreviation: "BIO",
			Unit:         "ton",
			Precision:    6,
		},
	}
	paramStore.Set(sdkCtx, v3.KeyCreditTypes, &ctypes)

	classInfoTableBuilder, err := regenorm.NewPrimaryKeyTableBuilder(v3.ClassInfoTablePrefix, ecocreditKey, &v3.ClassInfo{}, encCfg.Marshaler)
	require.NoError(t, err)

	classInfoTable := classInfoTableBuilder.Build()
	batchInfoTableBuilder, err := regenorm.NewPrimaryKeyTableBuilder(v3.BatchInfoTablePrefix, ecocreditKey, &v3.BatchInfo{}, encCfg.Marshaler)
	require.NoError(t, err)

	batchInfoTable := batchInfoTableBuilder.Build()

	creditTypeSeqTableBuilder, err := regenorm.NewPrimaryKeyTableBuilder(v3.CreditTypeSeqTablePrefix, ecocreditKey, &v3.CreditTypeSeq{}, encCfg.Marshaler)
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
	bd1 := formatBatchDenom("C01", 1, &startDate, &endDate)
	bd2 := formatBatchDenom("C01", 2, &startDate, &endDate)
	bd3 := formatBatchDenom("C01", 3, &startDate, &endDate)
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

	err = v3.MigrateState(sdkCtx, ecocreditKey, encCfg.Marshaler, ss, paramStore)
	require.NoError(t, err)

	ctx := sdk.WrapSDKContext(sdkCtx)

	// verify credit class data
	res, err := ss.ClassTable().GetById(ctx, "C01")
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
	res1, err := ss.ProjectTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-001")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "AB-CDE FG1 345")
	require.Equal(t, res1.ClassKey, uint64(1))
	require.NotNil(t, res1.Admin)

	// verify batch migration
	expbd1, err := core.FormatBatchDenom("C01-001", 1, &startDate, &endDate)
	require.NoError(t, err)
	expbd2, err := core.FormatBatchDenom("C01-002", 1, &startDate, &endDate)
	require.NoError(t, err)
	expbd3, err := core.FormatBatchDenom("C01-002", 2, &startDate, &endDate)
	require.NoError(t, err)
	batchRes, err := ss.BatchTable().GetByDenom(ctx, expbd1)
	require.NoError(t, err)
	require.Equal(t, expbd1, batchRes.Denom)
	batchRes, err = ss.BatchTable().GetByDenom(ctx, expbd2)
	require.NoError(t, err)
	require.Equal(t, expbd2, batchRes.Denom)
	batchRes, err = ss.BatchTable().GetByDenom(ctx, expbd3)
	require.NoError(t, err)
	require.Equal(t, expbd3, batchRes.Denom)

	// verify project sequence
	res2, err := ss.ProjectSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res2)
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
	// project C0102 contains two credit batch ==> expected nextBatchId is 3
	res4, err := ss.BatchSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(1))
	require.Equal(t, res4.NextSequence, uint64(2))

	res4, err = ss.BatchSequenceTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(2))
	require.Equal(t, res4.NextSequence, uint64(3))

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

	// verify credit types migration
	carbon, err := ss.CreditTypeTable().Get(ctx, "C")
	require.NoError(t, err)
	require.Equal(t, carbon.Abbreviation, ctypes[0].Abbreviation)
	require.Equal(t, carbon.Name, ctypes[0].Name)
	require.Equal(t, carbon.Precision, ctypes[0].Precision)
	require.Equal(t, carbon.Unit, ctypes[0].Unit)

	bio, err := ss.CreditTypeTable().Get(ctx, "BIO")
	require.NoError(t, err)
	require.Equal(t, bio.Abbreviation, ctypes[1].Abbreviation)
	require.Equal(t, bio.Name, ctypes[1].Name)
	require.Equal(t, bio.Precision, ctypes[1].Precision)
	require.Equal(t, bio.Unit, ctypes[1].Unit)

	// verify old state is deleted
	require.False(t, classInfoTable.Has(sdkCtx, regenorm.RowID("C01")))

	require.False(t, batchInfoTable.Has(sdkCtx, regenorm.RowID(bd1)))

	bz := store.Get(tradableBKey1)
	require.Nil(t, bz)

	bz = store.Get(tradableSKey1)
	require.Nil(t, bz)

}

func formatBatchDenom(classId string, batchSeqNo uint64, startDate *time.Time, endDate *time.Time) string {
	return fmt.Sprintf(
		"%s-%s-%s-%03d",

		// Class ID string
		classId,

		// Start Date as YYYYMMDD
		startDate.Format("20060102"),

		// End Date as YYYYMMDD
		endDate.Format("20060102"),

		// Batch sequence number padded to at least three digits
		batchSeqNo,
	)
}
