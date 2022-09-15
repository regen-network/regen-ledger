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

const OpWeightMsgMintBatchCredits = "op_weight_msg_mint_batch_credits" //nolint:gosec

var TypeMsgMintBatchCredits = sdk.MsgTypeURL(&types.MsgMintBatchCredits{})

const WeightMintBatchCredits = 33

// SimulateMsgMintBatchCredits generates a MsgMintBatchCredits with random values.
func SimulateMsgMintBatchCredits(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuerAcc, _ := simtypes.RandomAcc(r, accs)
		issuerAddr := issuerAcc.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgMintBatchCredits)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgMintBatchCredits, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgMintBatchCredits, project.Id)
		if batch == nil {
			return op, nil, err
		}

		if !batch.Open {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgMintBatchCredits, "batch is closed"), nil, nil
		}

		if batch.Issuer != issuerAddr {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgMintBatchCredits, "only batch issuer can mint additional credits"), nil, nil
		}

		msg := &types.MsgMintBatchCredits{
			Issuer:     issuerAddr,
			BatchDenom: batch.Denom,
			Issuance:   generateBatchIssuance(r, accs),
			OriginTx: &types.OriginTx{
				Source: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 64)),
				Id:     simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 64)),
			},
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, issuerAddr, TypeMsgUpdateClassIssuers)
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
