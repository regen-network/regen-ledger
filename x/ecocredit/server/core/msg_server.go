package core

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

// CreateClass creates a new class of ecocredit
//
// The admin is charged a fee for creating the class. This is controlled by
// the global parameter CreditClassFee, which can be updated through the
// governance process.
func (s serverImpl) CreateClass(ctx context.Context, req *v1beta1.MsgCreateClass) (*v1beta1.MsgCreateClassResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// Charge the admin a fee to create the credit class
	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	if err := s.checkAllowList(ctx, adminAddress); err != nil {
		return nil, err
	}

	if err = s.chargeCreditClassFee(ctx, adminAddress, req.FeeDenom); err != nil {
		return nil, err
	}

	creditType, err := s.getCreditType(ctx, req.CreditTypeName)
	if err != nil {
		return nil, err
	}

	classSeq, err := s.getClassSequenceNo(ctx, req.CreditTypeName)
	if err != nil {
		return nil, fmt.Errorf("error getting class sequence")
	}
	classID := ecocredit.FormatClassID(creditType.Abbreviation, classSeq)

	rowId, err := s.stateStore.ClassInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.ClassInfo{
		Name:       classID,
		Admin:      adminAddress,
		Metadata:   req.Metadata,
		CreditType: req.CreditTypeName,
	})
	if err != nil {
		return nil, err
	}

	for _, issuer := range req.Issuers {
		issuer, _ := sdk.AccAddressFromBech32(issuer)
		if err = s.stateStore.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: rowId,
			Issuer:  issuer,
		}); err != nil {
			return nil, err
		}
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventCreateClass{
		ClassId: classID,
		Admin:   req.Admin,
	})
	if err != nil {
		return nil, err
	}

	return &v1beta1.MsgCreateClassResponse{ClassId: classID}, nil
}

// CreateProject creates a new project.
func (s serverImpl) CreateProject(ctx context.Context, req *v1beta1.MsgCreateProject) (*v1beta1.MsgCreateProjectResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	classID := req.ClassId
	classInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, classID)
	if err != nil {
		return nil, err
	}

	err = s.assertClassIssuer(ctx, classInfo.Id, req.Issuer)
	if err != nil {
		return nil, err
	}

	projectID := req.ProjectId
	if projectID == "" {
		found := false
		for !found {
			projectID, err = s.genProjectID(ctx, classInfo.Id, classInfo.Name)
			if err != nil {
				return nil, err
			}
			found, err = s.stateStore.ProjectInfoStore().HasByClassIdName(ctx, classInfo.Id, projectID)
			if err != nil {
				return nil, err
			}
			sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "project id sequence")
		}
	}

	if err = s.stateStore.ProjectInfoStore().Insert(ctx, &ecocreditv1beta1.ProjectInfo{
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

// CreateBatch creates a new batch of credits.
// Credits in the batch must not have more decimal places than the credit type's specified precision.
func (s serverImpl) CreateBatch(ctx context.Context, req *v1beta1.MsgCreateBatch) (*v1beta1.MsgCreateBatchResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	projectID := req.ProjectId

	projectInfo, err := s.stateStore.ProjectInfoStore().GetByName(ctx, projectID)
	if err != nil {
		return nil, err
	}

	classInfo, err := s.stateStore.ClassInfoStore().Get(ctx, projectInfo.ClassId)
	if err != nil {
		return nil, err
	}

	err = s.assertClassIssuer(ctx, classInfo.Id, req.Issuer)
	if err != nil {
		return nil, err
	}

	creditType, err := s.getCreditType(ctx, classInfo.CreditType)
	if err != nil {
		return nil, err
	}

	maxDecimalPlaces := creditType.Precision

	batchSeqNo, err := s.getBatchSeqNo(ctx, projectID)
	if err != nil {
		return nil, err
	}

	batchDenom, err := ecocredit.FormatDenom(classInfo.Name, batchSeqNo, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	rowID, err := s.stateStore.BatchInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  projectInfo.Id,
		BatchDenom: batchDenom,
		Metadata:   req.Metadata,
		StartDate:  timestamppb.New(req.StartDate.UTC()),
		EndDate:    timestamppb.New(req.EndDate.UTC()),
	})
	if err != nil {
		return nil, err
	}

	tradableSupply, retiredSupply := math.NewDecFromInt64(0), math.NewDecFromInt64(0)

	for _, issuance := range req.Issuance {
		decs, err := getNonNegativeFixedDecs(maxDecimalPlaces, issuance.TradableAmount, issuance.RetiredAmount)
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
			if err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventRetire{
				Retirer:    recipient.String(),
				BatchDenom: batchDenom,
				Amount:     retired.String(),
				Location:   issuance.RetirementLocation,
			}); err != nil {
				return nil, err
			}
		}
		if err = s.stateStore.BatchBalanceStore().Insert(ctx, &ecocreditv1beta1.BatchBalance{
			Address:  recipient,
			BatchId:  rowID,
			Tradable: tradable.String(),
			Retired:  retired.String(),
		}); err != nil {
			return nil, err
		}

		if err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventReceive{
			Recipient:      recipient.String(),
			BatchDenom:     batchDenom,
			RetiredAmount:  tradable.String(),
			TradableAmount: retired.String(),
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "batch issuance")
	}

	if err = s.stateStore.BatchSupplyStore().Insert(ctx, &ecocreditv1beta1.BatchSupply{
		BatchId:         rowID,
		TradableAmount:  tradableSupply.String(),
		RetiredAmount:   retiredSupply.String(),
		CancelledAmount: math.NewDecFromInt64(0).String(),
	}); err != nil {
		return nil, err
	}

	return &v1beta1.MsgCreateBatchResponse{BatchDenom: batchDenom}, nil
}

