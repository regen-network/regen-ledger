package core

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	basketv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1beta1"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
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
	regenCtx := types.UnwrapSDKContext(ctx)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// Charge the admin a fee to create the credit class
	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	var params ecocredit.Params
	s.paramSpace.GetParamSet(regenCtx.Context, &params)
	if params.AllowlistEnabled && !s.isCreatorAllowListed(regenCtx, params.AllowedClassCreators, adminAddress) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("%s is not allowed to create credit classes", adminAddress.String())
	}

	err = s.chargeCreditClassFee(regenCtx.Context, adminAddress)
	if err != nil {
		return nil, err
	}

	creditType, err := s.getCreditType(sdkCtx, req.CreditTypeName)
	if err != nil {
		return nil, err
	}

	classSeq, err := s.getClassSequenceNo(ctx, req.CreditTypeName)
	if err != nil {
		return nil, fmt.Errorf("error getting class sequence")
	}
	classID := ecocredit.FormatClassID(creditType.Abbreviation, classSeq)

	rowID, err := s.stateStore.ClassInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.ClassInfo{
		Name:       classID,
		Admin:      req.Admin,
		Metadata:   req.Metadata,
		CreditType: req.CreditTypeName,
	})
	if err != nil {
		return nil, err
	}

	for _, issuer := range req.Issuers {
		if err = s.stateStore.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: classID,
			Issuer:  issuer,
		}); err != nil {
			return nil, err
		}
	}

	err = regenCtx.EventManager().EmitTypedEvent(&v1beta1.EventCreateClass{
		RowId:   rowID,
		ClassId: classID,
		Admin:   req.Admin,
	})
	if err != nil {
		return nil, err
	}

	return &v1beta1.MsgCreateClassResponse{RowId: rowID, ClassId: classID}, nil
}

// CreateProject creates a new project.
func (s serverImpl) CreateProject(ctx context.Context, req *v1beta1.MsgCreateProject) (*v1beta1.MsgCreateProjectResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	classID := req.ClassId
	classInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, classID)
	if err != nil {
		return nil, err
	}

	err = s.assertClassIssuer(ctx, classID, req.Issuer)
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

	rowID, err := s.stateStore.ProjectInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            projectID,
		ClassId:         classInfo.Id,
		ProjectLocation: req.ProjectLocation,
		Metadata:        req.Metadata,
	})
	if err != nil {
		return nil, err
	}

	if err := sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventCreateProject{
		RowId:           rowID,
		ClassId:         classID,
		ProjectId:       projectID,
		Issuer:          req.Issuer,
		ProjectLocation: req.ProjectLocation,
	}); err != nil {
		return nil, err
	}

	return &v1beta1.MsgCreateProjectResponse{
		RowId:     rowID,
		ProjectId: projectID,
	}, nil
}

