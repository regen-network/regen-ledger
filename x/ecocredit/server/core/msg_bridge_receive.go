package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// BridgeReceive bridges credits received from another chain.
func (k Keeper) BridgeReceive(ctx context.Context, req *core.MsgBridgeReceive) (*core.MsgBridgeReceiveResponse, error) {
	// check class id and get class information (specifically class key)
	class, err := k.stateStore.ClassTable().GetById(ctx, req.ClassId)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("credit class with id %s", req.ClassId)
		}
		return nil, err
	}

	// check if batch contract entry exists
	batchContract, err := k.stateStore.BatchContractTable().GetByClassKeyContract(ctx, class.Key, req.OriginTx.Contract)
	if err != nil {
		if !ormerrors.NotFound.Is(err) {
			return nil, err
		}
	}

	var event *core.EventBridgeReceive
	var response *core.MsgBridgeReceiveResponse

	// if batch contract entry with matching contract exists, and therefore a
	// project exists, dynamically mint credits to the existing credit batch,
	// otherwise search for an existing project based on credit class id and
	// project reference id and, if the project exists, create a credit batch
	// under the existing project, otherwise, create a new project and then a
	// new credit batch under the new project
	if batchContract != nil {

		// get batch information (specifically batch denom)
		batch, err := k.stateStore.BatchTable().Get(ctx, batchContract.BatchKey)
		if err != nil {
			return nil, err
		}

		// get project information (specifically project id)
		project, err := k.stateStore.ProjectTable().Get(ctx, batch.ProjectKey)
		if err != nil {
			return nil, err
		}

		// mint credits to the existing credit batch
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

		// set bridge receive event
		event = &core.EventBridgeReceive{
			BatchDenom: batch.Denom,
			ProjectId:  project.Id,
		}

		// set bridge receive response
		response = &core.MsgBridgeReceiveResponse{
			BatchDenom: batch.Denom,
			ProjectId:  project.Id,
		}
	} else {

		// attempt to find existing project based on credit class and reference id
		project, err := k.getProjectFromBridgeReq(ctx, req.Project, req.ClassId)
		if err != nil {
			return nil, err
		}

		// if no project exists that matches the credit class and project reference
		// id, then we create a new project with the information provided
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

			// set project id using the new project id
			project = &api.Project{Id: projectRes.ProjectId}
		}

		// create a new credit batch with the information provided
		batchRes, err := k.CreateBatch(ctx, &core.MsgCreateBatch{
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

		// set bridge receive event
		event = &core.EventBridgeReceive{
			BatchDenom: batchRes.BatchDenom,
			ProjectId:  project.Id,
		}

		// set bridge receive response
		response = &core.MsgBridgeReceiveResponse{
			BatchDenom: batchRes.BatchDenom,
			ProjectId:  project.Id,
		}
	}

	if err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(event); err != nil {
		return nil, err
	}

	return response, nil
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