// Send sends credits to a recipient.
// Send also retires credits if the amount to retire is specified in the request.
func (s serverImpl) Send(ctx context.Context, req *v1beta1.MsgSend) (*v1beta1.MsgSendResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	sender, _ := sdk.AccAddressFromBech32(req.Sender)
	recipient, _ := sdk.AccAddressFromBech32(req.Recipient)

	for _, credit := range req.Credits {
		err := s.sendEcocredits(ctx, credit, recipient, sender)
		if err != nil {
			return nil, err
		}
		if err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventReceive{
			Sender:         req.Sender,
			Recipient:      req.Recipient,
			BatchDenom:     credit.BatchDenom,
			TradableAmount: credit.TradableAmount,
			RetiredAmount:  credit.RetiredAmount,
		}); err != nil {
			return nil, err
		}
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "send ecocredits")
	}
	return &v1beta1.MsgSendResponse{}, nil
}

// Retire credits to the specified location.
// WARNING: retiring credits is permanent. Retired credits cannot be un-retired.
func (s serverImpl) Retire(ctx context.Context, req *v1beta1.MsgRetire) (*v1beta1.MsgRetireResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	holder, _ := sdk.AccAddressFromBech32(req.Holder)

	for _, credit := range req.Credits {
		batch, err := s.stateStore.BatchInfoStore().GetByBatchDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, err
		}
		creditType, err := s.getCreditTypeFromBatchDenom(ctx, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		userBalance, err := s.stateStore.BatchBalanceStore().Get(ctx, holder, batch.Id)
		if err != nil {
			return nil, err
		}

		decs, err := getNonNegativeFixedDecs(creditType.Precision, credit.Amount, userBalance.Tradable)
		if err != nil {
			return nil, err
		}
		amtToRetire, userTradableBalance := decs[0], decs[1]

		userTradableBalance, err = userTradableBalance.Sub(amtToRetire)
		if err != nil {
			return nil, err
		}
		if userTradableBalance.IsNegative() {
			return nil, ecocredit.ErrInsufficientFunds.Wrapf("cannot retire %s credits with a balance of %s", credit.Amount, userBalance.Tradable)
		}
		userRetiredBalance, err := math.NewNonNegativeFixedDecFromString(userBalance.Retired, creditType.Precision)
		if err != nil {
			return nil, err
		}
		userRetiredBalance, err = userRetiredBalance.Add(amtToRetire)
		if err != nil {
			return nil, err
		}
		batchSupply, err := s.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
		if err != nil {
			return nil, err
		}
		decs, err = getNonNegativeFixedDecs(creditType.Precision, batchSupply.RetiredAmount, batchSupply.TradableAmount)
		if err != nil {
			return nil, err
		}
		supplyRetired, supplyTradable := decs[0], decs[1]
		supplyRetired, err = supplyRetired.Add(amtToRetire)
		if err != nil {
			return nil, err
		}
		supplyTradable, err = supplyTradable.Sub(amtToRetire)
		if err != nil {
			return nil, err
		}

		if err = s.stateStore.BatchBalanceStore().Update(ctx, &ecocreditv1beta1.BatchBalance{
			Address:  holder,
			BatchId:  batch.Id,
			Tradable: userTradableBalance.String(),
			Retired:  userRetiredBalance.String(),
		}); err != nil {
			return nil, err
		}
		err = s.stateStore.BatchSupplyStore().Update(ctx, &ecocreditv1beta1.BatchSupply{
			BatchId:         batch.Id,
			TradableAmount:  supplyTradable.String(),
			RetiredAmount:   supplyRetired.String(),
			CancelledAmount: batchSupply.CancelledAmount,
		})
		if err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventRetire{
			Retirer:    req.Holder,
			BatchDenom: credit.BatchDenom,
			Amount:     credit.Amount,
			Location:   req.Location,
		}); err != nil {
			return nil, err
		}
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "retire ecocredits")
	}
	return &v1beta1.MsgRetireResponse{}, nil
}

