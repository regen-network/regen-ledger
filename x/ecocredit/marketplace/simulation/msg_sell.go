package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/simulation/utils"
)

const OpWeightMsgSell = "op_weight_msg_sell" //nolint:gosec

const WeightSell = 100

var TypeMsgSell = types.MsgSell{}.Route()

// SimulateMsgSell generates a Marketplace/MsgSell with random values.
func SimulateMsgSell(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient basetypes.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, baseApp *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		seller, _ := simtypes.RandomAcc(r, accs)
		sellerAddr := seller.Address.String()

		class, op, err := utils.GetRandomClass(ctx, r, qryClient, TypeMsgSell)
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
		maxVal := 1
		if len(batches) > 1 {
			maxVal = simtypes.RandIntBetween(r, 1, len(batches))
		}

		sellOrders := make([]*types.MsgSell_Order, maxVal)
		for i := 0; i < maxVal; i++ {
			bal, err := qryClient.Balance(ctx, &basetypes.QueryBalanceRequest{Address: sellerAddr, BatchDenom: batches[i].Denom})
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSell, err.Error()), nil, err
			}

			exp := ctx.BlockTime().AddDate(1, 0, 0)
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

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(ctx, bk, accs, sellerAddr, TypeMsgSell)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             baseApp,
			TxGen:           testutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
