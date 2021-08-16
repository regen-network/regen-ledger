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
	ak exported.AccountKeeper, bk exported.BankKeeper) simulation.WeightedOperations {

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
			SimulateMsgCreateClass(ak, bk),
		),
	}
}

// SimulateMsgCreateClass generates a MsgCreateClass with random values.
func SimulateMsgCreateClass(ak exported.AccountKeeper, bk exported.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		designer, _ := simtypes.RandomAcc(r, accs)
		issuers := []string{accs[0].Address.String(), accs[1].Address.String()}

		designerAcc := ak.GetAccount(ctx, designer.Address)
		spendableCoins := bk.SpendableCoins(ctx, designer.Address)
		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
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
