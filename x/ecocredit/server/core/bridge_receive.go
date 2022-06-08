package core

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// BridgeReceive bridges credits received from another chain.
func (k Keeper) BridgeReceive(ctx context.Context, req *core.MsgBridgeReceive) (*core.MsgBridgeReceiveResponse, error) {

	// first we check if there is an existing project
	it, err := k.stateStore.ProjectTable().List(ctx, api.ProjectReferenceIdIndexKey{}.WithReferenceId(req.ProjectRefId))
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
	// TODO: what do we do if theres multiple projects with this ref id?
	// project := projects[0]

	// if no project was found, we create one
	if len(projects) == 0 {
		projectRes, err := k.CreateProject(ctx, &core.MsgCreateProject{
			Issuer:       req.ServiceAddress,
			ClassId:      req.ClassId, // TODO: where would this come from??
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
		return &core.MsgBridgeReceiveResponse{BatchDenom: batchRes.BatchDenom}, nil
	}
	// otherwise, we can simply mint into the batch
	_, err = k.MintBatchCredits(ctx, &core.MsgMintBatchCredits{
		Issuer:     req.ServiceAddress,
		BatchDenom: "", // TODO(Tyler): does this come from the bridge? do we form it somehow?
		Issuance: []*core.BatchIssuance{
			{Recipient: req.Recipient, TradableAmount: req.Amount},
		},
		OriginTx: req.OriginTx,
		Note:     req.Note,
	})
	if err != nil {
		return nil, err
	}
	return &core.MsgBridgeReceiveResponse{}, nil
}
