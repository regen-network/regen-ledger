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
	bridgeServiceAddr, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err
	}

	class, err := k.stateStore.ClassTable().GetById(ctx, req.Project.ClassId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get class with id %s: %s", req.Project.ClassId, err.Error())
	}

	// first we check if there is an existing project
	idx := api.ProjectClassKeyReferenceIdIndexKey{}.WithClassKeyReferenceId(class.Key, req.Project.ReferenceId)
	it, err := k.stateStore.ProjectTable().List(ctx, idx)
	if err != nil {
		return nil, err
	}

	// we only want the first project that matches the reference ID, so we do not loop here.
	var project *api.Project
	if it.Next() {
		var err error
		project, err = it.Value()
		if err != nil {
			return nil, err
		}
	}
	it.Close()

	// if no project was found, create one + issue batch
	if project == nil {
		projectRes, err := k.CreateProject(ctx, &core.MsgCreateProject{
			Admin:        req.Issuer,
			ClassId:      req.Project.ClassId,
			Metadata:     req.Project.Metadata,
			Jurisdiction: req.Project.Jurisdiction,
			ReferenceId:  req.Project.ReferenceId,
		})
		if err != nil {
			return nil, err
		}
		batchRes, err := k.CreateBatch(ctx, &core.MsgCreateBatch{
			Issuer:    req.Issuer,
			ProjectId: projectRes.ProjectId,
			Issuance: []*core.BatchIssuance{
				{Recipient: req.Batch.Recipient, TradableAmount: req.Batch.Amount},
			},
			Metadata:  req.Batch.Metadata,
			StartDate: req.Batch.StartDate,
			EndDate:   req.Batch.EndDate,
			Open:      true,
			OriginTx:  req.Batch.OriginTx,
			Note:      req.Batch.Note,
		})
		if err != nil {
			return nil, err
		}
		return &core.MsgBridgeReceiveResponse{BatchDenom: batchRes.BatchDenom, ProjectId: projectRes.ProjectId}, nil
	}

	// batches are matched on their denom, iterating over all batches within the <ProjectId>-<StartDate>-<EndDate> range.
	// any batches in that iterator that have matching metadata, are added to the slice.
	// idx will be of form C01-001-20210107-20210125-" catching all batches with that project Id and in the date range.
	batchIdx := fmt.Sprintf("%s-%s-%s-", project.Id, req.Batch.StartDate.Format("20060102"), req.Batch.EndDate.Format("20060102"))
	bIt, err := k.stateStore.BatchTable().List(ctx, api.BatchDenomIndexKey{}.WithDenom(batchIdx))
	if err != nil {
		return nil, err
	}

	batches := make([]*api.Batch, 0)
	for bIt.Next() {
		batch, err := bIt.Value()
		if err != nil {
			return nil, err
		}
		// the timestamp stored in the batch is more granular than the date in the denom representation, so we match here.
		if batch.StartDate.AsTime().Equal(*req.Batch.StartDate) && batch.EndDate.AsTime().Equal(*req.Batch.EndDate) && batch.Metadata == req.Batch.Metadata {
			batches = append(batches, batch)
		}
	}
	it.Close()

	// TODO(Tyler): potentially select a batch by oldest issuance date?
	amtBatches := len(batches)
	if amtBatches > 1 {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("fatal error: bridge service %s has %d batches issued "+
			"with start %v and end %v dates in project %s", bridgeServiceAddr.String(), len(batches), req.Batch.StartDate, req.Batch.EndDate, project.Id)
	} else if amtBatches == 1 {
		batch := batches[0]
		// otherwise, we can simply mint into the batch
		_, err = k.MintBatchCredits(ctx, &core.MsgMintBatchCredits{
			Issuer:     req.Issuer,
			BatchDenom: batch.Denom,
			Issuance: []*core.BatchIssuance{
				{Recipient: req.Batch.Recipient, TradableAmount: req.Batch.Amount},
			},
			OriginTx: req.Batch.OriginTx,
			Note:     req.Batch.Note,
		})
		return &core.MsgBridgeReceiveResponse{BatchDenom: batch.Denom, ProjectId: project.Id}, nil
	}

	// len(batches) is not greater than or equal to 1, so its empty, meaning no batch exists yet.
	res, err := k.CreateBatch(ctx, &core.MsgCreateBatch{
		Issuer:    req.Issuer,
		ProjectId: project.Id,
		Issuance: []*core.BatchIssuance{
			{Recipient: req.Batch.Recipient, TradableAmount: req.Batch.Amount},
		},
		Metadata:  req.Batch.Metadata,
		StartDate: req.Batch.StartDate,
		EndDate:   req.Batch.EndDate,
		Open:      true,
		OriginTx:  req.Batch.OriginTx,
		Note:      req.Batch.Note,
	})
	if err != nil {
		return nil, err
	}
	return &core.MsgBridgeReceiveResponse{BatchDenom: res.BatchDenom, ProjectId: project.Id}, nil
}