// CreateBatch creates a new batch of credits.
// Credits in the batch must not have more decimal places than the credit type's specified precision.
func (s serverImpl) CreateBatch(ctx context.Context, req *v1beta1.MsgCreateBatch) (*v1beta1.MsgCreateBatchResponse, error) {
	regenCtx := types.UnwrapSDKContext(ctx)
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

	err = s.assertClassIssuer(ctx, classInfo.Name, req.Issuer)
	if err != nil {
		return nil, err
	}

	creditType, err := s.getCreditType(sdkCtx, classInfo.CreditType)
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
	newBatchID := rowID

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
			if err = regenCtx.EventManager().EmitTypedEvent(&v1beta1.EventRetire{
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
			BatchId:  newBatchID,
			Tradable: tradable.String(),
			Retired:  retired.String(),
		}); err != nil {
			return nil, err
		}

		if err = regenCtx.EventManager().EmitTypedEvent(&v1beta1.EventReceive{
			Recipient:      recipient.String(),
			BatchDenom:     batchDenom,
			RetiredAmount:  tradable.String(),
			TradableAmount: retired.String(),
		}); err != nil {
			return nil, err
		}

		regenCtx.GasMeter().ConsumeGas(gasCostPerIteration, "batch issuance")
	}

	if err = s.stateStore.BatchSupplyStore().Insert(ctx, &ecocreditv1beta1.BatchSupply{
		BatchId:         newBatchID,
		TradableAmount:  tradableSupply.String(),
		RetiredAmount:   retiredSupply.String(),
		CancelledAmount: math.NewDecFromInt64(0).String(),
	}); err != nil {
		return nil, err
	}

	return &v1beta1.MsgCreateBatchResponse{RowId: rowID, BatchDenom: batchDenom}, nil
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
	if classInfo.Admin != req.Admin {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}
	classInfo.Admin = req.NewAdmin
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
	if class.Admin != req.Admin {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", class.Admin, req.Admin)
	}

	// delete the old issuers
	if err = s.stateStore.ClassIssuerStore().DeleteBy(ctx, ecocreditv1beta1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(class.Name)); err != nil {
		return nil, err
	}

	// add the new issuers
	for _, issuer := range req.Issuers {
		if err = s.stateStore.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: req.ClassId,
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
	if classInfo.Admin != req.Admin {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("expected admin %s, got %s", classInfo.Admin, req.Admin)
	}
	classInfo.Metadata = req.Metadata
	if err = s.stateStore.ClassInfoStore().Update(ctx, classInfo); err != nil {
		return nil, err
	}
	return &v1beta1.MsgUpdateClassMetadataResponse{}, err
}

// Sell creates new sell orders for credits
// TODO: update this function with ORM
func (s serverImpl) Sell(ctx context.Context, sell *marketplacev1beta1.MsgSell) (*marketplacev1beta1.MsgSellResponse, error) {
	panic("implement me")
	//ctx := types.UnwrapSDKContext(goCtx)
	//owner := req.Owner
	//store := ctx.KVStore(s.storeKey)
	//
	//ownerAddr, err := sdk.AccAddressFromBech32(owner)
	//if err != nil {
	//	return nil, err
	//}
	//
	//sellOrderIds := make([]uint64, len(req.Orders))
	//
	//for i, order := range req.Orders {
	//
	//	// verify expiration is in the future
	//	if order.Expiration != nil && order.Expiration.Before(ctx.BlockTime()) {
	//		return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
	//	}
	//
	//	err = verifyCreditBalance(store, ownerAddr, order.BatchDenom, order.Quantity)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// TODO: Verify that AskPrice.Denom is in AllowAskDenom #624
	//
	//	orderID, err := s.createSellOrder(ctx, owner, order)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	sellOrderIds[i] = orderID
	//	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventSell{
	//		OrderId:           orderID,
	//		BatchDenom:        order.BatchDenom,
	//		Quantity:          order.Quantity,
	//		AskPrice:          order.AskPrice,
	//		DisableAutoRetire: order.DisableAutoRetire,
	//		Expiration:        order.Expiration,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	ctx.GasMeter().ConsumeGas(gasCostPerIteration, "create sell order")
	//}
	//
	//return &ecocredit.MsgSellResponse{SellOrderIds: sellOrderIds}, nil
}

// TODO: implement this with ORM
func (s serverImpl) createSellOrder(ctx types.Context, owner string, o *marketplacev1beta1.MsgSell_Order) (uint64, error) {
	panic("impl me!")
	//orderID := s.sellOrderTable.Sequence().PeekNextVal(ctx)
	//_, err := s.sellOrderTable.Create(ctx, &ecocredit.SellOrder{
	//	Owner:             owner,
	//	OrderId:           orderID,
	//	BatchDenom:        o.BatchDenom,
	//	Quantity:          o.Quantity,
	//	AskPrice:          o.AskPrice,
	//	DisableAutoRetire: o.DisableAutoRetire,
	//	Expiration:        o.Expiration,
	//})
	//return orderID, err
}