// Cancel credits, removing them from the supply and balance of the holder
func (s serverImpl) Cancel(ctx context.Context, req *v1beta1.MsgCancel) (*v1beta1.MsgCancelResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	holder, _ := sdk.AccAddressFromBech32(req.Holder)

	for _, credit := range req.Credits {
		batch, err := s.stateStore.BatchInfoStore().GetByBatchDenom(ctx, credit.BatchDenom)
		if err != nil {
			return nil, err

		}
		creditType, err := s.getCreditTypeFromBatchDenom(ctx, batch.BatchDenom)
		if err != nil {
			return nil, err
		}
		precision := creditType.Precision

		userBalance, err := s.stateStore.BatchBalanceStore().Get(ctx, holder, batch.Id)
		if err != nil {
			return nil, err
		}
		batchSupply, err := s.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
		if err != nil {
			return nil, err
		}
		decs, err := getNonNegativeFixedDecs(precision, credit.Amount, batchSupply.TradableAmount, userBalance.Tradable, batchSupply.CancelledAmount)
		if err != nil {
			return nil, err
		}
		amtToCancelDec, supplyTradable, userBalTradable, cancelledDec := decs[0], decs[1], decs[2], decs[3]
		userBalTradable, err = math.SafeSubBalance(userBalTradable, amtToCancelDec)
		if err != nil {
			return nil, err
		}
		supplyTradable, err = math.SafeSubBalance(supplyTradable, amtToCancelDec)
		if err != nil {
			return nil, err
		}
		cancelledDec, err = cancelledDec.Add(amtToCancelDec)
		if err != nil {
			return nil, err
		}
		if err = s.stateStore.BatchBalanceStore().Update(ctx, &ecocreditv1beta1.BatchBalance{
			Address:  holder,
			BatchId:  batch.Id,
			Tradable: userBalTradable.String(),
			Retired:  userBalance.Retired,
		}); err != nil {
			return nil, err
		}
		if err = s.stateStore.BatchSupplyStore().Update(ctx, &ecocreditv1beta1.BatchSupply{
			BatchId:         batch.Id,
			TradableAmount:  supplyTradable.String(),
			RetiredAmount:   batchSupply.RetiredAmount,
			CancelledAmount: cancelledDec.String(),
		}); err != nil {
			return nil, err
		}
		if err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventCancel{
			Canceller:  holder.String(),
			BatchDenom: credit.BatchDenom,
			Amount:     credit.Amount,
		}); err != nil {
			return nil, err
		}
		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "cancel ecocredits")
	}
	return &v1beta1.MsgCancelResponse{}, nil
}

