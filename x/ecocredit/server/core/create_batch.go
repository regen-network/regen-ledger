package core

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

// CreateBatch creates a new batch of credits.
// Credits in the batch must not have more decimal places than the credit type's specified precision.
func (k Keeper) CreateBatch(ctx context.Context, req *core.MsgCreateBatch) (*core.MsgCreateBatchResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	projectID := req.ProjectId

	projectInfo, err := k.stateStore.ProjectInfoTable().GetByName(ctx, projectID)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassInfoTable().Get(ctx, projectInfo.ClassId)
	if err != nil {
		return nil, err
	}

	err = k.assertClassIssuer(ctx, projectInfo.ClassId, req.Issuer)
	if err != nil {
		return nil, err
	}

	batchSeqNo, err := k.getBatchSeqNo(ctx, projectID)
	if err != nil {
		return nil, err
	}

	batchDenom, err := ecocredit.FormatDenom(classInfo.Name, batchSeqNo, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	startDate, endDate := timestamppb.New(req.StartDate.UTC()), timestamppb.New(req.EndDate.UTC())
	rowID, err := k.stateStore.BatchInfoTable().InsertReturningID(ctx, &api.BatchInfo{
		ProjectId:  projectInfo.Id,
		BatchDenom: batchDenom,
		Metadata:   req.Metadata,
		StartDate:  startDate,
		EndDate:    endDate,
	})
	if err != nil {
		return nil, err
	}

	creditType, err := server.GetCreditTypeFromBatchDenom(ctx, k.stateStore, k.params, batchDenom)
	if err != nil {
		return nil, err
	}
	maxDecimalPlaces := creditType.Precision

	tradableSupply, retiredSupply := math.NewDecFromInt64(0), math.NewDecFromInt64(0)
	for _, issuance := range req.Issuance {
		decs, err := server.GetNonNegativeFixedDecs(maxDecimalPlaces, issuance.TradableAmount, issuance.RetiredAmount)
		if err != nil {
			return nil, err
		}
		tradable, retired := decs[0], decs[1]

		recipient, _ := sdk.AccAddressFromBech32(issuance.Recipient)
		if !tradable.IsZero() {
			tradableSupply, err = tradableSupply.Add(tradable)
			if err != nil {
				return nil, err
			}
		}
		if !retired.IsZero() {
			retiredSupply, err = retiredSupply.Add(retired)
			if err != nil {
				return nil, err
			}
			if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventRetire{
				Retirer:    recipient.String(),
				BatchDenom: batchDenom,
				Amount:     retired.String(),
				Location:   issuance.RetirementLocation,
			}); err != nil {
				return nil, err
			}
		}
		if err = k.stateStore.BatchBalanceTable().Insert(ctx, &api.BatchBalance{
			Address:  recipient,
			BatchId:  rowID,
			Tradable: tradable.String(),
			Retired:  retired.String(),
		}); err != nil {
			return nil, err
		}

		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventReceive{
			Recipient:      recipient.String(),
			BatchDenom:     batchDenom,
			RetiredAmount:  retired.String(),
			TradableAmount: tradable.String(),
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "batch issuance")
	}

	if err = k.stateStore.BatchSupplyTable().Insert(ctx, &api.BatchSupply{
		BatchId:         rowID,
		TradableAmount:  tradableSupply.String(),
		RetiredAmount:   retiredSupply.String(),
		CancelledAmount: math.NewDecFromInt64(0).String(),
	}); err != nil {
		return nil, err
	}

	totalAmount, err := tradableSupply.Add(retiredSupply)
	if err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventCreateBatch{
		ClassId:         classInfo.Name,
		BatchDenom:      batchDenom,
		Issuer:          req.Issuer,
		TotalAmount:     totalAmount.String(),
		StartDate:       startDate.String(),
		EndDate:         endDate.String(),
		ProjectLocation: projectInfo.ProjectLocation,
		ProjectId:       projectInfo.Name,
	}); err != nil {
		return nil, err
	}

	return &core.MsgCreateBatchResponse{BatchDenom: batchDenom}, nil
}

// getBatchSeqNo gets the batch sequence number
func (k Keeper) getBatchSeqNo(ctx context.Context, projectID string) (uint64, error) {
	var seq uint64 = 1
	batchSeq, err := k.stateStore.BatchSequenceTable().Get(ctx, projectID)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return 0, err
		}
	} else {
		seq = batchSeq.NextBatchId
	}

	if err = k.stateStore.BatchSequenceTable().Save(ctx, &api.BatchSequence{
		ProjectId:   projectID,
		NextBatchId: seq + 1,
	}); err != nil {
		return 0, err
	}

	return seq, err
}
