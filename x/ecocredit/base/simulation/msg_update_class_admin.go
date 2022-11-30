package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/simulation/utils"
)

const OpWeightMsgUpdateClassAdmin = "op_weight_msg_update_class_admin" //nolint:gosec

var TypeMsgUpdateClassAdmin = sdk.MsgTypeURL(&types.MsgUpdateClassAdmin{})

const WeightUpdateClass = 30

// SimulateMsgUpdateClassAdmin generates a MsgUpdateClassAdmin with random values
func SimulateMsgUpdateClassAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassAdmin)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassAdmin)
		if spendable == nil {
			return op, nil, err
		}

		newAdmin, _ := simtypes.RandomAcc(r, accs)
		if newAdmin.Address.String() == admin.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassAdmin, "old and new account is same"), nil, nil // skip
		}

		msg := &types.MsgUpdateClassAdmin{
			Admin:    admin.String(),
			ClassId:  class.Id,
			NewAdmin: newAdmin.Address.String(),
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
