package v3_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/types/known/durationpb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	regenorm "github.com/regen-network/regen-ledger/orm"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	v3 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v3"
)

func TestMainnetMigrations(t *testing.T) {
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
	sdkCtx = sdkCtx.WithChainID("regen-1")

	paramStore.WithKeyTable(v3.ParamKeyTable())
	ctypes := []*v3.CreditType{
		{
			Name:         "carbon",
			Abbreviation: "C",
			Unit:         "metric ton CO2 equivalent",
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
	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C01",
		Admin:      "cosmos123a7e9gvgm53zvswc6daq7c85xtzt826warpxl",
		Metadata:   []byte("regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf"),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"cosmos1v2ncquer9r2ytlkxh2djmmsq3e8we6rj88m0lh"},
		NumBatches: 4,
	})
	require.NoError(t, err)

	startDate, endDate, err := v3.ParseBatchDenom("C01-20150101-20151231-003")
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20150101-20151231-003",
		Issuer:          "cosmos1v2ncquer9r2ytlkxh2djmmsq3e8we6rj88m0lh",
		TotalAmount:     "20",
		Metadata:        []byte("cmVnZW46MTN0b1ZnRjg0a1F3U1gxMURkaERhc1l0TUZVMVliNnFRd1F2dHYxcnZIOHBmNUU4VVR5YWpDWC5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate,
		EndDate:         endDate,
		ProjectLocation: "CD-MN",
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20150101-20151231-004",
		Issuer:          "cosmos1v2ncquer9r2ytlkxh2djmmsq3e8we6rj88m0lh",
		TotalAmount:     "3525",
		Metadata:        []byte("cmVnZW46MTN0b1ZnRjg0a1F3U1gxMURkaERhc1l0TUZVMVliNnFRd1F2dHYxcnZIOHBmNUU4VVR5YWpDWC5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate,
		EndDate:         endDate,
		ProjectLocation: "CD-MN",
	})
	require.NoError(t, err)

	startDate1, endDate1, err := v3.ParseBatchDenom("C01-20190101-20191231-001")
	require.NoError(t, err)
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20190101-20191231-001",
		Issuer:          "cosmos1v2ncquer9r2ytlkxh2djmmsq3e8we6rj88m0lh",
		TotalAmount:     "61",
		Metadata:        []byte("cmVnZW46MTN0b1ZnRjg0a1F3U1gxMURkaERhc1l0TUZVMVliNnFRd1F2dHYxcnZIOHBmNUU4VVR5YWpDWC5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate1,
		EndDate:         endDate1,
		ProjectLocation: "KE",
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20190101-20191231-002",
		Issuer:          "cosmos1v2ncquer9r2ytlkxh2djmmsq3e8we6rj88m0lh",
		TotalAmount:     "36",
		Metadata:        []byte("cmVnZW46MTN0b1ZnUjh4TDZOdXlyb3RqYWlrN2JxbWt1V1JuTXZpdDhrYTFmU0JMbmVielA3elVWYk1KMy5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate1,
		EndDate:         endDate1,
		ProjectLocation: "KE",
	})
	require.NoError(t, err)

	err = creditTypeSeqTable.Create(sdkCtx, &v3.CreditTypeSeq{
		Abbreviation: "C",
		SeqNumber:    2,
	})
	require.NoError(t, err)

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)
	ss, err := api.NewStateStore(ormdb)
	require.Nil(t, err)

	basketStore, err := basketapi.NewStateStore(ormdb)
	require.Nil(t, err)

	err = v3.MigrateState(sdkCtx, ecocreditKey, encCfg.Marshaler, ss, basketStore, paramStore)
	require.NoError(t, err)

	ctx := sdk.WrapSDKContext(sdkCtx)

	// verify credit class data
	res, err := ss.ClassTable().GetById(ctx, "C01")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, sdk.AccAddress(res.Admin).String(), "cosmos123a7e9gvgm53zvswc6daq7c85xtzt826warpxl")
	require.Equal(t, res.CreditTypeAbbrev, "C")
	require.Equal(t, res.Metadata, "regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf")
	require.Equal(t, res.Id, "C01")

	// verify class issuers migration
	itr, err := ss.ClassIssuerTable().List(ctx, api.ClassIssuerClassKeyIssuerIndexKey{}.WithClassKey(1))
	require.NoError(t, err)
	require.NotNil(t, itr)

	for itr.Next() {
		val, err := itr.Value()
		require.NoError(t, err)
		require.Equal(t, val.ClassKey, uint64(1))
		require.Equal(t, "cosmos1v2ncquer9r2ytlkxh2djmmsq3e8we6rj88m0lh", sdk.AccAddress(val.Issuer).String())
	}
	itr.Close()

	// verify project migration
	res1, err := ss.ProjectTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-001")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "CD-MN")
	require.Equal(t, res1.ClassKey, uint64(1))
	fmt.Println(res1)

	res1, err = ss.ProjectTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-002")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "KE")
	require.Equal(t, res1.ClassKey, uint64(1))

	// verify batch migration
	expbd1, err := core.FormatBatchDenom("C01-001", 1, startDate, endDate)
	require.NoError(t, err)
	expbd2, err := core.FormatBatchDenom("C01-001", 2, startDate, endDate)
	require.NoError(t, err)

	batchRes, err := ss.BatchTable().GetByDenom(ctx, expbd1)
	require.NoError(t, err)
	require.Equal(t, expbd1, batchRes.Denom)
	batchRes, err = ss.BatchTable().GetByDenom(ctx, expbd2)
	require.NoError(t, err)
	require.Equal(t, expbd2, batchRes.Denom)

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
	require.Equal(t, res3.NextSequence, uint64(2))

	// verify batch sequence table migration
	// project C0101 contains two credit batch ==> expected nextBatchId is 3
	// project C0102 contains two credit batch ==> expected nextBatchId is 3
	res4, err := ss.BatchSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(1))
	require.Equal(t, res4.NextSequence, uint64(3))

	res4, err = ss.BatchSequenceTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(2))
	require.Equal(t, res4.NextSequence, uint64(3))

	// verify credit types migration
	carbon, err := ss.CreditTypeTable().Get(ctx, "C")
	require.NoError(t, err)
	require.Equal(t, carbon.Abbreviation, ctypes[0].Abbreviation)
	require.Equal(t, carbon.Name, ctypes[0].Name)
	require.Equal(t, carbon.Precision, ctypes[0].Precision)
	require.Equal(t, carbon.Unit, ctypes[0].Unit)

	// verify old state is deleted
	require.False(t, classInfoTable.Has(sdkCtx, regenorm.RowID("C01")))

	// verify mainnet manual migrations
	// project location -> reference-id
	// KE    -> "VCS-612" (Kasigao)
	// CD-MN -> "VCS-934" (Mai Ndombe)
	assertProjectReferenceId(t, ctx, ss, "C01-002", "KE", "VCS-612")
	assertProjectReferenceId(t, ctx, ss, "C01-001", "CD-MN", "VCS-934")

	// batch issuance dates
	//  C01-001-20190101-20191231-001  -  "2022-05-06T01:33:13Z"
	//  C01-001-20190101-20191231-002  -  "2022-05-06T01:33:19Z"
	//  C01-002-20190101-20191231-001  -  "2022-05-06T01:33:25Z"
	//  C01-002-20190101-20191231-002  -  "2022-05-06T01:33:31Z"
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20150101-20151231-001", "2022-05-06T01:33:25Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20150101-20151231-002", "2022-05-06T01:33:31Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-002-20190101-20191231-001", "2022-05-06T01:33:13Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-002-20190101-20191231-002", "2022-05-06T01:33:19Z")
}

