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

const OpWeightMsgCreateClass = "op_weight_msg_create_class" //nolint:gosec

var TypeMsgCreateClass = sdk.MsgTypeURL(&types.MsgCreateClass{})

const WeightCreateClass = 10

// SimulateMsgCreateClass generates a MsgCreateClass with random values.
func SimulateMsgCreateClass(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		admin, _ := simtypes.RandomAcc(r, accs)
		issuers := randomIssuers(r, accs)

		ctx := sdk.WrapSDKContext(sdkCtx)
		res, err := qryClient.Params(ctx, &types.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
		}

		params := res.Params
		if params.AllowlistEnabled && !utils.Contains(params.AllowedClassCreators, admin.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not allowed to create credit class"), nil, nil // skip
		}

		spendable, neg := bk.SpendableCoins(sdkCtx, admin.Address).SafeSub(params.CreditClassFee...)
		if neg {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not enough balance"), nil, nil
		}

		creditTypes := []string{"C", "BIO"}
		msg := &types.MsgCreateClass{
			Admin:            admin.Address.String(),
			Issuers:          issuers,
			Metadata:         simtypes.RandStringOfLength(r, 10),
			CreditTypeAbbrev: creditTypes[r.Intn(len(creditTypes))],
			Fee:              &params.CreditClassFee[0],
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      admin,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}
