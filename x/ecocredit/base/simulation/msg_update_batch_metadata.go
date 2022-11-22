package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

const OpWeightMsgUpdateBatchMetadata = "op_weight_msg_update_batch_metadata" //nolint:gosec

var TypeMsgUpdateBatchMetadata = sdk.MsgTypeURL(&types.MsgUpdateBatchMetadata{})

const WeightUpdateBatchMetadata = 30

// SimulateMsgUpdateBatchMetadata generates a MsgUpdateBatchMetadata with random values.
func SimulateMsgUpdateBatchMetadata(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateBatchMetadata)
		if err != nil {
			return op, nil, err
		}

		ctx := sdk.WrapSDKContext(sdkCtx)

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgUpdateBatchMetadata, class.Id)
		if err != nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgUpdateBatchMetadata, project.Id)
		if batch == nil {
			return op, nil, err
		}

		issuer, err := sdk.AccAddressFromBech32(batch.Issuer)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateBatchMetadata, err.Error()), nil, err
		}

		msg := &types.MsgUpdateBatchMetadata{
			Issuer:      issuer.String(),
			BatchDenom:  batch.Denom,
			NewMetadata: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, base.MaxMetadataLength)),
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, issuer.String(), TypeMsgUpdateBatchMetadata)
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
