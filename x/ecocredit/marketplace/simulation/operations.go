package simulation

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

// Simulation operation weights constants
const (
	OpWeightMsgBuy                = "op_weight_msg_buy_direct"           //nolint:gosec
	OpWeightMsgSell               = "op_weight_msg_sell"                 //nolint:gosec
	OpWeightMsgUpdateSellOrder    = "op_weight_msg_update_sell_order"    //nolint:gosec
	OpWeightMsgCancelSellOrder    = "op_weight_msg_cancel_sell_order"    //nolint:gosec
	OpWeightMsgAddAllowedDenom    = "op_weight_msg_add_allowed_denom"    //nolint:gosec
	OpWeightMsgRemoveAllowedDenom = "op_weight_msg_remove_allowed_denom" //nolint:gosec
)

// basket operations weights
const (
	WeightBuyDirect          = 100
	WeightSell               = 100
	WeightUpdateSellOrder    = 100
	WeightCancelSellOrder    = 100
	WeightAddAllowedDenom    = 100
	WeightRemoveAllowedDenom = 100
)

// ecocredit message types
var (
	TypeMsgBuyDirect          = types.MsgBuyDirect{}.Route()
	TypeMsgSell               = types.MsgSell{}.Route()
	TypeMsgUpdateSellOrder    = types.MsgUpdateSellOrders{}.Route()
	TypeMsgCancelSellOrder    = types.MsgCancelSellOrder{}.Route()
	TypeMsgAddAllowedDenom    = types.MsgAddAllowedDenom{}.Route()
	TypeMsgRemoveAllowedDenom = types.MsgRemoveAllowedDenom{}.Route()
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient basetypes.QueryServer, mktQryClient types.QueryServer,
	govk ecocredit.GovKeeper, authority sdk.AccAddress) simulation.WeightedOperations {

	var (
		weightMsgBuyDirect          int
		weightMsgSell               int
		weightMsgUpdateSellOrder    int
		weightMsgCancelSellOrder    int
		weightMsgAddAllowedDenom    int
		weightMsgRemoveAllowedDenom int
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

	appParams.GetOrGenerate(cdc, OpWeightMsgAddAllowedDenom, &weightMsgAddAllowedDenom, nil,
		func(_ *rand.Rand) {
			weightMsgAddAllowedDenom = WeightAddAllowedDenom
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveAllowedDenom, &weightMsgRemoveAllowedDenom, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveAllowedDenom = WeightRemoveAllowedDenom
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
		simulation.NewWeightedOperation(
			weightMsgAddAllowedDenom,
			SimulateMsgAddAllowedDenom(ak, bk, mktQryClient, govk, authority),
		),
		simulation.NewWeightedOperation(
			weightMsgRemoveAllowedDenom,
			SimulateMsgRemoveAllowedDenom(ak, bk, mktQryClient, govk, authority),
		),
	}
}

// SimulateMsgBuyDirect generates a Marketplace/MsgBuyDirect with random values.
func SimulateMsgBuyDirect(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient basetypes.QueryServer, mktQryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		buyer, _ := simtypes.RandomAcc(r, accs)
		buyerAddr := buyer.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		result, err := mktQryClient.SellOrders(ctx, &types.QuerySellOrdersRequest{})
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

		buyOrders := make([]*types.MsgBuyDirect_Order, max)
		for i := 0; i < max; i++ {
			sellOrderAskAmount, err := strconv.Atoi(sellOrders[i].AskAmount)
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBuyDirect, "could not convert to int"), nil, nil
			}

			bidPrice := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, sellOrderAskAmount, sellOrderAskAmount+100))))
			buyOrders[i] = &types.MsgBuyDirect_Order{
				SellOrderId:            sellOrders[i].Id,
				Quantity:               sellOrders[i].Quantity,
				BidPrice:               &bidPrice,
				DisableAutoRetire:      sellOrders[i].DisableAutoRetire,
				RetirementJurisdiction: "AQ",
			}
		}

		msg := &types.MsgBuyDirect{
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
	qryClient basetypes.QueryServer) simtypes.Operation {
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

		batchRes, err := qryClient.BatchesByIssuer(ctx, &basetypes.QueryBatchesByIssuerRequest{Issuer: seller.Address.String()})
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

		sellOrders := make([]*types.MsgSell_Order, max)
		for i := 0; i < max; i++ {
			bal, err := qryClient.Balance(ctx, &basetypes.QueryBalanceRequest{Address: sellerAddr, BatchDenom: batches[i].Denom})
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
			sellOrders[i] = &types.MsgSell_Order{
				BatchDenom:        batches[i].Denom,
				Quantity:          fmt.Sprintf("%d", quantity),
				AskPrice:          &askPrice,
				DisableAutoRetire: r.Int63n(101) <= 30, // 30% chance of disable auto-retire
				Expiration:        &exp,
			}
		}

		msg := &types.MsgSell{
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
	baseQryClient basetypes.QueryServer, mktQryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		seller, _ := simtypes.RandomAcc(r, accs)
		sellerAddr := seller.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		result, err := mktQryClient.SellOrdersBySeller(ctx, &types.QuerySellOrdersBySellerRequest{Seller: sellerAddr})
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

		updatedOrders := make([]*types.MsgUpdateSellOrders_Update, len(orders))
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
			updatedOrders[i] = &types.MsgUpdateSellOrders_Update{
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

		msg := &types.MsgUpdateSellOrders{
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
	qryClient basetypes.QueryServer, mktQryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		seller, _ := simtypes.RandomAcc(r, accs)
		sellerAddr := seller.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		result, err := mktQryClient.SellOrdersBySeller(ctx, &types.QuerySellOrdersBySellerRequest{Seller: sellerAddr})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancelSellOrder, err.Error()), nil, err
		}

		orders := result.SellOrders
		if len(orders) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancelSellOrder, "no sell orders present"), nil, nil
		}

		// select random order
		order := orders[r.Intn(len(orders))]
		msg := &types.MsgCancelSellOrder{
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

func SimulateMsgAddAllowedDenom(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer, govk ecocredit.GovKeeper, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgAddAllowedDenom)
		if spendable == nil {
			return op, nil, err
		}

		bankDenom := simtypes.RandStringOfLength(r, 4)
		res, err := qryClient.AllowedDenoms(sdkCtx, &types.QueryAllowedDenomsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddAllowedDenom, err.Error()), nil, err
		}

		if isDenomExists(res.AllowedDenoms, bankDenom) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddAllowedDenom, fmt.Sprintf("denom %s already exists", bankDenom)), nil, nil
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, authority)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddAllowedDenom, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddAllowedDenom, "unable to generate deposit"), nil, err
		}

		msg := types.MsgAddAllowedDenom{
			Authority:    authority.String(),
			BankDenom:    bankDenom,
			DisplayDenom: bankDenom,
			Exponent:     6,
		}

		any, err := codectypes.NewAnyWithValue(&msg)
		if err != nil {
			return simtypes.NoOpMsg(TypeMsgAddAllowedDenom, TypeMsgAddAllowedDenom, err.Error()), nil, err
		}

		proposalMsg := govtypes.MsgSubmitProposal{
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
			Messages:       []*codectypes.Any{any},
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simapp.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &proposalMsg,
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

func isDenomExists(allowedDenom []*types.AllowedDenom, bankDenom string) bool {
	for _, denom := range allowedDenom {
		if denom.BankDenom == bankDenom {
			return true
		}
	}

	return false
}

func SimulateMsgRemoveAllowedDenom(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	mktClient types.QueryServer, govk ecocredit.GovKeeper, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		response, err := mktClient.AllowedDenoms(sdkCtx, &types.QueryAllowedDenomsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveAllowedDenom, err.Error()), nil, err
		}

		if len(response.AllowedDenoms) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveAllowedDenom, "no allowed denom present"), nil, nil
		}

		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgRemoveAllowedDenom)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, authority)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveAllowedDenom, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveAllowedDenom, "unable to generate deposit"), nil, err
		}

		msg := types.MsgRemoveAllowedDenom{
			Authority: authority.String(),
			Denom:     response.AllowedDenoms[r.Intn(len(response.AllowedDenoms))].BankDenom,
		}

		any, err := codectypes.NewAnyWithValue(&msg)
		if err != nil {
			return simtypes.NoOpMsg(TypeMsgAddAllowedDenom, TypeMsgAddAllowedDenom, err.Error()), nil, err
		}

		proposalMsg := govtypes.MsgSubmitProposal{
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
			Messages:       []*codectypes.Any{any},
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simapp.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             &proposalMsg,
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