func (s serverImpl) UpdateClassAdmin(ctx context.Context, req *v1beta1.MsgUpdateClassAdmin) (*v1beta1.MsgUpdateClassAdminResponse, error) {
	classInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}
	reqAddr, _ := sdk.AccAddressFromBech32(req.Admin)
	classAdmin := sdk.AccAddress(classInfo.Admin)
	if !classAdmin.Equals(reqAddr) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}
	classInfo.Admin = reqAddr
	if err = s.stateStore.ClassInfoStore().Update(ctx, classInfo); err != nil {
		return nil, err
	}
	return &v1beta1.MsgUpdateClassAdminResponse{}, err
}

func (s serverImpl) UpdateClassIssuers(ctx context.Context, req *v1beta1.MsgUpdateClassIssuers) (*v1beta1.MsgUpdateClassIssuersResponse, error) {
	class, err := s.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}
	reqAddr, _ := sdk.AccAddressFromBech32(req.Admin)
	admin := sdk.AccAddress(class.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", class.Admin, req.Admin)
	}

	for _, issuer := range req.RemoveIssuers {
		if err = s.stateStore.ClassIssuerStore().Delete(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: class.Id,
			Issuer:  issuer,
		}); err != nil {
			return nil, err
		}
	}

	// add the new issuers
	for _, issuer := range req.AddIssuers {
		if err = s.stateStore.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: class.Id,
			Issuer:  issuer,
		}); err != nil {
			return nil, err
		}
	}
	return &v1beta1.MsgUpdateClassIssuersResponse{}, nil
}

func (s serverImpl) UpdateClassMetadata(ctx context.Context, req *v1beta1.MsgUpdateClassMetadata) (*v1beta1.MsgUpdateClassMetadataResponse, error) {
	classInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}
	reqAddr, _ := sdk.AccAddressFromBech32(req.Admin)
	admin := sdk.AccAddress(classInfo.Admin)
	if !reqAddr.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}
	classInfo.Metadata = req.Metadata
	if err = s.stateStore.ClassInfoStore().Update(ctx, classInfo); err != nil {
		return nil, err
	}
	return &v1beta1.MsgUpdateClassMetadataResponse{}, err
}

// ------- UTILITIES ------

// Checks if the given address is in the allowlist of credit class creators
func (s serverImpl) isCreatorAllowListed(ctx sdk.Context, allowlist []string, designer sdk.Address) bool {
	for _, addr := range allowlist {
		ctx.GasMeter().ConsumeGas(gasCostPerIteration, "credit class creators allowlist")
		allowListedAddr, _ := sdk.AccAddressFromBech32(addr)
		if designer.Equals(allowListedAddr) {
			return true
		}
	}
	return false
}

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (s serverImpl) assertClassIssuer(goCtx context.Context, classID uint64, issuer string) error {
	addr, _ := sdk.AccAddressFromBech32(issuer)
	found, err := s.stateStore.ClassIssuerStore().Has(goCtx, classID, addr)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.ErrUnauthorized.Wrapf("%s is not an issuer for class %s", issuer, classID)
	}
	return nil
}

func (s serverImpl) genProjectID(ctx context.Context, classRowID uint64, classID string) (string, error) {
	var nextID uint64
	projectSeqNo, err := s.stateStore.ProjectSequenceStore().Get(ctx, classRowID)
	switch err {
	case ormerrors.NotFound:
		nextID = 1
	case nil:
		nextID = projectSeqNo.NextProjectId
	default:
		return "", err
	}

	if err = s.stateStore.ProjectSequenceStore().Save(ctx, &ecocreditv1beta1.ProjectSequence{
		ClassId:       classRowID,
		NextProjectId: nextID + 1,
	}); err != nil {
		return "", err
	}

	return ecocredit.FormatProjectID(classID, nextID), nil
}

