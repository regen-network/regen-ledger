package core

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// BridgeReceive bridges credits received from another chain.
func (k Keeper) BridgeReceive(ctx context.Context, req *core.MsgBridgeReceive) (*core.MsgBridgeReceiveResponse, error) {
	bridgeServiceAddr, err := sdk.AccAddressFromBech32(req.ServiceAddress)
	if err != nil {
		return nil, err
	}

	// first we check if there is an existing project
	idx := api.ProjectAdminReferenceIdIndexKey{}.WithAdminReferenceId(bridgeServiceAddr, req.ProjectRefId)
	it, err := k.stateStore.ProjectTable().List(ctx, idx)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	projects := make([]*api.Project, 0)
	for it.Next() {
		project, err := it.Value()
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if len(projects) > 1 {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("fatal error: bridge service %s has %d projects registered "+
			"with reference id %s", bridgeServiceAddr.String(), len(projects), req.ProjectRefId)
	}

	// if no project was found, create one + issue batch
	if len(projects) == 0 {
		projectRes, err := k.CreateProject(ctx, &core.MsgCreateProject{
			Issuer:       req.ServiceAddress,
			ClassId:      req.ClassId,
			Metadata:     req.ProjectMetadata,
			Jurisdiction: req.ProjectJurisdiction,
			ReferenceId:  req.ProjectRefId,
		})
		if err != nil {
			return nil, err
		}
		batchRes, err := k.CreateBatch(ctx, &core.MsgCreateBatch{
			Issuer:    req.ServiceAddress,
			ProjectId: projectRes.ProjectId,
			Issuance: []*core.BatchIssuance{
				{Recipient: req.Recipient, TradableAmount: req.Amount},
			},
			Metadata:  req.BatchMetadata,
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
			Open:      true,
			OriginTx:  req.OriginTx,
			Note:      req.Note,
		})
		if err != nil {
			return nil, err
		}
		return &core.MsgBridgeReceiveResponse{BatchDenom: batchRes.BatchDenom, ProjectId: projectRes.ProjectId}, nil
	}

	project := projects[0]

	// batches are matched on their denom, iterating over all batches within the <ProjectId>-<StartDate>-<EndDate> range.
	// any batches in that iterator that have matching metadata, are added to the slice.
	// idx will be of form C01-001-20210107-20210125-" catching all batches with that project Id and in the date range.
	batchIdx := fmt.Sprintf("%s-%s-%s-", project.Id, req.StartDate.Format("20060102"), req.EndDate.Format("20060102"))
	bIt, err := k.stateStore.BatchTable().List(ctx, api.BatchDenomIndexKey{}.WithDenom(batchIdx))
	if err != nil {
		return nil, err
	}
	defer bIt.Close()

	batches := make([]*api.Batch, 0)
	for bIt.Next() {
		batch, err := bIt.Value()
		if err != nil {
			return nil, err
		}
		// the timestamp stored in the batch is more granular than the date in the denom representation, so we match here.
		if batch.StartDate.AsTime().Equal(*req.StartDate) && batch.EndDate.AsTime().Equal(*req.EndDate) && batch.Metadata == req.BatchMetadata {
			batches = append(batches, batch)
		}
	}

	// TODO(Tyler): potentially select a batch by oldest issuance date?
	amtBatches := len(batches)
	if amtBatches > 1 { // change this to pick by issuance date
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("fatal error: bridge service %s has %d batches issued "+
			"with start %v and end %v dates in project %s", bridgeServiceAddr.String(), len(batches), req.StartDate, req.EndDate, project.Id)
	} else if amtBatches == 1 {
		batch := batches[0]
		// otherwise, we can simply mint into the batch
		_, err = k.MintBatchCredits(ctx, &core.MsgMintBatchCredits{
			Issuer:     req.ServiceAddress,
			BatchDenom: batch.Denom,
			Issuance: []*core.BatchIssuance{
				{Recipient: req.Recipient, TradableAmount: req.Amount},
			},
			OriginTx: req.OriginTx,
			Note:     req.Note,
		})
		return &core.MsgBridgeReceiveResponse{BatchDenom: batch.Denom, ProjectId: project.Id}, nil
	}
	// len(batches) is not greater than or equal to 1, so its empty, meaning no batch exists yet.
	res, err := k.CreateBatch(ctx, &core.MsgCreateBatch{
		Issuer:    req.ServiceAddress,
		ProjectId: project.Id,
		Issuance: []*core.BatchIssuance{
			{Recipient: req.Recipient, TradableAmount: req.Amount},
		},
		Metadata:  req.BatchMetadata,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Open:      true,
		OriginTx:  req.OriginTx,
		Note:      req.Note,
	})
	if err != nil {
		return nil, err
	}
	return &core.MsgBridgeReceiveResponse{BatchDenom: res.BatchDenom, ProjectId: project.Id}, nil
}
