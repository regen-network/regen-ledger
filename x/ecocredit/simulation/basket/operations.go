package basketsims

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreate = "op_weight_msg_create_basket"
	OpWeightMsgPut    = "op_weight_msg_put_into_basket"
	OpWeightMsgTake   = "op_weight_take_from_basket"
)

// basket operations weights
const (
	WeightCreate = 100
	WeightPut    = 100
	WeightTake   = 100
)

// ecocredit message types
var (
	TypeMsgCreate = basket.MsgCreate{}.Route()
	TypeMsgPut    = basket.MsgPut{}.Route()
	TypeMsgTake   = basket.MsgTake{}.Route()
)

func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient ecocredit.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgCreate int
		weightMsgPut    int
		weightMsgTake   int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreate, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgCreate = WeightCreate
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgPut, &weightMsgPut, nil,
		func(_ *rand.Rand) {
			weightMsgPut = WeightPut
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgTake, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgTake = WeightTake
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreate,
			SimulateMsgCreate(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgPut,
			SimulateMsgPut(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgTake,
			SimulateMsgTake(ak, bk, qryClient),
		),
	}
}

// SimulateMsgCreate generates a Basket/MsgCreate with random values.
func SimulateMsgCreate(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// TODO: implement

		return simtypes.NoOpMsg(TypeMsgCreate, TypeMsgCreate, ""), nil, nil
	}
}

// SimulateMsgPut generates a Basket/MsgPut with random values.
func SimulateMsgPut(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// TODO: implement

		return simtypes.NoOpMsg(TypeMsgPut, TypeMsgPut, ""), nil, nil
	}
}

// SimulateMsgTake generates a Basket/MsgTake with random values.
func SimulateMsgTake(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// TODO: implement

		return simtypes.NoOpMsg(TypeMsgTake, TypeMsgTake, ""), nil, nil
	}
}