func (s serverImpl) getBatchSeqNo(ctx context.Context, projectID string) (uint64, error) {
	var seq uint64
	batchSeq, err := s.stateStore.BatchSequenceStore().Get(ctx, projectID)

	switch err {
	case ormerrors.NotFound:
		seq = 1
	case nil:
		seq = batchSeq.NextBatchId
	default:
		return 0, err
	}
	if err = s.stateStore.BatchSequenceStore().Save(ctx, &ecocreditv1beta1.BatchSequence{
		ProjectId:   projectID,
		NextBatchId: seq + 1,
	}); err != nil {
		return 0, err
	}

	return seq, err
}

func (s serverImpl) sendEcocredits(ctx context.Context, credit *v1beta1.MsgSend_SendCredits, to, from sdk.AccAddress) error {
	batch, err := s.stateStore.BatchInfoStore().GetByBatchDenom(ctx, credit.BatchDenom)
	if err != nil {
		return err
	}
	creditType, err := s.getCreditTypeFromBatchDenom(ctx, batch.BatchDenom)
	if err != nil {
		return err
	}
	precision := creditType.Precision

	batchSupply, err := s.stateStore.BatchSupplyStore().Get(ctx, batch.Id)
	if err != nil {
		return err
	}
	fromBalance, err := s.stateStore.BatchBalanceStore().Get(ctx, from, batch.Id)
	if err != nil {
		if err == ormerrors.NotFound {
			return ecocredit.ErrInsufficientFunds.Wrapf("you do not have any credits from batch %s", batch.BatchDenom)
		}
		return err
	}

	toBalance, err := s.stateStore.BatchBalanceStore().Get(ctx, to, batch.Id)
	if err != nil {
		if err == ormerrors.NotFound {
			toBalance = &ecocreditv1beta1.BatchBalance{
				Address:  to,
				BatchId:  batch.Id,
				Tradable: "0",
				Retired:  "0",
			}
		} else {
			return err
		}
	}
	decs, err := getNonNegativeFixedDecs(precision, toBalance.Tradable, toBalance.Retired, fromBalance.Tradable, fromBalance.Retired, credit.TradableAmount, credit.RetiredAmount, batchSupply.TradableAmount, batchSupply.RetiredAmount)
	if err != nil {
		return err
	}
	toTradableBalance, toRetiredBalance,
		fromTradableBalance, fromRetiredBalance,
		sendAmtTradable, sendAmtRetired,
		batchSupplyTradable, batchSupplyRetired := decs[0], decs[1], decs[2], decs[3], decs[4], decs[5], decs[6], decs[7]

	if !sendAmtTradable.IsZero() {
		fromTradableBalance, err = math.SafeSubBalance(fromTradableBalance, sendAmtTradable)
		if err != nil {
			return err
		}
		toTradableBalance, err = toTradableBalance.Add(sendAmtTradable)
		if err != nil {
			return err
		}
	}

	didRetire := false
	if !sendAmtRetired.IsZero() {
		didRetire = true
		fromTradableBalance, err = math.SafeSubBalance(fromTradableBalance, sendAmtRetired)
		if err != nil {
			return err
		}
		toRetiredBalance, err = toRetiredBalance.Add(sendAmtRetired)
		if err != nil {
			return err
		}
		batchSupplyRetired, err = batchSupplyRetired.Add(sendAmtRetired)
		if err != nil {
			return err
		}
		batchSupplyTradable, err = batchSupplyTradable.Sub(sendAmtRetired)
		if err != nil {
			return err
		}
	}
	// update the "to" balance
	if err := s.stateStore.BatchBalanceStore().Save(ctx, &ecocreditv1beta1.BatchBalance{
		Address:  to,
		BatchId:  batch.Id,
		Tradable: toTradableBalance.String(),
		Retired:  toRetiredBalance.String(),
	}); err != nil {
		return err
	}

	// update the "from" balance
	if err := s.stateStore.BatchBalanceStore().Update(ctx, &ecocreditv1beta1.BatchBalance{
		Address:  from,
		BatchId:  batch.Id,
		Tradable: fromTradableBalance.String(),
		Retired:  fromRetiredBalance.String(),
	}); err != nil {
		return err
	}
	// update the "retired" balance only if credits were retired
	if didRetire {
		if err := s.stateStore.BatchSupplyStore().Update(ctx, &ecocreditv1beta1.BatchSupply{
			BatchId:         batch.Id,
			TradableAmount:  batchSupplyTradable.String(),
			RetiredAmount:   batchSupplyRetired.String(),
			CancelledAmount: batchSupply.CancelledAmount,
		}); err != nil {
			return err
		}
		if err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&ecocredit.EventRetire{
			Retirer:    to.String(),
			BatchDenom: credit.BatchDenom,
			Amount:     sendAmtRetired.String(),
			Location:   credit.RetirementLocation,
		}); err != nil {
			return err
		}
	}
	return nil
}

