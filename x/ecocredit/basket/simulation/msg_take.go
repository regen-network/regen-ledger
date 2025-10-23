package simulation

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/simulation/utils"
)

const OpWeightMsgTake = "op_weight_take_from_basket" //nolint:gosec

const WeightTake = 100

var TypeMsgTake = sdk.MsgTypeURL(&types.MsgTake{})

// SimulateMsgTake generates a Basket/MsgTake with random values.
func SimulateMsgTake(txCfg client.TxConfig, ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	_ basetypes.QueryServer, bsktQryClient types.QueryServer,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		owner, _ := simtypes.RandomAcc(r, accs)
		ownerAddr := owner.Address.String()

		res, err := bsktQryClient.Baskets(ctx, &types.QueryBasketsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, err
		}

		baskets := res.BasketsInfo
		if len(baskets) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "no baskets"), nil, nil
		}

		var rBasket *types.BasketInfo
		var bBalances []*types.BasketBalanceInfo
		for _, b := range baskets {
			balancesRes, err := bsktQryClient.BasketBalances(ctx, &types.QueryBasketBalancesRequest{
				BasketDenom: b.BasketDenom,
			})
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, err
			}
			balances := balancesRes.BalancesInfo
			if len(balances) != 0 {
				rBasket = b
				bBalances = balances
				break
			}
		}
		if rBasket == nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "no basket"), nil, nil
		}

		var amt int
		for _, b := range bBalances {
			iAmount, err := strconv.Atoi(b.Balance)
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, err.Error()), nil, nil
			}

			switch { //nolint:staticcheck
			case iAmount == 0:
				continue
			case iAmount == 1:
				amt = iAmount
			default:
				amt = simtypes.RandIntBetween(r, 1, iAmount)
			}
		}
		if amt == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgTake, "basket balance"), nil, nil
		}

		msg := &types.MsgTake{
			Owner:                  ownerAddr,
			BasketDenom:            rBasket.BasketDenom,
			Amount:                 fmt.Sprintf("%d", amt),
			RetirementJurisdiction: "AQ",
			RetireOnTake:           !rBasket.DisableAutoRetire,
		}

		spendable := bk.SpendableCoins(ctx, owner.Address)
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txCfg,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      owner,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCfg, txCtx)
	}
}
