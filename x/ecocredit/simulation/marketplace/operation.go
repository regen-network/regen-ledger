package marketplace

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

// Simulation operation weights constants
const (
	OpWeightMsgAllowAskDenom   = "op_weight_msg_allow_sdk_denom"
	OpWeightMsgBuy             = "op_weight_msg_buy"
	OpWeightMsgSell            = "op_weight_msg_sell"
	OpWeightMsgUpdateSellOrder = "op_weight_msg_update_sell_order"
	OpWeightMsgCancelSellOrder = "op_weight_msg_cancel_sell_order"
)

// basket operations weights
const (
	WeightAllowAskDenom   = 100
	WeightBuy             = 100
	WeightSell            = 100
	WeightUpdateSellOrder = 100
	WeightCancelSellOrder = 100
)

// ecocredit message types
var (
	TypeMsgAllowAskDenom   = marketplace.MsgAllowAskDenom{}.Route()
	TypeMsgBuy             = marketplace.MsgBuy{}.Route()
	TypeMsgSell            = marketplace.MsgSell{}.Route()
	TypeMsgUpdateSellOrder = marketplace.MsgUpdateSellOrders{}.Route()
	TypeMsgCancelSellOrder = marketplace.MsgCancelSellOrder{}.Route()
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient, mktQryClient marketplace.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgAllowAskDenom   int
		weightMsgBuy             int
		weightMsgSell            int
		weightMsgUpdateSellOrder int
		weightMsgCancelSellOrder int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAllowAskDenom, &weightMsgAllowAskDenom, nil,
		func(_ *rand.Rand) {
			weightMsgAllowAskDenom = WeightAllowAskDenom
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBuy, &weightMsgBuy, nil,
		func(_ *rand.Rand) {
			weightMsgBuy = WeightBuy
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
			weightMsgAllowAskDenom,
			SimulateMsgAllowAskDenom(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgBuy,
			SimulateMsgBuy(ak, bk, qryClient),
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

// SimulateMsgAllowAskDenom generates a Marketplace/MsgAllowAskDenom with random values.
func SimulateMsgAllowAskDenom(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAllowAskDenom, "not yet implemented"), nil, nil
	}
}

// SimulateMsgBuy generates a Marketplace/MsgBuy with random values.
func SimulateMsgBuy(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBuy, ""), nil, nil

	}
}

// SimulateMsgSell generates a Marketplace/MsgSell with random values.
func SimulateMsgSell(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, baseApp *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		owner, _ := simtypes.RandomAcc(r, accs)
		ownerAddr := owner.Address.String()

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgSell)
		if class == nil {
			return op, nil, err
		}

		batchRes, err := qryClient.BatchesByIssuer(ctx, &core.QueryBatchesByIssuerRequest{Issuer: owner.Address.String()})
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
			bal, err := qryClient.Balance(ctx, &core.QueryBalanceRequest{Account: ownerAddr, BatchDenom: batches[i].BatchDenom})
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, err.Error()), nil, err
			}

			exp := simtypes.RandTimestamp(r)
			d, err := math.NewNonNegativeDecFromString(bal.Balance.Tradable)
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, err.Error()), nil, err
			}

			if d.IsZero() {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, "no balance"), nil, nil
			}

			balInt, err := d.Int64()
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, err.Error()), nil, err
			}

			quantity := int(balInt)
			if balInt > 1 {
				quantity = simtypes.RandIntBetween(r, 1, int(quantity))
			}

			askPrice := sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simtypes.RandIntBetween(r, 1, 50)))
			sellOrders[i] = &marketplace.MsgSell_Order{
				BatchDenom:        batches[i].BatchDenom,
				Quantity:          fmt.Sprintf("%d", quantity),
				AskPrice:          &askPrice,
				DisableAutoRetire: r.Int63n(101) <= 30,
				Expiration:        &exp,
			}
		}

		msg := &marketplace.MsgSell{
			Owner:  owner.Address.String(),
			Orders: sellOrders,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, ownerAddr, TypeMsgSell)
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateSellOrder generates a Marketplace/MsgUpdateSellOrder with random values.
func SimulateMsgUpdateSellOrder(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	coreQryClient core.QueryClient, mktQryClient marketplace.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		owner, _ := simtypes.RandomAcc(r, accs)
		ownerAddr := owner.Address.String()

		ctx := regentypes.Context{Context: sdkCtx}
		result, err := mktQryClient.SellOrdersByAddress(ctx, &marketplace.QuerySellOrdersByAddressRequest{Address: ownerAddr})
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
			updatedOrders[i] = &marketplace.MsgUpdateSellOrders_Update{
				SellOrderId:       orders[i].Id,
				NewQuantity:       "1", //TODO: add new quantity
				NewAskPrice:       &askPrice,
				DisableAutoRetire: r.Int63n(101) <= 30,
				NewExpiration:     &exp,
			}
		}

		msg := &marketplace.MsgUpdateSellOrders{
			Owner:   ownerAddr,
			Updates: updatedOrders,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, ownerAddr, TypeMsgUpdateSellOrder)
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCancelSellOrder generates a Marketplace/MsgCancelSellOrder with random values.
func SimulateMsgCancelSellOrder(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient core.QueryClient, mktQryClient marketplace.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		owner, _ := simtypes.RandomAcc(r, accs)
		ownerAddr := owner.Address.String()

		ctx := regentypes.Context{Context: sdkCtx}
		result, err := mktQryClient.SellOrdersByAddress(ctx, &marketplace.QuerySellOrdersByAddressRequest{Address: ownerAddr})
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
			Seller:      ownerAddr,
			SellOrderId: order.Id,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, ownerAddr, TypeMsgCancelSellOrder)
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

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}
