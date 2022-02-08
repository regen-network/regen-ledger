package core

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"testing"
	"time"
)

var testAddr = sdk.AccAddress("foo")

func TestSetupTable(t *testing.T) {
	db := setupStore(t)
	ctx := context.Background()
	store, err := ecocreditv1beta1.NewStateStore(db)
	require.NoError(t, err)
	classInfoName := "CLS1"
	classRowId, err := store.ClassInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.ClassInfo{
		Name:       classInfoName,
		Admin:      testAddr,
		Metadata:   []byte("hello"),
		CreditType: "carbon",
	})
	_, err = store.ClassInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "CLS2",
		Admin:      testAddr,
		Metadata:   []byte("hello"),
		CreditType: "biodiversity",
	})
	require.NoError(t, err)
	err = store.CreditTypeStore().Insert(ctx, &ecocreditv1beta1.CreditType{
		Abbreviation: "C01",
		Name:         "carbon",
		Unit:         "chungus",
		Precision:    6,
	})
	require.NoError(t, err)
	projectRowId, err := store.ProjectInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "BigProject",
		ClassId:         classRowId,
		ProjectLocation: "90210",
		Metadata:        []byte("very cool project"),
	})
	require.NoError(t, err)
	start, end := time.Now(), time.Now()
	denom, err := ecocredit.FormatDenom(classInfoName, 1, &start, &end)
	require.NoError(t, err)
	batchRowID, err := store.BatchInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  projectRowId,
		BatchDenom: denom,
		Metadata:   []byte("cool batch"),
		StartDate:  timestamppb.New(start),
		EndDate:    timestamppb.New(end),
	})
	require.NoError(t, err)
	err = store.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
		ClassId: classInfoName,
		Issuer:  testAddr,
	})
	require.NoError(t, err)
	err = store.BatchSupplyStore().Insert(ctx, &ecocreditv1beta1.BatchSupply{
		BatchId:         batchRowID,
		TradableAmount:  "100000000000000000",
		RetiredAmount:   "100000000000000000",
		CancelledAmount: "5",
	})
	require.NoError(t, err)
	err = store.BatchBalanceStore().Insert(ctx, &ecocreditv1beta1.BatchBalance{
		Address:  testAddr,
		BatchId:  batchRowID,
		Tradable: "100000000000000000",
		Retired:  "100000000000000000",
	})
	require.NoError(t, err)
	target := ormjson.NewRawMessageTarget()
	err = db.ExportJSON(ctx, target)
	require.NoError(t, err)
	bz, err := target.JSON()
	require.NoError(t, err)
	err = os.WriteFile("testdata.json", bz, 0777)
	require.NoError(t, err)
}

func TestSomething(t *testing.T) {
	db := setupStore(t)
	ctx := context.Background()
	bz, err := os.ReadFile("testdata.json")
	require.NoError(t, err)
	jzn, err := ormjson.NewRawMessageSource(bz)
	require.NoError(t, err)
	err = db.ImportJSON(ctx, jzn)
	require.NoError(t, err)

	stores, err := ecocreditv1beta1.NewStateStore(db)
	require.NoError(t, err)

	it, err := stores.ClassInfoStore().List(ctx, &ecocreditv1beta1.ClassInfoPrimaryKey{})
	require.NoError(t, err)
	for it.Next() {
		val, err := it.Value()
		require.NoError(t, err)
		fmt.Println(val)
	}
}