func TestRedwoodMigrations(t *testing.T) {
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
	sdkCtx = sdkCtx.WithChainID("regen-redwood-1")

	paramStore.WithKeyTable(v3.ParamKeyTable())
	ctypes := []*v3.CreditType{
		{
			Name:         "carbon",
			Abbreviation: "C",
			Unit:         "metric ton CO2 equivalent",
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
	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C01",
		Admin:      "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		Metadata:   []byte("cmVnZW46MTN0b1ZnbzVDQ21Ra1BKRHdMZWd0ZjRVMWVzVzVycnRXcHdxRTZuU2RwMWhhOVc4OFJmdWY1TS5yZGY="),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"cosmos1wjul39t07ds68xasfc4mw8yszwappktmmrhqwn", "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7"},
		NumBatches: 9,
	})
	require.NoError(t, err)

	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C02",
		Admin:      "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		Metadata:   []byte("cmVnZW46YUhSMGNEb3ZMM0psWjJWdUxtNWxkSGR2Y21zdlZrTlRMVk4wWVc1a1lYSmsucmRm="),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7"},
		NumBatches: 2,
	})
	require.NoError(t, err)

	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C03",
		Admin:      "cosmos1mrvlgpmrjn9s7r7ct69euqfgxjazjt2ltafuwf",
		Metadata:   []byte("cmVnZW46MTN0b1ZncU40RVNYZ0xOUVpOejdpR1BOdXdyN0hnaFd4TkNrNFc1SnJzYVgyTndYZk5KVmc5Sy5yZGY="),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"cosmos1v2ncquer9r2ytlkxh2djmmsq3e8we6rj88m0lh"},
		NumBatches: 0,
	})
	require.NoError(t, err)

	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C04",
		Admin:      "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		Metadata:   nil,
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7"},
		NumBatches: 2,
	})
	require.NoError(t, err)

	startDate, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
	endDate, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170101-20180101-005",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "100",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US",
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170101-20180101-006",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "1000",
		Metadata:        []byte("cmVnZW46MTN0b1ZncmQ3dGZ1dmZpYjFiUFhNR29pRGdHaWtkc0FNUFpOV1dRRzVEcHl1Wm0ydEFYRHdlZC5yZGY="),
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US",
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170101-20180101-007",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "1000",
		Metadata:        []byte("cmVnZW46MTN0b1ZnRjg0a1F3U1gxMURkaERhc1l0TUZVMVliNnFRd1F2dHYxcnZIOHBmNUU4VVR5YWpDWC5yZGY="),
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2018-01-01T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20180101-20200101-001",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "10000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "AU-NSW 2453",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2018-09-09T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20180909-20200101-004",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "200",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "FR",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2019-12-31T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20190101-20191231-009",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "61",
		Metadata:        []byte("cmVnZW46MTN0b1ZnRjg0a1F3U1gxMURkaERhc1l0TUZVMVliNnFRd1F2dHYxcnZIOHBmNUU4VVR5YWpDWC5yZGY="),
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "KE",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20190101-20210101-002",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "100000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US",
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20190101-20210101-008",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "1000",
		Metadata:        []byte("cmVnZW46MTN0b1ZnRjg0a1F3U1gxMURkaERhc1l0TUZVMVliNnFRd1F2dHYxcnZIOHBmNUU4VVR5YWpDWC5yZGY="),
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US-FL 98765",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2021-09-09T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2022-01-01T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20210909-20220101-003",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "1000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2020-09-09T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2021-09-09T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C02",
		BatchDenom:      "C02-20200909-20210909-001",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "1234",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "FR",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2021-09-09T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2022-01-01T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C02",
		BatchDenom:      "C02-20210909-20220101-002",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "5678",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2018-02-02T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2019-02-02T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C04",
		BatchDenom:      "C04-20180202-20190202-001",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "30000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "FR",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2019-02-02T00:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "2020-02-02T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C04",
		BatchDenom:      "C04-20190202-20200202-002",
		Issuer:          "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7",
		TotalAmount:     "30000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US",
	})
	require.NoError(t, err)

	err = creditTypeSeqTable.Create(sdkCtx, &v3.CreditTypeSeq{
		Abbreviation: "C",
		SeqNumber:    5,
	})
	require.NoError(t, err)

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)
	ss, err := api.NewStateStore(ormdb)
	require.Nil(t, err)

	basketStore, err := basketapi.NewStateStore(ormdb)
	require.Nil(t, err)

	ctx := sdk.WrapSDKContext(sdkCtx)
	require.NoError(t, basketStore.BasketTable().Save(ctx, &basketapi.Basket{
		BasketDenom:       "eco.uC.rNCT",
		Name:              "rNCT",
		DisableAutoRetire: false,
		CreditTypeAbbrev:  "C",
		DateCriteria: &basketapi.DateCriteria{
			MinStartDate: nil,
			StartDateWindow: &durationpb.Duration{
				Seconds: 315576000,
			},
		},
		Exponent: 6,
	}))

	require.NoError(t, basketStore.BasketTable().Save(ctx, &basketapi.Basket{
		BasketDenom:       "eco.uC.NCT",
		Name:              "NCT",
		DisableAutoRetire: false,
		CreditTypeAbbrev:  "C",
		DateCriteria: &basketapi.DateCriteria{
			MinStartDate: nil,
			StartDateWindow: &durationpb.Duration{
				Seconds: 315576000,
			},
		},
		Exponent: 6,
	}))

	err = v3.MigrateState(sdkCtx, ecocreditKey, encCfg.Marshaler, ss, basketStore, paramStore)
	require.NoError(t, err)

	// verify credit class data
	res, err := ss.ClassTable().GetById(ctx, "C01")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, sdk.AccAddress(res.Admin).String(), "cosmos1df675r9vnf7pdedn4sf26svdsem3ugavhyscr7")
	require.Equal(t, res.CreditTypeAbbrev, "C")
	require.Equal(t, res.Id, "C01")

	// verify project migration
	// class - C01 project-id - C01-001 jurisdiction - US
	// class - C01 project-id - C01-002 jurisdiction - AU-NSW 2453
	// class - C01 project-id - C01-003 jurisdiction - FR
	// class - C01 project-id - C01-004 jurisdiction - KE
	// class - C01 project-id - C01-005 jurisdiction - US-FL 98765
	res1, err := ss.ProjectTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-001")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "US")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-002")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "AU-NSW 2453")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 3)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-003")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "FR")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 4)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-004")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "KE")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 5)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-005")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "US-FL 98765")
	require.Equal(t, res1.ClassKey, uint64(1))

	// class - C02 project-id - C02-001 jurisdiction - US
	// class - C02 project-id - C02-002 jurisdiction - FR
	res1, err = ss.ProjectTable().Get(ctx, 6)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C02-001")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "FR")
	require.Equal(t, res1.ClassKey, uint64(2))

	res1, err = ss.ProjectTable().Get(ctx, 7)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C02-002")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "US")
	require.Equal(t, res1.ClassKey, uint64(2))

	// class - C04 project-id - C04-001 jurisdiction - FR
	// class - C04 project-id - C04-002 jurisdiction - US
	res1, err = ss.ProjectTable().Get(ctx, 8)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C04-001")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "FR")
	require.Equal(t, res1.ClassKey, uint64(4))

	res1, err = ss.ProjectTable().Get(ctx, 9)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C04-002")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "US")
	require.Equal(t, res1.ClassKey, uint64(4))

	// verify project sequence
	res2, err := ss.ProjectSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res2)
	require.Equal(t, res2.ClassKey, uint64(1))
	require.Equal(t, res2.NextSequence, uint64(6))

	res2, err = ss.ProjectSequenceTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res2)
	require.Equal(t, res2.ClassKey, uint64(2))
	require.Equal(t, res2.NextSequence, uint64(3))

	res2, err = ss.ProjectSequenceTable().Get(ctx, 4)
	require.NoError(t, err)
	require.NotNil(t, res2)
	require.Equal(t, res2.ClassKey, uint64(4))
	require.Equal(t, res2.NextSequence, uint64(3))

	// verify class sequence table migration
	res3, err := ss.ClassSequenceTable().Get(ctx, "C")
	require.NoError(t, err)
	require.NotNil(t, res3)
	require.Equal(t, res3.CreditTypeAbbrev, "C")
	require.Equal(t, res3.NextSequence, uint64(5))

	// verify batch sequence table migration
	// project C01-001 contains 5 credit batches ==> expected nextBatchId is 6
	// project C01-002 contains 1 credit batch ==> expected nextBatchId is 2
	// project C01-003 contains 1 credit batch ==> expected nextBatchId is 2
	// project C01-004 contains 1 credit batch ==> expected nextBatchId is 2
	// project C01-005 contains 1 credit batch ==> expected nextBatchId is 2
	res4, err := ss.BatchSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(1))
	require.Equal(t, res4.NextSequence, uint64(6))

	res4, err = ss.BatchSequenceTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(2))
	require.Equal(t, res4.NextSequence, uint64(2))

	res4, err = ss.BatchSequenceTable().Get(ctx, 3)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(3))
	require.Equal(t, res4.NextSequence, uint64(2))

	res4, err = ss.BatchSequenceTable().Get(ctx, 4)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(4))
	require.Equal(t, res4.NextSequence, uint64(2))

	res4, err = ss.BatchSequenceTable().Get(ctx, 5)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(5))
	require.Equal(t, res4.NextSequence, uint64(2))

	// project C02-001 contains 1 credit batch ==> expected nextBatchId is 2
	// project C02-002 contains 1 credit batch ==> expected nextBatchId is 2
	res4, err = ss.BatchSequenceTable().Get(ctx, 6)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(6))
	require.Equal(t, res4.NextSequence, uint64(2))

	res4, err = ss.BatchSequenceTable().Get(ctx, 7)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(7))
	require.Equal(t, res4.NextSequence, uint64(2))

	// project C04-001 contains 1 credit batch ==> expected nextBatchId is 2
	// project C04-002 contains 1 credit batch ==> expected nextBatchId is 2
	res4, err = ss.BatchSequenceTable().Get(ctx, 8)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(8))
	require.Equal(t, res4.NextSequence, uint64(2))

	res4, err = ss.BatchSequenceTable().Get(ctx, 9)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(9))
	require.Equal(t, res4.NextSequence, uint64(2))

	// verify credit types migration
	carbon, err := ss.CreditTypeTable().Get(ctx, "C")
	require.NoError(t, err)
	require.Equal(t, carbon.Abbreviation, ctypes[0].Abbreviation)
	require.Equal(t, carbon.Name, ctypes[0].Name)
	require.Equal(t, carbon.Precision, ctypes[0].Precision)
	require.Equal(t, carbon.Unit, ctypes[0].Unit)

	// verify old state is deleted
	require.False(t, classInfoTable.Has(sdkCtx, regenorm.RowID("C01")))

	// verify mainnet manual migrations
	// project location -> reference-id
	// FR              -> ""
	// US              -> ""
	// AU-NSW 2453     -> ""
	// K               -> ""
	// US-FL 9876      -> ""
	assertProjectReferenceId(t, ctx, ss, "C01-001", "US", "")
	assertProjectReferenceId(t, ctx, ss, "C01-002", "AU-NSW 2453", "")
	assertProjectReferenceId(t, ctx, ss, "C01-003", "FR", "")
	assertProjectReferenceId(t, ctx, ss, "C01-004", "KE", "")
	assertProjectReferenceId(t, ctx, ss, "C01-005", "US-FL 98765", "")
	assertProjectReferenceId(t, ctx, ss, "C02-001", "FR", "")
	assertProjectReferenceId(t, ctx, ss, "C02-002", "US", "")
	assertProjectReferenceId(t, ctx, ss, "C04-001", "FR", "")
	assertProjectReferenceId(t, ctx, ss, "C04-002", "US", "")

	// batch issuance dates
	//  C01-20170101-20180101-005  ->  "2022-03-30T07:46:01Z"
	//  C01-20170101-20180101-006  ->  "2022-03-30T07:51:12Z"
	//  C01-20170101-20180101-007  ->  "2022-03-30T15:04:30Z"
	//  C01-20180101-20200101-001  ->  "2022-02-09T09:10:02Z"
	//  C01-20180909-20200101-004  ->  "2022-03-08T17:25:23Z"
	//  C01-20190101-20191231-009  ->  "2022-05-11T11:35:30Z"
	//  C01-20190101-20210101-002  ->  "2022-02-14T09:07:25Z"
	//  C01-20190101-20210101-008  ->  "2022-04-22T16:32:09Z"
	//  C01-20210909-20220101-003  ->  "2022-03-08T17:18:19Z"
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20170101-20180101-001", "2022-03-30T07:46:01Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20170101-20180101-002", "2022-03-30T07:51:12Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20170101-20180101-003", "2022-03-30T15:04:30Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-002-20180101-20200101-001", "2022-02-09T09:10:02Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-003-20180909-20200101-001", "2022-03-08T17:25:23Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-004-20190101-20191231-001", "2022-05-11T11:35:30Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20190101-20210101-004", "2022-02-14T09:07:25Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-005-20190101-20210101-001", "2022-04-22T16:32:09Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20210909-20220101-005", "2022-03-08T17:18:19Z")

	//  C02-20200909-20210909-001  ->  "2022-03-08T13:00:50Z"
	//  C02-20210909-20220101-002  ->  "2022-03-08T17:17:20Z"
	assertBatchIssuanceDate(t, ctx, ss, "C02-001-20200909-20210909-001", "2022-03-08T13:00:50Z")
	assertBatchIssuanceDate(t, ctx, ss, "C02-002-20210909-20220101-001", "2022-03-08T17:17:20Z")

	//  C04-20180202-20190202-001  ->  "2022-03-28T08:31:45Z"
	//  C04-20190202-20200202-002  ->  "2022-03-28T08:45:14Z"
	assertBatchIssuanceDate(t, ctx, ss, "C04-001-20180202-20190202-001", "2022-03-28T08:31:45Z")
	assertBatchIssuanceDate(t, ctx, ss, "C04-002-20190202-20200202-001", "2022-03-28T08:45:14Z")

	// verify basket curator
	// rNCT   -> regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46
	// NCT    -> regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46

	assertBasketCurator(t, ctx, basketStore, "rNCT", "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46")
	assertBasketCurator(t, ctx, basketStore, "NCT", "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46")
}

func assertBasketCurator(t *testing.T, ctx context.Context, ss basketapi.StateStore, name, curator string) {
	basket, err := ss.BasketTable().GetByName(ctx, name)
	require.NoError(t, err)
	require.Equal(t, string(basket.Curator), curator)
}

func assertBatchIssuanceDate(t *testing.T, ctx context.Context, ss api.StateStore, denom, exp string) {
	res, err := ss.BatchTable().GetByDenom(ctx, denom)
	require.NoError(t, err)
	parsed, _ := time.Parse(time.RFC3339, exp)
	require.Equal(t, res.IssuanceDate.AsTime(), parsed)
}

func assertProjectReferenceId(t *testing.T, ctx context.Context, ss api.StateStore, id, jurisdiction, referenceID string) {
	res, err := ss.ProjectTable().GetById(ctx, id)
	require.NoError(t, err)
	require.Equal(t, res.ReferenceId, referenceID)
	require.Equal(t, res.Jurisdiction, jurisdiction)
}
