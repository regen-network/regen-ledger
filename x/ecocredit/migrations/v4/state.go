package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	basketv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
)

// MigrateState performs in-place store migrations from ConsensusVersion 3 to 4.
func MigrateState(sdkCtx sdk.Context, baseStore ecocreditv1.StateStore, basketStore basketv1.StateStore) error {
	if sdkCtx.ChainID() == "regen-1" {

		// mainnet batch metadata migration
		batchUpdates := getBatchUpdates()
		for _, batchUpdate := range batchUpdates {
			if err := migrateBatchMetadata(sdkCtx, baseStore, batchUpdate); err != nil {
				return err
			}
		}

		// mainnet basket criteria migration
		if err := migrateBasketCriteria(sdkCtx, basketStore); err != nil {
			return err
		}
	}

	return nil
}

type Batch struct {
	Denom       string
	NewMetadata string
}

func getBatchUpdates() []Batch {
	return []Batch{
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-001
		{
			Denom:       "C01-001-20150101-20151231-001",
			NewMetadata: "regen:13toVg38ZRvFxPA2TBNnxGhabgogpJnv4LDm7YPgSuzuETiXz8GbnTF.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-002
		{
			Denom:       "C01-001-20150101-20151231-002",
			NewMetadata: "regen:13toVhTFtGtXxoHw7yy3QQVDGEpSQoVy4VARhtTWeuNQa5V25WUhagq.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-003
		{
			Denom:       "C01-001-20150101-20151231-003",
			NewMetadata: "regen:13toVghYySmmX9gm76MuwiPCC9AJy6Psb7wj6uj9JiBk4NvACGkpJDw.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-004
		{
			Denom:       "C01-001-20150101-20151231-004",
			NewMetadata: "regen:13toVgGhCxGuNrqKKugLY9thKAdLTgXHGxhbVutz2QLgtFmdZzPAKUB.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-005
		{
			Denom:       "C01-001-20150101-20151231-005",
			NewMetadata: "regen:13toVhaDUK1CHmqdZfKr6ZdF1L1ekTvUgjbEiGxWYqDWVZ937GUviFr.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-001
		{
			Denom:       "C01-002-20190101-20191231-001",
			NewMetadata: "regen:13toVgu5VbjKdfDKuwPfUoMeo2isi1ApbsCsaTCoyNknKnN9FE6j1hW.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-002
		{
			Denom:       "C01-002-20190101-20191231-002",
			NewMetadata: "regen:13toVgxDAxBev51DadD1he6gdkF6UPAoe25Y2xsSn3uDpzmHG3qGqRh.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-003
		{
			Denom:       "C01-002-20190101-20191231-003",
			NewMetadata: "regen:13toVgsTujvEeCS4hG9MXZii8eF9LBkrgzaw8mh2q54KfFtGFtH5DLi.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-004
		{
			Denom:       "C01-002-20190101-20191231-004",
			NewMetadata: "regen:13toVhKuR8NndSGZdciYTtCJf11hjYGwNvsWdjSPBmXNAqk8oL7u7XW.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-003-20150701-20160630-001
		{
			Denom:       "C01-003-20150701-20160630-001",
			NewMetadata: "regen:13toVhMT8c7hFZePMqa8raLBgCuLFo4MWbJ7QJRheFRC5dfPcmFZ4hk.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-001-20180101-20181231-001
		{
			Denom:       "C02-001-20180101-20181231-001",
			NewMetadata: "regen:13toVgpAwAm6fzYUUkD8UmioCYCP3GMbA3pdNkTM4wKeWc5UxmmCZW8.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-003-20200630-20220629-001
		{
			Denom:       "C02-003-20200630-20220629-001",
			NewMetadata: "regen:13toVh5g1AhGAWcTQCBXra1YfD2XJUbH35dvSLuEPAR8mQHJDY5ovVe.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-002-20211012-20241013-001
		{
			Denom:       "C02-002-20211012-20241013-001",
			NewMetadata: "regen:13toVhAukPXsjX5gMTADfUQzJQBjehJJoPwSavuU6GjyH5DtxZ5oVYS.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-004-20210102-20211207-001
		{
			Denom:       "C02-004-20210102-20211207-001",
			NewMetadata: "regen:13toVh1EoPoJJs1VSvmeQB3HHpXDgFBA19KqiP1tg4NByjWsFLJdjuq.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-005
		{
			Denom:       "C01-002-20190101-20191231-005",
			NewMetadata: "regen:13toVgwjqzxx3b9cRiRXBxrsUQB6D1WC4Kk8zZuXUfwnZ8WYtxRy4r5.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-006
		{
			Denom:       "C01-001-20150101-20151231-006",
			NewMetadata: "regen:13toVhaPG4MzeWcmoriPhQ5jRGx6ohhdzBPREoarrdqTRVCP8Xj7scM.rdf",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-006
		{
			Denom:       "C01-002-20190101-20191231-006",
			NewMetadata: "regen:13toVh3NyL4uDzLFcrf6rUFMQnV8af87tBdSzh1Dvsc8zgEx193Y7hr.rdf",
		},
	}
}

func migrateBatchMetadata(ctx sdk.Context, baseStore ecocreditv1.StateStore, batch Batch) error {
	b, err := baseStore.BatchTable().GetByDenom(ctx, batch.Denom)
	if err != nil {
		return err
	}

	b.Metadata = batch.NewMetadata

	return baseStore.BatchTable().Update(ctx, b)
}

func migrateBasketCriteria(ctx sdk.Context, basketStore basketv1.StateStore) error {
	b, err := basketStore.BasketTable().GetByBasketDenom(ctx, "eco.uC.NCT")
	if err != nil {
		return err
	}

	b.DisableAutoRetire = true

	b.DateCriteria = &basketv1.DateCriteria{
		YearsInThePast: 10,
	}

	return basketStore.BasketTable().Update(ctx, b)
}
