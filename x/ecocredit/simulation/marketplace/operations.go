package marketplace

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

// Simulation operation weights constants
const (
	OpWeightMsgBuy             = "op_weight_msg_buy_direct"        //nolint:gosec
	OpWeightMsgSell            = "op_weight_msg_sell"              //nolint:gosec
	OpWeightMsgUpdateSellOrder = "op_weight_msg_update_sell_order" //nolint:gosec
	OpWeightMsgCancelSellOrder = "op_weight_msg_cancel_sell_order" //nolint:gosec
)

// basket operations weights
const (
	WeightBuyDirect       = 100
	WeightSell            = 100
	WeightUpdateSellOrder = 100
	WeightCancelSellOrder = 100
)

// ecocredit message types
var (
	TypeMsgBuyDirect       = marketplace.MsgBuyDirect{}.Route()
	TypeMsgSell            = marketplace.MsgSell{}.Route()
	TypeMsgUpdateSellOrder = marketplace.MsgUpdateSellOrders{}.Route()
	TypeMsgCancelSellOrder = marketplace.MsgCancelSellOrder{}.Route()
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryServer, mktQryClient marketplace.QueryServer) simulation.WeightedOperations {

	var (
		weightMsgBuyDirect       int
		weightMsgSell            int
		weightMsgUpdateSellOrder int
		weightMsgCancelSellOrder int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBuy, &weightMsgBuyDirect, nil,
		func(_ *rand.Rand) {
			weightMsgBuyDirect = WeightBuyDirect
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSell, &weightMsgSell, nil,
		func(_ *rand.Rand) {
			weightMsgSell = WeightSell
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateSellOrder, &weightMsgUpdateSellOrder, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSellOrder = WeightUpdateSellOrder
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCancelSellOrder, &weightMsgCancelSellOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelSellOrder = WeightCancelSellOrder
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgBuyDirect,
			SimulateMsgBuyDirect(ak, bk, qryClient, mktQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSell,
			SimulateMsgSell(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateSellOrder,
			SimulateMsgUpdateSellOrder(ak, bk, qryClient, mktQryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCancelSellOrder,
			SimulateMsgCancelSellOrder(ak, bk, qryClient, mktQryClient),
		),
	}
}

// SimulateMsgBuyDirect generates a Marketplace/MsgBuyDirect with random values.
func SimulateMsgBuyDirect(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryServer, mktQryClient marketplace.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		buyer, _ := simtypes.RandomAcc(r, accs)
		buyerAddr := buyer.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		result, err := mktQryClient.SellOrders(ctx, &marketplace.QuerySellOrdersRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBuyDirect, err.Error()), nil, err
		}

		sellOrders := result.SellOrders
		if len(sellOrders) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBuyDirect, "no sell orders"), nil, nil
		}

		max := 1
		if len(sellOrders) > 1 {
			max = simtypes.RandIntBetween(r, 1, len(sellOrders))
		}

		buyOrders := make([]*marketplace.MsgBuyDirect_Order, max)
		for i := 0; i < max; i++ {
			sellOrderAskAmount, err := strconv.Atoi(sellOrders[i].AskAmount)
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBuyDirect, "could not convert to int"), nil, nil
			}

			bidPrice := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, sellOrderAskAmount, sellOrderAskAmount+100))))
			buyOrders[i] = &marketplace.MsgBuyDirect_Order{
				SellOrderId:            sellOrders[i].Id,
				Quantity:               sellOrders[i].Quantity,
				BidPrice:               &bidPrice,
				DisableAutoRetire:      sellOrders[i].DisableAutoRetire,
				RetirementJurisdiction: "AQ",
			}
		}

		msg := &marketplace.MsgBuyDirect{
			Buyer:  buyerAddr,
			Orders: buyOrders,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, buyerAddr, TypeMsgBuyDirect)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simapp.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)

	}
}

