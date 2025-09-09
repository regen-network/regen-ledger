package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/simulation/utils"
)

const OpWeightMsgUpdateProjectAdmin = "op_weight_msg_update_project_admin" //nolint:gosec

var TypeMsgUpdateProjectAdmin = sdk.MsgTypeURL(&types.MsgUpdateProjectAdmin{})

const WeightUpdateProjectAdmin = 30

// SimulateMsgUpdateProjectAdmin generates a MsgUpdateProjectAdmin with random values.
func SimulateMsgUpdateProjectAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(ctx, r, qryClient, TypeMsgUpdateProjectAdmin)
		if err != nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgUpdateProjectAdmin, class.Id)
		if project == nil {
			return op, nil, err
		}

		newAdmin, _ := simtypes.RandomAcc(r, accs)
		if project.Admin == newAdmin.Address.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateProjectAdmin, "old and new admin are same"), nil, nil
		}

		msg := &types.MsgUpdateProjectAdmin{
			Admin:     project.Admin,
			NewAdmin:  newAdmin.Address.String(),
			ProjectId: project.Id,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(ctx, bk, accs, project.Admin, TypeMsgUpdateProjectAdmin)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
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
