package core

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

// CreateBatch creates a new batch of credits.
// Credits in the batch must not have more decimal places than the credit type's specified precision.
func (k Keeper) CreateBatch(ctx context.Context, req *core.MsgCreateBatch) (*core.MsgCreateBatchResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	projectInfo, err := k.stateStore.ProjectTable().GetById(ctx, req.ProjectId)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassTable().Get(ctx, projectInfo.ClassKey)
	if err != nil {
		return nil, err
	}

	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err
	}

	err = k.assertClassIssuer(ctx, classInfo.Key, issuer)
	if err != nil {
		return nil, err
	}

	batchSeqNo, err := k.getBatchSeqNo(ctx, projectInfo.Key)
	if err != nil {
		return nil, err
	}

	batchDenom, err := core.FormatBatchDenom(projectInfo.Id, batchSeqNo, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	startDate, endDate := timestamppb.New(req.StartDate.UTC()), timestamppb.New(req.EndDate.UTC())
	issuanceDate := timestamppb.New(sdkCtx.BlockTime())
	key, err := k.stateStore.BatchTable().InsertReturningID(ctx, &api.Batch{
		ProjectKey:   projectInfo.Key,
		Issuer:       issuer,
		Denom:        batchDenom,
		Metadata:     req.Metadata,
		StartDate:    startDate,
		EndDate:      endDate,
		IssuanceDate: issuanceDate,
	})
	if err != nil {
		return nil, err
	}

	creditType, err := utils.GetCreditTypeFromBatchDenom(ctx, k.stateStore, batchDenom)
	if err != nil {
		return nil, err
	}
	maxDecimalPlaces := creditType.Precision

	tradableSupply, retiredSupply := math.NewDecFromInt64(0), math.NewDecFromInt64(0)
	moduleAddrString := k.moduleAddress.String()
	for _, issuance := range req.Issuance {
		decs, err := utils.GetNonNegativeFixedDecs(maxDecimalPlaces, issuance.TradableAmount, issuance.RetiredAmount)
		if err != nil {
			return nil, err
		}
		tradable, retired := decs[0], decs[1]

		tradableString := tradable.String()
		retiredString := retired.String()

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
				Owner:        issuance.Recipient,
				BatchDenom:   batchDenom,
				Amount:       retiredString,
				Jurisdiction: issuance.RetirementJurisdiction,
			}); err != nil {
				return nil, err
			}
		}
		if err = k.stateStore.BatchBalanceTable().Insert(ctx, &api.BatchBalance{
			BatchKey: key,
			Address:  recipient,
			TradableAmount: tradableString,
			RetiredAmount:  retiredString,
		}); err != nil {
			return nil, err
		}

		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventMint{
			BatchDenom:     batchDenom,
			RetiredAmount:  retiredString,
			TradableAmount: tradableString,
		}); err != nil {
			return nil, err
		}

		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventTransfer{
			Sender:         moduleAddrString, // ecocredit module
			Recipient:      issuance.Recipient,
			BatchDenom:     batchDenom,
			RetiredAmount:  retiredString,
			TradableAmount: tradableString,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgCreateBatch issuance iteration")
	}

	if err = k.stateStore.BatchSupplyTable().Insert(ctx, &api.BatchSupply{
		BatchKey:        key,
		TradableAmount:  tradableSupply.String(),
		RetiredAmount:   retiredSupply.String(),
		CancelledAmount: "0", // must be explicitly set to prevent zero-value ""
	}); err != nil {
		return nil, err
	}

	if req.OriginTx != nil {
		if err = k.stateStore.BatchOriginTxTable().Insert(ctx, &api.BatchOriginTx{
			Id:         req.OriginTx.Id,
			Source:     req.OriginTx.Source,
			Note:       req.Note,
			BatchDenom: batchDenom,
		}); err != nil {
			if ormerrors.PrimaryKeyConstraintViolation.Is(err) {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf("credits already issued with tx id: %s", req.OriginTx.Id)
			}
			return nil, err
		}
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventCreateBatch{
		BatchDenom: batchDenom,
		OriginTx:   req.OriginTx,
	}); err != nil {
		return nil, err
	}

	return &core.MsgCreateBatchResponse{BatchDenom: batchDenom}, nil
}

// getBatchSeqNo gets the batch sequence number
func (k Keeper) getBatchSeqNo(ctx context.Context, projectKey uint64) (uint64, error) {
	var seq uint64 = 1
	batchSeq, err := k.stateStore.BatchSequenceTable().Get(ctx, projectKey)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return 0, err
		}
	} else {
		seq = batchSeq.NextSequence
	}

	if err = k.stateStore.BatchSequenceTable().Save(ctx, &api.BatchSequence{
		ProjectKey:   projectKey,
		NextSequence: seq + 1,
	}); err != nil {
		return 0, err
	}

	return seq, err
}
