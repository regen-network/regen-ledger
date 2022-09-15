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

const OpWeightMsgCreateProject = "op_weight_msg_create_project" //nolint:gosec

var TypeMsgCreateProject = sdk.MsgTypeURL(&types.MsgCreateProject{})

const WeightCreateProject = 20

// SimulateMsgCreateProject generates a MsgCreateProject with random values.
func SimulateMsgCreateProject(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgCreateProject)
		if class == nil {
			return op, nil, err
		}

		issuers, op, err := getClassIssuers(sdkCtx, qryClient, class.Id, TypeMsgCreateProject)
		if len(issuers) == 0 {
			return op, nil, err
		}

		admin := issuers[r.Intn(len(issuers))]
		adminAddr, err := sdk.AccAddressFromBech32(admin)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, err.Error()), nil, err
		}

		adminAcc, found := simtypes.FindAccount(accs, adminAddr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, "not a simulation account"), nil, nil
		}

		spendable := bk.SpendableCoins(sdkCtx, adminAddr)

		msg := &types.MsgCreateProject{
			Admin:        admin,
			ClassId:      class.Id,
			Metadata:     simtypes.RandStringOfLength(r, 100),
			Jurisdiction: "AB-CDE FG1 345",
		}
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      adminAcc,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