// UpdateSellOrders updates existing sell orders for credits
// TODO: impl with ORM
func (s serverImpl) UpdateSellOrders(ctx context.Context, orders *marketplacev1beta1.MsgUpdateSellOrders) (*marketplacev1beta1.MsgUpdateSellOrdersResponse, error) {
	panic("implement me")
	//ctx := types.UnwrapSDKContext(goCtx)
	//owner := req.Owner
	//store := ctx.KVStore(s.storeKey)
	//
	//ownerAddr, err := sdk.AccAddressFromBech32(owner)
	//if err != nil {
	//	return nil, err
	//}
	//
	//for _, update := range req.Updates {
	//
	//	// verify expiration is in the future
	//	if update.NewExpiration != nil && update.NewExpiration.Before(ctx.BlockTime()) {
	//		return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", update.NewExpiration)
	//	}
	//
	//	sellOrder, err := s.getSellOrder(ctx, update.SellOrderId)
	//	if err != nil {
	//		return nil, ecocredit.ErrInvalidSellOrder.Wrapf("sell order id %d not found", update.SellOrderId)
	//	}
	//
	//	if req.Owner != sellOrder.Owner {
	//		return nil, sdkerrors.ErrUnauthorized.Wrapf("signer is not the owner of sell order id %d", update.SellOrderId)
	//	}
	//
	//	// TODO: Verify that NewAskPrice.Denom is in AllowAskDenom #624
	//
	//	err = verifyCreditBalance(store, ownerAddr, sellOrder.BatchDenom, update.NewQuantity)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	sellOrder.Quantity = update.NewQuantity
	//	sellOrder.AskPrice = update.NewAskPrice
	//	sellOrder.DisableAutoRetire = update.DisableAutoRetire
	//	sellOrder.Expiration = update.NewExpiration
	//
	//	err = s.sellOrderTable.Update(ctx, sellOrder.OrderId, sellOrder)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventUpdateSellOrder{
	//		Owner:             owner,
	//		SellOrderId:       sellOrder.OrderId,
	//		BatchDenom:        sellOrder.BatchDenom,
	//		NewQuantity:       sellOrder.Quantity,
	//		NewAskPrice:       sellOrder.AskPrice,
	//		DisableAutoRetire: sellOrder.DisableAutoRetire,
	//		NewExpiration:     sellOrder.Expiration,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	ctx.GasMeter().ConsumeGas(gasCostPerIteration, "update sell order")
	//}
	//
	//return &ecocredit.MsgUpdateSellOrdersResponse{}, nil
}

