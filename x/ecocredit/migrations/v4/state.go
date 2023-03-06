package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
)

// MigrateState performs in-place store migrations from ConsensusVersion 3 to 4.
func MigrateState(sdkCtx sdk.Context, baseStore ecocreditv1.StateStore) error {
	var batches []Batch

	if sdkCtx.ChainID() == "regen-1" {
		batches = getMainnetBatches()
	}

	for _, batch := range batches {
		if err := migrateBatchMetadata(sdkCtx, baseStore, batch); err != nil {
			return err
		}
	}

	return nil
}

type Batch struct {
	Denom       string
	NewMetadata string
}

func getMainnetBatches() []Batch {
	return []Batch{
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-001
		{
			Denom:       "C01-001-20150101-20151231-001",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-002
		{
			Denom:       "C01-001-20150101-20151231-002",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-003
		{
			Denom:       "C01-001-20150101-20151231-003",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-004
		{
			Denom:       "C01-001-20150101-20151231-004",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-005
		{
			Denom:       "C01-001-20150101-20151231-005",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-001
		{
			Denom:       "C01-002-20190101-20191231-001",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-002
		{
			Denom:       "C01-002-20190101-20191231-002",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-003
		{
			Denom:       "C01-002-20190101-20191231-003",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-004
		{
			Denom:       "C01-002-20190101-20191231-004",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-003-20150701-20160630-001
		{
			Denom:       "C01-003-20150701-20160630-001",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-001-20180101-20181231-001
		{
			Denom:       "C02-001-20180101-20181231-001",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-003-20200630-20220629-001
		{
			Denom:       "C02-003-20200630-20220629-001",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-002-20211012-20241013-001
		{
			Denom:       "C02-002-20211012-20241013-001",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C02-004-20210102-20211207-001
		{
			Denom:       "C02-004-20210102-20211207-001",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-005
		{
			Denom:       "C01-002-20190101-20191231-005",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-001-20150101-20151231-006
		{
			Denom:       "C01-001-20150101-20151231-006",
			NewMetadata: "regen1...",
		},
		// http://mainnet.regen.network:1317/regen/ecocredit/v1/batches/C01-002-20190101-20191231-006
		{
			Denom:       "C01-002-20190101-20191231-006",
			NewMetadata: "regen1...",
		},
	}
}

func migrateBatchMetadata(ctx sdk.Context, baseStore ecocreditv1.StateStore, batch Batch) error {
	b, err := baseStore.BatchTable().GetByDenom(ctx, batch.Denom)
	if err != nil {
		return err
	}

	b.Metadata = batch.NewMetadata

	if err := baseStore.BatchTable().Update(ctx, b); err != nil {
		return err
	}

	return nil
}
