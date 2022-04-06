package marketplace

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
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
	qryClient ecocredit.QueryClient, basketQryClient basket.QueryClient) simulation.WeightedOperations {

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
			SimulateMsgUpdateSellOrder(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCancelSellOrder,
			SimulateMsgCancelSellOrder(ak, bk, qryClient),
		),
	}
}

// SimulateMsgAllowAskDenom generates a Marketplace/MsgAllowAskDenom with random values.
func SimulateMsgAllowAskDenom(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAllowAskDenom, ""), nil, nil
	}
}

// SimulateMsgBuy generates a Marketplace/MsgBuy with random values.
func SimulateMsgBuy(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBuy, ""), nil, nil

	}
}

// SimulateMsgSell generates a Marketplace/MsgSell with random values.
func SimulateMsgSell(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, ""), nil, nil
	}
}

// SimulateMsgUpdateSellOrder generates a Marketplace/MsgUpdateSellOrder with random values.
func SimulateMsgUpdateSellOrder(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateSellOrder, ""), nil, nil
	}
}

// SimulateMsgCancelSellOrder generates a Marketplace/MsgCancelSellOrder with random values.
func SimulateMsgCancelSellOrder(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancelSellOrder, ""), nil, nil
	}
}
