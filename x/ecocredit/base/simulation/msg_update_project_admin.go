package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

const OpWeightMsgUpdateProjectAdmin = "op_weight_msg_update_project_admin" //nolint:gosec

var TypeMsgUpdateProjectAdmin = sdk.MsgTypeURL(&types.MsgUpdateProjectAdmin{})

const WeightUpdateProjectAdmin = 30

// SimulateMsgUpdateProjectAdmin generates a MsgUpdateProjectAdmin with random values.
func SimulateMsgUpdateProjectAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateProjectAdmin)
		if err != nil {
			return op, nil, err
		}

		ctx := sdk.WrapSDKContext(sdkCtx)
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

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, project.Admin, TypeMsgUpdateProjectAdmin)
		if spendable == nil {
			return op, nil, err
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
