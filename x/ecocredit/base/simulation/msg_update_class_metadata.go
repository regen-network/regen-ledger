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

const OpWeightMsgUpdateClassMetadata = "op_weight_msg_update_class_metadata" //nolint:gosec

var TypeMsgUpdateClassMetadata = sdk.MsgTypeURL(&types.MsgUpdateClassMetadata{})

const WeightUpdateClassMetadata = 33

// SimulateMsgUpdateClassMetadata generates a MsgUpdateClassMetadata with random metadata
func SimulateMsgUpdateClassMetadata(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassMetadata)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassMetadata)
		if spendable == nil {
			return op, nil, err
		}

		msg := &types.MsgUpdateClassMetadata{
			Admin:       admin.String(),
			ClassId:     class.Id,
			NewMetadata: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 256)),
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
