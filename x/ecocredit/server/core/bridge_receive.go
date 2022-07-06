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

	project, err := k.getProjectFromBridgeReq(ctx, req.Project, req.ClassId)
	if err != nil {
		return nil, err
	}

	var event *core.EventBridgeReceive
	var response *core.MsgBridgeReceiveResponse

	// if no project was found, create one + issue batch
	if project == nil {
		projectRes, err := k.CreateProject(ctx, &core.MsgCreateProject{
			Admin:        req.Issuer,
			ClassId:      req.ClassId,
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
				{
					Recipient:      req.Batch.Recipient,
					TradableAmount: req.Batch.Amount,
				},
			},
			Metadata:  req.Batch.Metadata,
			StartDate: req.Batch.StartDate,
			EndDate:   req.Batch.EndDate,
			Open:      true,
			OriginTx:  req.OriginTx,
		})
		if err != nil {
			return nil, err
		}
		event = &core.EventBridgeReceive{
			BatchDenom: batchRes.BatchDenom,
			ProjectId:  projectRes.ProjectId,
		}
		response = &core.MsgBridgeReceiveResponse{
			BatchDenom: batchRes.BatchDenom,
			ProjectId:  projectRes.ProjectId,
		}
	} else {
		batch, err := k.getBatchFromBridgeReq(ctx, req.Batch, project.Id, bridgeServiceAddr)
		if err != nil {
			return nil, err
		}

		if batch != nil {
			_, err = k.MintBatchCredits(ctx, &core.MsgMintBatchCredits{
				Issuer:     req.Issuer,
				BatchDenom: batch.Denom,
				Issuance: []*core.BatchIssuance{
					{
						Recipient:      req.Batch.Recipient,
						TradableAmount: req.Batch.Amount,
					},
				},
				OriginTx: req.OriginTx,
			})
			if err != nil {
				return nil, err
			}
			event = &core.EventBridgeReceive{
				BatchDenom: batch.Denom,
				ProjectId:  project.Id,
			}
			response = &core.MsgBridgeReceiveResponse{
				BatchDenom: batch.Denom,
				ProjectId:  project.Id,
			}
		} else {
			// batch was nil, so we need to create one.
			res, err := k.CreateBatch(ctx, &core.MsgCreateBatch{
				Issuer:    req.Issuer,
				ProjectId: project.Id,
				Issuance: []*core.BatchIssuance{
					{
						Recipient:      req.Batch.Recipient,
						TradableAmount: req.Batch.Amount,
					},
				},
				Metadata:  req.Batch.Metadata,
				StartDate: req.Batch.StartDate,
				EndDate:   req.Batch.EndDate,
				Open:      true,
				OriginTx:  req.OriginTx,
			})
			if err != nil {
				return nil, err
			}
			event = &core.EventBridgeReceive{
				BatchDenom: res.BatchDenom,
				ProjectId:  project.Id,
			}
			response = &core.MsgBridgeReceiveResponse{
				BatchDenom: res.BatchDenom,
				ProjectId:  project.Id,
			}
		}
	}

	if err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(event); err != nil {
		return nil, err
	}

	return response, nil
}

// getBatchFromBridgeReq attempts to retrieve a batch from state given the request.
// In the event that multiple batches are matched, the batch with the oldest issuance date is selected.
// When no batches are found, nil is returned for both return values.
func (k Keeper) getBatchFromBridgeReq(ctx context.Context, req *core.MsgBridgeReceive_Batch, projectId string, bridgeAddr sdk.AccAddress) (*api.Batch, error) {
	// batches are matched on their denom, iterating over all batches within the <ProjectId>-<StartDate>-<EndDate> range.
	// any batches in that iterator that were created by the same issuer and have matching metadata, are added to the slice.
	// idx will be of form C01-001-20210107-20210125-" catching all batches with that project Id and in the date range.
	batchIdx := fmt.Sprintf("%s-%s-%s-", projectId, req.StartDate.UTC().Format("20060102"), req.EndDate.UTC().Format("20060102"))
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
		if batch.StartDate.AsTime().UTC().Equal(req.StartDate.UTC()) &&
			batch.EndDate.AsTime().UTC().Equal(req.EndDate.UTC()) &&
			batch.Metadata == req.Metadata &&
			sdk.AccAddress(batch.Issuer).Equals(bridgeAddr) {
			batches = append(batches, batch)
		}
	}
	bIt.Close()

	if len(batches) == 1 {
		return batches[0], nil
	} else if len(batches) > 1 {
		oldestIssuedBatch := batches[0]
		for i := 1; i < len(batches); i++ {
			if oldestIssuedBatch.IssuanceDate.AsTime().UTC().After(batches[i].IssuanceDate.AsTime().UTC()) {
				oldestIssuedBatch = batches[i]
			}
		}
		return oldestIssuedBatch, nil
	}
	return nil, nil
}

// getProjectFromBridgeReq attempts to find a project with a matching reference id within the scope
// of the credit class. No more than one project will be returned when we list the projects based on
// class id and reference id because we enforce uniqueness on non-empty reference ids within the scope
// of a credit class (and we do this at the message server level and not the ORM level because reference
// id is optional when using Msg/CreateProject). If no project is found, nil is returned for both values.
func (k Keeper) getProjectFromBridgeReq(ctx context.Context, req *core.MsgBridgeReceive_Project, classId string) (*api.Project, error) {
	class, err := k.stateStore.ClassTable().GetById(ctx, classId)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get class with id %s: %s", classId, err.Error())
	}

	// first we check if there is an existing project
	idx := api.ProjectClassKeyReferenceIdIndexKey{}.WithClassKeyReferenceId(class.Key, req.ReferenceId)
	it, err := k.stateStore.ProjectTable().List(ctx, idx)
	if err != nil {
		return nil, err
	}

	// we only want the first project that matches the reference ID, so we do not loop here. We enforce
	// uniqueness for a non-empty reference id at the message service level so as long as the reference
	// id is not empty (verified in validate basic), no more than one project will ever be returned.
	var project *api.Project
	if it.Next() {
		var err error
		project, err = it.Value()
		if err != nil {
			return nil, err
		}
	}
	it.Close()

	return project, nil
}
