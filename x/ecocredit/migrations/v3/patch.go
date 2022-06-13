package v3

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// patchMigrate performs following migrations
// - add reference id to existing projects
// - update issuance date for credit batches
// - update curator address for baskets
func patchMigrate(ctx context.Context, sdkCtx sdk.Context, ss api.StateStore,
	basketStore basketapi.StateStore, oldBatchDenomToNewDenomMap map[string]string) error {
	if sdkCtx.ChainID() == "regen-1" {
		return patchMainnet(ctx, ss, oldBatchDenomToNewDenomMap)
	} else if sdkCtx.ChainID() == "regen-redwood-1" {
		return patchRedwood(ctx, ss, basketStore, oldBatchDenomToNewDenomMap)
	}

	return nil
}

func patchMainnet(ctx context.Context, ss api.StateStore, oldBatchDenomToNewDenomMap map[string]string) error {
	// project location -> reference-id
	// KE    -> "VCS-612" (Kasigao)
	// CD-MN -> "VCS-934" (Mai Ndombe)

	locationToReferenceIdMap := make(map[string]string)
	locationToReferenceIdMap["KE"] = "VCS-612"
	locationToReferenceIdMap["CD-MN"] = "VCS-934"

	// add reference id to existing projects
	if err := addReferenceIds(ctx, ss, locationToReferenceIdMap); err != nil {
		return err
	}

	// batch issuance dates
	//  C01-20190101-20191231-001  -  "2022-05-06T01:33:13Z"
	//  C01-20190101-20191231-002  -  "2022-05-06T01:33:19Z"
	//  C01-20150101-20151231-003  -  "2022-05-06T01:33:25Z"
	//  C01-20150101-20151231-004  -  "2022-05-06T01:33:31Z"

	batchIdToIssuanceDateMap := make(map[string]string)
	batchIdToIssuanceDateMap["C01-20190101-20191231-001"] = "2022-05-06T01:33:13Z"
	batchIdToIssuanceDateMap["C01-20190101-20191231-002"] = "2022-05-06T01:33:19Z"
	batchIdToIssuanceDateMap["C01-20150101-20151231-003"] = "2022-05-06T01:33:25Z"
	batchIdToIssuanceDateMap["C01-20150101-20151231-004"] = "2022-05-06T01:33:31Z"
	// update issuance date for credit batches
	if err := updateBatchIssueanceDate(ctx, ss, oldBatchDenomToNewDenomMap, batchIdToIssuanceDateMap); err != nil {
		return err
	}

	// we don't have baskets on mainnet

	return nil
}

