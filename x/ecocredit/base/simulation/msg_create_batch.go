package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

const OpWeightMsgCreateBatch = "op_weight_msg_create_batch" //nolint:gosec

var TypeMsgCreateBatch = sdk.MsgTypeURL(&types.MsgCreateBatch{})

const WeightCreateBatch = 50

// SimulateMsgCreateBatch generates a MsgCreateBatch with random values.
func SimulateMsgCreateBatch(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer, _ := simtypes.RandomAcc(r, accs)

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgCreateBatch)
		if class == nil {
			return op, nil, err
		}

		result, err := qryClient.ClassIssuers(ctx, &types.QueryClassIssuersRequest{ClassId: class.Id})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, err.Error()), nil, err
		}

		classIssuers := result.Issuers
		if len(classIssuers) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "no issuers"), nil, nil
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgCreateBatch, class.Id)
		if project == nil {
			return op, nil, err
		}

		if !utils.Contains(classIssuers, issuer.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "don't have permission to create credit batch"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		now := sdkCtx.BlockTime()
		tenHours := now.Add(10 * time.Hour)

		msg := &types.MsgCreateBatch{
			Issuer:    issuer.Address.String(),
			ProjectId: project.Id,
			Issuance:  generateBatchIssuance(r, accs),
			StartDate: &now,
			EndDate:   &tenHours,
			Metadata:  simtypes.RandStringOfLength(r, 10),
			Open:      r.Float32() < 0.3, // 30% chance of credit batch being dynamic batch
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      issuer,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