// SimulateMsgSell generates a Marketplace/MsgSell with random values.
func SimulateMsgSell(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, baseApp *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		seller, _ := simtypes.RandomAcc(r, accs)
		sellerAddr := seller.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgSell)
		if class == nil {
			return op, nil, err
		}

		batchRes, err := qryClient.BatchesByIssuer(ctx, &core.QueryBatchesByIssuerRequest{Issuer: seller.Address.String()})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, err.Error()), nil, err
		}

		batches := batchRes.Batches
		if len(batches) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, "no credit batches"), nil, nil
		}
		max := 1
		if len(batches) > 1 {
			max = simtypes.RandIntBetween(r, 1, len(batches))
		}

		sellOrders := make([]*marketplace.MsgSell_Order, max)
		for i := 0; i < max; i++ {
			bal, err := qryClient.Balance(ctx, &core.QueryBalanceRequest{Address: sellerAddr, BatchDenom: batches[i].Denom})
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, err.Error()), nil, err
			}

			exp := sdkCtx.BlockTime().AddDate(1, 0, 0)
			d, err := math.NewNonNegativeDecFromString(bal.Balance.TradableAmount)
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, err.Error()), nil, err
			}

			if d.IsZero() {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, "no balance"), nil, nil
			}

			balInt, err := d.Int64()
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, err.Error()), nil, nil
			}

			quantity := int(balInt)
			if balInt > 1 {
				quantity = simtypes.RandIntBetween(r, 1, quantity)
			}

			askPrice := sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simtypes.RandIntBetween(r, 1, 50)))
			sellOrders[i] = &marketplace.MsgSell_Order{
				BatchDenom:        batches[i].Denom,
				Quantity:          fmt.Sprintf("%d", quantity),
				AskPrice:          &askPrice,
				DisableAutoRetire: r.Int63n(101) <= 30, // 30% chance of disable auto-retire
				Expiration:        &exp,
			}
		}

		msg := &marketplace.MsgSell{
			Seller: seller.Address.String(),
			Orders: sellOrders,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, sellerAddr, TypeMsgSell)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             baseApp,
			TxGen:           simapp.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgUpdateSellOrder generates a Marketplace/MsgUpdateSellOrder with random values.
func SimulateMsgUpdateSellOrder(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	coreQryClient core.QueryServer, mktQryClient marketplace.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		seller, _ := simtypes.RandomAcc(r, accs)
		sellerAddr := seller.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		result, err := mktQryClient.SellOrdersBySeller(ctx, &marketplace.QuerySellOrdersBySellerRequest{Seller: sellerAddr})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateSellOrder, err.Error()), nil, err
		}

		orders := result.SellOrders
		if len(orders) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateSellOrder, "no sell orders present"), nil, nil
		}

		max := 1
		if len(orders) > 1 {
			max = simtypes.RandIntBetween(r, 1, len(orders))
		}

		updatedOrders := make([]*marketplace.MsgUpdateSellOrders_Update, len(orders))
		for i := 0; i < max; i++ {
			askPrice := sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simtypes.RandIntBetween(r, 1, 50)))
			exp := simtypes.RandTimestamp(r)
			if exp.Before(sdkCtx.BlockTime()) {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateSellOrder, "sell order expiration in the past"), nil, nil
			}
			q, err := strconv.Atoi(orders[i].Quantity)
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateSellOrder, err.Error()), nil, nil
			}

			newQuantity := 1
			if q > 1 {
				newQuantity = simtypes.RandIntBetween(r, 1, q)
			}
			updatedOrders[i] = &marketplace.MsgUpdateSellOrders_Update{
				SellOrderId: orders[i].Id,
				NewQuantity: func() string {
					// 30% chance of new quantity set to 0
					if r.Int63n(101) <= 30 {
						return "0"
					}
					return fmt.Sprintf("%d", newQuantity)
				}(),
				NewAskPrice:       &askPrice,
				DisableAutoRetire: r.Int63n(101) <= 30, // 30% chance of disable auto-retire
				NewExpiration:     &exp,
			}
		}

		msg := &marketplace.MsgUpdateSellOrders{
			Seller:  sellerAddr,
			Updates: updatedOrders,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, sellerAddr, TypeMsgUpdateSellOrder)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simapp.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgCancelSellOrder generates a Marketplace/MsgCancelSellOrder with random values.
func SimulateMsgCancelSellOrder(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryServer, mktQryClient marketplace.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		seller, _ := simtypes.RandomAcc(r, accs)
		sellerAddr := seller.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		result, err := mktQryClient.SellOrdersBySeller(ctx, &marketplace.QuerySellOrdersBySellerRequest{Seller: sellerAddr})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancelSellOrder, err.Error()), nil, err
		}

		orders := result.SellOrders
		if len(orders) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancelSellOrder, "no sell orders present"), nil, nil
		}

		// select random order
		order := orders[r.Intn(len(orders))]
		msg := &marketplace.MsgCancelSellOrder{
			Seller:      sellerAddr,
			SellOrderId: order.Id,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, sellerAddr, TypeMsgCancelSellOrder)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simapp.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