func getNonNegativeFixedDecs(precision uint32, decimals ...string) ([]math.Dec, error) {
	decs := make([]math.Dec, len(decimals))
	for i, decimal := range decimals {
		dec, err := math.NewNonNegativeFixedDecFromString(decimal, precision)
		if err != nil {
			return nil, err
		}
		decs[i] = dec
	}
	return decs, nil
}

func (s serverImpl) getCreditTypeFromBatchDenom(ctx context.Context, denom string) (*ecocreditv1beta1.CreditType, error) {
	classId := ecocredit.GetClassIdFromBatchDenom(denom)
	classInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, classId)
	if err != nil {
		return nil, err
	}
	return s.getCreditType(ctx, classInfo.CreditType)
}

func (s serverImpl) getClassSequenceNo(ctx context.Context, ctype string) (uint64, error) {
	var seq uint64
	classSeq, err := s.stateStore.ClassSequenceStore().Get(ctx, ctype)
	switch err {
	case nil:
		seq = classSeq.NextClassId
	case ormerrors.NotFound:
		seq = 1
	default:
		return 0, err
	}
	err = s.stateStore.ClassSequenceStore().Save(ctx, &ecocreditv1beta1.ClassSequence{
		CreditType:  ctype,
		NextClassId: seq + 1,
	})
	return seq, err
}

func (s serverImpl) getCreditType(ctx context.Context, creditTypeName string) (*ecocreditv1beta1.CreditType, error) {
	ct, err := s.stateStore.CreditTypeStore().GetByName(ctx, creditTypeName)
	return ct, err
}

func (s serverImpl) chargeCreditClassFee(ctx context.Context, creatorAddr sdk.AccAddress, denom string) error {
	fee, err := s.stateStore.CreditClassFeeStore().Get(ctx, denom)
	if err != nil {
		if err == ormerrors.NotFound {
			return sdkerrors.ErrInvalidRequest.Wrapf("cannot use %s to pay for credit class fees", denom)
		} else {
			return err
		}
	}

	amount, ok := sdk.NewIntFromString(fee.Amount)
	if !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid amount", fee.Amount)
	}

	coinFee := sdk.NewCoins(sdk.NewCoin(denom, amount))

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// Move the fee to the ecocredit module's account
	if err = s.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, creatorAddr, ecocredit.ModuleName, coinFee); err != nil {
		return err
	}

	// Burn the coins
	// TODO: Update this implementation based on the discussion at
	// https://github.com/regen-network/regen-ledger/issues/351
	return s.bankKeeper.BurnCoins(sdkCtx, ecocredit.ModuleName, coinFee)
}

func (s serverImpl) checkAllowList(ctx context.Context, address sdk.AccAddress) error {
	res, err := s.stateStore.AllowlistEnabledStore().Get(ctx)
	if err != nil {
		return err
	}
	if !res.Enabled {
		return nil
	}

	found, err := s.stateStore.AllowedClassCreatorsStore().Has(ctx, address)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.ErrUnauthorized.Wrapf("%s is not allowed to create credit classes", address.String())
	}

	return nil
}
