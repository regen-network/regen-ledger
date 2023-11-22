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

const OpWeightMsgSealBatch = "op_weight_msg_seal_batch" //nolint:gosec

var TypeMsgSealBatch = sdk.MsgTypeURL(&types.MsgSealBatch{})

const WeightSealBatch = 33

// SimulateMsgSealBatch generates a MsgSealBatch with random values.
func SimulateMsgSealBatch(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuerAcc, _ := simtypes.RandomAcc(r, accs)
		issuerAddr := issuerAcc.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgSealBatch)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgSealBatch, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgSealBatch, project.Id)
		if batch == nil {
			return op, nil, err
		}

		if batch.Issuer != issuerAddr {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSealBatch, "only batch issuer can seal batch"), nil, nil
		}

		if !batch.Open {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSealBatch, "batch is closed"), nil, nil
		}

		msg := &types.MsgSealBatch{
			Issuer:     issuerAddr,
			BatchDenom: batch.Denom,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, issuerAddr, TypeMsgSealBatch)
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
