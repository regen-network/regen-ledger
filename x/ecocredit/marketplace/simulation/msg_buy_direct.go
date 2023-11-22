package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/simulation/utils"
)

const OpWeightMsgBuy = "op_weight_msg_buy_direct" //nolint:gosec

const WeightBuyDirect = 100

var TypeMsgBuyDirect = types.MsgBuyDirect{}.Route()

// SimulateMsgBuyDirect generates a Marketplace/MsgBuyDirect with random values.
func SimulateMsgBuyDirect(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	_ basetypes.QueryServer, mktQryClient types.QueryServer) simtypes.Operation {
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

		var buyOrders []*types.MsgBuyDirect_Order
		for i := 0; i < max; i++ {
			if sellOrders[i].Seller == buyerAddr {
				continue
			}

			sellOrderAskAmount, err := strconv.Atoi(sellOrders[i].AskAmount)
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBuyDirect, "could not convert to int"), nil, nil
			}

			bidPrice := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, sellOrderAskAmount, sellOrderAskAmount+100))))
			buyOrders = append(buyOrders, &types.MsgBuyDirect_Order{
				SellOrderId:            sellOrders[i].Id,
				Quantity:               sellOrders[i].Quantity,
				BidPrice:               &bidPrice,
				DisableAutoRetire:      sellOrders[i].DisableAutoRetire,
				RetirementJurisdiction: "AQ",
			})
		}

		if len(buyOrders) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBuyDirect, "empty buy orders"), nil, nil
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