func patchRedwood(ctx context.Context, ss api.StateStore,
	basketStore basketapi.StateStore, oldBatchDenomToNewDenomMap map[string]string) error {
	// project location -> reference-id
	// FR              -> "" // TODO: add reference-id
	// US              -> ""
	// AU-NSW 2453     -> ""
	// K               -> ""
	// US-FL 9876      -> ""

	locationToReferenceIdMap := make(map[string]string)
	locationToReferenceIdMap["FR"] = ""
	locationToReferenceIdMap["US"] = ""
	locationToReferenceIdMap["AU-NSW 2453"] = ""
	locationToReferenceIdMap["K"] = ""
	locationToReferenceIdMap["US-FL 9876"] = ""

	// add reference id to existing projects
	if err := addReferenceIds(ctx, ss, locationToReferenceIdMap); err != nil {
		return err
	}

	// batch issuance dates
	//  C01-20210909-20220101-003  ->  "2022-03-08T17:18:19Z"
	//  C01-20190101-20210101-008  ->  "2022-04-22T16:32:09Z"
	//  C01-20190101-20210101-002  ->  "2022-02-14T09:07:25Z"
	//  C01-20190101-20191231-009  ->  "2022-05-11T11:35:30Z"
	//  C01-20180909-20200101-004  ->  "2022-03-08T17:25:23Z"
	//  C01-20180101-20200101-001  ->  "2022-02-09T09:10:02Z"
	//  C01-20170101-20180101-007  ->  "2022-03-30T15:04:30Z"
	//  C01-20170101-20180101-006  ->  "2022-03-30T07:51:12Z"
	//  C01-20170101-20180101-005  ->  "2022-03-30T07:46:01Z"
	//  C02-20200909-20210909-001  ->  "2022-03-08T13:00:50Z"
	//  C02-20210909-20220101-002  ->  "2022-03-08T17:17:20Z"
	//  C04-20180202-20190202-001  ->  "2022-03-28T08:31:45Z"
	//  C04-20190202-20200202-002  ->  "2022-03-28T08:45:14Z"

	batchIdToIssuanceDateMap := make(map[string]string)
	batchIdToIssuanceDateMap["C01-20210909-20220101-003"] = "2022-03-08T17:18:19Z"
	batchIdToIssuanceDateMap["C01-20190101-20210101-008"] = "2022-04-22T16:32:09Z"
	batchIdToIssuanceDateMap["C01-20190101-20210101-002"] = "2022-02-14T09:07:25Z"
	batchIdToIssuanceDateMap["C01-20190101-20191231-009"] = "2022-05-11T11:35:30Z"
	batchIdToIssuanceDateMap["C01-20180909-20200101-004"] = "2022-03-08T17:25:23Z"
	batchIdToIssuanceDateMap["C01-20180101-20200101-001"] = "2022-02-09T09:10:02Z"
	batchIdToIssuanceDateMap["C01-20170101-20180101-007"] = "2022-03-30T15:04:30Z"
	batchIdToIssuanceDateMap["C01-20170101-20180101-006"] = "2022-03-30T07:51:12Z"
	batchIdToIssuanceDateMap["C01-20170101-20180101-005"] = "2022-03-30T07:46:01Z"
	batchIdToIssuanceDateMap["C02-20200909-20210909-001"] = "2022-03-08T13:00:50Z"
	batchIdToIssuanceDateMap["C02-20210909-20220101-002"] = "2022-03-08T17:17:20Z"
	batchIdToIssuanceDateMap["C04-20180202-20190202-001"] = "2022-03-28T08:31:45Z"
	batchIdToIssuanceDateMap["C04-20190202-20200202-002"] = "2022-03-28T08:45:14Z"

	// update issuance date for credit batches
	if err := updateBatchIssueanceDate(ctx, ss, oldBatchDenomToNewDenomMap, batchIdToIssuanceDateMap); err != nil {
		return err
	}

	// basket name to curator
	// rNCT   -> regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46
	// NCT    -> regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46

	basketNameToCuratorMap := make(map[string]string)
	basketNameToCuratorMap["rNCT"] = "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46"
	basketNameToCuratorMap["NCT"] = "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46"
	if err := updateBasketCurator(ctx, ss, basketStore, basketNameToCuratorMap); err != nil {
		return err
	}

	return nil
}

func updateBasketCurator(ctx context.Context, ss api.StateStore, basketStore basketapi.StateStore,
	basketNameToCuratorMap map[string]string) error {
	for name, curator := range basketNameToCuratorMap {
		basket, err := basketStore.BasketTable().GetByName(ctx, name)
		if err != nil {
			return err
		}
		basket.Curator = sdk.AccAddress(curator)
		if err := basketStore.BasketTable().Update(ctx, basket); err != nil {
			return err
		}
	}

	return nil
}

func addReferenceIds(ctx context.Context, ss api.StateStore, locationToReferenceIdMap map[string]string) error {
	itr, err := ss.ProjectTable().List(ctx, api.ProjectKeyIndexKey{})
	if err != nil {
		return err
	}
	defer itr.Close()

	for itr.Next() {
		project, err := itr.Value()
		if err != nil {
			return err
		}

		project.ReferenceId = locationToReferenceIdMap[project.Jurisdiction]
		if err := ss.ProjectTable().Update(ctx, project); err != nil {
			return err
		}
	}

	return nil
}

func updateBatchIssueanceDate(ctx context.Context, ss api.StateStore,
	oldBatchDenomToNewDenomMap map[string]string, batchIdToIssuanceDateMap map[string]string) error {
	for denom, issuanceDate := range batchIdToIssuanceDateMap {
		batch, err := ss.BatchTable().GetByDenom(ctx, oldBatchDenomToNewDenomMap[denom])
		if err != nil {
			return err
		}

		parsed, err := time.Parse(time.RFC3339, issuanceDate)
		if err != nil {
			return err
		}
		batch.IssuanceDate = timestamppb.New(parsed)
		if err := ss.BatchTable().Update(ctx, batch); err != nil {
			return err
		}
	}

	return nil
}
