package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/data"
)

// Simulation operation weights constants
const (
	OpWeightMsgAnchor           = "op_weight_msg_anchor"
	OpWeightMsgAttest           = "op_weight_msg_attest"
	OpWeightMsgDefineResolver   = "op_weight_msg_define_resolver"
	OpWeightMsgRegisterResolver = "op_weight_msg_register_resolver"

	ModuleName = "data"
)

const (
	TypeMsgAnchor           = data.MsgAnchor{}.Route()
	TypeMsgAttest           = data.MsgAttest{}.Route()
	TypeMsgDefineResolver   = data.MsgDefineResolver{}.Route()
	TypeMsgRegisterResolver = data.MsgRegisterResolver{}.Route()
)

const (
	WeightAnchor           = 100
	WeightAttest           = 100
	WeightRegisterResolver = 100
	WeightDefineResolver   = 100
)

// WeightedOperations returns all the operations from the data module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgAnchor           int
		weightMsgAttest           int
		weightMsgDefineResolver   int
		weightMsgRegisterResolver int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAnchor, &weightMsgAnchor, nil,
		func(_ *rand.Rand) {
			weightMsgAnchor = WeightAnchor
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAttest, &weightMsgAttest, nil,
		func(_ *rand.Rand) {
			weightMsgAttest = WeightAttest
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgDefineResolver, &weightMsgDefineResolver, nil,
		func(_ *rand.Rand) {
			weightMsgDefineResolver = WeightDefineResolver
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRegisterResolver, &weightMsgRegisterResolver, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterResolver = WeightRegisterResolver
		},
	)

	ops := simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk, qryClient),
		),
	}

	return ops
}

// SimulateMsgAnchor generates a MsgAnchor with random values.
func SimulateMsgAnchor(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(ModuleName, TypeMsgAnchor, ""), nil, nil
	}
}

// SimulateMsgAttest generates a MsgAttest with random values.
func SimulateMsgAttest(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(ModuleName, TypeMsgAttest, ""), nil, nil
	}
}

// SimulateMsgDefineResolver generates a MsgDefineResolver with random values.
func SimulateMsgDefineResolver(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(ModuleName, TypeMsgDefineResolver, ""), nil, nil
	}
}

// SimulateMsgRegisterResolver generates a MsgRegisterResolver with random values.
func SimulateMsgRegisterResolver(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		return simtypes.NoOpMsg(ModuleName, TypeMsgRegisterResolver, ""), nil, nil
	}
}
