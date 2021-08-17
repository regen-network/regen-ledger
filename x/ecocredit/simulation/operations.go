package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/exported"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	regentypes "github.com/regen-network/regen-ledger/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateClass = "op_weight_msg_create_class"
)

// ecocredit operations weights
const (
	WeightCreateClass = 100
)

// ecocredit message types
var (
	TypeMsgCreateClass = sdk.MsgTypeURL(&ecocredit.MsgCreateClass{})
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak exported.AccountKeeper, bk exported.BankKeeper, qryClient ecocredit.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgCreateClass int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = WeightCreateClass
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateClass,
			SimulateMsgCreateClass(ak, bk, qryClient),
		),
	}
}

// SimulateMsgCreateClass generates a MsgCreateClass with random values.
func SimulateMsgCreateClass(ak exported.AccountKeeper, bk exported.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		designer, _ := simtypes.RandomAcc(r, accs)
		issuers := []string{accs[0].Address.String(), accs[1].Address.String()}

		designerAcc := ak.GetAccount(sdkCtx, designer.Address)

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Params(ctx, &ecocredit.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
		}

		params := res.Params
		if params.AllowlistEnabled && !contains(params.AllowedClassDesigners, designer.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not allowed to create credit class"), nil, nil // skip
		}

		spendableCoins := bk.SpendableCoins(sdkCtx, designer.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
		}

		if spendableCoins.IsAllLTE(params.CreditClassFee) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not enough balance"), nil, nil
		}

		msg := &ecocredit.MsgCreateClass{
			Designer:   designer.Address.String(),
			Issuers:    issuers,
			Metadata:   []byte(simtypes.RandStringOfLength(r, 10)),
			CreditType: "carbon",
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{designerAcc.GetAccountNumber()},
			[]uint64{designerAcc.GetSequence()},
			designer.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, err
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
