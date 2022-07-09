package v3

import (
	"context"
	"fmt"
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
	// KE              -> ""
	// US-FL 98765      -> ""

	locationToReferenceIdMap := make(map[string]string)
	locationToReferenceIdMap["FR"] = ""
	locationToReferenceIdMap["US"] = ""
	locationToReferenceIdMap["AU-NSW 2453"] = ""
	locationToReferenceIdMap["KE"] = ""
	locationToReferenceIdMap["US-FL 98765"] = ""

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
	//  C01-18540707-19870212-034  ->  "2022-06-28T18:10:53Z"
	//  C01-19900103-19990103-032  ->  "2022-06-28T08:36:25Z"
	//  C01-20130604-20130605-038  ->  "2022-06-30T07:16:34Z"
	//  C01-20140604-20200603-036  ->  "2022-06-29T08:45:12Z"
	//  C01-20170605-20180601-031  ->  "2022-06-28T07:22:43Z"
	//  C01-20170606-20210601-030  ->  "2022-06-28T07:06:46Z"
	//  C01-20170606-20230607-014  ->  "2022-06-21T11:10:39Z"
	//  C01-20170606-20230607-015  ->  "2022-06-21T11:33:53Z"
	//  C01-20170606-20230607-016  ->  "2022-06-21T12:17:52Z"
	//  C01-20170606-20230607-017  ->  "2022-06-21T12:36:28Z"
	//  C01-20170606-20230607-018  ->  "2022-06-21T12:43:25Z"
	//  C01-20170606-20230607-019  ->  "2022-06-21T12:44:24Z"
	//  C01-20170606-20230607-020  ->  "2022-06-21T12:45:23Z"
	//  C01-20170606-20230607-021  ->  "2022-06-21T12:47:52Z"
	//  C01-20170606-20230607-022  ->  "2022-06-21T12:49:44Z"
	//  C01-20170606-20230607-023  ->  "2022-06-21T12:51:15Z"
	//  C01-20170606-20230607-024  ->  "2022-06-21T14:45:26Z"
	//  C01-20170606-20230607-026  ->  "2022-06-23T09:22:58Z"
	//  C01-20170606-20230607-029  ->  "2022-06-27T18:44:05Z"
	//  C01-20170606-20230608-025  ->  "2022-06-21T23:17:48Z"
	//  C01-20170607-20220622-035  ->  "2022-06-29T07:13:52Z"
	//  C01-20170607-20230608-037  ->  "2022-06-29T13:57:54Z"
	//  C01-20170611-20230614-028  ->  "2022-06-27T16:39:28Z"
	//  C01-20170613-20230606-027  ->  "2022-06-27T14:16:55Z"
	//  C01-20170613-20230622-010  ->  "2022-06-16T10:38:23Z"
	//  C01-20170613-20230622-011  ->  "2022-06-16T10:41:35Z"
	//  C01-20170613-20230622-012  ->  "2022-06-16T13:14:01Z"
	//  C01-20170613-20230622-013  ->  "2022-06-16T17:47:30Z"
	//  C01-20180507-20240607-033  ->  "2022-06-28T18:08:34Z"

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
	batchIdToIssuanceDateMap["C01-18540707-19870212-034"] = "2022-06-28T18:10:53Z"
	batchIdToIssuanceDateMap["C01-19900103-19990103-032"] = "2022-06-28T08:36:25Z"
	batchIdToIssuanceDateMap["C01-20130604-20130605-038"] = "2022-06-30T07:16:34Z"
	batchIdToIssuanceDateMap["C01-20140604-20200603-036"] = "2022-06-29T08:45:12Z"
	batchIdToIssuanceDateMap["C01-20170605-20180601-031"] = "2022-06-28T07:22:43Z"
	batchIdToIssuanceDateMap["C01-20170606-20210601-030"] = "2022-06-28T07:06:46Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-014"] = "2022-06-21T11:10:39Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-015"] = "2022-06-21T11:33:53Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-016"] = "2022-06-21T12:17:52Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-017"] = "2022-06-21T12:36:28Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-018"] = "2022-06-21T12:43:25Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-019"] = "2022-06-21T12:44:24Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-020"] = "2022-06-21T12:45:23Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-021"] = "2022-06-21T12:47:52Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-022"] = "2022-06-21T12:49:44Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-023"] = "2022-06-21T12:51:15Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-024"] = "2022-06-21T14:45:26Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-026"] = "2022-06-23T09:22:58Z"
	batchIdToIssuanceDateMap["C01-20170606-20230607-029"] = "2022-06-27T18:44:05Z"
	batchIdToIssuanceDateMap["C01-20170606-20230608-025"] = "2022-06-21T23:17:48Z"
	batchIdToIssuanceDateMap["C01-20170607-20220622-035"] = "2022-06-29T07:13:52Z"
	batchIdToIssuanceDateMap["C01-20170607-20230608-037"] = "2022-06-29T13:57:54Z"
	batchIdToIssuanceDateMap["C01-20170611-20230614-028"] = "2022-06-27T16:39:28Z"
	batchIdToIssuanceDateMap["C01-20170613-20230606-027"] = "2022-06-27T14:16:55Z"
	batchIdToIssuanceDateMap["C01-20170613-20230622-010"] = "2022-06-16T10:38:23Z"
	batchIdToIssuanceDateMap["C01-20170613-20230622-011"] = "2022-06-16T10:41:35Z"
	batchIdToIssuanceDateMap["C01-20170613-20230622-012"] = "2022-06-16T13:14:01Z"
	batchIdToIssuanceDateMap["C01-20170613-20230622-013"] = "2022-06-16T17:47:30Z"
	batchIdToIssuanceDateMap["C01-20180507-20240607-033"] = "2022-06-28T18:08:34Z"

	// update issuance date for credit batches
	if err := updateBatchIssueanceDate(ctx, ss, oldBatchDenomToNewDenomMap, batchIdToIssuanceDateMap); err != nil {
		return err
	}

	// basket name to curator
	// rNCT   -> regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46
	// NCT    -> regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46
	// TYLER  -> regen1yqr0pf38v9j7ah79wmkacau5mdspsc7l0sjeva

	basketNameToCuratorMap := make(map[string]string)
	basketNameToCuratorMap["rNCT"] = "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46"
	basketNameToCuratorMap["NCT"] = "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46"
	basketNameToCuratorMap["TYLER"] = "regen1yqr0pf38v9j7ah79wmkacau5mdspsc7l0sjeva"
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

	var ok bool
	for itr.Next() {
		project, err := itr.Value()
		if err != nil {
			return err
		}

		project.ReferenceId, ok = locationToReferenceIdMap[project.Jurisdiction]
		if !ok {
			return fmt.Errorf("reference id is not exist for %s jurisdiction", project.Jurisdiction)
		}

		if err := ss.ProjectTable().Update(ctx, project); err != nil {
			return err
		}
	}

	return nil
}

func updateBatchIssueanceDate(ctx context.Context, ss api.StateStore,
	oldBatchDenomToNewDenomMap map[string]string, batchIdToIssuanceDateMap map[string]string) error {
	for denom, issuanceDate := range batchIdToIssuanceDateMap {
		newDenom, ok := oldBatchDenomToNewDenomMap[denom]
		if !ok {
			return fmt.Errorf("new batch denom not found for %s", denom)
		}

		batch, err := ss.BatchTable().GetByDenom(ctx, newDenom)
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
