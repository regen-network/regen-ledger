package v4_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	v4 "github.com/regen-network/regen-ledger/x/ecocredit/v3/migrations/v4"
)

func TestMainnetMigrations(t *testing.T) {
	sdkCtx, baseStore, basketStore := setup(t)
	sdkCtx = sdkCtx.WithChainID("regen-1")

	// issuer is the same for all credit batches
	issuer := sdk.MustAccAddressFromBech32("regen1v2ncquer9r2ytlkxh2djmmsq3e8we6rjc9snfn")

	// generic project key (we only need to test unchanged after migration)
	projectKey := uint64(1)

	// generic timestamp (we only need to test unchanged after migration)
	timestamp := timestamppb.Now()

	batches := []*baseapi.Batch{
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-001
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-001-20150101-20151231-001",
			Metadata:     "regen:13toVgyewRosPA4FVy4wWJgk7JGYc5K7TtE1nHaaHQJgvb6bBLtBBTC.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-002
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-001-20150101-20151231-002",
			Metadata:     "regen:13toVgbuejJuARiL27Js3Ek3bw3cFrpCN89agNL7pPksUjwQdLWnJRC.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-003
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-001-20150101-20151231-003",
			Metadata:     "regen:13toVgfKEu7dmUCsf6pfXKeWdNEaCAn8rhYB45gpJzoazQ1jEpRyapb.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-004
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-001-20150101-20151231-004",
			Metadata:     "regen:13toVhazYXg2LyQ7TXzEkBDsYi3wEyUzi56ZB6DgmbHGyLj9gUvfWHn.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-005
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-001-20150101-20151231-005",
			Metadata:     "regen:13toVgs8XiE5McGxNhf4hb4F6pxiRzQKXXwNNbbCs2VSDm8BWi94dQB.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-001
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-002-20190101-20191231-001",
			Metadata:     "regen:13toVgF84kQwSX11DdhDasYtMFU1Yb6qQwQvtv1rvH8pf5E8UTyajCX.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-002
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-002-20190101-20191231-002",
			Metadata:     "regen:13toVgR8xL6Nuyrotjaik7bqmkuWRnMvit8ka1fSBLnebzP7zUVbMJ3.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-003
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-002-20190101-20191231-003",
			Metadata:     "regen:13toVhAQCbMc2LJm44AV1enaqi27mRMkRjPJmVGdW11C4qcKuhrnGPA.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-004
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-002-20190101-20191231-004",
			Metadata:     "regen:13toVhJGk75xSgtrv831sgoEE6i3HhaWg7HtgtAxpFdk7Bp5Yz5Bz85.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-003-20150701-20160630-001
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-003-20150701-20160630-001",
			Metadata:     "regen:13toVhBNPukzs9mmH9j4yL7ZaAhBBbQ9D2M9Hiw2QkQR2MuFrmaUd5F.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-001-20180101-20181231-001
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C02-001-20180101-20181231-001",
			Metadata:     "regen:13toVgYscoLUM1ZVzavQv1wm4BBGRn36WGbtmr4Ppa8MrkKifYJxNUg.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-003-20200630-20220629-001
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C02-003-20200630-20220629-001",
			Metadata:     "regen:13toVfxW6oyujxDVKhFiUNK95FV92W1u5umcjED137HpnMcheqzzxUS.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-002-20211012-20241013-001
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C02-002-20211012-20241013-001",
			Metadata:     "regen:13toVh95YsiGdwkHXKeg2HtEXaiBn484JCK5wDAa6TijpZhUHkDjSBh.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-004-20210102-20211207-001
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C02-004-20210102-20211207-001",
			Metadata:     "regen:13toVgj1vCVsXxh713nC8joYfRgLRde8zryScve3PchGqQVWMqG85D9.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-005
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-002-20190101-20191231-005",
			Metadata:     "regen:13toVhUHUpYhfYLL2Pe4DhDrjFNmTQTnJBYsTQZDkZ9i9TboXgbDTJq.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-006
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-001-20150101-20151231-006",
			Metadata:     "regen:13toVgNQB7QfSsKEjFoRckAGLabSh8RoizLFtPxgkdswahbDuWiR2He.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-006
		{
			Issuer:       issuer,
			ProjectKey:   projectKey,
			Denom:        "C01-002-20190101-20191231-006",
			Metadata:     "regen:13toVgPD2BcAvUQ33EWf9CFQiuywRFoFeVoomKByrL68UGM6jpxtTiu.rdf",
			StartDate:    timestamp,
			EndDate:      timestamp,
			IssuanceDate: timestamp,
			Open:         false,
		},
	}

	// add batches to state
	for _, batch := range batches {
		require.NoError(t, baseStore.BatchTable().Insert(sdkCtx, batch))
	}

	curator := sdk.MustAccAddressFromBech32("regen1mrvlgpmrjn9s7r7ct69euqfgxjazjt2l5lzqcd")

	// http://mainnet.regen.network:1317/regen/ecocredit/basket/v1/baskets/eco.uC.NCT
	basket := &basketapi.Basket{
		BasketDenom:       "eco.uC.NCT",
		Name:              "NCT",
		DisableAutoRetire: false,
		CreditTypeAbbrev:  "C",
		DateCriteria:      nil,
		Exponent:          6,
		Curator:           curator,
	}

	// add basket to state
	require.NoError(t, basketStore.BasketTable().Insert(sdkCtx, basket))

	// execute state migrations
	require.NoError(t, v4.MigrateState(sdkCtx, baseStore, basketStore))

	b1, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-001-20150101-20151231-001")
	require.NoError(t, err)
	require.Equal(t, b1.Metadata, "regen:13toVg38ZRvFxPA2TBNnxGhabgogpJnv4LDm7YPgSuzuETiXz8GbnTF.rdf")

	b2, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-001-20150101-20151231-002")
	require.NoError(t, err)
	require.Equal(t, b2.Metadata, "regen:13toVhTFtGtXxoHw7yy3QQVDGEpSQoVy4VARhtTWeuNQa5V25WUhagq.rdf")

	b3, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-001-20150101-20151231-003")
	require.NoError(t, err)
	require.Equal(t, b3.Metadata, "regen:13toVghYySmmX9gm76MuwiPCC9AJy6Psb7wj6uj9JiBk4NvACGkpJDw.rdf")

	b4, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-001-20150101-20151231-004")
	require.NoError(t, err)
	require.Equal(t, b4.Metadata, "regen:13toVgGhCxGuNrqKKugLY9thKAdLTgXHGxhbVutz2QLgtFmdZzPAKUB.rdf")

	b5, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-001-20150101-20151231-005")
	require.NoError(t, err)
	require.Equal(t, b5.Metadata, "regen:13toVhaDUK1CHmqdZfKr6ZdF1L1ekTvUgjbEiGxWYqDWVZ937GUviFr.rdf")

	b6, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-002-20190101-20191231-001")
	require.NoError(t, err)
	require.Equal(t, b6.Metadata, "regen:13toVgu5VbjKdfDKuwPfUoMeo2isi1ApbsCsaTCoyNknKnN9FE6j1hW.rdf")

	b7, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-002-20190101-20191231-002")
	require.NoError(t, err)
	require.Equal(t, b7.Metadata, "regen:13toVgxDAxBev51DadD1he6gdkF6UPAoe25Y2xsSn3uDpzmHG3qGqRh.rdf")

	b8, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-002-20190101-20191231-003")
	require.NoError(t, err)
	require.Equal(t, b8.Metadata, "regen:13toVgsTujvEeCS4hG9MXZii8eF9LBkrgzaw8mh2q54KfFtGFtH5DLi.rdf")

	b9, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-002-20190101-20191231-004")
	require.NoError(t, err)
	require.Equal(t, b9.Metadata, "regen:13toVhKuR8NndSGZdciYTtCJf11hjYGwNvsWdjSPBmXNAqk8oL7u7XW.rdf")

	b10, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-003-20150701-20160630-001")
	require.NoError(t, err)
	require.Equal(t, b10.Metadata, "regen:13toVhMT8c7hFZePMqa8raLBgCuLFo4MWbJ7QJRheFRC5dfPcmFZ4hk.rdf")

	b11, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C02-001-20180101-20181231-001")
	require.NoError(t, err)
	require.Equal(t, b11.Metadata, "regen:13toVgpAwAm6fzYUUkD8UmioCYCP3GMbA3pdNkTM4wKeWc5UxmmCZW8.rdf")

	b12, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C02-003-20200630-20220629-001")
	require.NoError(t, err)
	require.Equal(t, b12.Metadata, "regen:13toVh5g1AhGAWcTQCBXra1YfD2XJUbH35dvSLuEPAR8mQHJDY5ovVe.rdf")

	b13, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C02-002-20211012-20241013-001")
	require.NoError(t, err)
	require.Equal(t, b13.Metadata, "regen:13toVhAukPXsjX5gMTADfUQzJQBjehJJoPwSavuU6GjyH5DtxZ5oVYS.rdf")

	b14, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C02-004-20210102-20211207-001")
	require.NoError(t, err)
	require.Equal(t, b14.Metadata, "regen:13toVh1EoPoJJs1VSvmeQB3HHpXDgFBA19KqiP1tg4NByjWsFLJdjuq.rdf")

	b15, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-002-20190101-20191231-005")
	require.NoError(t, err)
	require.Equal(t, b15.Metadata, "regen:13toVgwjqzxx3b9cRiRXBxrsUQB6D1WC4Kk8zZuXUfwnZ8WYtxRy4r5.rdf")

	b16, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-001-20150101-20151231-006")
	require.NoError(t, err)
	require.Equal(t, b16.Metadata, "regen:13toVhaPG4MzeWcmoriPhQ5jRGx6ohhdzBPREoarrdqTRVCP8Xj7scM.rdf")

	b17, err := baseStore.BatchTable().GetByDenom(sdkCtx, "C01-002-20190101-20191231-006")
	require.NoError(t, err)
	require.Equal(t, b17.Metadata, "regen:13toVh3NyL4uDzLFcrf6rUFMQnV8af87tBdSzh1Dvsc8zgEx193Y7hr.rdf")

	// ensure all other fields are unchanged
	for i := range batches {
		b, err := baseStore.BatchTable().Get(sdkCtx, uint64(i+1))
		require.NoError(t, err)

		require.Equal(t, b.Issuer, issuer.Bytes())
		require.Equal(t, b.ProjectKey, projectKey)
		require.Equal(t, b.StartDate.AsTime(), timestamp.AsTime())
		require.Equal(t, b.EndDate.AsTime(), timestamp.AsTime())
		require.Equal(t, b.IssuanceDate.AsTime(), timestamp.AsTime())
		require.Equal(t, b.Open, false)
	}

	b, err := basketStore.BasketTable().GetByBasketDenom(sdkCtx, "eco.uC.NCT")
	require.NoError(t, err)

	require.Equal(t, "NCT", b.Name)
	require.Equal(t, true, b.DisableAutoRetire)
	require.Equal(t, "C", b.CreditTypeAbbrev)
	require.Equal(t, (*timestamppb.Timestamp)(nil), b.DateCriteria.MinStartDate)
	require.Equal(t, (*durationpb.Duration)(nil), b.DateCriteria.StartDateWindow)
	require.Equal(t, uint32(10), b.DateCriteria.YearsInThePast)
	require.Equal(t, uint32(6), b.Exponent) //nolint:staticcheck
	require.Equal(t, curator.Bytes(), b.Curator)
}

func setup(t *testing.T) (sdk.Context, baseapi.StateStore, basketapi.StateStore) {
	ecocreditKey := sdk.NewKVStoreKey("ecocredit")

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(ecocreditKey, storetypes.StoreTypeIAVL, db)

	require.NoError(t, cms.LoadLatestVersion())

	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)

	modDB, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	baseStore, err := baseapi.NewStateStore(modDB)
	require.NoError(t, err)

	basketStore, err := basketapi.NewStateStore(modDB)
	require.NoError(t, err)

	return sdkCtx, baseStore, basketStore
}
