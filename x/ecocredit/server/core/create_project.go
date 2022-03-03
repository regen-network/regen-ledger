package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

func (k Keeper) CreateProject(ctx context.Context, req *v1beta1.MsgCreateProject) (*v1beta1.MsgCreateProjectResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	classID := req.ClassId
	classInfo, err := k.stateStore.ClassInfoStore().GetByName(ctx, classID)
	if err != nil {
		return nil, err
	}

	err = k.assertClassIssuer(ctx, classInfo.Id, req.Issuer)
	if err != nil {
		return nil, err
	}

	projectID := req.ProjectId
	if projectID == "" {
		exists := true
		for ; exists; sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "project id sequence") {
			projectID, err = k.genProjectID(ctx, classInfo.Id, classInfo.Name)
			if err != nil {
				return nil, err
			}
			exists, err = k.stateStore.ProjectInfoStore().HasByClassIdName(ctx, classInfo.Id, projectID)
			if err != nil {
				return nil, err
			}
		}
	}

	if err = k.stateStore.ProjectInfoStore().Insert(ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            projectID,
		ClassId:         classInfo.Id,
		ProjectLocation: req.ProjectLocation,
		Metadata:        req.Metadata,
	}); err != nil {
		return nil, err
	}

	if err := sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventCreateProject{
		ClassId:         classID,
		ProjectId:       projectID,
		Issuer:          req.Issuer,
		ProjectLocation: req.ProjectLocation,
	}); err != nil {
		return nil, err
	}

	return &v1beta1.MsgCreateProjectResponse{
		ProjectId: projectID,
	}, nil
}

func (k Keeper) genProjectID(ctx context.Context, classRowID uint64, classID string) (string, error) {
	var nextID uint64
	projectSeqNo, err := k.stateStore.ProjectSequenceStore().Get(ctx, classRowID)
	switch err {
	case ormerrors.NotFound:
		nextID = 1
	case nil:
		nextID = projectSeqNo.NextProjectId
	default:
		return "", err
	}

	if err = k.stateStore.ProjectSequenceStore().Save(ctx, &ecocreditv1beta1.ProjectSequence{
		ClassId:       classRowID,
		NextProjectId: nextID + 1,
	}); err != nil {
		return "", err
	}

	return ecocredit.FormatProjectID(classID, nextID), nil
}

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (k Keeper) assertClassIssuer(goCtx context.Context, classID uint64, issuer string) error {
	addr, _ := sdk.AccAddressFromBech32(issuer)
	found, err := k.stateStore.ClassIssuerStore().Has(goCtx, classID, addr)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.ErrUnauthorized.Wrapf("%s is not an issuer for the class", issuer)
	}
	return nil
}
