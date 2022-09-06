package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

const OpWeightMsgRetire = "op_weight_msg_retire" //nolint:gosec

var TypeMsgRetire = sdk.MsgTypeURL(&types.MsgRetire{})

const WeightRetire = 80

// SimulateMsgRetire generates a MsgRetire with random values.
func SimulateMsgRetire(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgRetire)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgRetire, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgRetire, project.Id)
		if batch == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(project.Admin).String()
		balanceRes, err := qryClient.Balance(ctx, &types.QueryBalanceRequest{
			Address:    admin,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "balance is zero"), nil, nil
		}

		randSub := math.NewDecFromInt64(int64(simtypes.RandIntBetween(r, 1, 10)))
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin, TypeMsgRetire)
		if spendable == nil {
			return op, nil, err
		}

		if !spendable.IsAllPositive() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "insufficient funds"), nil, nil
		}

		if tradableBalance.Cmp(randSub) != 1 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "insufficient funds"), nil, nil
		}

		msg := &types.MsgRetire{
			Owner: account.Address.String(),
			Credits: []*types.Credits{
				{
					BatchDenom: batch.Denom,
					Amount:     randSub.String(),
				},
			},
			Jurisdiction: "ST-UVW XY Z12",
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
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
