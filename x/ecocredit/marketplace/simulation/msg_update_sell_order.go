package simulation

import (
	"fmt"
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

const OpWeightMsgUpdateSellOrder = "op_weight_msg_update_sell_order" //nolint:gosec

const WeightUpdateSellOrder = 100

var TypeMsgUpdateSellOrder = types.MsgUpdateSellOrders{}.Route()

// SimulateMsgUpdateSellOrder generates a Marketplace/MsgUpdateSellOrder with random values.
func SimulateMsgUpdateSellOrder(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	_ basetypes.QueryServer, mktQryClient types.QueryServer) simtypes.Operation {
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
					// 30% chance of new quantity set to decimal
					if r.Int63n(101) <= 30 {
						return "0.123"
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