// Buy creates new buy orders for credits
func (s serverImpl) Buy(ctx context.Context, buy *marketplacev1beta1.MsgBuy) (*marketplacev1beta1.MsgBuyResponse, error) {
	panic("implement me")
	//ctx := types.UnwrapSDKContext(goCtx)
	//sdkCtx := sdk.UnwrapSDKContext(goCtx)
	//store := ctx.KVStore(s.storeKey)
	//buyer := req.Buyer
	//
	//buyerAddr, err := sdk.AccAddressFromBech32(buyer)
	//if err != nil {
	//	return nil, err
	//}
	//
	//buyOrderIds := make([]uint64, len(req.Orders))
	//
	//for i, order := range req.Orders {
	//
	//	// verify expiration is in the future
	//	if order.Expiration != nil && order.Expiration.Before(ctx.BlockTime()) {
	//		return nil, sdkerrors.ErrInvalidRequest.Wrapf("expiration must be in the future: %s", order.Expiration)
	//	}
	//
	//	balances := s.bankKeeper.SpendableCoins(sdkCtx, buyerAddr)
	//	bidPrice := order.BidPrice
	//	balanceAmount := balances.AmountOf(bidPrice.Denom)
	//
	//	// TODO: Verify that bidPrice.Denom is in AllowAskDenom #624
	//
	//	// get decimal amount of credits desired for purchase
	//	creditsDesired, err := math.NewPositiveDecFromString(order.Quantity)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// calculate the amount of coin to send for purchase
	//	coinToSend, err := getCoinNeeded(creditsDesired, bidPrice)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// verify buyer has sufficient balance in coin
	//	if balanceAmount.LT(coinToSend.Amount) {
	//		return nil, sdkerrors.ErrInsufficientFunds.Wrapf("insufficient balance: got %s, needed at least: %s", balanceAmount.String(), coinToSend.Amount.String())
	//	}
	//
	//	switch order.Selection.Sum.(type) {
	//	case *ecocredit.MsgBuy_Order_Selection_SellOrderId:
	//
	//		sellOrderId := order.Selection.GetSellOrderId()
	//		sellOrder, err := s.getSellOrder(ctx, sellOrderId)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		sellerAddr, err := sdk.AccAddressFromBech32(sellOrder.Owner)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		// verify bid price and ask price denoms match
	//		if bidPrice.Denom != sellOrder.AskPrice.Denom {
	//			return nil, sdkerrors.ErrInvalidRequest.Wrapf("bid price denom does not match ask price denom: got %s, expected: %s", bidPrice.Denom, sellOrder.AskPrice.Denom)
	//		}
	//
	//		// verify bid price is greater than or equal to ask price
	//		if bidPrice.Amount.LT(sellOrder.AskPrice.Amount) {
	//			return nil, sdkerrors.ErrInvalidRequest.Wrapf("bid price too low: got %s, needed at least: %s", bidPrice.String(), sellOrder.AskPrice.String())
	//		}
	//
	//		// verify seller has sufficient balance in credits
	//		err = verifyCreditBalance(store, sellerAddr, sellOrder.BatchDenom, sellOrder.Quantity)
	//		if err != nil {
	//			return nil, ecocredit.ErrInvalidSellOrder.Wrap(err.Error())
	//		}
	//
	//		// get decimal amount of credits available for purchase
	//		creditsAvailable, err := math.NewDecFromString(sellOrder.Quantity)
	//		if err != nil {
	//			return nil, ecocredit.ErrInvalidSellOrder.Wrap(err.Error())
	//		}
	//
	//		creditsToReceive := creditsDesired
	//
	//		// check if credits desired is more than credits available
	//		if creditsDesired.Cmp(creditsAvailable) == 1 {
	//
	//			// error if partial fill disabled
	//			if order.DisablePartialFill {
	//				return nil, ecocredit.ErrInsufficientFunds.Wrap("sell order does not have sufficient credits to fill the buy order")
	//			}
	//
	//			creditsToReceive = creditsAvailable
	//
	//			// recalculate coinToSend if creditsToReceive is not creditsDesired
	//			coinToSend, err = getCoinNeeded(creditsToReceive, bidPrice)
	//			if err != nil {
	//				return nil, err
	//			}
	//		}
	//
	//		// send coin to the seller account
	//		err = s.bankKeeper.SendCoins(sdkCtx, buyerAddr, sellerAddr, sdk.Coins{coinToSend})
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		// error if auto-retire is required for given sell order
	//		if !sellOrder.DisableAutoRetire && order.DisableAutoRetire {
	//			return nil, ecocredit.ErrInvalidBuyOrder.Wrapf("auto-retire is required for sell order %d", sellOrder.OrderId)
	//		}
	//
	//		// error if auto-retire is required and missing location
	//		if !sellOrder.DisableAutoRetire && order.RetirementLocation == "" {
	//			return nil, ecocredit.ErrInvalidBuyOrder.Wrapf("retirement location is required for sell order %d", sellOrder.OrderId)
	//		}
	//
	//		// declare credit for send message
	//		credit := &ecocredit.MsgSend_SendCredits{
	//			BatchDenom: sellOrder.BatchDenom,
	//		}
	//
	//		// set tradable or retired amount depending on auto-retire
	//		if sellOrder.DisableAutoRetire && order.DisableAutoRetire {
	//			credit.RetiredAmount = "0"
	//			credit.TradableAmount = creditsToReceive.String()
	//		} else {
	//			credit.RetiredAmount = creditsToReceive.String()
	//			credit.RetirementLocation = order.RetirementLocation
	//			credit.TradableAmount = "0"
	//		}
	//
	//		// send credits to the buyer account
	//		err = s.sendEcocredits(ctx, credit, store, sellerAddr, buyerAddr)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		// get remaining credits in sell order
	//		creditsRemaining, err := creditsAvailable.Sub(creditsToReceive)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		if creditsRemaining.IsZero() {
	//
	//			// delete sell order if no remaining credits
	//			if err := s.sellOrderTable.Delete(ctx, sellOrder.OrderId); err != nil {
	//				return nil, err
	//			}
	//
	//		} else {
	//			sellOrder.Quantity = creditsRemaining.String()
	//
	//			// update sell order quantity with remaining credits
	//			err = s.sellOrderTable.Update(ctx, sellOrder.OrderId, sellOrder)
	//			if err != nil {
	//				return nil, err
	//			}
	//		}
	//
	//		// TODO: do we want to store a direct buy order? #623
	//		buyOrderID := s.buyOrderTable.Sequence().NextVal(ctx)
	//		buyOrderIds[i] = buyOrderID
	//
	//		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventBuyOrderCreated{
	//			BuyOrderId:         buyOrderID,
	//			SellOrderId:        sellOrderId,
	//			Quantity:           order.Quantity,
	//			BidPrice:           order.BidPrice,
	//			DisableAutoRetire:  order.DisableAutoRetire,
	//			DisablePartialFill: order.DisablePartialFill,
	//			RetirementLocation: order.RetirementLocation,
	//			Expiration:         order.Expiration,
	//		})
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventBuyOrderFilled{
	//			BuyOrderId:  buyOrderID,
	//			SellOrderId: sellOrderId,
	//			BatchDenom:  sellOrder.BatchDenom,
	//			Quantity:    creditsToReceive.String(),
	//			TotalPrice:  &coinToSend,
	//		})
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//	// TODO: implement processing for filter option #623
	//	//case *ecocredit.MsgBuy_Order_Selection_Filter:
	//
	//	default:
	//		return nil, sdkerrors.ErrInvalidRequest
	//	}
	//
	//	ctx.GasMeter().ConsumeGas(gasCostPerIteration, "create buy order")
	//}
	//
	//return &ecocredit.MsgBuyResponse{BuyOrderIds: buyOrderIds}, nil
}

// TODO: impl with ORM
func (s serverImpl) createBuyOrder(ctx types.Context, buyer string, o *marketplacev1beta1.MsgBuy_Order) (uint64, error) {
	panic("impl me!")
	//orderID := s.buyOrderTable.Sequence().PeekNextVal(ctx)
	//selection := ecocredit.BuyOrder_Selection{
	//	Sum: &ecocredit.BuyOrder_Selection_SellOrderId{
	//		SellOrderId: o.Selection.GetSellOrderId(),
	//	},
	//}
	//_, err := s.buyOrderTable.Create(ctx, &ecocredit.BuyOrder{
	//	Buyer:              buyer,
	//	BuyOrderId:         orderID,
	//	Selection:          &selection,
	//	Quantity:           o.Quantity,
	//	BidPrice:           o.BidPrice,
	//	DisableAutoRetire:  o.DisableAutoRetire,
	//	DisablePartialFill: o.DisablePartialFill,
	//	Expiration:         o.Expiration,
	//})
	//return orderID, err
}

// AllowAskDenom adds a new ask denom
// TODO: impl with ORM
func (s serverImpl) AllowAskDenom(ctx context.Context, denom *marketplacev1beta1.MsgAllowAskDenom) (*marketplacev1beta1.MsgAllowAskDenomResponse, error) {
	panic("implement me")
	// ctx := types.UnwrapSDKContext(goCtx)
	//
	//rootAddress := s.accountKeeper.GetModuleAddress(govtypes.ModuleName).String()
	//
	//if req.RootAddress != rootAddress {
	//	return nil, sdkerrors.ErrUnauthorized.Wrapf("root address must be governance module address, got: %s, expected: %s", req.RootAddress, rootAddress)
	//}
	//
	//err := s.askDenomTable.Create(ctx, &ecocredit.AskDenom{
	//	Denom:        req.Denom,
	//	DisplayDenom: req.DisplayDenom,
	//	Exponent:     req.Exponent,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventAllowAskDenom{
	//	Denom:        req.Denom,
	//	DisplayDenom: req.DisplayDenom,
	//	Exponent:     req.Exponent,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ecocredit.MsgAllowAskDenomResponse{}, nil
}

func (s serverImpl) CreateBasket(ctx context.Context, basket *basketv1beta1.MsgCreateBasket) (*basketv1beta1.MsgCreateBasketResponse, error) {
	panic("implement me")
}

func (s serverImpl) AddToBasket(ctx context.Context, basket *basketv1beta1.MsgAddToBasket) (*basketv1beta1.MsgAddToBasketResponse, error) {
	panic("implement me")
}

func (s serverImpl) TakeFromBasket(ctx context.Context, basket *basketv1beta1.MsgTakeFromBasket) (*basketv1beta1.MsgTakeFromBasketResponse, error) {
	panic("implement me")
}

func (s serverImpl) PickFromBasket(ctx context.Context, basket *basketv1beta1.MsgPickFromBasket) (*basketv1beta1.MsgPickFromBasketResponse, error) {
	panic("implement me")
}

// ------- UTILITIES ------

// Checks if the given address is in the allowlist of credit class creators
func (s serverImpl) isCreatorAllowListed(ctx types.Context, allowlist []string, designer sdk.Address) bool {
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
func (s serverImpl) assertClassIssuer(goCtx context.Context, classID, issuer string) error {
	it, err := s.stateStore.ClassIssuerStore().List(goCtx, ecocreditv1beta1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(classID))
	if err != nil {
		return err
	}

	defer it.Close()
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return err
		}
		if v.Issuer == issuer {
			return nil
		}
	}
	return sdkerrors.ErrUnauthorized.Wrapf("%s is not an issuer for class %s", issuer, classID)
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

func (s serverImpl) getCreditTypeFromBatchDenom(ctx context.Context, denom string) (v1beta1.CreditType, error) {
	classId := ecocredit.GetClassIdFromBatchDenom(denom)
	classInfo, err := s.stateStore.ClassInfoStore().GetByName(ctx, classId)
	if err != nil {
		return v1beta1.CreditType{}, err
	}
	return s.getCreditType(sdk.UnwrapSDKContext(ctx), classInfo.CreditType)
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

func (s serverImpl) getCreditType(ctx sdk.Context, creditTypeName string) (v1beta1.CreditType, error) {
	creditTypes := s.getAllCreditTypes(ctx)
	creditTypeName = ecocredit.NormalizeCreditTypeName(creditTypeName)
	for _, creditType := range creditTypes {
		// credit type name's stored via params have enforcement on normalization, so we can be sure they will already
		// be normalized here.
		if creditType.Name == creditTypeName {
			return *creditType, nil
		}
	}
	return v1beta1.CreditType{}, sdkerrors.ErrInvalidType.Wrapf("%s is not a valid credit type", creditTypeName)
}

func (s serverImpl) getAllCreditTypes(ctx sdk.Context) []*v1beta1.CreditType {
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx, &params)
	creditTypes := params.CreditTypes
	v1beta1types := make([]*v1beta1.CreditType, len(creditTypes))
	for i, typ := range creditTypes {
		v1beta1types[i] = &v1beta1.CreditType{
			Abbreviation: typ.Abbreviation,
			Name:         typ.Name,
			Unit:         typ.Unit,
			Precision:    typ.Precision,
		}
	}
	return v1beta1types
}

func (s serverImpl) getCreditClassFee(ctx sdk.Context) sdk.Coins {
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx, &params)
	return params.CreditClassFee
}

func (s serverImpl) chargeCreditClassFee(ctx sdk.Context, creatorAddr sdk.AccAddress) error {
	creditClassFee := s.getCreditClassFee(ctx)

	// Move the fee to the ecocredit module's account
	err := s.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	// Burn the coins
	// TODO: Update this implementation based on the discussion at
	// https://github.com/regen-network/regen-ledger/issues/351
	err = s.bankKeeper.BurnCoins(ctx, ecocredit.ModuleName, creditClassFee)
	if err != nil {
		return err
	}

	return nil
}
