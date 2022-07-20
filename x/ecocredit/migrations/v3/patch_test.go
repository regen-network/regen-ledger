package v3_test

import (
	"context"
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
		Admin:      "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm",
		Metadata:   []byte("regen:13toVgo5CCmQkPJDwLegtf4U1esW5rrtWpwqE6nSdp1ha9W88Rfuf5M.rdf"),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn"},
		NumBatches: 8,
	})
	require.NoError(t, err)

	startDate, endDate, err := v3.ParseBatchDenom("C01-20150101-20151231-003")
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20150101-20151231-003",
		Issuer:          "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
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
		Issuer:          "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
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
		Issuer:          "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
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
		Issuer:          "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
		TotalAmount:     "36",
		Metadata:        []byte("cmVnZW46MTN0b1ZnUjh4TDZOdXlyb3RqYWlrN2JxbWt1V1JuTXZpdDhrYTFmU0JMbmVielA3elVWYk1KMy5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate1,
		EndDate:         endDate1,
		ProjectLocation: "KE",
	})
	require.NoError(t, err)

	startDate2, endDate2, err := v3.ParseBatchDenom("C01-20150101-20151231-005")
	require.NoError(t, err)
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20150101-20151231-005",
		Issuer:          "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
		TotalAmount:     "16",
		Metadata:        []byte("cmVnZW46MTN0b1ZnZktFdTdkbVVDc2Y2cGZYS2VXZE5FYUNBbjhyaFlCNDVncEp6b2F6UTFqRXBSeWFwYi5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate2,
		EndDate:         endDate2,
		ProjectLocation: "CD-MN",
	})
	require.NoError(t, err)

	startDate2, endDate2, err = v3.ParseBatchDenom("C01-20190101-20191231-006")
	require.NoError(t, err)
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20190101-20191231-006",
		Issuer:          "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
		TotalAmount:     "31",
		Metadata:        []byte("cmVnZW46MTN0b1ZoQVFDYk1jMkxKbTQ0QVYxZW5hcWkyN21STWtSalBKbVZHZFcxMUM0cWNLdWhybkdQQS5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate2,
		EndDate:         endDate2,
		ProjectLocation: "KE",
	})
	require.NoError(t, err)

	startDate2, endDate2, err = v3.ParseBatchDenom("C01-20150101-20151231-007")
	require.NoError(t, err)
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20150101-20151231-007",
		Issuer:          "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
		TotalAmount:     "51",
		Metadata:        []byte("cmVnZW46MTN0b1ZoYXpZWGcyTHlRN1RYekVrQkRzWWkzd0V5VXppNTZaQjZEZ21iSEd5TGo5Z1V2ZldIbi5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate2,
		EndDate:         endDate2,
		ProjectLocation: "CD-MN",
	})
	require.NoError(t, err)

	startDate2, endDate2, err = v3.ParseBatchDenom("C01-20150101-20151231-008")
	require.NoError(t, err)
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20150101-20151231-008",
		Issuer:          "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn",
		TotalAmount:     "512",
		Metadata:        []byte("cmVnZW46MTN0b1ZnczhYaUU1TWNHeE5oZjRoYjRGNnB4aVJ6UUtYWHdOTmJiQ3MyVlNEbThCV2k5NGRRQi5yZGY="),
		AmountCancelled: "0",
		StartDate:       startDate2,
		EndDate:         endDate2,
		ProjectLocation: "CD-MN",
	})
	require.NoError(t, err)

	err = creditTypeSeqTable.Create(sdkCtx, &v3.CreditTypeSeq{
		Abbreviation: "C",
		SeqNumber:    1,
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
	require.Equal(t, sdk.AccAddress(res.Admin).String(), "regen123a7e9gvgm53zvswc6daq7c85xtzt8263lgasm")
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
		require.Equal(t, "regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn", sdk.AccAddress(val.Issuer).String())
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
	// project C0101 contains five credit batch ==> expected nextBatchId is 6
	// project C0102 contains three credit batch ==> expected nextBatchId is 4
	res4, err := ss.BatchSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(1))
	require.Equal(t, res4.NextSequence, uint64(6))

	res4, err = ss.BatchSequenceTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(2))
	require.Equal(t, res4.NextSequence, uint64(4))

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
	// C01-001-20150101-20151231-003   -  "2022-06-17T00:04:31Z"
	// C01-001-20150101-20151231-004   -  "2022-06-17T00:04:43Z"
	// C01-001-20150101-20151231-005   -  "2022-06-17T00:04:51Z"
	// C01-002-20190101-20191231-003   -  "2022-06-17T00:04:37Z"
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20150101-20151231-001", "2022-05-06T01:33:25Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20150101-20151231-002", "2022-05-06T01:33:31Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-002-20190101-20191231-001", "2022-05-06T01:33:13Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-002-20190101-20191231-002", "2022-05-06T01:33:19Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20150101-20151231-003", "2022-06-17T00:04:31Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20150101-20151231-004", "2022-06-17T00:04:43Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20150101-20151231-005", "2022-06-17T00:04:51Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-002-20190101-20191231-003", "2022-06-17T00:04:37Z")
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
		Admin:      "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		Metadata:   []byte("cmVnZW46MTN0b1ZnbzVDQ21Ra1BKRHdMZWd0ZjRVMWVzVzVycnRXcHdxRTZuU2RwMWhhOVc4OFJmdWY1TS5yZGY="),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers: []string{
			"regen1wjul39t07ds68xasfc4mw8yszwappktmypuuch",
			"regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
			"regen1sv6a7ry6nrls84z0w5lauae4mgmw3kh2mg97ht",
			"regen1ql2tzktyc5clwgzp43g60khdrsecl9n8vfe70u",
			"regen1pcjkd4fwzn36phvnxx2lw26ww8jc0cq9x8k6wj",
			"regen1lhzh0dl7hqq4xgfhps67wrlruejzxcnq5z2sk0",
			"regen1rn2mn8p0j3kqgglf7kpn8eshymgy5sm8w4wmj4",
			"regen15h5eszss2wtavw2x73f66jqp00sh4kewltwppt",
		},
		NumBatches: 38,
	})
	require.NoError(t, err)

	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C02",
		Admin:      "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		Metadata:   []byte("cmVnZW46YUhSMGNEb3ZMM0psWjJWdUxtNWxkSGR2Y21zdlZrTlRMVk4wWVc1a1lYSmsucmRm="),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46"},
		NumBatches: 2,
	})
	require.NoError(t, err)

	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C03",
		Admin:      "regen1mrvlgpmrjn9s7r7ct69euqfgxjazjt2l5lzqcd",
		Metadata:   []byte("cmVnZW46MTN0b1ZncU40RVNYZ0xOUVpOejdpR1BOdXdyN0hnaFd4TkNrNFc1SnJzYVgyTndYZk5KVmc5Sy5yZGY="),
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn"},
		NumBatches: 0,
	})
	require.NoError(t, err)

	err = classInfoTable.Create(sdkCtx, &v3.ClassInfo{
		ClassId:    "C04",
		Admin:      "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		Metadata:   nil,
		CreditType: &v3.CreditType{Name: "carbon", Abbreviation: "C", Precision: 6, Unit: "metric ton CO2 equivalent"},
		Issuers:    []string{"regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46"},
		NumBatches: 2,
	})
	require.NoError(t, err)

	startDate, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
	endDate, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170101-20180101-005",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
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
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
		ProjectLocation: "US",
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2018-05-07T18:07:46Z")
	endDate, _ = time.Parse(time.RFC3339, "2024-06-07T18:07:53Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20180507-20240607-033",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-13T11:20:11Z")
	endDate, _ = time.Parse(time.RFC3339, "2023-06-22T11:20:17Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170613-20230622-013",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170613-20230622-012",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170613-20230622-011",
		ProjectLocation: "US",
		Issuer:          "regen1sv6a7ry6nrls84z0w5lauae4mgmw3kh2mg97ht",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170613-20230622-010",
		ProjectLocation: "US",
		Issuer:          "regen1sv6a7ry6nrls84z0w5lauae4mgmw3kh2mg97ht",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-13T13:15:29Z")
	endDate, _ = time.Parse(time.RFC3339, "2023-06-06T13:15:36Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170613-20230606-027",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-11T16:38:56Z")
	endDate, _ = time.Parse(time.RFC3339, "2023-06-14T16:39:03Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170611-20230614-028",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-07T11:54:54Z")
	endDate, _ = time.Parse(time.RFC3339, "2023-06-08T11:55:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170607-20230608-037",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-07T07:11:55Z")
	endDate, _ = time.Parse(time.RFC3339, "2022-06-22T07:11:59Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170607-20220622-035",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "1854-07-07T06:59:56Z")
	endDate, _ = time.Parse(time.RFC3339, "1987-02-12T07:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-18540707-19870212-034",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "1990-01-03T23:00:00Z")
	endDate, _ = time.Parse(time.RFC3339, "1999-01-03T23:00:00Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-19900103-19990103-032",
		ProjectLocation: "US",
		Issuer:          "regen1ql2tzktyc5clwgzp43g60khdrsecl9n8vfe70u",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2013-06-04T07:15:01Z")
	endDate, _ = time.Parse(time.RFC3339, "2013-06-05T07:15:04Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20130604-20130605-038",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2014-06-04T08:38:25Z")
	endDate, _ = time.Parse(time.RFC3339, "2020-06-03T08:38:29Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20140604-20200603-036",
		ProjectLocation: "US",
		Issuer:          "regen1ql2tzktyc5clwgzp43g60khdrsecl9n8vfe70u",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-05T07:15:48Z")
	endDate, _ = time.Parse(time.RFC3339, "2018-06-01T07:15:51Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170605-20180601-031",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-06T07:03:10Z")
	endDate, _ = time.Parse(time.RFC3339, "2021-06-01T07:03:15Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20210601-030",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "1000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-06T09:18:08Z")
	endDate, _ = time.Parse(time.RFC3339, "2023-06-07T09:18:14Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-014",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "5",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-015",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "5",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-016",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "282",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-017",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "282",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-018",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "282",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-019",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "282",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-06T18:08:55Z")
	endDate, _ = time.Parse(time.RFC3339, "2023-06-07T18:09:01Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-020",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-021",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-022",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-023",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-024",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-026",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230607-029",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	startDate, _ = time.Parse(time.RFC3339, "2017-06-06T17:46:06Z")
	endDate, _ = time.Parse(time.RFC3339, "2023-06-08T17:46:12Z")
	err = batchInfoTable.Create(sdkCtx, &v3.BatchInfo{
		ClassId:         "C01",
		BatchDenom:      "C01-20170606-20230608-025",
		ProjectLocation: "US",
		Issuer:          "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
		TotalAmount:     "30000",
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	})
	require.NoError(t, err)

	err = creditTypeSeqTable.Create(sdkCtx, &v3.CreditTypeSeq{
		Abbreviation: "C",
		SeqNumber:    4,
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

	require.NoError(t, basketStore.BasketTable().Save(ctx, &basketapi.Basket{
		BasketDenom:       "eco.uC.TYLER",
		Name:              "TYLER",
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		DateCriteria:      nil,
		Exponent:          6,
	}))

	err = v3.MigrateState(sdkCtx, ecocreditKey, encCfg.Marshaler, ss, basketStore, paramStore)
	require.NoError(t, err)

	// verify credit class data
	res, err := ss.ClassTable().GetById(ctx, "C01")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, sdk.AccAddress(res.Admin).String(), "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46")
	require.Equal(t, res.CreditTypeAbbrev, "C")
	require.Equal(t, res.Id, "C01")

	// verify project migration
	// class - C01 project-id - C01-001 jurisdiction - US
	// class - C01 project-id - C01-002 jurisdiction - US
	// class - C01 project-id - C01-003 jurisdiction - US
	// class - C01 project-id - C01-004 jurisdiction - AU-NSW 2453
	// class - C01 project-id - C01-005 jurisdiction - FR
	// class - C01 project-id - C01-006 jurisdiction - KE
	// class - C01 project-id - C01-007 jurisdiction - US-FL 98765
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
	require.Equal(t, res1.Jurisdiction, "US")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 3)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-003")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "US")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 4)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-004")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "AU-NSW 2453")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 5)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-005")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "FR")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 6)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-006")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "KE")
	require.Equal(t, res1.ClassKey, uint64(1))

	res1, err = ss.ProjectTable().Get(ctx, 7)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C01-007")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "US-FL 98765")
	require.Equal(t, res1.ClassKey, uint64(1))

	// class - C02 project-id - C02-001 jurisdiction - US
	// class - C02 project-id - C02-002 jurisdiction - FR
	res1, err = ss.ProjectTable().Get(ctx, 8)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C02-001")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "FR")
	require.Equal(t, res1.ClassKey, uint64(2))

	res1, err = ss.ProjectTable().Get(ctx, 9)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C02-002")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "US")
	require.Equal(t, res1.ClassKey, uint64(2))

	// class - C04 project-id - C04-001 jurisdiction - FR
	// class - C04 project-id - C04-002 jurisdiction - US
	res1, err = ss.ProjectTable().Get(ctx, 10)
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, res1.Id, "C04-001")
	require.Equal(t, res1.Metadata, "")
	require.Equal(t, res1.Jurisdiction, "FR")
	require.Equal(t, res1.ClassKey, uint64(4))

	res1, err = ss.ProjectTable().Get(ctx, 11)
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
	require.Equal(t, res2.NextSequence, uint64(8))

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
	// project C01-001 contains 30 credit batches ==> expected nextBatchId is 31
	// project C01-002 contains 2 credit batches  ==> expected nextBatchId is 3
	// project C01-003 contains 2 credit batches  ==> expected nextBatchId is 3
	// project C01-004 contains 1 credit batch    ==> expected nextBatchId is 2
	// project C01-005 contains 1 credit batch    ==> expected nextBatchId is 2
	// project C01-006 contains 1 credit batch    ==> expected nextBatchId is 2
	// project C01-007 contains 1 credit batch    ==> expected nextBatchId is 2
	res4, err := ss.BatchSequenceTable().Get(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(1))
	require.Equal(t, res4.NextSequence, uint64(31))

	res4, err = ss.BatchSequenceTable().Get(ctx, 2)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(2))
	require.Equal(t, res4.NextSequence, uint64(3))

	res4, err = ss.BatchSequenceTable().Get(ctx, 3)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(3))
	require.Equal(t, res4.NextSequence, uint64(3))

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

	// project C02-001 contains 1 credit batch ==> expected nextBatchId is 2
	// project C02-002 contains 1 credit batch ==> expected nextBatchId is 2
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

	// project C04-001 contains 1 credit batch ==> expected nextBatchId is 2
	// project C04-002 contains 1 credit batch ==> expected nextBatchId is 2
	res4, err = ss.BatchSequenceTable().Get(ctx, 10)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(10))
	require.Equal(t, res4.NextSequence, uint64(2))

	res4, err = ss.BatchSequenceTable().Get(ctx, 11)
	require.NoError(t, err)
	require.NotNil(t, res4)
	require.Equal(t, res4.ProjectKey, uint64(11))
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

	// verify redwood manual migrations
	// project location -> reference-id
	// FR              -> ""
	// US              -> ""
	// AU-NSW 2453     -> ""
	// KE               -> ""
	// US-FL 9876      -> ""
	assertProjectReferenceId(t, ctx, ss, "C01-001", "US", "")
	assertProjectReferenceId(t, ctx, ss, "C01-002", "US", "")
	assertProjectReferenceId(t, ctx, ss, "C01-003", "US", "")
	assertProjectReferenceId(t, ctx, ss, "C01-004", "AU-NSW 2453", "")
	assertProjectReferenceId(t, ctx, ss, "C01-005", "FR", "")
	assertProjectReferenceId(t, ctx, ss, "C01-006", "KE", "")
	assertProjectReferenceId(t, ctx, ss, "C01-007", "US-FL 98765", "")
	assertProjectReferenceId(t, ctx, ss, "C02-001", "FR", "")
	assertProjectReferenceId(t, ctx, ss, "C02-002", "US", "")
	assertProjectReferenceId(t, ctx, ss, "C04-001", "FR", "")
	assertProjectReferenceId(t, ctx, ss, "C04-002", "US", "")

	// batch issuance dates
	// C01-18540707-19870212-034  ->  "2022-06-28T18:10:53Z"
	// C01-19900103-19990103-032  ->  "2022-06-28T08:36:25Z"
	// C01-20130604-20130605-038  ->  "2022-06-30T07:16:34Z"
	// C01-20140604-20200603-036  ->  "2022-06-29T08:45:12Z"
	// C01-20170101-20180101-005  ->  "2022-03-30T07:46:01Z"
	// C01-20190101-20191231-009  ->  "2022-05-11T11:35:30Z"
	// C01-20190101-20210101-002  ->  "2022-02-14T09:07:25Z"
	// C01-20190101-20210101-008  ->  "2022-04-22T16:32:09Z"
	// C01-20210909-20220101-003  ->  "2022-03-08T17:18:19Z"
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-18540707-19870212-001", "2022-06-28T18:10:53Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-002-19900103-19990103-001", "2022-06-28T08:36:25Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20130604-20130605-002", "2022-06-30T07:16:34Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-002-20140604-20200603-002", "2022-06-29T08:45:12Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20170101-20180101-003", "2022-03-30T07:46:01Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-006-20190101-20191231-001", "2022-05-11T11:35:30Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20190101-20210101-029", "2022-02-14T09:07:25Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-007-20190101-20210101-001", "2022-04-22T16:32:09Z")
	assertBatchIssuanceDate(t, ctx, ss, "C01-001-20210909-20220101-030", "2022-03-08T17:18:19Z")

	//  C02-20200909-20210909-001  ->  "2022-03-08T13:00:50Z"
	//  C02-20210909-20220101-002  ->  "2022-03-08T17:17:20Z"
	assertBatchIssuanceDate(t, ctx, ss, "C02-001-20200909-20210909-001", "2022-03-08T13:00:50Z")
	assertBatchIssuanceDate(t, ctx, ss, "C02-002-20210909-20220101-001", "2022-03-08T17:17:20Z")

	//  C04-20180202-20190202-001  ->  "2022-03-28T08:31:45Z"
	//  C04-20190202-20200202-002  ->  "2022-03-28T08:45:14Z"
	assertBatchIssuanceDate(t, ctx, ss, "C04-001-20180202-20190202-001", "2022-03-28T08:31:45Z")
	assertBatchIssuanceDate(t, ctx, ss, "C04-002-20190202-20200202-001", "2022-03-28T08:45:14Z")

	// verify basket curator
	// rNCT     -> regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46
	// NCT      -> regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46
	// TYLER    -> regen1yqr0pf38v9j7ah79wmkacau5mdspsc7l0sjeva

	assertBasketCurator(t, ctx, basketStore, "rNCT", "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46")
	assertBasketCurator(t, ctx, basketStore, "NCT", "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46")
	assertBasketCurator(t, ctx, basketStore, "TYLER", "regen1yqr0pf38v9j7ah79wmkacau5mdspsc7l0sjeva")
}

func assertBasketCurator(t *testing.T, ctx context.Context, ss basketapi.StateStore, name, curator string) {
	basket, err := ss.BasketTable().GetByName(ctx, name)
	require.NoError(t, err)
	require.Equal(t, sdk.AccAddress(basket.Curator).String(), curator)
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
